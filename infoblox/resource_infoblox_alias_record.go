package infoblox

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func resourceAliasRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliasRecordCreate,
		Read:   resourceAliasRecordRead,
		Update: resourceAliasRecordUpdate,
		Delete: resourceAliasRecordDelete,
		Importer: &schema.ResourceImporter{
			State: resourceAliasRecordImport,
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
				Description: "Name of the alias record.",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comment for the alias record.",
			},
			"disable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "A boolean flag which indicates if the alias record is disabled.",
			},
			"target_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Target name in FQDN format.",
			},
			"target_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of the target object.",
			},
			"dns_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     defaultDNSView,
				Description: "Name of the DNS view in which the alias record is created.",
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  ttlUndef,
				Description: "The Time To Live (TTL) value for record. A 32-bit unsigned integer that represents the duration, " +
					"in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Extensible attributes of the  Alias Record to be added/updated, as a map in JSON format",
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

func resourceAliasRecordCreate(d *schema.ResourceData, m interface{}) error {
	// Check if internal_id is set manually
	if intId := d.Get("internal_id"); intId.(string) != "" {
		return fmt.Errorf("the value of 'internal_id' field must not be set manually")
	}

	name := d.Get("name").(string)
	comment := d.Get("comment").(string)
	disable := d.Get("disable").(bool)
	targetName := d.Get("target_name").(string)
	targetType := d.Get("target_type").(string)
	dnsView := d.Get("dns_view").(string)

	var ttl uint32
	useTtl := false
	tempVal := d.Get("ttl")
	tempTTL := tempVal.(int)
	if tempTTL >= 0 {
		useTtl = true
		ttl = uint32(tempTTL)
	} else if tempTTL != ttlUndef {
		return fmt.Errorf("TTL value must be 0 or higher")
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

	// create alias record
	aliasRecord, err := objMgr.CreateAliasRecord(name, dnsView, targetName, targetType, comment, disable, extAttrs, ttl, useTtl)
	if err != nil {
		return fmt.Errorf("failed to create alias record: %w", err)
	}
	d.SetId(aliasRecord.Ref)
	if err = d.Set("internal_id", internalId.String()); err != nil {
		return err
	}
	if err = d.Set("ref", aliasRecord.Ref); err != nil {
		return err
	}
	return resourceAliasRecordRead(d, m)
}

func resourceAliasRecordRead(d *schema.ResourceData, m interface{}) error {
	var ttl int
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}

	rec, err := searchObjectByRefOrInternalId("AliasRecord", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); ok {
			d.SetId("")
			return nil
		} else {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		}
	}
	var recordAlias *ibclient.RecordAlias

	recJson, err := json.Marshal(rec)
	if err != nil {
		return fmt.Errorf("failed to marshal Alias record : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &recordAlias)
	if err != nil {
		return fmt.Errorf("failed getting Alias record : %s", err.Error())
	}

	delete(recordAlias.Ea, eaNameForInternalId)
	omittedEAs := omitEAs(recordAlias.Ea, extAttrs)

	if omittedEAs != nil && len(omittedEAs) > 0 {
		eaJSON, err := terraformSerializeEAs(omittedEAs)
		if err != nil {
			return err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return err
		}
	}

	if recordAlias.Name != nil {
		if err = d.Set("name", *recordAlias.Name); err != nil {
			return err
		}
	}
	if recordAlias.Comment != nil {
		if err = d.Set("comment", *recordAlias.Comment); err != nil {
			return err
		}
	}
	if recordAlias.Disable != nil {
		if err = d.Set("disable", *recordAlias.Disable); err != nil {
			return err
		}
	}
	if recordAlias.TargetName != nil {
		if err = d.Set("target_name", *recordAlias.TargetName); err != nil {
			return err
		}
	}
	if err = d.Set("target_type", recordAlias.TargetType); err != nil {
		return err
	}
	if recordAlias.View != nil {
		if err = d.Set("dns_view", *recordAlias.View); err != nil {
			return err
		}
	}

	if recordAlias.Ttl != nil {
		ttl = int(*recordAlias.Ttl)
	}
	if !*recordAlias.UseTtl {
		ttl = ttlUndef
	}

	if err = d.Set("ttl", ttl); err != nil {
		return err
	}

	if err = d.Set("ref", recordAlias.Ref); err != nil {
		return err
	}

	d.SetId(recordAlias.Ref)

	return nil
}

func resourceAliasRecordUpdate(d *schema.ResourceData, m interface{}) error {
	var updateSuccessful bool
	defer func() {
		if !updateSuccessful {
			prevName, _ := d.GetChange("name")
			prevComment, _ := d.GetChange("comment")
			prevDisable, _ := d.GetChange("disable")
			prevTargetName, _ := d.GetChange("target_name")
			prevTargetType, _ := d.GetChange("target_type")
			prevDnsView, _ := d.GetChange("dns_view")
			prevTTL, _ := d.GetChange("ttl")
			prevExtAttrs, _ := d.GetChange("ext_attrs")

			_ = d.Set("name", prevName)
			_ = d.Set("comment", prevComment)
			_ = d.Set("disable", prevDisable)
			_ = d.Set("target_name", prevTargetName)
			_ = d.Set("target_type", prevTargetType)
			_ = d.Set("dns_view", prevDnsView)
			_ = d.Set("ttl", prevTTL)
			_ = d.Set("ext_attrs", prevExtAttrs)
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
	comment := d.Get("comment").(string)
	disable := d.Get("disable").(bool)
	targetName := d.Get("target_name").(string)
	targetType := d.Get("target_type").(string)
	dnsView := d.Get("dns_view").(string)

	var ttl uint32
	useTtl := false
	tempVal := d.Get("ttl")
	tempTTL := tempVal.(int)
	if tempTTL >= 0 {
		useTtl = true
		ttl = uint32(tempTTL)
	} else if tempTTL != ttlUndef {
		return fmt.Errorf("TTL value must be 0 or higher")
	}

	if d.HasChange("dns_view") {
		return fmt.Errorf("changing the value of 'dns_view' field is not allowed")
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	var recordAlias *ibclient.RecordAlias

	rec, err := searchObjectByRefOrInternalId("AliasRecord", d, m)
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
		return fmt.Errorf("failed to marshal alias record : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &recordAlias)
	if err != nil {
		return fmt.Errorf("failed getting alias record : %s", err.Error())
	}

	// If 'internal_id' is not set, then generate a new one and set it to the EA.
	internalId := d.Get("internal_id").(string)
	if internalId == "" {
		internalId = generateInternalId().String()
	}
	newInternalId := newInternalResourceIdFromString(internalId)
	newExtAttrs[eaNameForInternalId] = newInternalId.String()

	newExtAttrs, err = mergeEAs(recordAlias.Ea, newExtAttrs, oldExtAttrs, connector)
	if err != nil {
		return err
	}

	updatedRecord, err := objMgr.UpdateAliasRecord(d.Id(), name, dnsView, targetName, targetType, comment, disable, newExtAttrs, ttl, useTtl)
	if err != nil {
		return fmt.Errorf("Failed to update alias Record with %s, ", err.Error())
	}

	updateSuccessful = true

	if err = d.Set("internal_id", newInternalId.String()); err != nil {
		return err
	}
	if err = d.Set("ref", updatedRecord.Ref); err != nil {
		return err
	}
	d.SetId(updatedRecord.Ref)

	return nil
}

func resourceAliasRecordDelete(d *schema.ResourceData, m interface{}) error {
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

	rec, err := searchObjectByRefOrInternalId("AliasRecord", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		} else {
			d.SetId("")
			return nil
		}
	}

	var aliasRecord *ibclient.RecordAlias
	recJson, err := json.Marshal(rec)
	if err != nil {
		return fmt.Errorf("failed to marshal alias record : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &aliasRecord)
	if err != nil {
		return err
	}
	_, err = objMgr.DeleteAliasRecord(aliasRecord.Ref)
	if err != nil {
		return fmt.Errorf("failed to delete alias : %s", err.Error())
	}
	return nil
}

func resourceAliasRecordImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	var ttl int
	extAttrJSON := d.Get("ext_attrs").(string)
	_, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return nil, err
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	aliasRecord, err := objMgr.GetAliasRecordByRef(d.Id())
	if err != nil {
		return nil, fmt.Errorf("failed getting alias record: %w", err)
	}

	if aliasRecord.Ea != nil && len(aliasRecord.Ea) > 0 {
		eaJSON, err := terraformSerializeEAs(aliasRecord.Ea)
		if err != nil {
			return nil, err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return nil, err
		}
	}
	delete(aliasRecord.Ea, eaNameForInternalId)

	if aliasRecord.Name != nil {
		if err = d.Set("name", *aliasRecord.Name); err != nil {
			return nil, err
		}
	}
	if aliasRecord.Comment != nil {
		if err = d.Set("comment", *aliasRecord.Comment); err != nil {
			return nil, err
		}
	}
	if aliasRecord.Disable != nil {
		if err = d.Set("disable", *aliasRecord.Disable); err != nil {
			return nil, err
		}
	}
	if aliasRecord.TargetName != nil {
		if err = d.Set("target_name", *aliasRecord.TargetName); err != nil {
			return nil, err
		}
	}
	if err = d.Set("target_type", aliasRecord.TargetType); err != nil {
		return nil, err
	}
	if aliasRecord.View != nil {
		if err = d.Set("dns_view", *aliasRecord.View); err != nil {
			return nil, err
		}
	}

	if aliasRecord.Ttl != nil {
		ttl = int(*aliasRecord.Ttl)
	}
	if !*aliasRecord.UseTtl {
		ttl = ttlUndef
	}

	if err = d.Set("ttl", ttl); err != nil {
		return nil, err
	}

	if err = d.Set("ref", aliasRecord.Ref); err != nil {
		return nil, err
	}
	d.SetId(aliasRecord.Ref)
	err = resourceAliasRecordUpdate(d, m)

	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
