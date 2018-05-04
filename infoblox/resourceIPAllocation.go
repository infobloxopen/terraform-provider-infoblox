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
				Description: "The name of the record.",
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
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("vmid", nil),
				Description: "instance name.",
			},
			"tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("tenantID", nil),
				Description: "Unique identifier of your tenant in cloud.",
			},
		},
	}
}

//We are using creation of host record for IPAM purposes for the reason that
//if we create a fixed address,t has mac address as a required parameter and
//a mac address is not generated before provisioning of vm hence we use create
//host record method used below
func resourceIPAllocationRequest(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning to request a next free IP from a required network block", resourceIPAllocationIDString(d))

	networkViewName := d.Get("network_view_name").(string)
	recordName := d.Get("host_name").(string)
	ipAddr := d.Get("ip_addr").(string)
	cidr := d.Get("cidr").(string)
	macAddr := d.Get("mac_addr").(string)
	vmID := d.Get("vm_id").(string)
	tenantID := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)

	hostrecordObj, err := objMgr.CreateHostRecordWithoutDNS(recordName, networkViewName, cidr, ipAddr, macAddr, vmID)
	if err != nil {
		return fmt.Errorf("Error allocating IP from network block(%s): %s", cidr, err)
	}

	d.Set("ip_addr", hostrecordObj.Ipv4Addrs[0].Ipv4Addr)
	d.SetId(hostrecordObj.Ref)

	log.Printf("[DEBUG] %s:completing Request of IP from required network block", resourceIPAllocationIDString(d))
	return nil
}

func resourceIPAllocationGet(d *schema.ResourceData, m interface{}) error {

	log.Printf("[DEBUG] %s:Reading the required IP from network block", resourceIPAllocationIDString(d))

	tenantID := d.Get("tenant_id").(string)
	cidr := d.Get("cidr").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)

	_, err := objMgr.GetHostRecord(d.Id())
	if err != nil {
		return fmt.Errorf("Error getting IP from network block(%s): %s", cidr, err)
	}

	log.Printf("[DEBUG] %s: Completed Reading IP from the network block", resourceIPAllocationIDString(d))
	return nil
}

func resourceIPAllocationUpdate(d *schema.ResourceData, m interface{}) error {

	log.Printf("[DEBUG] %s: Updating the Parameters of the allocated IP in the specified network block", resourceIPAllocationIDString(d))

	macAddr := d.Get("mac_addr").(string)
	tenantID := d.Get("tenant_id").(string)
	vmID := d.Get("vm_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)

	hostRecordObj, _ := objMgr.GetHostRecord(d.Id())
	IPAddrObj, _ := objMgr.GetIpAddressFromHostRecord(*hostRecordObj)
	objMgr.UpdateHostRecordWithoutDNS(d.Id(), IPAddrObj, macAddr, vmID)

	log.Printf("[DEBUG] %s: Updation of Parameters of allocated IP complete in the specified network block", resourceIPAllocationIDString(d))
	return nil
}

func resourceIPAllocationRelease(d *schema.ResourceData, m interface{}) error {

	log.Printf("[DEBUG] %s: Beginning Release of an allocated IP in the specified network block", resourceIPAllocationIDString(d))

	tenantID := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)
	_, err := objMgr.DeleteHostRecord(d.Id())
	if err != nil {
		return fmt.Errorf("Error Releasing IP from network block having reference (%s): %s", d.Id(), err)
	}

	d.SetId("")

	log.Printf("[DEBUG] %s: Finishing Release of allocated IP in the specified network block", resourceIPAllocationIDString(d))
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
