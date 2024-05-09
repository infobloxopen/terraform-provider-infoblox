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

func dataSourceIpv6NetworkContainer() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIpv6NetworkContainerRead,
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
							Type:        schema.TypeString,
							Optional:    true,
							Default:     defaultNetView,
							Description: "Newtwork view's name the network container belongs to.",
						},
						"cidr": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The CIDR value of the network container.",
						},
						"comment": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Network container's description.",
						},
						"ext_attrs": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Extensible attributes for the network container.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIpv6NetworkContainerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(ibclient.IBConnector)

	n := &ibclient.Ipv6NetworkContainer{}
	n.SetReturnFields(append(n.ReturnFields(), "extattrs"))

	filters := filterFromMap(d.Get("filters").(map[string]interface{}))
	qp := ibclient.NewQueryParams(false, filters)
	var res []ibclient.Ipv6NetworkContainer

	err := connector.GetObject(n, "", qp, &res)
	if err != nil {
		return diag.FromErr(fmt.Errorf("getting NetworkContainer failed : %w", err))
	}

	// TODO: temporary scaffold, need to rework marshalling/unmarshalling of EAs
	//       (avoiding additional layer of keys ("value" key)
	results := make([]interface{}, 0, len(res))
	for _, nc := range res {
		networkContainerFlat, err := flattenIpv6NetworkContainer(nc)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to flatten network container: %w", err))
		}

		results = append(results, networkContainerFlat)
	}

	err = d.Set("results", results)
	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return nil
}

func flattenIpv6NetworkContainer(nc ibclient.Ipv6NetworkContainer) (map[string]interface{}, error) {
	var eaMap map[string]interface{}
	if nc.Ea != nil && len(nc.Ea) > 0 {
		eaMap = nc.Ea
	} else {
		eaMap = make(map[string]interface{})
	}
	ea, err := json.Marshal(eaMap)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":           nc.Ref,
		"network_view": nc.NetworkView,
		"cidr":         nc.Network,
		"ext_attrs":    string(ea),
	}

	if nc.Comment != nil {
		res["comment"] = *nc.Comment
	}

	return res, nil
}
