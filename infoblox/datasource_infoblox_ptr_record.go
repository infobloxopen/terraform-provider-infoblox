package infoblox

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func dataSourcePtrRecord() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePtrRecordRead,

		Schema: map[string]*schema.Schema{
			"dns_view": {
				Type:        schema.TypeString,
				Default:     defaultDNSView,
				Optional:    true,
				Description: "DNS view which the record's zone belongs to.",
			},
			"ptrdname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The domain name in FQDN which the record points to.",
			},
			"record_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "The name of the DNS PTR-record in FQDN format",
			},
			"ip_addr": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "IPv4/IPv6 address the PTR-record points from.",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "TTL attribute value for the record.",
			},
			"comment": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "PTR-record's description.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Extensible attributes of the PTR-record.",
			},
		},
	}
}

func dataSourcePtrRecordRead(d *schema.ResourceData, m interface{}) error {
	dnsView := d.Get("dns_view").(string)
	ptrdname := d.Get("ptrdname").(string)
	ipAddr := d.Get("ip_addr").(string)
	recordName := d.Get("record_name").(string)

	if ipAddr == "" && recordName == "" {
		return fmt.Errorf("either IP address or record's FQDN must be specified")
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	ptrRecord, err := objMgr.GetPTRRecord(dnsView, ptrdname, recordName, ipAddr)
	if err != nil {
		return fmt.Errorf("failed getting PTR-record: %s", err.Error())
	}

	d.SetId(ptrRecord.Ref)
	if err := d.Set("ttl", ptrRecord.Ttl); err != nil {
		return err
	}
	if err := d.Set("comment", ptrRecord.Comment); err != nil {
		return err
	}

	if err := d.Set("record_name", ptrRecord.Name); err != nil {
		return err
	}

	if ptrRecord.Ipv4Addr != "" {
		ipAddr = ptrRecord.Ipv4Addr
	} else {
		ipAddr = ptrRecord.Ipv6Addr
	}
	if err := d.Set("ip_addr", ipAddr); err != nil {
		return err
	}

	dsExtAttrsVal := ptrRecord.Ea
	dsExtAttrs, err := dsExtAttrsVal.MarshalJSON()
	if err != nil {
		return err
	}
	if err := d.Set("ext_attrs", string(dsExtAttrs)); err != nil {
		return err
	}
	return nil
}
