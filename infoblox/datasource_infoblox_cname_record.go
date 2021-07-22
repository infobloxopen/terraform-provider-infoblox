package infoblox

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func dataSourceCNameRecord() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCNameRecordRead,

		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_view": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Dns View under which the zone has been created.",
			},
			"canonical": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Canonical name in FQDN format.",
			},
			"alias": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The alias name in FQDN format.",
			},
			"zone": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Zone under which record has been created.",
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

	recordCNAME, err := objMgr.GetCNAMERecord(dnsView, canonical, alias)
	if err != nil {
		return fmt.Errorf("Getting CNAME Record failed : %s", err.Error())
	}

	d.SetId(recordCNAME.Ref)
	if err := d.Set("zone", recordCNAME.Zone); err != nil {
		return err
	}
	if err := d.Set("ttl", recordCNAME.Ttl); err != nil {
		return err
	}
	if err := d.Set("comment", recordCNAME.Comment); err != nil {
		return err
	}

	dsExtAttrsVal := recordCNAME.Ea
	dsExtAttrs, err := dsExtAttrsVal.MarshalJSON()
	if err != nil {
		return err
	}
	if err := d.Set("ext_attrs", string(dsExtAttrs)); err != nil {
		return err
	}

	return nil
}
