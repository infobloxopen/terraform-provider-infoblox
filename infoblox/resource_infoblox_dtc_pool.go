package infoblox

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func addDefaultValues(input string) (string, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		return "", err
	}
	if _, ok := data["invert_monitor_metric"]; !ok {
		data["invert_monitor_metric"] = false
	}
	if _, ok := data["method"]; !ok {
		data["method"] = "MONITOR"
	}
	if _, ok := data["monitor_metric"]; !ok {
		data["monitor_metric"] = ""
	}
	if _, ok := data["monitor_weighing"]; !ok {
		data["monitor_weighing"] = "RATIO"
	}

	output, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func convertInterfaceToList(input []interface{}) []map[string]interface{} {
	var result []map[string]interface{}

	for _, item := range input {
		if itemMap, ok := item.(map[string]interface{}); ok {
			monitor := ibclient.Monitor{
				Name: itemMap["monitor_name"].(string),
				Type: itemMap["monitor_type"].(string),
			}
			// Remove monitor_name and monitor_type from the map
			delete(itemMap, "monitor_name")
			delete(itemMap, "monitor_type")
			// Append the Monitor struct to the map
			itemMap["monitor"] = monitor
			if members, ok1 := itemMap["members"].([]interface{}); ok1 {
				var membersStr []string
				for _, member := range members {
					if memberStr, ok2 := member.(string); ok2 {
						membersStr = append(membersStr, memberStr)
					}
				}
				itemMap["members"] = membersStr
			}
			result = append(result, itemMap)
		}
	}
	return result
}

func ConvertDynamicRatioPreferredToInterface(jsonStr string) (map[string]interface{}, error) {
	var lbDynamicRatioPreferred map[string]interface{}
	if jsonStr != "" {
		err := json.Unmarshal([]byte(jsonStr), &lbDynamicRatioPreferred)
		if err != nil {
			return nil, err
		}
		monitor := ibclient.Monitor{}
		if monitorName, ok := lbDynamicRatioPreferred["monitor_name"]; ok {
			monitor.Name = monitorName.(string)
			delete(lbDynamicRatioPreferred, "monitor_name")
		}
		if monitorType, ok := lbDynamicRatioPreferred["monitor_type"]; ok {
			monitor.Type = monitorType.(string)
			delete(lbDynamicRatioPreferred, "monitor_type")
		}
		lbDynamicRatioPreferred["monitor"] = monitor
	}
	return lbDynamicRatioPreferred, nil
}

func ConvertInterfaceToServers(serversInterface []interface{}) []*ibclient.DtcServerLink {
	var servers []*ibclient.DtcServerLink
	for _, serverInterface := range serversInterface {
		server := serverInterface.(map[string]interface{})
		dtcServerLink := &ibclient.DtcServerLink{
			Server: server["server"].(string),
			Ratio:  uint32(server["ratio"].(int)),
		}
		servers = append(servers, dtcServerLink)
	}
	return servers
}

func ConvertInterfaceToMonitors(monitorsInterface []interface{}) []ibclient.Monitor {
	var monitors []ibclient.Monitor

	for _, monitor := range monitorsInterface {
		monitorMap := monitor.(map[string]interface{})

		monitorStruct := ibclient.Monitor{
			Name: monitorMap["monitor_name"].(string),
			Type: monitorMap["monitor_type"].(string),
		}

		monitors = append(monitors, monitorStruct)
	}

	return monitors
}

func resourceDtcPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceDtcPoolCreate,
		Read:   resourceDtcPoolGet,
		Update: resourceDtcPoolUpdate,
		Delete: resourceDtcPoolDelete,
		Importer: &schema.ResourceImporter{
			State: resourceDtcPoolImport,
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
				Description: "Name of the DTC pool.",
			},
			"auto_consolidated_monitors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Flag for enabling auto managing DTC Consolidated Monitors in DTC Pool.",
			},
			"availability": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "ALL",
				Description: "A resource in the pool is available if ANY, at least QUORUM, or ALL monitors for the pool say that it is up.",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Description of the Dtc pool.",
			},
			"consolidated_monitors": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of monitors and associated members statuses of which are shared across members and consolidated in server availability determination.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"monitor_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the monitor ",
						},
						"monitor_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The type of the monitor",
						},
						"members": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "Members whose monitor statuses are shared across other members in a pool",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"availability": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Servers assigned to a pool with monitor defined are healthy if ANY or ALL members report healthy status.",
						},
						"full_health_communication": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Flag for switching health performing and sharing behavior to perform health checks on each DTC grid member that serves related LBDN(s) and send them across all DTC grid members from both selected and non-selected lists.",
						},
					},
				},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					autoConsolidated, ok := d.GetOk("auto_consolidated_monitors")
					if ok && autoConsolidated.(bool) {
						return true // Suppress differences when auto_consolidated_monitors is true
					}
					return false
				},
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
				Default:     "",
				Description: "Extensible attributes of the  Dtc Pool to be added/updated, as a map in JSON format",
			},
			"lb_preferred_method": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Load Balancing Preferred Method of the DTC pool.",
			},
			"lb_dynamic_ratio_preferred": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The DTC Pool settings for dynamic ratio when it’s selected as preferred method.",
				StateFunc: func(val interface{}) string {
					input := val.(string)
					output, err := addDefaultValues(input)
					if err != nil {
						return input // Return the original input in case of error
					}
					return output
				},
			},
			"lb_preferred_topology": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The preferred topology for load balancing.",
			},
			"lb_alternate_method": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "NONE",
				Description: "The alternate load balancing method. Use this to select a method type from the pool if the preferred method does not return any results.",
			},
			"lb_dynamic_ratio_alternate": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The DTC Pool settings for dynamic ratio when it’s selected as alternate method.",
				StateFunc: func(val interface{}) string {
					input := val.(string)
					output, err := addDefaultValues(input)
					if err != nil {
						return input // Return the original input in case of error
					}
					return output
				},
			},
			"lb_alternate_topology": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The alternate topology for load balancing.",
			},
			"monitors": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Monitors associated with the DTC pool.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"monitor_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the monitor",
						},
						"monitor_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The type of the monitor.",
						},
					},
				},
			},
			"servers": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Servers of the DTC pool.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the server for the pool",
						},
						"ratio": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The weight of server.",
						},
					},
				},
			},
			"quorum": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "For availability mode QUORUM, at least this many monitors must report the resource as up for it to be available",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     ttlUndef,
				Description: "TTL value for the Dtc Pool.",
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

func resourceDtcPoolCreate(d *schema.ResourceData, m interface{}) error {
	if intId := d.Get("internal_id"); intId.(string) != "" {
		return fmt.Errorf("the value of 'internal_id' field must not be set manually")
	}
	name := d.Get("name").(string)
	comment := d.Get("comment").(string)
	lbPreferredMethod := d.Get("lb_preferred_method").(string)
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

	serversInterface := d.Get("servers").([]interface{})
	servers := ConvertInterfaceToServers(serversInterface)

	monitorsInterface := d.Get("monitors").([]interface{})
	monitors := ConvertInterfaceToMonitors(monitorsInterface)

	lbDynamicRatioJson := d.Get("lb_dynamic_ratio_preferred").(string)
	lbDynamicRatioPreferred, err := ConvertDynamicRatioPreferredToInterface(lbDynamicRatioJson)
	if err != nil {
		return err
	}
	lbPreferredTopologyValue := d.Get("lb_preferred_topology").(string)
	var lbPreferredTopology *string
	if lbPreferredTopologyValue != "" {
		lbPreferredTopology = &lbPreferredTopologyValue
	}
	lbAlternateMethod := d.Get("lb_alternate_method").(string)
	autoConsolidatedMonitors := d.Get("auto_consolidated_monitors").(bool)
	disable := d.Get("disable").(bool)
	availability := d.Get("availability").(string)
	lbAlternateTopologyValue := d.Get("lb_alternate_topology").(string)
	var lbAlternateTopology *string
	if lbAlternateTopologyValue != "" {
		lbAlternateTopology = &lbAlternateTopologyValue
	}
	lbDynamicRatioAlternateJson := d.Get("lb_dynamic_ratio_alternate").(string)
	lbDynamicRatioAlternate, err := ConvertDynamicRatioPreferredToInterface(lbDynamicRatioAlternateJson)
	if err != nil {
		return err
	}
	quorum := uint32(d.Get("quorum").(int))
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	newDtcPool, err := objMgr.CreateDtcPool(comment, name, lbPreferredMethod, lbDynamicRatioPreferred, servers, monitors, lbPreferredTopology, lbAlternateMethod, lbAlternateTopology, lbDynamicRatioAlternate, extAttrs, autoConsolidatedMonitors, availability, ttl, useTtl, disable, quorum)
	if err != nil {
		return err
	}
	d.SetId(newDtcPool.Ref)
	if err = d.Set("ref", newDtcPool.Ref); err != nil {
		return err
	}
	if err = d.Set("internal_id", internalId.String()); err != nil {
		return err
	}
	return resourceDtcPoolGet(d, m)
}

func resourceDtcPoolGet(d *schema.ResourceData, m interface{}) error {
	var ttl int
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}

	connector := m.(ibclient.IBConnector)
	rec, err := searchObjectByRefOrInternalId("DtcPool", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		} else {
			d.SetId("")
			return nil
		}
	}
	var dtcPool *ibclient.DtcPool
	recJson, err := json.Marshal(rec)
	if err != nil {
		return fmt.Errorf("failed to marshal DTC Pool : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &dtcPool)
	if err != nil {
		return fmt.Errorf("failed getting DTC pool : %s", err.Error())
	}
	delete(dtcPool.Ea, eaNameForInternalId)
	omittedEAs := omitEAs(dtcPool.Ea, extAttrs)

	if omittedEAs != nil && len(omittedEAs) > 0 {
		eaJSON, err := terraformSerializeEAs(omittedEAs)
		if err != nil {
			return err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return err
		}
	}
	if dtcPool.Ttl != nil {
		ttl = int(*dtcPool.Ttl)
	}
	if !*dtcPool.UseTtl {
		ttl = ttlUndef
	}
	if err = d.Set("ttl", ttl); err != nil {
		return err
	}
	if err = d.Set("name", dtcPool.Name); err != nil {
		return err
	}
	if err = d.Set("comment", dtcPool.Comment); err != nil {
		return err
	}
	if err = d.Set("disable", dtcPool.Disable); err != nil {
		return err
	}
	if err = d.Set("lb_preferred_method", dtcPool.LbPreferredMethod); err != nil {
		return err
	}
	if err = d.Set("auto_consolidated_monitors", dtcPool.AutoConsolidatedMonitors); err != nil {
		return err
	}
	consolidatedMonitorsInterface, err := convertConsolidatedMonitorsToInterface(dtcPool.ConsolidatedMonitors, connector)
	if err != nil {
		return err
	}
	if err = d.Set("consolidated_monitors", consolidatedMonitorsInterface); err != nil {
		return err
	}
	slInterface, err := convertDtcServerLinksToInterface(dtcPool.Servers, connector)
	if err != nil {
		return err
	}
	if err = d.Set("servers", slInterface); err != nil {
		return err
	}
	monitorsInterface := convertMonitorsToInterface(dtcPool.Monitors, connector)
	if err = d.Set("monitors", monitorsInterface); err != nil {
		return err
	}
	if dtcPool.LbPreferredTopology != nil {
		var topologies ibclient.DtcTopology
		err = connector.GetObject(&ibclient.DtcTopology{}, *dtcPool.LbPreferredTopology, nil, &topologies)
		topologyPreferredName := topologies.Name
		if err = d.Set("lb_preferred_topology", topologyPreferredName); err != nil {
			return err
		}
	} else {
		if err = d.Set("lb_preferred_topology", nil); err != nil {
			return err
		}
	}

	if dtcPool.LbDynamicRatioPreferred != nil && dtcPool.LbPreferredMethod == "DYNAMIC_RATIO" {
		dynamicRatioInterface, _ := serializeSettingDynamicRatio(dtcPool.LbDynamicRatioPreferred, connector)
		if err := d.Set("lb_dynamic_ratio_preferred", dynamicRatioInterface); err != nil {
			return err
		}
	} else {
		if err := d.Set("lb_dynamic_ratio_preferred", nil); err != nil {
			return err
		}
	}

	if err = d.Set("lb_alternate_method", dtcPool.LbAlternateMethod); err != nil {
		return err
	}
	if dtcPool.LbDynamicRatioAlternate != nil && dtcPool.LbAlternateMethod == "DYNAMIC_RATIO" {
		dynamicRatioInterface, _ := serializeSettingDynamicRatio(dtcPool.LbDynamicRatioAlternate, connector)
		if err := d.Set("lb_dynamic_ratio_alternate", dynamicRatioInterface); err != nil {
			return err
		}
	} else {
		if err := d.Set("lb_dynamic_ratio_alternate", nil); err != nil {
			return err
		}
	}
	if dtcPool.LbAlternateTopology != nil {
		var topologiesAlternate ibclient.DtcTopology
		err = connector.GetObject(&ibclient.DtcTopology{}, *dtcPool.LbAlternateTopology, nil, &topologiesAlternate)
		topologyAlternateName := topologiesAlternate.Name
		if err = d.Set("lb_alternate_topology", topologyAlternateName); err != nil {
			return err
		}
	} else {
		if err = d.Set("lb_alternate_topology", nil); err != nil {
			return err
		}
	}

	return nil
}

func resourceDtcPoolUpdate(d *schema.ResourceData, m interface{}) error {

	var updateSuccessful bool
	defer func() {
		// Reverting the state back, in case of a failure,
		// otherwise Terraform will keep the values, which leaded to the failure, in the state file.
		if !updateSuccessful {
			prevComment, _ := d.GetChange("comment")
			prevName, _ := d.GetChange("name")
			prevAutoConsolidatedMonitors, _ := d.GetChange("auto_consolidated_monitors")
			prevAvailability, _ := d.GetChange("availability")
			prevConsolidatedMonitors, _ := d.GetChange("consolidated_monitors")
			prevDisable, _ := d.GetChange("disable")
			prevEa, _ := d.GetChange("ext_attrs")
			prevLbPreferredMethod, _ := d.GetChange("lb_preferred_method")
			prevLbDynamicRatioPreferred, _ := d.GetChange("lb_dynamic_ratio_preferred")
			prevLbPreferredTopology, _ := d.GetChange("lb_preferred_topology")
			prevLbAlternateMethod, _ := d.GetChange("lb_alternate_method")
			prevLbDynamicRatioAlternate, _ := d.GetChange("lb_dynamic_ratio_alternate")
			prevLbAlternateTopology, _ := d.GetChange("lb_alternate_topology")
			prevMonitors, _ := d.GetChange("monitors")
			prevServers, _ := d.GetChange("servers")
			prevQuorum, _ := d.GetChange("quorum")
			prevTTL, _ := d.GetChange("ttl")

			_ = d.Set("comment", prevComment.(string))
			_ = d.Set("name", prevName.(string))
			_ = d.Set("auto_consolidated_monitors", prevAutoConsolidatedMonitors.(bool))
			_ = d.Set("availability", prevAvailability.(string))
			_ = d.Set("consolidated_monitors", prevConsolidatedMonitors)
			_ = d.Set("disable", prevDisable.(bool))
			_ = d.Set("ext_attrs", prevEa.(string))
			_ = d.Set("lb_preferred_method", prevLbPreferredMethod.(string))
			_ = d.Set("lb_dynamic_ratio_preferred", prevLbDynamicRatioPreferred.(string))
			_ = d.Set("lb_preferred_topology", prevLbPreferredTopology.(string))
			_ = d.Set("lb_alternate_method", prevLbAlternateMethod.(string))
			_ = d.Set("lb_dynamic_ratio_alternate", prevLbDynamicRatioAlternate.(string))
			_ = d.Set("lb_alternate_topology", prevLbAlternateTopology.(string))
			_ = d.Set("monitors", prevMonitors)
			_ = d.Set("servers", prevServers)
			_ = d.Set("quorum", prevQuorum.(int))
			_ = d.Set("ttl", prevTTL.(int))
		}
	}()
	if d.HasChange("internal_id") {
		return fmt.Errorf("changing the value of 'internal_id' field is not allowed")
	}
	name := d.Get("name").(string)
	comment := d.Get("comment").(string)
	lbPreferredMethod := d.Get("lb_preferred_method").(string)

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

	serversInterface := d.Get("servers").([]interface{})
	servers := ConvertInterfaceToServers(serversInterface)

	monitorsInterface := d.Get("monitors").([]interface{})
	monitors := ConvertInterfaceToMonitors(monitorsInterface)

	lbDynamicRatioJson := d.Get("lb_dynamic_ratio_preferred").(string)
	lbDynamicRatioPreferred, err := ConvertDynamicRatioPreferredToInterface(lbDynamicRatioJson)
	if err != nil {
		return err
	}
	lbPreferredTopologyValue := d.Get("lb_preferred_topology").(string)
	var lbPreferredTopology *string
	if lbPreferredTopologyValue != "" {
		lbPreferredTopology = &lbPreferredTopologyValue
	}
	lbAlternateMethod := d.Get("lb_alternate_method").(string)
	autoConsolidatedMonitors := d.Get("auto_consolidated_monitors").(bool)
	disable := d.Get("disable").(bool)
	availability := d.Get("availability").(string)
	lbAlternateTopologyValue := d.Get("lb_alternate_topology").(string)
	var lbAlternateTopology *string
	if lbAlternateTopologyValue != "" {
		lbAlternateTopology = &lbAlternateTopologyValue
	}
	lbDynamicRatioAlternateJson := d.Get("lb_dynamic_ratio_alternate").(string)
	lbDynamicRatioAlternate, err := ConvertDynamicRatioPreferredToInterface(lbDynamicRatioAlternateJson)
	if err != nil {
		return err
	}
	quorum := uint32(d.Get("quorum").(int))

	consolidatedMonitorsInterface := d.Get("consolidated_monitors").([]interface{})
	consolidatedMonitors := convertInterfaceToList(consolidatedMonitorsInterface)
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

	var dtcPool *ibclient.DtcPool

	rec, err := searchObjectByRefOrInternalId("DtcPool", d, m)
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
		return fmt.Errorf("failed to marshal Dtc Pool : %s", err.Error())
	}
	err = json.Unmarshal(recJson, &dtcPool)
	if err != nil {
		return fmt.Errorf("failed getting Dtc Pool : %s", err.Error())
	}

	// If 'internal_id' is not set, then generate a new one and set it to the EA.
	internalId := d.Get("internal_id").(string)
	if internalId == "" {
		internalId = generateInternalId().String()
	}
	newInternalId := newInternalResourceIdFromString(internalId)
	newExtAttrs[eaNameForInternalId] = newInternalId.String()

	newExtAttrs, err = mergeEAs(dtcPool.Ea, newExtAttrs, oldExtAttrs, connector)
	if err != nil {
		return err
	}
	dtcPool, err = objMgr.UpdateDtcPool(d.Id(), comment, name, lbPreferredMethod, lbDynamicRatioPreferred, servers, monitors, lbPreferredTopology, lbAlternateMethod, lbAlternateTopology, lbDynamicRatioAlternate, newExtAttrs, autoConsolidatedMonitors, availability, consolidatedMonitors, ttl, useTtl, disable, quorum)
	if err != nil {
		return fmt.Errorf("error updating dtc-pool: %w", err)
	}
	updateSuccessful = true
	d.SetId(dtcPool.Ref)
	if err = d.Set("ref", dtcPool.Ref); err != nil {
		return err
	}
	if err = d.Set("internal_id", newInternalId.String()); err != nil {
		return err
	}
	return resourceDtcPoolGet(d, m)
}

func resourceDtcPoolDelete(d *schema.ResourceData, m interface{}) error {
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

	rec, err := searchObjectByRefOrInternalId("DtcPool", d, m)
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
	var dtcPool *ibclient.DtcPool
	recJson, _ := json.Marshal(rec)
	err = json.Unmarshal(recJson, &dtcPool)

	_, err = objMgr.DeleteDtcPool(dtcPool.Ref)
	if err != nil {
		return fmt.Errorf("deletion of Dtc Pool failed: %w", err)
	}
	d.SetId("")

	return nil
}

func resourceDtcPoolImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
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
	obj, err := objMgr.GetDtcPoolByRef(d.Id())
	if err != nil {
		return nil, fmt.Errorf("getting DtcPool with ID: %s failed: %w", d.Id(), err)
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

	if err = d.Set("ttl", ttl); err != nil {
		return nil, err
	}
	if err = d.Set("name", obj.Name); err != nil {
		return nil, err
	}
	if err = d.Set("comment", obj.Comment); err != nil {
		return nil, err
	}
	if err = d.Set("disable", obj.Disable); err != nil {
		return nil, err
	}
	if err = d.Set("lb_preferred_method", obj.LbPreferredMethod); err != nil {
		return nil, err
	}
	if err = d.Set("auto_consolidated_monitors", obj.AutoConsolidatedMonitors); err != nil {
		return nil, err
	}
	consolidatedMonitorsInterface, err := convertConsolidatedMonitorsToInterface(obj.ConsolidatedMonitors, connector)
	if err != nil {
		return nil, err
	}
	if err = d.Set("consolidated_monitors", consolidatedMonitorsInterface); err != nil {
		return nil, err
	}
	slInterface, err := convertDtcServerLinksToInterface(obj.Servers, connector)
	if err != nil {
		return nil, err
	}
	if err = d.Set("servers", slInterface); err != nil {
		return nil, err
	}
	monitorsInterface := convertMonitorsToInterface(obj.Monitors, connector)
	if err = d.Set("monitors", monitorsInterface); err != nil {
		return nil, err
	}

	if obj.LbPreferredTopology != nil {
		var topologies ibclient.DtcTopology
		err = connector.GetObject(&ibclient.DtcTopology{}, *obj.LbPreferredTopology, nil, &topologies)
		topologyPreferredName := topologies.Name
		if err = d.Set("lb_preferred_topology", topologyPreferredName); err != nil {
			return nil, err
		}
	} else {
		if err = d.Set("lb_preferred_topology", nil); err != nil {
			return nil, err
		}
	}

	if obj.LbDynamicRatioPreferred != nil && obj.LbPreferredMethod == "DYNAMIC_RATIO" {
		dynamicRatioInterface, _ := serializeSettingDynamicRatio(obj.LbDynamicRatioPreferred, connector)
		if err := d.Set("lb_dynamic_ratio_preferred", dynamicRatioInterface); err != nil {
			return nil, err
		}
	} else {
		if err := d.Set("lb_dynamic_ratio_preferred", nil); err != nil {
			return nil, err
		}
	}
	if err = d.Set("lb_alternate_method", obj.LbAlternateMethod); err != nil {
		return nil, err
	}
	if obj.LbDynamicRatioAlternate != nil && obj.LbAlternateMethod == "DYNAMIC_RATIO" {
		dynamicRatioInterface, _ := serializeSettingDynamicRatio(obj.LbDynamicRatioAlternate, connector)
		if err := d.Set("lb_dynamic_ratio_alternate", dynamicRatioInterface); err != nil {
			return nil, err
		}
	} else {
		if err := d.Set("lb_dynamic_ratio_alternate", nil); err != nil {
			return nil, err
		}
	}
	if obj.LbAlternateTopology != nil {
		var topologiesAlternate ibclient.DtcTopology
		err = connector.GetObject(&ibclient.DtcTopology{}, *obj.LbAlternateTopology, nil, &topologiesAlternate)
		topologyAlternateName := topologiesAlternate.Name
		if err = d.Set("lb_alternate_topology", topologyAlternateName); err != nil {
			return nil, err
		}
	} else {
		if err = d.Set("lb_alternate_topology", nil); err != nil {
			return nil, err
		}
	}
	d.SetId(obj.Ref)

	err = resourceDtcPoolUpdate(d, m)
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
