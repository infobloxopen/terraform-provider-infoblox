package infoblox

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func resourceNetwork() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"network_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
				Description: "Network view name available in NIOS Server.",
			},
			"parent_cidr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The parent network container block in cidr format to allocate from.",
			},
			"allocate_prefix_len": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Set the parameter's value > 0 to allocate next available network with corresponding prefix length from the network container defined by 'parent_cidr'",
			},
			"cidr": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The network block in cidr format.",
			},
			"reserve_ip": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "The number of IP's you want to reserve in IPv4 Network.",
			},
			"reserve_ipv6": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "The number of IP's you want to reserve in IPv6 Network",
			},
			"gateway": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Gateways's IP address of your network block. By default first IPv4 address is set as gateway address.",
				Computed:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A string describing the network",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "The Extensible attributes of the Network, as a map in JSON format",
			},
		},
	}
}

func resourceNetworkCreate(d *schema.ResourceData, m interface{}, isIPv6 bool) error {

	networkViewName := d.Get("network_view").(string)
	parentCidr := d.Get("parent_cidr").(string)
	prefixLen := d.Get("allocate_prefix_len").(int)
	cidr := d.Get("cidr").(string)
	reserveIP := d.Get("reserve_ip").(int)
	reserveIPv6 := d.Get("reserve_ipv6").(int)
	gateway := d.Get("gateway").(string)

	comment := d.Get("comment").(string)
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
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

	ZeroMacAddr := "00:00:00:00:00:00"
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)
	ea := make(map[string]interface{})

	var network *ibclient.Network
	var err error
	if cidr == "" && parentCidr != "" && prefixLen > 1 {
		_, err := objMgr.GetNetworkContainer(networkViewName, parentCidr, isIPv6, nil)
		if err != nil {
			return fmt.Errorf(
				"Allocation of network block within network container '%s' under network view '%s' failed: %s", parentCidr, networkViewName, err.Error())
		}

		network, err = objMgr.AllocateNetwork(networkViewName, parentCidr, isIPv6, uint(prefixLen), comment, extAttrs)
		if err != nil {
			return fmt.Errorf("Allocation of network block failed in network view (%s) : %s", networkViewName, err)
		}
		d.Set("cidr", network.Cidr)
	} else if cidr != "" {
		network, err = objMgr.CreateNetwork(networkViewName, cidr, isIPv6, comment, extAttrs)
		if err != nil {
			return fmt.Errorf("Creation of network block failed in network view (%s) : %s", networkViewName, err)
		}
	} else {
		return fmt.Errorf("Creation of network block failed: neither cidr nor parentCidr with allocate_prefix_len was specified.")
	}

	if isIPv6 {
		// We need Zeroduid since AWS mandates first 3 IPv6 addresses to be reserved
		Zeroduid := "00"
		for i := 1; i <= reserveIPv6; i++ {
			Zeroduid += ":00"
			_, err = objMgr.AllocateIP(networkViewName, network.Cidr, gateway, isIPv6, Zeroduid, "", comment, ea)
			if err != nil {
				return fmt.Errorf("Reservation in network block failed in network view(%s):%s", networkViewName, err)
			}
		}
	} else {
		// Check whether gateway or ip address already allocated
		if gateway != "none" {
			gatewayIP, err := objMgr.GetFixedAddress(networkViewName, network.Cidr, gateway, false, "")
			if err == nil && gatewayIP != nil {
				fmt.Printf("Gateway already created")
			} else if gatewayIP == nil {
				gatewayIP, err = objMgr.AllocateIP(networkViewName, network.Cidr, gateway, isIPv6, ZeroMacAddr, "", comment, ea)
				if err != nil {
					return fmt.Errorf("Gateway Creation failed in network block(%s) error: %s", network.Cidr, err)
				}
			}
			d.Set("gateway", gatewayIP.IPv4Address)
		}

		for i := 1; i <= reserveIP; i++ {
			_, err = objMgr.AllocateIP(networkViewName, network.Cidr, gateway, isIPv6, ZeroMacAddr, "", comment, ea)
			if err != nil {
				return fmt.Errorf("Reservation in network block failed in network view(%s):%s", networkViewName, err)
			}
		}

	}
	d.SetId(network.Ref)
	return nil
}
func resourceNetworkRead(d *schema.ResourceData, m interface{}) error {

	networkViewName := d.Get("network_view").(string)
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
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

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	obj, err := objMgr.GetNetworkByRef(d.Id())
	if err != nil {
		return fmt.Errorf("Getting Network block from network view (%s) failed : %s", networkViewName, err)
	}
	d.SetId(obj.Ref)
	return nil
}
func resourceNetworkUpdate(d *schema.ResourceData, m interface{}) error {

	networkViewName := d.Get("network_view").(string)
	if d.HasChange("network_view") {
		return fmt.Errorf("changing the value of 'network_view' field is not allowed")
	}
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
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

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)
	var Network *ibclient.Network

	comment := ""
	commentVal, commentFieldFound := d.GetOk("comment")
	if commentFieldFound {
		comment = commentVal.(string)
	}

	var err error

	Network, err = objMgr.UpdateNetwork(d.Id(), extAttrs, comment)
	if err != nil {
		return fmt.Errorf("Updation of IPv4 Network under network view '%s' failed: '%s'", networkViewName, err.Error())
	}

	d.SetId(Network.Ref)
	return nil
}

func resourceNetworkDelete(d *schema.ResourceData, m interface{}) error {
	networkViewName := d.Get("network_view").(string)
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
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

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	_, err := objMgr.DeleteNetwork(d.Id())
	if err != nil {
		return fmt.Errorf("Deletion of Network block failed from network view(%s): %s", networkViewName, err)
	}
	d.SetId("")
	return nil
}

func resourceIPv4NetworkCreate(d *schema.ResourceData, m interface{}) error {
	return resourceNetworkCreate(d, m, false)
}

func resourceIPv4Network() *schema.Resource {
	nw := resourceNetwork()
	nw.Create = resourceIPv4NetworkCreate
	nw.Read = resourceNetworkRead
	nw.Update = resourceNetworkUpdate
	nw.Delete = resourceNetworkDelete

	return nw
}

func resourceIPv6NetworkCreate(d *schema.ResourceData, m interface{}) error {
	return resourceNetworkCreate(d, m, true)
}

func resourceIPv6Network() *schema.Resource {
	nw := resourceNetwork()
	nw.Create = resourceIPv6NetworkCreate
	nw.Read = resourceNetworkRead
	nw.Update = resourceNetworkUpdate
	nw.Delete = resourceNetworkDelete

	return nw
}
