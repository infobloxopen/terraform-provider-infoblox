package infoblox

import (
	"context"
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
				Description: "FQDN for the A-record.",
			},
			"ip_addr": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true, // making this optional because of possible dynamic IP allocation (CIDR)
				Description: "IP address to associate with the A-record. For static allocation, set the field with a valid IP address. For dynamic allocation, leave this field empty and set 'cidr' and 'network_view' fields" +
					"or 'filter_params' and optional 'network_view' fields.",
			},
			"network_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Network view to use when allocating an IP address from a network dynamically. For static allocation, leave this field empty.",
			},
			"filter_params": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The parent network block's extensible attributes (dynamic allocation). For static allocation, leave this field empty.",
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

func resourceARecordCreate(d *schema.ResourceData, m interface{}) error {
	// Check if internal_id is set manually
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
	nextAvailableFilter := d.Get("filter_params").(string)
	if ipAddr == "" && cidr == "" && nextAvailableFilter == "" {
		return fmt.Errorf("either of 'ip_addr' or 'cidr' or 'filter_params' values is required")
	}

	if ipAddr != "" && cidr != "" && nextAvailableFilter == "" {
		return fmt.Errorf("only one of 'ip_addr' or 'cidr' or 'filter_params' values is allowed to be defined")
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

	var newRecord *ibclient.RecordA
	if cidr == "" && ipAddr == "" && nextAvailableFilter != "" {
		var (
			eaMap map[string]string
		)
		err = json.Unmarshal([]byte(nextAvailableFilter), &eaMap)
		eaMap["network_view"] = networkView
		if err != nil {
			return fmt.Errorf("error unmarshalling extra attributes of network container: %s", err)
		}
		rec, err := objMgr.AllocateNextAvailableIp(fqdn, "record:a", eaMap, nil, false, extAttrs, comment, false, nil, "IPV4",
			false, false, "", "", networkView, dnsViewName, useTtl, ttl, nil)
		if err != nil {
			return fmt.Errorf("error allocating next available IP: %w", err)
		}
		var ok bool
		newRecord, ok = rec.(*ibclient.RecordA)
		if !ok {
			return fmt.Errorf("failed to convert rec to *ibclient.RecordA")
		}
	} else {
		newRecord, err = objMgr.CreateARecord(
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
	}
	d.SetId(newRecord.Ref)
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

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	// Get by Ref, if not found, then by Terraform Internal ID from EA
	rec, err := searchObjectByRefOrInternalId("A", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); ok {
			d.SetId("")
			return nil
		} else {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		}
	}

	// Assertion of object type and error handling
	var recA *ibclient.RecordA
	recJson, _ := json.Marshal(rec)
	err = json.Unmarshal(recJson, &recA)

	if err != nil {
		return fmt.Errorf("failed getting A-record: %w", err)
	}

	if err = d.Set("ip_addr", recA.Ipv4Addr); err != nil {
		return err
	}

	if recA.Ttl != nil {
		ttl = int(*recA.Ttl)
	}
	if !*recA.UseTtl {
		ttl = ttlUndef
	}
	if err = d.Set("ttl", ttl); err != nil {
		return err
	}
	delete(recA.Ea, eaNameForInternalId)
	omittedEAs := omitEAs(recA.Ea, extAttrs)

	if omittedEAs != nil && len(omittedEAs) > 0 {
		eaJSON, err := terraformSerializeEAs(omittedEAs)
		if err != nil {
			return err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return err
		}
	}

	if err = d.Set("comment", recA.Comment); err != nil {
		return err
	}

	if err = d.Set("dns_view", recA.View); err != nil {
		return err
	}
	if err = d.Set("ref", recA.Ref); err != nil {
		return err
	}
	if val, ok := d.GetOk("network_view"); !ok || val.(string) == "" {
		dnsView, err := objMgr.GetDNSView(recA.View)
		if err != nil {
			return fmt.Errorf(
				"error while retrieving information about DNS view '%s': %w",
				recA.View, err)
		}
		if err = d.Set("network_view", dnsView.NetworkView); err != nil {
			return err
		}
	}

	if err = d.Set("fqdn", recA.Name); err != nil {
		return err
	}

	d.SetId(recA.Ref)

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
			prevNextAvailableFilter, _ := d.GetChange("filter_params")

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
			_ = d.Set("filter_params", prevNextAvailableFilter.(string))
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

	// Get by Ref
	recA, err := objMgr.GetARecordByRef(d.Id())
	if err != nil {
		return fmt.Errorf("failed to read A Record for update operation: %w", err)
	}

	// If 'internal_id' is not set, then generate a new one and set it to the EA.
	internalId := d.Get("internal_id").(string)
	if internalId == "" {
		internalId = generateInternalId().String()
	}
	newInternalId := newInternalResourceIdFromString(internalId)
	newExtAttrs[eaNameForInternalId] = newInternalId.String()

	newExtAttrs, err = mergeEAs(recA.Ea, newExtAttrs, oldExtAttrs, connector)
	if err != nil {
		return err
	}

	obj, err := objMgr.UpdateARecord(
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
	d.SetId(obj.Ref)
	if err = d.Set("ref", obj.Ref); err != nil {
		return err
	}
	if err = d.Set("internal_id", newInternalId.String()); err != nil {
		return err
	}

	if err = d.Set("ip_addr", obj.Ipv4Addr); err != nil {
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

	rec, err := searchObjectByRefOrInternalId("A", d, m)
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
	var recA *ibclient.RecordA
	recJson, _ := json.Marshal(rec)
	err = json.Unmarshal(recJson, &recA)

	_, err = objMgr.DeleteARecord(recA.Ref)
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

	// Resource ARecord update Terraform Internal ID and Ref on NIOS side
	// After the record is imported, call the update function
	err = resourceARecordUpdate(d, m)
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
