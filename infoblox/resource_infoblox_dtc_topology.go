package infoblox

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

// convert tf data structure to API-client object dtc:topology:rule
// TODO: user-visible error handling
func convertTfListToDtcTopologyRules(tf []any) []*ibclient.DtcTopologyRule {
	result := make([]*ibclient.DtcTopologyRule, len(tf))
	for ruleIdx, rule := range tf {
		if ruleMap, ok := rule.(map[string]any); ok {
			dtcRule := &ibclient.DtcTopologyRule{}
			if ruleDestType, ok := ruleMap["dest_type"]; ok {
				dtcRule.DestType = ruleDestType.(string)
			}
			if ruleDestination, ok := ruleMap["destination"]; ok {
				rd := ruleDestination.(string)
				dtcRule.DestinationLink = &rd
			}
			if ruleReturnType, ok := ruleMap["return_type"]; ok {
				dtcRule.ReturnType = ruleReturnType.(string)
			}
			if ruleSourcesTf, ok := ruleMap["sources"]; ok {
				if ruleSourcesListTf, ok := ruleSourcesTf.([]any); ok {
					dtcRule.Sources = make([]*ibclient.DtcTopologyRuleSource, 0, len(ruleSourcesListTf))
					for _, ruleSourceTf := range ruleSourcesListTf {
						if ruleSource, ok := ruleSourceTf.(map[string]any); ok {
							var ruleSource = ibclient.DtcTopologyRuleSource{
								SourceOp:    ruleSource["source_op"].(string),
								SourceType:  ruleSource["source_type"].(string),
								SourceValue: ruleSource["source_value"].(string),
							}
							dtcRule.Sources = append(dtcRule.Sources, &ruleSource)
						}
					}
				}
			}
			if ruleValid, ok := ruleMap["valid"]; ok {
				dtcRule.Valid = ruleValid.(bool)
			}
			result[ruleIdx] = dtcRule
		}
	}
	return result
}

// convert API-client object dtc:topology:rule to tf data structure
func convertDtcTopologyRulesToTfList(rules []*ibclient.DtcTopologyRule) []map[string]any {
	result := make([]map[string]any, 0, len(rules))
	for _, dtcRule := range rules {
		ruleMap := make(map[string]any)
		ruleMap["dest_type"] = dtcRule.DestType
		if dest := dtcRule.DestinationLink; dest != nil {
			ruleMap["destination"] = *dest
		}
		ruleMap["return_type"] = dtcRule.ReturnType
		ruleSources := make([]map[string]string, 0, len(dtcRule.Sources))
		for _, dtcRuleSource := range dtcRule.Sources {
			ruleSources = append(ruleSources, map[string]string{
				"source_op":    dtcRuleSource.SourceOp,
				"source_type":  dtcRuleSource.SourceType,
				"source_value": dtcRuleSource.SourceValue,
			})
		}
		ruleMap["sources"] = ruleSources
		ruleMap["valid"] = dtcRule.Valid
		result = append(result, ruleMap)
	}
	return result
}

func resourceDtcTopology() *schema.Resource {
	return &schema.Resource{
		Description: `A Topology Ruleset. Can be used by dtc_lbdn and dtc_pool when setting lb_method = "TOPOLOGY".`,
		Create:      resourceDtcTopologyCreate,
		Read:        resourceDtcTopologyGet,
		Update:      resourceDtcTopologyUpdate,
		Delete:      resourceDtcTopologyDelete,
		Importer: &schema.ResourceImporter{
			State: resourceDtcTopologyImport,
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
				Default:     "",
				Description: "Description of the Dtc topology.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Extensible attributes of the Dtc Topology to be added/updated, as a map in JSON format",
			},
			"rules": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of DTC topology rules.\n\n",
				Elem: &schema.Resource{
					//check the required part once
					Schema: map[string]*schema.Schema{
						"dest_type": {
							Description: "The type of the destination for this DTC Topology rule. Valid values are: POOL SERVER .",
							Required:    true,
							Type:        schema.TypeString,
						},
						"destination": {
							Description: "The reference to the destination DTC pool or DTC server.",
							Required:    true,
							Type:        schema.TypeString,
						},
						"return_type": {
							Description: `Type of the DNS response for rule. Valid values are:
								NOERR
								NXDOMAIN
								REGULAR
							.`,
							Required: true,
							Type:     schema.TypeString,
						},
						"sources": {
							Description: "The conditions for matching sources. Should be empty to set rule as default destination.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"source_op": {
										Description: "The operation used to match the value. IS or IS_NOT",
										Type:        schema.TypeString,
										Required:    true,
									},
									"source_type": {
										Description: `The source type. Valid values are:
											CITY
											CONTINENT
											COUNTRY
											EA0
											EA1
											EA2
											EA3
											SUBDIVISION
											SUBNET
										.`,
										Type:     schema.TypeString,
										Required: true,
									},
									"source_value": {
										Description: "The source value.",
										Type:        schema.TypeString,
										Required:    true,
									},
								},
							},
						},
						"valid": {
							Description: "True if the label in the rule exists in the current Topology DB. Always true for SUBNET rules. " +
								"Rules with non-existent labels may be configured but will never match.",
							Type:     schema.TypeBool,
							Optional: false,
							Computed: true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The DTC Topology display name.",
			},
			"ref": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "NIOS object's reference, not to be set by a user.",
			},
			"internal_id": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "Internal ID of an object at NIOS side," +
					" used by Infoblox Terraform plugin to search for a NIOS's object" +
					" which corresponds to the Terraform resource.",
			},
		},
	}
}

func resourceDtcTopologyCreate(d *schema.ResourceData, m interface{}) error {
	// Check if internal_id is set manually
	if intId := d.Get("internal_id"); intId.(string) != "" {
		return fmt.Errorf("the value of 'internal_id' field must not be set manually")
	}

	comment := d.Get("comment").(string)
	name := d.Get("name").(string)
	rules := d.Get("rules").([]any)
	dtcTopologyRules := convertTfListToDtcTopologyRules(rules)
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

	newDtcTopology, err := objMgr.CreateDtcTopology(comment, name, dtcTopologyRules, extAttrs)
	if err != nil {
		return err
	}
	d.SetId(newDtcTopology.Ref)
	if err = d.Set("internal_id", internalId.String()); err != nil {
		return err
	}
	if err = d.Set("ref", newDtcTopology.Ref); err != nil {
		return err
	}
	return resourceDtcTopologyGet(d, m)
}

func resourceDtcTopologyGet(d *schema.ResourceData, m interface{}) error {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}

	rec, err := searchObjectByRefOrInternalId("DtcTopology", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); ok {
			d.SetId("")
			return nil
		} else {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		}
	}
	recJson, err := json.Marshal(rec)
	if err != nil {
		return fmt.Errorf("failed to marshal DTC Topology : %s", err.Error())
	}
	var dtcTopology *ibclient.DtcTopology
	err = json.Unmarshal(recJson, &dtcTopology)
	if err != nil {
		return fmt.Errorf("failed getting DTC Topology : %s", err.Error())
	}
	delete(dtcTopology.Ea, eaNameForInternalId)
	omittedEAs := omitEAs(dtcTopology.Ea, extAttrs)

	if omittedEAs != nil && len(omittedEAs) > 0 {
		eaJSON, err := terraformSerializeEAs(omittedEAs)
		if err != nil {
			return err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return err
		}
	}

	if err = d.Set("name", dtcTopology.Name); err != nil {
		return err
	}
	if err = d.Set("comment", dtcTopology.Comment); err != nil {
		return err
	}
	rulesInterface := convertDtcTopologyRulesToTfList(dtcTopology.Rules)
	if err = d.Set("rules", rulesInterface); err != nil {
		return err
	}
	if err = d.Set("ref", dtcTopology.Ref); err != nil {
		return err
	}
	d.SetId(dtcTopology.Ref)
	return nil
}

func resourceDtcTopologyUpdate(d *schema.ResourceData, m interface{}) error {
	var updateSuccessful bool
	defer func() {
		if !updateSuccessful {
			prevName, _ := d.GetChange("name")
			prevComment, _ := d.GetChange("comment")
			prevEa, _ := d.GetChange("ext_attrs")
			prevRules, _ := d.GetChange("rules")

			_ = d.Set("comment", prevComment.(string))
			_ = d.Set("name", prevName.(string))
			_ = d.Set("ext_attrs", prevEa.(string))
			_ = d.Set("rules", prevRules)
		}
	}()
	comment := d.Get("comment").(string)
	name := d.Get("name").(string)
	rules := d.Get("rules").([]interface{})
	dtcTopologyRules := convertTfListToDtcTopologyRules(rules)
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

	var dtcTopology *ibclient.DtcTopology

	rec, err := searchObjectByRefOrInternalId("DtcTopology", d, m)
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
		return fmt.Errorf("failed to marshal Dtc Topology : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &dtcTopology)
	if err != nil {
		return fmt.Errorf("failed getting Dtc Topology : %s", err.Error())
	}

	// If 'internal_id' is not set, then generate a new one and set it to the EA.
	internalId := d.Get("internal_id").(string)
	if internalId == "" {
		internalId = generateInternalId().String()
	}
	newInternalId := newInternalResourceIdFromString(internalId)
	newExtAttrs[eaNameForInternalId] = newInternalId.String()

	newExtAttrs, err = mergeEAs(dtcTopology.Ea, newExtAttrs, oldExtAttrs, connector)
	if err != nil {
		return err
	}
	dtcTopology, err = objMgr.UpdateDtcTopology(d.Id(), comment, name, newExtAttrs, dtcTopologyRules)
	if err != nil {
		return fmt.Errorf("error updating dtc-topology: %w", err)
	}
	updateSuccessful = true
	d.SetId(dtcTopology.Ref)
	if err = d.Set("ref", dtcTopology.Ref); err != nil {
		return err
	}
	if err = d.Set("internal_id", newInternalId.String()); err != nil {
		return err
	}
	return resourceDtcTopologyGet(d, m)
}

func resourceDtcTopologyDelete(d *schema.ResourceData, m interface{}) error {
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

	rec, err := searchObjectByRefOrInternalId("DtcTopology", d, m)
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
	var DtcTopology *ibclient.DtcTopology
	recJson, _ := json.Marshal(rec)
	err = json.Unmarshal(recJson, &DtcTopology)

	_, err = objMgr.DeleteDtcTopology(DtcTopology.Ref)
	if err != nil {
		return fmt.Errorf("deletion of Dtc Topology failed: %w", err)
	}
	d.SetId("")

	return nil
}

func resourceDtcTopologyImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return nil, err
	}

	var tenantID string
	if tempVal, ok := extAttrs[eaNameForTenantId]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)
	obj, err := objMgr.GetDtcTopologyByRef(d.Id())

	if err != nil {
		return nil, fmt.Errorf("getting DtcTopology with ID: %s failed: %w", d.Id(), err)
	}

	// Set ref
	if err = d.Set("ref", obj.Ref); err != nil {
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

	if err = d.Set("name", obj.Name); err != nil {
		return nil, err
	}
	if err = d.Set("comment", obj.Comment); err != nil {
		return nil, err
	}
	rulesInterface := convertDtcTopologyRulesToTfList(obj.Rules)
	if err = d.Set("rules", rulesInterface); err != nil {
		return nil, err
	}

	d.SetId(obj.Ref)
	err = resourceDtcTopologyUpdate(d, m)
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
