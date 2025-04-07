package infoblox

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func diffSuppressServerAssociationType(k, old, new string, d *schema.ResourceData) bool {
	if old == "MEMBER" && d.Get("member") != nil {
		return true
	}
	return false
}

func resourceRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceRangeCreate,
		Read:   resourceRangeRead,
		Update: resourceRangeUpdate,
		Delete: resourceRangeDelete,
		Importer: &schema.ResourceImporter{
			State: resourceRangeImport,
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
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comment for the range; maximum 256 characters.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the range.",
			},
			"network": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The network to which this range belongs, in IPv4 Address/CIDR format.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// Suppress the diff if the new value is empty and the old value is not empty
					return new == "" && old != ""
				},
			},
			"network_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     defaultNetView,
				Description: "The name of the network view in which this range resides.",
			},
			"start_addr": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The IPv4 Address starting address of the range.",
			},
			"end_addr": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The IPv4 Address end address of the range.",
			},
			"disable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Determines whether a range is disabled or not. When this is set to False, the range is enabled.\n\n",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Extensible attributes of the range to be added/updated, as a map in JSON format.",
			},
			"failover_association": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "TThe name of the failover association: the server in this failover association will serve the IPv4 range in case the main server is out of service.",
			},
			"options": {
				Type:     schema.TypeList,
				Optional: true,
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
			"member": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The member that will provide service for this range.",
			},
			"use_options": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Use flag for options.",
			},
			"server_association_type": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "NONE",
				Description:      "The type of server that is going to serve the range. The valid values are: 'FAILOVER', 'MEMBER', 'NONE'.'MS_FAILOVER','MS_SERVER'",
				DiffSuppressFunc: diffSuppressServerAssociationType,
			},
			"template": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "If set on creation, the range will be created according to the values specified in the named template.",
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

func resourceRangeCreate(d *schema.ResourceData, m interface{}) error {
	// Check if internal_id is set manually
	if intId := d.Get("internal_id"); intId.(string) != "" {
		return fmt.Errorf("the value of 'internal_id' field must not be set manually")
	}
	comment := d.Get("comment").(string)
	name := d.Get("name").(string)
	network := d.Get("network").(string)
	networkView := d.Get("network_view").(string)
	if networkView == "" {
		networkView = defaultNetView
	}
	startAddr := d.Get("start_addr").(string)
	endAddr := d.Get("end_addr").(string)
	disable := d.Get("disable").(bool)
	extAttrJSON := d.Get("ext_attrs").(string)
	useOptions := d.Get("use_options").(bool)
	optionsInterface := d.Get("options").([]interface{})
	options := ConvertInterfaceToDhcpOptions(optionsInterface)
	serverAssociationType := d.Get("server_association_type").(string)
	failOverAssociation := d.Get("failover_association").(string)
	template := d.Get("template").(string)
	member := d.Get("member").(string)
	memberStructure, err := ConvertJSONToDhcpMember(member)
	if err != nil {
		return err
	}
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
	newNetworkRange, err := objMgr.CreateNetworkRange(comment, name, network, networkView, startAddr, endAddr, disable, extAttrs, memberStructure, failOverAssociation, options, useOptions, serverAssociationType, template)
	if err != nil {
		return err
	}
	d.SetId(newNetworkRange.Ref)
	if err = d.Set("ref", newNetworkRange.Ref); err != nil {
		return err
	}
	if err = d.Set("internal_id", internalId.String()); err != nil {
		return err
	}
	return resourceRangeRead(d, m)

}
func resourceRangeRead(d *schema.ResourceData, m interface{}) error {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}

	rec, err := searchObjectByRefOrInternalId("Range", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		} else {
			d.SetId("")
			return nil
		}
	}
	var networkRange *ibclient.Range
	recJson, err := json.Marshal(rec)
	if err != nil {
		return fmt.Errorf("failed to marshal network range: %s", err.Error())
	}
	err = json.Unmarshal(recJson, &networkRange)
	if err != nil {
		return fmt.Errorf("failed getting network range : %s", err.Error())
	}

	delete(networkRange.Ea, eaNameForInternalId)
	omittedEAs := omitEAs(networkRange.Ea, extAttrs)

	if omittedEAs != nil && len(omittedEAs) > 0 {
		eaJSON, err := terraformSerializeEAs(omittedEAs)
		if err != nil {
			return err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return err
		}
	}
	// Assertion of object type and error handling
	if err = d.Set("comment", networkRange.Comment); err != nil {
		return err
	}
	if err = d.Set("name", networkRange.Name); err != nil {
		return err
	}
	if err = d.Set("network", networkRange.Network); err != nil {
		return err
	}
	if err = d.Set("network_view", networkRange.NetworkView); err != nil {
		return err
	}
	if err = d.Set("start_addr", networkRange.StartAddr); err != nil {
		return err
	}
	if err = d.Set("end_addr", networkRange.EndAddr); err != nil {
		return err
	}
	if err = d.Set("disable", networkRange.Disable); err != nil {
		return err
	}
	if err = d.Set("failover_association", networkRange.FailoverAssociation); err != nil {
		return err
	}
	serializedMember, err := serializeDhcpMember(networkRange.Member)
	if err != nil {
		return err
	}
	if err = d.Set("member", serializedMember); err != nil {
		return err
	}
	if err = d.Set("server_association_type", networkRange.ServerAssociationType); err != nil {
		return err
	}
	if err = d.Set("options", convertDhcpOptionsToInterface(networkRange.Options)); err != nil {
		return err
	}
	if err = d.Set("use_options", networkRange.UseOptions); err != nil {
		return err
	}
	d.SetId(networkRange.Ref)
	return nil
}
func resourceRangeUpdate(d *schema.ResourceData, m interface{}) error {
	var updateSuccessful bool
	defer func() {
		// Reverting the state back, in case of a failure,
		// otherwise Terraform will keep the values, which leaded to the failure,
		// in the state file.
		if !updateSuccessful {
			prevComment, _ := d.GetChange("comment")
			prevName, _ := d.GetChange("name")
			prevNetwork, _ := d.GetChange("network")
			prevNetworkView, _ := d.GetChange("network_view")
			prevStartAddr, _ := d.GetChange("start_addr")
			prevEndAddr, _ := d.GetChange("end_addr")
			prevEa, _ := d.GetChange("ext_attrs")
			prevMember, _ := d.GetChange("member")
			prevDisable, _ := d.GetChange("disable")
			prevOptions, _ := d.GetChange("options")
			prevUseOptions, _ := d.GetChange("use_options")
			prevServerAssociationType, _ := d.GetChange("server_association_type")
			prevFailOverAssociation, _ := d.GetChange("failover_association")

			// TODO: move to the new Terraform plugin framework and
			// process all the errors instead of ignoring them here.
			_ = d.Set("name", prevName.(string))
			_ = d.Set("network", prevNetwork.(string))
			_ = d.Set("start_addr", prevStartAddr.(string))
			_ = d.Set("end_addr", prevEndAddr.(string))
			_ = d.Set("comment", prevComment.(string))
			_ = d.Set("ext_attrs", prevEa.(string))
			_ = d.Set("network_view", prevNetworkView.(string))
			_ = d.Set("member", prevMember)
			_ = d.Set("disable", prevDisable.(bool))
			_ = d.Set("options", prevOptions)
			_ = d.Set("use_options", prevUseOptions.(bool))
			_ = d.Set("server_association_type", prevServerAssociationType.(string))
			_ = d.Set("failover_association", prevFailOverAssociation.(string))
		}
	}()
	if d.HasChange("internal_id") {
		return fmt.Errorf("changing the value of 'internal_id' field is not allowed")
	}
	if d.HasChange("network_view") {
		return fmt.Errorf("changing the value of 'network_view' field is not allowed")
	}
	comment := d.Get("comment").(string)
	name := d.Get("name").(string)
	network := d.Get("network").(string)
	startAddr := d.Get("start_addr").(string)
	endAddr := d.Get("end_addr").(string)
	disable := d.Get("disable").(bool)
	useOptions := d.Get("use_options").(bool)
	optionsInterface := d.Get("options").([]interface{})
	options := ConvertInterfaceToDhcpOptions(optionsInterface)
	member := d.Get("member").(string)
	memberStructure, err := ConvertJSONToDhcpMember(member)
	if err != nil {
		return err
	}
	failoverAssociation := d.Get("failover_association").(string)
	serverAssociationType := d.Get("server_association_type").(string)
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

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	var networkRange *ibclient.Range

	rec, err := searchObjectByRefOrInternalId("Range", d, m)
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
		return fmt.Errorf("failed to marshal Network Range : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &networkRange)
	if err != nil {
		return fmt.Errorf("failed getting Network Range : %s", err.Error())
	}

	// If 'internal_id' is not set, then generate a new one and set it to the EA.
	internalId := d.Get("internal_id").(string)
	if internalId == "" {
		internalId = generateInternalId().String()
	}
	newInternalId := newInternalResourceIdFromString(internalId)
	newExtAttrs[eaNameForInternalId] = newInternalId.String()

	newExtAttrs, err = mergeEAs(networkRange.Ea, newExtAttrs, oldExtAttrs, connector)
	networkRange, err = objMgr.UpdateNetworkRange(d.Id(), comment, name, network, startAddr, endAddr, disable, newExtAttrs, memberStructure, failoverAssociation, options, useOptions, serverAssociationType)
	if err != nil {
		return fmt.Errorf("Failed to update network range with %s, ", err.Error())
	}

	updateSuccessful = true

	if err = d.Set("internal_id", newInternalId.String()); err != nil {
		return err
	}
	if err = d.Set("ref", networkRange.Ref); err != nil {
		return err
	}
	d.SetId(networkRange.Ref)
	return nil

}
func resourceRangeDelete(d *schema.ResourceData, m interface{}) error {
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

	rec, err := searchObjectByRefOrInternalId("Range", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		} else {
			d.SetId("")
			return nil
		}
	}

	var networkRange *ibclient.Range
	recJson, err := json.Marshal(rec)
	if err != nil {
		return fmt.Errorf("failed to marshal network range : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &networkRange)
	if err != nil {
		return err
	}
	_, err = objMgr.DeleteNetworkRange(networkRange.Ref)
	if err != nil {
		return fmt.Errorf("failed to delete network range : %s", err.Error())
	}

	return nil

}
func resourceRangeImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	extAttrJSON := d.Get("ext_attrs").(string)
	_, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return nil, err
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	networkRange, err := objMgr.GetNetworkRangeByRef(d.Id())
	if err != nil {
		return nil, fmt.Errorf("failed getting network range : %w", err)
	}

	if networkRange.Ea != nil && len(networkRange.Ea) > 0 {
		eaJSON, err := terraformSerializeEAs(networkRange.Ea)
		if err != nil {
			return nil, err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return nil, err
		}
	}
	delete(networkRange.Ea, eaNameForInternalId)
	if err = d.Set("comment", networkRange.Comment); err != nil {
		return nil, err
	}
	if err = d.Set("name", networkRange.Name); err != nil {
		return nil, err
	}
	if err = d.Set("network", networkRange.Network); err != nil {
		return nil, err
	}
	if err = d.Set("network_view", networkRange.NetworkView); err != nil {
		return nil, err
	}
	if err = d.Set("start_addr", networkRange.StartAddr); err != nil {
		return nil, err
	}
	if err = d.Set("end_addr", networkRange.EndAddr); err != nil {
		return nil, err
	}
	if err = d.Set("disable", networkRange.Disable); err != nil {
		return nil, err
	}
	if err = d.Set("failover_association", networkRange.FailoverAssociation); err != nil {
		return nil, err
	}
	serializedMember, err := serializeDhcpMember(networkRange.Member)
	if err != nil {
		return nil, err
	}
	if err = d.Set("member", serializedMember); err != nil {
		return nil, err
	}
	if err = d.Set("server_association_type", networkRange.ServerAssociationType); err != nil {
		return nil, err
	}
	if err = d.Set("options", convertDhcpOptionsToInterface(networkRange.Options)); err != nil {
		return nil, err
	}
	if err = d.Set("use_options", networkRange.UseOptions); err != nil {
		return nil, err
	}
	if err = d.Set("template", networkRange.Template); err != nil {
		return nil, err
	}
	d.SetId(networkRange.Ref)

	// Update the resource with the EA Terraform Internal ID
	err = resourceRangeUpdate(d, m)
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil

}
func ConvertJSONToDhcpMember(jsonStr string) (*ibclient.Dhcpmember, error) {
	if jsonStr == "" {
		return nil, nil
	}
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return nil, err
	}
	ipv4Addr, _ := data["ipv4addr"].(string)
	ipv6Addr, _ := data["ipv6addr"].(string)
	name, _ := data["name"].(string)
	member := &ibclient.Dhcpmember{
		Ipv4Addr: ipv4Addr,
		Ipv6Addr: ipv6Addr,
		Name:     name,
	}
	return member, nil
}
func ConvertInterfaceToDhcpOptions(optionsInterface []interface{}) []*ibclient.Dhcpoption {
	var options []*ibclient.Dhcpoption
	for _, optionInterface := range optionsInterface {
		option := optionInterface.(map[string]interface{})
		dhcpOption := &ibclient.Dhcpoption{
			Name:        option["name"].(string),
			Num:         uint32(option["num"].(int)),
			VendorClass: option["vendor_class"].(string),
			Value:       option["value"].(string),
			UseOption:   option["use_option"].(bool),
		}
		options = append(options, dhcpOption)
	}
	return options
}

func convertDhcpOptionsToInterface(options []*ibclient.Dhcpoption) []map[string]interface{} {
	optionsInterface := make([]map[string]interface{}, 0, len(options))
	for _, option := range options {
		optionMap := make(map[string]interface{})
		optionMap["name"] = option.Name
		optionMap["num"] = option.Num
		optionMap["vendor_class"] = option.VendorClass
		optionMap["value"] = option.Value
		optionMap["use_option"] = option.UseOption
		optionsInterface = append(optionsInterface, optionMap)
	}
	return optionsInterface
}

func serializeDhcpMember(member *ibclient.Dhcpmember) (string, error) {
	if member == nil {
		return "", nil
	}
	data, err := json.Marshal(member)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
