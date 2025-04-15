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

func dataSourceRangeTemplate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRangeTemplateRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:     schema.TypeMap,
				Required: true,
			},
			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of PTR Records matching filters",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the Range Template record.",
						},
						"comment": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Comment for the Range Template record.",
						},
						"number_of_addresses": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of addresses for this range.",
						},
						"offset": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The start address offset for the range.",
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
						"server_association_type": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "The type of server that is going to serve the range. Valid values are: `FAILOVER`, `MEMBER`, `MS_FAILOVER`, " +
								"`MS_SERVER`, `NONE`. Default value is `NONE`",
						},
						"failover_association": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "The name of the failover association: the server in this failover association will serve the IPv4 range in case the " +
								"main server is out of service. `server_association_type` must be set to ‘FAILOVER’ or ‘FAILOVER_MS’ if you want the " +
								"failover association specified here to serve the range.",
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
						"ext_attrs": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Extensible attributes of the  Range Template Record to be added/updated, as a map in JSON format",
						},
					},
				},
			},
		},
	}
}

func dataSourceRangeTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(ibclient.IBConnector)

	var diags diag.Diagnostics

	filters := filterFromMap(d.Get("filters").(map[string]interface{}))

	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	qp := ibclient.NewQueryParams(false, filters)
	res, err := objMgr.GetAllRangeTemplate(qp)
	if err != nil {
		return diag.FromErr(err)
	}

	if res == nil {
		return diag.FromErr(fmt.Errorf("API returns a nil/empty ID for Range Template"))
	}
	results := make([]interface{}, 0, len(res))
	for _, r := range res {
		rangeTemplate, err := flattenRangeTemplate(r, connector)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to flatten Range Template : %w", err))
		}
		results = append(results, rangeTemplate)
	}

	err = d.Set("results", results)
	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenRangeTemplate(rangeTemplate ibclient.Rangetemplate, connector ibclient.IBConnector) (map[string]interface{}, error) {
	var eaMap map[string]interface{}
	if rangeTemplate.Ea != nil && len(rangeTemplate.Ea) > 0 {
		eaMap = rangeTemplate.Ea
	} else {
		eaMap = make(map[string]interface{})
	}

	ea, err := json.Marshal(eaMap)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":                      rangeTemplate.Ref,
		"name":                    rangeTemplate.Name,
		"ext_attrs":               string(ea),
		"number_of_addresses":     int(*rangeTemplate.NumberOfAddresses),
		"offset":                  int(*rangeTemplate.Offset),
		"server_association_type": rangeTemplate.ServerAssociationType,
	}
	if rangeTemplate.Comment != nil {
		res["comment"] = *rangeTemplate.Comment
	}
	if rangeTemplate.UseOptions != nil {
		res["use_options"] = *rangeTemplate.UseOptions
	}
	if rangeTemplate.Options != nil {
		options := convertDhcpOptionsToInterface(rangeTemplate.Options)
		res["options"] = options
	}
	if rangeTemplate.FailoverAssociation != nil {
		res["failover_association"] = *rangeTemplate.FailoverAssociation
	}
	if rangeTemplate.Member != nil {
		res["member"] = convertDhcpMemberToMap(rangeTemplate.Member)
	}
	return res, nil
}
