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

func dataSourceTXTRecord() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTXTRecordRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:     schema.TypeMap,
				Required: true,
			},

			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of TXT Records matching filters",
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
							Description: "FQDN for the TXT-Record.",
						},
						"text": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Data of the TXT-Record.",
						},
						"zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The zone which the record belongs to.",
						},
						"ttl": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "TTL value for the TXT-Record.",
						},
						"comment": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of the TXT-Record.",
						},
						"ext_attrs": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Extensible attributes of the TXT-record, as a map in JSON format.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTXTRecordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(ibclient.IBConnector)

	var diags diag.Diagnostics

	n := &ibclient.RecordTXT{}
	n.SetReturnFields(append(n.ReturnFields(), "extattrs", "zone", "comment", "ttl"))

	filters := filterFromMap(d.Get("filters").(map[string]interface{}))
	qp := ibclient.NewQueryParams(false, filters)
	var res []ibclient.RecordTXT

	err := connector.GetObject(n, "", qp, &res)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed getting TXT-Record: %s", err))
	}

	if res == nil {
		return diag.FromErr(fmt.Errorf("API returns a nil/empty ID for the TXT Records"))
	}

	// TODO: temporary scaffold, need to rework marshalling/unmarshalling of EAs
	//       (avoiding additional layer of keys ("value" key)
	results := make([]interface{}, 0, len(res))
	for _, tx := range res {
		recordtxtFlat, err := flattenRecordTXT(tx)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to flatten TXT Record: %w", err))
		}

		results = append(results, recordtxtFlat)
	}

	err = d.Set("results", results)
	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenRecordTXT(recordtxt ibclient.RecordTXT) (map[string]interface{}, error) {
	var eaMap map[string]interface{}
	if recordtxt.Ea != nil && len(recordtxt.Ea) > 0 {
		eaMap = (map[string]interface{})(recordtxt.Ea)
	} else {
		eaMap = make(map[string]interface{})
	}
	ea, err := json.Marshal(eaMap)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":        recordtxt.Ref,
		"zone":      recordtxt.Zone,
		"ext_attrs": string(ea),
	}

	if recordtxt.Text != nil {
		res["text"] = *recordtxt.Text
	}

	if recordtxt.UseTtl != nil {
		if !*recordtxt.UseTtl {
			res["ttl"] = ttlUndef
		}
	}

	if recordtxt.Ttl != nil && *recordtxt.Ttl > 0 {
		res["ttl"] = *recordtxt.Ttl
	} else {
		res["ttl"] = ttlUndef
	}

	if recordtxt.View != nil {
		res["dns_view"] = *recordtxt.View
	}

	if recordtxt.Name != nil {
		res["fqdn"] = *recordtxt.Name
	}

	if recordtxt.Comment != nil {
		res["comment"] = *recordtxt.Comment
	}

	return res, nil
}
