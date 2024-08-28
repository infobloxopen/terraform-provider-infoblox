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

func dataSourceZoneDelegated() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceZoneDelegatedRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:     schema.TypeMap,
				Required: true,
			},
			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of Forward Zones matching filters",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
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
					},
				},
			},
		},
	}
}

func dataSourceZoneDelegatedRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(ibclient.IBConnector)

	var diags diag.Diagnostics

	filters := filterFromMap(d.Get("filters").(map[string]interface{}))

	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	qp := ibclient.NewQueryParams(false, filters)
	res, err := objMgr.GetZoneDelegatedByFilters(qp)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to get zone delegated records: %w", err))
	}

	if res == nil {
		return diag.FromErr(fmt.Errorf("API returns a nil/empty ID for zone delegated"))
	}
	// TODO: temporary scaffold, need to rework marshalling/unmarshalling of EAs
	//       (avoiding additional layer of keys ("value" key)
	results := make([]interface{}, 0, len(res))
	for _, r := range res {
		zoneDelegatedFlat, err := flattenZoneDelegated(r)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to flatten zone delegated  : %w", err))
		}
		results = append(results, zoneDelegatedFlat)
	}

	err = d.Set("results", results)
	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags

}

func flattenZoneDelegated(zoneDelegated ibclient.ZoneDelegated) (map[string]interface{}, error) {
	var eaMap map[string]interface{}
	if zoneDelegated.Ea != nil && len(zoneDelegated.Ea) > 0 {
		eaMap = zoneDelegated.Ea
	} else {
		eaMap = make(map[string]interface{})
	}

	ea, err := json.Marshal(eaMap)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":          zoneDelegated.Ref,
		"fqdn":        zoneDelegated.Fqdn,
		"ext_attrs":   string(ea),
		"zone_format": zoneDelegated.ZoneFormat,
		"view":        *zoneDelegated.View,
	}
	if zoneDelegated.Comment != nil {
		res["comment"] = *zoneDelegated.Comment
	}
	if zoneDelegated.Disable != nil {
		res["disable"] = *zoneDelegated.Disable
	}
	if zoneDelegated.Locked != nil {
		res["locked"] = *zoneDelegated.Locked
	}
	if zoneDelegated.NsGroup != nil {
		res["ns_group"] = *zoneDelegated.NsGroup
	}
	if zoneDelegated.DelegatedTtl != nil {
		res["delegated_ttl"] = *zoneDelegated.DelegatedTtl
	}
	if zoneDelegated.DelegateTo.IsNull == false {
		nsInterface := convertNullableNameServersToInterface(zoneDelegated.DelegateTo)
		res["delegate_to"] = nsInterface
	}

	return res, nil
}
