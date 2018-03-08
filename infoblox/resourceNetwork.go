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
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("network_view_name", "default"),
				Description: "Network view name available in NIOS Server",
			},
			"network_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("network_name", nil),
				Description: "The name of the network",
			},
			"cidr": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("net_address", nil),
				Description: "Give the address in cidr format",
			},
			"tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("tenant_id", nil),
				Description: "Unique identifier of your instance in cloud",
			},
		},
	}
}

func resourceNetworkCreate(d *schema.ResourceData, m interface{}) error {
	network_view_name := d.Get("network_view_name").(string)
	cidr := d.Get("cidr").(string)
	network_name := d.Get("network_name").(string)
	tenant_id := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenant_id)

	nwname, err := objMgr.CreateNetwork(network_view_name, cidr, network_name)
	if err != nil {
		return fmt.Errorf("Creation of network failed in network view (%s) : %s", network_view_name, err)
	}
	d.SetId(nwname.Cidr)

	return nil
}
func resourceNetworkRead(d *schema.ResourceData, m interface{}) error {
	network_view_name := d.Get("network_view_name").(string)
	cidr := d.Get("cidr").(string)
	tenant_id := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenant_id)

	_, err := objMgr.GetNetwork(network_view_name, cidr, nil)
	if err != nil {
		return fmt.Errorf("Getting Network from network view (%s) failed : %s", network_view_name, err)
	}

	return nil
}
func resourceNetworkUpdate(d *schema.ResourceData, m interface{}) error {
	//not supported by Infoblox Go Client for now
	return nil
}
func resourceNetworkDelete(d *schema.ResourceData, m interface{}) error {
	network_view_name := d.Get("network_view_name").(string)
	cidr := d.Get("cidr").(string)
	tenant_id := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenant_id)

	ref, err := objMgr.GetNetwork(network_view_name, cidr, nil)
	if err != nil {
		return fmt.Errorf("Getting Network failed from network view(%s) for deletion : %s", network_view_name, err)
	}

	_, err = objMgr.DeleteNetwork(ref.Ref, d.Get("network_view_name").(string))
	if err != nil {
		return fmt.Errorf("Deletion of Network failed from network view(%s) for deletion : %s", network_view_name, err)
	}
	d.SetId("")

	return nil
}
