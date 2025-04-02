package infoblox

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"strings"
)

func resourceFixedRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceFixedRecordCreate,
		Read:   resourceFixedRecordRead,
		Update: resourceFixedRecordUpdate,
		Delete: resourceFixedRecordDelete,
		Importer: &schema.ResourceImporter{
			State: resourceFixedRecordImport,
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
			"agent_circuit_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The agent circuit ID for the fixed address.",
			},
			"agent_remote_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The agent remote ID for the fixed address.",
			},
			"client_identifier_prepend_zero": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "This field controls whether there is a prepend for the dhcp-client-identifier of a fixed address.",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comment for the fixed address; maximum 256 characters.",
			},
			"dhcp_client_identifier": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The DHCP client ID for the fixed address. The field is required only when match_client is set to CLIENT_ID.",
			},
			"disable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Determines whether a fixed address is disabled or not. When this is set to False, the fixed address is enabled.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Extensible attributes of the A-record to be added/updated, as a map in JSON format",
			},
			"ipv4addr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The IPv4 Address of the fixed address.",
				DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
					oldNetwork, _ := d.GetChange("network")
					if oldValue != "" && newValue == "" && oldNetwork != "" {
						return true
					}
					return false
				},
			},
			"mac": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The MAC address value for this fixed address.",
				DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
					oldValue = strings.ToUpper(oldValue)
					newValue = strings.ToUpper(newValue)
					if d.Get("match_client").(string) != "MAC_ADDRESS" && newValue != "" {
						return true
					}
					// Suppress diff if MAC addresses match after normalization
					return oldValue == newValue
				},
			},
			"match_client": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "MAC_ADDRESS",
				Description: "The match client for the fixed address.Valid values are CIRCUIT_ID, CLIENT_ID , MAC_ADDRESS, REMOTE_ID and RESERVED",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "This field contains the name of this fixed address.",
			},
			"network": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The network to which this fixed address belongs, in IPv4 Address/CIDR format.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if d.Get("ipv4addr").(string) != "" && new == "" {
						return true
					}
					return false
				},
			},
			"network_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     defaultNetView,
				Description: "The name of the network view in which this fixed address resides.",
			},
			"options": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "An array of DHCP option structs that lists the DHCP options associated with the object.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the DHCP option.",
						},
						"num": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The code of the DHCP option.",
						},
						"use_option": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Only applies to special options that are displayed separately from other options and have a use flag.",
						},
						"value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Value of the DHCP option",
						},
						"vendor_class": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "DHCP",
							Description: "The name of the space this DHCP option is associated to.",
						},
					},
				},
				DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
					if newValue == "0" {
						return false
					}
					oldList, newList := d.GetChange("options")
					return CompareSortedList(oldList, newList, "name", "num")
				},
			},
			"use_option": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Use option is a flag that indicates whether the options field are used or not.",
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
func resourceFixedRecordCreate(d *schema.ResourceData, m interface{}) error {
	// Check if internal_id is set manually
	if intId := d.Get("internal_id"); intId.(string) != "" {
		return fmt.Errorf("the value of 'internal_id' field must not be set manually")
	}
	agentCircuitId := d.Get("agent_circuit_id").(string)
	agentRemoteId := d.Get("agent_remote_id").(string)
	clientIdentifierPrependZero := d.Get("client_identifier_prepend_zero").(bool)
	comment := d.Get("comment").(string)
	dhcpClientIdentifier := d.Get("dhcp_client_identifier").(string)
	disable := d.Get("disable").(bool)
	ipAddr := d.Get("ipv4addr").(string)
	mac := d.Get("mac").(string)
	matchClient := d.Get("match_client").(string)
	if matchClient == "MAC_ADDRESS" && mac == "" {
		return fmt.Errorf("MAC address is required when match_client set to MAC_ADDRESS")
	}
	name := d.Get("name").(string)
	network := d.Get("network").(string)
	if ipAddr == "" && network == "" {
		return fmt.Errorf("either 'ipv4addr' or 'network' fields needs to provided to allocate a fixed address")
	}
	networkView := d.Get("network_view").(string)

	optionsInterface := d.Get("options").([]interface{})
	options := ConvertInterfaceToDhcpOptions(optionsInterface)
	useOptions := d.Get("use_option").(bool)
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

	fixedAddress, err := objMgr.AllocateIP(networkView, network, ipAddr, false, mac, name, comment, extAttrs, matchClient, agentCircuitId, agentRemoteId, clientIdentifierPrependZero, dhcpClientIdentifier, disable, options, useOptions)
	if err != nil {
		return err
	}
	d.SetId(fixedAddress.Ref)
	if err = d.Set("internal_id", internalId.String()); err != nil {
		return err
	}
	if err = d.Set("ref", fixedAddress.Ref); err != nil {
		return err
	}
	return resourceFixedRecordRead(d, m)
}
func resourceFixedRecordRead(d *schema.ResourceData, m interface{}) error {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}

	rec, err := searchObjectByRefOrInternalId("FixedAddress", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		} else {
			d.SetId("")
			return nil
		}
	}
	var fixedAddress *ibclient.FixedAddress
	recJson, err := json.Marshal(rec)
	if err != nil {
		return fmt.Errorf("failed to marshal fixed address : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &fixedAddress)
	if err != nil {
		return fmt.Errorf("failed getting fixed address : %s", err.Error())
	}

	delete(fixedAddress.Ea, eaNameForInternalId)
	omittedEAs := omitEAs(fixedAddress.Ea, extAttrs)

	if omittedEAs != nil && len(omittedEAs) > 0 {
		eaJSON, err := terraformSerializeEAs(omittedEAs)
		if err != nil {
			return err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return err
		}
	}
	if err = d.Set("comment", fixedAddress.Comment); err != nil {
		return err
	}
	if err = d.Set("disable", fixedAddress.Disable); err != nil {
		return err
	}
	if err = d.Set("ipv4addr", fixedAddress.IPv4Address); err != nil {
		return err
	}
	if err = d.Set("mac", fixedAddress.Mac); err != nil {
		return err
	}
	if err = d.Set("match_client", fixedAddress.MatchClient); err != nil {
		return err
	}
	if err = d.Set("name", fixedAddress.Name); err != nil {
		return err
	}
	if err = d.Set("network", fixedAddress.Cidr); err != nil {
		return err
	}
	if err = d.Set("network_view", fixedAddress.NetviewName); err != nil {
		return err
	}

	if err = d.Set("agent_circuit_id", fixedAddress.AgentCircuitId); err != nil {
		return err
	}
	if err = d.Set("agent_remote_id", fixedAddress.AgentRemoteId); err != nil {
		return err
	}
	if err = d.Set("client_identifier_prepend_zero", fixedAddress.ClientIdentifierPrependZero); err != nil {
		return err
	}
	if err = d.Set("dhcp_client_identifier", fixedAddress.DhcpClientIdentifier); err != nil {
		return err
	}
	if err = d.Set("use_option", fixedAddress.UseOptions); err != nil {
		return err
	}
	if err = d.Set("options", convertDhcpOptionsToInterface(fixedAddress.Options)); err != nil {
		return err
	}
	d.SetId(fixedAddress.Ref)
	return nil
}
func resourceFixedRecordUpdate(d *schema.ResourceData, m interface{}) error {
	var updateSuccessful bool
	defer func() {
		if !updateSuccessful {
			// Reverting the state back, in case of a failure,
			// otherwise Terraform will keep the values, which leaded to the failure,
			// in the state file.
			prevComment, _ := d.GetChange("comment")
			prevDisable, _ := d.GetChange("disable")
			prevIpv4addr, _ := d.GetChange("ipv4addr")
			prevMac, _ := d.GetChange("mac")
			prevMatchClient, _ := d.GetChange("match_client")
			prevName, _ := d.GetChange("name")
			prevNetwork, _ := d.GetChange("network")
			prevNetworkView, _ := d.GetChange("network_view")
			prevAgentCircuitId, _ := d.GetChange("agent_circuit_id")
			prevAgentRemoteId, _ := d.GetChange("agent_remote_id")
			prevClientIdentifierPrependZero, _ := d.GetChange("client_identifier_prepend_zero")
			prevDhcpClientIdentifier, _ := d.GetChange("dhcp_client_identifier")
			prevUseOption, _ := d.GetChange("use_option")
			prevOptions, _ := d.GetChange("options")
			prevExtAttrsJSON, _ := d.GetChange("ext_attrs")

			_ = d.Set("comment", prevComment.(string))
			_ = d.Set("disable", prevDisable.(bool))
			_ = d.Set("ipv4addr", prevIpv4addr.(string))
			_ = d.Set("mac", prevMac.(string))
			_ = d.Set("match_client", prevMatchClient.(string))
			_ = d.Set("name", prevName.(string))
			_ = d.Set("network", prevNetwork.(string))
			_ = d.Set("network_view", prevNetworkView.(string))
			_ = d.Set("agent_circuit_id", prevAgentCircuitId.(string))
			_ = d.Set("agent_remote_id", prevAgentRemoteId.(string))
			_ = d.Set("client_identifier_prepend_zero", prevClientIdentifierPrependZero.(bool))
			_ = d.Set("dhcp_client_identifier", prevDhcpClientIdentifier.(string))
			_ = d.Set("use_option", prevUseOption.(bool))
			_ = d.Set("options", prevOptions)
			_ = d.Set("ext_attrs", prevExtAttrsJSON.(string))

		}
	}()
	if d.HasChange("internal_id") {
		return fmt.Errorf("changing the value of 'internal_id' field is not allowed")
	}
	if d.HasChange("network_view") {
		return fmt.Errorf("changing the value of 'network_view' field is not allowed")
	}

	network := d.Get("network").(string)
	ipv4addr := d.Get("ipv4addr").(string)
	mac := d.Get("mac").(string)
	name := d.Get("name").(string)
	comment := d.Get("comment").(string)
	disable := d.Get("disable").(bool)
	matchClient := d.Get("match_client").(string)
	agentCircuitId := d.Get("agent_circuit_id").(string)
	agentRemoteId := d.Get("agent_remote_id").(string)
	clientIdentifierPrependZero := d.Get("client_identifier_prepend_zero").(bool)
	dhcpClientIdentifier := d.Get("dhcp_client_identifier").(string)
	useOptions := d.Get("use_option").(bool)
	optionsInterface := d.Get("options").([]interface{})
	options := ConvertInterfaceToDhcpOptions(optionsInterface)
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

	var fixedAddress *ibclient.FixedAddress

	rec, err := searchObjectByRefOrInternalId("FixedAddress", d, m)
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
		return fmt.Errorf("failed to marshal fixedAddress : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &fixedAddress)
	if err != nil {
		return fmt.Errorf("failed getting fixed addresss: %s", err.Error())
	}

	// If 'internal_id' is not set, then generate a new one and set it to the EA.
	internalId := d.Get("internal_id").(string)
	if internalId == "" {
		internalId = generateInternalId().String()
	}
	newInternalId := newInternalResourceIdFromString(internalId)
	newExtAttrs[eaNameForInternalId] = newInternalId.String()

	newExtAttrs, err = mergeEAs(fixedAddress.Ea, newExtAttrs, oldExtAttrs, connector)

	fixedAddress, err = objMgr.UpdateFixedAddress(d.Id(), "", name, network, ipv4addr, matchClient, mac, comment, newExtAttrs, agentCircuitId, agentRemoteId, clientIdentifierPrependZero, dhcpClientIdentifier, disable, options, useOptions)
	if err != nil {
		return fmt.Errorf("error updating Fixed address: %w", err)
	}
	updateSuccessful = true
	d.SetId(fixedAddress.Ref)
	if err = d.Set("ref", fixedAddress.Ref); err != nil {
		return err
	}
	if err = d.Set("internal_id", newInternalId.String()); err != nil {
		return err
	}
	return resourceFixedRecordRead(d, m)
}
func resourceFixedRecordDelete(d *schema.ResourceData, m interface{}) error {
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

	rec, err := searchObjectByRefOrInternalId("FixedAddress", d, m)
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
	var fixedAddress *ibclient.FixedAddress
	recJson, _ := json.Marshal(rec)
	err = json.Unmarshal(recJson, &fixedAddress)

	_, err = objMgr.DeleteARecord(fixedAddress.Ref)
	if err != nil {
		return fmt.Errorf("deletion of Fixed address failed: %w", err)
	}
	d.SetId("")

	return nil
}

func resourceFixedRecordImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return nil, err
	}

	var tenantID string
	tempVal, found := extAttrs[eaNameForTenantId]
	if found {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	obj, err := objMgr.GetFixedAddressByRef(d.Id())
	if err != nil {
		return nil, fmt.Errorf("failed getting fixed address: %w", err)
	}
	if err = d.Set("comment", obj.Comment); err != nil {
		return nil, err
	}
	if err = d.Set("disable", obj.Disable); err != nil {
		return nil, err
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
	if err = d.Set("ipv4addr", obj.IPv4Address); err != nil {
		return nil, err
	}
	if err = d.Set("mac", obj.Mac); err != nil {
		return nil, err
	}
	if err = d.Set("match_client", obj.MatchClient); err != nil {
		return nil, err
	}
	if err = d.Set("name", obj.Name); err != nil {
		return nil, err
	}
	if err = d.Set("agent_circuit_id", obj.AgentCircuitId); err != nil {
		return nil, err
	}
	if err = d.Set("agent_remote_id", obj.AgentRemoteId); err != nil {
		return nil, err
	}
	if err = d.Set("network", obj.Cidr); err != nil {
		return nil, err
	}
	if err = d.Set("network_view", obj.NetviewName); err != nil {
		return nil, err
	}
	if err = d.Set("client_identifier_prepend_zero", obj.ClientIdentifierPrependZero); err != nil {
		return nil, err
	}
	if err = d.Set("dhcp_client_identifier", obj.DhcpClientIdentifier); err != nil {
		return nil, err
	}
	if err = d.Set("use_option", obj.UseOptions); err != nil {
		return nil, err
	}
	OptionsInterface := convertDhcpOptionsToInterface(obj.Options)
	if err = d.Set("options", OptionsInterface); err != nil {
		return nil, err
	}
	d.SetId(obj.Ref)

	// Update the resource with the EA Terraform Internal ID
	err = resourceFixedRecordUpdate(d, m)
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
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
