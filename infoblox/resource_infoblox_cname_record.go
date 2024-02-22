package infoblox

import (
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
			"ext_attrs": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The Extensible attributes of CNAME record, as a map in JSON format",
			},
		},
	}
}

func resourceCNAMERecordCreate(d *schema.ResourceData, m interface{}) error {

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

	// Generate internal ID and add it to the extensible attributes
	internalId := generateInternalId()
	extAttrs[eaNameForInternalId] = internalId.String()

	var tenantID string
	if tempVal, ok := extAttrs[eaNameForTenantId]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	recordCNAME, err := objMgr.CreateCNAMERecord(dnsView, canonical, alias, useTtl, ttl, comment, extAttrs)
	if err != nil {
		return fmt.Errorf("creation of CNAME Record under %s DNS View failed: %s", dnsView, err.Error())
	}

	//d.SetId(recordCNAME.Ref)
	d.SetId(internalId.String())
	if err = d.Set("ref", recordCNAME.Ref); err != nil {
		return err
	}
	if err = d.Set("internal_id", internalId.String()); err != nil {
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

	var tenantID string
	var (
		recordCname *ibclient.RecordCNAME
		cnameList   []ibclient.RecordCNAME
	)
	if tempVal, ok := extAttrs[eaNameForTenantId]; ok {
		tenantID = tempVal.(string)
	}
	connector := m.(ibclient.IBConnector)
	ibclient.NewObjectManager(connector, "Terraform", tenantID)

	//obj, err := objMgr.GetCNAMERecordByRef(d.Id())
	obj, err := getOrFindRecord("CNAME", d, m)

	if err != nil {
		d.SetId("")
		return nil
		//return fmt.Errorf("getting CNAME Record with ID: %s failed: %s", d.Id(), err.Error())
	}

	recJson, _ := json.Marshal(obj)
	json.Unmarshal(recJson, &cnameList)
	byteObj, _ := json.Marshal(cnameList[0])
	err = json.Unmarshal(byteObj, &recordCname)
	if err != nil {
		return fmt.Errorf("unable to unmarshal %v", err.Error())
	}

	if err = d.Set("alias", recordCname.Name); err != nil {
		return err
	}
	if err = d.Set("canonical", recordCname.Canonical); err != nil {
		return err
	}
	if err = d.Set("comment", recordCname.Comment); err != nil {
		return err
	}

	if recordCname.Ttl != nil {
		ttl = int(*recordCname.Ttl)
	}

	if !*recordCname.UseTtl {
		ttl = ttlUndef
	}
	if err = d.Set("ttl", ttl); err != nil {
		return err
	}

	omittedEAs := omitEAs(recordCname.Ea, extAttrs)

	if omittedEAs != nil && len(omittedEAs) > 0 {
		eaJSON, err := terraformSerializeEAs(omittedEAs)
		if err != nil {
			return err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return err
		}
	}

	if err = d.Set("dns_view", recordCname.View); err != nil {
		return err
	}

	d.SetId(recordCname.Ref)

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

	dnsView := d.Get("dns_view").(string)
	if d.HasChange("dns_view") {
		return fmt.Errorf("changing the value of 'dns_view' field is not allowed")
	}
	canonical := d.Get("canonical").(string)
	alias := d.Get("alias").(string)

	comment := d.Get("comment").(string)

	var (
		crec      *ibclient.RecordCNAME
		cnameList []ibclient.RecordCNAME
	)

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

	//crec, err := objMgr.GetCNAMERecordByRef(d.Id())
	obj, err := getOrFindRecord("A", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); ok {
			d.SetId("")
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find apropriate object on NIOS side for resource with ID '%s': %s;"+
					" removing the resource from Terraform state",
				d.Id(), err))
		}

		return err
	}
	recJson, _ := json.Marshal(obj)
	json.Unmarshal(recJson, &cnameList)
	byteObj, _ := json.Marshal(cnameList[0])
	err = json.Unmarshal(byteObj, &crec)
	if err != nil {
		return fmt.Errorf("unable to unmarshal json %v", err.Error())
	}
	if d.HasChange("internal_id") {
		return fmt.Errorf("changing the value of 'internal_id' field is not allowed")
	}
	internalId := newInternalResourceIdFromString(d.Get("internal_id").(string))
	newExtAttrs[eaNameForInternalId] = internalId.String()
	//if err != nil {
	//	return fmt.Errorf("failed to read CNAME record for update operation: %w", err)
	//}

	newExtAttrs, err = mergeEAs(crec.Ea, newExtAttrs, oldExtAttrs, connector)
	if err != nil {
		return err
	}

	recordCNAME, err := objMgr.UpdateCNAMERecord(d.Id(), canonical, alias, useTtl, ttl, comment, newExtAttrs)
	if err != nil {
		return fmt.Errorf("updation of CNAME Record under %s DNS View failed: %s", dnsView, err.Error())
	}
	updateSuccessful = true

	d.SetId(recordCNAME.Ref)
	if err = d.Set("ref", crec.Ref); err != nil {
		return err
	}

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

	_, err = objMgr.DeleteCNAMERecord(d.Id())
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

	// Internal ID
	// Generate internal ID and add it to the extensible attributes
	internalId := generateInternalId()
	extAttrs[eaNameForInternalId] = internalId.String()
	if err = d.Set("internal_id", internalId.String()); err != nil {
		return nil, err
	}
	// Set ref
	if err = d.Set("ref", obj.Ref); err != nil {
		return nil, err
	}
	delete(obj.Ea, eaNameForInternalId)
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

	// Resource CNAME Record update Terraform Internal ID and Ref on NIOS side
	// After the record is imported, call the update function
	_, err = connector.UpdateObject(&ibclient.RecordCNAME{
		Ea: ibclient.EA{
			eaNameForInternalId: internalId.String(),
		},
	}, obj.Ref)
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
