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

func dataSourceIPv4Network() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIPv4NetworkRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:     schema.TypeMap,
				Required: true,
			},

			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of networks matching filters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_view": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  defaultNetView,
						},
						"cidr": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"comment": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A string describing the network",
						},
						"ext_attrs": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Extensible attributes for network datasource, as a map in JSON format",
						},
					},
				},
			},
		},
	}
}

func dataSourceIPv4NetworkRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(ibclient.IBConnector)

	var diags diag.Diagnostics

	n := &ibclient.Ipv4Network{}
	n.SetReturnFields(append(n.ReturnFields(), "extattrs"))

	filters := filterFromMap(d.Get("filters").(map[string]interface{}))
	qp := ibclient.NewQueryParams(false, filters)
	var res []ibclient.Ipv4Network

	err := connector.GetObject(n, "", qp, &res)
	if err != nil {
		return diag.FromErr(fmt.Errorf("getting network failed: %s", err))
	}
	if res == nil {
		return diag.FromErr(fmt.Errorf("API returns a nil/empty ID for the network"))
	}

	// TODO: temporary scaffold, need to rework marshalling/unmarshalling of EAs
	//       (avoiding additional layer of keys ("value" key)
	results := make([]interface{}, 0, len(res))
	for _, n := range res {
		networkFlat, err := flattenNetwork(n)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to flatten network: %w", err))
		}

		results = append(results, networkFlat)
	}

	err = d.Set("results", results)
	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenNetwork(network ibclient.Ipv4Network) (map[string]interface{}, error) {
	var eaMap map[string]interface{}
	if network.Ea != nil && len(network.Ea) > 0 {
		eaMap = (map[string]interface{})(network.Ea)
	} else {
		eaMap = make(map[string]interface{})
	}
	ea, err := json.Marshal(eaMap)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":           network.Ref,
		"network_view": network.NetworkView,
		"ext_attrs":    string(ea),
	}

	if network.Network != nil {
		res["cidr"] = *network.Network
	}

	if network.Comment != nil {
		res["comment"] = *network.Comment
	}

	return res, nil
}
