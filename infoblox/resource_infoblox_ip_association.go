package infoblox

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client"
)

func resourceIPAssociation() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"network_view_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
				Description: "Network view name of NIOS server.",
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
			"dns_view": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
				Description: "view in which record has to be created.",
			},
			"zone": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "zone under which record has been created.",
			},
			"cidr": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
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
			"host_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The host name",
			},
			"comment": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "A description of the IP association.",
			},
			"extensible_attributes": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "The Extensible attributes of the network container to be added/updated, as a map in JSON format",
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

	if err := Resource(d, m, isIPv6); err != nil {
		return err
	}

	return nil
}

func resourceIPAssociationRead(d *schema.ResourceData, m interface{}) error {

	enableDns := d.Get("enable_dns").(bool)
	enableDhcp := d.Get("enable_dhcp").(bool)
	hostName := d.Get("host_name").(string)

	cidr := d.Get("cidr").(string)
	dnsView := d.Get("dns_view").(string)
	zone := d.Get("zone").(string)
	extAttrJSON := d.Get("extensible_attributes").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'extensible_attributes' field: %s", err.Error())
		}
	}

	var tenantID string
	for attrName, attrValueInf := range extAttrs {
		attrValue, _ := attrValueInf.(string)
		if attrName == "Tenant ID" {
			tenantID = attrValue
		}
	}

	if hostName == "" {
		return fmt.Errorf("'host_name' is mandatory to be passed for Association of IP")
	}

	connector := m.(*ibclient.Connector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	if (zone != "" || len(zone) != 0) && (dnsView != "" || len(dnsView) != 0) && enableDns {
		obj, err := objMgr.GetHostRecordByRef(d.Id())
		if err != nil {
			return fmt.Errorf("Error getting IP from network block %s: %s", cidr, err.Error())
		}
		d.SetId(obj.Ref)
	} else if enableDhcp || !enableDns {
		obj, err := objMgr.GetHostRecordByRef(d.Id())
		if err != nil {
			return fmt.Errorf("Error getting IP from network block %s: %s", cidr, err.Error())
		}
		d.SetId(obj.Ref)
	} else {
		obj, err := objMgr.GetFixedAddressByRef(d.Id())
		if err != nil {
			return fmt.Errorf("Error getting IP from network block %s: %s", cidr, err.Error())
		}
		d.SetId(obj.Ref)
	}
	return nil
}

//we are updating the record with an empty mac address after the vm has been
//destroyed because if we implement the delete hostrecord method here then there
//will be a conflict of resources
func resourceIPAssociationDelete(d *schema.ResourceData, m interface{}, isIPv6 bool) error {

	ipAddr := d.Get("ip_addr").(string)
	hostName := d.Get("host_name").(string)

	enableDns := d.Get("enable_dns").(bool)
	enableDhcp := d.Get("enable_dhcp").(bool)
	dnsView := d.Get("dns_view").(string)
	zone := d.Get("zone").(string)
	duid := d.Get("duid").(string)

	comment := d.Get("comment").(string)
	extAttrJSON := d.Get("extensible_attributes").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'extensible_attributes' field: %s", err.Error())
		}
	}
	var tenantID string
	for attrName, attrValueInf := range extAttrs {
		attrValue, _ := attrValueInf.(string)
		if attrName == "Tenant ID" {
			tenantID = attrValue
		}
	}

	if hostName == "" {
		return fmt.Errorf("'host_name' is mandatory to be passed for removal of Association")
	}

	var recFQDN string
	if len(zone) > 0 {
		recFQDN = hostName + "." + zone
	} else {
		recFQDN = hostName
	}

	ZeroMacAddr := "00:00:00:00:00:00"
	connector := m.(*ibclient.Connector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	if (zone != "" || len(zone) != 0) && (dnsView != "" || len(dnsView) != 0) && enableDns {
		if isIPv6 {
			obj, err := objMgr.UpdateHostRecord(
				d.Id(),
				enableDns,
				enableDhcp,
				recFQDN,
				"", ipAddr,
				"", duid,
				comment,
				extAttrs, []string{})
			if err != nil {
				return fmt.Errorf("Error updating Host record of IP from network block having reference %s: %s", d.Id(), err.Error())
			}
			d.SetId(obj.Ref)
		} else {
			obj, err := objMgr.UpdateHostRecord(
				d.Id(),
				enableDns,
				enableDhcp,
				recFQDN,
				ipAddr, "",
				ZeroMacAddr, "",
				comment,
				extAttrs, []string{})
			if err != nil {
				return fmt.Errorf("Error updating Host record of IP from network block having reference %s: %s", d.Id(), err.Error())
			}
			d.SetId(obj.Ref)
		}
	} else if enableDhcp || !enableDns {
		// default value of enableDns is true. When user sets enableDhcp as true and does not pass a enableDns flag
		enableDns = false

		if isIPv6 {
			obj, err := objMgr.UpdateHostRecord(
				d.Id(),
				enableDns,
				enableDhcp,
				recFQDN,
				"", ipAddr,
				"", duid,
				comment,
				extAttrs, []string{})
			if err != nil {
				return fmt.Errorf("Error updating Host record of IP from network block having reference %s: %s", d.Id(), err.Error())
			}
			d.SetId(obj.Ref)
		} else {
			obj, err := objMgr.UpdateHostRecord(
				d.Id(),
				enableDns,
				enableDhcp,
				recFQDN,
				ipAddr, "",
				ZeroMacAddr, "",
				comment,
				extAttrs, []string{})
			if err != nil {
				return fmt.Errorf("Error updating Host record of IP from network block having reference %s: %s", d.Id(), err.Error())
			}
			d.SetId(obj.Ref)
		}
	} else {
		matchClient := "MAC_ADDRESS"
		if isIPv6 {
			matchClient = ""
			_, err := objMgr.UpdateFixedAddress(d.Id(), hostName, matchClient, duid, comment, extAttrs)
			if err != nil {
				return fmt.Errorf("Error Releasing IP from network block having reference %s: %s", d.Id(), err.Error())
			}
		} else {
			_, err := objMgr.UpdateFixedAddress(d.Id(), hostName, matchClient, ZeroMacAddr, comment, extAttrs)
			if err != nil {
				return fmt.Errorf("Error Releasing IP from network block having reference %s: %s", d.Id(), err.Error())
			}
		}
		d.SetId("")
	}

	return nil
}

func Resource(d *schema.ResourceData, m interface{}, isIPv6 bool) error {

	networkViewName := d.Get("network_view_name").(string)
	enableDns := d.Get("enable_dns").(bool)
	enableDhcp := d.Get("enable_dhcp").(bool)
	dnsView := d.Get("dns_view").(string)
	zone := d.Get("zone").(string)
	hostName := d.Get("host_name").(string)

	cidr := d.Get("cidr").(string)
	ipAddr := d.Get("ip_addr").(string)
	macAddr := d.Get("mac_addr").(string)
	duid := d.Get("duid").(string)

	comment := d.Get("comment").(string)
	extAttrJSON := d.Get("extensible_attributes").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'extensible_attributes' field: %s", err.Error())
		}
	}
	var tenantID string
	for attrName, attrValue := range extAttrs {
		attrValue, _ := attrValue.(string)
		if attrName == "Tenant ID" {
			tenantID = attrValue
		}
	}

	if hostName == "" {
		return fmt.Errorf("'host_name' is mandatory to be passed for Association")
	}

	//conversion from bit reversed EUI-48 format to hexadecimal EUI-48 format
	macAddr = strings.Replace(macAddr, "-", ":", -1)
	var recFQDN string
	if len(zone) > 0 {
		recFQDN = hostName + "." + zone
	} else {
		recFQDN = hostName
	}

	connector := m.(*ibclient.Connector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	//var err error
	if (zone != "" || len(zone) != 0) && (dnsView != "" || len(dnsView) != 0) && enableDns {
		if isIPv6 {
			hostRecordObj, err := objMgr.GetHostRecord(recFQDN, "", ipAddr)
			if err != nil {
				return fmt.Errorf("GetHostRecord failed from IPv6 network block %s:%s", cidr, err.Error())
			}
			if hostRecordObj == nil {
				return fmt.Errorf("HostRecord %s not found.", recFQDN)
			}
			Obj, err := objMgr.UpdateHostRecord(
				hostRecordObj.Ref,
				enableDns,
				enableDhcp,
				recFQDN,
				"", ipAddr,
				"", duid,
				comment,
				extAttrs, []string{})
			if err != nil {
				return fmt.Errorf("UpdateHost Record failed from IPv6 network block %s: %s", cidr, err.Error())
			}
			d.SetId(Obj.Ref)
		} else {
			hostRecordObj, err := objMgr.GetHostRecord(recFQDN, ipAddr, "")
			if err != nil {
				return fmt.Errorf("GetHostRecord failed from IPv4 network block %s:%s", cidr, err.Error())
			}
			if hostRecordObj == nil {
				return fmt.Errorf("HostRecord %s not found.", recFQDN)
			}
			Obj, err := objMgr.UpdateHostRecord(
				hostRecordObj.Ref,
				enableDns,
				enableDhcp,
				recFQDN,
				ipAddr, "",
				macAddr, "",
				comment,
				extAttrs, []string{})
			if err != nil {
				return fmt.Errorf("UpdateHost Record failed from network block %s:%s", cidr, err.Error())
			}
			d.SetId(Obj.Ref)
		}
	} else if enableDhcp || !enableDns {
		// default value of enableDns is true. When user sets enableDhcp as true and does not pass a enableDns flag
		enableDns = false

		if isIPv6 {
			hostRecordObj, err := objMgr.GetHostRecord(recFQDN, "", ipAddr)
			if err != nil {
				return fmt.Errorf("GetHostRecord failed from IPv6 network block %s:%s", cidr, err.Error())
			}
			if hostRecordObj == nil {
				return fmt.Errorf("HostRecord %s not found.", recFQDN)
			}
			Obj, err := objMgr.UpdateHostRecord(
				hostRecordObj.Ref,
				enableDns,
				enableDhcp,
				recFQDN,
				"", ipAddr,
				"", duid,
				comment,
				extAttrs, []string{})
			if err != nil {
				return fmt.Errorf("UpdateHost Record failed from IPv6 network block %s: %s", cidr, err.Error())
			}
			d.SetId(Obj.Ref)
		} else {
			hostRecordObj, err := objMgr.GetHostRecord(recFQDN, ipAddr, "")
			if err != nil {
				return fmt.Errorf("GetHostRecord failed from IPv4 network block %s:%s", cidr, err.Error())
			}
			if hostRecordObj == nil {
				return fmt.Errorf("HostRecord %s not found.", recFQDN)
			}
			Obj, err := objMgr.UpdateHostRecord(
				hostRecordObj.Ref,
				enableDns,
				enableDhcp,
				recFQDN,
				ipAddr, "",
				macAddr, "",
				comment,
				extAttrs, []string{})
			if err != nil {
				return fmt.Errorf("UpdateHost Record failed from network block %s:%s", cidr, err.Error())
			}
			d.SetId(Obj.Ref)
		}
	} else {
		fixedAddressObj, err := objMgr.GetFixedAddress(networkViewName, cidr, ipAddr, isIPv6, "")
		if err != nil {
			return fmt.Errorf("GetFixedAddress failed from network block %s:%s", cidr, err.Error())
		}
		if fixedAddressObj == nil {
			return fmt.Errorf("FixedAddress %s not found in network %s.", ipAddr, cidr)
		}

		var macOrDuid string
		macOrDuid = macAddr
		matchClient := "MAC_ADDRESS"
		if isIPv6 {
			matchClient = ""
			macOrDuid = duid
		}
		_, err = objMgr.UpdateFixedAddress(fixedAddressObj.Ref, hostName, matchClient, macOrDuid, comment, extAttrs)
		if err != nil {
			return fmt.Errorf("UpdateFixedAddress error from network block %s:%s", cidr, err.Error())
		}
		d.SetId(fixedAddressObj.Ref)
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
