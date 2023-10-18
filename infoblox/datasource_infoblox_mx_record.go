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

func dataSourceMXRecord() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMXRecordRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:     schema.TypeMap,
				Required: true,
			},

			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of MX Records matching filters",
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
							Description: "FQDN for the MX-Record.",
						},
						"mail_exchanger": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "A record used to specify mail server.",
						},
						"preference": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Configures the preference (0-65535) for this MX record.",
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

func dataSourceMXRecordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(ibclient.IBConnector)

	var diags diag.Diagnostics

	n := &ibclient.RecordMX{}
	n.SetReturnFields(append(n.ReturnFields(), "extattrs", "zone", "ttl", "comment"))

	filters := filterFromMap(d.Get("filters").(map[string]interface{}))
	qp := ibclient.NewQueryParams(false, filters)
	var res []ibclient.RecordMX

	err := connector.GetObject(n, "", qp, &res)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed getting MX-Record: %s", err))
	}

	if res == nil {
		return diag.FromErr(fmt.Errorf("API returns a nil/empty ID for MX Record"))
	}

	// TODO: temporary scaffold, need to rework marshalling/unmarshalling of EAs
	//       (avoiding additional layer of keys ("value" key)
	results := make([]interface{}, 0, len(res))
	for _, mr := range res {
		recordmxFlat, err := flattenRecordMX(mr)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to flatten MX Record: %w", err))
		}

		results = append(results, recordmxFlat)
	}

	err = d.Set("results", results)
	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenRecordMX(recordmx ibclient.RecordMX) (map[string]interface{}, error) {
	var eaMap map[string]interface{}
	if recordmx.Ea != nil && len(recordmx.Ea) > 0 {
		eaMap = recordmx.Ea
	} else {
		eaMap = make(map[string]interface{})
	}

	ea, err := json.Marshal(eaMap)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":        recordmx.Ref,
		"zone":      recordmx.Zone,
		"ext_attrs": string(ea),
	}

	if recordmx.Preference != nil {
		tempInt := int(*recordmx.Preference)
		if err := ibclient.CheckIntRange("preference", tempInt, 0, 65535); err != nil {
			return nil, err
		}
		res["preference"] = *recordmx.Preference
	}

	if recordmx.View != nil {
		res["dns_view"] = *recordmx.View
	}

	if recordmx.MailExchanger != nil {
		res["mail_exchanger"] = *recordmx.MailExchanger
	}

	if recordmx.UseTtl != nil {
		if !*recordmx.UseTtl {
			res["ttl"] = ttlUndef
		}
	}

	if recordmx.Ttl != nil && *recordmx.Ttl > 0 {
		res["ttl"] = *recordmx.Ttl
	} else {
		res["ttl"] = ttlUndef
	}

	if recordmx.Name != nil {
		res["fqdn"] = *recordmx.Name
	}

	if recordmx.Comment != nil {
		res["comment"] = *recordmx.Comment
	}

	return res, nil
}
