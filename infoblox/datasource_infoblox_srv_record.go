package infoblox

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func dataSourceSRVRecord() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSRVRecordRead,

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
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Combination of service name, protocol name and zone name.",
			},
			"priority": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Configures the priority (0-65535) for this SRV record.",
			},
			"weight": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Configures weight of the SRV record, valid values (0-65535).",
			},
			"port": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Configures port (0-65535) for this SRV record.",
			},
			"target": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Provides service for domain name in the SRV record.",
			},
			"zone": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The zone which the record belongs to.",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "TTL value for the SRV record.",
			},
			"comment": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of the SRV record.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Extensible attributes of the SRV-record to be added/updated, as a map in JSON format.",
			},
		},
	}
}

func dataSourceSRVRecordRead(d *schema.ResourceData, m interface{}) error {
	dnsView := d.Get("dns_view").(string)
	name := d.Get("name").(string)
	target := d.Get("target").(string)
	tempPortVal := d.Get("port").(int)
	if err := ibclient.CheckIntRange("port", tempPortVal, 0, 65535); err != nil {
		return err
	}
	port := uint32(tempPortVal)

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")
	obj, err := objMgr.GetSRVRecord(dnsView, name, target, port)
	if err != nil {
		return fmt.Errorf("failed getting SRV-Record: %s", err)
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

	if err := d.Set("priority", obj.Priority); err != nil {
		return err
	}
	if err := d.Set("weight", obj.Weight); err != nil {
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
