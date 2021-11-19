package infoblox

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

// Code snipper for IP (IPv4 and IPv6) Allocation
func resourceIPAllocation() *schema.Resource {
	return &schema.Resource{
		Create: resourceAllocationRequest,
		Read:   resourceAllocationGet,
		Update: resourceAllocationUpdate,
		Delete: resourceAllocationRelease,
		Schema: map[string]*schema.Schema{
			"network_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
				Description: "Network view name of NIOS server.",
			},
			"dns_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
				Description: "Dns View under which the zone has been created.",
			},
			"enable_dns": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "flag that defines if the host record is to be used for DNS Purposes.",
			},
			"ipv4_cidr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The IPv4 cidr from which an IPv4 address will be allocated.",
			},
			"ipv6_cidr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The IPv6 cidr from which an IPv6 address will be allocated.",
			},
			"ipv4_addr": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
				Description: "IPv4 address of cloud instance." +
					"Set a valid IP address for static allocation and leave empty if dynamically allocated.",
			},
			"allocated_ipv4_addr": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Value which comes from 'ipv4_addr' (if specified) or from auto-allocation function (using 'ipv4_cidr').",
			},
			"ipv6_addr": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
				Description: "IPv6 address of cloud instance." +
					"Set a valid IP address for static allocation and leave empty if dynamically allocated.",
			},
			"allocated_ipv6_addr": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Value which comes from 'ipv6_addr' (if specified) or from auto-allocation function (using 'ipv6_cidr').",
			},
			"fqdn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The host name for Host Record in FQDN format.",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     ttlUndef,
				Description: "TTL attribute value for the record.",
			},
			"comment": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "A description of IP address allocation.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "The extensible attributes for IP address allocation, as a map in JSON format",
			},
			"internal_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "Internal ID of an object at NIOS side," +
					" used by Infoblox Terraform plugin to search for a NIOS's object" +
					" which corresponds to the Terraform resource.",
			},
		},
	}
}

func getAndRenewHostRecAltId(
	d *schema.ResourceData, m interface{}) (

	hostRec *ibclient.HostRecord, err error) {

	internalId, ref, err := getAltIdFields(d.Id())
	if err != nil {
		return nil, err
	}

	objMgr := ibclient.NewObjectManager(m.(ibclient.IBConnector), "Terraform", "")
	hostRec, err = objMgr.SearchHostRecordByAltId(internalId, ref, eaNameForInternalId)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return nil, fmt.Errorf(
				"error getting the allocated host record with ID '%s': %s",
				d.Id(), err.Error())
		}
		log.Printf("resource with the ID '%s' has been lost, removing it", d.Id())
		d.SetId("")
		return nil, nil
	}
	if ref != hostRec.Ref {
		newAltId, err := renewAltId(d.Id(), hostRec.Ref)
		if err != nil {
			return nil, nil
		}
		d.SetId(newAltId)
	}

	return
}

func resourceAllocationRequest(d *schema.ResourceData, m interface{}) error {
	networkView := d.Get("network_view").(string)
	dnsView := d.Get("dns_view").(string)
	enableDns := d.Get("enable_dns").(bool)
	fqdn := d.Get("fqdn").(string)
	if intId := d.Get("internal_id"); intId.(string) != "" {
		return fmt.Errorf("the value of 'internal_id' field must not be set manually")
	}

	ipv4Cidr := d.Get("ipv4_cidr").(string)
	ipv6Cidr := d.Get("ipv6_cidr").(string)
	ipv4Addr := d.Get("ipv4_addr").(string)
	ipv6Addr := d.Get("ipv6_addr").(string)
	if ipv4Cidr == "" && ipv6Cidr == "" && ipv4Addr == "" && ipv6Addr == "" {
		return fmt.Errorf("allocation through host address record creation needs an IPv4/IPv6 address" +
			" or IPv4/IPv6 cidr")
	}

	ZeroMacAddr := "00:00:00:00:00:00"
	var macAddr string
	if ipv4Cidr != "" || ipv4Addr != "" {
		macAddr = ZeroMacAddr
	}

	var ttl uint32
	useTtl := false
	tempVal := d.Get("ttl")
	tempTTL := tempVal.(int)
	if tempTTL >= 0 {
		useTtl = true
		ttl = uint32(tempTTL)
	} else if tempTTL == ttlUndef {
		ttl = 0
		useTtl = false
	} else {
		return fmt.Errorf("TTL value must be 0 or higher")
	}

	comment := d.Get("comment").(string)
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
		}
	}

	var tenantID string
	if tempVal, ok := extAttrs["Tenant ID"]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	internalId := generateInternalId()
	extAttrs[eaNameForInternalId] = internalId

	// enableDns and enableDhcp flags used to create host record with respective flags.
	// By default enableDns is true.
	hostRec, err := objMgr.CreateHostRecord(
		enableDns,
		false,
		fqdn,
		networkView,
		dnsView,
		ipv4Cidr, ipv6Cidr,
		ipv4Addr, ipv6Addr,
		macAddr, "",
		useTtl, ttl,
		comment,
		extAttrs, []string{})
	if err != nil {
		return fmt.Errorf("error while creating a host record: %s", err.Error())
	}
	if hostRec.Ipv6Addrs == nil || len(hostRec.Ipv6Addrs) < 1 {
		d.Set("allocated_ipv6_addr", "")
	} else {
		d.Set("allocated_ipv6_addr", hostRec.Ipv6Addrs[0].Ipv6Addr)
	}
	if hostRec.Ipv4Addrs == nil || len(hostRec.Ipv4Addrs) < 1 {
		d.Set("allocated_ipv4_addr", "")
	} else {
		d.Set("allocated_ipv4_addr", hostRec.Ipv4Addrs[0].Ipv4Addr)
	}
	d.SetId(generateAltId(internalId, hostRec.Ref))
	d.Set("internal_id", internalId)

	return nil
}

func resourceAllocationGet(d *schema.ResourceData, m interface{}) error {
	hostRecObj, err := getAndRenewHostRecAltId(d, m)
	if err != nil || hostRecObj == nil {
		return err
	}

	if hostRecObj.Ipv6Addrs == nil || len(hostRecObj.Ipv6Addrs) < 1 {
		d.Set("allocated_ipv6_addr", "")
	} else {
		d.Set("allocated_ipv6_addr", hostRecObj.Ipv6Addrs[0].Ipv6Addr)
	}
	if hostRecObj.Ipv4Addrs == nil || len(hostRecObj.Ipv4Addrs) < 1 {
		d.Set("allocated_ipv4_addr", "")
	} else {
		d.Set("allocated_ipv4_addr", hostRecObj.Ipv4Addrs[0].Ipv4Addr)
	}

	internalId, _, err := getAltIdFields(d.Id())
	if err != nil {
		return err
	}

	d.Set("internal_id", internalId)

	return nil
}

func resourceAllocationUpdate(d *schema.ResourceData, m interface{}) error {
	var err error

	hostRecObj, err := getAndRenewHostRecAltId(d, m)
	if err != nil || hostRecObj == nil {
		return err
	}

	if d.HasChange("network_view") {
		return fmt.Errorf("changing the value of 'network_view' field is not allowed")
	}
	if d.HasChange("dns_view") {
		return fmt.Errorf("changing the value of 'dns_view' field is not allowed")
	}
	if d.HasChange("internal_id") {
		return fmt.Errorf("changing the value of 'internal_id' field is not allowed")
	}
	fqdn := d.Get("fqdn").(string)

	ipv4Cidr := d.Get("ipv4_cidr").(string)
	ipv6Cidr := d.Get("ipv6_cidr").(string)
	ipv4Addr := d.Get("ipv4_addr").(string)
	ipv6Addr := d.Get("ipv6_addr").(string)

	// If 'ipv4_cidr' or 'ipv6_cidr' are unchanged, then nothing to update here.
	// making them empty to skip dynamic allocation of a new IP address again.
	// (This is to prevent record renewal for the case when 'cidr' is used for IP address allocation,
	// otherwise the address will be changing during every 'update' operation).
	if !d.HasChange("ipv4_cidr") {
		ipv4Cidr = ""
	}
	if !d.HasChange("ipv6_cidr") {
		ipv6Cidr = ""
	}

	var ttl uint32
	useTtl := false
	tempVal := d.Get("ttl")
	tempTTL := tempVal.(int)
	if tempTTL >= 0 {
		useTtl = true
		ttl = uint32(tempTTL)
	} else if tempTTL == ttlUndef {
		tempTTL = 0
		useTtl = false
	} else {
		return fmt.Errorf("TTL value must be 0 or higher")
	}

	enableDns := d.Get("enable_dns").(bool)
	comment := d.Get("comment").(string)
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
		}
	}
	var tenantID string
	if tempVal, ok := extAttrs["Tenant ID"]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	// Retrive the IP of Host or Fixed Address record.
	// When IP is allocated using cidr and an empty IP is passed for updation
	needIpv4Addr := ipv4Cidr == "" && ipv4Addr == ""
	needIpv6Addr := ipv6Cidr == "" && ipv6Addr == ""
	var (
		macAddr, duid string
	)
	if needIpv4Addr || needIpv6Addr {
		if _, ipv4_cidrFlag := d.GetOk("ipv4_cidr"); ipv4_cidrFlag && len(hostRecObj.Ipv4Addrs) > 0 {
			ipv4Addr = hostRecObj.Ipv4Addrs[0].Ipv4Addr
			macAddr = hostRecObj.Ipv4Addrs[0].Mac
		}
		if _, ipv6_cidrFlag := d.GetOk("ipv6_cidr"); ipv6_cidrFlag && len(hostRecObj.Ipv6Addrs) > 0 {
			ipv6Addr = hostRecObj.Ipv6Addrs[0].Ipv6Addr
			duid = hostRecObj.Ipv6Addrs[0].Duid
		}
	}

	internalId, _, err := getAltIdFields(d.Id())
	if err != nil {
		return err
	}
	extAttrs[eaNameForInternalId] = internalId

	var (
		recIpV4Addr *ibclient.HostRecordIpv4Addr
		recIpV6Addr *ibclient.HostRecordIpv6Addr
	)
	if len(hostRecObj.Ipv4Addrs) > 0 {
		recIpV4Addr = &hostRecObj.Ipv4Addrs[0]
	}
	if len(hostRecObj.Ipv6Addrs) > 0 {
		recIpV6Addr = &hostRecObj.Ipv6Addrs[0]
	}

	enableDhcp := false
	if enableDhcpV4 := recIpV4Addr != nil && recIpV4Addr.EnableDhcp; enableDhcpV4 {
		macAddr = recIpV4Addr.Mac
		enableDhcp = true
	}
	if enableDhcpV6 := recIpV6Addr != nil && recIpV6Addr.EnableDhcp; enableDhcpV6 {
		duid = recIpV6Addr.Duid
		enableDhcp = true
	}

	hostRecObj, err = objMgr.UpdateHostRecord(
		hostRecObj.Ref,
		enableDns,
		enableDhcp,
		fqdn,
		hostRecObj.NetworkView,
		ipv4Cidr, ipv6Cidr,
		ipv4Addr, ipv6Addr,
		macAddr, duid,
		useTtl, ttl,
		comment,
		extAttrs, []string{})
	if err != nil {
		return fmt.Errorf(
			"error while updating IP addresses of the host record with ID '%s': %s", d.Id(), err.Error())
	}
	d.SetId(generateAltId(internalId, hostRecObj.Ref))

	if hostRecObj.Ipv6Addrs == nil || len(hostRecObj.Ipv6Addrs) < 1 {
		d.Set("allocated_ipv6_addr", "")
	} else {
		d.Set("allocated_ipv6_addr", hostRecObj.Ipv6Addrs[0].Ipv6Addr)
	}
	if hostRecObj.Ipv4Addrs == nil || len(hostRecObj.Ipv4Addrs) < 1 {
		d.Set("allocated_ipv4_addr", "")
	} else {
		d.Set("allocated_ipv4_addr", hostRecObj.Ipv4Addrs[0].Ipv4Addr)
	}

	d.Set("internal_id", internalId)

	return nil
}

func resourceAllocationRelease(d *schema.ResourceData, m interface{}) error {
	_, err := getAndRenewHostRecAltId(d, m)
	if err != nil {
		return err
	}

	if d.HasChange("network_view") {
		return fmt.Errorf("changing the value of 'network_view' field is not allowed")
	}
	if d.HasChange("dns_view") {
		return fmt.Errorf("changing the value of 'dns_view' field is not allowed")
	}
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
		}
	}
	var tenantID string
	if tempVal, ok := extAttrs["Tenant ID"]; ok {
		tenantID = tempVal.(string)
	}

	if d.Id() == "" {
		log.Printf("WARNING: getting an error while determining ID of the resource to be cleanned up (probably non-existent resource, continuing): ): %s", err)
		return nil
	}
	_, ref, err := getAltIdFields(d.Id())
	if err != nil {
		return err
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)
	_, err = objMgr.DeleteHostRecord(ref)
	if err != nil {
		return fmt.Errorf("error while releasing the IP address with ID '%s': %s", d.Id(), err.Error())
	}
	d.SetId("")
	return nil
}
