package infoblox

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func dataSourceTXTRecord() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTXTRecordRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_view": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "DNS View which the zone exists within.",
			},
			"zone": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "zone under which record has been created.",
			},
			"fqdn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "FQDN for the TXT-Record.",
			},
			"text": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Data of the TXT-Record.",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "TTL value for the TXT-Record.",
			},
			"comment": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of the TXT-Record.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Extensible attributes of the TXT-record, as a map in JSON format.",
			},
		},
	}
}

func dataSourceTXTRecordRead(d *schema.ResourceData, m interface{}) error {
	dnsView := d.Get("dns_view").(string)
	fqdn := d.Get("fqdn").(string)

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	obj, err := objMgr.GetTXTRecord(dnsView, fqdn)
	if err != nil {
		return fmt.Errorf("failed getting TXT-Record: %s", err)
	}

	if err = d.Set("text", obj.Text); err != nil {
		return err
	}

	ttl := int(obj.Ttl)
	if !obj.UseTtl {
		ttl = ttlUndef
	}
	if err = d.Set("ttl", ttl); err != nil {
		return err
	}

	if obj.Ea != nil && len(obj.Ea) > 0 {
		eaMap := (map[string]interface{})(obj.Ea)
		ea, err := json.Marshal(eaMap)
		if err != nil {
			return err
		}
		if err = d.Set("ext_attrs", string(ea)); err != nil {
			return err
		}
	}

	if err = d.Set("comment", obj.Comment); err != nil {
		return err
	}

	if err = d.Set("zone", obj.Zone); err != nil {
		return err
	}

	d.SetId(obj.Ref)

	return nil
}
