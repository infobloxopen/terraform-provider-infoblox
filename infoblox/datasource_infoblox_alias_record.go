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

func dataSourceAliasRecord() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceAliasRecordRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:     schema.TypeMap,
				Required: true,
			},
			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of Alias records matching filters",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the alias record.",
						},
						"comment": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Comment for the alias record.",
						},
						"disable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "A boolean flag which indicates if the alias record is disabled.",
						},
						"target_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Target name in FQDN format.",
						},
						"target_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of the target object.",
						},
						"dns_view": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the DNS view in which the alias record is created.",
						},
						"ttl": {
							Type:     schema.TypeInt,
							Computed: true,
							Description: "The Time To Live (TTL) value for record. A 32-bit unsigned integer that represents the duration, " +
								"in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached.",
						},
						"ext_attrs": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Extensible attributes of the  Alias Record to be added/updated, as a map in JSON format",
						},
						"creator": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creator of the Alias Record. Valid value is STATIC",
						},
						"dns_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for an Alias record in punycode format.",
						},
						"dns_target_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "DNS target name of the Alias Record in punycode format.",
						},
						"zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the zone in which the record resides. Example: “zone.com”.",
						},
						"cloud_info": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Structure containing all cloud API related information for this object.",
						},
					},
				},
			},
		},
	}
}

func datasourceAliasRecordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(ibclient.IBConnector)

	var diags diag.Diagnostics

	filters := filterFromMap(d.Get("filters").(map[string]interface{}))

	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	qp := ibclient.NewQueryParams(false, filters)
	res, err := objMgr.GetAllAliasRecord(qp)
	if err != nil {
		return diag.FromErr(err)
	}

	if res == nil {
		return diag.FromErr(fmt.Errorf("API returns a nil/empty ID for alias record"))
	}
	// TODO: temporary scaffold, need to rework marshalling/unmarshalling of EAs
	//       (avoiding additional layer of keys ("value" key)
	results := make([]interface{}, 0, len(res))
	for _, r := range res {
		aliasRecord, err := flattenAliasRecord(r)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to flatten alias record  : %w", err))
		}
		results = append(results, aliasRecord)
	}

	err = d.Set("results", results)
	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenAliasRecord(aliasRecord ibclient.RecordAlias) (interface{}, error) {
	var eaMap map[string]interface{}
	if aliasRecord.Ea != nil && len(aliasRecord.Ea) > 0 {
		eaMap = aliasRecord.Ea
	} else {
		eaMap = make(map[string]interface{})
	}

	ea, err := json.Marshal(eaMap)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":              aliasRecord.Ref,
		"name":            aliasRecord.Name,
		"ext_attrs":       string(ea),
		"target_name":     aliasRecord.TargetName,
		"target_type":     aliasRecord.TargetType,
		"dns_view":        *aliasRecord.View,
		"dns_name":        aliasRecord.DnsName,
		"dns_target_name": aliasRecord.DnsTargetName,
		"creator":         aliasRecord.Creator,
		"zone":            aliasRecord.Zone,
	}
	if aliasRecord.Comment != nil {
		res["comment"] = *aliasRecord.Comment
	}
	if aliasRecord.Disable != nil {
		res["disable"] = *aliasRecord.Disable
	}
	if aliasRecord.UseTtl != nil {
		if !*aliasRecord.UseTtl {
			res["ttl"] = ttlUndef
		}
	}
	if aliasRecord.Ttl != nil && *aliasRecord.Ttl > 0 {
		res["ttl"] = *aliasRecord.Ttl
	} else {
		res["ttl"] = ttlUndef
	}
	if aliasRecord.CloudInfo != nil {
		cloudInfo, err := serializeGridCloudApiInfo(aliasRecord.CloudInfo)
		if err != nil {
			return nil, err
		}
		res["cloud_info"] = cloudInfo
	}

	return res, nil
}
