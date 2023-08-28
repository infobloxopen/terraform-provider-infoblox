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

func dataSourceZoneAuth() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceZoneAuthRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:     schema.TypeMap,
				Required: true,
			},

			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of zones matching filters",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"view": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     defaultDNSView,
							Description: "The name of the DNS view in which the zone resides.",
						},
						"fqdn": {
							Type:     schema.TypeString,
							Required: true,
							Description: "The name of this DNS zone. For a reverse zone, this is in 'address/cidr' " +
								"format. For other zones, this is in FQDN format. This value can be in " +
								"unicode format.",
						},
						"ns_group": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name server group that serves DNS for this zone.",
						},
						"zone_format": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Determines the format of this zone. Valid values are: FORWARD, IPV4, IPV6.",
						},
						"comment": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of Authoritative Zone Object.",
						},
						"ext_attrs": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Extensible attributes of the zone, as a map in JSON format",
						},
					},
				},
			},
		},
	}
}

func dataSourceZoneAuthRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(ibclient.IBConnector)

	var diags diag.Diagnostics

	n := &ibclient.ZoneAuth{}
	n.SetReturnFields(append(n.ReturnFields(), "extattrs", "comment", "zone_format", "ns_group"))

	filters := filterFromMap(d.Get("filters").(map[string]interface{}))
	qp := ibclient.NewQueryParams(false, filters)
	var res []ibclient.ZoneAuth

	err := connector.GetObject(n, "", qp, &res)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed getting Zone Auth: %s", err.Error()))
	}

	if res == nil {
		return diag.FromErr(fmt.Errorf("API returns a nil/empty ID for the Zone Auth"))
	}

	// TODO: temporary scaffold, need to rework marshalling/unmarshalling of EAs
	//       (avoiding additional layer of keys ("value" key)
	results := make([]interface{}, 0, len(res))
	for _, za := range res {
		zoneauthFlat, err := flattenZoneAuth(za)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to flatten Zone Auth : %w", err))
		}

		results = append(results, zoneauthFlat)
	}

	err = d.Set("results", results)
	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenZoneAuth(zoneauth ibclient.ZoneAuth) (map[string]interface{}, error) {
	var eaMap map[string]interface{}
	if zoneauth.Ea != nil && len(zoneauth.Ea) > 0 {
		eaMap = zoneauth.Ea
	} else {
		eaMap = make(map[string]interface{})
	}

	ea, err := json.Marshal(eaMap)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":          zoneauth.Ref,
		"zone_format": zoneauth.ZoneFormat,
		"fqdn":        zoneauth.Fqdn,
		"ext_attrs":   string(ea),
	}

	if zoneauth.View != nil {
		res["view"] = *zoneauth.View
	}

	if zoneauth.Comment != nil {
		res["comment"] = *zoneauth.Comment
	}

	if zoneauth.NsGroup != nil {
		res["ns_group"] = *zoneauth.NsGroup
	}

	return res, nil
}
