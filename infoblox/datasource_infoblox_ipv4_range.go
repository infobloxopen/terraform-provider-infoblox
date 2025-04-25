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

func dataSourceRange() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRangeRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:     schema.TypeMap,
				Required: true,
			},
			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of Network Range matching filters",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"comment": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Comment for the range; maximum 256 characters.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the range.",
						},
						"network": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The network to which this range belongs, in IPv4 Address/CIDR format.",
						},
						"network_view": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the network view in which this range resides.",
						},
						"start_addr": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IPv4 Address starting address of the range.",
						},
						"end_addr": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IPv4 Address end address of the range.",
						},
						"disable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Determines whether a range is disabled or not. When this is set to False, the range is enabled.\n\n",
						},
						"ext_attrs": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Extensible attributes of the range to be added/updated, as a map in JSON format.",
						},
						"failover_association": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "TThe name of the failover association: the server in this failover association will serve the IPv4 range in case the main server is out of service.",
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
						"ms_server": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "The Microsoft server that will provide service for this range. `server_association_type` needs to be set to `MS_SERVER` +" +
								"if you want the server specified here to serve the range. For searching by this field you should use a HTTP method that contains a" +
								"body (POST or PUT) with MS DHCP server structure and the request should have option _method=GET.",
						},
						"member": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The member that will provide service for this range. server_association_type needs to be set to ‘MEMBER’ if you want" +
								"the server specified here to serve the range.",
						},
						"use_options": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Use flag for options.",
						},
						"server_association_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of server that is going to serve the range. The valid values are: 'FAILOVER', 'MEMBER', 'NONE'.'MS_FAILOVER','MS_SERVER'",
						},
						"cloud_info": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Structure containing all cloud API related information for this object.",
						},
					},
				},
			},
		},
	}
}
func dataSourceRangeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(ibclient.IBConnector)

	var diags diag.Diagnostics
	filters := filterFromMap(d.Get("filters").(map[string]interface{}))

	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")
	qp := ibclient.NewQueryParams(false, filters)
	res, err := objMgr.GetNetworkRange(qp)
	if err != nil {
		return diag.FromErr(err)
	}
	if res == nil {
		return diag.FromErr(fmt.Errorf("API returns a nil/empty ID for Network Range"))
	}
	results := make([]interface{}, 0, len(res))
	for _, r := range res {
		dtcPool, err := flattenNetworkRange(r)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to flatten Network Range: %w", err))
		}
		results = append(results, dtcPool)
	}

	err = d.Set("results", results)
	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenNetworkRange(networkRange ibclient.Range) (map[string]interface{}, error) {
	var eaMap map[string]interface{}
	if networkRange.Ea != nil && len(networkRange.Ea) > 0 {
		eaMap = networkRange.Ea
	} else {
		eaMap = make(map[string]interface{})
	}

	ea, err := json.Marshal(eaMap)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":                      networkRange.Ref,
		"ext_attrs":               string(ea),
		"server_association_type": networkRange.ServerAssociationType,
	}
	if networkRange.Network != nil {
		res["network"] = *networkRange.Network
	}
	if networkRange.NetworkView != nil {
		res["network_view"] = *networkRange.NetworkView
	}
	if networkRange.StartAddr != nil {
		res["start_addr"] = *networkRange.StartAddr
	}
	if networkRange.EndAddr != nil {
		res["end_addr"] = *networkRange.EndAddr

	}
	if networkRange.FailoverAssociation != nil {
		res["failover_association"] = *networkRange.FailoverAssociation
	}
	if networkRange.Member != nil {
		res["member"] = convertDhcpMemberToMap(networkRange.Member)
	}
	if networkRange.Disable != nil {
		res["disable"] = *networkRange.Disable
	}
	if networkRange.UseOptions != nil {
		res["use_options"] = *networkRange.UseOptions
	}
	if networkRange.Options != nil {
		res["options"] = convertDhcpOptionsToInterface(networkRange.Options)
	}
	if networkRange.Comment != nil {
		res["comment"] = *networkRange.Comment
	}
	if networkRange.Name != nil {
		res["name"] = *networkRange.Name
	}
	if networkRange.MsServer != nil {
		res["ms_server"] = networkRange.MsServer.Ipv4Addr
	}
	if networkRange.CloudInfo != nil {
		cloudInfo, err := serializeGridCloudApiInfo(networkRange.CloudInfo)
		if err != nil {
			return nil, err
		}
		res["cloud_info"] = cloudInfo
	}
	return res, nil
}
