package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/infobloxopen/infoblox-go-client"
)

func resourceNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkCreate,
		Read:   resourceNetworkRead,
		Update: resourceNetworkUpdate,
		Delete: resourceNetworkDelete,

		Schema: map[string]*schema.Schema{
			"networkviewname": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("nv_view_name", nil),
				Description: "give the nnetviewname you created",
			},
			"network_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("network_name", nil),
				Description: "The name you want to give to your network",
			},
			"cidr": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("nv_address", nil),
				Description: "",
			},
		},
	}
}

func resourceNetworkCreate(d *schema.ResourceData, m interface{}) error {

	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", "goclient1")

	nwname, err := objMgr.CreateNetwork(d.Get("networkviewname").(string), d.Get("cidr").(string), d.Get("network_name").(string))
	if err != nil {
		fmt.Errorf("Error creating network")
	}
	d.SetId(nwname.Cidr)
	return nil
}
func resourceNetworkRead(d *schema.ResourceData, m interface{}) error {
	return nil
}
func resourceNetworkUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}
func resourceNetworkDelete(d *schema.ResourceData, m interface{}) error {
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", "goclient1")

	ref, err := objMgr.GetNetwork(d.Get("networkviewname").(string), d.Get("cidr").(string), nil)

	if err != nil {
		fmt.Errorf("cant get ")
	}
	objMgr.DeleteNetwork(ref.Ref, d.Get("networkviewname").(string))

	d.SetId("")
	return nil
}
