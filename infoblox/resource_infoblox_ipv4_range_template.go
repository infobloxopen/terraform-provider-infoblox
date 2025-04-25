package infoblox

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"reflect"
)

func resourceRangeTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceRangeTemplateCreate,
		Read:   resourceRangeTemplateRead,
		Update: resourceRangeTemplateUpdate,
		Delete: resourceRangeTemplateDelete,
		Importer: &schema.ResourceImporter{
			State: resourceRangeTemplateImport,
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
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the Range Template record.",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comment for the Range Template record.",
			},
			"number_of_addresses": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The number of addresses for this range.",
			},
			"offset": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The start address offset for the range.",
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
							Default:  false,
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
				DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
					if newValue == "0" && oldValue >= "1" {
						return false
					}
					oldOptions, newOptions := d.GetChange("options")
					oldList, okOld := oldOptions.([]interface{})
					newList, okNew := newOptions.([]interface{})
					// Ensure type assertions are successful
					if !okOld || !okNew {
						return false
					}

					sortOptions(oldList, "name")
					sortOptions(newList, "name")
					//return reflect.DeepEqual(oldOptions, newOptions)

					// Filter out default values from both old and new lists
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
						return false
					}

					// Check for changes in subfields of default elements
					if len(filteredOldList) == len(filteredNewList) {
						for i := range filteredOldList {
							oldOptMap := filteredOldList[i].(map[string]interface{})
							newOptMap := filteredNewList[i].(map[string]interface{})
							if !reflect.DeepEqual(oldOptMap, newOptMap) {
								return false
							}
						}
					}
					return true
				},
			},
			"server_association_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NONE",
				Description: "The type of server that is going to serve the range. Valid values are: `FAILOVER`, `MEMBER`, `MS_FAILOVER`, " +
					"`MS_SERVER`, `NONE`. Default value is `NONE`",
			},
			"failover_association": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "The name of the failover association: the server in this failover association will serve the IPv4 range in case the " +
					"main server is out of service. `server_association_type` must be set to ‘FAILOVER’ or ‘FAILOVER_MS’ if you want the " +
					"failover association specified here to serve the range.",
			},
			"member": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The member that will provide service for this range. server_association_type needs to be set to ‘MEMBER’ if you want" +
					"the server specified here to serve the range.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Extensible attributes of the  Range Template Record to be added/updated, as a map in JSON format",
			},
			"cloud_api_compatible": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "This flag controls whether this template can be used to create network objects in a cloud-computing deployment.",
			},
			"ms_server": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "The Microsoft server that will provide service for this range. `server_association_type` needs to be set to `MS_SERVER` +" +
					"if you want the server specified here to serve the range. For searching by this field you should use a HTTP method that contains a" +
					"body (POST or PUT) with MS DHCP server structure and the request should have option _method=GET.",
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

func resourceRangeTemplateCreate(d *schema.ResourceData, m interface{}) error {

	if intId := d.Get("internal_id"); intId.(string) != "" {
		return fmt.Errorf("the value of 'internal_id' field must not be set manually")
	}
	name := d.Get("name").(string)
	numberOfAddresses := d.Get("number_of_addresses").(int)
	offset := d.Get("offset").(int)
	comment := d.Get("comment").(string)
	useOptions := d.Get("use_options").(bool)
	options := d.Get("options").([]interface{})
	optionsList, err := validateDhcpOptions(options)
	if err != nil {
		return fmt.Errorf("failed to validate options: %w", err)
	}

	serverAssociationType := d.Get("server_association_type").(string)
	failoverAssociation := d.Get("failover_association").(string)
	member := d.Get("member").(map[string]interface{})
	cloudApiCompatible := d.Get("cloud_api_compatible").(bool)
	msServer := d.Get("ms_server").(string)
	dhcpMemeber, err := ConvertMapToDhcpMember(member)
	if err != nil {
		return fmt.Errorf("failed to convert member to dhcpmember: %w", err)
	}
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return fmt.Errorf("failed to allocate IP: %w", err)
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

	// Create the Range Template record
	newRecord, err := objMgr.CreateRangeTemplate(name, uint32(numberOfAddresses), uint32(offset), comment, extAttrs, optionsList, useOptions, serverAssociationType, failoverAssociation, dhcpMemeber, cloudApiCompatible, msServer)
	if err != nil {
		return fmt.Errorf("failed to create Range Template record: %w", err)
	}
	d.SetId(newRecord.Ref)
	if err = d.Set("internal_id", internalId.String()); err != nil {
		return err
	}
	if err = d.Set("ref", newRecord.Ref); err != nil {
		return err
	}

	return resourceRangeTemplateRead(d, m)
}

func resourceRangeTemplateRead(d *schema.ResourceData, m interface{}) error {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}
	rec, err := searchObjectByRefOrInternalId("RangeTemplate", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); ok {
			d.SetId("")
			return nil
		} else {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		}
	}

	var rangeTemplate *ibclient.Rangetemplate
	recJson, err := json.Marshal(rec)
	if err != nil {
		return fmt.Errorf("failed to marshal Range Template record : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &rangeTemplate)
	if err != nil {
		return fmt.Errorf("failed getting Range Template record : %s", err.Error())
	}

	delete(rangeTemplate.Ea, eaNameForInternalId)
	omittedEAs := omitEAs(rangeTemplate.Ea, extAttrs)

	if omittedEAs != nil && len(omittedEAs) > 0 {
		eaJSON, err := terraformSerializeEAs(omittedEAs)
		if err != nil {
			return err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return err
		}
	}
	if rangeTemplate.Name != nil {
		if err = d.Set("name", *rangeTemplate.Name); err != nil {
			return err
		}
	}
	if rangeTemplate.Comment != nil {
		if err = d.Set("comment", *rangeTemplate.Comment); err != nil {
			return err
		}
	}
	if rangeTemplate.NumberOfAddresses != nil {
		if err = d.Set("number_of_addresses", *rangeTemplate.NumberOfAddresses); err != nil {
			return err
		}
	}
	if rangeTemplate.Offset != nil {
		if err = d.Set("offset", *rangeTemplate.Offset); err != nil {
			return err
		}
	}
	if rangeTemplate.UseOptions != nil {
		if err = d.Set("use_options", rangeTemplate.UseOptions); err != nil {
			return err
		}
	}
	if rangeTemplate.Options != nil {
		options := convertDhcpOptionsToInterface(rangeTemplate.Options)
		if err = d.Set("options", options); err != nil {
			return err
		}
	}
	if rangeTemplate.ServerAssociationType != "" {
		if err = d.Set("server_association_type", rangeTemplate.ServerAssociationType); err != nil {
			return err
		}
	}
	if rangeTemplate.FailoverAssociation != nil {
		if err = d.Set("failover_association", *rangeTemplate.FailoverAssociation); err != nil {
			return err
		}
	}
	if rangeTemplate.Member != nil {
		member := convertDhcpMemberToMap(rangeTemplate.Member)
		if err = d.Set("member", member); err != nil {
			return err
		}
	}
	if err = d.Set("cloud_api_compatible", rangeTemplate.CloudApiCompatible); err != nil {
		return err
	}
	if rangeTemplate.MsServer != nil {
		if err = d.Set("ms_server", rangeTemplate.MsServer.Ipv4Addr); err != nil {
			return err
		}
	}
	if err = d.Set("ref", rangeTemplate.Ref); err != nil {
		return err
	}
	d.SetId(rangeTemplate.Ref)
	return nil
}

func resourceRangeTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	var updateSuccessful bool
	defer func() {
		if !updateSuccessful {
			prevName, _ := d.GetChange("name")
			prevNumberOfAddresses, _ := d.GetChange("number_of_addresses")
			prevOffset, _ := d.GetChange("offset")
			prevComment, _ := d.GetChange("comment")
			prevUseOptions, _ := d.GetChange("use_options")
			prevOptions, _ := d.GetChange("options")
			prevServerAssociationType, _ := d.GetChange("server_association_type")
			prevFailoverAssociation, _ := d.GetChange("failover_association")
			prevMember, _ := d.GetChange("member")
			prevExtAttrs, _ := d.GetChange("ext_attrs")
			prevCloudApiCompatible, _ := d.GetChange("cloud_api_compatible")
			prevMsServer, _ := d.GetChange("ms_server")

			_ = d.Set("name", prevName.(string))
			_ = d.Set("number_of_addresses", prevNumberOfAddresses.(int))
			_ = d.Set("offset", prevOffset.(int))
			_ = d.Set("comment", prevComment.(string))
			_ = d.Set("use_options", prevUseOptions.(bool))
			_ = d.Set("options", prevOptions.([]interface{}))
			_ = d.Set("server_association_type", prevServerAssociationType.(string))
			_ = d.Set("failover_association", prevFailoverAssociation.(string))
			_ = d.Set("member", prevMember.(map[string]interface{}))
			_ = d.Set("ext_attrs", prevExtAttrs.(string))
			_ = d.Set("cloud_api_compatible", prevCloudApiCompatible.(bool))
			_ = d.Set("ms_server", prevMsServer.(string))
		}
	}()

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
	numberOfAddresses := d.Get("number_of_addresses").(int)
	offset := d.Get("offset").(int)
	comment := d.Get("comment").(string)
	useOptions := d.Get("use_options").(bool)
	serverAssociationType := d.Get("server_association_type").(string)
	failoverAssociation := d.Get("failover_association").(string)
	cloudApiCompatible := d.Get("cloud_api_compatible").(bool)
	msServer := d.Get("ms_server").(string)
	member := d.Get("member").(map[string]interface{})
	dhcpMemeber, err := ConvertMapToDhcpMember(member)
	if err != nil {
		return fmt.Errorf("failed to convert member to dhcpmember: %w", err)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)
	var rangeTemplate *ibclient.Rangetemplate

	rec, err := searchObjectByRefOrInternalId("RangeTemplate", d, m)
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
		return fmt.Errorf("failed to marshal Range Template record : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &rangeTemplate)
	if err != nil {
		return fmt.Errorf("failed getting Range Template record : %s", err.Error())
	}

	// If 'internal_id' is not set, then generate a new one and set it to the EA.
	internalId := d.Get("internal_id").(string)
	if internalId == "" {
		internalId = generateInternalId().String()
	}
	newInternalId := newInternalResourceIdFromString(internalId)
	newExtAttrs[eaNameForInternalId] = newInternalId.String()
	newExtAttrs, err = mergeEAs(rangeTemplate.Ea, newExtAttrs, oldExtAttrs, connector)
	if err != nil {
		return err
	}

	// Check if the options field has changes
	oldOptions, newOptions := d.GetChange("options")
	oldList, okOld := oldOptions.([]interface{})
	newList, okNew := newOptions.([]interface{})

	if !okOld || !okNew {
		return fmt.Errorf("options is not a slice of interfaces")
	}

	optimizedOptions := optimizeDhcpOptions(oldList, newList)
	options, err := validateDhcpOptions(optimizedOptions)
	if err != nil {
		return fmt.Errorf("failed to validate options: %w", err)
	}

	rangeTemplate, err = objMgr.UpdateRangeTemplate(d.Id(), name, uint32(numberOfAddresses), uint32(offset), comment, newExtAttrs, options, useOptions, serverAssociationType, failoverAssociation, dhcpMemeber, cloudApiCompatible, msServer)
	if err != nil {
		return fmt.Errorf("failed to update Range Template: %s.", err.Error())
	}
	updateSuccessful = true

	if err = d.Set("internal_id", newInternalId.String()); err != nil {
		return err
	}
	if err = d.Set("ref", rangeTemplate.Ref); err != nil {
		return err
	}
	d.SetId(rangeTemplate.Ref)
	return resourceRangeTemplateRead(d, m)
}

func resourceRangeTemplateDelete(d *schema.ResourceData, m interface{}) error {
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

	rec, err := searchObjectByRefOrInternalId("RangeTemplate", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		} else {
			d.SetId("")
			return nil
		}
	}

	var rangeTemplate *ibclient.Rangetemplate
	recJson, err := json.Marshal(rec)
	if err != nil {
		return fmt.Errorf("failed to marshal Range Template record : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &rangeTemplate)
	if err != nil {
		return err
	}
	_, err = objMgr.DeleteRangeTemplate(rangeTemplate.Ref)
	if err != nil {
		return fmt.Errorf("failed to delete Range Template : %s", err.Error())
	}
	return nil
}

func resourceRangeTemplateImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	extAttrJSON := d.Get("ext_attrs").(string)
	_, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return nil, err
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	rangeTemplate, err := objMgr.GetRangeTemplateByRef(d.Id())
	if err != nil {
		return nil, fmt.Errorf("failed getting Range Template record: %w", err)
	}

	if rangeTemplate.Ea != nil && len(rangeTemplate.Ea) > 0 {
		eaJSON, err := terraformSerializeEAs(rangeTemplate.Ea)
		if err != nil {
			return nil, err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return nil, err
		}
	}
	delete(rangeTemplate.Ea, eaNameForInternalId)

	if rangeTemplate.Name != nil {
		if err = d.Set("name", *rangeTemplate.Name); err != nil {
			return nil, err
		}
	}
	if rangeTemplate.Comment != nil {
		if err = d.Set("comment", *rangeTemplate.Comment); err != nil {
			return nil, err
		}
	}
	if rangeTemplate.NumberOfAddresses != nil {
		if err = d.Set("number_of_addresses", *rangeTemplate.NumberOfAddresses); err != nil {
			return nil, err
		}
	}
	if rangeTemplate.Offset != nil {
		if err = d.Set("offset", *rangeTemplate.Offset); err != nil {
			return nil, err
		}
	}
	if rangeTemplate.UseOptions != nil {
		if err = d.Set("use_options", rangeTemplate.UseOptions); err != nil {
			return nil, err
		}
	}
	if rangeTemplate.Options != nil {
		options := convertDhcpOptionsToInterface(rangeTemplate.Options)
		if err = d.Set("options", options); err != nil {
			return nil, err
		}
	}
	if rangeTemplate.ServerAssociationType != "" {
		if err = d.Set("server_association_type", rangeTemplate.ServerAssociationType); err != nil {
			return nil, err
		}
	}
	if rangeTemplate.FailoverAssociation != nil {
		if err = d.Set("failover_association", *rangeTemplate.FailoverAssociation); err != nil {
			return nil, err
		}
	}
	if rangeTemplate.Member != nil {
		member := convertDhcpMemberToMap(rangeTemplate.Member)
		if err = d.Set("member", member); err != nil {
			return nil, err
		}
	}
	if err = d.Set("cloud_api_compatible", rangeTemplate.CloudApiCompatible); err != nil {
		return nil, err
	}
	if rangeTemplate.MsServer != nil {
		if err = d.Set("ms_server", rangeTemplate.MsServer.Ipv4Addr); err != nil {
			return nil, err
		}
	}

	d.SetId(rangeTemplate.Ref)

	// Update the resource with the EA Terraform Internal ID
	err = resourceRangeTemplateUpdate(d, m)
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func ConvertMapToDhcpMember(m map[string]interface{}) (*ibclient.Dhcpmember, error) {
	if len(m) == 0 {
		return nil, nil
	}
	var dhcpMember ibclient.Dhcpmember
	name, ok := m["name"]
	if ok {
		dhcpMember.Name = name.(string)
	}
	ipv4Addr, ok := m["ipv4addr"]
	if ok {
		dhcpMember.Ipv4Addr = ipv4Addr.(string)
	}
	ipv6Addr, ok := m["ipv6addr"]
	if ok {
		dhcpMember.Ipv6Addr = ipv6Addr.(string)
	}
	return &dhcpMember, nil
}

func convertDhcpMemberToMap(member *ibclient.Dhcpmember) interface{} {
	if member == nil {
		return nil
	}
	return map[string]interface{}{
		"name":     member.Name,
		"ipv4addr": member.Ipv4Addr,
		"ipv6addr": member.Ipv6Addr,
	}
}
