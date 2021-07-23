package infoblox

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
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
				Default:     "default",
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
				Default:     ttlUndef,
				Description: "TTL value for the A-record.",
			},
			"comment": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "Description of the A-record.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "Extensible attributes of the A-record to be added/updated, as a map in JSON format",
			},
		},
	}
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
	tempVal, found := extAttrs["Tenant ID"]
	if found {
		tenantID = tempVal.(string)
	}
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	newRecord, err := objMgr.CreateARecord(
		networkView, dnsView, fqdn, cidr, ipAddr, ttl, useTtl, comment, extAttrs)
	if err != nil {
		return fmt.Errorf("error creating A-record: %s", err.Error())
	}
	d.SetId(newRecord.Ref)
	return nil
}

func resourceARecordGet(d *schema.ResourceData, m interface{}) error {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
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
	if d.HasChange("network_view") {
		return fmt.Errorf("changing the value of 'network_view' field is not allowed")
	}
	if d.HasChange("dns_view") {
		return fmt.Errorf("changing the value of 'dns_view' field is not allowed")
	}

	cidr := d.Get("cidr").(string)
	fqdn := d.Get("fqdn").(string)
	ipAddr := d.Get("ip_addr").(string)
	if ipAddr == "" && cidr == "" {
		return fmt.Errorf("error updating A-record: either 'ip_addr' or 'cidr' value must not be empty.")
	}

	// If 'cidr' is unchanged, then making it empty to skip the update.
	// (This is to prevent record renewal for the case when 'cidr' is
	// used for IP address allocation, otherwise the address will be changing
	// during every 'update' operation).
	if !d.HasChange("cidr") {
		cidr = ""
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

	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
		}
	}

	var tenantID string
	tempVal, found := extAttrs["Tenant ID"]
	if found {
		tenantID = tempVal.(string)
	}
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	// Get the existing IP address
	if ipAddr == "" && cidr == "" {
		aRec, err := objMgr.GetARecordByRef(d.Id())
		if err != nil {
			return fmt.Errorf("failed getting A-record: %s", err.Error())
		}
		ipAddr = aRec.Ipv4Addr
	}

	rec, err := objMgr.UpdateARecord(
		d.Id(), fqdn, ipAddr, cidr, networkView, ttl, useTtl, comment, extAttrs)
	if err != nil {
		return fmt.Errorf("error updating A-record: %s", err.Error())
	}
	d.SetId(rec.Ref)
	return nil
}

func resourceARecordDelete(d *schema.ResourceData, m interface{}) error {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
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
