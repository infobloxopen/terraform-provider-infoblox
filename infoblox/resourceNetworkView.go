package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/infobloxopen/infoblox-go-client"
)

func resourceNetworkView() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkViewCreate,
		Read:   resourceNetworkViewRead,
		Update: resourceNetworkViewUpdate,
		Delete: resourceNetworkViewDelete,

		Schema: map[string]*schema.Schema{
			"network_view_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("net_view_name", nil),
				Description: "The name you want to give to your  network view",
			},
		},
	}
}

func resourceNetworkViewCreate(d *schema.ResourceData, m interface{}) error {
	connector := m.(*ibclient.Connector)
	objMgr := ibclient.NewObjectManager(connector, "terraform", "goclient1")
	network_view_name, err := objMgr.CreateNetworkView(d.Get("network_view_name").(string))
	if err != nil {
		fmt.Errorf("Failed to create Network View")
	}
	d.SetId(network_view_name.Name)

	return nil
}
func resourceNetworkViewRead(d *schema.ResourceData, m interface{}) error {

	return nil
}
func resourceNetworkViewUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}
func resourceNetworkViewDelete(d *schema.ResourceData, m interface{}) error {

	return nil
}
