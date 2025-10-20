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

func dataSourceNetwork() *schema.Resource {
	return &schema.Resource{
		//ReadContext: dataSourceIPv4NetworkRead,
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
						"utilization": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The utilization of the network",
						},
						"options": {
							Type:     schema.TypeList,
							Computed: true,
							Description: "An array of DHCP option structs that lists the DHCP options associated with the object. An option sets the" +
								"value of a DHCP option that has been defined in an option space. DHCP options describe network configuration settings" +
								"and various services available on the network. These options occur as variable-length fields at the end of DHCP messages." +
								"When defining a DHCP option, at least a ‘name’ or a ‘num’ is required.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Name of the DHCP option.",
									},
									"num": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The code of the DHCP option.",
									},
									"use_option": {
										Type:     schema.TypeBool,
										Computed: true,
										Description: "Only applies to special options that are displayed separately from other options and have a use flag. " +
											"These options are: `routers`, `router-templates`, `domain-name-servers`, `domain-name`, `broadcast-address`, " +
											"`broadcast-address-offset`, `dhcp-lease-time`, `dhcp6.name-servers`",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Value of the DHCP option.",
									},
									"vendor_class": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the space this DHCP option is associated to.",
									},
								},
							},
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
	n.SetReturnFields(append(n.ReturnFields(), "extattrs", "options", "utilization"))

	filters := filterFromMap(d.Get("filters").(map[string]interface{}))
	qp := ibclient.NewQueryParams(false, filters)
	var res []ibclient.Ipv4Network

	err := connector.GetObject(n, "", qp, &res)
	if err != nil {
		// Check if it's a "not found" error for data source - this is acceptable
		if _, ok := err.(*ibclient.NotFoundError); ok {
			// For data sources, empty results are valid - just return empty results
			res = []ibclient.Ipv4Network{}
		} else {
			return diag.FromErr(fmt.Errorf("Getting IPv4 network failed with filters %v: %s", filters, err.Error()))
		}
	}

	// TODO: temporary scaffold, need to rework marshalling/unmarshalling of EAs
	//       (avoiding additional layer of keys ("value" key)
	results := make([]interface{}, 0, len(res))
	for _, n := range res {
		networkFlat, err := flattenIpv4Network(n)
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

func flattenIpv4Network(network ibclient.Ipv4Network) (map[string]interface{}, error) {
	var eaMap map[string]interface{}
	if network.Ea != nil && len(network.Ea) > 0 {
		eaMap = network.Ea
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
		"utilization":  network.Utilization,
	}

	if network.Network != nil {
		res["cidr"] = *network.Network
	}

	if network.Comment != nil {
		res["comment"] = *network.Comment
	}

	if network.Options != nil {
		res["options"] = convertDhcpOptionsToInterface(network.Options)
	}
	return res, nil
}

func flattenIpv6Network(network ibclient.Ipv6Network) (map[string]interface{}, error) {
	var eaMap map[string]interface{}
	if network.Ea != nil && len(network.Ea) > 0 {
		eaMap = network.Ea
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

	if network.Options != nil {
		res["options"] = convertDhcpOptionsToInterface(network.Options)
	}
	return res, nil
}

func dataSourceIPv6NetworkRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(ibclient.IBConnector)

	var diags diag.Diagnostics

	n := &ibclient.Ipv6Network{}
	n.SetReturnFields(append(n.ReturnFields(), "extattrs", "options"))

	filters := filterFromMap(d.Get("filters").(map[string]interface{}))
	qp := ibclient.NewQueryParams(false, filters)
	var res []ibclient.Ipv6Network

	err := connector.GetObject(n, "", qp, &res)
	if err != nil {
		// Check if it's a "not found" error for data source - this is acceptable
		if _, ok := err.(*ibclient.NotFoundError); ok {
			// For data sources, empty results are valid - just return empty results
			res = []ibclient.Ipv6Network{}
		} else {
			return diag.FromErr(fmt.Errorf("Getting IPv6 network failed with filters %v: %s", filters, err.Error()))
		}
	}

	// TODO: temporary scaffold, need to rework marshalling/unmarshalling of EAs
	//       (avoiding additional layer of keys ("value" key)
	results := make([]interface{}, 0, len(res))
	for _, n := range res {
		networkFlat, err := flattenIpv6Network(n)
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

func dataSourceIPv4Network() *schema.Resource {
	nw := dataSourceNetwork()
	nw.ReadContext = dataSourceIPv4NetworkRead
	return nw
}

func dataSourceIPv6Network() *schema.Resource {
	nw := dataSourceNetwork()
	nw.ReadContext = dataSourceIPv6NetworkRead
	return nw
}
