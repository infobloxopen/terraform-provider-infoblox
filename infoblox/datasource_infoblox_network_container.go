package infoblox

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func dataSourceIpv4NetworkContainer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIpv4NetworkContainerRead,
		Schema: map[string]*schema.Schema{
			"network_view": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Newtwork view's name the network container belongs to.",
			},
			"cidr": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The CIDR value of the network container.",
			},
			"comment": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Network container's description.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Extensible attributes for the network container.",
			},
		},
	}
}

func dataSourceIpv4NetworkContainerRead(d *schema.ResourceData, m interface{}) error {
	networkView := d.Get("network_view").(string)
	cidr := d.Get("cidr").(string)

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")

	networkContainer, err := objMgr.GetNetworkContainer(networkView, cidr, false, nil)
	if err != nil {
		return fmt.Errorf("Getting NetworkContainer %s failed : %s", cidr, err.Error())
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
