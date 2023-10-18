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

func dataSourceCNameRecord() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCNameRecordRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:     schema.TypeMap,
				Required: true,
			},

			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of CNAME Records matching filters",
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
						"canonical": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The Canonical name in FQDN format.",
						},
						"alias": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The alias name in FQDN format.",
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
							Description: "A description about CNAME record.",
						},
						"ext_attrs": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Extensible attributes of CNAME record, as a map in JSON format",
						},
					},
				},
			},
		},
	}
}

func dataSourceCNameRecordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(ibclient.IBConnector)

	var diags diag.Diagnostics

	n := &ibclient.RecordCNAME{}
	n.SetReturnFields(append(n.ReturnFields(), "extattrs", "zone", "comment", "ttl"))

	filters := filterFromMap(d.Get("filters").(map[string]interface{}))
	qp := ibclient.NewQueryParams(false, filters)
	var res []ibclient.RecordCNAME

	err := connector.GetObject(n, "", qp, &res)
	if err != nil {
		return diag.FromErr(fmt.Errorf("Getting CNAME Record failed : %s", err.Error()))
	}

	if res == nil {
		return diag.FromErr(fmt.Errorf("API returns a nil/empty ID for the CNAME Record"))
	}

	// TODO: temporary scaffold, need to rework marshalling/unmarshalling of EAs
	//       (avoiding additional layer of keys ("value" key)
	results := make([]interface{}, 0, len(res))
	for _, cn := range res {
		recordcnameFlat, err := flattenRecordCNAME(cn)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to flatten CNAME Record : %w", err))
		}

		results = append(results, recordcnameFlat)
	}

	err = d.Set("results", results)
	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenRecordCNAME(recordcname ibclient.RecordCNAME) (map[string]interface{}, error) {
	var eaMap map[string]interface{}
	if recordcname.Ea != nil && len(recordcname.Ea) > 0 {
		eaMap = recordcname.Ea
	} else {
		eaMap = make(map[string]interface{})
	}

	ea, err := json.Marshal(eaMap)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":        recordcname.Ref,
		"zone":      recordcname.Zone,
		"ext_attrs": string(ea),
	}

	if recordcname.UseTtl != nil {
		if !*recordcname.UseTtl {
			res["ttl"] = ttlUndef
		}
	}

	if recordcname.Ttl != nil && *recordcname.Ttl > 0 {
		res["ttl"] = *recordcname.Ttl
	} else {
		res["ttl"] = ttlUndef
	}

	if recordcname.View != nil {
		res["dns_view"] = *recordcname.View
	}

	if recordcname.Canonical != nil {
		res["canonical"] = *recordcname.Canonical
	}

	if recordcname.Name != nil {
		res["alias"] = *recordcname.Name
	}

	if recordcname.Comment != nil {
		res["comment"] = *recordcname.Comment
	}

	return res, nil
}
