package infoblox

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client"
)

func dataSourceCNameRecord() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCNameRecordRead,

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
			"canonical": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Canonical name.",
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

func dataSourceCNameRecordRead(d *schema.ResourceData, m interface{}) error {
	var records []ibclient.RecordCNAME

	zone := d.Get("zone").(string)
	dnsView := d.Get("dns_view").(string)
	fqdn := d.Get("fqdn").(string)
	canonical := d.Get("canonical").(string)
	first_record := d.Get("first_record").(bool)

	connector := m.(*ibclient.Connector)

	cnRec := ibclient.NewEmptyRecordCNAME()
	sf := map[string]string{
		"canonical": canonical,
		"name":      fqdn,
		"zone":      zone,
		"view":      dnsView,
	}
	queryParams := ibclient.NewQueryParams(false, sf)
	err := connector.GetObject(cnRec, "", queryParams, &records)
	d.SetId("")
	if err != nil {
		return fmt.Errorf("Read CNAME record failed: %s", err)
	}
	if len(records) == 0 {
		return fmt.Errorf("No CNAME record found. view(%s) zone(%s) fqdn(%s) canonical(%s)", dnsView, zone, fqdn, canonical)
	}
	if len(records) > 1 && !first_record {
		return fmt.Errorf("Expect single record but found %d CNAME records. view(%s) zone(%s) fqdn(%s) canonical(%s)", len(records), dnsView, zone, fqdn, canonical)
	}
	d.Set("canonical", records[0].Canonical)
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
