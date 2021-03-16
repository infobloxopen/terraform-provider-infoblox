package infoblox

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client"
)

func dataSourceARecord() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceARecordRead,

		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"zone": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Zone under which record has been created.",
			},
			"dns_view": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Dns View under which the zone has been created.",
			},
			"fqdn": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "A record FQDN.",
			},
			"ip_addr": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "IP address.",
			},
			"eas": &schema.Schema{
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Extension attributes",
			},
			"first_record": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Return first found record. Raise error if set to false and more than one record found.",
			},
		},
	}
}

func dataSourceARecordRead(d *schema.ResourceData, m interface{}) error {
	var records []ibclient.RecordA

	zone := d.Get("zone").(string)
	dnsView := d.Get("dns_view").(string)
	fqdn := d.Get("fqdn").(string)
	ip_addr := d.Get("ip_addr").(string)
	first_record := d.Get("first_record").(bool)

	connector := m.(*ibclient.Connector)

	search_data := ibclient.NewRecordA(
		ibclient.RecordA{
			Ipv4Addr: ip_addr,
			Name:     fqdn,
			Zone:     zone,
			View:     dnsView,
		})
	err := connector.GetObject(search_data, "", &records)
	d.SetId("")
	if err != nil {
		return fmt.Errorf("Read A record failed: %s", err)
	}
	if len(records) == 0 {
		return fmt.Errorf("No A record found. view(%s) zone(%s) fqdn(%s) ip_addr(%s)", dnsView, zone, fqdn, ip_addr)
	}
	if len(records) > 1 && !first_record {
		return fmt.Errorf("Expect single record but found %d A records. view(%s) zone(%s) fqdn(%s) ip_addr(%s)", len(records), dnsView, zone, fqdn, ip_addr)
	}
	d.Set("ip_addr", records[0].Ipv4Addr)
	d.Set("zone", records[0].Zone)
	d.Set("dns_view", records[0].View)
	d.Set("fqdn", records[0].Name)

	eas := make(map[string]string)
	for key, value := range records[0].Ea {
		eas[key] = fmt.Sprintf("%v", value)
	}
	d.Set("eas", eas)

	d.SetId(records[0].Ref)

	return nil
}
