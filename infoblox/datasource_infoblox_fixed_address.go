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

func dataSourceFixedAddress() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFixedAddressRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:     schema.TypeMap,
				Required: true,
			},

			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of fixed address matching filters",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"agent_circuit_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The agent circuit ID for the fixed address.",
						},
						"agent_remote_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The agent remote ID for the fixed address.",
						},
						"client_identifier_prepend_zero": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "This field controls whether there is a prepend for the dhcp-client-identifier of a fixed address.",
						},
						"comment": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Comment for the fixed address; maximum 256 characters.",
						},
						"dhcp_client_identifier": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The DHCP client ID for the fixed address. The field is required only when match_client is set to CLIENT_ID.",
						},
						"disable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Determines whether a fixed address is disabled or not. When this is set to False, the fixed address is enabled.",
						},
						"ext_attrs": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Extensible attributes of the A-record to be added/updated, as a map in JSON format",
						},
						"ipv4addr": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IPv4 Address of the fixed address.",
						},
						"mac": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The MAC address value for this fixed address.",
						},
						"match_client": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The match client for the fixed address.Valid values are CIRCUIT_ID, CLIENT_ID , MAC_ADDRESS, REMOTE_ID and RESERVED",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "This field contains the name of this fixed address.",
						},
						"network": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The network to which this fixed address belongs, in IPv4 Address/CIDR format.",
						},
						"network_view": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the network view in which this fixed address resides.",
						},
						"options": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "An array of DHCP option structs that lists the DHCP options associated with the object.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the DHCP option.",
									},
									"num": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The code of the DHCP option.",
									},
									"use_option": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Only applies to special options that are displayed separately from other options and have a use flag.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Value of the DHCP option",
									},
									"vendor_class": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the space this DHCP option is associated to.",
									},
								},
							},
						},
						"use_options": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Use option is a flag that indicates whether the options field are used or not.",
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

func dataSourceFixedAddressRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(ibclient.IBConnector)

	var diags diag.Diagnostics

	filters := filterFromMap(d.Get("filters").(map[string]interface{}))
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	qp := ibclient.NewQueryParams(false, filters)
	res, err := objMgr.GetAllFixedAddress(qp, false)
	if err != nil {
		return diag.FromErr(err)
	}
	if res == nil {
		return diag.FromErr(fmt.Errorf("API returns a nil/empty ID for Fixed address "))
	}
	results := make([]interface{}, 0, len(res))
	for _, r := range res {
		fixedAddress, err := flattenFixedAddress(r)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to flatten Fixed address : %w", err))
		}
		results = append(results, fixedAddress)
	}

	err = d.Set("results", results)
	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenFixedAddress(fixedAddress ibclient.FixedAddress) (map[string]interface{}, error) {
	var eaMap map[string]interface{}
	if fixedAddress.Ea != nil && len(fixedAddress.Ea) > 0 {
		eaMap = fixedAddress.Ea
	} else {
		eaMap = make(map[string]interface{})
	}

	ea, err := json.Marshal(eaMap)
	if err != nil {
		return nil, err
	}
	res := map[string]interface{}{
		"id":                             fixedAddress.Ref,
		"network_view":                   fixedAddress.NetviewName,
		"network":                        fixedAddress.Cidr,
		"comment":                        fixedAddress.Comment,
		"ipv4addr":                       fixedAddress.IPv4Address,
		"mac":                            fixedAddress.Mac,
		"match_client":                   fixedAddress.MatchClient,
		"agent_circuit_id":               fixedAddress.AgentCircuitId,
		"agent_remote_id":                fixedAddress.AgentRemoteId,
		"client_identifier_prepend_zero": fixedAddress.ClientIdentifierPrependZero,
		"use_options":                    fixedAddress.UseOptions,
		"disable":                        fixedAddress.Disable,
		"dhcp_client_identifier":         fixedAddress.DhcpClientIdentifier,
		"ext_attrs":                      string(ea),
		"name":                           fixedAddress.Name,
	}
	if fixedAddress.Options != nil {
		res["options"] = convertDhcpOptionsToInterface(fixedAddress.Options)
	}
	if fixedAddress.CloudInfo != nil {
		cloudInfo, err := serializeGridCloudApiInfo(fixedAddress.CloudInfo)
		if err != nil {
			return nil, err
		}
		res["cloud_info"] = cloudInfo
	}
	return res, nil
}
