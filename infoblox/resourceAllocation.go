package infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/infobloxopen/infoblox-go-client"
)

func resourceAllocation() *schema.Resource {
	return &schema.Resource{
		Create: resourceAllocationRequest,
		Read:   resourceAllocationGet,
		Update: resourceAllocationUpdate,
		Delete: resourceAllocationRelease,

		Schema: map[string]*schema.Schema{
			"networkviewname": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("net_view_name", nil),
				Description: "give the created network view name",
			},
			"network_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("network_name", nil),
				Description: "The name you have given to your network.",
			},
			"cidr": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("net_address", nil),
			},
			"ipaddr": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ipaddr", nil),
			},
			"macaddr": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("macaddr", nil),
				Description: "mac address of your instance",
			},
			"vmid": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("vmid", nil),
				Description: "VM name",
			},
		},
	}
}

func resourceAllocationRequest(d *schema.ResourceData, m interface{}) error {
	networkviewname := d.Get("networkviewname").(string)
	ipaddr :=d.Get("ipaddr").(string)
	cidr := d.Get("cidr").(string)
	macaddr := d.Get("macaddr").(string)
	vmID := d.Get("vmid").(string)
	connector := m.(*ibclient.Connector)
	objMgr := ibclient.NewObjectManager(connector, "terraform", "goclient1")
	objMgr.AllocateIP(networkviewname, cidr, ipaddr, macaddr, vmID)
	d.SetId(macaddr)
	return nil
}
func resourceAllocationGet(d *schema.ResourceData, m interface{}) error {
	networkviewname := d.Get("networkviewname").(string)
	ipaddr :=d.Get("ipaddr").(string)
	cidr := d.Get("cidr").(string)
	macaddr := d.Get("macaddr").(string)
	connector := m.(*ibclient.Connector)
	objMgr := ibclient.NewObjectManager(connector, "terraform", "goclient1")
	objMgr.GetFixedAddress(networkviewname, cidr, ipaddr, macaddr)
	d.SetId(macaddr)
	return nil
}
func resourceAllocationUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}
func resourceAllocationRelease(d *schema.ResourceData, m interface{}) error {
	networkviewname := d.Get("networkviewname").(string)
	ipaddr :=d.Get("ipaddr").(string)
	cidr := d.Get("cidr").(string)
	macaddr := d.Get("macaddr").(string)
	connector := m.(*ibclient.Connector)
	objMgr := ibclient.NewObjectManager(connector, "terraform", "goclient1")
	objMgr.ReleaseIP(networkviewname, cidr, ipaddr, macaddr)
	
	d.SetId("")
	return nil
}
