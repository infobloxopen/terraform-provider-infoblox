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
				DefaultFunc: schema.EnvDefaultFunc("networkViewName", "default"),
				Description: "Network view name available in Nios server.",
			},
			"network_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("networkName", nil),
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
				DefaultFunc: schema.EnvDefaultFunc("tenantID", nil),
				Description: "Unique identifier of your instance in cloud.",
			},
			"gateway": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("tenantID", nil),
				Description: "gateway ip address of your network block.First IPv4 address.",
				Computed:    true,
			},
		},
	}
}

func resourceIPAllocationRequest(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning to request a next free IP from a required network block", resourceIPAllocationIDString(d))

	networkViewName := d.Get("network_view_name").(string)
	ipAddr := d.Get("ip_addr").(string)
	cidr := d.Get("cidr").(string)
	macAddr := d.Get("mac_addr").(string)
	vmID := d.Get("vm_id").(string)
	tenantID := d.Get("tenant_id").(string)
	gateway := d.Get("gateway").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)

	// Check whether gateway or ip address already allocated
	gatewayIP, err := objMgr.GetFixedAddress(networkViewName, cidr, gateway, "")
	if err == nil && gatewayIP != nil {
		fmt.Printf("Gateway alreadt created")
	} else if gatewayIP == nil {
		gatewayIP, err = objMgr.AllocateIP(networkViewName, cidr, gateway, "00:00:00:00:00:00", "")
		if err != nil {
			return fmt.Errorf("Gateway Creation failed in network block(%s) error: %s", cidr, err)
		}
	}

	ipAddrObj, err := objMgr.AllocateIP(networkViewName, cidr, ipAddr, macAddr, vmID)
	if err != nil {
		return fmt.Errorf("Error allocating IP from network block(%s): %s", cidr, err)
	}

	d.Set("gateway", gatewayIP.IPAddress)
	d.Set("ipAddr", ipAddrObj.IPAddress)
	// TODO what happens in case of a VM have 2 network interfaces.
	d.SetId(vmID)

	log.Printf("[DEBUG] %s:completing Request of IP from  required network block", resourceIPAllocationIDString(d))
	return nil
}

func resourceIPAllocationGet(d *schema.ResourceData, m interface{}) error {

	log.Printf("[DEBUG] %s:Reading the required IP from  network block", resourceIPAllocationIDString(d))

	networkViewName := d.Get("network_view_name").(string)
	tenantID := d.Get("tenant_id").(string)
	ipAddr := d.Get("ip_addr").(string)
	cidr := d.Get("cidr").(string)
	macAddr := d.Get("mac_addr").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)

	_, err := objMgr.GetFixedAddress(networkViewName, cidr, ipAddr, macAddr)
	if err != nil {
		return fmt.Errorf("Error getting IP from network block(%s): %s", cidr, err)
	}

	log.Printf("[DEBUG] %s: Completed Reading IP from the network block", resourceIPAllocationIDString(d))
	return nil
}

func resourceIPAllocationUpdate(d *schema.ResourceData, m interface{}) error {

	log.Printf("[DEBUG] %s: Updating the Parameters of the allocated IP", resourceIPAllocationIDString(d))

	networkViewName := d.Get("network_view_ame").(string)
	networkName := d.Get("network_name").(string)
	ipAddr := d.Get("ip_addr").(string)
	cidr := d.Get("cidr").(string)
	macAddr := d.Get("mac_addr").(string)
	vmID := d.Get("vm_id").(string)
	tenantID := d.Get("tenant_id").(string)

	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)

	ref, err := objMgr.GetFixedAddress(networkViewName, networkName, ipAddr, "")
	if err != nil {
		return fmt.Errorf("GetFixedAddress error from network(%s):%s", cidr, err)
	}

	_, err = objMgr.UpdateFixedAddress(ref.Ref, macAddr, vmID)
	if err != nil {
		return fmt.Errorf("UpdateFixedAddress error from network block(%s):%s", cidr, err)
	}

	log.Printf("[DEBUG] %s: Updation of Parameters of allocated IP complete", resourceIPAllocationIDString(d))
	return nil
}

func resourceIPAllocationRelease(d *schema.ResourceData, m interface{}) error {

	log.Printf("[DEBUG] %s: Beginning Release of an allocated IP", resourceIPAllocationIDString(d))

	networkViewName := d.Get("network_view_name").(string)
	ipAddr := d.Get("ip_addr").(string)
	cidr := d.Get("cidr").(string)
	macAddr := d.Get("mac_addr").(string)
	tenantID := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)

	_, err := objMgr.ReleaseIP(networkViewName, cidr, ipAddr, macAddr)
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
