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

func dataSourceDNSView() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDNSViewRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:     schema.TypeMap,
				Required: true,
			},

			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of DNS View matching filters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of th DNS View.",
						},
						"network_view": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of Network View in which DNS View exists.",
						},
						"comment": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Description of the DNS View.",
						},
						"ext_attrs": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The Extensible attributes of the DNS view to be added/updated, as a map in JSON format",
						},
					},
				},
			},
		},
	}
}

func dataSourceDNSViewRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(ibclient.IBConnector)

	var diags diag.Diagnostics

	dv := &ibclient.View{}
	dv.SetReturnFields(append(dv.ReturnFields(), "extattrs", "network_view"))

	filters := filterFromMap(d.Get("filters").(map[string]interface{}))
	qp := ibclient.NewQueryParams(false, filters)
	var res []ibclient.View

	err := connector.GetObject(dv, "", qp, &res)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed getting DNS View: %s", err.Error()))
	}

	if res == nil {
		return diag.FromErr(fmt.Errorf("API returns a nil/empty ID for the DNS View"))
	}

	// TODO: temporary scaffold, need to rework marshalling/unmarshalling of EAs
	//       (avoiding additional layer of keys ("value" key)
	results := make([]interface{}, 0, len(res))
	for _, r := range res {
		dnsviewFlat, err := flattenDNSView(r)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to flatten DNS View  : %w", err))
		}

		results = append(results, dnsviewFlat)
	}

	err = d.Set("results", results)
	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenDNSView(dnsview ibclient.View) (map[string]interface{}, error) {
	var eaMap map[string]interface{}
	if dnsview.Ea != nil && len(dnsview.Ea) > 0 {
		eaMap = dnsview.Ea
	} else {
		eaMap = make(map[string]interface{})
	}

	ea, err := json.Marshal(eaMap)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":        dnsview.Ref,
		"ext_attrs": string(ea),
	}

	if dnsview.Name != nil {
		res["name"] = *dnsview.Name
	}

	if dnsview.NetworkView != nil {
		res["network_view"] = *dnsview.NetworkView
	}

	if dnsview.Comment != nil {
		res["comment"] = *dnsview.Comment
	}

	return res, nil
}
