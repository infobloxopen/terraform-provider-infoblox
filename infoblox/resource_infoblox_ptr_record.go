package infoblox

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func resourcePTRRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourcePTRRecordCreate,
		Read:   resourcePTRRecordGet,
		Update: resourcePTRRecordUpdate,
		Delete: resourcePTRRecordDelete,

		Importer: &schema.ResourceImporter{
			State: stateImporter,
		},

		Schema: map[string]*schema.Schema{
			"network_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
				Description: "Network view name of NIOS server.",
			},
			"cidr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The network address in cidr format under which record has to be created.",
			},
			"ip_addr": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "IPv4/IPv6 address for record creation. Set the field with valid IP for static allocation. If to be dynamically allocated set cidr field",
			},
			"dns_view": {
				Type:        schema.TypeString,
				Default:     "default",
				Optional:    true,
				Description: "Dns View under which the zone has been created.",
			},
			"ptrdname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The domain name in FQDN to which the record should point to.",
			},
			"record_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The name of the DNS PTR record in FQDN format",
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
				Description: "A description about PTR record.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "The Extensible attributes of PTR record to be added/updated, as a map in JSON format",
			},
		},
	}
}

func resourcePTRRecordCreate(d *schema.ResourceData, m interface{}) error {
	networkView := d.Get("network_view").(string)
	cidr := d.Get("cidr").(string)
	ipAddr := d.Get("ip_addr").(string)

	dnsView := d.Get("dns_view").(string)
	ptrdname := d.Get("ptrdname").(string)
	recordName := d.Get("record_name").(string)

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

	if recordName == "" {
		if ipAddr == "" && cidr == "" {
			return fmt.Errorf(
				"Creation of PTR record failed: 'ip_addr' or 'cidr' are mandatory in reverse mapping zone and 'record_name' is mandatory in forward mapping zone")
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

	if err = d.Set("ip_addr", recordPTR.Ipv4Addr); err != nil {
		return err
	}
	d.SetId(recordPTR.Ref)

	return nil
}

func resourcePTRRecordGet(d *schema.ResourceData, m interface{}) error {
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

	obj, err := objMgr.GetPTRRecordByRef(d.Id())
	if err != nil {
		return fmt.Errorf("getting PTR Record with ID %s failed : %s", d.Id(), err.Error())
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

	if val, ok := d.GetOk("network_view"); !ok || val.(string) == "" {
		if err = d.Set("network_view", "default"); err != nil {
			return err
		}
	}

	if err = d.Set("ptrdname", obj.PtrdName); err != nil {
		return err
	}

	var ipAddr string
	if obj.Ipv4Addr != "" {
		ipAddr = obj.Ipv4Addr
	} else {
		ipAddr = obj.Ipv6Addr
	}
	if err = d.Set("ip_addr", ipAddr); err != nil {
		return err
	}

	d.SetId(obj.Ref)

	return nil
}

func resourcePTRRecordUpdate(d *schema.ResourceData, m interface{}) error {
	networkView := d.Get("network_view").(string)
	if d.HasChange("network_view") {
		return fmt.Errorf("changing the value of 'network_view' field is not allowed")
	}
	dnsView := d.Get("dns_view").(string)
	if d.HasChange("dns_view") {
		return fmt.Errorf("changing the value of 'dns_view' field is not allowed")
	}
	ptrdname := d.Get("ptrdname").(string)
	recordName := d.Get("record_name").(string)

	ipAddr := d.Get("ip_addr").(string)
	cidr := d.Get("cidr").(string)
	// If 'cidr' is unchanged, then nothing to update here, making them empty to skip the update.
	// (This is to prevent record renewal for the case when 'cidr' is
	// used for IP address allocation, otherwise the address will be changing
	// during every 'update' operation).
	if !d.HasChange("cidr") {
		cidr = ""
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

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	// Retrive the IP of PTR record.
	// When IP is allocated using cidr and an empty IP is passed for updation
	if cidr == "" && ipAddr == "" {
		recordPTR, err := objMgr.GetPTRRecordByRef(d.Id())
		if err != nil {
			return fmt.Errorf("Getting PTR Record with ID %s failed : %s", d.Id(), err.Error())
		}

		ipv4 := recordPTR.Ipv4Addr
		ipv6 := recordPTR.Ipv6Addr
		if len(ipv4) > 0 {
			ipAddr = ipv4
		} else {
			ipAddr = ipv6
		}
	}

	recordPTRUpdated, err := objMgr.UpdatePTRRecord(d.Id(), networkView, ptrdname, recordName, cidr, ipAddr, useTtl, ttl, comment, extAttrs)
	if err != nil {
		return fmt.Errorf("Updating of PTR Record from dns view %s failed : %s", dnsView, err.Error())
	}

	if err = d.Set("ip_addr", recordPTRUpdated.Ipv4Addr); err != nil {
		return err
	}
	d.SetId(recordPTRUpdated.Ref)

	return nil
}

func resourcePTRRecordDelete(d *schema.ResourceData, m interface{}) error {
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

	_, err := objMgr.DeletePTRRecord(d.Id())
	if err != nil {
		return fmt.Errorf("Deletion of PTR Record from dns view %s failed : %s", dnsView, err.Error())
	}
	d.SetId("")

	return nil
}
