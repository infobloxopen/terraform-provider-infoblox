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
			"network_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
				Description: "Network view name of NIOS server.",
			},
			"cidr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The network address in cidr format under which record has to be created.",
			},
			"ip_addr": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "IPv4/IPv6 address for record creation. Set the field with valid IP for static allocation. If to be dynamically allocated set cidr field",
			},
			"dns_view": {
				Type:        schema.TypeString,
				Default:     "default",
				Optional:    true,
				Description: "Dns View under which the zone has been created.",
			},
			"ptrdname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The domain name in FQDN to which the record should point to.",
			},
			"record_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The name of the DNS PTR record in FQDN format",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     ttlUndef,
				Description: "TTL attribute value for the record.",
			},
			"comment": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "A description about PTR record.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "The Extensible attributes of PTR record to be added/updated, as a map in JSON format",
			},
		},
	}
}

func dataSourcePtrRecordRead(d *schema.ResourceData, m interface{}) error {

	dnsView := d.Get("dns_view").(string)
	ptrdname := d.Get("ptrdname").(string)
	ipAddr := d.Get("ip_addr").(string)
	recordName := d.Get("record_name").(string)

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	ptrRecord, err := objMgr.GetPTRRecord(dnsView, ptrdname, recordName, ipAddr)
	if err != nil {
		return fmt.Errorf("failed getting PTR-record: %s", err.Error())
	}

	d.SetId(ptrRecord.Ref)
	if err := d.Set("zone", ptrRecord.Zone); err != nil {
		return err
	}
	if err := d.Set("ttl", ptrRecord.Ttl); err != nil {
		return err
	}
	if err := d.Set("comment", ptrRecord.Comment); err != nil {
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
