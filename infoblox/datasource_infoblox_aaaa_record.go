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

func dataSourceAAAARecord() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAAAARecordRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:     schema.TypeMap,
				Required: true,
			},

			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of AAAA Records matching filters.",
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
						"fqdn": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the AAAA-record in FQDN format.",
						},
						"ipv6_addr": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "IPv6 address of the AAAA-record.",
						},
						"zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The zone which the record belongs to.",
						},
						"ttl": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "TTL attribute value for the record.",
						},
						"comment": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The content of 'comment' field at the AAAA-record.",
						},
						"ext_attrs": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Extensible attributes of the AAAA-record",
						},
					},
				},
			},
		},
	}
}

func dataSourceAAAARecordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(ibclient.IBConnector)

	var diags diag.Diagnostics

	n := &ibclient.RecordAAAA{}
	n.SetReturnFields(append(n.ReturnFields(), "extattrs", "zone", "comment", "ttl"))

	filters := filterFromMap(d.Get("filters").(map[string]interface{}))
	qp := ibclient.NewQueryParams(false, filters)
	var res []ibclient.RecordAAAA

	err := connector.GetObject(n, "", qp, &res)
	if err != nil {
		// Check if it's a "not found" error for data source - this is acceptable
		if _, ok := err.(*ibclient.NotFoundError); ok {
			// For data sources, empty results are valid - just return empty results
			res = []ibclient.RecordAAAA{}
		} else {
			return diag.FromErr(fmt.Errorf("Getting AAAA Record failed : %s", err.Error()))
		}
	}

	// TODO: temporary scaffold, need to rework marshalling/unmarshalling of EAs
	//       (avoiding additional layer of keys ("value" key)
	results := make([]interface{}, 0, len(res))
	for _, qa := range res {
		qarecordFlat, err := flattenRecordAAAA(qa)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to flatten AAAA Record: %w", err))
		}

		results = append(results, qarecordFlat)
	}

	err = d.Set("results", results)
	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenRecordAAAA(qarecord ibclient.RecordAAAA) (map[string]interface{}, error) {
	var eaMap map[string]interface{}
	if qarecord.Ea != nil && len(qarecord.Ea) > 0 {
		eaMap = qarecord.Ea
	} else {
		eaMap = make(map[string]interface{})
	}
	ea, err := json.Marshal(eaMap)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":        qarecord.Ref,
		"dns_view":  qarecord.View,
		"zone":      qarecord.Zone,
		"ext_attrs": string(ea),
	}

	if qarecord.Ipv6Addr != nil {
		res["ipv6_addr"] = *qarecord.Ipv6Addr
	}

	if qarecord.UseTtl != nil {
		if !*qarecord.UseTtl {
			res["ttl"] = ttlUndef
		}
	}

	if qarecord.Ttl != nil && *qarecord.Ttl > 0 {
		res["ttl"] = *qarecord.Ttl
	} else {
		res["ttl"] = ttlUndef
	}

	if qarecord.Name != nil {
		res["fqdn"] = *qarecord.Name
	}

	if qarecord.Comment != nil {
		res["comment"] = *qarecord.Comment
	}

	return res, nil
}
