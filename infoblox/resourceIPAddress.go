package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/infobloxopen/infoblox-go-client"
	"log"
)

func resourceIPAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceIPAddressRequest,
		Read:   resourceIPAddressGet,
		Update: resourceIPAddressUpdate,
		Delete: resourceIPAddressRelease,

		Schema: map[string]*schema.Schema{
			"network_view_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("network_view_name", "default"),
				Description: "Network view name available in Nios server",
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
				Description: "Give the address in cidr format",
			},
			"ip_addr": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ipaddr", nil),
				Description: "IP address of your instance in cloud",
				Computed:    true,
			},
			"mac_addr": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("macaddr", nil),
				Description: "mac address of your instance in cloud",
			},
			"vm_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("vmid", nil),
				Description: "VM name",
			},
			"tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("tenant_id", nil),
				Description: "Unique identifier of your instance in cloud",
			},
			"gateway": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("tenant_id", nil),
				Description: "gateway ip address of your network block.First IPv4 address",
				Computed:    true,
			},
		},
	}
}

func resourceIPAddressRequest(d *schema.ResourceData, m interface{}) error {
	network_view_name := d.Get("network_view_name").(string)
	network_name := d.Get("network_name").(string)
	ip_addr := d.Get("ip_addr").(string)
	cidr := d.Get("cidr").(string)
	mac_addr := d.Get("mac_addr").(string)
	vm_id := d.Get("vm_id").(string)
	tenant_id := d.Get("tenant_id").(string)
	gateway := d.Get("gateway").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenant_id)

	gatewayIp, err := objMgr.AllocateIP(network_view_name, cidr, gateway, "", "")
	if err != nil {
		log.Printf("Gateway creation failed with error:'%s'", err)
	}

	ip_addr_obj, err := objMgr.AllocateIP(network_view_name, cidr, ip_addr, mac_addr, vm_id)
	if err != nil {
		return fmt.Errorf("Error allocating IP from Network(%s) : %s", network_name, err)
	}

	d.Set("gateway", gatewayIp.IPAddress)
	d.Set("ip_addr", ip_addr_obj.IPAddress)
	d.SetId(vm_id)

	return nil
}
func resourceIPAddressGet(d *schema.ResourceData, m interface{}) error {
	network_view_name := d.Get("network_view_name").(string)
	network_name := d.Get("network_name").(string)
	tenant_id := d.Get("tenant_id").(string)
	ip_addr := d.Get("ip_addr").(string)
	cidr := d.Get("cidr").(string)
	mac_addr := d.Get("mac_addr").(string)
	gateway := d.Get("gateway").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenant_id)

	_, err := objMgr.GetFixedAddress(network_view_name, cidr, gateway, "")
	if err != nil {
		log.Printf("Gateway creation failed with error:'%s'", err)
	}

	_, err = objMgr.GetFixedAddress(network_view_name, cidr, ip_addr, mac_addr)
	if err != nil {
		return fmt.Errorf("Error getting IP from network (%s) : %s", network_name, err)
	}

	return nil
}
func resourceIPAddressUpdate(d *schema.ResourceData, m interface{}) error {
	//Not Supported by Infoblox Go Client for now
	return nil
}
func resourceIPAddressRelease(d *schema.ResourceData, m interface{}) error {
	network_view_name := d.Get("network_view_name").(string)
	network_name := d.Get("network_name").(string)
	ip_addr := d.Get("ip_addr").(string)
	cidr := d.Get("cidr").(string)
	mac_addr := d.Get("mac_addr").(string)
	tenant_id := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenant_id)

	_, err := objMgr.ReleaseIP(network_view_name, cidr, ip_addr, mac_addr)
	if err != nil {
		return fmt.Errorf("Error Releasing IP from network(%s) : %s", network_name, err)
	}

	d.SetId("")

	return nil
}
