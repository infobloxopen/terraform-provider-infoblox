package infoblox

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"reflect"
	"regexp"
	"sort"
	"strings"
)

func resourceIpv4SharedNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpv4SharedNetworkCreate,
		Read:   resourceIpv4SharedNetworkRead,
		Update: resourceIpv4SharedNetworkUpdate,
		Delete: resourceIpv4SharedNetworkDelete,
		Importer: &schema.ResourceImporter{
			State: resourceIpv4SharedNetworkImport,
		},
		CustomizeDiff: func(context context.Context, d *schema.ResourceDiff, meta interface{}) error {
			if internalID := d.Get("internal_id"); internalID == "" || internalID == nil {
				err := d.SetNewComputed("internal_id")
				if err != nil {
					return err
				}
			}

			// Helper function to check if an option is the default value
			isDefault := func(opt map[string]interface{}) bool {
				return opt["name"] == "dhcp-lease-time" && opt["num"] == 51 && opt["use_option"] == false && opt["value"] == "43200" && opt["vendor_class"] == "DHCP"
			}

			// Check if newList contains dhcp-lease-time
			containsDhcpLeaseTime := func(opt map[string]interface{}) bool {
				if opt["name"] == "dhcp-lease-time" {
					return true
				}
				return false
			}

			// Get the old and new values for the options field
			oldVal, newVal := d.GetChange("options")

			// Ensure oldVal and newVal are not nil
			if oldVal == nil || newVal == nil {
				return nil
			}

			oldList, okOld := oldVal.([]interface{})
			newList, okNew := newVal.([]interface{})

			// Ensure type assertions are successful
			if !okOld || !okNew {
				return nil
			}

			// If newList is empty, set it to an empty list to clear old values
			if len(newList) == 0 {
				if err := d.SetNew("options", []interface{}{}); err != nil {
					return err
				}
				return nil
			}

			// Add default values to oldList if they are missing
			defaultOpt := map[string]interface{}{
				"name":         "dhcp-lease-time",
				"num":          51,
				"use_option":   false,
				"value":        "43200",
				"vendor_class": "DHCP",
			}

			hasDefaultOld := false
			oldListContainsDhcpLeaseTime := false
			for _, oldOpt := range oldList {
				if containsDhcpLeaseTime(oldOpt.(map[string]interface{})) {
					oldListContainsDhcpLeaseTime = true
				}
				if isDefault(oldOpt.(map[string]interface{})) {
					hasDefaultOld = true
					break
				}
			}

			if !hasDefaultOld && !oldListContainsDhcpLeaseTime {
				oldList = append(oldList, defaultOpt)
			}

			// Add default values to newList if they are missing
			hasDefaultNew := false
			newListContainsDhcpLeaseTime := false
			for _, newOpt := range newList {
				if containsDhcpLeaseTime(newOpt.(map[string]interface{})) {
					newListContainsDhcpLeaseTime = true
				}
				if isDefault(newOpt.(map[string]interface{})) {
					hasDefaultNew = true
					break
				}
			}

			if !hasDefaultNew && !newListContainsDhcpLeaseTime {
				newList = append(newList, defaultOpt)
			}

			// Filter out default values from both old and new lists for comparison
			filteredOldList := []interface{}{}
			for _, oldOpt := range oldList {
				oldOptMap, ok := oldOpt.(map[string]interface{})
				if ok && !isDefault(oldOptMap) {
					filteredOldList = append(filteredOldList, oldOpt)
				}
			}

			filteredNewList := []interface{}{}
			for _, newOpt := range newList {
				newOptMap, ok := newOpt.(map[string]interface{})
				if ok && !isDefault(newOptMap) {
					filteredNewList = append(filteredNewList, newOpt)
				}
			}

			// Compare the filtered lists
			if len(filteredOldList) != len(filteredNewList) {
				if err := d.SetNew("options", newList); err != nil {
					return err
				}
				return nil
			}

			for i := range filteredOldList {
				oldOptMap := filteredOldList[i].(map[string]interface{})
				newOptMap := filteredNewList[i].(map[string]interface{})
				// Iterate through newList and check if num is 0, find the corresponding old option and set the num attribute to its old value
				if newOptMap["num"] == 0 {
					newList[i].(map[string]interface{})["num"] = oldOptMap["num"]
				}
				if !reflect.DeepEqual(oldOptMap, newOptMap) {
					if err := d.SetNew("options", newList); err != nil {
						return err
					}
					return nil
				}
			}

			// Ensure that the plan shows changes when non-default values are removed
			if len(filteredNewList) == 0 && len(filteredOldList) > 0 {
				if err := d.SetNew("options", newList); err != nil {
					return err
				}
				return nil
			}

			// If no changes in non-default values, set newList to oldList to avoid showing plan diff
			if err := d.SetNew("options", oldList); err != nil {
				return err
			}

			return nil
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the IPv4 shared network object.",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The comment for the IPv4 shared network object.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Extensible attributes of the IPv4 Shared Network record to be added/updated, as a map in JSON format.",
			},
			"disable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "The disable flag for the IPv4 shared network object.",
			},
			"networks": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "A list of networks belonging to the shared network",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					oldVal, newVal := d.GetChange("networks")
					oldList := oldVal.([]interface{})
					newList := newVal.([]interface{})
					return compareNetworkReferences(oldList, newList)
				},
			},
			"network_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     defaultNetView,
				Description: "The name of the network view in which this shared network resides.",
			},
			"use_options": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Use flag for options.",
			},
			"options": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Description: "An array of DHCP option structs that lists the DHCP options associated with the object. An option sets the" +
					"value of a DHCP option that has been defined in an option space. DHCP options describe network configuration settings" +
					"and various services available on the network. These options occur as variable-length fields at the end of DHCP messages." +
					"When defining a DHCP option, at least a ‘name’ or a ‘num’ is required.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the DHCP option.",
						},
						"num": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The code of the DHCP option.",
						},
						"use_option": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
							Description: "Only applies to special options that are displayed separately from other options and have a use flag. " +
								"These options are: `routers`, `router-templates`, `domain-name-servers`, `domain-name`, `broadcast-address`, " +
								"`broadcast-address-offset`, `dhcp-lease-time`, `dhcp6.name-servers`",
						},
						"value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Value of the DHCP option.",
						},
						"vendor_class": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "DHCP",
							Description: "The name of the space this DHCP option is associated to.",
						},
					},
				},
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

// Helper function to compare network references
func compareNetworkReferences(oldList, newList []interface{}) bool {
	if len(oldList) != len(newList) {
		return false
	}
	oldNetworks := make([]string, len(oldList))
	newNetworks := make([]string, len(newList))
	oldCidrs := make([]string, 0)
	newCidrs := make([]string, 0)
	for i, v := range oldList {
		oldNetworks[i] = v.(string)
		oldCidrs = append(oldCidrs, extractCIDR(oldNetworks[i]))
	}
	for i, v := range newList {
		newNetworks[i] = v.(string)
		newCidrs = append(newCidrs, extractCIDR(newNetworks[i]))
	}
	sort.Strings(oldCidrs)
	sort.Strings(newCidrs)

	for i, _ := range oldCidrs {
		return oldCidrs[i] == newCidrs[i]
	}
	return true

}

func extractCIDR(network string) string {
	// Regular expression to match CIDR format
	isCIDR := regexp.MustCompile(`^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\/\d{1,2}$`)

	// If the input is already a CIDR, return it as is
	if isCIDR.MatchString(network) {
		return network
	}

	// Otherwise, extract the CIDR part from the network reference
	parts := strings.Split(network, ":")
	if len(parts) > 1 {
		cidrParts := strings.Split(parts[1], "/")
		if len(cidrParts) > 1 {
			return cidrParts[0] + "/" + cidrParts[1]
		}
	}
	return network
}

func resourceIpv4SharedNetworkCreate(d *schema.ResourceData, m interface{}) error {

	if intId := d.Get("internal_id"); intId.(string) != "" {
		return fmt.Errorf("the value of 'internal_id' field must not be set manually")
	}
	name := d.Get("name").(string)
	comment := d.Get("comment").(string)
	disable := d.Get("disable").(bool)

	networks := d.Get("networks").([]interface{})
	networksList := make([]string, len(networks))
	for i, network := range networks {
		networksList[i] = network.(string)
	}

	networkView := d.Get("network_view").(string)
	useOptions := d.Get("use_options").(bool)
	options := d.Get("options").([]interface{})
	optionsList, err := validateDhcpOptions(options)
	if err != nil {
		return fmt.Errorf("failed to validate options: %w", err)
	}

	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}
	// Generate internal ID and add it to the extensible attributes
	internalId := generateInternalId()
	extAttrs[eaNameForInternalId] = internalId.String()

	var tenantID string
	if tempVal, found := extAttrs[eaNameForTenantId]; found {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	// create a sharedNetwork object
	sharedNetwork, err := objMgr.CreateIpv4SharedNetwork(name, networksList, networkView, extAttrs, comment, disable, useOptions, optionsList)
	if err != nil {
		return fmt.Errorf("failed to create a sharedNetwork object: %s", err)
	}
	d.SetId(sharedNetwork.Ref)
	if err = d.Set("internal_id", internalId.String()); err != nil {
		return err
	}
	if err = d.Set("ref", sharedNetwork.Ref); err != nil {
		return err
	}

	return resourceIpv4SharedNetworkRead(d, m)
}

func resourceIpv4SharedNetworkRead(d *schema.ResourceData, m interface{}) error {

	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}

	rec, err := searchObjectByRefOrInternalId("SharedNetwork", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); ok {
			d.SetId("")
			return nil
		} else {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		}
	}

	var sharedNetwork *ibclient.SharedNetwork
	recJson, err := json.Marshal(rec)
	if err != nil {
		return fmt.Errorf("failed to marshal shared network record : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &sharedNetwork)
	if err != nil {
		return fmt.Errorf("failed getting shared network record : %s", err.Error())
	}

	delete(sharedNetwork.Ea, eaNameForInternalId)
	omittedEAs := omitEAs(sharedNetwork.Ea, extAttrs)

	if omittedEAs != nil && len(omittedEAs) > 0 {
		eaJSON, err := terraformSerializeEAs(omittedEAs)
		if err != nil {
			return err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return err
		}
	}
	if sharedNetwork.Name != nil {
		if err = d.Set("name", *sharedNetwork.Name); err != nil {
			return err
		}
	}
	if sharedNetwork.Comment != nil {
		if err = d.Set("comment", *sharedNetwork.Comment); err != nil {
			return err
		}
	}
	if sharedNetwork.Disable != nil {
		if err = d.Set("disable", *sharedNetwork.Disable); err != nil {
			return err
		}
	}
	if sharedNetwork.Networks != nil {
		networks := setNetworksRef(sharedNetwork.Networks)
		if err = d.Set("networks", networks); err != nil {
			return err
		}
	}
	if err = d.Set("network_view", sharedNetwork.NetworkView); err != nil {
		return err
	}
	if sharedNetwork.UseOptions != nil {
		if err = d.Set("use_options", *sharedNetwork.UseOptions); err != nil {
			return err
		}
	}
	if sharedNetwork.Options != nil {
		networksInterface := convertDhcpOptionsToInterface(sharedNetwork.Options)
		if err = d.Set("options", networksInterface); err != nil {
			return err
		}
	}

	if err = d.Set("ref", sharedNetwork.Ref); err != nil {
		return err
	}
	d.SetId(sharedNetwork.Ref)
	return nil
}

func setNetworksRef(networks []*ibclient.Ipv4Network) interface{} {
	if len(networks) == 0 {
		return nil
	}
	ipv4Networks := make([]interface{}, len(networks))
	for i, network := range networks {
		ipv4Networks[i] = network.Ref
	}
	return ipv4Networks
}

func resourceIpv4SharedNetworkUpdate(d *schema.ResourceData, m interface{}) error {
	var updateSuccessful bool
	defer func() {
		if !updateSuccessful {
			prevName, _ := d.GetChange("name")
			prevComment, _ := d.GetChange("comment")
			prevDisable, _ := d.GetChange("disable")
			prevNetworks, _ := d.GetChange("networks")
			prevNetworkView, _ := d.GetChange("network_view")
			prevUseOptions, _ := d.GetChange("use_options")
			prevOptions, _ := d.GetChange("options")
			prevExtAttrs, _ := d.GetChange("ext_attrs")

			_ = d.Set("name", prevName)
			_ = d.Set("comment", prevComment)
			_ = d.Set("disable", prevDisable)
			_ = d.Set("networks", prevNetworks)
			_ = d.Set("network_view", prevNetworkView)
			_ = d.Set("use_options", prevUseOptions)
			_ = d.Set("options", prevOptions)
			_ = d.Set("ext_attrs", prevExtAttrs)
		}
	}()

	if d.HasChange("network_view") {
		return fmt.Errorf("changing the value of 'network_view' field is not allowed")
	}

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
	if tempVal, found := newExtAttrs[eaNameForTenantId]; found {
		tenantID = tempVal.(string)
	}

	name := d.Get("name").(string)
	comment := d.Get("comment").(string)
	disable := d.Get("disable").(bool)
	networks := d.Get("networks").([]interface{})
	networksList := make([]string, len(networks))
	for i, network := range networks {
		networksList[i] = network.(string)
	}

	networkView := d.Get("network_view").(string)
	useOptions := d.Get("use_options").(bool)
	options := d.Get("options").([]interface{})
	optionsList, err := validateDhcpOptions(options)
	if err != nil {
		return fmt.Errorf("failed to validate options: %w", err)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)
	var sharedNetwork *ibclient.SharedNetwork

	rec, err := searchObjectByRefOrInternalId("SharedNetwork", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		} else {
			d.SetId("")
			return nil
		}
	}
	recJson, err := json.Marshal(rec)
	if err != nil {
		return fmt.Errorf("failed to marshal sharedNetwork record : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &sharedNetwork)
	if err != nil {
		return fmt.Errorf("failed getting sharedNetwork record : %s", err.Error())
	}

	// If 'internal_id' is not set, then generate a new one and set it to the EA.
	internalId := d.Get("internal_id").(string)
	if internalId == "" {
		internalId = generateInternalId().String()
	}
	newInternalId := newInternalResourceIdFromString(internalId)
	newExtAttrs[eaNameForInternalId] = newInternalId.String()

	newExtAttrs, err = mergeEAs(sharedNetwork.Ea, newExtAttrs, oldExtAttrs, connector)
	if err != nil {
		return err
	}

	sharedNetwork, err = objMgr.UpdateIpv4SharedNetwork(d.Id(), name, networksList, networkView, comment, newExtAttrs, disable, useOptions, optionsList)
	if err != nil {
		return fmt.Errorf("failed to update sharedNetwork: %s.", err.Error())
	}

	updateSuccessful = true

	if err = d.Set("internal_id", newInternalId.String()); err != nil {
		return err
	}
	if err = d.Set("ref", sharedNetwork.Ref); err != nil {
		return err
	}
	d.SetId(sharedNetwork.Ref)
	return resourceIpv4SharedNetworkRead(d, m)
}

func resourceIpv4SharedNetworkDelete(d *schema.ResourceData, m interface{}) error {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}

	var tenantID string
	tempVal, found := extAttrs[eaNameForTenantId]
	if found {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	rec, err := searchObjectByRefOrInternalId("SharedNetwork", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		} else {
			d.SetId("")
			return nil
		}
	}

	var sharedNetwork *ibclient.SharedNetwork
	recJson, err := json.Marshal(rec)
	if err != nil {
		return fmt.Errorf("failed to marshal shared network record : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &sharedNetwork)
	if err != nil {
		return err
	}
	_, err = objMgr.DeleteIpv4SharedNetwork(sharedNetwork.Ref)
	if err != nil {
		return fmt.Errorf("failed to delete shared network : %s", err.Error())
	}

	return nil
}

func resourceIpv4SharedNetworkImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	extAttrJSON := d.Get("ext_attrs").(string)
	_, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return nil, err
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	sharedNetwork, err := objMgr.GetIpv4SharedNetworkByRef(d.Id())
	if err != nil {
		return nil, fmt.Errorf("failed getting shared network record: %w", err)
	}

	if sharedNetwork.Ea != nil && len(sharedNetwork.Ea) > 0 {
		eaJSON, err := terraformSerializeEAs(sharedNetwork.Ea)
		if err != nil {
			return nil, err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return nil, err
		}
	}
	delete(sharedNetwork.Ea, eaNameForInternalId)

	if sharedNetwork.Name != nil {
		if err = d.Set("name", *sharedNetwork.Name); err != nil {
			return nil, err
		}
	}
	if sharedNetwork.Comment != nil {
		if err = d.Set("comment", *sharedNetwork.Comment); err != nil {
			return nil, err
		}
	}
	if sharedNetwork.Disable != nil {
		if err = d.Set("disable", *sharedNetwork.Disable); err != nil {
			return nil, err
		}
	}
	if sharedNetwork.Networks != nil {
		networks := setNetworksRef(sharedNetwork.Networks)
		if err = d.Set("networks", networks); err != nil {
			return nil, err
		}
	}
	if err = d.Set("network_view", sharedNetwork.NetworkView); err != nil {
		return nil, err
	}
	if sharedNetwork.UseOptions != nil {
		if err = d.Set("use_options", *sharedNetwork.UseOptions); err != nil {
			return nil, err
		}
	}
	if sharedNetwork.Options != nil {
		networksInterface := convertDhcpOptionsToInterface(sharedNetwork.Options)
		if err = d.Set("options", networksInterface); err != nil {
			return nil, err
		}
	}

	if err = d.Set("ref", sharedNetwork.Ref); err != nil {
		return nil, err
	}

	d.SetId(sharedNetwork.Ref)

	// Update the resource with the EA Terraform Internal ID
	err = resourceIpv4SharedNetworkUpdate(d, m)
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func validateDhcpOptions(dhcpOptions []interface{}) ([]*ibclient.Dhcpoption, error) {
	if dhcpOptions == nil {
		return nil, nil
	}
	dhcpOptionsList := make([]*ibclient.Dhcpoption, 0, len(dhcpOptions))
	for _, option := range dhcpOptions {
		// Assert the type of dhcpOption to map[string]interface{}
		dhcpOptionsMap, ok := option.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("dhcpOption is not of type map[string]interface{}")
		}

		// Create a new dhcpOption and populate its fields
		dhcpOption := ibclient.Dhcpoption{}
		if name, ok := dhcpOptionsMap["name"].(string); ok {
			dhcpOption.Name = name
		}
		if value, ok := dhcpOptionsMap["value"].(string); ok {
			dhcpOption.Value = value
		}
		if num, ok := dhcpOptionsMap["num"].(int); ok {
			dhcpOption.Num = uint32(num)
		}
		if vendorClass, ok := dhcpOptionsMap["vendor_class"].(string); ok {
			dhcpOption.VendorClass = vendorClass
		}
		if useOption, ok := dhcpOptionsMap["use_option"].(bool); ok {
			dhcpOption.UseOption = useOption
		}
		dhcpOptionsList = append(dhcpOptionsList, &dhcpOption)
	}

	return dhcpOptionsList, nil
}

func convertDhcpOptionsToInterface(dhcpOptions []*ibclient.Dhcpoption) []map[string]interface{} {
	options := make([]map[string]interface{}, 0, len(dhcpOptions))
	for _, option := range dhcpOptions {
		sMap := make(map[string]interface{})
		sMap["name"] = option.Name
		sMap["num"] = option.Num
		sMap["value"] = option.Value
		sMap["vendor_class"] = option.VendorClass
		sMap["use_option"] = option.UseOption
		options = append(options, sMap)
	}
	return options
}
