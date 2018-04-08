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
			"host_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("hostName", nil),
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
				Optional:    true,
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
	recordName := d.Get("host_name").(string)
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

	hostrecordObj, err := objMgr.CreateRecordHostWithoutDNS(recordName, networkViewName, cidr, ipAddr, macAddr, vmID)
	if err != nil {
		return fmt.Errorf("Error allocating IP from network block(%s): %s", cidr, err)
	}

	d.Set("gateway", gatewayIP.IPAddress)
	d.Set("ip_addr", hostrecordObj.Ipv4Addrs[0].Ipv4Addr)
	// TODO what happens in case of a VM have 2 network interfaces.
	d.SetId(hostrecordObj.Ref)

	log.Printf("[DEBUG] %s:completing Request of IP from  required network block", resourceIPAllocationIDString(d))
	return nil
}

func resourceIPAllocationGet(d *schema.ResourceData, m interface{}) error {

	log.Printf("[DEBUG] %s:Reading the required IP from  network block", resourceIPAllocationIDString(d))

	tenantID := d.Get("tenant_id").(string)
	cidr := d.Get("cidr").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)

	_, err := objMgr.GetRecordHost(d.Id())
	if err != nil {
		return fmt.Errorf("Error getting IP from network block(%s): %s", cidr, err)
	}

	log.Printf("[DEBUG] %s: Completed Reading IP from the network block", resourceIPAllocationIDString(d))
	return nil
}

func resourceIPAllocationUpdate(d *schema.ResourceData, m interface{}) error {

	log.Printf("[DEBUG] %s: Updating the Parameters of the allocated IP", resourceIPAllocationIDString(d))

	macAddr := d.Get("mac_addr").(string)
	tenantID := d.Get("tenant_id").(string)

	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)

	hostRecordObj, _ := objMgr.GetRecordHost(d.Id())
	IPAddrObj, _ := objMgr.GetIpAddressFromRecordHost(*hostRecordObj)
	objMgr.UpdateRecordHostWithoutDNS(d.Id(), IPAddrObj, macAddr)

	log.Printf("[DEBUG] %s: Updation of Parameters of allocated IP complete", resourceIPAllocationIDString(d))
	return nil
}

func resourceIPAllocationRelease(d *schema.ResourceData, m interface{}) error {

	log.Printf("[DEBUG] %s: Beginning Release of an allocated IP", resourceIPAllocationIDString(d))

	cidr := d.Get("cidr").(string)
	tenantID := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)
	_, err := objMgr.DeleteRecordHost(d.Id())
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
