package infoblox

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func resourceTXTRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceTXTRecordCreate,
		Read:   resourceTXTRecordGet,
		Update: resourceTXTRecordUpdate,
		Delete: resourceTXTRecordDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the TXT record.",
			},
			"text": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The text for the record to contain.",
			},
			"dns_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
				Description: "Dns View under which the zone has been created.",
				ForceNew:    true,
			},
			"ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     ttlUndef,
				Description: "TTL value for the TXT record.",
			},
			"comment": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "Description of the TXT record.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "Extensible attributes of the TXT record to be added/updated, as a map in JSON format",
			},
		},
	}
}

func resourceTXTRecordCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning to create TXT record from  required network block", resourceTXTRecordIDString(d))

	//This is for record Name
	recordName := d.Get("name").(string)
	text := d.Get("text").(string)
	dnsView := d.Get("dns_view").(string)

	comment := d.Get("comment").(string)
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
		}
	}

	var tenantID string
	if tempVal, ok := extAttrs["Tenant ID"]; ok {
		tenantID = tempVal.(string)
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

	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	recordTXT, err := objMgr.CreateTXTRecord(
		recordName,
		text,
		dnsView,
		useTtl,
		ttl,
		comment,
		extAttrs)
	if err != nil {
		return fmt.Errorf("Error creating TXT record: %s", err)
	}

	d.SetId(recordTXT.Ref)

	log.Printf("[DEBUG] %s: Creation of TXT Record complete", resourceTXTRecordIDString(d))
	return nil
}

func resourceTXTRecordGet(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Begining to Get TXT Record", resourceTXTRecordIDString(d))

	dnsView := d.Get("dns_view").(string)

	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
		}
	}

	var tenantID string
	if tempVal, ok := extAttrs["Tenant ID"]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	obj, err := objMgr.GetTXTRecordByRef(d.Id())
	if err != nil {
		return fmt.Errorf("Getting TXT record failed from dns view (%s) : %s", dnsView, err)
	}
	d.SetId(obj.Ref)
	log.Printf("[DEBUG] %s: Completed reading required TXT Record ", resourceTXTRecordIDString(d))
	return nil
}

func resourceTXTRecordUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning to create TXT record from  required network block", resourceTXTRecordIDString(d))

	dnsView := d.Get("dns_view").(string)
	if d.HasChange("dns_view") {
		return fmt.Errorf("changing the value of 'dns_view' field is not allowed")
	}

	//This is for record Name
	recordName := d.Get("name").(string)
	text := d.Get("text").(string)

	comment := d.Get("comment").(string)
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
		}
	}

	var tenantID string
	if tempVal, ok := extAttrs["Tenant ID"]; ok {
		tenantID = tempVal.(string)
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

	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	recordTXTUpdated, err := objMgr.UpdateTXTRecord(d.Id(), recordName, text, useTtl, ttl, comment, extAttrs)
	if err != nil {
		return fmt.Errorf("Updating of TXT Record from dns view %s failed : %s", dnsView, err.Error())
	}

	d.SetId(recordTXTUpdated.Ref)
	return nil
}

func resourceTXTRecordDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Deletion of TXT Record", resourceTXTRecordIDString(d))

	dnsView := d.Get("dns_view").(string)

	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
		}
	}

	var tenantID string
	if tempVal, ok := extAttrs["Tenant ID"]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	_, err := objMgr.DeleteTXTRecord(d.Id())
	if err != nil {
		return fmt.Errorf("Deletion of TXT Record failed from dns view(%s) : %s", dnsView, err)
	}
	d.SetId("")

	log.Printf("[DEBUG] %s: Deletion of TXT Record complete", resourceTXTRecordIDString(d))
	return nil
}

type resourceTXTRecordIDStringInterface interface {
	Id() string
}

func resourceTXTRecordIDString(d resourceTXTRecordIDStringInterface) string {
	id := d.Id()
	if id == "" {
		id = "<new resource>"
	}
	return fmt.Sprintf("infoblox_txt_record (ID = %s)", id)
}
