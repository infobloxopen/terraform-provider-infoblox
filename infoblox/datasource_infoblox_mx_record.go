package infoblox

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func dataSourceMXRecord() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMXRecordRead,

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
			"fqdn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "FQDN for the MX-Record.",
			},
			"mail_exchanger": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A record used to specify mail server.",
			},
			"preference": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Configures the preference (0-65535) for this MX record.",
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

func dataSourceMXRecordRead(d *schema.ResourceData, m interface{}) error {

	dnsView := d.Get("dns_view").(string)
	fqdn := d.Get("fqdn").(string)
	mx := d.Get("mail_exchanger").(string)

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	obj, err := objMgr.GetMXRecord(dnsView, fqdn, mx)
	if err != nil {
		return fmt.Errorf("failed getting MX-Record: %s", err)
	}

	ttl := int(obj.Ttl)
	if !obj.UseTtl {
		ttl = ttlUndef
	}
	if err = d.Set("ttl", ttl); err != nil {
		return err
	}

	if obj.Ea != nil && len(obj.Ea) > 0 {
		// TODO: temporary scaffold, need to rework marshalling/unmarshalling of EAs
		//       (avoiding additional layer of keys ("value" key)
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
	if err = d.Set("preference", obj.Preference); err != nil {
		return err
	}
	d.SetId(obj.Ref)

	return nil
}
