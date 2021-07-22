package infoblox

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func resourceIPAssociation() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"network_view": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
				Description: "Network view name of NIOS server.",
			},
			"dns_view": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
				Description: "view in which record has to be created.",
			},
			"enable_dns": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "flag that defines if the host record is to be used for DNS Purposes",
			},
			"enable_dhcp": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "flag that defines if the host record is to be used for IPAM Purposes.",
			},
			"cidr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The address in cidr format.",
			},
			"ip_addr": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "IP address of cloud instance.",
			},
			"mac_addr": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "mac address of cloud instance.",
			},
			"duid": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "DHCP unique identifier for IPv6.",
			},
			"fqdn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The host name for Host Record in FQDN format.",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "TTL attribute value for the record.",
			},
			"comment": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "A description of the IP association.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "The Extensible attributes for IP Association, as a map in JSON format",
			},
		},
	}
}

//This method has an update call for the reason that,we are creating
//a reservation which doesnt have the details of the mac address
//at the beginig and we are using this update call to update the mac address
//of the record after the VM has been provisined. It is in the create method
//because for this resource we are doing association instead of allocation.
func resourceIPAssociationCreate(d *schema.ResourceData, m interface{}, isIPv6 bool) error {

	if err := Resource(d, m, isIPv6); err != nil {
		return err
	}

	return nil
}

func resourceIPAssociationUpdate(d *schema.ResourceData, m interface{}, isIPv6 bool) error {

	if d.HasChange("network_view") {
		return fmt.Errorf("changing the value of 'networkView' field is not allowed")
	}
	if d.HasChange("dns_view") {
		return fmt.Errorf("changing the value of 'dns_view' field is not allowed")
	}

	if err := Resource(d, m, isIPv6); err != nil {
		return err
	}

	return nil
}

func resourceIPAssociationRead(d *schema.ResourceData, m interface{}) error {

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
		return fmt.Errorf("Error getting Allocated HostRecord with ID: %s failed : %s",
			d.Id(), err.Error())
	}
	d.SetId(hostRec.Ref)
	return nil
}

//we are updating the record with an empty mac address after the vm has been
//destroyed because if we implement the delete hostrecord method here then there
//will be a conflict of resources
func resourceIPAssociationDelete(d *schema.ResourceData, m interface{}, isIPv6 bool) error {

	networkView := d.Get("network_view").(string)
	if d.HasChange("network_view") {
		return fmt.Errorf("changing the value of 'networkView' field is not allowed")
	}
	if d.HasChange("dns_view") {
		return fmt.Errorf("changing the value of 'dns_view' field is not allowed")
	}
	enableDns := d.Get("enable_dns").(bool)
	enableDhcp := d.Get("enable_dhcp").(bool)
	fqdn := d.Get("fqdn").(string)
	cidr := d.Get("cidr").(string)
	ipAddr := d.Get("ip_addr").(string)
	duid := d.Get("duid").(string)

	var ttl uint32
	tempVal, useTtl := d.GetOk("ttl")
	if useTtl {
		tempTtl := tempVal.(int)
		if tempTtl < 0 {
			return fmt.Errorf("TTL value must be 0 or higher")
		}
		ttl = uint32(tempTtl)
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

	ZeroMacAddr := "00:00:00:00:00:00"
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

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
			return fmt.Errorf("Error updating Host record with ID %s: %s", d.Id(), err.Error())
		}
		d.SetId(hostRec.Ref)
	} else {
		hostRec, err := objMgr.UpdateHostRecord(
			d.Id(),
			enableDns,
			enableDhcp,
			fqdn,
			networkView,
			cidr, "",
			ipAddr, "",
			ZeroMacAddr, "",
			useTtl, ttl,
			comment,
			extAttrs, []string{})
		if err != nil {
			return fmt.Errorf("Error updating Host record with ID %s: %s", d.Id(), err.Error())
		}
		d.SetId(hostRec.Ref)
	}
	return nil
}

func Resource(d *schema.ResourceData, m interface{}, isIPv6 bool) error {

	networkView := d.Get("network_view").(string)
	dnsView := d.Get("dns_view").(string)
	enableDhcp := d.Get("enable_dhcp").(bool)
	enableDns := d.Get("enable_dns").(bool)
	// dnsView made empty so that searching of host record to be done at IPAM end
	if !enableDns {
		dnsView = ""
	}

	fqdn := d.Get("fqdn").(string)
	cidr := d.Get("cidr").(string)
	ipAddr := d.Get("ip_addr").(string)
	macAddr := d.Get("mac_addr").(string)
	//conversion from bit reversed EUI-48 format to hexadecimal EUI-48 format
	macAddr = strings.Replace(macAddr, "-", ":", -1)
	duid := d.Get("duid").(string)

	var ttl uint32
	tempVal, useTtl := d.GetOk("ttl")
	if useTtl {
		tempTtl := tempVal.(int)
		if tempTtl < 0 {
			return fmt.Errorf("TTL value must be 0 or higher")
		}
		ttl = uint32(tempTtl)
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

	if isIPv6 {
		hostRecordObj, err := objMgr.GetHostRecord(networkView, dnsView, fqdn, "", ipAddr)
		if err != nil {
			return fmt.Errorf("Failed to get HostRecord for 'fqdn': %s and 'IP':%s in"+
				"'network view': %s and 'dns view':%s. Error:%s",
				fqdn, ipAddr, networkView, dnsView, err.Error())
		}
		if hostRecordObj == nil {
			return fmt.Errorf("HostRecord %s not found.", fqdn)
		}
		hostRec, err := objMgr.UpdateHostRecord(
			hostRecordObj.Ref,
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
			return fmt.Errorf("UpdateHost Record failed with ID %s: %s", d.Id(), err.Error())
		}
		d.SetId(hostRec.Ref)
	} else {
		hostRecordObj, err := objMgr.GetHostRecord(networkView, dnsView, fqdn, ipAddr, "")
		if err != nil {
			return fmt.Errorf("Failed to get HostRecord for 'fqdn': %s and 'IP':%s in"+
				"'network view': %s and 'dns view':%s. Error:%s",
				fqdn, ipAddr, networkView, dnsView, err.Error())
		}
		if hostRecordObj == nil {
			return fmt.Errorf("HostRecord %s not found.", fqdn)
		}
		hostRec, err := objMgr.UpdateHostRecord(
			hostRecordObj.Ref,
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
			return fmt.Errorf("UpdateHost Record failed with ID %s: %s", d.Id(), err.Error())
		}
		d.SetId(hostRec.Ref)
	}
	return nil
}

// Code snippet for IPv4 IP Association
func resourceIPv4AssociationCreate(d *schema.ResourceData, m interface{}) error {
	return resourceIPAssociationCreate(d, m, false)
}

func resourceIPv4AssociationGet(d *schema.ResourceData, m interface{}) error {
	return resourceIPAssociationRead(d, m)
}

func resourceIPv4AssociationUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceIPAssociationUpdate(d, m, false)
}

func resourceIPv4AssociationDelete(d *schema.ResourceData, m interface{}) error {
	return resourceIPAssociationDelete(d, m, false)
}

func resourceIPv4Association() *schema.Resource {
	ipv4Association := resourceIPAssociation()
	ipv4Association.Create = resourceIPv4AssociationCreate
	ipv4Association.Read = resourceIPv4AssociationGet
	ipv4Association.Update = resourceIPv4AssociationUpdate
	ipv4Association.Delete = resourceIPv4AssociationDelete

	return ipv4Association
}

// Code snippet for IPv6 IP Association
func resourceIPv6AssociationCreate(d *schema.ResourceData, m interface{}) error {
	return resourceIPAssociationCreate(d, m, true)
}

func resourceIPv6AssociationRead(d *schema.ResourceData, m interface{}) error {
	return resourceIPAssociationRead(d, m)
}

func resourceIPv6AssociationUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceIPAssociationUpdate(d, m, true)
}

func resourceIPv6AssociationDelete(d *schema.ResourceData, m interface{}) error {
	return resourceIPAssociationDelete(d, m, true)
}

func resourceIPv6Association() *schema.Resource {
	ipv6Association := resourceIPAssociation()
	ipv6Association.Create = resourceIPv6AssociationCreate
	ipv6Association.Read = resourceIPv6AssociationRead
	ipv6Association.Update = resourceIPv6AssociationUpdate
	ipv6Association.Delete = resourceIPv6AssociationDelete

	return ipv6Association
}
