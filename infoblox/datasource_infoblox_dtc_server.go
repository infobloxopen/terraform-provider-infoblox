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

func dataSourceDtcServer() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDtcServerRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:     schema.TypeMap,
				Required: true,
			},
			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of A records matching filters",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"auto_create_host_record": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enabling this option will auto-create a single read-only A/AAAA/CNAME record corresponding to the configured hostname and update it if the hostname changes.\n\n",
						},
						"comment": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "Description of the Dtc server.",
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
							Default:     "",
							Description: "Extensible attributes of the  Dtc Server to be added/updated, as a map in JSON format",
						},
						"host": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The address or FQDN of the server.",
						},
						"monitors": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of IP/FQDN and monitor pairs to be used for additional monitoring.\n\n",
							Elem: &schema.Resource{
								//check the required part once
								Schema: map[string]*schema.Schema{
									"host": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "IP address or FQDN of the server used for monitoring.",
									},
									"monitor_name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The monitor name related to server.",
									},
									"monitor_type": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The monitor type related to server.",
									},
								},
							},
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The DTC Server display name.",
						},
						"sni_hostname": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The hostname for Server Name Indication (SNI) in FQDN format.",
						},
						"use_sni_hostname": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Use flag for: sni_hostname",
						},
					},
				},
			},
		},
	}
}
func dataSourceDtcServerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(ibclient.IBConnector)

	var diags diag.Diagnostics

	filters := filterFromMap(d.Get("filters").(map[string]interface{}))

	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	qp := ibclient.NewQueryParams(false, filters)

	res, err := objMgr.GetAllDtcServer(qp)
	if err != nil {
		return diag.FromErr(err)
	}
	if res == nil {
		return diag.FromErr(fmt.Errorf("API returns a nil/empty ID for zone forward"))
	}
	// TODO: temporary scaffold, need to rework marshalling/unmarshalling of EAs
	//       (avoiding additional layer of keys ("value" key)
	results := make([]interface{}, 0, len(res))
	for _, r := range res {
		dsFlat, err := flattenDtcServer(r, connector)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to flatten dtc server  : %w", err))
		}
		results = append(results, dsFlat)
	}
	err = d.Set("results", results)
	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
func flattenDtcServer(dtcServer ibclient.DtcServer, connector ibclient.IBConnector) (map[string]interface{}, error) {
	var eaMap map[string]interface{}
	if dtcServer.Ea != nil && len(dtcServer.Ea) > 0 {
		eaMap = dtcServer.Ea
	} else {
		eaMap = make(map[string]interface{})
	}

	ea, err := json.Marshal(eaMap)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":        dtcServer.Ref,
		"ext_attrs": string(ea),
		"name":      *dtcServer.Name,
		"host":      *dtcServer.Host,
	}
	if dtcServer.AutoCreateHostRecord != nil {
		res["auto_create_host_record"] = *dtcServer.AutoCreateHostRecord
	}
	if dtcServer.Comment != nil {
		res["comment"] = *dtcServer.Comment
	}
	if dtcServer.Disable != nil {
		res["disable"] = *dtcServer.Disable
	}
	if dtcServer.Monitors != nil {
		monitorInterface := convertDtcServerMonitorsToInterface(dtcServer.Monitors, connector)
		res["monitors"] = monitorInterface
	}
	if dtcServer.SniHostname != nil {
		res["sni_hostname"] = *dtcServer.SniHostname
	}
	if dtcServer.UseSniHostname != nil {
		res["use_sni_hostname"] = *dtcServer.UseSniHostname
	}
	return res, nil
}
