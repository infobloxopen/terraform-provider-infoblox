package infoblox

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"strconv"
	"time"
)

func dataSourceZoneForward() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceZoneForwardRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:     schema.TypeMap,
				Required: true,
			},
			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of Forward Zones matching filters",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"fqdn": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of this DNS zone",
						},
						"forward_to": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "The information for the remote name servers to which you want the Infoblox appliance to forward queries for a specified domain name.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The IP address of the remote name server to which you want the Infoblox appliance to forward queries for a specified domain name.",
									},
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The name of the remote name server to which you want the Infoblox appliance to forward queries for a specified domain name.",
									},
								},
							},
						},
						"view": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "default",
							Description: "The DNS view in which the zone is created.",
						},
						"zone_format": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "FORWARD",
							Description: "The format of the zone. Valid values are: FORWARD, IPV4, IPV6.",
						},
						"ns_group": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A forwarding member name server group.",
						},
						"external_ns_group": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A forwarding member name server group.",
						},
						"comment": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A descriptive comment.",
						},
						"disable": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Determines if the zone is disabled or not.",
						},
						"ext_attrs": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Extensible attributes of the zone forward to be added/updated, as a map in JSON format.",
						},
						"forwarders_only": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Determines if the appliance sends queries to forwarders only, and not to other internal or Internet root servers.",
						},
						"forwarding_servers": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The information for the remote name servers to which you want the Infoblox appliance to forward queries for a specified domain name.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The name of this Grid member in FQDN format.",
									},
									"forwarders_only": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     false,
										Description: "Determines if the appliance sends queries to forwarders only, and not to other internal or Internet root servers.",
									},
									"use_override_forwarders": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Determines if the appliance sends queries to name servers.",
									},
									"forward_to": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The information for the remote name servers to which you want the Infoblox appliance to forward queries for a specified domain name.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"address": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The IP address of the remote name server to which you want the Infoblox appliance to forward queries for a specified domain name.",
												},
												"name": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The name of the remote name server to which you want the Infoblox appliance to forward queries for a specified domain name.",
												},
											},
										},
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

func dataSourceZoneForwardRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(ibclient.IBConnector)

	var diags diag.Diagnostics

	filters := filterFromMap(d.Get("filters").(map[string]interface{}))

	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	qp := ibclient.NewQueryParams(false, filters)
	res, err := objMgr.GetZoneForwardFilters(qp)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to get zone forward records: %w", err))
	}

	if res == nil {
		return diag.FromErr(fmt.Errorf("API returns a nil/empty ID for zone forward"))
	}
	// TODO: temporary scaffold, need to rework marshalling/unmarshalling of EAs
	//       (avoiding additional layer of keys ("value" key)
	results := make([]interface{}, 0, len(res))
	for _, r := range res {
		zfFlat, err := flattenZoneForward(r)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to flatten zone forward  : %w", err))
		}
		results = append(results, zfFlat)
	}

	err = d.Set("results", results)
	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags

}

func flattenZoneForward(zf ibclient.ZoneForward) (map[string]interface{}, error) {
	var eaMap map[string]interface{}
	if zf.Ea != nil && len(zf.Ea) > 0 {
		eaMap = zf.Ea
	} else {
		eaMap = make(map[string]interface{})
	}

	ea, err := json.Marshal(eaMap)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":          zf.Ref,
		"fqdn":        zf.Fqdn,
		"ext_attrs":   string(ea),
		"zone_format": zf.ZoneFormat,
		"view":        *zf.View,
	}
	if zf.Comment != nil {
		res["comment"] = *zf.Comment
	}
	if zf.Disable != nil {
		res["disable"] = *zf.Disable
	}
	if zf.NsGroup != nil {
		res["ns_group"] = *zf.NsGroup
	}

	if zf.ExternalNsGroup != nil {
		res["external_ns_group"] = *zf.ExternalNsGroup
	}

	if zf.ForwardersOnly != nil {
		res["forwarders_only"] = *zf.ForwardersOnly
	}

	if zf.ForwardTo.IsNull == false {
		nsInterface := convertForwardToInterface(zf.ForwardTo)
		res["forward_to"] = nsInterface
	}

	if zf.ForwardingServers.Servers != nil {
		fwServersInterface, _ := convertForwardingServersToInterface(zf.ForwardingServers.Servers)
		res["forwarding_servers"] = fwServersInterface
	}
	return res, nil
}

func convertForwardingServersToInterface(zf []*ibclient.Forwardingmemberserver) ([]map[string]interface{}, error) {
	if zf == nil {
		return nil, errors.New("forwarding servers is nil")
	}
	fwServers := make([]map[string]interface{}, 0, len(zf))
	for _, fs := range zf {
		sMap := make(map[string]interface{})
		sMap["name"] = fs.Name
		sMap["forwarders_only"] = fs.ForwardersOnly
		sMap["use_override_forwarders"] = fs.UseOverrideForwarders
		if fs.ForwardTo.IsNull == false {
			nsInterface := convertForwardToInterface(fs.ForwardTo)
			sMap["forward_to"] = nsInterface
		}
		fwServers = append(fwServers, sMap)
	}
	return fwServers, nil
}
func convertForwardToInterface(nameServers ibclient.NullForwardTo) []map[string]interface{} {
	nsInterface := make([]map[string]interface{}, 0, len(nameServers.ForwardTo))
	for _, ns := range nameServers.ForwardTo {
		nsMap := make(map[string]interface{})
		nsMap["address"] = ns.Address
		nsMap["name"] = ns.Name
		nsInterface = append(nsInterface, nsMap)
	}
	return nsInterface
}
