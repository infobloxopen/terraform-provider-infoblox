package infoblox

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	// "log"
	// "reflect"
)

func resourceZoneForward() *schema.Resource {
	return &schema.Resource{
		Create: resourceZoneForwardCreate,
		Read:   resourceZoneForwardRead,
		Update: resourceZoneForwardUpdate,
		Delete: resourceZoneForwardDelete,
		Importer: &schema.ResourceImporter{
			State: resourceZoneForwardImport,
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
			"fqdn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of this DNS zone",
			},
			"forward_to": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "The information for the remote name servers to which you want the Infoblox appliance to forward queries for a specified domain name.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The IP address of the remote name server to which you want the Infoblox appliance to forward queries for a specified domain name.",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the remote name server to which you want the Infoblox appliance to forward queries for a specified domain name.",
						},
					},
				},
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
			"ns_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A forwarding member name server group.",
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
			"ext_attrs": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Extensible attributes of the zone forward to be added/updated, as a map in JSON format.",
			},
			"forwarders_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Determines if the appliance sends queries to forwarders only, and not to other internal or Internet root servers.",
			},
			"forwarding_servers": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The information for the remote name servers to which you want the Infoblox appliance to forward queries for a specified domain name.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of this Grid member in FQDN format.",
						},
						"forwarders_only": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Determines if the appliance sends queries to forwarders only, and not to other internal or Internet root servers.",
						},
						"use_override_forwarders": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Determines if the appliance sends queries to name servers.",
						},
						"forward_to": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The information for the remote name servers to which you want the Infoblox appliance to forward queries for a specified domain name.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The IP address of the remote name server to which you want the Infoblox appliance to forward queries for a specified domain name.",
									},
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The name of the remote name server to which you want the Infoblox appliance to forward queries for a specified domain name.",
									},
								},
							},
						},
					},
				},
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

func resourceZoneForwardCreate(d *schema.ResourceData, m interface{}) error {
	// Check if internal_id is set manually
	if intId := d.Get("internal_id"); intId.(string) != "" {
		return fmt.Errorf("the value of 'internal_id' field must not be set manually")
	}

	_, nsGroupOk := d.GetOk("ns_group")
	fsInterface, forwardingServersOk := d.GetOk("forwarding_servers")

	if nsGroupOk && forwardingServersOk {
		return fmt.Errorf("ns_group and forwarding_servers are mutually exclusive")
	}

	fqdn := d.Get("fqdn").(string)
	nsGroup := d.Get("ns_group").(string)
	view := d.Get("view").(string)
	zoneFormat := d.Get("zone_format").(string)

	comment := d.Get("comment").(string)
	disable := d.Get("disable").(bool)
	forwardersOnly := d.Get("forwarders_only").(bool)
	ftInterface := d.Get("forward_to")

	ftSlice, ok := ftInterface.([]interface{})
	if !ok {
		return fmt.Errorf("forward_to is not a slice of Nameservers")
	}
	forwardTo, err := validateForwardTo(ftSlice)
	if err != nil {
		return err
	}

	fsSlice, ok := fsInterface.([]interface{})
	if !ok {
		return fmt.Errorf("forwarding_servers is not a slice of Forwardingmemberserver pointers")
	}
	forwardingServer, err := validateForwardingServers(fsSlice)
	if err != nil {
		return err
	}

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

	newForwardZone, err := objMgr.CreateZoneForward(comment, disable, extAttrs, forwardTo, forwardersOnly, forwardingServer, fqdn, nsGroup, view, zoneFormat)
	if err != nil {
		return fmt.Errorf("failed to create zone forward : %s", err)
	}
	d.SetId(newForwardZone.Ref)
	if err = d.Set("internal_id", internalId.String()); err != nil {
		return err
	}
	if err = d.Set("ref", newForwardZone.Ref); err != nil {
		return err
	}
	return nil
}

func validateForwardingServers(fsSlice []interface{}) ([]*ibclient.Forwardingmemberserver, error) {
	fsStr, err := json.Marshal(fsSlice)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal forwarding_servers: %s", err)
	}
	var forwardingServer []*ibclient.Forwardingmemberserver
	err = json.Unmarshal(fsStr, &forwardingServer)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal forwarding_servers: %s", err)
	}
	return forwardingServer, nil
}

func validateForwardTo(ftSlice []interface{}) ([]ibclient.NameServer, error) {
	nsStr, err := json.Marshal(ftSlice)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal forward_to: %s", err)
	}
	var forwardTo []ibclient.NameServer
	err = json.Unmarshal(nsStr, &forwardTo)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal forward_to: %s", err)
	}
	return forwardTo, nil
}

func resourceZoneForwardRead(d *schema.ResourceData, m interface{}) error {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}

	rec, err := searchObjectByRefOrInternalId("ZoneForward", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		} else {
			d.SetId("")
			return nil
		}
	}
	var zoneForward *ibclient.ZoneForward
	recJson, err := json.Marshal(rec)
	if err != nil {
		return fmt.Errorf("failed to marshal zone forward record : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &zoneForward)
	if err != nil {
		return fmt.Errorf("failed getting zone forward record : %s", err.Error())
	}

	delete(zoneForward.Ea, eaNameForInternalId)
	omittedEAs := omitEAs(zoneForward.Ea, extAttrs)

	if omittedEAs != nil && len(omittedEAs) > 0 {
		eaJSON, err := terraformSerializeEAs(omittedEAs)
		if err != nil {
			return err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return err
		}
	}

	if err := d.Set("fqdn", zoneForward.Fqdn); err != nil {
		return err
	}

	if zoneForward.View != nil {
		if err := d.Set("view", *zoneForward.View); err != nil {
			return err
		}
	}

	if err := d.Set("zone_format", zoneForward.ZoneFormat); err != nil {
		return err
	}

	if zoneForward.NsGroup != nil {
		if err := d.Set("ns_group", *zoneForward.NsGroup); err != nil {
			return err
		}
	} else {
		if err := d.Set("ns_group", ""); err != nil {
			return err
		}
	}

	if zoneForward.Comment != nil {
		if err := d.Set("comment", *zoneForward.Comment); err != nil {
			return err
		}
	}

	if zoneForward.Disable != nil {
		if err := d.Set("disable", *zoneForward.Disable); err != nil {
			return err
		}
	}

	if zoneForward.ForwardersOnly != nil {
		if err := d.Set("forwarders_only", *zoneForward.ForwardersOnly); err != nil {
			return err
		}
	}

	if zoneForward.ForwardTo != nil {
		nsInterface := convertForwardToInterface(zoneForward.ForwardTo)
		if err := d.Set("forward_to", nsInterface); err != nil {
			return err
		}
	}

	if zoneForward.ForwardingServers.Servers != nil {
		fwServerInterface, _ := convertForwardingServersToInterface(zoneForward.ForwardingServers.Servers)
		if err := d.Set("forwarding_servers", fwServerInterface); err != nil {
			return err
		}
	} else {
		if err := d.Set("forwarding_servers", nil); err != nil {
			return err
		}
	}

	d.SetId(zoneForward.Ref)
	return nil
}

func resourceZoneForwardUpdate(d *schema.ResourceData, m interface{}) error {
	var updateSuccessful bool
	defer func() {
		// Reverting the state back, in case of a failure,
		// otherwise Terraform will keep the values, which leaded to the failure, in the state file.
		if !updateSuccessful {
			prevComment, _ := d.GetChange("comment")
			prevDisable, _ := d.GetChange("disable")
			prevForwardersOnly, _ := d.GetChange("forwarders_only")
			prevNsGroup, _ := d.GetChange("ns_group")
			prevForwardTo, _ := d.GetChange("forward_to")
			prevForwardingServers, _ := d.GetChange("forwarding_servers")
			prevExtAttrs, _ := d.GetChange("ext_attrs")

			_ = d.Set("comment", prevComment.(string))
			_ = d.Set("disable", prevDisable.(bool))
			_ = d.Set("forwarders_only", prevForwardersOnly.(bool))
			_ = d.Set("ns_group", prevNsGroup.(string))
			_ = d.Set("forward_to", prevForwardTo)
			_ = d.Set("forwarding_servers", prevForwardingServers)
			_ = d.Set("ext_attrs", prevExtAttrs.(string))
		}
	}()

	_, nsGroupOk := d.GetOk("ns_group")
	fsInterface, forwardingServersOk := d.GetOk("forwarding_servers")

	if nsGroupOk && forwardingServersOk {
		return fmt.Errorf("ns_group and forwarding_servers are mutually exclusive")
	}

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

	var zf *ibclient.ZoneForward

	rec, err := searchObjectByRefOrInternalId("ZoneForward", d, m)
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
		return fmt.Errorf("failed to marshal zone forward record : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &zf)
	if err != nil {
		return fmt.Errorf("failed getting zone forward record : %s", err.Error())
	}

	// If 'internal_id' is not set, then generate a new one and set it to the EA.
	internalId := d.Get("internal_id").(string)
	if internalId == "" {
		internalId = generateInternalId().String()
	}
	newInternalId := newInternalResourceIdFromString(internalId)
	newExtAttrs[eaNameForInternalId] = newInternalId.String()

	newExtAttrs, err = mergeEAs(zf.Ea, newExtAttrs, oldExtAttrs, connector)
	if err != nil {
		return err
	}

	comment := d.Get("comment").(string)
	disable := d.Get("disable").(bool)
	var nsGroup string
	if d.Get("ns_group") != "" {
		nsGroup = d.Get("ns_group").(string)
	} else {
		nsGroup = ""
	}
	ftInterface := d.Get("forward_to")
	forwardersOnly := d.Get("forwarders_only").(bool)
	//fsInterface := d.Get("forwarding_servers")

	ftSlice, ok := ftInterface.([]interface{})
	if !ok {
		return fmt.Errorf("forward_to is not a slice of inetrfaces")
	}
	forwardTo, err := validateForwardTo(ftSlice)
	if err != nil {
		return err
	}

	var nullFWS *ibclient.NullableForwardingServers

	if forwardingServersOk && d.HasChange("forwarding_servers") {
		fsSlice, ok := fsInterface.([]interface{})
		if !ok {
			return fmt.Errorf("forwarding_servers is not a slice of Forwardingmemberserver pointers")
		}
		forwardingServer, err := validateForwardingServers(fsSlice)
		if err != nil {
			return err
		}
		nullFWS = &ibclient.NullableForwardingServers{IsNull: false, Servers: forwardingServer}
	} else if !forwardingServersOk && d.HasChange("forwarding_servers") {
		nullFWS = &ibclient.NullableForwardingServers{IsNull: false, Servers: []*ibclient.Forwardingmemberserver{}}
	} else {
		nullFWS = &ibclient.NullableForwardingServers{IsNull: true, Servers: nil}
	}

	//TODO: Check if forwarding_servers is nil

	zf, err = objMgr.UpdateZoneForward(d.Id(), comment, disable, newExtAttrs, forwardTo, forwardersOnly, nullFWS, nsGroup)
	if err != nil {
		return fmt.Errorf("Failed to update Zone Forward with %s, ", err.Error())
	}

	updateSuccessful = true

	if err = d.Set("internal_id", newInternalId.String()); err != nil {
		return err
	}
	if err = d.Set("ref", zf.Ref); err != nil {
		return err
	}
	d.SetId(zf.Ref)

	return nil
}

func resourceZoneForwardDelete(d *schema.ResourceData, m interface{}) error {
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

	rec, err := searchObjectByRefOrInternalId("ZoneForward", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		} else {
			d.SetId("")
			return nil
		}
	}

	var zf *ibclient.ZoneForward
	recJson, err := json.Marshal(rec)
	if err != nil {
		return fmt.Errorf("failed to marshal zone forward record : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &zf)
	if err != nil {
		return err
	}
	_, err = objMgr.DeleteZoneForward(zf.Ref)
	if err != nil {
		return fmt.Errorf("Failed to delete Zone Forward : %s", err.Error())
	}

	return nil
}

func resourceZoneForwardImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	extAttrJSON := d.Get("ext_attrs").(string)
	_, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return nil, err
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	zf, err := objMgr.GetZoneForwardByRef(d.Id())
	if err != nil {
		return nil, fmt.Errorf("failed getting zone forward record: %w", err)
	}

	if zf.Ea != nil && len(zf.Ea) > 0 {
		eaJSON, err := terraformSerializeEAs(zf.Ea)
		if err != nil {
			return nil, err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return nil, err
		}
	}

	if err = d.Set("fqdn", zf.Fqdn); err != nil {
		return nil, err
	}

	if zf.View != nil {
		if err = d.Set("view", *zf.View); err != nil {
			return nil, err
		}
	}

	if zf.NsGroup != nil {
		if err = d.Set("ns_group", *zf.NsGroup); err != nil {
			return nil, err
		}
	}

	if err = d.Set("zone_format", zf.ZoneFormat); err != nil {
		return nil, err
	}

	if zf.Comment != nil {
		if err = d.Set("comment", *zf.Comment); err != nil {
			return nil, err
		}
	}

	if zf.Disable != nil {
		if err = d.Set("disable", *zf.Disable); err != nil {
			return nil, err
		}
	}

	if zf.ForwardTo != nil {
		nsInterface := convertForwardToInterface(zf.ForwardTo)
		if err = d.Set("forward_to", nsInterface); err != nil {
			return nil, err
		}
	}

	if zf.ForwardersOnly != nil {
		if err = d.Set("forwarders_only", *zf.ForwardersOnly); err != nil {
			return nil, err
		}
	}

	if zf.ForwardingServers.Servers != nil {
		fwServerInterface, _ := convertForwardingServersToInterface(zf.ForwardingServers.Servers)
		if err = d.Set("forwarding_servers", fwServerInterface); err != nil {
			return nil, err
		}
	} else {
		if err := d.Set("forwarding_servers", nil); err != nil {
			return nil, err
		}
	}

	d.SetId(zf.Ref)

	// Update the resource with the EA Terraform Internal ID
	err = resourceZoneForwardUpdate(d, m)
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
