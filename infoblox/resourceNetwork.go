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
			"network_view_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("network_view_name", nil),
				Description: "Give the network view name  you created",
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
				Description: "Give the address in cidr format",
			},
			"tennant_id": &schema.Schema{
				Type: schema.TypeString,
				Optional:true,
				DefaultFunc: schema.EnvDefaultFunc("tennant_id",nil),
				Description:"Unique identifier of your instance",
				},
		},
	}
}

func resourceNetworkCreate(d *schema.ResourceData, m interface{}) error {
	network_view_name := d.Get("network_view_name").(string)
	cidr := d.Get("cidr").(string)
	network_name := d.Get("network_name").(string)
	tennant_id := d.Get("tennant_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tennant_id)

	nwname, err := objMgr.CreateNetwork(network_view_name, cidr, network_name)
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
	network_view_name := d.Get("network_view_name").(string)
	cidr := d.Get("cidr").(string)
	tennant_id := d.Get("tennant_id").(string)
	connector := m.(*ibclient.Connector)
	objMgr := ibclient.NewObjectManager(connector, "terraform", tennant_id)
	ref, err := objMgr.GetNetwork(network_view_name, cidr, nil)

	if err != nil {
		fmt.Errorf("cant get ")
	}
	objMgr.DeleteNetwork(ref.Ref, d.Get("network_view_name").(string))

	d.SetId("")
	return nil
}
