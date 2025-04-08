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

func dataSourceIpv4SharedNetwork() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIpv4SharedNetworkRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:     schema.TypeMap,
				Required: true,
			},

			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of AAAA Records matching filters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the IPv4 shared network object.",
						},
						"comment": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The comment for the IPv4 shared network object.",
						},
						"ext_attrs": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Extensible attributes of the IPv4 Shared Network record to be added/updated, as a map in JSON format.",
						},
						"disable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The disable flag for the IPv4 shared network object.",
						},
						"networks": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A list of networks belonging to the shared network",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"network_view": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the network view in which this shared network resides.",
						},
						"use_options": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Use flag for options.",
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
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIpv4SharedNetworkRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(ibclient.IBConnector)
	var diags diag.Diagnostics
	filters := filterFromMap(d.Get("filters").(map[string]interface{}))
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	qp := ibclient.NewQueryParams(false, filters)
	res, err := objMgr.GetAllIpv4SharedNetwork(qp)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to get shared network records: %w", err))
	}

	if res == nil {
		return diag.FromErr(fmt.Errorf("API returns a nil/empty ID for shared network"))
	}
	// TODO: temporary scaffold, need to rework marshalling/unmarshalling of EAs
	//       (avoiding additional layer of keys ("value" key)
	results := make([]interface{}, 0, len(res))
	for _, r := range res {
		record, err := flattenSharedNetwork(r)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to flatten shared network: %w", err))
		}
		results = append(results, record)
	}

	err = d.Set("results", results)
	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenSharedNetwork(sharedNetwork ibclient.SharedNetwork) (interface{}, error) {
	var eaMap map[string]interface{}
	if sharedNetwork.Ea != nil && len(sharedNetwork.Ea) > 0 {
		eaMap = sharedNetwork.Ea
	} else {
		eaMap = make(map[string]interface{})
	}

	ea, err := json.Marshal(eaMap)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":           sharedNetwork.Ref,
		"name":         *sharedNetwork.Name,
		"ext_attrs":    string(ea),
		"network_view": sharedNetwork.NetworkView,
	}
	if sharedNetwork.Comment != nil {
		res["comment"] = *sharedNetwork.Comment
	}
	if sharedNetwork.Disable != nil {
		res["disable"] = *sharedNetwork.Disable
	}
	if sharedNetwork.UseOptions != nil {
		res["use_options"] = *sharedNetwork.UseOptions
	}
	if sharedNetwork.Options != nil {
		networksInterface := convertDhcpOptionsToInterface(sharedNetwork.Options)
		res["options"] = networksInterface
	}
	if sharedNetwork.Networks != nil {
		networks := setNetworksRef(sharedNetwork.Networks)
		res["networks"] = networks
	}
	return res, nil
}
