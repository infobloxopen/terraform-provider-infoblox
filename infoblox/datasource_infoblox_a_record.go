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

func dataSourceARecord() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceARecordRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:     schema.TypeMap,
				Required: true,
			},

			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of A records matching filters",
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
							Description: "A record FQDN.",
						},
						"ip_addr": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "IP address the A-record points to",
						},
						"zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The zone which the record belongs to.",
						},
						"ttl": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "TTL value for the A-record.",
						},
						"comment": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of the A-record.",
						},
						"ext_attrs": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Extensible attributes of the A-record, as a map in JSON format",
						},
					},
				},
			},
		},
	}
}

func dataSourceARecordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(ibclient.IBConnector)

	var diags diag.Diagnostics

	n := &ibclient.RecordA{}
	n.SetReturnFields(append(n.ReturnFields(), "extattrs", "comment", "zone", "ttl"))

	filters := filterFromMap(d.Get("filters").(map[string]interface{}))
	qp := ibclient.NewQueryParams(false, filters)
	var res []ibclient.RecordA

	err := connector.GetObject(n, "", qp, &res)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed getting A-record: %s", err.Error()))
	}

	if res == nil {
		return diag.FromErr(fmt.Errorf("API returns a nil/empty ID for the A Record"))
	}

	// TODO: temporary scaffold, need to rework marshalling/unmarshalling of EAs
	//       (avoiding additional layer of keys ("value" key)
	results := make([]interface{}, 0, len(res))
	for _, r := range res {
		recordaFlat, err := flattenRecordA(r)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to flatten A Record  : %w", err))
		}

		results = append(results, recordaFlat)
	}

	err = d.Set("results", results)
	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenRecordA(recorda ibclient.RecordA) (map[string]interface{}, error) {
	var eaMap map[string]interface{}
	if recorda.Ea != nil && len(recorda.Ea) > 0 {
		eaMap = recorda.Ea
	} else {
		eaMap = make(map[string]interface{})
	}

	ea, err := json.Marshal(eaMap)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":        recorda.Ref,
		"zone":      recorda.Zone,
		"dns_view":  recorda.View,
		"ext_attrs": string(ea),
	}

	if recorda.Ipv4Addr != nil {
		res["ip_addr"] = *recorda.Ipv4Addr
	}

	if recorda.UseTtl != nil {
		if !*recorda.UseTtl {
			res["ttl"] = ttlUndef
		}
	}

	if recorda.Ttl != nil && *recorda.Ttl > 0 {
		res["ttl"] = *recorda.Ttl
	} else {
		res["ttl"] = ttlUndef
	}

	if recorda.Name != nil {
		res["fqdn"] = *recorda.Name
	}

	if recorda.Comment != nil {
		res["comment"] = *recorda.Comment
	}

	return res, nil

}
