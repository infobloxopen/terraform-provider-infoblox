package infoblox

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func dataSourceARecord() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceARecordRead,

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
			"zone": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Zone under which record has been created.",
			},
			"fqdn": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "A record FQDN.",
			},
			"ip_addr": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "IP address the A-record points to",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "TTL value for the A-record.",
			},
			"comment": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of the A-record.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Extensible attributes of the A-record, as a map in JSON format",
			},
		},
	}
}

func dataSourceARecordRead(d *schema.ResourceData, m interface{}) error {

	dnsView := d.Get("dns_view").(string)
	fqdn := d.Get("fqdn").(string)
	ipAddr := d.Get("ip_addr").(string)

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	aRec, err := objMgr.GetARecord(dnsView, fqdn, ipAddr)
	if err != nil {
		return fmt.Errorf("failed getting A-record: %s", err.Error())
	}

	d.SetId(aRec.Ref)
	if err := d.Set("zone", aRec.Zone); err != nil {
		return err
	}
	if err := d.Set("ttl", aRec.Ttl); err != nil {
		return err
	}
	if err := d.Set("comment", aRec.Comment); err != nil {
		return err
	}

	dsExtAttrsVal := aRec.Ea
	dsExtAttrs, err := dsExtAttrsVal.MarshalJSON()
	if err != nil {
		return err
	}
	if err := d.Set("ext_attrs", string(dsExtAttrs)); err != nil {
		return err
	}
	return nil
}
