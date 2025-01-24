package infoblox

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"strconv"
	"time"
)

func dataSourceDtcLbdnRecord() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDtcLbdnRecordRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:     schema.TypeMap,
				Required: true,
			},
			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of DTC LBDN Records matching filters",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The display name of the DTC LBDN.",
						},
						"auth_zones": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of linked auth zones.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"auto_consolidated_monitors": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Flag for enabling auto managing DTC Consolidated Monitors on related DTC Pools.",
						},
						"comment": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of the DTC LBDN record.",
						},
						"ext_attrs": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Extensible attributes of the DTC LBDN record to be added/updated, as a map in JSON format.",
						},
						"disable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Determines whether the DTC LBDN is disabled or not. When this is set to False, the fixed address is enabled.",
						},
						"lb_method": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The load balancing method. Used to select pool. Valid values are GLOBAL_AVAILABILITY, RATIO, ROUND_ROBIN, SOURCE_IP_HASH and TOPOLOGY.",
						},
						"patterns": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "LBDN wildcards for pattern match.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"persistence": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum time, in seconds, for which client specific LBDN responses will be cached. Zero specifies no caching.",
						},
						"pools": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Pools associated with an LBDN are collections of load-balanced servers",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pool": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The pool to link with.",
									},
									"ratio": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "The weight of pool.",
									},
								},
							},
						},
						"priority": {
							Type:     schema.TypeInt,
							Computed: true,
							Description: "The LBDN pattern match priority for “overlapping” DTC LBDN objects. LBDNs are “overlapping” if " +
								"they are simultaneously assigned to a zone and have patterns that can match the same FQDN. The matching LBDN with highest priority (lowest ordinal) will be used.",
						},
						"topology": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The topology rules for TOPOLOGY method.",
						},
						"ttl": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The Time To Live (TTL) value for the DTC LBDN. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached.",
						},
						"types": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of resource record types supported by LBDN. Valid values are A, AAAA, CNAME, NAPTR, SRV. Default value is A and AAAA",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"health": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The LBDN health information.",
						},
					},
				},
			},
		},
	}
}

func dataSourceDtcLbdnRecordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(ibclient.IBConnector)

	var diags diag.Diagnostics

	filters := filterFromMap(d.Get("filters").(map[string]interface{}))

	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	qp := ibclient.NewQueryParams(false, filters)
	res, err := objMgr.GetAllDtcLbdn(qp)
	if err != nil {
		return diag.FromErr(err)
	}

	if res == nil {
		return diag.FromErr(fmt.Errorf("API returns a nil/empty ID for DTC LBDN"))
	}
	results := make([]interface{}, 0, len(res))
	for _, r := range res {
		dtcLbdn, err := flattenDtcLbdn(r, connector)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to flatten DTC LBDN : %w", err))
		}
		results = append(results, dtcLbdn)
	}

	err = d.Set("results", results)
	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenDtcLbdn(lbdn ibclient.DtcLbdn, connector ibclient.IBConnector) (map[string]interface{}, error) {
	var eaMap map[string]interface{}
	if lbdn.Ea != nil && len(lbdn.Ea) > 0 {
		eaMap = lbdn.Ea
	} else {
		eaMap = make(map[string]interface{})
	}

	ea, err := json.Marshal(eaMap)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":        lbdn.Ref,
		"name":      lbdn.Name,
		"ext_attrs": string(ea),
		"lb_method": lbdn.LbMethod,
	}
	if lbdn.Comment != nil {
		res["comment"] = *lbdn.Comment
	}
	if lbdn.Disable != nil {
		res["disable"] = *lbdn.Disable
	}
	if lbdn.AuthZones != nil {
		authZones, err := ConvertAuthZonesToInterface(connector, &lbdn)
		if err != nil {
			return nil, err
		}
		res["auth_zones"] = authZones
	}

	if lbdn.AutoConsolidatedMonitors != nil {
		res["auto_consolidated_monitors"] = *lbdn.AutoConsolidatedMonitors
	}

	if lbdn.Patterns != nil {
		res["patterns"] = convertSliceToInterface(lbdn.Patterns)
	}

	if lbdn.Persistence != nil {
		res["persistence"] = lbdn.Persistence
	}

	if lbdn.Pools != nil {
		pools, err := convertPoolsToInterface(&lbdn, connector)
		if err != nil {
			return nil, err
		} else {
			res["pools"] = pools
		}
	}
	if lbdn.Priority != nil {
		res["priority"] = lbdn.Priority
	}
	if lbdn.Topology != nil {
		var topology ibclient.DtcTopology
		err := connector.GetObject(&ibclient.DtcTopology{}, *lbdn.Topology, nil, &topology)
		if err != nil {
			return nil, fmt.Errorf("error getting %s DtcTopology object: %s", *lbdn.Topology, err)
		}
		res["topology"] = *topology.Name
	}
	if lbdn.Types != nil {
		res["types"] = convertSliceToInterface(lbdn.Types)
	}
	if lbdn.UseTtl != nil {
		if !*lbdn.UseTtl {
			res["ttl"] = ttlUndef
		}
	}
	if lbdn.Ttl != nil && *lbdn.Ttl > 0 {
		res["ttl"] = *lbdn.Ttl
	} else {
		res["ttl"] = ttlUndef
	}
	if lbdn.Health != nil {
		res["health"] = ConvertDtcHealthToMap(lbdn.Health)
	}
	return res, nil
}
