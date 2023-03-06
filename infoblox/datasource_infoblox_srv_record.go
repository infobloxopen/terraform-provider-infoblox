package infoblox

import (
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
				Required:    true,
				Description: "DNS View which the zone exists within.",
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
				Computed:    true,
				Description: "Configures port (0-65535) for this SRV record.",
			},
			"target": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Provides service for domain name in the SRV record.",
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
			"extattrs": {
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

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	srvRec, err := objMgr.GetSRVRecord(dnsView, name)
	if err != nil {
		return fmt.Errorf("failed getting SRV-Record: %s", err.Error())
	}
	d.SetId(srvRec.Ref)
	if err := d.Set("priority", srvRec.Priority); err != nil {
		return err
	}
	if err := d.Set("weight", srvRec.Weight); err != nil {
		return err
	}
	if err := d.Set("port", srvRec.Port); err != nil {
		return err
	}
	if err := d.Set("target", srvRec.Target); err != nil {
		return err
	}
	if err := d.Set("ttl", srvRec.Ttl); err != nil {
		return err
	}
	if err := d.Set("comment", srvRec.Comment); err != nil {
		return err
	}

	dsExtAttrsVal := srvRec.Ea
	dsExtAttrs, err := dsExtAttrsVal.MarshalJSON()
	if err != nil {
		return err
	}
	if err := d.Set("extattrs", string(dsExtAttrs)); err != nil {
		return err
	}
	return nil
}
