package infoblox

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func dataSourceAAAARecord() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAAAARecordRead,

		Schema: map[string]*schema.Schema{
			"dns_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     defaultDNSView,
				Description: "DNS view which the record's zone belongs to.",
			},
			"fqdn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the AAAA-record in FQDN format.",
			},
			"ipv6_addr": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "IPv6 address of the AAAA-record.",
			},
			"zone": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Zone which the record belongs to.",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "TTL attribute value for the record.",
			},
			"comment": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The content of 'comment' field at the AAAA-record.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Extensible attributes of the AAAA-record",
			},
		},
	}
}

func dataSourceAAAARecordRead(d *schema.ResourceData, m interface{}) error {

	dnsView := d.Get("dns_view").(string)
	fqdn := d.Get("fqdn").(string)
	ipAddr := d.Get("ipv6_addr").(string)

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	aaaaRec, err := objMgr.GetAAAARecord(dnsView, fqdn, ipAddr)
	if err != nil {
		return fmt.Errorf("failed getting AAAA-record: %s", err.Error())
	}

	d.SetId(aaaaRec.Ref)
	if err := d.Set("zone", aaaaRec.Zone); err != nil {
		return err
	}
	if aaaaRec.UseTtl {
		if err := d.Set("ttl", aaaaRec.Ttl); err != nil {
			return err
		}
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
