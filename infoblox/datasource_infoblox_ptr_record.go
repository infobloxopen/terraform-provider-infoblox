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

func dataSourcePtrRecord() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePtrRecordRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:     schema.TypeMap,
				Required: true,
			},

			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of PTR Records matching filters",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dns_view": {
							Type:        schema.TypeString,
							Default:     defaultDNSView,
							Optional:    true,
							Description: "DNS view which the record's zone belongs to.",
						},
						"ptrdname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The domain name in FQDN which the record points to.",
						},
						"record_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "The name of the DNS PTR-record in FQDN format",
						},
						"ip_addr": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "IPv4/IPv6 address the PTR-record points from.",
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
							Description: "PTR-record's description.",
						},
						"ext_attrs": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Extensible attributes of the PTR-record.",
						},
					},
				},
			},
		},
	}
}

func dataSourcePtrRecordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(ibclient.IBConnector)

	var diags diag.Diagnostics

	n := &ibclient.RecordPTR{}
	n.SetReturnFields(append(n.ReturnFields(), "extattrs", "zone", "comment", "name", "ipv4addr", "ipv6addr", "ttl"))

	filters := filterFromMap(d.Get("filters").(map[string]interface{}))
	qp := ibclient.NewQueryParams(false, filters)
	var res []ibclient.RecordPTR

	err := connector.GetObject(n, "", qp, &res)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed getting PTR-record: %s", err.Error()))
	}
	if res == nil {
		return diag.FromErr(fmt.Errorf("API returns a nil/empty ID for the PTR Record"))
	}

	// TODO: temporary scaffold, need to rework marshalling/unmarshalling of EAs
	//       (avoiding additional layer of keys ("value" key)
	results := make([]interface{}, 0, len(res))
	for _, pt := range res {
		recordptrFlat, err := flattenRecordPTR(pt)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to flatten PTR Record: %w", err))
		}

		results = append(results, recordptrFlat)
	}

	err = d.Set("results", results)
	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenRecordPTR(recordptr ibclient.RecordPTR) (map[string]interface{}, error) {
	var eaMap map[string]interface{}
	if recordptr.Ea != nil && len(recordptr.Ea) > 0 {
		eaMap = recordptr.Ea
	} else {
		eaMap = make(map[string]interface{})
	}

	ea, err := json.Marshal(eaMap)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":        recordptr.Ref,
		"dns_view":  recordptr.View,
		"zone":      recordptr.Zone,
		"ext_attrs": string(ea),
	}

	if recordptr.Ipv4Addr != nil || recordptr.Ipv6Addr != nil {
		if *recordptr.Ipv4Addr != "" {
			res["ip_addr"] = *recordptr.Ipv4Addr
		} else {
			res["ip_addr"] = *recordptr.Ipv6Addr
		}
	}

	if recordptr.UseTtl != nil {
		if !*recordptr.UseTtl {
			res["ttl"] = ttlUndef
		}
	}

	if recordptr.Ttl != nil && *recordptr.Ttl > 0 {
		res["ttl"] = *recordptr.Ttl
	} else {
		res["ttl"] = ttlUndef
	}

	if recordptr.PtrdName != nil {
		res["ptrdname"] = *recordptr.PtrdName
	}

	if recordptr.Name != nil {
		res["record_name"] = *recordptr.Name
	}

	if recordptr.Comment != nil {
		res["comment"] = *recordptr.Comment
	}

	return res, nil
}
