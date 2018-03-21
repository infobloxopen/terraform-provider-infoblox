package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/infobloxopen/infoblox-go-client"
	"log"
)

func resourceIPAllocation() *schema.Resource {
	return &schema.Resource{
		Create: resourceIPAllocationRequest,
		Read:   resourceIPAllocationGet,
		Update: resourceIPAllocationUpdate,
		Delete: resourceIPAllocationRelease,

		Schema: map[string]*schema.Schema{
			"network_view_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("network_view_name", "default"),
				Description: "Network view name available in Nios server.",
			},
			"network_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("network_name", nil),
				Description: "The name of the network.",
			},
			"cidr": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("net_address", nil),
				Description: "Give the address in cidr format.",
			},
			"ip_addr": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ipaddr", nil),
				Description: "IP address of your instance in cloud.",
				Computed:    true,
			},
			"mac_addr": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("macaddr", nil),
				Description: "mac address of your instance in cloud.",
			},
			"vm_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("vmid", nil),
				Description: "Virtual Machine name.",
			},
			"tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("tenant_id", nil),
				Description: "Unique identifier of your instance in cloud.",
			},
			"gateway": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("tenant_id", nil),
				Description: "gateway ip address of your network block.First IPv4 address.",
				Computed:    true,
			},
		},
	}
}

func resourceIPAllocationRequest(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning to request a next free IP from a required network block", resourceIPAllocationIDString(d))

	network_view_name := d.Get("network_view_name").(string)
	//network_name := d.Get("network_name").(string)
	ip_addr := d.Get("ip_addr").(string)
	cidr := d.Get("cidr").(string)
	mac_addr := d.Get("mac_addr").(string)
	vm_id := d.Get("vm_id").(string)
	tenant_id := d.Get("tenant_id").(string)
	gateway := d.Get("gateway").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenant_id)

	// Check whether gateway or ip address already allocated
	gatewayIp, err := objMgr.GetFixedAddress(network_view_name, cidr, gateway, "")
	if err == nil && gatewayIp != nil {
		fmt.Printf("Gateway alreadt created")
	} else if gatewayIp == nil {
		gatewayIp, err = objMgr.AllocateIP(network_view_name, cidr, gateway, "00:00:00:00:00:00", "")
		if err != nil {
			return fmt.Errorf("Gateway Creation failed in network block(%s) error: %s", cidr, err)
		}
	}

	ip_addr_obj, err := objMgr.AllocateIP(network_view_name, cidr, ip_addr, mac_addr, vm_id)
	if err != nil {
		return fmt.Errorf("Error allocating IP from network block(%s): %s", cidr, err)
	}

	d.Set("gateway", gatewayIp.IPAddress)
	d.Set("ip_addr", ip_addr_obj.IPAddress)
	// TODO what happens in case of a VM have 2 network interfaces.
	d.SetId(vm_id)

	log.Printf("[DEBUG] %s:completing Request of IP from  required network block", resourceIPAllocationIDString(d))
	return nil
}

func resourceIPAllocationGet(d *schema.ResourceData, m interface{}) error {

	log.Printf("[DEBUG] %s:Reading the required IP from  network block", resourceIPAllocationIDString(d))

	network_view_name := d.Get("network_view_name").(string)
	//network_name := d.Get("network_name").(string)
	tenant_id := d.Get("tenant_id").(string)
	ip_addr := d.Get("ip_addr").(string)
	cidr := d.Get("cidr").(string)
	mac_addr := d.Get("mac_addr").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenant_id)

	_, err := objMgr.GetFixedAddress(network_view_name, cidr, ip_addr, mac_addr)
	if err != nil {
		return fmt.Errorf("Error getting IP from network block(%s): %s", cidr, err)
	}

	log.Printf("[DEBUG] %s: Completed Reading IP from the network block", resourceIPAllocationIDString(d))
	return nil
}

func resourceIPAllocationUpdate(d *schema.ResourceData, m interface{}) error {

	log.Printf("[DEBUG] %s: Updating the Parameters of the allocated IP", resourceIPAllocationIDString(d))

	network_view_name := d.Get("network_view_name").(string)
	network_name := d.Get("network_name").(string)
	ip_addr := d.Get("ip_addr").(string)
	cidr := d.Get("cidr").(string)
	mac_addr := d.Get("mac_addr").(string)
	vm_id := d.Get("vm_id").(string)
	tenant_id := d.Get("tenant_id").(string)

	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenant_id)

	ref, err := objMgr.GetFixedAddress(network_view_name, network_name, ip_addr, "")
	if err != nil {
		return fmt.Errorf("GetFixedAddress error from network(%s):%s", cidr, err)
	}

	_, err = objMgr.UpdateFixedAddress(ref.Ref, mac_addr, vm_id)
	if err != nil {
		return fmt.Errorf("UpdateFixedAddress error from network block(%s):%s", cidr, err)
	}

	log.Printf("[DEBUG] %s: Updation of Parameters of allocated IP complete", resourceIPAllocationIDString(d))
	return nil
}

func resourceIPAllocationRelease(d *schema.ResourceData, m interface{}) error {

	log.Printf("[DEBUG] %s: Beginning Release of an allocated IP", resourceIPAllocationIDString(d))

	network_view_name := d.Get("network_view_name").(string)
	ip_addr := d.Get("ip_addr").(string)
	cidr := d.Get("cidr").(string)
	mac_addr := d.Get("mac_addr").(string)
	tenant_id := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenant_id)

	_, err := objMgr.ReleaseIP(network_view_name, cidr, ip_addr, mac_addr)
	if err != nil {
		return fmt.Errorf("Error Releasing IP from network(%s): %s", cidr, err)
	}

	d.SetId("")

	log.Printf("[DEBUG] %s: Finishing Release of allocated IP", resourceIPAllocationIDString(d))
	return nil
}

type resourceIPAllocationIDStringInterface interface {
	Id() string
}

func resourceIPAllocationIDString(d resourceIPAllocationIDStringInterface) string {
	id := d.Id()
	if id == "" {
		id = "<new resource>"
	}
	return fmt.Sprintf("infoblox_ip_allocation (ID = %s)", id)
}
