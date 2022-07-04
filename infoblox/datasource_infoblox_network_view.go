package infoblox

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func dataSourceNetworkView() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetworkViewRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Desired name of the view shown in NIOS appliance.",
			},
			"comment": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "A description of the IP allocation.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "The Extensible attributes of the network container to be added/updated, as a map in JSON format",
			},
		},
	}
}

func dataSourceNetworkViewRead(d *schema.ResourceData, m interface{}) error {

	name := d.Get("name").(string)

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	networkView, err := objMgr.GetNetworkView(name)
	if err != nil {
		return fmt.Errorf("Getting networkView %s failed : %s", name, err.Error())
	}
	if networkView == nil {
		return fmt.Errorf("API returns a nil/empty id on networkView %s", name)
	}

	d.SetId(networkView.Ref)

	if err := d.Set("comment", networkView.Comment); err != nil {
		return err
	}

	dsExtAttrsVal := networkView.Ea
	dsExtAttrs, err := dsExtAttrsVal.MarshalJSON()
	if err != nil {
		return err
	}

	if err := d.Set("ext_attrs", string(dsExtAttrs)); err != nil {
		return err
	}
	return nil
}
