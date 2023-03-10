package infoblox

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func resourceSRVRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceSRVRecordCreate,
		Read:   resourceSRVRecordGet,
		Update: resourceSRVRecordUpdate,
		Delete: resourceSRVRecordDelete,

		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"dns_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     defaultDNSView,
				Description: "DNS view which the zone does exist within",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Combination of service's name, protocol's name and zone's name",
			},
			"priority": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Configures the priority (0..65535) for this SRV-record.",
			},
			"weight": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Configures weight of the SRV-record, valid values are 0..65535.",
			},
			"port": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Configures port number (0..65535) for this SRV-record.",
			},
			"target": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Provides service for domain name in the SRV-record.",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     ttlUndef,
				Description: "TTL value for the SRV-record.",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Description of the SRV-record",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Extensible attributes of the SRV-record to be added/updated, as a map in JSON format.",
			},
		},
	}
}

func resourceSRVRecordCreate(d *schema.ResourceData, m interface{}) error {
	dnsView := d.Get("dns_view").(string)

	// the next group of parameters will be validated inside ibclient.CreateSRVRecord()
	name := d.Get("name").(string)
	priority := d.Get("priority").(int)
	weight := d.Get("weight").(int)
	port := d.Get("port").(int)
	target := d.Get("target").(string)

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

	newRecord, err := objMgr.CreateSRVRecord(
		dnsView, name, uint32(priority), uint32(weight), uint32(port), target, ttl, useTtl, comment, extAttrs)

	if err != nil {
		return fmt.Errorf("error creating SRV-record: %s", err.Error())
	}

	d.SetId(newRecord.Ref)

	return nil
}

func resourceSRVRecordGet(d *schema.ResourceData, m interface{}) error {
	extAttrJSON := d.Get("ext_attrs").(string)
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

	obj, err := objMgr.GetSRVRecordByRef(d.Id())
	if err != nil {
		return fmt.Errorf("failed getting SRV-Record: %s", err.Error())
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
	if err = d.Set("name", obj.Name); err != nil {
		return err
	}
	if err = d.Set("priority", obj.Priority); err != nil {
		return err
	}
	if err = d.Set("weight", obj.Weight); err != nil {
		return err
	}
	if err = d.Set("port", obj.Port); err != nil {
		return err
	}
	if err = d.Set("target", obj.Target); err != nil {
		return err
	}
	d.SetId(obj.Ref)

	return nil
}

func resourceSRVRecordUpdate(d *schema.ResourceData, m interface{}) error {
	var updateSuccessful bool
	defer func() {
		// Reverting the state back, in case of a failure,
		// otherwise Terraform will keep the values, which leaded to the failure,
		// in the state file.

		if !updateSuccessful {
			prevDNSView, _ := d.GetChange("dns_view")
			prevName, _ := d.GetChange("name")
			prevPriority, _ := d.GetChange("priority")
			prevWeight, _ := d.GetChange("weight")
			prevPort, _ := d.GetChange("port")
			prevTarget, _ := d.GetChange("target")
			prevTTL, _ := d.GetChange("ttl")
			prevComment, _ := d.GetChange("comment")
			prevEa, _ := d.GetChange("ext_attrs")

			_ = d.Set("dns_view", prevDNSView.(string))
			_ = d.Set("name", prevName.(string))
			_ = d.Set("priority", prevPriority.(uint32))
			_ = d.Set("weight", prevWeight.(uint32))
			_ = d.Set("port", prevPort.(uint32))
			_ = d.Set("target", prevTarget.(string))
			_ = d.Set("ttl", prevTTL.(uint32))
			_ = d.Set("comment", prevComment.(string))
			_ = d.Set("ext_attrs", prevEa.(string))
		}
	}()

	// the next group of parameters will be validated inside ibclient.UpdateSRVRecord()
	name := d.Get("name").(string)
	priority := d.Get("priority").(int)
	weight := d.Get("weight").(int)
	port := d.Get("port").(int)
	target := d.Get("target").(string)

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

	rec, err := objMgr.UpdateSRVRecord(
		d.Id(), name, uint32(priority), uint32(weight), uint32(port), target, ttl, useTtl, comment, extAttrs)
	if err != nil {
		return fmt.Errorf("error updating SRV-Record: %s", err.Error())
	}
	updateSuccessful = true
	d.SetId(rec.Ref)

	return nil
}

func resourceSRVRecordDelete(d *schema.ResourceData, m interface{}) error {
	extAttrJSON := d.Get("ext_attrs").(string)
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

	_, err := objMgr.DeleteSRVRecord(d.Id())
	if err != nil {
		return fmt.Errorf("deletion of MX-Record failed: %s", err.Error())
	}
	d.SetId("")

	return nil

}
