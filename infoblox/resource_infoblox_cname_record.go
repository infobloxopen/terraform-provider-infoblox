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

		Importer: &schema.ResourceImporter{},

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
		},
	}
}

func resourceCNAMERecordCreate(d *schema.ResourceData, m interface{}) error {
	dnsView := d.Get("dns_view").(string)
	canonical := d.Get("canonical").(string)
	alias := d.Get("alias").(string)

	comment := d.Get("comment").(string)
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
		}
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

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	recordCNAME, err := objMgr.CreateCNAMERecord(dnsView, canonical, alias, useTtl, ttl, comment, extAttrs)
	if err != nil {
		return fmt.Errorf("creation of CNAME Record under %s DNS View failed: %s", dnsView, err.Error())
	}

	d.SetId(recordCNAME.Ref)

	return nil
}

func resourceCNAMERecordGet(d *schema.ResourceData, m interface{}) error {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
		}
	}

	var tenantID string
	if tempVal, ok := extAttrs[eaNameForTenantId]; ok {
		tenantID = tempVal.(string)
	}
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	obj, err := objMgr.GetCNAMERecordByRef(d.Id())
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
	ttl := int(*obj.Ttl)
	if !*obj.UseTtl {
		ttl = ttlUndef
	}
	if err = d.Set("ttl", ttl); err != nil {
		return err
	}

	if obj.Ea != nil && len(obj.Ea) > 0 {
		// TODO: temporary scaffold, need to rework marshalling/unmarshalling of EAs
		//       (avoiding additional layer of keys ("value" key)
		eaMap := (map[string]interface{})(obj.Ea)
		ea, err := json.Marshal(eaMap)
		if err != nil {
			return err
		}
		if err = d.Set("ext_attrs", string(ea)); err != nil {
			return err
		}
	}
	if err = d.Set("dns_view", obj.View); err != nil {
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

	dnsView := d.Get("dns_view").(string)
	if d.HasChange("dns_view") {
		return fmt.Errorf("changing the value of 'dns_view' field is not allowed")
	}
	canonical := d.Get("canonical").(string)
	alias := d.Get("alias").(string)

	comment := d.Get("comment").(string)
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
		}
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
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	recordCNAME, err := objMgr.UpdateCNAMERecord(d.Id(), canonical, alias, useTtl, ttl, comment, extAttrs)
	if err != nil {
		return fmt.Errorf("updation of CNAME Record under %s DNS View failed: %s", dnsView, err.Error())
	}
	updateSuccessful = true

	d.SetId(recordCNAME.Ref)

	return nil
}

func resourceCNAMERecordDelete(d *schema.ResourceData, m interface{}) error {
	dnsView := d.Get("dns_view").(string)
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
		}
	}

	var tenantID string
	if tempVal, ok := extAttrs[eaNameForTenantId]; ok {
		tenantID = tempVal.(string)
	}
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	_, err := objMgr.DeleteCNAMERecord(d.Id())
	if err != nil {
		return fmt.Errorf("deletion of CNAME Record from dns view %s failed: %s", dnsView, err.Error())
	}
	d.SetId("")

	return nil
}
