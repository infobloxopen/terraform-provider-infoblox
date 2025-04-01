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

func dataSourceNSRecord() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNSRecordRead,
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
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the NS record in FQDN format. This value can be in unicode format.",
						},
						"nameserver": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The domain name of an authoritative server for the redirected zone.",
						},
						"addresses": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of zone name servers.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "he address of the Zone Name Server.",
									},
									"auto_create_ptr": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Flag to indicate if ptr records need to be auto created.",
									},
								},
							},
						},
						"dns_view": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the DNS view in which the record resides.Example: “external”.",
						},
						"zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the zone in which the record resides.If a view is not specified when searching by zone, the default view is used.",
						},
						"creator": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The record creator. The valid values are 'STATIC' and 'SYSTEM'",
						},
						"dns_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the NS record in punycode format.",
						},
						"policy": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The host name policy for the record.",
						},
						"ms_delegation_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The MS delegation point name.",
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

func dataSourceNSRecordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(ibclient.IBConnector)

	var diags diag.Diagnostics

	filters := filterFromMap(d.Get("filters").(map[string]interface{}))

	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	qp := ibclient.NewQueryParams(false, filters)
	res, err := objMgr.GetAllRecordNS(qp)
	if err != nil {
		return diag.FromErr(err)
	}

	if res == nil {
		return diag.FromErr(fmt.Errorf("API returns a nil/empty ID for the NS Record"))
	}

	// TODO: temporary scaffold, need to rework marshalling/unmarshalling of EAs
	//       (avoiding additional layer of keys ("value" key)
	results := make([]interface{}, 0, len(res))
	for _, r := range res {
		recordaFlat, err := flattenRecordNS(r)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to flatten NS Record  : %w", err))
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

func flattenRecordNS(recordNS ibclient.RecordNS) (map[string]interface{}, error) {

	res := map[string]interface{}{
		"id":                 recordNS.Ref,
		"zone":               recordNS.Zone,
		"dns_view":           recordNS.View,
		"name":               recordNS.Name,
		"creator":            recordNS.Creator,
		"policy":             recordNS.Policy,
		"dns_name":           recordNS.DnsName,
		"nameserver":         recordNS.Nameserver,
		"ms_delegation_name": recordNS.MsDelegationName,
	}
	if recordNS.Addresses != nil {
		addressInterface := convertZoneNameServersToInterface(recordNS.Addresses)
		res["addresses"] = addressInterface
	}

	if recordNS.CloudInfo != nil {
		cloudInfo, err := serializeGridCloudApiInfo(recordNS.CloudInfo)
		if err != nil {
			return nil, err
		}
		res["cloud_info"] = cloudInfo
	}
	return res, nil

}

func serializeGridCloudApiInfo(gci *ibclient.GridCloudapiInfo) (string, error) {
	// Create a map to hold the serialized values
	gciMap := map[string]interface{}{}

	// Check and add fields to the map if they are non-zero
	if gci.DelegatedMember != nil {
		// If DelegatedMember exists, add it to the map
		gciMap["delegated_member"] = gci.DelegatedMember
	}
	if gci.DelegatedScope != "" {
		gciMap["delegated_scope"] = gci.DelegatedScope
	}
	if gci.DelegatedRoot != "" {
		gciMap["delegated_root"] = gci.DelegatedRoot
	}
	gciMap["owned_by_adaptor"] = gci.OwnedByAdaptor // boolean field, will be added regardless
	if gci.Usage != "" {
		gciMap["usage"] = gci.Usage
	}
	if gci.Tenant != "" {
		gciMap["tenant"] = gci.Tenant
	}
	if gci.MgmtPlatform != "" {
		gciMap["mgmt_platform"] = gci.MgmtPlatform
	}
	if gci.AuthorityType != "" {
		gciMap["authority_type"] = gci.AuthorityType
	}

	// If no fields to serialize, return an empty string
	if len(gciMap) == 0 {
		return "", nil
	}
	// Marshal the map to JSON
	gciJSON, err := json.Marshal(gciMap)
	if err != nil {
		return "", err
	}
	// Return the serialized JSON string
	return string(gciJSON), nil
}
