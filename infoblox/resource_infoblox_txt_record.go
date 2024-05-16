package infoblox

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func resourceTXTRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceTXTRecordCreate,
		Read:   resourceTXTRecordGet,
		Update: resourceTXTRecordUpdate,
		Delete: resourceTXTRecordDelete,

		Importer: &schema.ResourceImporter{
			State: resourceTXTRecordImport,
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
			"dns_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     defaultDNSView,
				Description: "DNS view in which the record's zone exists.",
			},
			"fqdn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "FQDN for the TXT-Record.",
			},
			"text": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Data to be associated with TXT_Record.",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     ttlUndef,
				Description: "TTL value of the TXT-Record",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Description of the TXT-record.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Extensible attributes of the TXT-record to be added/updated, as a map in JSON format",
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

func resourceTXTRecordCreate(d *schema.ResourceData, m interface{}) error {

	if intId := d.Get("internal_id"); intId.(string) != "" {
		return fmt.Errorf("the value of 'internal_id' field must not be set manually")
	}

	dnsView := d.Get("dns_view").(string)
	fqdn := d.Get("fqdn").(string)
	text := d.Get("text").(string)
	if text == "" {
		return fmt.Errorf("empty 'text' value is not allowed")
	}

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

	comment := d.Get("comment").(string)

	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}

	// Generate internal ID and add it to the extensible attributes
	internalId := generateInternalId()
	extAttrs[eaNameForInternalId] = internalId.String()

	var tenantID string
	tempVal, found := extAttrs[eaNameForTenantId]
	if found {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	newRecord, err := objMgr.CreateTXTRecord(dnsView, fqdn, text, ttl, useTtl, comment, extAttrs)

	if err != nil {
		return fmt.Errorf("error creating TXT-Record: %s", err)
	}

	d.SetId(newRecord.Ref)

	if err = d.Set("ref", newRecord.Ref); err != nil {
		return err
	}
	if err = d.Set("internal_id", internalId.String()); err != nil {
		return err
	}

	return nil
}

func resourceTXTRecordGet(d *schema.ResourceData, m interface{}) error {
	var ttl int
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}

	rec, err := searchObjectByRefOrInternalId("TXT", d, m)
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
	var obj *ibclient.RecordTXT
	recJson, _ := json.Marshal(rec)
	err = json.Unmarshal(recJson, &obj)

	if err != nil && obj.Ref != "" {
		return fmt.Errorf("failed getting TXT-Record: %s", err)
	}

	if err = d.Set("text", obj.Text); err != nil {
		return err
	}

	if obj.Ttl != nil {
		ttl = int(*obj.Ttl)
	}

	if !*obj.UseTtl {
		ttl = ttlUndef
	}
	if err = d.Set("ttl", ttl); err != nil {
		return err
	}

	delete(obj.Ea, eaNameForInternalId)
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

	if err = d.Set("dns_view", obj.View); err != nil {
		return err
	}

	if err = d.Set("fqdn", obj.Name); err != nil {
		return err
	}

	if err = d.Set("ref", obj.Ref); err != nil {
		return err
	}

	d.SetId(obj.Ref)

	return nil
}

func resourceTXTRecordUpdate(d *schema.ResourceData, m interface{}) error {
	var updateSuccessful bool
	defer func() {
		// Reverting the state back, in case of a failure,
		// otherwise Terraform will keep the values, which leaded to the failure
		// in the state file.
		if !updateSuccessful {
			prevDNSView, _ := d.GetChange("dns_view")
			prevFQDN, _ := d.GetChange("fqdn")
			prevTEXT, _ := d.GetChange("text")
			prevTTL, _ := d.GetChange("ttl")
			prevComment, _ := d.GetChange("comment")
			prevEa, _ := d.GetChange("ext_attrs")

			_ = d.Set("dns_view", prevDNSView.(string))
			_ = d.Set("fqdn", prevFQDN.(string))
			_ = d.Set("text", prevTEXT.(string))
			_ = d.Set("ttl", prevTTL.(int))
			_ = d.Set("comment", prevComment.(string))
			_ = d.Set("ext_attrs", prevEa.(string))
		}
	}()
	if d.HasChange("internal_id") {
		return fmt.Errorf("changing the value of 'internal_id' field is not allowed")
	}

	if d.HasChange("dns_view") {
		return fmt.Errorf("changing the value of 'dns_view' field is not allowed")
	}

	text := d.Get("text").(string)
	if text == "" {
		return fmt.Errorf("empty 'text' value is not allowed")
	}

	fqdn := d.Get("fqdn").(string)

	var ttl uint32
	useTtl := false
	tempVal := d.Get("ttl")
	tempTTL := tempVal.(int)
	if tempTTL >= 0 {
		useTtl = true
		ttl = uint32(tempTTL)
	} else if tempTTL != ttlUndef {
		return fmt.Errorf("TTL must be 0 or higher")
	}

	comment := d.Get("comment").(string)

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
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	txtrec, err := objMgr.GetTXTRecordByRef(d.Id())
	if err != nil {
		return fmt.Errorf("failed to read TXT Record for update operation: %w", err)
	}

	internalId := d.Get("internal_id").(string)

	if internalId == "" {
		internalId = generateInternalId().String()
	}

	newInternalId := newInternalResourceIdFromString(internalId)
	newExtAttrs[eaNameForInternalId] = newInternalId.String()

	newExtAttrs, err = mergeEAs(txtrec.Ea, newExtAttrs, oldExtAttrs, connector)
	if err != nil {
		return err
	}

	rec, err := objMgr.UpdateTXTRecord(
		d.Id(), fqdn, text, ttl, useTtl, comment, newExtAttrs)
	if err != nil {
		return fmt.Errorf("error updating TXT-Record: %s", err)
	}
	updateSuccessful = true
	d.SetId(rec.Ref)

	if err = d.Set("ref", rec.Ref); err != nil {
		return err
	}
	if err = d.Set("internal_id", newInternalId.String()); err != nil {
		return err
	}
	return nil
}

func resourceTXTRecordDelete(d *schema.ResourceData, m interface{}) error {
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
	txtrec, err := searchObjectByRefOrInternalId("TXT", d, m)
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
	var obj ibclient.RecordTXT
	recJson, _ := json.Marshal(txtrec)
	err = json.Unmarshal(recJson, &obj)

	if err != nil {
		return fmt.Errorf("failed getting TXT-Record: %s", err)
	}

	_, err = objMgr.DeleteTXTRecord(obj.Ref)
	if err != nil {
		return fmt.Errorf("deletion of TXT-Record failed: %s", err)
	}
	d.SetId("")

	return nil
}

func resourceTXTRecordImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	var ttl int
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

	obj, err := objMgr.GetTXTRecordByRef(d.Id())
	if err != nil {
		return nil, fmt.Errorf("failed getting TXT-Record: %s", err)
	}

	if err = d.Set("text", obj.Text); err != nil {
		return nil, err
	}

	if obj.Ttl != nil {
		ttl = int(*obj.Ttl)
	}

	if !*obj.UseTtl {
		ttl = ttlUndef
	}

	// Set ref
	if err = d.Set("ref", obj.Ref); err != nil {
		return nil, err
	}

	if err = d.Set("ttl", ttl); err != nil {
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

	if err = d.Set("comment", obj.Comment); err != nil {
		return nil, err
	}

	if err = d.Set("dns_view", obj.View); err != nil {
		return nil, err
	}

	if err = d.Set("fqdn", obj.Name); err != nil {
		return nil, err
	}

	d.SetId(obj.Ref)

	err = resourceTXTRecordUpdate(d, m)
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
