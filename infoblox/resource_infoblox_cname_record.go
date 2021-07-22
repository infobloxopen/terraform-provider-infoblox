package infoblox

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func resourceCNAMERecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceCNAMERecordCreate,
		Read:   resourceCNAMERecordGet,
		Update: resourceCNAMERecordUpdate,
		Delete: resourceCNAMERecordDelete,

		Schema: map[string]*schema.Schema{
			"dns_view": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
				Description: "Dns View under which the zone has been created.",
			},
			"canonical": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Canonical name in FQDN format.",
			},
			"alias": &schema.Schema{
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
				Default:     "",
				Optional:    true,
				Description: "A description about CNAME record.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
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
	if tempVal, ok := extAttrs["Tenant ID"]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	recordCNAME, err := objMgr.CreateCNAMERecord(dnsView, canonical, alias, useTtl, ttl, comment, extAttrs)
	if err != nil {
		return fmt.Errorf("Creation of CNAME Record under %s DNS View failed : %s", dnsView, err.Error())
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
	if tempVal, ok := extAttrs["Tenant ID"]; ok {
		tenantID = tempVal.(string)
	}
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	recordCNAME, err := objMgr.GetCNAMERecordByRef(d.Id())
	if err != nil {
		return fmt.Errorf("Getting CNAME Record with ID: %s failed : %s", d.Id(), err.Error())
	}
	d.SetId(recordCNAME.Ref)
	return nil
}

func resourceCNAMERecordUpdate(d *schema.ResourceData, m interface{}) error {

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
	if tempVal, ok := extAttrs["Tenant ID"]; ok {
		tenantID = tempVal.(string)
	}
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	recordCNAME, err := objMgr.UpdateCNAMERecord(d.Id(), canonical, alias, useTtl, ttl, comment, extAttrs)
	if err != nil {
		return fmt.Errorf("Updation of CNAME Record under %s DNS View failed : %s", dnsView, err.Error())
	}

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
	if tempVal, ok := extAttrs["Tenant ID"]; ok {
		tenantID = tempVal.(string)
	}
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	_, err := objMgr.DeleteCNAMERecord(d.Id())
	if err != nil {
		return fmt.Errorf("Deletion of CNAME Record from dns view %s failed : %s", dnsView, err.Error())
	}
	d.SetId("")
	return nil
}
