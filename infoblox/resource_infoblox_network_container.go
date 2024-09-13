package infoblox

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"regexp"
)

var (
	netContainerIPv4Regexp = regexp.MustCompile("^networkcontainer/.+")
	netContainerIPv6Regexp = regexp.MustCompile("^ipv6networkcontainer/.+")
)

func resourceNetworkContainer() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: resourceNetworkContainerImport,
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
				Description: "The name of network view for the network container.",
			},
			"cidr": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The network container's address, in CIDR format.",
			},
			"parent_cidr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The parent network container block in CIDR format to allocate from.",
			},
			"filter_params": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The parent network container block's extensible attributes.",
			},
			"allocate_prefix_len": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Set the parameter's value > 0 to allocate next available network container with corresponding prefix length from the network container defined by 'parent_cidr'",
			},
			"comment": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "A description of the network container.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "The Extensible attributes of the network container to be added/updated, as a map in JSON format",
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

func resourceNetworkContainerCreate(d *schema.ResourceData, m interface{}, isIPv6 bool) error {
	if intId := d.Get("internal_id"); intId.(string) != "" {
		return fmt.Errorf("the value of 'internal_id' field must not be set manually")
	}
	var (
		nc  *ibclient.NetworkContainer
		err error
	)

	nvName := d.Get("network_view").(string)
	cidr := d.Get("cidr").(string)
	parentCidr := d.Get("parent_cidr").(string)
	prefixLen := d.Get("allocate_prefix_len").(int)
	nextAvailableFilter := d.Get("filter_params").(string)
	comment := d.Get("comment").(string)

	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return fmt.Errorf("failed to create network container: %w", err)
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

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	// Attempt to allocate next available network container
	if cidr == "" && parentCidr != "" && prefixLen > 1 {
		_, err = objMgr.GetNetworkContainer(nvName, parentCidr, isIPv6, nil)
		if err != nil {
			return fmt.Errorf(
				"allocation of network block within network container '%s' under network view '%s' failed: %w", parentCidr, nvName, err)
		}

		nc, err = objMgr.AllocateNetworkContainer(nvName, parentCidr, isIPv6, uint(prefixLen), comment, extAttrs)
		if err != nil {
			return fmt.Errorf("allocation of network block in network view '%s' failed: %w", nvName, err)
		}
		if err = d.Set("cidr", nc.Cidr); err != nil {
			return err
		}
	} else if cidr == "" && nextAvailableFilter != "" && prefixLen > 1 {
		var eaMap map[string]string
		err = json.Unmarshal([]byte(nextAvailableFilter), &eaMap)
		if err != nil {
			return fmt.Errorf("error unmarshalling extra attributes of network container: %s", err)
		}

		nc, err = objMgr.AllocateNetworkContainerByEA(nvName, isIPv6, comment, extAttrs, eaMap, prefixLen)
		if err != nil {
			return fmt.Errorf("allocation of network block failed in network with extra attributes (%s) : %s", nextAvailableFilter, err)
		}

	} else if cidr != "" {
		nc, err = objMgr.CreateNetworkContainer(nvName, cidr, isIPv6, comment, extAttrs)
		if err != nil {
			return fmt.Errorf(
				"creation of IPv6 network container block in network view '%s' failed: %w",
				nvName, err)
		}
	} else {
		return fmt.Errorf("creation of network block failed: neither cidr nor parentCidr with allocate_prefix_len was specified")
	}

	d.SetId(nc.Ref)
	if err = d.Set("internal_id", internalId.String()); err != nil {
		return err
	}
	if err = d.Set("ref", nc.Ref); err != nil {
		return err
	}
	return nil
}

func resourceNetworkContainerRead(d *schema.ResourceData, m interface{}) error {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return fmt.Errorf("failed to read network containter: %w", err)
	}

	//var tenantID string
	//tempVal, found := extAttrs[eaNameForTenantId]
	//if found {
	//	tenantID = tempVal.(string)
	//}

	//connector := m.(ibclient.IBConnector)
	//objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	nc, err := searchObjectByRefOrInternalId("NetworkContainer", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		} else {
			d.SetId("")
			return nil
		}
	}
	var obj *ibclient.NetworkContainer
	recJson, _ := json.Marshal(nc)
	err = json.Unmarshal(recJson, &obj)
	if err != nil {
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

	if err = d.Set("network_view", obj.NetviewName); err != nil {
		return err
	}

	if err = d.Set("cidr", obj.Cidr); err != nil {
		return err
	}

	if err = d.Set("ref", obj.Ref); err != nil {
		return err
	}
	d.SetId(obj.Ref)

	return nil
}

func resourceNetworkContainerUpdate(d *schema.ResourceData, m interface{}) error {
	var updateSuccessful bool
	defer func() {
		// Reverting the state back, in case of a failure,
		// otherwise Terraform will keep the values, which leaded to the failure,
		// in the state file.
		if !updateSuccessful {
			prevNetView, _ := d.GetChange("network_view")
			prevCIDR, _ := d.GetChange("cidr")
			prevParCIDR, _ := d.GetChange("parent_cidr")
			prevPrefLen, _ := d.GetChange("allocate_prefix_len")
			prevNextAvailableFilter, _ := d.GetChange("filter_params")
			prevComment, _ := d.GetChange("comment")
			prevEa, _ := d.GetChange("ext_attrs")

			_ = d.Set("network_view", prevNetView.(string))
			_ = d.Set("cidr", prevCIDR.(string))
			_ = d.Set("parent_cidr", prevParCIDR.(string))
			_ = d.Set("allocate_prefix_len", prevPrefLen.(int))
			_ = d.Set("filter_params", prevNextAvailableFilter.(string))
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

	if d.HasChange("parent_cidr") {
		return fmt.Errorf("changing the value of 'parent_cidr' field is not allowed")
	}

	if d.HasChange("allocate_prefix_len") {
		return fmt.Errorf("changing the value of 'allocate_prefix_len' field is not allowed")
	}

	if d.HasChange("filter_params") {
		return fmt.Errorf("changing the value of 'filter_params' field is not allowed")
	}

	nvName := d.Get("network_view").(string)
	cidr := d.Get("cidr").(string)

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
	tempVal, found := newExtAttrs[eaNameForTenantId]
	if found {
		tenantID = tempVal.(string)
	}

	// Generate UUID for internal_id and add to the EA if it is not set
	internalId := d.Get("internal_id").(string)

	if internalId == "" {
		internalId = generateInternalId().String()
	}

	newInternalId := newInternalResourceIdFromString(internalId)
	newExtAttrs[eaNameForInternalId] = newInternalId.String()

	if cidr == "" || nvName == "" {
		return fmt.Errorf(
			"tenant ID, network view's name and CIDR are required to update a network container")
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	nc, err := objMgr.GetNetworkContainerByRef(d.Id())
	if err != nil {
		return fmt.Errorf("failed to read network container for update operation: %w", err)
	}

	newExtAttrs, err = mergeEAs(nc.Ea, newExtAttrs, oldExtAttrs, connector)
	if err != nil {
		return err
	}

	comment := ""
	commentText, commentFieldFound := d.GetOk("comment")
	if commentFieldFound {
		comment = commentText.(string)
	}

	nc, err = objMgr.UpdateNetworkContainer(d.Id(), newExtAttrs, comment)
	if err != nil {
		return fmt.Errorf(
			"failed to update the network container in network view '%s': %w",
			nvName, err)
	}
	updateSuccessful = true
	d.SetId(nc.Ref)
	if err = d.Set("internal_id", newInternalId.String()); err != nil {
		return err
	}
	if err = d.Set("ref", nc.Ref); err != nil {
		return err
	}
	return nil
}

func resourceNetworkContainerDelete(d *schema.ResourceData, m interface{}) error {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return fmt.Errorf("failed to delete network container: %w", err)
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

	obj, err := searchObjectByRefOrInternalId("NetworkContainer", d, m)
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
	var nc *ibclient.NetworkContainer
	recJson, _ := json.Marshal(obj)
	err = json.Unmarshal(recJson, &nc)
	if err != nil {
		return fmt.Errorf("failed to read network container for deletion: %w", err)
	}

	if _, err := objMgr.DeleteNetworkContainer(nc.Ref); err != nil {
		return fmt.Errorf(
			"deletion of the network container failed: %w", err)
	}

	return nil
}

// TODO: implement this after infoblox-go-client refactoring
//func resourceNetworkContainerExists(d *schema.ResourceData, m interface{}, isIPv6 bool) (bool, error) {
//	return false, nil
//}

func resourceIPv4NetworkContainerCreate(d *schema.ResourceData, m interface{}) error {
	return resourceNetworkContainerCreate(d, m, false)
}

func resourceIPv4NetworkContainerRead(d *schema.ResourceData, m interface{}) error {
	ref := d.Id()
	if !netContainerIPv4Regexp.MatchString(ref) {
		return fmt.Errorf("reference '%s' for 'networkcontainer' object has an invalid format", ref)
	}

	return resourceNetworkContainerRead(d, m)
}

func resourceIPv4NetworkContainerUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceNetworkContainerUpdate(d, m)
}

func resourceIPv4NetworkContainerDelete(d *schema.ResourceData, m interface{}) error {
	return resourceNetworkContainerDelete(d, m)
}

//func resourceIPv4NetworkContainerExists(d *schema.ResourceData, m interface{}) (bool, error) {
//	return resourceNetworkContainerExists(d, m, false)
//}

func resourceIPv4NetworkContainer() *schema.Resource {
	nc := resourceNetworkContainer()
	nc.Create = resourceIPv4NetworkContainerCreate
	nc.Read = resourceIPv4NetworkContainerRead
	nc.Update = resourceIPv4NetworkContainerUpdate
	nc.Delete = resourceIPv4NetworkContainerDelete
	//nc.Exists = resourceIPv4NetworkContainerExists

	return nc
}

func resourceIPv6NetworkContainerCreate(d *schema.ResourceData, m interface{}) error {
	return resourceNetworkContainerCreate(d, m, true)
}

func resourceIPv6NetworkContainerRead(d *schema.ResourceData, m interface{}) error {
	ref := d.Id()
	if !netContainerIPv6Regexp.MatchString(ref) {
		return fmt.Errorf("reference '%s' for 'ipv6networkcontainer' object has an invalid format", ref)
	}

	return resourceNetworkContainerRead(d, m)
}

func resourceIPv6NetworkContainerUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceNetworkContainerUpdate(d, m)
}

func resourceIPv6NetworkContainerDelete(d *schema.ResourceData, m interface{}) error {
	return resourceNetworkContainerDelete(d, m)
}

//func resourceIPv6NetworkContainerExists(d *schema.ResourceData, m interface{}) (bool, error) {
//	return resourceNetworkContainerExists(d, m, true)
//}

func resourceIPv6NetworkContainer() *schema.Resource {
	nc := resourceNetworkContainer()
	nc.Create = resourceIPv6NetworkContainerCreate
	nc.Read = resourceIPv6NetworkContainerRead
	nc.Update = resourceIPv6NetworkContainerUpdate
	nc.Delete = resourceIPv6NetworkContainerDelete
	//nc.Exists = resourceIPv6NetworkContainerExists

	return nc
}

func resourceNetworkContainerImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to read network containter: %w", err)
	}

	var tenantID string
	tempVal, found := extAttrs[eaNameForTenantId]
	if found {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	obj, err := objMgr.GetNetworkContainerByRef(d.Id())
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve network container: %w", err)
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

	if err = d.Set("network_view", obj.NetviewName); err != nil {
		return nil, err
	}

	if err = d.Set("cidr", obj.Cidr); err != nil {
		return nil, err
	}

	d.SetId(obj.Ref)

	// Set the Terraform Internal ID to the NIOS EA if it is not already set
	err = resourceNetworkContainerUpdate(d, m)
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
