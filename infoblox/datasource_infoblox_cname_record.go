package infoblox

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func dataSourceCNameRecord() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCNameRecordRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     defaultDNSView,
				Description: "DNS view which the record's zone belongs to.",
			},
			"canonical": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Canonical name in FQDN format.",
			},
			"alias": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The alias name in FQDN format.",
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
				Description: "A description about CNAME record.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Extensible attributes of CNAME record, as a map in JSON format",
			},
		},
	}
}

func dataSourceCNameRecordRead(d *schema.ResourceData, m interface{}) error {
	dnsView := d.Get("dns_view").(string)
	canonical := d.Get("canonical").(string)
	alias := d.Get("alias").(string)

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")
	obj, err := objMgr.GetCNAMERecord(dnsView, canonical, alias)
	if err != nil {
		return fmt.Errorf("Getting CNAME Record failed : %s", err.Error())
	}

	ttl := int(obj.Ttl)
	if !obj.UseTtl {
		ttl = ttlUndef
	}
	if err = d.Set("ttl", ttl); err != nil {
		return err
	}

	// TODO: temporary scaffold, need to rework marshalling/unmarshalling of EAs
	//       (avoiding additional layer of keys ("value" key)
	var eaMap map[string]interface{}
	if obj.Ea != nil && len(obj.Ea) > 0 {
		eaMap = (map[string]interface{})(obj.Ea)
	} else {
		eaMap = make(map[string]interface{})
	}
	ea, err := json.Marshal(eaMap)
	if err != nil {
		return err
	}
	if err = d.Set("ext_attrs", string(ea)); err != nil {
		return err
	}

	if err := d.Set("zone", obj.Zone); err != nil {
		return err
	}
	if err := d.Set("comment", obj.Comment); err != nil {
		return err
	}

	d.SetId(obj.Ref)

	return nil
}
