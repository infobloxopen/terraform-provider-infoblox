package infoblox

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func dataSourceNetworkContainer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIPv4NetworkRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_view": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cidr": {
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

func dataSourceNetworkContainerRead(d *schema.ResourceData, m interface{}) error {

	networkView := d.Get("network_view").(string)
	cidr := d.Get("cidr").(string)

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	networkContainer, err := objMgr.GetNetworkContainer(networkView, cidr, false, nil)
	if err != nil {
		return fmt.Errorf("Getting NetworkContainer %s failed : %s", cidr, err.Error())
	}
	if networkContainer == nil {
		return fmt.Errorf("API returns a nil/empty id on networkContainer %s", cidr)
	}

	d.SetId(networkContainer.Ref)

	if err := d.Set("comment", networkContainer.Comment); err != nil {
		return err
	}

	dsExtAttrsVal := networkContainer.Ea
	dsExtAttrs, err := dsExtAttrsVal.MarshalJSON()
	if err != nil {
		return err
	}

	if err := d.Set("ext_attrs", string(dsExtAttrs)); err != nil {
		return err
	}
	return nil
}
