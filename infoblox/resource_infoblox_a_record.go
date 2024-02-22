package infoblox

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func resourceARecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceARecordCreate,
		Read:   resourceARecordGet,
		Update: resourceARecordUpdate,
		Delete: resourceARecordDelete,
		Importer: &schema.ResourceImporter{
			State: resourceARecordImport,
		},

		Schema: map[string]*schema.Schema{
			"dns_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     defaultDNSView,
				Description: "DNS view which the zone does exist within.",
			},
			"fqdn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "FQDN for the A-record.",
			},
			"ip_addr": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true, // making this optional because of possible dynamic IP allocation (CIDR)
				Description: "IP address to associate with the A-record. For static allocation, set the field with a valid IP address. For dynamic allocation, leave this field empty and set 'cidr' and 'network_view' fields.",
			},
			"network_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Network view to use when allocating an IP address from a network dynamically. For static allocation, leave this field empty.",
			},
			"cidr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Network to allocate an IP address from, when the 'ip_addr' field is empty (dynamic allocation). The address is in CIDR format. For static allocation, leave this field empty.",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     ttlUndef,
				Description: "TTL value for the A-record.",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Description of the A-record.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Extensible attributes of the A-record to be added/updated, as a map in JSON format",
			},
			"internal_id": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "Internal ID of an object at NIOS side," +
					" used by Infoblox Terraform plugin to search for a NIOS's object" +
					" which corresponds to the Terraform resource.",
			},
			"ref": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "NIOS object's reference, not to be set by a user.",
			},
		},
	}
}

//func getOrFindRecordA(d *schema.ResourceData, m interface{}) (
//	hostRec interface{},
//	err error) {
//
//	var (
//		ref         string
//		actualIntId *internalResourceId
//	)
//
//	if r, found := d.GetOk("ref"); found {
//		ref = r.(string)
//		fmt.Printf("Reference: %v", ref)
//	} else {
//		_, ref = getAltIdFields(d.Id())
//		fmt.Printf("Reference in else: %v", ref)
//	}
//
//	if id, found := d.GetOk("internal_id"); !found {
//		return nil, fmt.Errorf("internal_id value is required for the resource but it is not defined")
//	} else {
//		actualIntId = newInternalResourceIdFromString(id.(string))
//		if actualIntId == nil {
//			return nil, fmt.Errorf("internal_id value is not in a proper format")
//		}
//	}
//
//	// TODO: use proper Tenant ID
//	objMgr := ibclient.NewObjectManager(m.(ibclient.IBConnector), "Terraform", "")
//	res, err := objMgr.SearchDnsObjectByAltId("A", ref, actualIntId.String(), eaNameForInternalId)
//	// Json marshall
//	var record []ibclient.RecordA
//	recJson, _ := json.Marshal(res)
//	json.Unmarshal(recJson, &record)
//	fmt.Printf("Debug point a: val %v", record)
//	if err != nil {
//		return nil, fmt.Errorf("failng to fetch %v", err)
//	}
//
//	return record[0], nil
//}

func resourceARecordCreate(d *schema.ResourceData, m interface{}) error {

	if intId := d.Get("internal_id"); intId.(string) != "" {
		return fmt.Errorf("the value of 'internal_id' field must not be set manually")
	}

	networkView := d.Get("network_view").(string)
	if networkView == "" {
		networkView = defaultNetView
	}
	cidr := d.Get("cidr").(string)
	dnsViewName := d.Get("dns_view").(string)
	fqdn := d.Get("fqdn").(string)
	ipAddr := d.Get("ip_addr").(string)
	if ipAddr == "" && cidr == "" {
		return fmt.Errorf("either of 'ip_addr' and 'cidr' values is required")
	}

	if ipAddr != "" && cidr != "" {
		return fmt.Errorf("only one of 'ip_addr' and 'cidr' values is allowed to be defined")
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
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)

	// Generate internal ID and add it to the extensible attributes
	internalId := generateInternalId()
	extAttrs[eaNameForInternalId] = internalId.String()

	if err != nil {
		return err
	}

	var tenantID string
	if tempVal, found := extAttrs[eaNameForTenantId]; found {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	newRecord, err := objMgr.CreateARecord(
		networkView,
		dnsViewName,
		fqdn,
		cidr,
		ipAddr,
		ttl,
		useTtl,
		comment,
		extAttrs)
	if err != nil {
		return fmt.Errorf("creation of A-record under DNS view '%s' failed: %w", dnsViewName, err)
	}
	//d.SetId(newRecord.Ref)
	d.SetId(internalId.String())
	if err = d.Set("ref", newRecord.Ref); err != nil {
		return err
	}
	if err = d.Set("internal_id", internalId.String()); err != nil {
		return err
	}

	if err = d.Set("ip_addr", newRecord.Ipv4Addr); err != nil {
		return err
	}
	if val, ok := d.GetOk("network_view"); !ok || val.(string) == "" {
		dnsViewObj, err := objMgr.GetDNSView(dnsViewName)
		if err != nil {
			return fmt.Errorf(
				"error while retrieving information about DNS view '%s': %w",
				dnsViewName, err)
		}
		if err = d.Set("network_view", dnsViewObj.NetworkView); err != nil {
			return err
		}
	}

	return nil
}

func resourceARecordGet(d *schema.ResourceData, m interface{}) error {
	var ttl int
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}

	var tenantID string
	var (
		recordA *ibclient.RecordA
		aList   []ibclient.RecordA
	)
	tempVal, found := extAttrs[eaNameForTenantId]
	if found {
		tenantID = tempVal.(string)
	}
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	obj, err := getOrFindRecord("A", d, m)
	fmt.Printf("Some value %v", obj)
	//obj, err := objMgr.GetARecordByRef(d.Id())

	if err != nil {
		//return nil
		d.SetId("")
		return nil
		//return fmt.Errorf("failed getting A-record: %w", err)
	}

	recJson, _ := json.Marshal(obj)
	json.Unmarshal(recJson, &aList)
	byteObj, _ := json.Marshal(aList[0])
	err = json.Unmarshal(byteObj, &recordA)
	if err != nil {
		return fmt.Errorf("unable to unmarshal %v", err.Error())
	}

	if err = d.Set("ip_addr", recordA.Ipv4Addr); err != nil {
		return err
	}

	if recordA.Ttl != nil {
		ttl = int(*recordA.Ttl)
	}
	if !*recordA.UseTtl {
		ttl = ttlUndef
	}
	if err = d.Set("ttl", ttl); err != nil {
		return err
	}

	omittedEAs := omitEAs(recordA.Ea, extAttrs)

	if omittedEAs != nil && len(omittedEAs) > 0 {
		eaJSON, err := terraformSerializeEAs(omittedEAs)
		if err != nil {
			return err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return err
		}
	}

	if err = d.Set("comment", recordA.Comment); err != nil {
		return err
	}

	if err = d.Set("dns_view", recordA.View); err != nil {
		return err
	}
	if val, ok := d.GetOk("network_view"); !ok || val.(string) == "" {
		dnsView, err := objMgr.GetDNSView(recordA.View)
		if err != nil {
			return fmt.Errorf(
				"error while retrieving information about DNS view '%s': %w",
				recordA.View, err)
		}
		if err = d.Set("network_view", dnsView.NetworkView); err != nil {
			return err
		}
	}

	if err = d.Set("fqdn", recordA.Name); err != nil {
		return err
	}

	d.SetId(recordA.Ref)

	return nil
}

func resourceARecordUpdate(d *schema.ResourceData, m interface{}) error {
	var updateSuccessful bool
	defer func() {
		// Reverting the state back, in case of a failure,
		// otherwise Terraform will keep the values, which leaded to the failure,
		// in the state file.
		if !updateSuccessful {
			prevNetView, _ := d.GetChange("network_view")
			prevDNSView, _ := d.GetChange("dns_view")
			prevFQDN, _ := d.GetChange("fqdn")
			prevIPAddr, _ := d.GetChange("ip_addr")
			prevCIDR, _ := d.GetChange("cidr")
			prevTTL, _ := d.GetChange("ttl")
			prevComment, _ := d.GetChange("comment")
			prevEa, _ := d.GetChange("ext_attrs")

			// TODO: move to the new Terraform plugin framework and
			// process all the errors instead of ignoring them here.
			_ = d.Set("network_view", prevNetView.(string))
			_ = d.Set("dns_view", prevDNSView.(string))
			_ = d.Set("fqdn", prevFQDN.(string))
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

	if d.HasChange("dns_view") {
		return fmt.Errorf("changing the value of 'dns_view' field is not allowed")
	}

	fqdn := d.Get("fqdn").(string)
	cidr := d.Get("cidr").(string)
	ipAddr := d.Get("ip_addr").(string)

	// for readability
	dynamicAllocation := cidr != ""
	cidrChanged := d.HasChange("cidr")

	// If 'cidr' is not empty (dynamic allocation) and is unchanged,
	// then making it empty to skip the update.
	// (This is to prevent record renewal for the case when 'cidr' is
	// used for IP address allocation, otherwise the address will be changing
	// during every 'update' operation).
	// And making ipAddr empty in case 'cidr' gets changed, to make it possible
	// to allocate an IP address from another network.

	ipaddrChanged := d.HasChange("ip_addr")

	if dynamicAllocation {
		if !cidrChanged {
			cidr = ""
		} else if ipaddrChanged && cidrChanged {
			return fmt.Errorf("only one of 'ip_addr' and 'cidr' values is allowed to update")
		} else {
			ipAddr = ""
		}
	}

	var (
		arec  *ibclient.RecordA
		aList []ibclient.RecordA
	)

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

	oldExtAttrsJSON, newExtAttrsJSON := d.GetChange("ext_attrs")

	newExtAttrs, err := terraformDeserializeEAs(newExtAttrsJSON.(string))
	if err != nil {
		return err
	}

	oldExtAttrs, err := terraformDeserializeEAs(oldExtAttrsJSON.(string))
	if err != nil {
		return err
	}

	var tenantID string
	if tempVal, found := newExtAttrs[eaNameForTenantId]; found {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	//arec, err := objMgr.GetARecordByRef(d.Id())
	obj, err := getOrFindRecord("A", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); ok {
			d.SetId("")
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find apropriate object on NIOS side for resource with ID '%s': %s;"+
					" removing the resource from Terraform state",
				d.Id(), err))
		}

		return err
	}
	recJson, _ := json.Marshal(obj)
	json.Unmarshal(recJson, &aList)
	byteObj, _ := json.Marshal(aList[0])
	err = json.Unmarshal(byteObj, &arec)
	if err != nil {
		return fmt.Errorf("unable to unmarshal json %v", err.Error())
	}
	if d.HasChange("internal_id") {
		return fmt.Errorf("changing the value of 'internal_id' field is not allowed")
	}
	internalId := newInternalResourceIdFromString(d.Get("internal_id").(string))
	newExtAttrs[eaNameForInternalId] = internalId.String()

	//if err != nil {
	//	return fmt.Errorf("failed to read A Record for update operation: %w", err)
	//}

	newExtAttrs, err = mergeEAs(arec.Ea, newExtAttrs, oldExtAttrs, connector)
	if err != nil {
		return err
	}

	rec, err := objMgr.UpdateARecord(
		d.Id(),
		fqdn,
		ipAddr,
		cidr,
		networkView,
		ttl,
		useTtl,
		comment,
		newExtAttrs)
	if err != nil {
		return fmt.Errorf("error updating A-record: %w", err)
	}
	updateSuccessful = true
	d.SetId(rec.Ref)
	if err = d.Set("ref", rec.Ref); err != nil {
		return err
	}

	if err = d.Set("ip_addr", rec.Ipv4Addr); err != nil {
		return err
	}

	return nil
}

func resourceARecordDelete(d *schema.ResourceData, m interface{}) error {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}

	var tenantID string
	tempVal, found := extAttrs[eaNameForTenantId]
	if found {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	_, err = objMgr.DeleteARecord(d.Id())
	if err != nil {
		return fmt.Errorf("deletion of A-record failed: %w", err)
	}
	d.SetId("")

	return nil
}

func resourceARecordImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	var ttl int
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return nil, err
	}

	var tenantID string
	tempVal, found := extAttrs[eaNameForTenantId]
	if found {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	obj, err := objMgr.GetARecordByRef(d.Id())
	if err != nil {
		return nil, fmt.Errorf("failed getting A-record: %w", err)
	}

	if err = d.Set("ip_addr", obj.Ipv4Addr); err != nil {
		return nil, err
	}

	if obj.Ttl != nil {
		ttl = int(*obj.Ttl)
	}
	if !*obj.UseTtl {
		ttl = ttlUndef
	}

	// Internal ID
	// Generate internal ID and add it to the extensible attributes
	internalId := generateInternalId()
	extAttrs[eaNameForInternalId] = internalId.String()
	if err = d.Set("internal_id", internalId.String()); err != nil {
		return nil, err
	}
	// Set ref
	if err = d.Set("ref", obj.Ref); err != nil {
		return nil, err
	}

	if err = d.Set("ttl", ttl); err != nil {
		return nil, err
	}
	delete(obj.Ea, eaNameForInternalId)
	if obj.Ea != nil && len(obj.Ea) > 0 {
		eaJSON, err := terraformSerializeEAs(obj.Ea)
		if err != nil {
			return nil, err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return nil, err
		}
	}

	if err = d.Set("comment", obj.Comment); err != nil {
		return nil, err
	}

	if err = d.Set("dns_view", obj.View); err != nil {
		return nil, err
	}
	if val, ok := d.GetOk("network_view"); !ok || val.(string) == "" {
		dnsView, err := objMgr.GetDNSView(obj.View)
		if err != nil {
			return nil, fmt.Errorf(
				"error while retrieving information about DNS view '%s': %w",
				obj.View, err)
		}
		if err = d.Set("network_view", dnsView.NetworkView); err != nil {
			return nil, err
		}
	}

	if err = d.Set("fqdn", obj.Name); err != nil {
		return nil, err
	}

	d.SetId(obj.Ref)

	// Resource ARecord update Terraform Internal ID and Ref on NOIS side
	// After the record is imported, call the update function
	//err = resourceARecordUpdate(d, m)
	//if err != nil {
	//	return nil, err
	//}

	_, err = connector.UpdateObject(&ibclient.RecordA{
		Ea: ibclient.EA{
			eaNameForInternalId: internalId.String(),
		},
	}, obj.Ref)
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
