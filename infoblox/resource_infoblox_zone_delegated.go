package infoblox

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func resourceZoneDelegated() *schema.Resource {
	return &schema.Resource{
		Create: resourceZoneDelegatedCreate,
		Read:   resourceZoneDelegatedRead,
		Update: resourceZoneDelegatedUpdate,
		Delete: resourceZoneDelegatedDelete,
		Importer: &schema.ResourceImporter{
			State: resourceZoneDelegatedImport,
		},
		Schema: map[string]*schema.Schema{
			"fqdn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The FQDN of the delegated zone.",
			},
			"delegate_to": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The Infoblox appliance redirects queries for data for the delegated zone to this remote name server.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The IPv4 Address or IPv6 Address of the server.",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "A resolvable domain name for the external DNS server.",
						},
					},
				},
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A descriptive comment.",
			},
			"disable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Determines if the zone is disabled or not.",
			},
			"locked": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				Description: "If you enable this flag, other administrators cannot make conflicting changes. This is for administration purposes only. " +
					"The zone will continue to serve DNS data even when it is locked.",
			},
			"ns_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The delegation NS group bound with delegated zone.",
			},
			"delegated_ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     ttlUndef,
				Description: "TTL value for zone-delegated.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "Extensible attributes, as a map in JSON format",
			},
			"view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
				Description: "The DNS view in which the zone is created.",
			},
			"zone_format": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "FORWARD",
				Description: "The format of the zone. Valid values are: FORWARD, IPV4, IPV6.",
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

func resourceZoneDelegatedCreate(d *schema.ResourceData, m interface{}) error {

	if intId := d.Get("internal_id"); intId.(string) != "" {
		return fmt.Errorf("the value of 'internal_id' field must not be set manually")
	}
	fqdn := d.Get("fqdn").(string)

	nsGroup, nsGroupOk := d.GetOk("ns_group")
	dtInterface, delegateToOk := d.GetOk("delegate_to")

	var delegateTo []ibclient.NameServer
	var nullDT ibclient.NullableNameServers
	if !nsGroupOk && !delegateToOk {
		return fmt.Errorf("either 'ns_group' or 'delegate_to' must be set")
	}
	if delegateToOk {
		dtSlice := dtInterface.(*schema.Set).List()
		var err error
		delegateTo, err = validateForwardTo(dtSlice)
		if err != nil {
			return err
		}
		nullDT = ibclient.NullableNameServers{IsNull: false, NameServers: delegateTo}
	}

	comment := d.Get("comment").(string)
	disable := d.Get("disable").(bool)
	locked := d.Get("locked").(bool)
	delegatedTtl := d.Get("delegated_ttl")

	var ttl uint32
	useTtl := false
	tempTTL := delegatedTtl.(int)
	if tempTTL >= 0 {
		useTtl = true
		ttl = uint32(tempTTL)
	} else if tempTTL != ttlUndef {
		return fmt.Errorf("TTL value must be 0 or higher")
	}

	view := d.Get("view").(string)
	zoneFormat := d.Get("zone_format").(string)

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

	newZoneDelegated, err := objMgr.CreateZoneDelegated(fqdn, nullDT, comment, disable, locked, nsGroup.(string), ttl, useTtl, extAttrs, view, zoneFormat)
	if err != nil {
		return fmt.Errorf("failed to create zone delegation : %s", err)
	}
	d.SetId(newZoneDelegated.Ref)
	if err = d.Set("internal_id", internalId.String()); err != nil {
		return err
	}
	if err = d.Set("ref", newZoneDelegated.Ref); err != nil {
		return err
	}
	return resourceZoneDelegatedRead(d, m)
}

func resourceZoneDelegatedRead(d *schema.ResourceData, m interface{}) error {

	var ttl int
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}

	rec, err := searchObjectByRefOrInternalId("ZoneDelegated", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		} else {
			d.SetId("")
			return nil
		}
	}

	var zoneDelegated *ibclient.ZoneDelegated
	recJson, err := json.Marshal(rec)
	if err != nil {
		return fmt.Errorf("failed to marshal zone delegated record : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &zoneDelegated)
	if err != nil {
		return fmt.Errorf("failed getting zone delegated record : %s", err.Error())
	}

	delete(zoneDelegated.Ea, eaNameForInternalId)
	omittedEAs := omitEAs(zoneDelegated.Ea, extAttrs)

	if omittedEAs != nil && len(omittedEAs) > 0 {
		eaJSON, err := terraformSerializeEAs(omittedEAs)
		if err != nil {
			return err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return err
		}
	}

	if zoneDelegated.DelegatedTtl != nil {
		ttl = int(*zoneDelegated.DelegatedTtl)
	}
	if !*zoneDelegated.UseDelegatedTtl {
		ttl = ttlUndef
	}
	if err = d.Set("delegated_ttl", ttl); err != nil {
		return err
	}

	if err := d.Set("fqdn", zoneDelegated.Fqdn); err != nil {
		return err
	}

	if zoneDelegated.View != nil {
		if err := d.Set("view", *zoneDelegated.View); err != nil {
			return err
		}
	}

	if err := d.Set("zone_format", zoneDelegated.ZoneFormat); err != nil {
		return err
	}

	if zoneDelegated.NsGroup != nil {
		if err := d.Set("ns_group", *zoneDelegated.NsGroup); err != nil {
			return err
		}
	} else {
		if err := d.Set("ns_group", ""); err != nil {
			return err
		}
	}

	if zoneDelegated.Comment != nil {
		if err := d.Set("comment", *zoneDelegated.Comment); err != nil {
			return err
		}
	}

	if zoneDelegated.Disable != nil {
		if err := d.Set("disable", *zoneDelegated.Disable); err != nil {
			return err
		}
	}

	if zoneDelegated.DelegateTo.NameServers != nil {
		nsInterface := convertNullableNameServersToInterface(zoneDelegated.DelegateTo)
		if err = d.Set("delegate_to", nsInterface); err != nil {
			return err
		}
	} else {
		if err := d.Set("delegate_to", nil); err != nil {
			return err
		}
	}

	if zoneDelegated.Locked != nil {
		if err := d.Set("locked", *zoneDelegated.Locked); err != nil {
			return err
		}
	}

	d.SetId(zoneDelegated.Ref)
	return nil
}

func resourceZoneDelegatedUpdate(d *schema.ResourceData, m interface{}) error {
	var updateSuccessful bool
	defer func() {
		// Reverting the state back, in case of a failure,
		// otherwise Terraform will keep the values, which leaded to the failure, in the state file.
		if !updateSuccessful {
			prevComment, _ := d.GetChange("comment")
			prevDisable, _ := d.GetChange("disable")
			prevLocked, _ := d.GetChange("locked")
			prevNsGroup, _ := d.GetChange("ns_group")
			prevDelegateTo, _ := d.GetChange("delegate_to")
			prevExtAttrs, _ := d.GetChange("ext_attrs")
			prevTtl, _ := d.GetChange("delegated_ttl")

			_ = d.Set("comment", prevComment.(string))
			_ = d.Set("disable", prevDisable.(bool))
			_ = d.Set("locked", prevLocked.(bool))
			_ = d.Set("ns_group", prevNsGroup.(string))
			_ = d.Set("delegate_to", prevDelegateTo)
			_ = d.Set("ext_attrs", prevExtAttrs.(string))
			_ = d.Set("delegated_ttl", prevTtl.(int))
		}
	}()

	_, nsGroupOk := d.GetOk("ns_group")
	dtInterface, delegateToOk := d.GetOk("delegate_to")

	if d.HasChange("internal_id") {
		return fmt.Errorf("changing the value of 'internal_id' field is not allowed")
	}
	if d.HasChange("fqdn") {
		return fmt.Errorf("changing the value of 'fqdn' field is not allowed")
	}
	if d.HasChange("view") {
		return fmt.Errorf("changing the value of 'view' field is not allowed")
	}
	if d.HasChange("zone_format") {
		return fmt.Errorf("changing the value of 'zone_format' field is not allowed")
	}

	var delegateTo []ibclient.NameServer
	var nullDT ibclient.NullableNameServers
	if !nsGroupOk && !delegateToOk {
		return fmt.Errorf("either ns_group or delegate_to must be set")
	} else if !delegateToOk {
		nullDT = ibclient.NullableNameServers{IsNull: false, NameServers: []ibclient.NameServer{}}
	} else {
		dtSlice := dtInterface.(*schema.Set).List()
		var err error
		delegateTo, err = validateForwardTo(dtSlice)
		if err != nil {
			return err
		}
		nullDT = ibclient.NullableNameServers{IsNull: false, NameServers: delegateTo}
	}

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

	var zoneDelegated *ibclient.ZoneDelegated

	rec, err := searchObjectByRefOrInternalId("ZoneDelegated", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		} else {
			d.SetId("")
			return nil
		}
	}
	recJson, err := json.Marshal(rec)
	if err != nil {
		return fmt.Errorf("failed to marshal zone delegated record : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &zoneDelegated)
	if err != nil {
		return fmt.Errorf("failed getting zone delegated record : %s", err.Error())
	}

	// If 'internal_id' is not set, then generate a new one and set it to the EA.
	internalId := d.Get("internal_id").(string)
	if internalId == "" {
		internalId = generateInternalId().String()
	}
	newInternalId := newInternalResourceIdFromString(internalId)
	newExtAttrs[eaNameForInternalId] = newInternalId.String()

	newExtAttrs, err = mergeEAs(zoneDelegated.Ea, newExtAttrs, oldExtAttrs, connector)
	if err != nil {
		return err
	}

	comment := d.Get("comment").(string)
	disable := d.Get("disable").(bool)
	locked := d.Get("locked").(bool)
	delegatedTtl := d.Get("delegated_ttl")
	var nsGroup string
	if d.Get("ns_group") != "" {
		nsGroup = d.Get("ns_group").(string)
	} else {
		nsGroup = ""
	}
	var ttl uint32
	useTtl := false
	tempTTL := delegatedTtl.(int)
	if tempTTL >= 0 {
		useTtl = true
		ttl = uint32(tempTTL)
	} else if tempTTL != ttlUndef {
		return fmt.Errorf("TTL value must be 0 or higher")
	}

	zoneDelegated, err = objMgr.UpdateZoneDelegated(d.Id(), nullDT, comment, disable, locked, nsGroup, ttl, useTtl, newExtAttrs)
	if err != nil {
		return fmt.Errorf("Failed to update zone delegated with %s, ", err.Error())
	}

	updateSuccessful = true

	if err = d.Set("internal_id", newInternalId.String()); err != nil {
		return err
	}
	if err = d.Set("ref", zoneDelegated.Ref); err != nil {
		return err
	}
	d.SetId(zoneDelegated.Ref)

	return nil
}

func resourceZoneDelegatedDelete(d *schema.ResourceData, m interface{}) error {
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

	rec, err := searchObjectByRefOrInternalId("ZoneDelegated", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		} else {
			d.SetId("")
			return nil
		}
	}

	var zd *ibclient.ZoneDelegated
	recJson, err := json.Marshal(rec)
	if err != nil {
		return fmt.Errorf("failed to marshal zone delegated record : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &zd)
	if err != nil {
		return err
	}
	_, err = objMgr.DeleteZoneDelegated(zd.Ref)
	if err != nil {
		return fmt.Errorf("failed to delete zone delegated : %s", err.Error())
	}

	return nil
}

func resourceZoneDelegatedImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	extAttrJSON := d.Get("ext_attrs").(string)
	_, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return nil, err
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	zoneDelegated, err := objMgr.GetZoneDelegatedByRef(d.Id())
	if err != nil {
		return nil, fmt.Errorf("failed getting zone delegated record: %w", err)
	}

	if zoneDelegated.Ea != nil && len(zoneDelegated.Ea) > 0 {
		eaJSON, err := terraformSerializeEAs(zoneDelegated.Ea)
		if err != nil {
			return nil, err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return nil, err
		}
	}
	delete(zoneDelegated.Ea, eaNameForInternalId)

	if err = d.Set("fqdn", zoneDelegated.Fqdn); err != nil {
		return nil, err
	}

	if zoneDelegated.View != nil {
		if err = d.Set("view", *zoneDelegated.View); err != nil {
			return nil, err
		}
	}

	if zoneDelegated.NsGroup != nil {
		if err = d.Set("ns_group", *zoneDelegated.NsGroup); err != nil {
			return nil, err
		}
	}

	if err = d.Set("zone_format", zoneDelegated.ZoneFormat); err != nil {
		return nil, err
	}

	if zoneDelegated.Comment != nil {
		if err = d.Set("comment", *zoneDelegated.Comment); err != nil {
			return nil, err
		}
	}

	if zoneDelegated.Disable != nil {
		if err = d.Set("disable", *zoneDelegated.Disable); err != nil {
			return nil, err
		}
	}

	if zoneDelegated.Locked != nil {
		if err = d.Set("locked", *zoneDelegated.Locked); err != nil {
			return nil, err
		}
	}

	if zoneDelegated.DelegatedTtl != nil {
		if err = d.Set("delegated_ttl", *zoneDelegated.DelegatedTtl); err != nil {
			return nil, err
		}
	}

	if zoneDelegated.DelegateTo.NameServers != nil {
		nsInterface := convertNullableNameServersToInterface(zoneDelegated.DelegateTo)
		if err = d.Set("delegate_to", nsInterface); err != nil {
			return nil, err
		}
	} else {
		if err := d.Set("delegate_to", nil); err != nil {
			return nil, err
		}
	}

	d.SetId(zoneDelegated.Ref)

	// Update the resource with the EA Terraform Internal ID
	err = resourceZoneDelegatedUpdate(d, m)
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
