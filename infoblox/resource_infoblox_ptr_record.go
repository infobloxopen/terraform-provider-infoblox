package infoblox

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client"
)

func resourcePTRRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourcePTRRecordCreate,
		Read:   resourcePTRRecordGet,
		Update: resourcePTRRecordUpdate,
		Delete: resourcePTRRecordDelete,

		Schema: map[string]*schema.Schema{
			"network_view_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
				Description: "Network view name of NIOS server.",
			},
			"cidr": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The network address in cidr format under which record has to be created.",
			},
			"ip_addr": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPv4/IPv6 address for record creation. Set the field with valid IP for static allocation. If to be dynamically allocated set cidr field",
			},
			"dns_view": &schema.Schema{
				Type:        schema.TypeString,
				Default:     "default",
				Optional:    true,
				Description: "Dns View under which the zone has been created.",
			},
			"ptrdname": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The domain name in FQDN to which the record should point to.",
			},
			"record_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the DNS PTR record in FQDN format",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "TTL attribute value for the record.",
			},
			"comment": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "A description about PTR record.",
			},
			"extensible_attributes": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "The Extensible attributes of PTR record to be added/updated, as a map in JSON format",
			},
		},
	}
}

func resourcePTRRecordCreate(d *schema.ResourceData, m interface{}) error {

	networkView := d.Get("network_view_name").(string)
	cidr := d.Get("cidr").(string)
	ipAddr := d.Get("ip_addr").(string)

	dnsView := d.Get("dns_view").(string)
	ptrdname := d.Get("ptrdname").(string)
	recordName := d.Get("record_name").(string)

	comment := d.Get("comment").(string)
	extAttrJSON := d.Get("extensible_attributes").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'extensible_attributes' field: %s", err.Error())
		}
	}

	var tenantID string
	if tempVal, ok := extAttrs["Tenant ID"]; ok {
		tenantID = tempVal.(string)
	}

	if recordName == "" {
		if ipAddr == "" && cidr == "" {
			return fmt.Errorf(
				"Creation of PTR record failed: 'ip_addr' or 'cidr' are mandatory in reverse mapping zone and 'record_name' is mandatory in forward mapping zone")
		}
	}

	var ttl uint32
	tempVal, useTtl := d.GetOk("ttl")
	if useTtl {
		tempTtl := tempVal.(int)
		if tempTtl < 0 {
			return fmt.Errorf("TTL value must be 0 or higher")
		}
		ttl = uint32(tempTtl)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	recordPTR, err := objMgr.CreatePTRRecord(
		networkView,
		dnsView,
		ptrdname,
		recordName,
		cidr,
		ipAddr,
		useTtl,
		ttl,
		comment,
		extAttrs)
	if err != nil {
		return fmt.Errorf("Creation of PTR Record under %s DNS View failed : %s", dnsView, err.Error())
	}

	d.SetId(recordPTR.Ref)
	return nil
}

func resourcePTRRecordGet(d *schema.ResourceData, m interface{}) error {

	extAttrJSON := d.Get("extensible_attributes").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'extensible_attributes' field: %s", err.Error())
		}
	}
	var tenantID string
	if tempVal, ok := extAttrs["Tenant ID"]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	recordPTR, err := objMgr.GetPTRRecordByRef(d.Id())
	if err != nil {
		return fmt.Errorf("Getting PTR Record with ID %s failed : %s", d.Id(), err.Error())
	}
	d.SetId(recordPTR.Ref)
	return nil
}

func resourcePTRRecordUpdate(d *schema.ResourceData, m interface{}) error {

	dnsView := d.Get("dns_view").(string)
	if d.HasChange("dns_view") {
		return fmt.Errorf("changing the value of 'dns_view' field is not allowed")
	}
	ptrdname := d.Get("ptrdname").(string)
	recordName := d.Get("record_name").(string)

	cidr := d.Get("cidr").(string)
	ipAddr := d.Get("ip_addr").(string)

	comment := d.Get("comment").(string)
	extAttrJSON := d.Get("extensible_attributes").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'extensible_attributes' field: %s", err.Error())
		}
	}

	var tenantID string
	if tempVal, ok := extAttrs["Tenant ID"]; ok {
		tenantID = tempVal.(string)
	}

	if recordName == "" {
		if ipAddr == "" && cidr == "" {
			return fmt.Errorf(
				"Creation of PTR record failed: 'ip_addr' or 'cidr' are mandatory in reverse mapping zone and 'record_name' is mandatory in forward mapping zone")
		}
	}

	var ttl uint32
	tempVal, useTtl := d.GetOk("ttl")
	if useTtl {
		tempTtl := tempVal.(int)
		if tempTtl < 0 {
			return fmt.Errorf("TTL value must be 0 or higher")
		}
		ttl = uint32(tempTtl)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	recordPTRUpdated, err := objMgr.UpdatePTRRecord(d.Id(), ptrdname, recordName, ipAddr, useTtl, ttl, comment, extAttrs)
	if err != nil {
		return fmt.Errorf("Updating of PTR Record from dns view %s failed : %s", dnsView, err.Error())
	}

	d.SetId(recordPTRUpdated.Ref)
	return nil
}

func resourcePTRRecordDelete(d *schema.ResourceData, m interface{}) error {

	dnsView := d.Get("dns_view").(string)

	extAttrJSON := d.Get("extensible_attributes").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'extensible_attributes' field: %s", err.Error())
		}
	}

	var tenantID string
	if tempVal, ok := extAttrs["Tenant ID"]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	_, err := objMgr.DeletePTRRecord(d.Id())
	if err != nil {
		return fmt.Errorf("Deletion of PTR Record from dns view %s failed : %s", dnsView, err.Error())
	}
	d.SetId("")
	return nil
}
