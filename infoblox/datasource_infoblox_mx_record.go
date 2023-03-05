package infoblox

import (
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
				Computed:    true,
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
			"extattrs": {
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

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	mxRec, err := objMgr.GetMXRecord(dnsView, fqdn)
	if err != nil {
		return fmt.Errorf("failed getting MX-Record: %s", err.Error())
	}
	d.SetId(mxRec.ref)
	if err := d.Set("mail_exchanger", mxRec.MX); err != nil {
		return err
	}
	if err := d.Set("preference", mxRec.Priority); err != nil {
		return err
	}
	if err := d.Set("ttl", mxRec.Ttl); err != nil {
		return err
	}
	if err := d.Set("comment", mxRec.Comment); err != nil {
		return err
	}

	dsExtAttrsVal := mxRec.Ea
	dsExtAttrs, err := dsExtAttrsVal.MarshalJSON()
	if err != nil {
		return err
	}
	if err := d.Set("extattrs", string(dsExtAttrs)); err != nil {
		return err
	}
	return nil
}
