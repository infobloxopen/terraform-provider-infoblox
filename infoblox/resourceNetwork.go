package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/infobloxopen/infoblox-go-client"
	"log"
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
				Description: "Network view name available in NIOS Server.",
			},
			"network_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("networkName", nil),
				Description: "The name of your network block.",
			},
			"cidr": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("net_address", nil),
				Description: "Give the network block in cidr format.",
			},
			"tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("tenantID", nil),
				Description: "Unique identifier of your tenant in cloud.",
			},
			"reserve_ip": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("reserveIP", 0),
				Description: "The no of IP's you want to reserve.",
			},
			"gateway": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("gateway", nil),
				Description: "gateway ip address of your network block.By default first IPv4 address is set as gateway address.",
				Computed:    true,
			},
		},
	}
}

func resourceNetworkCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning network block Creation", resourceNetworkIDString(d))

	networkViewName := d.Get("network_view_name").(string)
	cidr := d.Get("cidr").(string)
	networkName := d.Get("network_name").(string)
	reserveIP := d.Get("reserve_ip").(int)
	gateway := d.Get("gateway").(string)
	tenantID := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)

	nwname, err := objMgr.CreateNetwork(networkViewName, cidr, networkName)
	if err != nil {
		return fmt.Errorf("Creation of network block failed in network view (%s) : %s", networkViewName, err)
	}

	gatewayIP, err := objMgr.GetFixedAddress(networkViewName, cidr, gateway, "")
	if err == nil && gatewayIP != nil {
		fmt.Printf("Gateway already created")
	} else if gatewayIP == nil {
		gatewayIP, err = objMgr.AllocateIP(networkViewName, cidr, gateway, "00:00:00:00:00:00", "", "")
		if err != nil {
			return fmt.Errorf("Gateway Creation failed in network block(%s) error: %s", cidr, err)
		}
	}

	for i := 1; i <= reserveIP; i++ {
		_, err = objMgr.AllocateIP(networkViewName, cidr, gateway, "00:00:00:00:00:00", "", "")
		if err != nil {
			return fmt.Errorf("Reservation in network block failed in network view(%s):%s", networkViewName, err)
		}
	}

	d.Set("gateway", gatewayIP.IPAddress)
	d.SetId(nwname.Ref)

	log.Printf("[DEBUG] %s: Creation on network block complete", resourceNetworkIDString(d))
	return nil
}
func resourceNetworkRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Reading the required network block", resourceNetworkIDString(d))

	networkViewName := d.Get("network_view_name").(string)
	tenantID := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)

	_, err := objMgr.GetNetworkwithref(d.Id())
	if err != nil {
		return fmt.Errorf("Getting Network block from network view (%s) failed : %s", networkViewName, err)
	}

	log.Printf("[DEBUG] %s: Completed reading network block", resourceNetworkIDString(d))
	return nil
}
func resourceNetworkUpdate(d *schema.ResourceData, m interface{}) error {
	//not supported by Infoblox Go Client for now
	return nil
}

func resourceNetworkDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Deletion of network block", resourceNetworkIDString(d))

	networkViewName := d.Get("network_view_name").(string)
	tenantID := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)

	_, err := objMgr.DeleteNetwork(d.Id(), d.Get("network_view_name").(string))
	if err != nil {
		return fmt.Errorf("Deletion of Network block failed from network view(%s) for deletion : %s", networkViewName, err)
	}
	d.SetId("")

	log.Printf("[DEBUG] %s: Deletion of network block complete", resourceNetworkIDString(d))
	return nil
}

type resourceNetworkIDStringInterface interface {
	Id() string
}

func resourceNetworkIDString(d resourceNetworkIDStringInterface) string {
	id := d.Id()
	if id == "" {
		id = "<new resource>"
	}
	return fmt.Sprintf("infoblox_ip_allocation (ID = %s)", id)
}

// Check whether gateway or ip address already allocated
