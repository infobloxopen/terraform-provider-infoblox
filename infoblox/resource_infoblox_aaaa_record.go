package infoblox

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func resourceAAAARecord() *schema.Resource {
	return &schema.Resource{
		Create:   resourceAAAARecordCreate,
		Read:     resourceAAAARecordGet,
		Update:   resourceAAAARecordUpdate,
		Delete:   resourceAAAARecordDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"network_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Network view name of NIOS server.",
			},
			"dns_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     defaultDNSView,
				Description: "Dns View under which the zone has been created.",
			},
			"cidr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The network address in cidr format under which record has to be created.",
			},
			"ipv6_addr": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "IPv6 address for record creation. Set the field with valid IP for static allocation. If to be dynamically allocated set cidr field",
			},
			"fqdn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the AAAA record in FQDN format.",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     ttlUndef,
				Description: "TTL attribute value for the record.",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "A description about AAAA record.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The Extensible attributes of AAAA record to be added/updated, as a map in JSON format",
			},
		},
	}
}

func resourceAAAARecordCreate(d *schema.ResourceData, m interface{}) error {
	networkView := d.Get("network_view").(string)
	if networkView == "" {
		networkView = defaultNetView
	}
	cidr := d.Get("cidr").(string)
	ipv6Addr := d.Get("ipv6_addr").(string)

	dnsViewName := d.Get("dns_view").(string)
	fqdn := d.Get("fqdn").(string)

	comment := d.Get("comment").(string)
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
		}
	}

	var tenantID string
	if tempVal, ok := extAttrs[eaNameForTenantId]; ok {
		tenantID = tempVal.(string)
	}

	if ipv6Addr == "" && cidr == "" {
		return fmt.Errorf(
			"creation of AAAA record failed: 'ipv6_addr' or 'cidr' are mandatory")
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

	recordAAAA, err := objMgr.CreateAAAARecord(
		networkView,
		dnsViewName,
		fqdn,
		cidr,
		ipv6Addr,
		useTtl,
		ttl,
		comment,
		extAttrs)
	if err != nil {
		return fmt.Errorf("creation of AAAA Record under %s DNS View failed: %s", dnsViewName, err.Error())
	}
	d.SetId(recordAAAA.Ref)

	if err = d.Set("ipv6_addr", recordAAAA.Ipv6Addr); err != nil {
		return err
	}
	if val, ok := d.GetOk("network_view"); !ok || val.(string) == "" {
		dnsViewObj, err := objMgr.GetDNSView(dnsViewName)
		if err != nil {
			return fmt.Errorf(
				"error while retrieving information about DNS view '%s': %s",
				dnsViewName, err)
		}
		if err = d.Set("network_view", dnsViewObj.NetworkView); err != nil {
			return err
		}
	}

	return nil
}

func resourceAAAARecordGet(d *schema.ResourceData, m interface{}) error {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
		}
	}
	var tenantID string
	if tempVal, ok := extAttrs[eaNameForTenantId]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	obj, err := objMgr.GetAAAARecordByRef(d.Id())
	if err != nil {
		return fmt.Errorf("getting AAAA Record with ID: %s failed: %s", d.Id(), err.Error())
	}
	if err = d.Set("ipv6_addr", obj.Ipv6Addr); err != nil {
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
		dnsView, err := objMgr.GetDNSView(obj.View)
		if err != nil {
			return fmt.Errorf(
				"error while retrieving information about DNS view '%s': %s",
				obj.View, err)
		}
		if err = d.Set("network_view", dnsView.NetworkView); err != nil {
			return err
		}
	}

	if err = d.Set("fqdn", obj.Name); err != nil {
		return err
	}

	d.SetId(obj.Ref)

	return nil
}

func resourceAAAARecordUpdate(d *schema.ResourceData, m interface{}) error {
	var updateSuccessful bool
	defer func() {
		// Reverting the state back, in case of a failure,
		// otherwise Terraform will keep the values, which leaded to the failure,
		// in the state file.
		if !updateSuccessful {
			prevNetView, _ := d.GetChange("network_view")
			prevDNSView, _ := d.GetChange("dns_view")
			prevFQDN, _ := d.GetChange("fqdn")
			prevIPAddr, _ := d.GetChange("ipv6_addr")
			prevCIDR, _ := d.GetChange("cidr")
			prevTTL, _ := d.GetChange("ttl")
			prevComment, _ := d.GetChange("comment")
			prevEa, _ := d.GetChange("ext_attrs")

			_ = d.Set("network_view", prevNetView.(string))
			_ = d.Set("dns_view", prevDNSView.(string))
			_ = d.Set("fqdn", prevFQDN.(string))
			_ = d.Set("ipv6_addr", prevIPAddr.(string))
			_ = d.Set("cidr", prevCIDR.(string))
			_ = d.Set("ttl", prevTTL.(int))
			_ = d.Set("comment", prevComment.(string))
			_ = d.Set("ext_attrs", prevEa.(string))

		}
	}()

	networkView := d.Get("network_view").(string)
	if d.HasChange("network_view") {
		return fmt.Errorf("changing the value of 'network_view' field is not allowed")
	}
	cidr := d.Get("cidr").(string)
	ipv6Addr := d.Get("ipv6_addr").(string)

	// If 'cidr' is unchanged, then nothing to update here, making them empty to skip the update.
	// (This is to prevent record renewal for the case when 'cidr' is
	// used for IP address allocation, otherwise the address will be changing
	// during every 'update' operation).
	if !d.HasChange("cidr") {
		cidr = ""
	}

	dnsView := d.Get("dns_view").(string)
	if d.HasChange("dns_view") {
		return fmt.Errorf("changing the value of 'dns_view' field is not allowed")
	}
	fqdn := d.Get("fqdn").(string)

	comment := d.Get("comment").(string)
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
		}
	}

	var tenantID string
	if tempVal, ok := extAttrs[eaNameForTenantId]; ok {
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

	// Get the existing IP address
	if ipv6Addr == "" && cidr == "" {
		aaaaRec, err := objMgr.GetAAAARecordByRef(d.Id())
		if err != nil {
			return fmt.Errorf("getting AAAA Record with ID: %s failed: %s", d.Id(), err.Error())
		}
		ipv6Addr = aaaaRec.Ipv6Addr
	}

	recordAAAA, err := objMgr.UpdateAAAARecord(
		d.Id(),
		networkView,
		fqdn,
		ipv6Addr,
		cidr,
		useTtl,
		ttl,
		comment,
		extAttrs)
	if err != nil {
		return fmt.Errorf("updation of AAAA Record under %s DNS View failed: %s", dnsView, err.Error())
	}
	updateSuccessful = true
	d.SetId(recordAAAA.Ref)

	if err = d.Set("ipv6_addr", recordAAAA.Ipv6Addr); err != nil {
		return err
	}

	return nil
}

func resourceAAAARecordDelete(d *schema.ResourceData, m interface{}) error {
	dnsView := d.Get("dns_view").(string)

	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
		}
	}

	var tenantID string
	if tempVal, ok := extAttrs[eaNameForTenantId]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	_, err := objMgr.DeleteAAAARecord(d.Id())
	if err != nil {
		return fmt.Errorf("deletion of AAAA Record from dns view %s failed: %s", dnsView, err.Error())
	}
	d.SetId("")

	return nil
}
