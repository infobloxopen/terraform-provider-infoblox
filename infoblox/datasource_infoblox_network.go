package infoblox

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client"
)

func dataSourceNetwork() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetworkRead,
		Schema: map[string]*schema.Schema{
			"network_view_name": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
			},
			"network_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cidr": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tenant_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"gateway": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceNetworkRead(d *schema.ResourceData, m interface{}) error {
	connector := m.(*ibclient.Connector)

	cidr := d.Get("cidr").(string)
	networkViewName := d.Get("network_view_name").(string)
	tenantID := d.Get("tenant_id").(string)

	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	obj, err := objMgr.GetNetwork(networkViewName, cidr, nil)
	if err != nil {
		return fmt.Errorf("Getting Network block from network (%s) failed : %s", cidr, err)
	}

	if obj == nil {
		return fmt.Errorf("API returns a nil/empty id on network (%s) failed", cidr)
	}
	d.SetId(obj.Ref)
	if obj.Ea["Network Name"] != nil {
		d.Set("network_name", obj.Ea["Network Name"])
	}
	return nil
}
