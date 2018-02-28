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
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("network_view_name", "default"),
				Description: "Desired name of the view shown in NIOS appliance",
			},
			"tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("tennant_id", nil),
				Description: "Unique identifier of your instace in cloud",
			},
		},
	}
}

func resourceNetworkViewCreate(d *schema.ResourceData, m interface{}) error {
	tenant_id := d.Get("tenant_id").(string)
	Connector := m.(*ibclient.Connector)
	objMgr := ibclient.NewObjectManager(Connector, "terraform", tenant_id)

	network_view_name, err := objMgr.CreateNetworkView(d.Get("network_view_name").(string))
	if err != nil {
		return fmt.Errorf("Failed to create Network View : %s", err)
	}

	d.SetId(network_view_name.Name)

	return nil
}
func resourceNetworkViewRead(d *schema.ResourceData, m interface{}) error {
	// Not Supported by Infoblox Go Client for Now

	return nil
}
func resourceNetworkViewUpdate(d *schema.ResourceData, m interface{}) error {
	// Not Supported by Infoblox Go Client for Now
	return nil
}
func resourceNetworkViewDelete(d *schema.ResourceData, m interface{}) error {
	// Not Supported by Infoblox Go Client for Now
	return nil
}
