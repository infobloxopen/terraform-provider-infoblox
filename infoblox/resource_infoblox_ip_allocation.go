package infoblox

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client"
)

func resourceIPAllocation() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"network_view_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
				Description: "Network view name of NIOS server.",
			},
			"enable_dns": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "flag that defines if the host record is to be used for DNS Purposes.",
			},
			"enable_dhcp": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "flag that defines if the host record is to be used for IPAM Purposes.",
			},
			"dns_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Dns View under which the zone has been created.",
			},
			"zone": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Zone under which host record has to be created.",
			},
			"cidr": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The address in cidr format.",
			},
			"ip_addr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IP address of cloud instance. Set a valid IP for static allocation and leave empty if dynamically allocated.",
				Computed:    true,
			},
			"mac_addr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "MAC Address of cloud instance.",
			},
			"duid": {
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
				Description: "A description of the IP allocation.",
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

func resourceIPAllocationRequest(d *schema.ResourceData, m interface{}, isIPv6 bool) error {
	networkViewName := d.Get("network_view_name").(string)
	dnsView := d.Get("dns_view").(string)
	enableDns := d.Get("enable_dns").(bool)
	enableDhcp := d.Get("enable_dhcp").(bool)
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

	ZeroMacAddr := "00:00:00:00:00:00"
	connector := m.(*ibclient.Connector)

	var tenantID string
	for attrName, attrValueInf := range extAttrs {
		attrValue, _ := attrValueInf.(string)
		if attrName == "Tenant ID" {
			tenantID = attrValue
		}
	}

	if hostName == "" {
		return fmt.Errorf("'host_name' is mandatory to be passed for Allocation of IP")
	}

	var recFQDN string
	if len(zone) > 0 {
		recFQDN = hostName + "." + zone
	} else {
		recFQDN = hostName
	}

	if macAddr == "" {
		macAddr = ZeroMacAddr
	}

	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	if (zone != "" || len(zone) != 0) && (dnsView != "" || len(dnsView) != 0) && enableDns {
		if isIPv6 {
			Obj, err := objMgr.CreateHostRecord(
				enableDns,
				enableDhcp,
				recFQDN,
				networkViewName,
				dnsView,
				"", cidr,
				"", ipAddr,
				"", duid,
				comment,
				extAttrs, []string{})
			if err != nil {
				return fmt.Errorf(
					"Error in allocating an IPv6 address and creating a host record in cidr %s: %s", cidr, err.Error())
			}
			d.Set("ip_addr", Obj.Ipv6Addrs[0].Ipv6Addr)
			d.SetId(Obj.Ref)
		} else {
			Obj, err := objMgr.CreateHostRecord(
				enableDns,
				enableDhcp,
				recFQDN,
				networkViewName,
				dnsView,
				cidr, "",
				ipAddr, "",
				macAddr, "",
				comment,
				extAttrs, []string{})
			if err != nil {
				return fmt.Errorf(
					"Error in allocating an IPv4 address and creating a host record in cidr %s: %s", cidr, err.Error())
			}
			d.Set("ip_addr", Obj.Ipv4Addrs[0].Ipv4Addr)
			d.SetId(Obj.Ref)
		}

	} else if enableDhcp || !enableDns {
		// default value of enableDns is true. When user sets enableDhcp as true and does not pass a enableDns flag
		enableDns = false

		if isIPv6 {
			Obj, err := objMgr.CreateHostRecord(
				enableDns,
				enableDhcp,
				recFQDN,
				networkViewName,
				dnsView,
				"", cidr,
				"", ipAddr,
				"", duid,
				comment,
				extAttrs, []string{})
			if err != nil {
				return fmt.Errorf(
					"Error in allocating an IPv6 address and creating a host record in cidr %s: %s", cidr, err.Error())
			}
			d.Set("ip_addr", Obj.Ipv6Addrs[0].Ipv6Addr)
			d.SetId(Obj.Ref)
		} else {
			Obj, err := objMgr.CreateHostRecord(
				enableDns,
				enableDhcp,
				recFQDN,
				networkViewName,
				dnsView,
				cidr, "",
				ipAddr, "",
				macAddr, "",
				comment,
				extAttrs, []string{})
			if err != nil {
				return fmt.Errorf(
					"Error in allocating an IPv4 address and creating a host record in cidr %s: %s", cidr, err.Error())
			}
			d.Set("ip_addr", Obj.Ipv4Addrs[0].Ipv4Addr)
			d.SetId(Obj.Ref)
		}

	} else {
		if isIPv6 {
			Obj, err := objMgr.AllocateIP(networkViewName, cidr, ipAddr, isIPv6, duid, hostName, comment, extAttrs)
			if err != nil {
				return fmt.Errorf("Error allocating IP from network block %s: %s", cidr, err.Error())
			}
			d.Set("ip_addr", Obj.IPv6Address)
			d.SetId(Obj.Ref)
		} else {
			Obj, err := objMgr.AllocateIP(networkViewName, cidr, ipAddr, isIPv6, macAddr, hostName, comment, extAttrs)
			if err != nil {
				return fmt.Errorf("Error allocating IP from network block %s: %s", cidr, err.Error())
			}
			d.Set("ip_addr", Obj.IPv4Address)
			d.SetId(Obj.Ref)
		}
	}
	return nil
}

func resourceIPAllocationGet(d *schema.ResourceData, m interface{}, isIPv6 bool) error {

	dnsView := d.Get("dns_view").(string)
	zone := d.Get("zone").(string)
	cidr := d.Get("cidr").(string)

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
			break
		}
	}

	connector := m.(*ibclient.Connector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	if (zone != "" || len(zone) != 0) && (dnsView != "" || len(dnsView) != 0) {
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

func resourceIPAllocationUpdate(d *schema.ResourceData, m interface{}, isIPv6 bool) error {

	enableDns := d.Get("enable_dns").(bool)
	enableDhcp := d.Get("enable_dhcp").(bool)
	dnsView := d.Get("dns_view").(string)
	zone := d.Get("zone").(string)
	hostName := d.Get("host_name").(string)

	duid := d.Get("duid").(string)
	macAddr := d.Get("mac_addr").(string)

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
		return fmt.Errorf("'hostName' is mandatory to be passed for Allocation of IP")
	}

	var recFQDN string
	if len(zone) > 0 {
		recFQDN = hostName + "." + zone
	} else {
		recFQDN = hostName
	}

	connector := m.(*ibclient.Connector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	if (zone != "" || len(zone) != 0) && (dnsView != "" || len(dnsView) != 0) && enableDns {
		hostRecordObj, _ := objMgr.GetHostRecordByRef(d.Id())

		if isIPv6 {
			IPAddrObj := hostRecordObj.Ipv6Addr
			obj, err := objMgr.UpdateHostRecord(
				d.Id(),
				enableDns,
				enableDhcp,
				recFQDN,
				"", IPAddrObj,
				"", duid,
				comment,
				extAttrs, []string{})
			if err != nil {
				return fmt.Errorf(
					"Error updating Host record of IPv6 from network block having reference %s: %s", d.Id(), err.Error())
			}
			d.SetId(obj.Ref)
		} else {
			IPAddrObj := hostRecordObj.Ipv4Addr
			obj, err := objMgr.UpdateHostRecord(
				d.Id(),
				enableDns,
				enableDhcp,
				recFQDN,
				IPAddrObj, "",
				macAddr, "",
				comment,
				extAttrs, []string{})
			if err != nil {
				return fmt.Errorf(
					"Error updating Host record of IPv4 from network block having reference %s: %s", d.Id(), err.Error())
			}
			d.SetId(obj.Ref)
		}
	} else if enableDhcp || !enableDns {
		hostRecordObj, _ := objMgr.GetHostRecordByRef(d.Id())

		// default value of enableDns is true. When user sets enableDhcp as true and does not pass a enableDns flag
		enableDns = false

		if isIPv6 {
			IPAddrObj := hostRecordObj.Ipv6Addr
			obj, err := objMgr.UpdateHostRecord(
				d.Id(),
				enableDns,
				enableDhcp,
				recFQDN,
				"", IPAddrObj,
				"", duid,
				comment,
				extAttrs, []string{})
			if err != nil {
				return fmt.Errorf(
					"Error updating Host record of IPv6 from network block having reference %s: %s", d.Id(), err.Error())
			}
			d.SetId(obj.Ref)
		} else {
			IPAddrObj := hostRecordObj.Ipv4Addr
			obj, err := objMgr.UpdateHostRecord(
				d.Id(),
				enableDns,
				enableDhcp,
				recFQDN,
				IPAddrObj, "",
				macAddr, "",
				comment,
				extAttrs, []string{})
			if err != nil {
				return fmt.Errorf(
					"Error updating Host record of IPv4 from network block having reference %s: %s", d.Id(), err.Error())
			}
			d.SetId(obj.Ref)
		}
	} else {
		var macOrDuid string
		match_client := "MAC_ADDRESS"
		if isIPv6 {
			macOrDuid = duid
			match_client = ""
		} else {
			macOrDuid = macAddr
		}
		obj, err := objMgr.UpdateFixedAddress(d.Id(), hostName, match_client, macOrDuid, comment, extAttrs)
		if err != nil {
			return fmt.Errorf("Error updating IP from network block having reference %s: %s", d.Id(), err.Error())
		}
		d.SetId(obj.Ref)
	}
	return nil
}

func resourceIPAllocationRelease(d *schema.ResourceData, m interface{}, isIPv6 bool) error {

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
	for attrName, attrValue := range extAttrs {
		if attrName == "Tenant ID" {
			tenantID = attrValue.(string)
			break
		}
	}

	connector := m.(*ibclient.Connector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	if (zone != "" || len(zone) != 0) && (dnsView != "" || len(dnsView) != 0) {
		_, err := objMgr.DeleteHostRecord(d.Id())
		if err != nil {
			return fmt.Errorf("Error Releasing IP from network block having reference %s: %s", d.Id(), err.Error())
		}
	} else {
		_, err := objMgr.DeleteFixedAddress(d.Id())
		if err != nil {
			return fmt.Errorf("Error Releasing IP from network block having reference %s: %s", d.Id(), err.Error())
		}
	}
	d.SetId("")

	return nil
}

// Code snippet for IPv4 IP Allocation
func resourceIPv4AllocationRequest(d *schema.ResourceData, m interface{}) error {
	return resourceIPAllocationRequest(d, m, false)
}

func resourceIPv4AllocationGet(d *schema.ResourceData, m interface{}) error {
	return resourceIPAllocationGet(d, m, false)
}

func resourceIPv4AllocationUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceIPAllocationUpdate(d, m, false)
}

func resourceIPv4AllocationRelease(d *schema.ResourceData, m interface{}) error {
	return resourceIPAllocationRelease(d, m, false)
}

func resourceIPv4Allocation() *schema.Resource {
	ipv4Allocation := resourceIPAllocation()
	ipv4Allocation.Create = resourceIPv4AllocationRequest
	ipv4Allocation.Read = resourceIPv4AllocationGet
	ipv4Allocation.Update = resourceIPv4AllocationUpdate
	ipv4Allocation.Delete = resourceIPv4AllocationRelease

	return ipv4Allocation
}

// Code snippet for IPv6 IP allocation
func resourceIPv6AllocationRequest(d *schema.ResourceData, m interface{}) error {
	return resourceIPAllocationRequest(d, m, true)
}

func resourceIPv6AllocationGet(d *schema.ResourceData, m interface{}) error {
	return resourceIPAllocationGet(d, m, true)
}

func resourceIPv6AllocationUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceIPAllocationUpdate(d, m, true)
}

func resourceIPv6AllocationRelease(d *schema.ResourceData, m interface{}) error {
	return resourceIPAllocationRelease(d, m, true)
}

func resourceIPv6Allocation() *schema.Resource {
	ipv6Allocation := resourceIPAllocation()
	ipv6Allocation.Create = resourceIPv6AllocationRequest
	ipv6Allocation.Read = resourceIPv6AllocationGet
	ipv6Allocation.Update = resourceIPv6AllocationUpdate
	ipv6Allocation.Delete = resourceIPv6AllocationRelease

	return ipv6Allocation
}
