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
				DefaultFunc: schema.EnvDefaultFunc("net_view_name", nil),
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
				DefaultFunc: schema.EnvDefaultFunc("net_address", nil),
				Description: "",
			},
		},
	}
}

func resourceNetworkCreate(d *schema.ResourceData, m interface{}) error {
	networkviewname := d.Get("networkviewname").(string)
	cidr := d.Get("cidr").(string)
	networkname :=d.Get("network_name").(string)

	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", "goclient1")

	nwname, err := objMgr.CreateNetwork(networkviewname, cidr, networkname)
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
	networkviewname := d.Get("networkviewname").(string)
	cidr := d.Get("cidr").(string)
	connector := m.(*ibclient.Connector)
	objMgr := ibclient.NewObjectManager(connector, "terraform", "goclient1")
	ref, err := objMgr.GetNetwork(networkviewname, cidr, nil)

	if err != nil {
		fmt.Errorf("cant get ")
	}
	objMgr.DeleteNetwork(ref.Ref, d.Get("networkviewname").(string))

	d.SetId("")
	return nil
}
