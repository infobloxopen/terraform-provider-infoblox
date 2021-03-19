package infoblox

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client"
)

func resourceNetworkView() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkViewCreate,
		Read:   resourceNetworkViewRead,
		Update: resourceNetworkViewUpdate,
		Delete: resourceNetworkViewDelete,

		Schema: map[string]*schema.Schema{
			"network_view_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Desired name of the view shown in NIOS appliance.",
			},
			"network_view_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_id": {
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
	d.Set("network_view_ref", obj.Ref)

	log.Printf("[DEBUG] %s: got Network View", resourceNetworkViewIDString(d))

	return nil
}

func resourceNetworkViewUpdate(d *schema.ResourceData, m interface{}) error {
	return fmt.Errorf("network view updation is not supported")
}

func resourceNetworkViewDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Deletion of network block", resourceNetworkIDString(d))

	networkViewName := d.Get("network_view_name").(string)
	networkViewRef := d.Get("network_view_ref").(string)
	tenantID := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	_, err := objMgr.DeleteNetworkView(networkViewRef)
	if err != nil {
		return fmt.Errorf("Deletion of Network view (%s) failed: %s", networkViewName, err)
	}
	d.SetId("")

	log.Printf("[DEBUG] %s: Deletion of network block complete", resourceNetworkViewIDString(d))
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
