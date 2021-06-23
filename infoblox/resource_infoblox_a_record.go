package infoblox

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client"
)

func resourceARecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceARecordCreate,
		Read:   resourceARecordGet,
		Update: resourceARecordUpdate,
		Delete: resourceARecordDelete,

		Schema: map[string]*schema.Schema{
			"network_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Network view to use when allocating an IP address from a network dynamically. For static allocation, leave this field empty.",
			},
			"cidr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Network to allocate an IP address from, when the 'ip_addr' field is empty (dynamic allocation). The address is in CIDR format. For static allocation, leave this field empty.",
			},
			"dns_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
				Description: "DNS view which the zone does exist within.",
			},
			"fqdn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "FQDN for the A-record.",
			},
			"ip_addr": {
				Type:        schema.TypeString,
				Optional:    true, // making this optional because of possible dynalmic IP allocation (CIDR)
				Description: "IP address to associate with the A-record. For static allocation, set the field with a valid IP address. For dynamic allocation, leave this field empty and set 'cidr' and 'network_view' fields.",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "TTL value for the A-record.",
			},
			"comment": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "Description of the A-record.",
			},
			"extensible_attributes": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "Extensible attributes of the A-record to be added/updated, as a map in JSON format",
			},
		},
	}
}

func setTFFieldsForRecordA(d *schema.ResourceData, rec *ibclient.RecordA) error {
	d.SetId(rec.Ref)
	d.Set("ip_addr", rec.Ipv4Addr)

	return nil
}

func resourceARecordCreate(d *schema.ResourceData, m interface{}) error {
	networkView := d.Get("network_view").(string)
	cidr := d.Get("cidr").(string)
	dnsView := d.Get("dns_view").(string)
	fqdn := d.Get("fqdn").(string)
	ipAddr := d.Get("ip_addr").(string)
	if ipAddr == "" && cidr == "" {
		return fmt.Errorf("error creating A-record: 'ip_addr' is empty and either 'cidr' or 'network_view' values are absent.")
	}

	var ttl uint32
	tempVal, useTTL := d.GetOk("ttl")
	if useTTL {
		tempTTL := tempVal.(int)
		if tempTTL < 0 {
			return fmt.Errorf("TTL value must be 0 or higher")
		}
		ttl = uint32(tempTTL)
	}

	comment := d.Get("comment").(string)

	extAttrJSON := d.Get("extensible_attributes").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'extensible_attributes' field: %s", err.Error())
		}
	}

	var tenantID string
	tempVal, found := extAttrs["Tenant ID"]
	if found {
		tenantID = tempVal.(string)
	}
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	newRecord, err := objMgr.CreateARecord(
		networkView, dnsView, fqdn, cidr, ipAddr, ttl, useTTL, comment, extAttrs)
	if err != nil {
		return fmt.Errorf("error creating A-record: %s", err.Error())
	}
	return setTFFieldsForRecordA(d, newRecord)
}

func resourceARecordGet(d *schema.ResourceData, m interface{}) error {
	extAttrJSON := d.Get("extensible_attributes").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'extensible_attributes' field: %s", err.Error())
		}
	}
	var tenantID string
	tempVal, found := extAttrs["Tenant ID"]
	if found {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	obj, err := objMgr.GetARecordByRef(d.Id())
	if err != nil {
		return fmt.Errorf("failed getting A-record: %s", err.Error())
	}

	d.SetId(obj.Ref)

	return nil
}

func resourceARecordUpdate(d *schema.ResourceData, m interface{}) error {
	networkView := d.Get("network_view").(string)
	cidr := d.Get("cidr").(string)
	fqdn := d.Get("fqdn").(string)
	ipAddr := d.Get("ip_addr").(string)
	if ipAddr == "" && cidr == "" {
		return fmt.Errorf("error updating A-record: either 'ip_addr' or 'cidr' value must not be empty.")
	}

	if d.HasChange("dns_view") {
		return fmt.Errorf("changing the value of 'dns_view' field is not allowed")
	}

	// If both 'cidr' and 'ip_addr' are unchanged, then nothing to update here,
	// making them empty to skip the update.
	// (This is to prevent record renewal for the case when 'cidr' is
	// used for IP address allocation, otherwise the address will be changing
	// during every 'update' operation).
	if !d.HasChange("ip_addr") && !d.HasChange("cidr") {
		ipAddr = ""
		cidr = ""
	}

	var ttl uint32
	tempVal, useTTL := d.GetOk("ttl")
	if useTTL {
		tempTTL := tempVal.(int)
		if tempTTL < 0 {
			return fmt.Errorf("TTL value must be 0 or higher")
		}
		ttl = uint32(tempTTL)
	}

	comment := d.Get("comment").(string)

	extAttrJSON := d.Get("extensible_attributes").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'extensible_attributes' field: %s", err.Error())
		}
	}

	var tenantID string
	tempVal, found := extAttrs["Tenant ID"]
	if found {
		tenantID = tempVal.(string)
	}
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	rec, err := objMgr.UpdateARecord(
		d.Id(), fqdn, ipAddr, cidr, networkView, ttl, useTTL, comment, extAttrs)
	if err != nil {
		return fmt.Errorf("error updating A-record: %s", err.Error())
	}

	return setTFFieldsForRecordA(d, rec)
}

func resourceARecordDelete(d *schema.ResourceData, m interface{}) error {
	extAttrJSON := d.Get("extensible_attributes").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'extensible_attributes' field: %s", err.Error())
		}
	}
	var tenantID string
	tempVal, found := extAttrs["Tenant ID"]
	if found {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	_, err := objMgr.DeleteARecord(d.Id())
	if err != nil {
		return fmt.Errorf("deletion of A-record failed: %s", err.Error())
	}
	d.SetId("")

	return nil
}
