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

func dataSourceNetworkView() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNetworkViewRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:     schema.TypeMap,
				Required: true,
			},

			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of network views matching filters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the view shown in NIOS's UI.",
						},
						"comment": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Network view's description",
						},
						"ext_attrs": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Extensible attributes of the network view.",
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkViewRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	connector := m.(ibclient.IBConnector)

	n := &ibclient.NetworkView{}
	n.SetReturnFields(append(n.ReturnFields(), "extattrs"))
	filters := filterFromMap(d.Get("filters").(map[string]interface{}))
	qp := ibclient.NewQueryParams(false, filters)
	var res []ibclient.NetworkView

	err := connector.GetObject(n, "", qp, &res)
	if err != nil {
		// Check if it's a "not found" error for data source - this is acceptable
		if _, ok := err.(*ibclient.NotFoundError); ok {
			// For data sources, empty results are valid - just return empty results
			res = []ibclient.NetworkView{}
		} else {
			return diag.FromErr(fmt.Errorf("Getting network view failed with filters %v: %s", filters, err.Error()))
		}
	}

	// TODO: temporary scaffold, need to rework marshalling/unmarshalling of EAs
	//       (avoiding additional layer of keys ("value" key)
	results := make([]interface{}, 0, len(res))
	for _, n := range res {
		networkViewFlat, err := flattenNetworkView(n)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to flatten network view: %w", err))
		}

		results = append(results, networkViewFlat)
	}

	err = d.Set("results", results)
	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenNetworkView(nv ibclient.NetworkView) (map[string]interface{}, error) {
	var eaMap map[string]interface{}
	if nv.Ea != nil && len(nv.Ea) > 0 {
		eaMap = nv.Ea
	} else {
		eaMap = make(map[string]interface{})
	}
	ea, err := json.Marshal(eaMap)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":        nv.Ref,
		"ext_attrs": string(ea),
	}

	if nv.Name != nil {
		res["name"] = *nv.Name
	}

	if nv.Comment != nil {
		res["comment"] = *nv.Comment
	}

	return res, nil
}
