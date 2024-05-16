package infoblox

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func resourceCNAMERecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceCNAMERecordCreate,
		Read:   resourceCNAMERecordGet,
		Update: resourceCNAMERecordUpdate,
		Delete: resourceCNAMERecordDelete,

		Importer: &schema.ResourceImporter{
			State: resourceCNAMERecordImport,
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
				Description: "Dns View under which the zone has been created.",
			},
			"canonical": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Canonical name in FQDN format.",
			},
			"alias": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The alias name in FQDN format.",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     ttlUndef,
				Description: "TTL attribute value for the record.",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "A description about CNAME record.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The Extensible attributes of CNAME record, as a map in JSON format",
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

func resourceCNAMERecordCreate(d *schema.ResourceData, m interface{}) error {
	// Check if internal_id is set manually
	if intId := d.Get("internal_id"); intId.(string) != "" {
		return fmt.Errorf("the value of 'internal_id' field must not be set manually")
	}
	dnsView := d.Get("dns_view").(string)
	canonical := d.Get("canonical").(string)
	alias := d.Get("alias").(string)

	comment := d.Get("comment").(string)
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
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

	var tenantID string
	if tempVal, ok := extAttrs[eaNameForTenantId]; ok {
		tenantID = tempVal.(string)
	}
	// Generate internal ID and add it to the extensible attributes
	internalId := generateInternalId()
	extAttrs[eaNameForInternalId] = internalId.String()

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	recordCNAME, err := objMgr.CreateCNAMERecord(dnsView, canonical, alias, useTtl, ttl, comment, extAttrs)
	if err != nil {
		return fmt.Errorf("creation of CNAME Record under %s DNS View failed: %s", dnsView, err.Error())
	}

	d.SetId(recordCNAME.Ref)

	if err = d.Set("internal_id", internalId.String()); err != nil {
		return err
	}
	if err = d.Set("ref", recordCNAME.Ref); err != nil {
		return err
	}
	return nil
}

func resourceCNAMERecordGet(d *schema.ResourceData, m interface{}) error {
	var ttl int
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}

	rec, err := searchObjectByRefOrInternalId("CNAME", d, m)
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
	var obj *ibclient.RecordCNAME
	recJson, _ := json.Marshal(rec)
	err = json.Unmarshal(recJson, &obj)
	if err != nil {
		return fmt.Errorf("getting CNAME Record with ID: %s failed: %s", d.Id(), err.Error())
	}

	if err = d.Set("alias", obj.Name); err != nil {
		return err
	}
	if err = d.Set("canonical", obj.Canonical); err != nil {
		return err
	}
	if err = d.Set("comment", obj.Comment); err != nil {
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

	if err = d.Set("dns_view", obj.View); err != nil {
		return err
	}
	if err = d.Set("ref", obj.Ref); err != nil {
		return err
	}

	d.SetId(obj.Ref)

	return nil
}

func resourceCNAMERecordUpdate(d *schema.ResourceData, m interface{}) error {
	var updateSuccessful bool
	defer func() {
		// Reverting the state back, in case of a failure,
		// otherwise Terraform will keep the values, which leaded to the failure,
		// in the state file.
		if !updateSuccessful {
			prevDNSView, _ := d.GetChange("dns_view")
			prevCanonical, _ := d.GetChange("canonical")
			prevAlias, _ := d.GetChange("alias")
			prevTTL, _ := d.GetChange("ttl")
			prevComment, _ := d.GetChange("comment")
			prevEa, _ := d.GetChange("ext_attrs")

			_ = d.Set("dns_view", prevDNSView.(string))
			_ = d.Set("canonical", prevCanonical.(string))
			_ = d.Set("alias", prevAlias.(string))
			_ = d.Set("ttl", prevTTL.(int))
			_ = d.Set("comment", prevComment.(string))
			_ = d.Set("ext_attrs", prevEa.(string))
		}
	}()

	if d.HasChange("dns_view") {
		return fmt.Errorf("changing the value of 'dns_view' field is not allowed")
	}
	if d.HasChange("internal_id") {
		return fmt.Errorf("changing the value of 'internal_id' field is not allowed")
	}

	dnsView := d.Get("dns_view").(string)
	canonical := d.Get("canonical").(string)
	alias := d.Get("alias").(string)
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

	var tenantID string
	if tempVal, ok := newExtAttrs[eaNameForTenantId]; ok {
		tenantID = tempVal.(string)
	}
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	crec, err := objMgr.GetCNAMERecordByRef(d.Id())
	if err != nil {
		return fmt.Errorf("failed to read CNAME record for update operation: %w", err)
	}

	// Generate internal ID and add it to the extensible attributes if not set
	internalId := d.Get("internal_id").(string)
	if internalId == "" {
		internalId = generateInternalId().String()
	}
	newInternalId := newInternalResourceIdFromString(internalId)
	newExtAttrs[eaNameForInternalId] = newInternalId.String()

	newExtAttrs, err = mergeEAs(crec.Ea, newExtAttrs, oldExtAttrs, connector)
	if err != nil {
		return err
	}

	recordCNAME, err := objMgr.UpdateCNAMERecord(d.Id(), canonical, alias, useTtl, ttl, comment, newExtAttrs)
	if err != nil {
		return fmt.Errorf("updation of CNAME Record under %s DNS View failed: %s", dnsView, err.Error())
	}
	updateSuccessful = true

	if err = d.Set("ref", recordCNAME.Ref); err != nil {
		return err
	}
	if err = d.Set("internal_id", newInternalId.String()); err != nil {
		return err
	}
	d.SetId(recordCNAME.Ref)

	return nil
}

func resourceCNAMERecordDelete(d *schema.ResourceData, m interface{}) error {
	dnsView := d.Get("dns_view").(string)
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}

	var tenantID string
	if tempVal, ok := extAttrs[eaNameForTenantId]; ok {
		tenantID = tempVal.(string)
	}
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	rec, err := searchObjectByRefOrInternalId("CNAME", d, m)
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
	var crec *ibclient.RecordCNAME
	recJson, _ := json.Marshal(rec)
	err = json.Unmarshal(recJson, &crec)

	_, err = objMgr.DeleteCNAMERecord(crec.Ref)
	if err != nil {
		return fmt.Errorf("deletion of CNAME Record from dns view %s failed: %s", dnsView, err.Error())
	}
	d.SetId("")

	return nil
}

func resourceCNAMERecordImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	var ttl int
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

	obj, err := objMgr.GetCNAMERecordByRef(d.Id())
	if err != nil {
		return nil, fmt.Errorf("getting CNAME Record with ID: %s failed: %s", d.Id(), err.Error())
	}

	if err = d.Set("alias", obj.Name); err != nil {
		return nil, err
	}
	if err = d.Set("canonical", obj.Canonical); err != nil {
		return nil, err
	}
	if err = d.Set("comment", obj.Comment); err != nil {
		return nil, err
	}

	if obj.Ttl != nil {
		ttl = int(*obj.Ttl)
	}

	if !*obj.UseTtl {
		ttl = ttlUndef
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

	if err = d.Set("dns_view", obj.View); err != nil {
		return nil, err
	}

	d.SetId(obj.Ref)
	// Set ref in the state
	if err = d.Set("ref", obj.Ref); err != nil {
		return nil, err
	}

	// Update the resource with the EA Terraform Internal ID
	err = resourceCNAMERecordUpdate(d, m)
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
