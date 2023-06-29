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

		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"network_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
				Default:     defaultDNSView,
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
				Computed:    true,
				Optional:    true,
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
	networkView, trimmed := checkAndTrimSpaces(d.Get("network_view").(string))
	if trimmed {
		return fmt.Errorf(errMsgFormatLeadingTrailingSpaces, "network_view")
	}
	if networkView == "" {
		networkView = defaultNetView
	}

	ipAddrSrcCounter := 0

	cidr, trimmed := checkAndTrimSpaces(d.Get("cidr").(string))
	if trimmed {
		return fmt.Errorf(errMsgFormatLeadingTrailingSpaces, "cidr")
	}
	if cidr != "" {
		ipAddrSrcCounter = ipAddrSrcCounter + 1
	}

	ipAddr, trimmed := checkAndTrimSpaces(d.Get("ip_addr").(string))
	if trimmed {
		return fmt.Errorf(errMsgFormatLeadingTrailingSpaces, "ip_addr")
	}
	if ipAddr != "" {
		ipAddrSrcCounter = ipAddrSrcCounter + 1
	}

	dnsViewName, trimmed := checkAndTrimSpaces(d.Get("dns_view").(string))
	if trimmed {
		return fmt.Errorf(errMsgFormatLeadingTrailingSpaces, "dns_view")
	}

	ptrdname, trimmed := checkAndTrimSpaces(d.Get("ptrdname").(string))
	if trimmed {
		return fmt.Errorf(errMsgFormatLeadingTrailingSpaces, "ptrdname")
	}

	recordName, trimmed := checkAndTrimSpaces(d.Get("record_name").(string))
	if trimmed {
		return fmt.Errorf(errMsgFormatLeadingTrailingSpaces, "record_name")
	}
	if recordName != "" {
		ipAddrSrcCounter = ipAddrSrcCounter + 1
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
	if tempVal, ok := extAttrs[eaNameForTenantId]; ok {
		tenantID = tempVal.(string)
	}

	if ipAddrSrcCounter == 0 {
		return fmt.Errorf(
			"'ip_addr' or 'cidr' are mandatory in reverse mapping zone and 'record_name' is mandatory in forward mapping zone")
	}

	if ipAddrSrcCounter != 1 {
		return fmt.Errorf(
			"only one of 'ip_addr', 'cidr' and 'record_name' must be defined")
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
		dnsViewName,
		ptrdname,
		recordName,
		cidr,
		ipAddr,
		useTtl,
		ttl,
		comment,
		extAttrs)
	if err != nil {
		return fmt.Errorf("creation of PTR-record under the DNS view '%s' failed: %s", dnsViewName, err)
	}

	// After reading a newly created object, IP address will be
	// set even if it is not specified directly in the configuration of the resource,
	if *recordPTR.Ipv4Addr != "" {
		ipAddr = *recordPTR.Ipv4Addr
	} else {
		ipAddr = *recordPTR.Ipv6Addr
	}

	if err = d.Set("ip_addr", ipAddr); err != nil {
		return err
	}
	if err = d.Set("record_name", recordPTR.Name); err != nil {
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

	if err = d.Set("comment", comment); err != nil {
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
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err)
		}
	}
	var tenantID string
	if tempVal, ok := extAttrs[eaNameForTenantId]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	obj, err := objMgr.GetPTRRecordByRef(d.Id())
	if err != nil {
		return fmt.Errorf("getting PTR-record with ID '%s' failed: %s", d.Id(), err)
	}

	ttl := int(*obj.Ttl)
	if !*obj.UseTtl {
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

	if err = d.Set("ptrdname", obj.PtrdName); err != nil {
		return err
	}

	var ipAddr string
	if *obj.Ipv4Addr != "" {
		ipAddr = *obj.Ipv4Addr
	} else {
		ipAddr = *obj.Ipv6Addr
	}
	if err = d.Set("ip_addr", ipAddr); err != nil {
		return err
	}
	if err = d.Set("record_name", obj.Name); err != nil {
		return err
	}

	d.SetId(obj.Ref)

	return nil
}

func resourcePTRRecordUpdate(d *schema.ResourceData, m interface{}) error {
	var updateSuccessful bool
	defer func() {
		// Reverting the state back, in case of a failure,
		// otherwise Terraform will keep the values, which leaded to the failure,
		// in the state file.
		if !updateSuccessful {
			prevNetView, _ := d.GetChange("network_view")
			prevDNSView, _ := d.GetChange("dns_view")
			prevPtrDName, _ := d.GetChange("ptrdname")
			prevName, _ := d.GetChange("record_name")
			prevIPAddr, _ := d.GetChange("ip_addr")
			prevCIDR, _ := d.GetChange("cidr")
			prevTTL, _ := d.GetChange("ttl")
			prevComment, _ := d.GetChange("comment")
			prevEa, _ := d.GetChange("ext_attrs")

			_ = d.Set("network_view", prevNetView.(string))
			_ = d.Set("dns_view", prevDNSView.(string))
			_ = d.Set("ptrdname", prevPtrDName.(string))
			_ = d.Set("record_name", prevName.(string))
			_ = d.Set("ip_addr", prevIPAddr.(string))
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
	dnsView := d.Get("dns_view").(string)
	if d.HasChange("dns_view") {
		return fmt.Errorf("changing the value of 'dns_view' field is not allowed")
	}
	ptrdname := d.Get("ptrdname").(string)

	ipAddrSrcChangesCounter := 0
	ipAddrSrcCounter := 0

	recordName, trimmed := checkAndTrimSpaces(d.Get("record_name").(string))
	if trimmed {
		return fmt.Errorf(errMsgFormatLeadingTrailingSpaces, "record_name")
	}
	if recordName != "" {
		ipAddrSrcCounter = ipAddrSrcCounter + 1
	}
	if d.HasChange("record_name") && recordName != "" {
		ipAddrSrcChangesCounter = ipAddrSrcChangesCounter + 1
	}

	ipAddr, trimmed := checkAndTrimSpaces(d.Get("ip_addr").(string))
	if trimmed {
		return fmt.Errorf(errMsgFormatLeadingTrailingSpaces, "ip_addr")
	}
	if ipAddr != "" {
		ipAddrSrcCounter = ipAddrSrcCounter + 1
	}
	if d.HasChange("ip_addr") && ipAddr != "" {
		recordName = "" // In go-client, 'record_name' takes precedence over 'cidr' and 'ip_addr', we need to disable it.
		ipAddrSrcChangesCounter = ipAddrSrcChangesCounter + 1
	}

	cidr, trimmed := checkAndTrimSpaces(d.Get("cidr").(string))
	if trimmed {
		return fmt.Errorf(errMsgFormatLeadingTrailingSpaces, "cidr")
	}
	if cidr != "" {
		ipAddrSrcCounter = ipAddrSrcCounter + 1
	}
	// If 'cidr' is unchanged, then nothing to update here, making them empty to skip the update.
	// (This is to prevent record renewal for the case when 'cidr' is
	// used for IP address allocation, otherwise the address will be changing
	// during every 'update' operation).
	if !d.HasChange("cidr") {
		cidr = ""
	} else {
		if cidr != "" {
			recordName = "" // In go-client, 'record_name' takes precedence over 'cidr' and 'ip_addr', we need to disable it.
			ipAddr = ""     // In go-client, 'ip_addr' takes precedence over 'cidr', we need to disable it.
			ipAddrSrcChangesCounter = ipAddrSrcChangesCounter + 1
		}
	}

	if ipAddrSrcCounter == 0 {
		return fmt.Errorf(
			"'ip_addr' or 'cidr' are mandatory in reverse mapping zone and 'record_name' is mandatory in forward mapping zone")
	}

	if ipAddrSrcChangesCounter > 1 {
		return fmt.Errorf("only one of 'cidr', 'ip_addr' and 'record_name' is allowed to be non-empty")
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

	// Retrieve the IP of PTR record.
	// When IP is allocated using cidr and an empty IP is passed for an update.
	if cidr == "" && ipAddr == "" {
		recordPTR, err := objMgr.GetPTRRecordByRef(d.Id())
		if err != nil {
			return fmt.Errorf("getting PTR-record with ID '%s' failed: %s", d.Id(), err)
		}

		ipv4 := recordPTR.Ipv4Addr
		ipv6 := recordPTR.Ipv6Addr
		if len(*ipv4) > 0 {
			ipAddr = *ipv4
		} else {
			ipAddr = *ipv6
		}
	}

	recordPTRUpdated, err := objMgr.UpdatePTRRecord(d.Id(), networkView, ptrdname, recordName, cidr, ipAddr, useTtl, ttl, comment, extAttrs)
	if err != nil {
		return fmt.Errorf("update operaiton failed for the PTR-record with ID '%s' under the DNS view '%s': %s", d.Id(), dnsView, err)
	}
	updateSuccessful = true
	d.SetId(recordPTRUpdated.Ref)

	// After reading a newly created object, IP address will be
	// set even if it is not specified directly in the configuration of the resource,
	if *recordPTRUpdated.Ipv4Addr != "" {
		ipAddr = *recordPTRUpdated.Ipv4Addr
	} else {
		ipAddr = *recordPTRUpdated.Ipv6Addr
	}

	if err = d.Set("ip_addr", ipAddr); err != nil {
		return err
	}
	if err = d.Set("record_name", recordPTRUpdated.Name); err != nil {
		return err
	}

	return nil
}

func resourcePTRRecordDelete(d *schema.ResourceData, m interface{}) error {
	dnsView := d.Get("dns_view").(string)

	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err)
		}
	}

	var tenantID string
	if tempVal, ok := extAttrs[eaNameForTenantId]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	_, err := objMgr.DeletePTRRecord(d.Id())
	if err != nil {
		if isNotFoundError(err) {
			d.SetId("")
		}
		return fmt.Errorf("deletion of PTR-record with ID '%s'under the DNS view '%s' failed: %s", d.Id(), dnsView, err)
	}

	return nil
}
