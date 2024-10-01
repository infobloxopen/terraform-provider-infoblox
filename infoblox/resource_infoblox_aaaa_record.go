package infoblox

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func resourceAAAARecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceAAAARecordCreate,
		Read:   resourceAAAARecordGet,
		Update: resourceAAAARecordUpdate,
		Delete: resourceAAAARecordDelete,
		Importer: &schema.ResourceImporter{
			State: resourceAAAARecordImport,
		},
		CustomizeDiff: func(context context.Context, d *schema.ResourceDiff, meta interface{}) error {
			if internalID := d.Get("internal_id"); internalID == "" || internalID == nil {
				err := d.SetNewComputed("internal_id")
				if err != nil {
					return err
				}
			}
			return nil
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
				Description: "FQDN for the AAAA-record.",
			},
			"ipv6_addr": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true, // making this optional because of possible dynamic IP allocation (CIDR)
				Description: "IP address to associate with the AAAA-record. For static allocation, set the field with a valid IP address. For dynamic allocation, leave this field empty and set 'cidr' and 'network_view' fields.",
			},
			"filter_params": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The parent network's Ip or extensible attributes.",
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
				Description: "Network to allocate an IP address from, when the 'ipv6_addr' field is empty (dynamic allocation). The address is in CIDR format. For static allocation, leave this field empty.",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     ttlUndef,
				Description: "TTL value for the AAAA-record.",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Description of the AAAA-record.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Extensible attributes of the AAAA-record to be added/updated, as a map in JSON format",
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

func resourceAAAARecordCreate(d *schema.ResourceData, m interface{}) error {

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
	ipv6Addr := d.Get("ipv6_addr").(string)
	nextAvailableFilter := d.Get("filter_params").(string)
	if ipv6Addr == "" && cidr == "" && nextAvailableFilter == "" {
		return fmt.Errorf("any one of 'ipv6_addr', 'cidr' and 'filter_params' values is required")
	}

	if ipv6Addr != "" && cidr != "" && nextAvailableFilter != "" {
		return fmt.Errorf("only one of 'ipv6_addr', 'cidr' and 'filter_params' values is allowed to be defined")
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
	if err != nil {
		return err
	}

	// Generate internal ID and add it to the extensible attributes
	internalId := generateInternalId()
	extAttrs[eaNameForInternalId] = internalId.String()

	var tenantID string
	if tempVal, found := extAttrs[eaNameForTenantId]; found {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	var (
		newRecordAAAA interface{}
		eaMap         map[string]string
	)

	if nextAvailableFilter != "" {
		err = json.Unmarshal([]byte(nextAvailableFilter), &eaMap)
		if err != nil {
			return fmt.Errorf("error unmarshalling extra attributes of network: %s", err)
		}
		newRecordAAAA, err = objMgr.AllocateNextAvailableIp(fqdn, "record:aaaa", eaMap, nil, false, true, extAttrs, comment, false, nil)

	} else {
		newRecordAAAA, err = objMgr.CreateAAAARecord(networkView, dnsViewName, fqdn, cidr, ipv6Addr, useTtl, ttl, comment, extAttrs)
	}
	if err != nil {
		return fmt.Errorf("creation of AAAA-record under DNS view '%s' failed: %w", dnsViewName, err)
	}

	recordAAAA := newRecordAAAA.(*ibclient.RecordAAAA)
	d.SetId(recordAAAA.Ref)

	if err = d.Set("ref", recordAAAA.Ref); err != nil {
		return err
	}
	if err = d.Set("internal_id", internalId.String()); err != nil {
		return err
	}

	if err = d.Set("ipv6_addr", recordAAAA.Ipv6Addr); err != nil {
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

func resourceAAAARecordGet(d *schema.ResourceData, m interface{}) error {
	var ttl int
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}

	var tenantID string
	if tempVal, ok := extAttrs[eaNameForTenantId]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	rec, err := searchObjectByRefOrInternalId("AAAA", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		} else {
			d.SetId("")
			return nil
		}
	}

	// Assertion of object type and error handling
	var obj *ibclient.RecordAAAA
	recJson, _ := json.Marshal(rec)
	err = json.Unmarshal(recJson, &obj)

	if err != nil && obj.Ref != "" {
		return fmt.Errorf("getting AAAA Record with ID: %s failed: %w", d.Id(), err)
	}
	if err = d.Set("ipv6_addr", obj.Ipv6Addr); err != nil {
		return err
	}

	if obj.Ttl != nil {
		ttl = int(*obj.Ttl)
	}
	if !*obj.UseTtl {
		ttl = ttlUndef
	}
	if err = d.Set("ttl", ttl); err != nil {
		return err
	}

	delete(obj.Ea, eaNameForInternalId)
	omittedEAs := omitEAs(obj.Ea, extAttrs)

	if omittedEAs != nil && len(omittedEAs) > 0 {
		eaJSON, err := terraformSerializeEAs(omittedEAs)
		if err != nil {
			return err
		}

		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return err
		}
	}

	if err = d.Set("comment", obj.Comment); err != nil {
		return err
	}

	if err = d.Set("dns_view", obj.View); err != nil {
		return err
	}
	if err = d.Set("ref", obj.Ref); err != nil {
		return err
	}
	if val, ok := d.GetOk("network_view"); !ok || val.(string) == "" {
		dnsView, err := objMgr.GetDNSView(obj.View)
		if err != nil {
			return fmt.Errorf(
				"error while retrieving information about DNS view '%s': %w",
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
			prevNextAvailableFilter, _ := d.GetChange("filter_params")
			prevCIDR, _ := d.GetChange("cidr")
			prevTTL, _ := d.GetChange("ttl")
			prevComment, _ := d.GetChange("comment")
			prevEa, _ := d.GetChange("ext_attrs")

			_ = d.Set("network_view", prevNetView.(string))
			_ = d.Set("dns_view", prevDNSView.(string))
			_ = d.Set("fqdn", prevFQDN.(string))
			_ = d.Set("ipv6_addr", prevIPAddr.(string))
			_ = d.Set("filter_params", prevNextAvailableFilter.(string))
			_ = d.Set("cidr", prevCIDR.(string))
			_ = d.Set("ttl", prevTTL.(int))
			_ = d.Set("comment", prevComment.(string))
			_ = d.Set("ext_attrs", prevEa.(string))

		}
	}()

	if d.HasChange("internal_id") {
		return fmt.Errorf("changing the value of 'internal_id' field is not allowed")
	}

	if d.HasChange("network_view") {
		return fmt.Errorf("changing the value of 'network_view' field is not allowed")
	}

	if d.HasChange("dns_view") {
		return fmt.Errorf("changing the value of 'dns_view' field is not allowed")
	}

	if d.HasChange("filter_params") {
		return fmt.Errorf("changing the value of 'filter_params' field is not allowed")
	}

	networkView := d.Get("network_view").(string)
	fqdn := d.Get("fqdn").(string)
	cidr := d.Get("cidr").(string)
	ipv6Addr := d.Get("ipv6_addr").(string)

	// for readability
	dynamicAllocation := cidr != ""
	cidrChanged := d.HasChange("cidr")

	// If 'cidr' is not empty (dynamic allocation) and is unchanged,
	// then making it empty to skip the update.
	// (This is to prevent record renewal for the case when 'cidr' is
	// used for IP address allocation, otherwise the address will be changing
	// during every 'update' operation).
	// And making ipv6Addr empty in case 'cidr' gets changed, to make it possible
	// to allocate an IP address from another network.

	// to get the change status of ipv6 address
	ipaddrChanged := d.HasChange("ipv6_addr")

	if dynamicAllocation {
		if !cidrChanged {
			cidr = ""
		} else if ipaddrChanged && cidrChanged {
			return fmt.Errorf("only one of 'ipv6_addr' and 'cidr' values is allowed to update")
		} else {
			ipv6Addr = ""
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

	qarec, err := objMgr.GetAAAARecordByRef(d.Id())
	if err != nil {
		return fmt.Errorf("failed to read AAAA Record for update operation: %w", err)
	}

	internalId := d.Get("internal_id").(string)

	if internalId == "" {
		internalId = generateInternalId().String()
	}

	newInternalId := newInternalResourceIdFromString(internalId)
	newExtAttrs[eaNameForInternalId] = newInternalId.String()

	newExtAttrs, err = mergeEAs(qarec.Ea, newExtAttrs, oldExtAttrs, connector)
	if err != nil {
		return err
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
		newExtAttrs)
	if err != nil {
		return fmt.Errorf("error updating AAAA-record: %w", err)
	}
	updateSuccessful = true
	d.SetId(recordAAAA.Ref)
	if err = d.Set("ref", recordAAAA.Ref); err != nil {
		return err
	}
	if err = d.Set("internal_id", newInternalId.String()); err != nil {
		return err
	}

	if err = d.Set("ipv6_addr", recordAAAA.Ipv6Addr); err != nil {
		return err
	}

	return nil
}

func resourceAAAARecordDelete(d *schema.ResourceData, m interface{}) error {
	dnsView := d.Get("dns_view").(string)

	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}

	var tenantID string
	if tempVal, ok := extAttrs[eaNameForTenantId]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	qarec, err := searchObjectByRefOrInternalId("AAAA", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		} else {
			d.SetId("")
			return nil
		}
	}

	// Assertion of object type and error handling
	var obj ibclient.RecordAAAA
	recJson, _ := json.Marshal(qarec)
	err = json.Unmarshal(recJson, &obj)

	if err != nil {
		return fmt.Errorf("getting AAAA Record with ID: %s failed: %w", d.Id(), err)
	}

	_, err = objMgr.DeleteAAAARecord(obj.Ref)
	if err != nil {
		return fmt.Errorf("deletion of AAAA Record from dns view %s failed: %w", dnsView, err)
	}
	d.SetId("")

	return nil
}

func resourceAAAARecordImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	var ttl int
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return nil, err
	}

	var tenantID string
	if tempVal, ok := extAttrs[eaNameForTenantId]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	obj, err := objMgr.GetAAAARecordByRef(d.Id())
	if err != nil {
		return nil, fmt.Errorf("getting AAAA Record with ID: %s failed: %w", d.Id(), err)
	}
	if err = d.Set("ipv6_addr", obj.Ipv6Addr); err != nil {
		return nil, err
	}

	if obj.Ttl != nil {
		ttl = int(*obj.Ttl)
	}

	if !*obj.UseTtl {
		ttl = ttlUndef
	}

	// Set ref
	if err = d.Set("ref", obj.Ref); err != nil {
		return nil, err
	}

	if err = d.Set("ttl", ttl); err != nil {
		return nil, err
	}

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

	err = resourceAAAARecordUpdate(d, m)
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
