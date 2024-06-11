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

func dataSourceHostRecord() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHostRecordRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:     schema.TypeMap,
				Required: true,
			},

			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of Host records matching filters",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dns_view": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     defaultDNSView,
							Description: "DNS view under which the zone has been created.",
						},
						"fqdn": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The host name for Host Record in FQDN format.",
						},
						"ipv4_addr": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "IPv4 address of host record.",
						},
						"ipv6_addr": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "IPv6 address of host record.",
						},
						"zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The zone which the record belongs to.",
						},
						"ttl": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "TTL attribute value for the record.",
						},
						"comment": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of the Host-record.",
						},
						"ext_attrs": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Extensible attributes of the Host-record, as a map in JSON format",
						},
						"mac_addr": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "MAC address of a cloud instance.",
						},
						"duid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "DHCP unique identifier for IPv6.",
						},
						"enable_dhcp": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "The flag which defines if the host record is to be used for IPAM purposes.",
						},
						"enable_dns": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							Description: "flag that defines if the host record is to be used for DNS purposes.",
						},
					},
				},
			},
		},
	}
}

func dataSourceHostRecordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(ibclient.IBConnector)

	var diags diag.Diagnostics

	n := &ibclient.HostRecord{}
	n.SetReturnFields(append(n.ReturnFields(), "extattrs", "comment", "zone", "ttl", "configure_for_dns"))

	filters := filterFromMap(d.Get("filters").(map[string]interface{}))
	qp := ibclient.NewQueryParams(false, filters)
	var res []ibclient.HostRecord

	err := connector.GetObject(n, "", qp, &res)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed getting Host-record: %s", err.Error()))
	}

	if res == nil {
		return diag.FromErr(fmt.Errorf("API returns a nil/empty ID for the Host Record"))
	}

	// TODO: temporary scaffold, need to rework marshalling/unmarshalling of EAs
	//       (avoiding additional layer of keys ("value" key)
	results := make([]interface{}, 0, len(res))
	for _, r := range res {
		recordaFlat, err := flattenRecordHost(r)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to flatten Host Record: %w", err))
		}

		results = append(results, recordaFlat)
	}

	err = d.Set("results", results)
	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenRecordHost(hostRecord ibclient.HostRecord) (map[string]interface{}, error) {
	var eaMap map[string]interface{}
	if hostRecord.Ea != nil && len(hostRecord.Ea) > 0 {
		eaMap = hostRecord.Ea
	} else {
		eaMap = make(map[string]interface{})
	}

	ea, err := json.Marshal(eaMap)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"id":        hostRecord.Ref,
		"zone":      hostRecord.Zone,
		"dns_view":  hostRecord.View,
		"ext_attrs": string(ea),
	}

	if hostRecord.Ipv4Addrs != nil {
		res["ipv4_addr"] = hostRecord.Ipv4Addrs[0].Ipv4Addr
		res["mac_addr"] = hostRecord.Ipv4Addrs[0].Mac
		res["enable_dhcp"] = hostRecord.Ipv4Addrs[0].EnableDhcp
		res["enable_dns"] = hostRecord.EnableDns

	}
	if hostRecord.Ipv6Addrs != nil {
		res["ipv6_addr"] = hostRecord.Ipv6Addrs[0].Ipv6Addr
		res["duid"] = hostRecord.Ipv6Addrs[0].Duid
		res["enable_dhcp"] = hostRecord.Ipv6Addrs[0].EnableDhcp
		res["enable_dns"] = hostRecord.EnableDns
	}

	if hostRecord.UseTtl != nil {
		if !*hostRecord.UseTtl {
			res["ttl"] = ttlUndef
		}
	}

	if hostRecord.Ttl != nil && *hostRecord.Ttl > 0 {
		res["ttl"] = *hostRecord.Ttl
	} else {
		res["ttl"] = ttlUndef
	}

	if hostRecord.Name != nil {
		res["fqdn"] = *hostRecord.Name
	}

	if hostRecord.Comment != nil {
		res["comment"] = *hostRecord.Comment
	}

	return res, nil

}
