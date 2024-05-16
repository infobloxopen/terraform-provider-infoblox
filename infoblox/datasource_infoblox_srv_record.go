package infoblox

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func dataSourceSRVRecord() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSRVRecordRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:     schema.TypeMap,
				Required: true,
			},

			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of SRV Records matching filters",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dns_view": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     defaultDNSView,
							Description: "DNS view which the record's zone belongs to.",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Combination of service name, protocol name and zone name.",
						},
						"priority": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Configures the priority (0-65535) for this SRV record.",
						},
						"weight": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Configures weight of the SRV record, valid values (0-65535).",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Configures port (0-65535) for this SRV record.",
						},
						"target": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Provides service for domain name in the SRV record.",
						},
						"zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The zone which the record belongs to.",
						},
						"ttl": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "TTL value for the SRV record.",
						},
						"comment": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of the SRV record.",
						},
						"ext_attrs": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Extensible attributes of the SRV-record to be added/updated, as a map in JSON format.",
						},
					},
				},
			},
		},
	}
}

func dataSourceSRVRecordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(ibclient.IBConnector)

	var diags diag.Diagnostics

	n := &ibclient.RecordSRV{}
	n.SetReturnFields(append(n.ReturnFields(), "extattrs", "zone", "comment", "ttl"))

	filters := filterFromMap(d.Get("filters").(map[string]interface{}))
	qp := ibclient.NewQueryParams(false, filters)
	var res []ibclient.RecordSRV

	err := connector.GetObject(n, "", qp, &res)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed getting SRV-Record: %s", err))
	}
	if res == nil {
		return diag.FromErr(fmt.Errorf("API returns a nil/empty ID for the SRV Record"))
	}

	// TODO: temporary scaffold, need to rework marshalling/unmarshalling of EAs
	//       (avoiding additional layer of keys ("value" key)
	results := make([]interface{}, 0, len(res))
	for _, sv := range res {
		recordsrvFlat, err := flattenRecordSRV(sv)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to flatten SRV Record: %w", err))
		}

		results = append(results, recordsrvFlat)
	}

	err = d.Set("results", results)
	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenRecordSRV(recordsrv ibclient.RecordSRV) (map[string]interface{}, error) {
	var eaMap map[string]interface{}
	if recordsrv.Ea != nil && len(recordsrv.Ea) > 0 {
		eaMap = recordsrv.Ea
	} else {
		eaMap = make(map[string]interface{})
	}
	ea, err := json.Marshal(eaMap)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":        recordsrv.Ref,
		"dns_view":  recordsrv.View,
		"zone":      recordsrv.Zone,
		"ext_attrs": string(ea),
	}

	if recordsrv.Port != nil {
		tempPortVal := int(*recordsrv.Port)
		if err := ibclient.CheckIntRange("port", tempPortVal, 0, 65535); err != nil {
			return nil, err
		}
		res["port"] = *recordsrv.Port
	}

	if recordsrv.UseTtl != nil {
		if !*recordsrv.UseTtl {
			res["ttl"] = ttlUndef
		}
	}

	if recordsrv.Ttl != nil && *recordsrv.Ttl > 0 {
		res["ttl"] = *recordsrv.Ttl
	} else {
		res["ttl"] = ttlUndef
	}

	if recordsrv.Target != nil {
		res["target"] = *recordsrv.Target
	}

	if recordsrv.Weight != nil {
		res["weight"] = *recordsrv.Weight
	}

	if recordsrv.Priority != nil {
		res["priority"] = *recordsrv.Priority
	}

	if recordsrv.Name != nil {
		res["name"] = *recordsrv.Name
	}

	if recordsrv.Comment != nil {
		res["comment"] = *recordsrv.Comment
	}

	return res, nil
}
