package infoblox

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func resourceIPAlloc() *schema.Resource {
	return &schema.Resource{
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
			"cidr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The address in cidr format.",
			},
			"ip_addr": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "IP address of cloud instance." +
					"Set a valid IP for static allocation and leave empty if dynamically allocated.",
				Computed: true,
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
		},
	}
}

func resourceIPAllocRequest(d *schema.ResourceData, m interface{}, isIPv6 bool) error {
	networkView := d.Get("network_view").(string)
	dnsView := d.Get("dns_view").(string)
	enableDns := d.Get("enable_dns").(bool)
	fqdn := d.Get("fqdn").(string)

	cidr := d.Get("cidr").(string)
	ipAddr := d.Get("ip_addr").(string)
	if ipAddr == "" && cidr == "" {
		return fmt.Errorf("'ipAddr' or 'cidr' mandatory for allocation through Host Address Record creation")
	}

	var ttl uint32
	useTtl := false
	tempVal := d.Get("ttl")
	tempTTL := tempVal.(int)
	if tempTTL >= 0 {
		useTtl = true
		ttl = uint32(tempTTL)
	} else if tempTTL != ttlUndef {
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

	// enableDns and enableDhcp flags used to create host record with respective flags.
	// By default enableDns is true.
	if isIPv6 {
		hostRec, err := objMgr.CreateHostRecord(
			enableDns,
			false,
			fqdn,
			networkView,
			dnsView,
			"", cidr,
			"", ipAddr,
			"", "",
			useTtl, ttl,
			comment,
			extAttrs, []string{})
		if err != nil {
			return fmt.Errorf(
				"error while creating a host record: %s", err.Error())
		}
		if err = d.Set("ip_addr", hostRec.Ipv6Addrs[0].Ipv6Addr); err != nil {
			return err
		}
		d.SetId(hostRec.Ref)
	} else {
		hostRec, err := objMgr.CreateHostRecord(
			enableDns,
			false,
			fqdn,
			networkView,
			dnsView,
			cidr, "",
			ipAddr, "",
			"", "",
			useTtl, ttl,
			comment,
			extAttrs, []string{})
		if err != nil {
			return fmt.Errorf(
				"error while creating a host record: %s", err.Error())
		}
		if err = d.Set("ip_addr", hostRec.Ipv4Addrs[0].Ipv4Addr); err != nil {
			return err
		}
		d.SetId(hostRec.Ref)
	}
	return nil
}

func resourceIPAllocGet(d *schema.ResourceData, m interface{}, isIPv6 bool) error {

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

	hostRec, err := objMgr.GetHostRecordByRef(d.Id())
	if err != nil {
		return fmt.Errorf("error while getting the host record with ID '%s': %s", d.Id(), err.Error())
	}
	d.SetId(hostRec.Ref)

	return nil
}

func resourceIPAllocUpdate(d *schema.ResourceData, m interface{}, isIPv6 bool) error {
	networkView := d.Get("network_view").(string)
	if d.HasChange("network_view") {
		return fmt.Errorf("changing the value of 'network_view' field is not allowed")
	}
	if d.HasChange("dns_view") {
		return fmt.Errorf("changing the value of 'dns_view' field is not allowed")
	}
	enableDns := d.Get("enable_dns").(bool)
	fqdn := d.Get("fqdn").(string)

	cidr := d.Get("cidr").(string)
	ipAddr := d.Get("ip_addr").(string)
	// If 'cidr' is unchanged, then nothing to update here, making them empty to skip the update.
	// (This is to prevent record renewal for the case when 'cidr' is
	// used for IP address allocation, otherwise the address will be changing
	// during every 'update' operation).
	if !d.HasChange("cidr") {
		cidr = ""
	}

	var ttl uint32
	useTtl := false
	tempVal := d.Get("ttl")
	tempTTL := tempVal.(int)
	if tempTTL >= 0 {
		useTtl = true
		ttl = uint32(tempTTL)
	} else if tempTTL != ttlUndef {
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

	var (
		recIpV4Addr   *ibclient.HostRecordIpv4Addr
		recIpV6Addr   *ibclient.HostRecordIpv6Addr
		macAddr, duid string
	)

	// Retrive the IP of Host or Fixed Address record.
	// When IP is allocated using cidr and an empty IP is passed for updation
	hostRecObj, err := objMgr.GetHostRecordByRef(d.Id())
	if err != nil {
		return fmt.Errorf("error while getting the host record with ID '%s': %s", d.Id(), err.Error())
	}
	if len(hostRecObj.Ipv4Addrs) > 0 {
		recIpV4Addr = &hostRecObj.Ipv4Addrs[0]
	}
	if len(hostRecObj.Ipv6Addrs) > 0 {
		recIpV6Addr = &hostRecObj.Ipv6Addrs[0]
	}

	enableDhcp := false
	if recIpV4Addr != nil {
		macAddr = recIpV4Addr.Mac
		enableDhcp = recIpV4Addr.EnableDhcp
	} else if recIpV6Addr != nil {
		duid = recIpV6Addr.Duid
		enableDhcp = recIpV6Addr.EnableDhcp
	}

	if cidr == "" && ipAddr == "" {
		if isIPv6 {
			ipAddr = recIpV6Addr.Ipv6Addr
		} else {
			ipAddr = recIpV4Addr.Ipv4Addr
		}
	}

	if isIPv6 {
		hostRec, err := objMgr.UpdateHostRecord(
			d.Id(),
			enableDns,
			enableDhcp,
			fqdn,
			networkView,
			"", cidr,
			"", ipAddr,
			"", duid,
			useTtl, ttl,
			comment,
			extAttrs, []string{})
		if err != nil {
			return fmt.Errorf(
				"error updating IPv6 Host record with ID %s: %s", d.Id(), err.Error())
		}
		d.SetId(hostRec.Ref)
		if err = d.Set("ip_addr", hostRec.Ipv6Addrs[0].Ipv6Addr); err != nil {
			return err
		}
	} else {
		hostRec, err := objMgr.UpdateHostRecord(
			d.Id(),
			enableDns,
			enableDhcp,
			fqdn,
			networkView,
			cidr, "",
			ipAddr, "",
			macAddr, "",
			useTtl, ttl,
			comment,
			extAttrs, []string{})
		if err != nil {
			return fmt.Errorf(
				"error while updating IPv4 host record with ID '%s': %s", d.Id(), err.Error())
		}
		d.SetId(hostRec.Ref)
		if err = d.Set("ip_addr", hostRec.Ipv4Addrs[0].Ipv4Addr); err != nil {
			return err
		}
	}
	return nil
}

func resourceIPAllocRelease(d *schema.ResourceData, m interface{}, isIPv6 bool) error {
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

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	_, err := objMgr.DeleteHostRecord(d.Id())
	if err != nil {
		return fmt.Errorf("error releasing the IP address with ID '%s': %s", d.Id(), err.Error())
	}

	d.SetId("")
	return nil
}

// Code snippet for IPv4 IP Allocation
func resourceIPv4AllocationRequest(d *schema.ResourceData, m interface{}) error {
	return resourceIPAllocRequest(d, m, false)
}

func resourceIPv4AllocationGet(d *schema.ResourceData, m interface{}) error {
	return resourceIPAllocGet(d, m, false)
}

func resourceIPv4AllocationUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceIPAllocUpdate(d, m, false)
}

func resourceIPv4AllocationRelease(d *schema.ResourceData, m interface{}) error {
	return resourceIPAllocRelease(d, m, false)
}

func resourceIPv4Allocation() *schema.Resource {
	ipv4Allocation := resourceIPAlloc()
	ipv4Allocation.Create = resourceIPv4AllocationRequest
	ipv4Allocation.Read = resourceIPv4AllocationGet
	ipv4Allocation.Update = resourceIPv4AllocationUpdate
	ipv4Allocation.Delete = resourceIPv4AllocationRelease

	return ipv4Allocation
}

// Code snippet for IPv6 IP allocation
func resourceIPv6AllocationRequest(d *schema.ResourceData, m interface{}) error {
	return resourceIPAllocRequest(d, m, true)
}

func resourceIPv6AllocationGet(d *schema.ResourceData, m interface{}) error {
	return resourceIPAllocGet(d, m, true)
}

func resourceIPv6AllocationUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceIPAllocUpdate(d, m, true)
}

func resourceIPv6AllocationRelease(d *schema.ResourceData, m interface{}) error {
	return resourceIPAllocRelease(d, m, true)
}

func resourceIPv6Allocation() *schema.Resource {
	ipv6Allocation := resourceIPAlloc()
	ipv6Allocation.Create = resourceIPv6AllocationRequest
	ipv6Allocation.Read = resourceIPv6AllocationGet
	ipv6Allocation.Update = resourceIPv6AllocationUpdate
	ipv6Allocation.Delete = resourceIPv6AllocationRelease

	return ipv6Allocation
}
