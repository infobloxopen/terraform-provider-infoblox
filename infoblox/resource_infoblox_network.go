package infoblox

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func resourceNetwork() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: stateImporter,
		},

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
				Description: "The number of IP's you want to reserve in IPv4 Network.",
			},
			"reserve_ipv6": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of IP's you want to reserve in IPv6 Network",
			},
			"gateway": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Gateway's IP address of the network. By default, the first IP address is set as gateway address; if the value is 'none' then the network has no gateway.",
				Computed:    true,
				// TODO: implement full support for this field
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
	reserveIPv4 := d.Get("reserve_ip").(int)
	reserveIPv6 := d.Get("reserve_ipv6").(int)
	if reserveIPv6 > 255 || reserveIPv6 < 0 {
		return fmt.Errorf("reserve_ipv6 value must be in range 0..255")
	}

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

	autoAllocateGateway := gateway == ""

	if !autoAllocateGateway && gateway != "none" {
		_, err = objMgr.AllocateIP(networkViewName, network.Cidr, gateway, isIPv6, ZeroMacAddr, "", "", nil)
		if err != nil {
			return fmt.Errorf(
				"reservation of the IP address '%s' in network block '%s' from network view '%s' failed: %s",
				gateway, network.Cidr, networkViewName, err.Error())
		}
	}

	if isIPv6 {
		for i := 1; i <= reserveIPv6; i++ {
			reservedDuid := fmt.Sprintf("00:%.2x", i)
			newAddr, err := objMgr.AllocateIP(
				networkViewName, network.Cidr, "", isIPv6, reservedDuid, "", "", nil)
			if err != nil {
				return fmt.Errorf(
					"reservation in network block '%s' from network view '%s' failed: %s",
					network.Cidr, networkViewName, err.Error())
			}
			if autoAllocateGateway && i == 1 {
				gateway = newAddr.IPv4Address
			}
		}
	} else {
		for i := 1; i <= reserveIPv4; i++ {
			newAddr, err := objMgr.AllocateIP(
				networkViewName, network.Cidr, "", isIPv6, ZeroMacAddr, "", "", nil)
			if err != nil {
				return fmt.Errorf(
					"reservation in network block '%s' from network view '%s' failed: %s",
					network.Cidr, networkViewName, err.Error())
			}
			if autoAllocateGateway && i == 1 {
				gateway = newAddr.IPv4Address
			}
		}
	}

	d.Set("gateway", gateway)
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
		return fmt.Errorf("getting Network block from network view (%s) failed : %s", networkViewName, err)
	}

	if obj.Ea != nil && len(obj.Ea) > 0 {
		// TODO: temporary scaffold, need to rework marshalling/unmarshalling of EAs
		//       (avoiding additional layer of keys ("value" key)
		eaMap := (map[string]interface{})(obj.Ea)
		ea, err := json.Marshal(eaMap)
		if err != nil {
			return err
		}
		if err = d.Set("ext_attrs", string(ea)); err != nil {
			return err
		}
	}

	if err = d.Set("comment", obj.Comment); err != nil {
		return err
	}

	if obj.NetviewName != "" {
		if err = d.Set("network_view", obj.NetviewName); err != nil {
			return err
		}
	} else {
		if err = d.Set("network_view", "default"); err != nil {
			return err
		}
	}

	if err = d.Set("cidr", obj.Cidr); err != nil {
		return err
	}

	d.SetId(obj.Ref)

	return nil
}
func resourceNetworkUpdate(d *schema.ResourceData, m interface{}) (err error) {
	defer func() {
		if err != nil {
			d.Partial(true)
		}
	}()

	networkViewName := d.Get("network_view").(string)
	if d.HasChange("network_view") {
		return fmt.Errorf("changing the value of 'network_view' field is not allowed")
	}
	if d.HasChange("cidr") {
		return fmt.Errorf("changing the value of 'cidr' field is not allowed")
	}
	if d.HasChange("reserve_ip") {
		return fmt.Errorf("changing the value of 'reserve_ip' field is not allowed")
	}
	if d.HasChange("reserve_ipv6") {
		return fmt.Errorf("changing the value of 'reserve_ipv6' field is not allowed")
	}
	if d.HasChange("gateway") {
		return fmt.Errorf("changing the value of 'gateway' field is not allowed")
	}
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err = json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
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
