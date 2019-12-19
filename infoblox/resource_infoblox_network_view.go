package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/infobloxopen/infoblox-go-client"
	"log"
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
				Default:     "default",
				Description: "Desired name of the view shown in NIOS appliance.",
			},
			"tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique identifier of your tenant in cloud.",
			},
		},
	}
}

func resourceNetworkViewCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning network view Creation", resourceNetworkViewIDString(d))

	tenantID := d.Get("tenant_id").(string)
	Connector := m.(*ibclient.Connector)
	objMgr := ibclient.NewObjectManager(Connector, "Terraform", tenantID)

	networkViewName, err := objMgr.CreateNetworkView(d.Get("network_view_name").(string))
	if err != nil {
		return fmt.Errorf("Failed to create Network View : %s", err)
	}

	d.SetId(networkViewName.Name)

	log.Printf("[DEBUG] %s: Completed network view Creation", resourceNetworkViewIDString(d))

	return resourceNetworkViewRead(d, m)
}
func resourceNetworkViewRead(d *schema.ResourceData, m interface{}) error {

	log.Printf("[DEBUG] %s: Beginning to get network view ", resourceNetworkViewIDString(d))

	tenantID := d.Get("tenant_id").(string)
	Connector := m.(*ibclient.Connector)
	objMgr := ibclient.NewObjectManager(Connector, "Terraform", tenantID)

	obj, err := objMgr.GetNetworkView(d.Id())
	if err != nil {
		return fmt.Errorf("Failed to get Network View : %s", err)
	}
	d.SetId(obj.Name)

	log.Printf("[DEBUG] %s: got Network View", resourceNetworkViewIDString(d))

	return nil
}
func resourceNetworkViewUpdate(d *schema.ResourceData, m interface{}) error {

	return fmt.Errorf("network view updation is not supported")
}
func resourceNetworkViewDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}

type resourceNetworkViewIDStringInterface interface {
	Id() string
}

func resourceNetworkViewIDString(d resourceNetworkViewIDStringInterface) string {
	id := d.Id()
	if id == "" {
		id = "<new resource>"
	}
	return fmt.Sprintf("infoblox_network_view(ID = %s)", id)
}
