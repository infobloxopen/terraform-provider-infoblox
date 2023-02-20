package infoblox

import (
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
			"extattrs": {
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

	txtRec, err := objMgr.GetTXTRecord(dnsView, fqdn)
	if err != nil {
		return fmt.Errorf("failed getting TXT-Record: %s", err.Error())
	}

	d.SetId(txtRec.Ref)
	if err := d.Set("zone", txtRec.Zone); err != nil {
		return err
	}
	if err := d.Set("ttl", txtRec.Ttl); err != nil {
		return err
	}
	if err := d.Set("comment", txtRec.Comment); err != nil {
		return err
	}

	dsExtAttrsVal := txtRec.Ea
	dsExtAttrs, err := dsExtAttrsVal.MarshalJSON()
	if err != nil {
		return err
	}
	if err := d.Set("extattrs", string(dsExtAttrs)); err != nil {
		return err
	}
	return nil
}
