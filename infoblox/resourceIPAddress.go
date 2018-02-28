package infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/infobloxopen/infoblox-go-client"
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
				Description: "Network view name available in Nios Appliance",
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
				Description: "IP address of your instance",
			},
			"mac_addr": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("macaddr", nil),
				Description: "mac address of your instance",
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
		},
	}
}

func resourceIPAddressRequest(d *schema.ResourceData, m interface{}) error {
	network_view_name := d.Get("network_view_name").(string)
	ip_addr := d.Get("ip_addr").(string)
	cidr := d.Get("cidr").(string)
	mac_addr := d.Get("mac_addr").(string)
	vmID := d.Get("vm_id").(string)
	tenant_id := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)
	objMgr := ibclient.NewObjectManager(connector, "terraform", tenant_id)
	objMgr.AllocateIP(network_view_name, cidr, ip_addr, mac_addr, vmID)
	d.SetId(mac_addr)
	return nil
}
func resourceIPAddressGet(d *schema.ResourceData, m interface{}) error {
	network_view_name := d.Get("network_view_name").(string)
	tenant_id := d.Get("tenant_id").(string)
	ip_addr := d.Get("ip_addr").(string)
	cidr := d.Get("cidr").(string)
	mac_addr := d.Get("mac_addr").(string)
	connector := m.(*ibclient.Connector)
	objMgr := ibclient.NewObjectManager(connector, "terraform", tenant_id)
	objMgr.GetFixedAddress(network_view_name, cidr, ip_addr, mac_addr)
	d.SetId(mac_addr)
	return nil
}
func resourceIPAddressUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}
func resourceIPAddressRelease(d *schema.ResourceData, m interface{}) error {
	network_view_name := d.Get("network_view_name").(string)
	ip_addr := d.Get("ip_addr").(string)
	cidr := d.Get("cidr").(string)
	mac_addr := d.Get("mac_addr").(string)
	tenant_id := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)
	objMgr := ibclient.NewObjectManager(connector, "terraform", tenant_id)
	objMgr.ReleaseIP(network_view_name, cidr, ip_addr, mac_addr)

	d.SetId("")
	return nil
}
