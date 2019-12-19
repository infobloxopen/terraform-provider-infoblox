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
				Default:     "default",
				Description: "Network view name available in Nios server.",
			},
			"vm_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the VM.",
			},
			"cidr": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The address in cidr format.",
			},
			"zone": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Zone under which host record has to be created.",
			},
			"enable_dns": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "flag that defines if the host reocrd is used for DNS or IPAM Purposes.",
			},
			"dns_view": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Dns View under which the zone has been created.",
			},
			"ip_addr": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IP address your instance in cloud.For static allocation ,set the field with valid IP. For dynamic allocation, leave this field empty.",
				Computed:    true,
			},
			"mac_addr": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "mac address of your instance in cloud.",
			},
			"vm_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "instance id.",
			},
			"tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique identifier of your tenant in cloud.",
			},
		},
	}
}

func resourceIPAllocationRequest(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning to request a next free IP from a required network block", resourceIPAllocationIDString(d))

	networkViewName := d.Get("network_view_name").(string)
	//This is for record Name
	recordName := d.Get("vm_name").(string)
	ipAddr := d.Get("ip_addr").(string)
	cidr := d.Get("cidr").(string)
	macAddr := d.Get("mac_addr").(string)
	//This is for EA's
	vmName := d.Get("vm_name").(string)
	vmID := d.Get("vm_id").(string)
	tenantID := d.Get("tenant_id").(string)
	zone := d.Get("zone").(string)
	enableDns := d.Get("enable_dns").(bool)
	dnsView := d.Get("dns_view").(string)

	connector := m.(*ibclient.Connector)
	ZeroMacAddr := "00:00:00:00:00:00"
	//fqdn
	name := recordName + "." + zone
	ea := make(ibclient.EA)
	if vmName != "" {
		ea["VM Name"] = vmName
	}
	if vmID != "" {
		ea["VM ID"] = vmID
	}
	if macAddr == "" {
		macAddr = ZeroMacAddr
	}
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	if (zone != "" || len(zone) != 0) && (dnsView != "" || len(dnsView) != 0) {
		hostAddressObj, err := objMgr.CreateHostRecord(enableDns, name, networkViewName, dnsView, cidr, ipAddr, macAddr, ea)
		if err != nil {
			return fmt.Errorf("Error allocating IP from network block(%s): %s", cidr, err)
		}
		d.Set("ip_addr", hostAddressObj.Ipv4Addrs[0].Ipv4Addr)
		d.SetId(hostAddressObj.Ref)
	} else {
		fixedAddressObj, err := objMgr.AllocateIP(networkViewName, cidr, ipAddr, macAddr, recordName, ea)
		if err != nil {
			return fmt.Errorf("Error allocating IP from network block(%s): %s", cidr, err)
		}
		d.Set("ip_addr", fixedAddressObj.IPAddress)
		d.SetId(fixedAddressObj.Ref)
	}
	log.Printf("[DEBUG] %s:completing Request of IP from required network block", resourceIPAllocationIDString(d))
	return resourceIPAllocationGet(d, m)
}

func resourceIPAllocationGet(d *schema.ResourceData, m interface{}) error {

	log.Printf("[DEBUG] %s:Reading the required IP from network block", resourceIPAllocationIDString(d))

	tenantID := d.Get("tenant_id").(string)
	cidr := d.Get("cidr").(string)
	zone := d.Get("zone").(string)
	dnsView := d.Get("dns_view").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	if (zone != "" || len(zone) != 0) && (dnsView != "" || len(dnsView) != 0) {
		obj, err := objMgr.GetHostRecordByRef(d.Id())
		if err != nil {
			return fmt.Errorf("Error getting IP from network block(%s): %s", cidr, err)
		}
		d.SetId(obj.Ref)
	} else {
		obj, err := objMgr.GetFixedAddressByRef(d.Id())
		if err != nil {
			return fmt.Errorf("Error getting IP from network block(%s): %s", cidr, err)
		}
		d.SetId(obj.Ref)
	}
	log.Printf("[DEBUG] %s: Completed Reading IP from the network block", resourceIPAllocationIDString(d))
	return nil
}

func resourceIPAllocationUpdate(d *schema.ResourceData, m interface{}) error {

	match_client := "MAC_ADDRESS"

	log.Printf("[DEBUG] %s: Updating the Parameters of the allocated IP in the specified network block", resourceIPAllocationIDString(d))

	macAddr := d.Get("mac_addr").(string)
	tenantID := d.Get("tenant_id").(string)
	vmID := d.Get("vm_id").(string)
	vmName := d.Get("vm_name").(string)
	zone := d.Get("zone").(string)
	dnsView := d.Get("dns_view").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	if (zone != "" || len(zone) != 0) && (dnsView != "" || len(dnsView) != 0) {
		hostRecordObj, _ := objMgr.GetHostRecordByRef(d.Id())
		IPAddrObj, _ := objMgr.GetIpAddressFromHostRecord(*hostRecordObj)
		obj, err := objMgr.UpdateHostRecord(d.Id(), IPAddrObj, macAddr, vmID, vmName)
		if err != nil {
			return fmt.Errorf("Error updating IP from network block having reference (%s): %s", d.Id(), err)
		}
		d.SetId(obj)
	} else {
		obj, err := objMgr.UpdateFixedAddress(d.Id(), match_client, macAddr, vmID, vmName)
		if err != nil {
			return fmt.Errorf("Error updating IP from network block having reference (%s): %s", d.Id(), err)
		}
		d.SetId(obj.Ref)
	}
	log.Printf("[DEBUG] %s: Updation of Parameters of allocated IP complete in the specified network block", resourceIPAllocationIDString(d))
	return resourceIPAllocationGet(d, m)
}

func resourceIPAllocationRelease(d *schema.ResourceData, m interface{}) error {

	log.Printf("[DEBUG] %s: Beginning Release of an allocated IP in the specified network block", resourceIPAllocationIDString(d))

	tenantID := d.Get("tenant_id").(string)
	zone := d.Get("zone").(string)
	dnsView := d.Get("dns_view").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)
	if (zone != "" || len(zone) != 0) && (dnsView != "" || len(dnsView) != 0) {
		_, err := objMgr.DeleteHostRecord(d.Id())
		if err != nil {
			return fmt.Errorf("Error Releasing IP from network block having reference (%s): %s", d.Id(), err)
		}
	} else {
		_, err := objMgr.DeleteFixedAddress(d.Id())
		if err != nil {
			return fmt.Errorf("Error Releasing IP from network block having reference (%s): %s", d.Id(), err)
		}
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
