package infoblox

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"reflect"
)

func resourceDtcLbdnRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceDtcLbdnCreate,
		Read:   resourceDtcLbdnGet,
		Update: resourceDtcLbdnUpdate,
		Delete: resourceDtcLbdnDelete,
		Importer: &schema.ResourceImporter{
			State: resourceDtcLbdnImport,
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
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The display name of the DTC LBDN.",
			},
			"auth_zones": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of linked auth zones.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"auto_consolidated_monitors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Flag for enabling auto managing DTC Consolidated Monitors on related DTC Pools.",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Description of the DTC LBDN record.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Extensible attributes of the DTC LBDN record to be added/updated, as a map in JSON format.",
			},
			"disable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Determines whether the DTC LBDN is disabled or not. When this is set to False, the fixed address is enabled.",
			},
			"lb_method": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The load balancing method. Used to select pool. Valid values are GLOBAL_AVAILABILITY, RATIO, ROUND_ROBIN, SOURCE_IP_HASH and TOPOLOGY.",
			},
			"patterns": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "LBDN wildcards for pattern match.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"persistence": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      0,
				Description:  "Maximum time, in seconds, for which client specific LBDN responses will be cached. Zero specifies no caching.",
				ValidateFunc: validation.IntBetween(0, 7200),
			},
			"pools": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Pools associated with the LBDN are collections of load-balanced servers",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pool": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The pool to link with.",
						},
						"ratio": {
							Type:         schema.TypeInt,
							Required:     true,
							Description:  "The weight of pool.",
							ValidateFunc: validation.IntBetween(1, 65535),
						},
					},
				},
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
				Description: "The LBDN pattern match priority for “overlapping” DTC LBDN objects. LBDNs are “overlapping” if " +
					"they are simultaneously assigned to a zone and have patterns that can match the same FQDN. The matching LBDN with highest priority (lowest ordinal) will be used.",
				ValidateFunc: validation.IntBetween(1, 3),
			},
			"topology": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The topology rules for TOPOLOGY method.",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     ttlUndef,
				Description: "The Time To Live (TTL) value for the DTC LBDN. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached.",
			},
			"types": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The list of resource record types supported by LBDN. Valid values are A, AAAA, CNAME, NAPTR, SRV. Default value is A and AAAA",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
					oldTypes, newTypes := d.GetChange("types")
					if d.Get("types") == nil && oldValue != newValue {
						return true
					}
					return reflect.DeepEqual(oldTypes, newTypes)
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

func resourceDtcLbdnCreate(d *schema.ResourceData, m interface{}) error {
	// Check if internal_id is set manually
	if intId := d.Get("internal_id"); intId.(string) != "" {
		return fmt.Errorf("the value of 'internal_id' field must not be set manually")
	}

	name := d.Get("name").(string)
	authZones := d.Get("auth_zones").([]interface{})
	authZoneList := make([]string, len(authZones))
	for i, authZone := range authZones {
		authZoneList[i] = authZone.(string)
	}
	autoConsolidatedMonitors := d.Get("auto_consolidated_monitors").(bool)
	comment := d.Get("comment").(string)
	disable := d.Get("disable").(bool)
	lbMethod := d.Get("lb_method").(string)
	poolsLink := d.Get("pools").([]interface{})

	pools, err := validatePoolsLink(poolsLink)
	if err != nil {
		return fmt.Errorf("failed to validate pools: %w", err)
	}

	patterns := d.Get("patterns").([]interface{})
	patternsList := make([]string, len(patterns))
	for i, pattern := range patterns {
		patternsList[i] = pattern.(string)
	}
	tempPersistence := d.Get("persistence").(int)
	persistence := uint32(tempPersistence)

	tempPriority := d.Get("priority").(int)
	priority := uint32(tempPriority)

	topology := d.Get("topology").(string)

	types := d.Get("types").([]interface{})
	typesList := make([]string, len(types))
	for i, j := range types {
		typesList[i] = j.(string)
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

	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return fmt.Errorf("failed to allocate IP: %w", err)
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

	// Create the DTC LBDN record
	newRecord, err := objMgr.CreateDtcLbdn(name, authZoneList, comment, disable, autoConsolidatedMonitors, extAttrs, lbMethod, patternsList, persistence, pools, priority, topology, typesList, ttl, useTtl)
	if err != nil {
		return fmt.Errorf("failed to create DTC LBDN record: %w", err)
	}
	d.SetId(newRecord.Ref)
	if err = d.Set("internal_id", internalId.String()); err != nil {
		return err
	}
	if err = d.Set("ref", newRecord.Ref); err != nil {
		return err
	}

	return resourceDtcLbdnGet(d, m)
}

func resourceDtcLbdnGet(d *schema.ResourceData, m interface{}) error {

	var ttl int
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}

	rec, err := searchObjectByRefOrInternalId("DtcLbdn", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		} else {
			d.SetId("")
			return nil
		}
	}

	var dtcLbdn *ibclient.DtcLbdn
	var listInterface []interface{}
	connector := m.(ibclient.IBConnector)

	recJson, err := json.Marshal(rec)
	if err != nil {
		return fmt.Errorf("failed to marshal DTC LBDN record : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &dtcLbdn)
	if err != nil {
		return fmt.Errorf("failed getting DTC LBDN record : %s", err.Error())
	}

	delete(dtcLbdn.Ea, eaNameForInternalId)
	omittedEAs := omitEAs(dtcLbdn.Ea, extAttrs)

	if omittedEAs != nil && len(omittedEAs) > 0 {
		eaJSON, err := terraformSerializeEAs(omittedEAs)
		if err != nil {
			return err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return err
		}
	}

	if dtcLbdn.Name != nil {
		if err = d.Set("name", *dtcLbdn.Name); err != nil {
			return err
		}
	}
	if dtcLbdn.AuthZones != nil {
		authZoneInterface, err := ConvertAuthZonesToInterface(connector, dtcLbdn)
		if err != nil {
			return fmt.Errorf("failed to convert auth zones to interface: %w", err)
		}
		if err = d.Set("auth_zones", authZoneInterface); err != nil {
			return err
		}
	}

	if dtcLbdn.AutoConsolidatedMonitors != nil {
		if err = d.Set("auto_consolidated_monitors", *dtcLbdn.AutoConsolidatedMonitors); err != nil {
			return err
		}
	}
	if dtcLbdn.Comment != nil {
		if err = d.Set("comment", *dtcLbdn.Comment); err != nil {
			return err
		}
	}

	if dtcLbdn.Disable != nil {
		if err = d.Set("disable", *dtcLbdn.Disable); err != nil {
			return err
		}
	}
	if dtcLbdn.LbMethod != "" {
		if err = d.Set("lb_method", dtcLbdn.LbMethod); err != nil {
			return err
		}
	}
	if dtcLbdn.Patterns != nil {
		listInterface = convertSliceToInterface(dtcLbdn.Patterns)
		if err = d.Set("patterns", listInterface); err != nil {
			return err
		}
	}

	listInterface = convertSliceToInterface(dtcLbdn.Types)
	if err = d.Set("types", listInterface); err != nil {
		return err
	}

	if dtcLbdn.Persistence != nil {
		if err = d.Set("persistence", *dtcLbdn.Persistence); err != nil {
			return err
		}
	}
	if dtcLbdn.Priority != nil {
		if err = d.Set("priority", *dtcLbdn.Priority); err != nil {
			return err
		}
	}
	if dtcLbdn.Pools != nil {
		poolsInterface, err := convertPoolsToInterface(dtcLbdn, connector)
		if err != nil {
			return fmt.Errorf("failed to convert pools to interface: %w", err)
		}
		if err = d.Set("pools", poolsInterface); err != nil {
			return err
		}
	}

	if dtcLbdn.Topology != nil {
		var res ibclient.DtcTopology
		err := connector.GetObject(&ibclient.DtcTopology{}, *dtcLbdn.Topology, nil, &res)
		if err != nil {
			return fmt.Errorf("failed to get %s topology: %w", *dtcLbdn.Topology, err)
		}
		if err = d.Set("topology", *res.Name); err != nil {
			return err
		}
	}
	if dtcLbdn.Ttl != nil {
		ttl = int(*dtcLbdn.Ttl)
	}
	if !*dtcLbdn.UseTtl {
		ttl = ttlUndef
	}
	if err = d.Set("ttl", ttl); err != nil {
		return err
	}

	if err = d.Set("ref", dtcLbdn.Ref); err != nil {
		return err
	}

	return nil
}

func ConvertAuthZonesToInterface(connector ibclient.IBConnector, dtcLbdn *ibclient.DtcLbdn) ([]interface{}, error) {
	if len(dtcLbdn.AuthZones) == 0 {
		return nil, nil
	}
	authZoneInterface := make([]interface{}, len(dtcLbdn.AuthZones))
	for i, authZone := range dtcLbdn.AuthZones {
		var res ibclient.ZoneAuth
		err := connector.GetObject(&ibclient.ZoneAuth{}, authZone.Ref, nil, &res)
		if err != nil {
			return nil, err
		}
		authZoneInterface[i] = res.Fqdn
	}
	return authZoneInterface, nil
}

func convertSliceToInterface(list []string) []interface{} {
	if len(list) == 0 {
		return nil
	}
	listInterface := make([]interface{}, len(list))
	for i, j := range list {
		listInterface[i] = j
	}
	return listInterface
}

func convertPoolsToInterface(dtcLbdn *ibclient.DtcLbdn, connector ibclient.IBConnector) ([]interface{}, error) {

	poolsInterface := make([]interface{}, len(dtcLbdn.Pools))
	for i, pool := range dtcLbdn.Pools {
		var res ibclient.DtcPool
		err := connector.GetObject(&ibclient.DtcPool{}, pool.Pool, nil, &res)
		if err != nil {
			return nil, err
		}
		poolInterface := map[string]interface{}{
			"pool":  *res.Name,
			"ratio": pool.Ratio,
		}
		poolsInterface[i] = poolInterface
	}
	return poolsInterface, nil
}

func resourceDtcLbdnUpdate(d *schema.ResourceData, m interface{}) error {
	var updateSuccessful bool
	defer func() {
		if !updateSuccessful {
			prevName, _ := d.GetChange("name")
			prevAuthZones, _ := d.GetChange("auth_zones")
			prevAutoConsolidatedMonitors, _ := d.GetChange("auto_consolidated_monitors")
			prevComment, _ := d.GetChange("comment")
			prevDisable, _ := d.GetChange("disable")
			prevLbMethod, _ := d.GetChange("lb_method")
			prevPatterns, _ := d.GetChange("patterns")
			prevPersistence, _ := d.GetChange("persistence")
			prevPools, _ := d.GetChange("pools")
			prevPriority, _ := d.GetChange("priority")
			prevTopology, _ := d.GetChange("topology")
			prevTypes, _ := d.GetChange("types")
			prevTtl, _ := d.GetChange("ttl")
			prevExtAttrs, _ := d.GetChange("ext_attrs")

			_ = d.Set("name", prevName.(string))
			_ = d.Set("auth_zones", prevAuthZones)
			_ = d.Set("auto_consolidated_monitors", prevAutoConsolidatedMonitors.(bool))
			_ = d.Set("comment", prevComment.(string))
			_ = d.Set("disable", prevDisable.(bool))
			_ = d.Set("lb_method", prevLbMethod.(string))
			_ = d.Set("patterns", prevPatterns)
			_ = d.Set("persistence", prevPersistence.(int))
			_ = d.Set("pools", prevPools)
			_ = d.Set("priority", prevPriority.(int))
			_ = d.Set("topology", prevTopology.(string))
			_ = d.Set("types", prevTypes)
			_ = d.Set("ttl", prevTtl.(int))
			_ = d.Set("ext_attrs", prevExtAttrs.(string))
		}
	}()

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

	name := d.Get("name").(string)
	authZones := d.Get("auth_zones").([]interface{})
	authZoneList := make([]string, len(authZones))
	for i, authZone := range authZones {
		authZoneList[i] = authZone.(string)
	}
	autoConsolidatedMonitors := d.Get("auto_consolidated_monitors").(bool)
	comment := d.Get("comment").(string)
	disable := d.Get("disable").(bool)
	lbMethod := d.Get("lb_method").(string)
	poolsLink := d.Get("pools").([]interface{})

	pools, err := validatePoolsLink(poolsLink)
	if err != nil {
		return fmt.Errorf("failed to validate pools: %w", err)
	}

	patterns := d.Get("patterns").([]interface{})
	patternsList := make([]string, len(patterns))
	for i, pattern := range patterns {
		patternsList[i] = pattern.(string)
	}
	tempPersistence := d.Get("persistence").(int)
	persistence := uint32(tempPersistence)

	tempPriority := d.Get("priority").(int)
	priority := uint32(tempPriority)

	topology := d.Get("topology").(string)

	types := d.Get("types").([]interface{})
	typesList := make([]string, len(types))
	for i, j := range types {
		typesList[i] = j.(string)
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

	var lbdn *ibclient.DtcLbdn

	rec, err := searchObjectByRefOrInternalId("DtcLbdn", d, m)
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
		return fmt.Errorf("failed to marshal DTC LBDN record : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &lbdn)
	if err != nil {
		return fmt.Errorf("failed getting DTC LBDN record : %s", err.Error())
	}

	// If 'internal_id' is not set, then generate a new one and set it to the EA.
	internalId := d.Get("internal_id").(string)
	if internalId == "" {
		internalId = generateInternalId().String()
	}
	newInternalId := newInternalResourceIdFromString(internalId)
	newExtAttrs[eaNameForInternalId] = newInternalId.String()

	newExtAttrs, err = mergeEAs(lbdn.Ea, newExtAttrs, oldExtAttrs, connector)
	if err != nil {
		return err
	}

	lbdn, err = objMgr.UpdateDtcLbdn(d.Id(), name, authZoneList, comment, disable, autoConsolidatedMonitors, newExtAttrs, lbMethod, patternsList, persistence, pools, priority, topology, typesList, ttl, useTtl)
	if err != nil {
		return fmt.Errorf("failed to update DTC LBDN: %s.", err.Error())
	}

	updateSuccessful = true

	if err = d.Set("internal_id", newInternalId.String()); err != nil {
		return err
	}
	if err = d.Set("ref", lbdn.Ref); err != nil {
		return err
	}
	d.SetId(lbdn.Ref)
	return resourceDtcLbdnGet(d, m)
}

func resourceDtcLbdnDelete(d *schema.ResourceData, m interface{}) error {
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

	rec, err := searchObjectByRefOrInternalId("DtcLbdn", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		} else {
			d.SetId("")
			return nil
		}
	}

	var lbdn *ibclient.DtcLbdn
	recJson, err := json.Marshal(rec)
	if err != nil {
		return fmt.Errorf("failed to marshal DTC LBDN record : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &lbdn)
	if err != nil {
		return err
	}
	_, err = objMgr.DeleteDtcLbdn(lbdn.Ref)
	if err != nil {
		return fmt.Errorf("failed to delete DTC LBDN : %s", err.Error())
	}

	return nil
}

func resourceDtcLbdnImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	extAttrJSON := d.Get("ext_attrs").(string)
	_, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return nil, err
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	lbdn, err := objMgr.GetDtcLbdnByRef(d.Id())
	if err != nil {
		return nil, fmt.Errorf("failed getting DTC LBDN record: %w", err)
	}

	if lbdn.Ea != nil && len(lbdn.Ea) > 0 {
		eaJSON, err := terraformSerializeEAs(lbdn.Ea)
		if err != nil {
			return nil, err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return nil, err
		}
	}
	delete(lbdn.Ea, eaNameForInternalId)

	var listInterface []interface{}
	var ttl int

	if lbdn.Name != nil {
		if err = d.Set("name", *lbdn.Name); err != nil {
			return nil, err
		}
	}
	if lbdn.AuthZones != nil {
		authZoneInterface, err := ConvertAuthZonesToInterface(connector, lbdn)
		if err != nil {
			return nil, fmt.Errorf("failed to convert auth zones to interface: %w", err)
		}
		if err = d.Set("auth_zones", authZoneInterface); err != nil {
			return nil, err
		}
	}

	if lbdn.AutoConsolidatedMonitors != nil {
		if err = d.Set("auto_consolidated_monitors", *lbdn.AutoConsolidatedMonitors); err != nil {
			return nil, err
		}
	}
	if lbdn.Comment != nil {
		if err = d.Set("comment", *lbdn.Comment); err != nil {
			return nil, err
		}
	}

	if lbdn.Disable != nil {
		if err = d.Set("disable", *lbdn.Disable); err != nil {
			return nil, err
		}
	}
	if lbdn.LbMethod != "" {
		if err = d.Set("lb_method", lbdn.LbMethod); err != nil {
			return nil, err
		}
	}
	if lbdn.Patterns != nil {
		listInterface = convertSliceToInterface(lbdn.Patterns)
		if err = d.Set("patterns", listInterface); err != nil {
			return nil, err
		}
	}

	listInterface = convertSliceToInterface(lbdn.Types)
	if err = d.Set("types", listInterface); err != nil {
		return nil, err
	}

	if lbdn.Persistence != nil {
		if err = d.Set("persistence", *lbdn.Persistence); err != nil {
			return nil, err
		}
	}
	if lbdn.Priority != nil {
		if err = d.Set("priority", *lbdn.Priority); err != nil {
			return nil, err
		}
	}
	if lbdn.Pools != nil {
		poolsInterface, err := convertPoolsToInterface(lbdn, connector)
		if err != nil {
			return nil, fmt.Errorf("failed to convert pools to interface: %w", err)
		}
		if err = d.Set("pools", poolsInterface); err != nil {
			return nil, err
		}
	}

	if lbdn.Topology != nil {
		var res ibclient.DtcTopology
		err := connector.GetObject(&ibclient.DtcTopology{}, *lbdn.Topology, nil, &res)
		if err != nil {
			return nil, fmt.Errorf("failed to get %s topology: %w", *lbdn.Topology, err)
		}
		if err = d.Set("topology", *res.Name); err != nil {
			return nil, err
		}
	}
	if lbdn.Ttl != nil {
		ttl = int(*lbdn.Ttl)
	}
	if !*lbdn.UseTtl {
		ttl = ttlUndef
	}
	if err = d.Set("ttl", ttl); err != nil {
		return nil, err
	}

	d.SetId(lbdn.Ref)

	// Update the resource with the EA Terraform Internal ID
	err = resourceDtcLbdnUpdate(d, m)
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func validatePoolsLink(poolsLink []interface{}) ([]*ibclient.DtcPoolLink, error) {
	if poolsLink == nil {
		return nil, nil
	}
	dtcPoolLinks := make([]*ibclient.DtcPoolLink, 0, len(poolsLink))
	for _, item := range poolsLink {
		// Assert the type of item to map[string]interface{}
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("item is not of type map[string]interface{}")
		}

		// Create a new DtcPoolLink and populate its fields
		dtcPoolLink := &ibclient.DtcPoolLink{}
		if pool, ok := itemMap["pool"].(string); ok {
			dtcPoolLink.Pool = pool
		}
		if tempRatio, ok := itemMap["ratio"].(int); ok {
			dtcPoolLink.Ratio = uint32(tempRatio)
		}
		dtcPoolLinks = append(dtcPoolLinks, dtcPoolLink)
	}

	return dtcPoolLinks, nil
}
