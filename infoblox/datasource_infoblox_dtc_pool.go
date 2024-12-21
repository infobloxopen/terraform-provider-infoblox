package infoblox

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"strconv"
	//"strings"
	"time"
)

func datasourceDtcPool() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDtcPoolRecordRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:     schema.TypeMap,
				Required: true,
			},
			"results": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
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
						},
						"lb_preferred_topology": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The preferred topology for load balancing.",
						},
						"lb_alternate_method": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The alternate load balancing method. Use this to select a method type from the pool if the preferred method does not return any results.",
						},
						"lb_dynamic_ratio_alternate": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The DTC Pool settings for dynamic ratio when it’s selected as alternate method.",
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
					},
				},
			},
		},
	}
}

func dataSourceDtcPoolRecordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(ibclient.IBConnector)

	var diags diag.Diagnostics

	filters := filterFromMap(d.Get("filters").(map[string]interface{}))

	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	qp := ibclient.NewQueryParams(false, filters)
	res, err := objMgr.GetAllDtcPool(qp)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to get DTC pool records: %w", err))
	}

	if res == nil {
		return diag.FromErr(fmt.Errorf("API returns a nil/empty ID for DTC pool"))
	}
	results := make([]interface{}, 0, len(res))
	for _, r := range res {
		dtcPool, err := flattenDtcPool(r, connector)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to flatten DTC Pool : %w", err))
		}
		results = append(results, dtcPool)
	}

	err = d.Set("results", results)
	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenDtcPool(pool ibclient.DtcPool, connector ibclient.IBConnector) (map[string]interface{}, error) {
	var eaMap map[string]interface{}
	if pool.Ea != nil && len(pool.Ea) > 0 {
		eaMap = pool.Ea
	} else {
		eaMap = make(map[string]interface{})
	}

	ea, err := json.Marshal(eaMap)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":                  pool.Ref,
		"name":                pool.Name,
		"ext_attrs":           string(ea),
		"lb_preferred_method": pool.LbPreferredMethod,
	}
	if pool.AutoConsolidatedMonitors != nil {
		res["auto_consolidated_monitors"] = *pool.AutoConsolidatedMonitors
	}
	if pool.Availability != "" {
		res["availability"] = pool.Availability
	}
	if pool.Comment != nil {
		res["comment"] = *pool.Comment
	}
	if pool.ConsolidatedMonitors != nil {
		consolidatedMonitors, err := convertConsolidatedMonitorsToInterface(pool.ConsolidatedMonitors, connector)
		if err != nil {
			return nil, err
		}
		res["consolidated_monitors"] = consolidatedMonitors
	}
	if pool.Disable != nil {
		res["disable"] = *pool.Disable
	}
	if pool.LbAlternateMethod != "" {
		res["lb_alternate_method"] = pool.LbAlternateMethod
	}
	if pool.LbAlternateTopology != nil {
		res["lb_alternate_topology"] = *pool.LbAlternateTopology
	}
	if pool.LbDynamicRatioAlternate != nil && pool.LbAlternateMethod == "DYNAMIC_RATIO" {
		lbDynamicRatioAlternate, err := serializeSettingDynamicRatio(pool.LbDynamicRatioAlternate, connector)
		if err != nil {
			return nil, err
		}
		res["lb_dynamic_ratio_alternate"] = lbDynamicRatioAlternate
	}
	if pool.LbDynamicRatioPreferred != nil && pool.LbPreferredMethod == "DYNAMIC_RATIO" {
		lbDynamicRatioPreferred, err := serializeSettingDynamicRatio(pool.LbDynamicRatioPreferred, connector)
		if err != nil {
			return nil, err
		}
		res["lb_dynamic_ratio_preferred"] = lbDynamicRatioPreferred
	}
	if pool.LbPreferredTopology != nil {
		res["lb_preferred_topology"] = *pool.LbPreferredTopology
	}
	if pool.Monitors != nil {
		res["monitors"] = convertMonitorsToInterface(pool.Monitors, connector)
	}
	if pool.Quorum != nil {
		res["quorum"] = *pool.Quorum
	}
	if pool.Servers != nil {
		servers, err := convertDtcServerLinksToInterface(pool.Servers, connector)
		if err != nil {
			return nil, err
		}
		res["servers"] = servers
	}
	if pool.UseTtl != nil {
		if !*pool.UseTtl {
			res["ttl"] = ttlUndef
		}
	}
	if pool.Ttl != nil && *pool.Ttl > 0 {
		res["ttl"] = *pool.Ttl
	} else {
		res["ttl"] = ttlUndef
	}

	return res, nil
}

//func convertDtcServerLinksToInterface(serverLinks []*ibclient.DtcServerLink, connector ibclient.IBConnector) []map[string]interface{} {
//	slInterface := make([]map[string]interface{}, 0, len(serverLinks))
//	for _, sl := range serverLinks {
//		slMap := make(map[string]interface{})
//		var serverResult ibclient.DtcServer
//		err := connector.GetObject(&ibclient.DtcServer{}, sl.Server, nil, &serverResult)
//		//check for this err thing this is wrong to return nil
//		if err != nil {
//			return nil
//		}
//		slMap["server"] = serverResult.Name
//		slMap["ratio"] = sl.Ratio
//		slInterface = append(slInterface, slMap)
//	}
//	return slInterface
//}
//
//func convertConsolidatedMonitorsToInterface(monitors []*ibclient.DtcPoolConsolidatedMonitorHealth, connector ibclient.IBConnector) []map[string]interface{} {
//	monitorsInterface := make([]map[string]interface{}, 0, len(monitors))
//	for _, monitor := range monitors {
//		monitorMap := make(map[string]interface{})
//		var monitorResult ibclient.DtcMonitorHttp
//		err := connector.GetObject(&ibclient.DtcMonitorHttp{}, monitor.Monitor, nil, &monitorResult)
//		if err != nil {
//			return nil
//		}
//		referenceParts := strings.Split(monitor.Monitor, ":")
//		monitorType := strings.Split(referenceParts[2], "/")[0]
//		monitorMap["monitor_name"] = monitorResult.Name
//		monitorMap["monitor_type"] = monitorType
//		monitorMap["members"] = monitor.Members
//		monitorMap["availability"] = monitor.Availability
//		monitorMap["full_health_communication"] = monitor.FullHealthCommunication
//		monitorsInterface = append(monitorsInterface, monitorMap)
//	}
//	return monitorsInterface
//}
//
//func convertMonitorsToInterface(monitors []*ibclient.DtcMonitorHttp, connector ibclient.IBConnector) []map[string]interface{} {
//	monitorsInterface := make([]map[string]interface{}, 0, len(monitors))
//	for _, monitor := range monitors {
//		monitorMap := make(map[string]interface{})
//		var monitorResult ibclient.DtcMonitorHttp
//		err := connector.GetObject(&ibclient.DtcMonitorHttp{}, monitor.Ref, nil, &monitorResult)
//		if err != nil {
//			return nil
//		}
//		referenceParts := strings.Split(monitor.Ref, ":")
//		monitorType := strings.Split(referenceParts[2], "/")[0]
//		monitorMap["monitor_name"] = monitorResult.Name
//		monitorMap["monitor_type"] = monitorType
//		monitorsInterface = append(monitorsInterface, monitorMap)
//	}
//	return monitorsInterface
//}

func convertToMapList(input []interface{}) []map[string]interface{} {
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
