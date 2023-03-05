package infoblox

import (
	"encoding/json"
	"fmt"
	"regexp"

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
				Description: "FQDN for the MX Record.",
			},
			"mail_exchanger": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A record used to specify mail server.",
			},
			"preference": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Configures the preference (0-65535) for this MX record.",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "TTL value for the MX-record.",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Description of the MX-Record.",
			},
			"extattrs": {
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
	fqdnPattern := `^[a-z]+\.[a-z0-9-]+\.[a-z]+$`

	if dnsView == "" {
		dnsView = defaultDNSView
	}

	valid, _ := regexp.MatchString(fqdnPattern, fqdn)
	if !valid {
		return fmt.Errorf("'fqdn' is not in valid format")
	}

	mx := d.Get("mail_exchanger").(string)

	if mx == "" {
		return fmt.Errorf("'mail_exchanger' must not be empty")
	}

	var priority uint32
	tempPref := d.Get("preference")
	tempPriority := tempPref.(int)

	if tempPriority >= 0 || tempPriority < 65535 {
		priority = uint32(tempPriority)
	} else if tempPriority < 0 || tempPriority > 65535 {
		return fmt.Errorf("'preference' ranges between 0 to 65535")
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

	newRecord, err := objMgr.CreateMXRecord(dnsView, fqdn, mx, priority, ttl, useTtl, comment, extAttrs)

	if err != nil {
		return fmt.Errorf("error creating MX-record: %s", err.Error())
	}
	d.SetId(newRecord.Ref)

	return nil

}

func resourceMXRecordGet(d *schema.ResourceData, m interface{}) error {
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

	obj, err := objMgr.GetMXRecordByRef(d.Id())
	if err != nil {
		return fmt.Errorf("failed getting MX-Record: %s", err.Error())
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
	if err = d.Set("fqdn", obj.Fqdn); err != nil {
		return err
	}
	if err = d.Set("mail_exchanger", obj.MX); err != nil {
		return err
	}
	if err = d.Set("preference", obj.Priority); err != nil {
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
			prevPriority, _ := d.GetChange("preference")
			prevTTL, _ := d.GetChange("ttl")
			prevComment, _ := d.GetChange("comment")
			prevEa, _ := d.GetChange("extattrs")

			_ = d.Set("dns_view", prevDNSView.(string))
			_ = d.Set("fqdn", prevFQDN.(string))
			_ = d.Set("mail_exchanger", prevMX.(string))
			_ = d.Set("preference", prevPriority.(uint32))
			_ = d.Set("ttl", prevTTL.(uint32))
			_ = d.Set("comment", prevComment.(string))
			_ = d.Set("extattrs", prevEa.(string))
		}
	}()

	fqdn := d.Get("fqdn").(string)
	dnsView := d.Get("dns_view").(string)
	mx := d.Get("mail_exchanger").(string)
	fqdnPattern := `^[a-z]+\.[a-z0-9-]+\.[a-z]+$`

	valid, _ := regexp.MatchString(fqdnPattern, fqdn)
	if !valid {
		return fmt.Errorf("'fqdn' not in valid format")
	}

	var priority uint32
	tempPref := d.Get("preference")
	tempPriority := tempPref.(int)

	if tempPriority >= 0 || tempPriority < 65535 {
		priority = uint32(tempPriority)
	} else if tempPriority < 0 || tempPriority > 65535 {
		return fmt.Errorf("'preference' ranges between 0 to 65535")
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

	//Get the existing mail_exchanger
	if mx == "" {
		mxRec, err := objMgr.GetMXRecordByRef(d.Id())
		if err != nil {
			return fmt.Errorf("failed getting MX-record: %s", err.Error())
		}
		mx = mxRec.MX
	}

	if fqdn == "" {
		mxRec, err := objMgr.GetMXRecordByRef(d.Id())
		if err != nil {
			return fmt.Errorf("failed getting MX-Record: %s", err.Error())
		}
		fqdn = mxRec.Fqdn
	}

	if dnsView == "" {
		mxRec, err := objMgr.GetMXRecordByRef(d.Id())
		if err != nil {
			return fmt.Errorf("failed getting MX-Record: %s", err.Error())
		}
		dnsView = mxRec.View
	}

	rec, err := objMgr.UpdateMXRecord(
		d.Id(), dnsView, fqdn, mx, ttl, useTtl, comment, priority, extAttrs)
	if err != nil {
		return fmt.Errorf("error updating MX-Record: %s", err.Error())
	}
	updateSuccessful = true
	d.SetId(rec.Ref)

	return nil
}

func resourceMXRecordDelete(d *schema.ResourceData, m interface{}) error {
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

	_, err := objMgr.DeleteMXRecord(d.Id())
	if err != nil {
		return fmt.Errorf("deletion of MX-Record failed: %s", err.Error())
	}
	d.SetId("")

	return nil
}
