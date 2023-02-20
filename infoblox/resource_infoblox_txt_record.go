package infoblox

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"regexp"

	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func resourceTXTRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceTXTRecordCreate,
		Read:   resourceTXTRecordGet,
		Update: resourceTXTRecordUpdate,
		Delete: resourceTXTRecordDelete,

		Importer: &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			"dns_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     defaultDNSView,
				Description: "DNS View in which the zone exists.",
			},
			"fqdn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "FQDN for the TXT-Record.",
			},
			"text": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Data to be associated with TXT_Record, this field can be empty.",
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
			"extattrs": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Extensible attributes of the TXT-record to be added/updated, as a map in JSON format",
			},
		},
	}
}

func resourceTXTRecordCreate(d *schema.ResourceData, m interface{}) error {
	dnsView := d.Get("dns_view").(string)
	fqdn := d.Get("fqdn").(string)

	fqdnPattern := `^[a-z]+\.[a-z0-9-]+\.[a-z]+$`

	if dnsView == "" {
		dnsView = defaultDNSView
	}

	valid, _ := regexp.MatchString(fqdnPattern, fqdn)

	if !valid {
		return fmt.Errorf("'fqdn is not in valid format'")
	}

	text := d.Get("text").(string)
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

	extAttrJSON := d.Get("extattrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
		}
	}

	var tenantID string
	tempVal, found := extAttrs[eaNameForTenantId]
	if found {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	newRecord, err := objMgr.CreateTXTRecord(dnsView, fqdn, text, ttl, useTtl, comment, extAttrs)

	if err != nil {
		return fmt.Errorf("error creating TXT-Record: %s", err.Error())
	}

	d.SetId(newRecord.Ref)

	if err = d.Set("text", newRecord.Text); err != nil {
		return err
	}

	return nil
}

func resourceTXTRecordGet(d *schema.ResourceData, m interface{}) error {
	extAttrJSON := d.Get("extattrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'extattrs' field: %s", err.Error())
		}
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
		return fmt.Errorf("failed getting TXT-Record: %s", err.Error())
	}

	if err = d.Set("text", obj.Text); err != nil {
		return err
	}

	ttl := int(obj.Ttl)
	if !obj.UseTtl {
		ttl = ttlUndef
	}
	if err = d.Set("ttl", ttl); err != nil {
		return err
	}

	if obj.Ea != nil && len(obj.Ea) > 0 {

		eaMap := (map[string]interface{})(obj.Ea)
		ea, err := json.Marshal(eaMap)
		if err != nil {
			return err
		}
		if err = d.Set("extattrs", string(ea)); err != nil {
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
			prevEa, _ := d.GetChange("extattrs")

			_ = d.Set("dns_view", prevDNSView.(string))
			_ = d.Set("fqdn", prevFQDN.(string))
			_ = d.Set("text", prevTEXT.(string))
			_ = d.Set("ttl", prevTTL.(int))
			_ = d.Set("comment", prevComment.(string))
			_ = d.Set("extattrs", prevEa.(string))
		}
	}()

	text := d.Get("text").(string)
	fqdn := d.Get("fqdn").(string)
	fqdnPattern := `^[a-z]+\.[a-z0-9-]+\.[a-z]+$`

	valid, _ := regexp.MatchString(fqdnPattern, fqdn)
	if !valid {
		return fmt.Errorf("fqdn not in valid format")
	}
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
	extAttrJSON := d.Get("extattrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'extattrs' field: %s", err.Error())
		}
	}

	var tenantID string
	tempVal, found := extAttrs[eaNameForTenantId]
	if found {
		tenantID = tempVal.(string)
	}
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	//Get the existing text value
	if text == "" {
		txtRec, err := objMgr.GetTXTRecordByRef(d.Id())
		if err != nil {
			return fmt.Errorf("failed getting TXT-Record: %s", err.Error())
		}
		text = txtRec.Text
	}

	if fqdn == "" {
		txtRec, err := objMgr.GetTXTRecordByRef(d.Id())
		if err != nil {
			return fmt.Errorf("failed getting TXT-Record: %s", err.Error())
		}
		fqdn = txtRec.Name
	}

	rec, err := objMgr.UpdateTXTRecord(
		d.Id(), fqdn, text, ttl, useTtl, comment, extAttrs)
	if err != nil {
		return fmt.Errorf("error updating TXT-Record: %s", err.Error())
	}
	updateSuccessful = true
	d.SetId(rec.Ref)

	return nil
}

func resourceTXTRecordDelete(d *schema.ResourceData, m interface{}) error {
	extAttrJSON := d.Get("extattrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'extattrs' field: %s", err.Error())
		}
	}
	var tenantID string
	tempVal, found := extAttrs[eaNameForTenantId]
	if found {
		tenantID = tempVal.(string)
	}
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	_, err := objMgr.DeleteTXTRecord(d.Id())
	if err != nil {
		return fmt.Errorf("deletion of TXT-Record failed: %s", err.Error())
	}
	d.SetId("")

	return nil
}
