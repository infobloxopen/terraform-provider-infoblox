package infoblox

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"regexp"
)

var (
	networkIPv4Regexp = regexp.MustCompile("^network/.+")
	networkIPv6Regexp = regexp.MustCompile("^ipv6network/.+")
)

func resourceNetwork() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: resourceNetworkImport,
		},
		CustomizeDiff: func(context context.Context, d *schema.ResourceDiff, meta interface{}) error {
			if internalID := d.Get("internal_id"); internalID == "" || internalID == nil {
				err := d.SetNewComputed("internal_id")
				if err != nil {
					return err
				}
			}
			return nil
		},

		Schema: map[string]*schema.Schema{
			"network_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     defaultNetView,
				Description: "Network view name available in NIOS Server.",
			},
			"parent_cidr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The parent network container block in cidr format to allocate from.",
			},
			"filter_params": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The parent network container block's extensible attributes.",
			},
			"object": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The object type to allocate from the parent network container block.",
				Default:     "networkcontainer",
				ValidateFunc: validation.StringInSlice([]string{
					"networkcontainer", "network",
				}, false),
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
				Computed:    true,
				Description: "The number of IP's you want to reserve in IPv4 Network.",
			},
			"reserve_ipv6": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
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
				Optional:    true,
				Default:     "",
				Description: "The Extensible attributes of the Network",
			},
			"internal_id": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "Internal ID of an object at NIOS side," +
					" used by Infoblox Terraform plugin to search for a NIOS's object" +
					" which corresponds to the Terraform resource.",
			},
			"ref": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "NIOS object's reference, not to be set by a user.",
			},
		},
	}
}

func resourceNetworkCreate(d *schema.ResourceData, m interface{}, isIPv6 bool) error {
	// Check if internal_id is set manually
	if intId := d.Get("internal_id"); intId.(string) != "" {
		return fmt.Errorf("the value of 'internal_id' field must not be set manually")
	}
	networkViewName := d.Get("network_view").(string)
	parentCidr := d.Get("parent_cidr").(string)
	nextAvailableFilter := d.Get("filter_params").(string)
	object := d.Get("object").(string)
	prefixLen := d.Get("allocate_prefix_len").(int)
	cidr := d.Get("cidr").(string)
	reserveIPv4 := d.Get("reserve_ip").(int)
	reserveIPv6 := d.Get("reserve_ipv6").(int)
	if reserveIPv6 > 255 || reserveIPv6 < 0 {
		return fmt.Errorf("reserve_ipv6 value must be in range 0..255")
	}

	gateway := d.Get("gateway").(string)

	comment := d.Get("comment").(string)

	extAttrsJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrsJSON)
	if err != nil {
		return err
	}

	// Generate UUID for internal_id and add to the EA
	internalId := generateInternalId()
	extAttrs[eaNameForInternalId] = internalId.String()

	var tenantID string
	for attrName, attrValueInf := range extAttrs {
		attrValue, _ := attrValueInf.(string)
		if attrName == eaNameForTenantId {
			tenantID = attrValue
			break
		}
	}

	ZeroMacAddr := "00:00:00:00:00:00"
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	var network *ibclient.Network
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

	} else if cidr == "" && nextAvailableFilter != "" && prefixLen > 1 {
		var (
			eaMap map[string]string
		)
		err = json.Unmarshal([]byte(nextAvailableFilter), &eaMap)
		if err != nil {
			return fmt.Errorf("error unmarshalling extra attributes of network container: %s", err)
		}

		network, err = objMgr.AllocateNetworkByEA(networkViewName, isIPv6, comment, extAttrs, eaMap, prefixLen, object)
		if err != nil {
			return fmt.Errorf("allocation of network block failed in network with extra attributes (%s) : %s", nextAvailableFilter, err)
		}

	} else if cidr != "" {
		network, err = objMgr.CreateNetwork(networkViewName, cidr, isIPv6, comment, extAttrs)
		if err != nil {
			return fmt.Errorf("Creation of network block failed in network view (%s) : %s", networkViewName, err)
		}
	} else {
		return fmt.Errorf("creation of network block failed: neither cidr nor parentCidr with allocate_prefix_len was specified")
	}

	d.SetId(network.Ref)
	if err = d.Set("internal_id", internalId.String()); err != nil {
		return err
	}
	if err = d.Set("ref", network.Ref); err != nil {
		return err
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
				gateway = newAddr.IPv6Address
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

	return nil
}

func resourceNetworkRead(d *schema.ResourceData, m interface{}) error {
	//networkViewName := d.Get("network_view").(string)

	extAttrsJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrsJSON)
	if err != nil {
		return err
	}

	//var tenantID string
	//for attrName, attrValueInf := range extAttrs {
	//	attrValue, _ := attrValueInf.(string)
	//	if attrName == eaNameForTenantId {
	//		tenantID = attrValue
	//		break
	//	}
	//}

	//connector := m.(ibclient.IBConnector)
	//objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	network, err := searchObjectByRefOrInternalId("Network", d, m)
	// Assertion of object type and error handling
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); ok {
			d.SetId("")
			return nil
		} else {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		}
	}
	var obj *ibclient.Network
	recJson, _ := json.Marshal(network)
	err = json.Unmarshal(recJson, &obj)
	if err != nil {
		return err
	}

	if err = d.Set("ref", obj.Ref); err != nil {
		return err
	}
	delete(extAttrs, eaNameForInternalId)

	omittedEAs := omitEAs(obj.Ea, extAttrs)

	if omittedEAs != nil && len(omittedEAs) > 0 {
		eaJSON, err := terraformSerializeEAs(omittedEAs)
		if err != nil {
			return err
		}

		if err = d.Set("ext_attrs", eaJSON); err != nil {
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
		if err = d.Set("network_view", defaultNetView); err != nil {
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
	var updateSuccessful bool
	defer func() {
		// Reverting the state back, in case of a failure,
		// otherwise Terraform will keep the values, which leaded to the failure,
		// in the state file.
		if !updateSuccessful {
			d.Partial(true)

			prevNetView, _ := d.GetChange("network_view")
			prevCIDR, _ := d.GetChange("cidr")
			prevParCIDR, _ := d.GetChange("parent_cidr")
			prevGW, _ := d.GetChange("gateway")
			prevPrefLen, _ := d.GetChange("allocate_prefix_len")
			prevNextAvailableFilter, _ := d.GetChange("filter_params")
			prevResIPv4, _ := d.GetChange("reserve_ip")
			prevResIPv6, _ := d.GetChange("reserve_ipv6")
			prevComment, _ := d.GetChange("comment")
			prevEa, _ := d.GetChange("ext_attrs")

			_ = d.Set("network_view", prevNetView.(string))
			_ = d.Set("cidr", prevCIDR.(string))
			_ = d.Set("parent_cidr", prevParCIDR.(string))
			_ = d.Set("gateway", prevGW.(string))
			_ = d.Set("allocate_prefix_len", prevPrefLen.(int))
			_ = d.Set("filter_params", prevNextAvailableFilter.(string))
			_ = d.Set("reserve_ip", prevResIPv4.(int))
			_ = d.Set("reserve_ipv6", prevResIPv6.(int))
			_ = d.Set("comment", prevComment.(string))
			_ = d.Set("ext_attrs", prevEa.(string))
		}
	}()

	if d.HasChange("internal_id") {
		return fmt.Errorf("changing the value of 'internal_id' field is not allowed")
	}
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
	if d.HasChange("filter_params") {
		return fmt.Errorf("changing the value of 'filter_params' field is not allowed")
	}

	networkViewName := d.Get("network_view").(string)
	oldExtAttrsJSON, newExtAttrsJSON := d.GetChange("ext_attrs")

	newExtAttrs, err := terraformDeserializeEAs(newExtAttrsJSON.(string))
	if err != nil {
		return err
	}

	oldExtAttrs, err := terraformDeserializeEAs(oldExtAttrsJSON.(string))
	if err != nil {
		return err
	}

	var tenantID string
	for attrName, attrValueInf := range newExtAttrs {
		attrValue, _ := attrValueInf.(string)
		if attrName == eaNameForTenantId {
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

	net, err := objMgr.GetNetworkByRef(d.Id())
	if err != nil {
		return fmt.Errorf("failed to read network for update operation: %w", err)
	}

	newExtAttrs, err = mergeEAs(net.Ea, newExtAttrs, oldExtAttrs, connector)
	if err != nil {
		return err
	}

	// Set the Terraform Internal ID to the NIOS EA if it is not already set
	internalId := d.Get("internal_id").(string)

	if internalId == "" {
		internalId = generateInternalId().String()
	}

	newInternalId := newInternalResourceIdFromString(internalId)
	newExtAttrs[eaNameForInternalId] = newInternalId.String()
	Network, err = objMgr.UpdateNetwork(net.Ref, newExtAttrs, comment)
	if err != nil {
		return fmt.Errorf("Updation of IP Network under network view '%s' failed: '%s'", networkViewName, err.Error())
	}
	updateSuccessful = true
	d.SetId(Network.Ref)
	if err = d.Set("internal_id", newInternalId.String()); err != nil {
		return err
	}
	if err = d.Set("ref", Network.Ref); err != nil {
		return err
	}

	return nil
}

func resourceNetworkDelete(d *schema.ResourceData, m interface{}) error {
	networkViewName := d.Get("network_view").(string)

	extAttrsJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrsJSON)
	if err != nil {
		return err
	}

	var tenantID string
	for attrName, attrValueInf := range extAttrs {
		attrValue, _ := attrValueInf.(string)
		if attrName == eaNameForTenantId {
			tenantID = attrValue
			break
		}
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	network, err := searchObjectByRefOrInternalId("Network", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		} else {
			d.SetId("")
			return nil
		}
	}

	// Assertion of object type and error handling
	var net *ibclient.Network
	recJson, _ := json.Marshal(network)
	err = json.Unmarshal(recJson, &net)
	if err != nil {
		return fmt.Errorf("failed to read network for delete operation: %w", err)
	}

	_, err = objMgr.DeleteNetwork(net.Ref)
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
	nw.Read = resourceIPv4NetworkRead
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
	nw.Read = resourceIPv6NetworkRead
	nw.Update = resourceNetworkUpdate
	nw.Delete = resourceNetworkDelete
	return nw
}

func resourceIPv4NetworkRead(d *schema.ResourceData, m interface{}) error {
	ref := d.Id()
	if !networkIPv4Regexp.MatchString(ref) {
		return fmt.Errorf("reference '%s' for 'network' object has an invalid format", ref)
	}

	return resourceNetworkRead(d, m)
}

func resourceIPv6NetworkRead(d *schema.ResourceData, m interface{}) error {
	ref := d.Id()
	if !networkIPv6Regexp.MatchString(ref) {
		return fmt.Errorf("reference '%s' for 'ipv6network' object has an invalid format", ref)
	}

	return resourceNetworkRead(d, m)
}

func resourceNetworkImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	networkViewName := d.Get("network_view").(string)

	extAttrsJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrsJSON)
	if err != nil {
		return nil, err
	}

	var tenantID string
	for attrName, attrValueInf := range extAttrs {
		attrValue, _ := attrValueInf.(string)
		if attrName == eaNameForTenantId {
			tenantID = attrValue
			break
		}
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	obj, err := objMgr.GetNetworkByRef(d.Id())
	if err != nil {
		return nil, fmt.Errorf("getting Network block from network view (%s) failed : %s", networkViewName, err)
	}

	if obj.Ea != nil && len(obj.Ea) > 0 {
		eaJSON, err := terraformSerializeEAs(obj.Ea)
		if err != nil {
			return nil, err
		}

		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return nil, err
		}
	}

	if err = d.Set("comment", obj.Comment); err != nil {
		return nil, err
	}

	if obj.NetviewName != "" {
		if err = d.Set("network_view", obj.NetviewName); err != nil {
			return nil, err
		}
	} else {
		if err = d.Set("network_view", defaultNetView); err != nil {
			return nil, err
		}
	}

	if err = d.Set("cidr", obj.Cidr); err != nil {
		return nil, err
	}

	d.SetId(obj.Ref)
	err = resourceNetworkUpdate(d, m)
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
