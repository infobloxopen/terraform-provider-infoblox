package infoblox

import (
	"encoding/json"
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
			"zone": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The zone which the record belongs to.",
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
	obj, err := objMgr.GetPTRRecord(dnsView, ptrdname, recordName, ipAddr)
	if err != nil {
		return fmt.Errorf("failed getting PTR-record: %s", err.Error())
	}

	if obj.Ipv4Addr != "" {
		ipAddr = obj.Ipv4Addr
	} else {
		ipAddr = obj.Ipv6Addr
	}
	if err := d.Set("ip_addr", ipAddr); err != nil {
		return err
	}

	ttl := int(obj.Ttl)
	if !obj.UseTtl {
		ttl = ttlUndef
	}
	if err = d.Set("ttl", ttl); err != nil {
		return err
	}

	// TODO: temporary scaffold, need to rework marshalling/unmarshalling of EAs
	//       (avoiding additional layer of keys ("value" key)
	var eaMap map[string]interface{}
	if obj.Ea != nil && len(obj.Ea) > 0 {
		eaMap = (map[string]interface{})(obj.Ea)
	} else {
		eaMap = make(map[string]interface{})
	}
	ea, err := json.Marshal(eaMap)
	if err != nil {
		return err
	}
	if err = d.Set("ext_attrs", string(ea)); err != nil {
		return err
	}

	if err := d.Set("record_name", obj.Name); err != nil {
		return err
	}
	if err = d.Set("zone", obj.Zone); err != nil {
		return err
	}
	if err := d.Set("comment", obj.Comment); err != nil {
		return err
	}

	d.SetId(obj.Ref)

	return nil
}
