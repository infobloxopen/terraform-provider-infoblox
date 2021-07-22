package infoblox

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func dataSourceIPv4Network() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIPv4NetworkRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_view": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cidr": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A string describing the network",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Extensible attributes for network datasource, as a map in JSON format",
			},
		},
	}
}

func dataSourceIPv4NetworkRead(d *schema.ResourceData, m interface{}) error {

	networkView := d.Get("network_view").(string)
	cidr := d.Get("cidr").(string)

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	network, err := objMgr.GetNetwork(networkView, cidr, false, nil)
	if err != nil {
		return fmt.Errorf("Getting Network %s failed : %s", cidr, err.Error())
	}
	if network == nil {
		return fmt.Errorf("API returns a nil/empty id on network %s", cidr)
	}

	d.SetId(network.Ref)

	if err := d.Set("comment", network.Comment); err != nil {
		return err
	}

	dsExtAttrsVal := network.Ea
	dsExtAttrs, err := dsExtAttrsVal.MarshalJSON()
	if err != nil {
		return err
	}

	if err := d.Set("ext_attrs", string(dsExtAttrs)); err != nil {
		return err
	}
	return nil
}
