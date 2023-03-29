package infoblox

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func dataSourceARecord() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceARecordRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     defaultDNSView,
				Description: "DNS view which the record's zone belongs to.",
			},
			"fqdn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A record FQDN.",
			},
			"ip_addr": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "IP address the A-record points to",
			},
			"zone": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The zone which the record belongs to.",
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
	obj, err := objMgr.GetARecord(dnsView, fqdn, ipAddr)
	if err != nil {
		return fmt.Errorf("failed getting A-record: %s", err.Error())
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

	if err := d.Set("zone", obj.Zone); err != nil {
		return err
	}
	if err := d.Set("comment", obj.Comment); err != nil {
		return err
	}

	d.SetId(obj.Ref)

	return nil
}
