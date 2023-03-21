package infoblox

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func resourceMXRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceMXRecordCreate,
		Read:   resourceMXRecordGet,
		Update: resourceMXRecordUpdate,
		Delete: resourceMXRecordDelete,

		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"dns_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     defaultDNSView,
				Description: "DNS view which the zone does exist within",
			},
			"fqdn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "FQDN for the MX-record.",
			},
			"mail_exchanger": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A record used to specify mail server.",
			},
			"preference": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Configures the preference (0-65535) for this MX-record.",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     ttlUndef,
				Description: "TTL value for the MX-record.",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Description of the MX-Record.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Extensible attributes of the MX-record to be added/updated, as a map in JSON format.",
			},
		},
	}
}

func resourceMXRecordCreate(d *schema.ResourceData, m interface{}) error {
	dnsView := d.Get("dns_view").(string)

	fqdn := d.Get("fqdn").(string)
	if fqdn == "" {
		return fmt.Errorf("'fqdn' must not be empty")
	}

	mx := d.Get("mail_exchanger").(string)
	if mx == "" {
		return fmt.Errorf("'mail_exchanger' must not be empty")
	}

	tempInt := d.Get("preference").(int)
	if err := ibclient.CheckIntRange("preference", tempInt, 0, 65535); err != nil {
		return err
	}
	preference := uint32(tempInt)

	var ttl uint32
	useTtl := false
	tempTTL := d.Get("ttl").(int)
	if tempTTL >= 0 {
		useTtl = true
		ttl = uint32(tempTTL)
	} else if tempTTL != ttlUndef {
		return fmt.Errorf("TTL value must be 0 or higher")
	}

	comment := d.Get("comment").(string)

	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err)
		}
	}

	var tenantID string
	tempVal, found := extAttrs[eaNameForTenantId]
	if found {
		tenantID = tempVal.(string)
	}
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	newRecord, err := objMgr.CreateMXRecord(dnsView, fqdn, mx, preference, ttl, useTtl, comment, extAttrs)
	if err != nil {
		return fmt.Errorf("error creating MX-record: %s", err)
	}
	d.SetId(newRecord.Ref)

	return nil

}

func resourceMXRecordGet(d *schema.ResourceData, m interface{}) error {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err)
		}
	}
	var tenantID string
	tempVal, found := extAttrs[eaNameForTenantId]
	if found {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	obj, err := objMgr.GetMXRecordByRef(d.Id())
	if err != nil {
		return fmt.Errorf("failed getting MX-Record: %s", err)
	}

	ttl := int(obj.Ttl)
	if !obj.UseTtl {
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
	if err = d.Set("comment", obj.Comment); err != nil {
		return err
	}
	if err = d.Set("dns_view", obj.View); err != nil {
		return err
	}
	if err = d.Set("fqdn", obj.Fqdn); err != nil {
		return err
	}
	if err = d.Set("mail_exchanger", obj.MX); err != nil {
		return err
	}
	if err = d.Set("preference", obj.Preference); err != nil {
		return err
	}
	d.SetId(obj.Ref)

	return nil
}

func resourceMXRecordUpdate(d *schema.ResourceData, m interface{}) error {
	var updateSuccessful bool
	defer func() {
		// Reverting the state back, in case of a failure,
		// otherwise Terraform will keep the values, which leaded to the failure,
		// in the state file.

		if !updateSuccessful {
			prevDNSView, _ := d.GetChange("dns_view")
			prevFQDN, _ := d.GetChange("fqdn")
			prevMX, _ := d.GetChange("mail_exchanger")
			prevPreference, _ := d.GetChange("preference")
			prevTTL, _ := d.GetChange("ttl")
			prevComment, _ := d.GetChange("comment")
			prevEa, _ := d.GetChange("ext_attrs")

			_ = d.Set("dns_view", prevDNSView.(string))
			_ = d.Set("fqdn", prevFQDN.(string))
			_ = d.Set("mail_exchanger", prevMX.(string))
			_ = d.Set("preference", prevPreference.(int))
			_ = d.Set("ttl", prevTTL.(int))
			_ = d.Set("comment", prevComment.(string))
			_ = d.Set("ext_attrs", prevEa.(string))
		}
	}()

	dnsView := d.Get("dns_view").(string)
	fqdn := d.Get("fqdn").(string)
	mx := d.Get("mail_exchanger").(string)

	tempInt := d.Get("preference").(int)
	if err := ibclient.CheckIntRange("preference", tempInt, 0, 65535); err != nil {
		return err
	}
	preference := uint32(tempInt)

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
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err)
		}
	}

	var tenantID string
	tempVal, found := extAttrs[eaNameForTenantId]
	if found {
		tenantID = tempVal.(string)
	}
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	rec, err := objMgr.UpdateMXRecord(
		d.Id(), dnsView, fqdn, mx, preference, ttl, useTtl, comment, extAttrs)
	if err != nil {
		return fmt.Errorf("error updating MX-Record: %s", err)
	}
	updateSuccessful = true
	d.SetId(rec.Ref)

	return nil
}

func resourceMXRecordDelete(d *schema.ResourceData, m interface{}) error {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err)
		}
	}
	var tenantID string
	tempVal, found := extAttrs[eaNameForTenantId]
	if found {
		tenantID = tempVal.(string)
	}
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	_, err := objMgr.DeleteMXRecord(d.Id())
	if err != nil {
		return fmt.Errorf("deletion of MX-Record failed: %s", err)
	}
	d.SetId("")

	return nil
}
