package infoblox

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client"
)

func dataSourceAAAARecord() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAAAARecordRead,

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
				Description: "AAAA record FQDN.",
			},
			"ipv6_addr": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "IP address the AAAA-record points to",
			},
			"ttl": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "TTL value for the AAAA-record.",
			},
			"comment": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of the AAAA-record.",
			},
			"ext_attrs": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Extensible attributes of the AAAA-record, as a map in JSON format",
			},
		},
	}
}

func dataSourceAAAARecordRead(d *schema.ResourceData, m interface{}) error {

	dnsView := d.Get("dns_view").(string)
	fqdn := d.Get("fqdn").(string)
	ipv6Addr := d.Get("ipv6_addr").(string)

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	aaaaRec, err := objMgr.GetAAAARecord(dnsView, fqdn, ipv6Addr)
	if err != nil {
		return fmt.Errorf("failed getting AAAA-record: %s", err.Error())
	}

	d.SetId(aaaaRec.Ref)
	if err := d.Set("zone", aaaaRec.Zone); err != nil {
		return err
	}
	if err := d.Set("ttl", aaaaRec.Ttl); err != nil {
		return err
	}
	if err := d.Set("comment", aaaaRec.Comment); err != nil {
		return err
	}

	dsExtAttrsVal := aaaaRec.Ea
	dsExtAttrs, err := dsExtAttrsVal.MarshalJSON()
	if err != nil {
		return err
	}
	if err := d.Set("ext_attrs", string(dsExtAttrs)); err != nil {
		return err
	}
	return nil
}
