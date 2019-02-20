package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/infobloxopen/infoblox-go-client"
	"log"
	"strings"
)

func resourceIPAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceIPAssociationCreate,
		Update: resourceIPAssociationUpdate,
		Delete: resourceIPAssociationDelete,
		Read:   resourceIPAssociationRead,

		Schema: map[string]*schema.Schema{
			"network_view_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("network_view_name", "default"),
				Description: "Network view name available in Nios server.",
			},
			"vm_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("host_name", nil),
				Description: "The name of the vm.",
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
			"dns_view": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("dns_view", nil),
				Description: "view in which record has to be created.",
			},
			"zone": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("zone", nil),
				Description: "zone under which record has been created.",
			},
			"vm_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("vmid", nil),
				Description: "instance id.",
			},
			"tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("tenant_id", nil),
				Description: "Unique identifier of your tenant in cloud.",
			},
		},
	}
}

//This method has an update call for the reason that,we are creating
//a reservation which doesnt have the details of the mac address
//at the beginig and we are using this update call to update the mac address
//of the record after the VM has been provisined.It is in the create method
//because for this resource we are doing association instead of allocation.
func resourceIPAssociationCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Association of IP address in specified network block", resourceIPAssociationIDString(d))

	Resource(d, m)

	log.Printf("[DEBUG] %s:completing Association of IP address in specified network block", resourceIPAssociationIDString(d))
	return nil
}

func resourceIPAssociationUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s:update operation on Association of IP address in specified network block", resourceIPAssociationIDString(d))

	Resource(d, m)

	log.Printf("[DEBUG] %s:completing updation on Association of IP address in specified network block", resourceIPAssociationIDString(d))
	return nil
}

func resourceIPAssociationRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

//we are updating the record with an empty mac address after the vm has been
//destroyed because if we implement the delete hostrecord method here then there
//will be a conflict of resources
func resourceIPAssociationDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Reassociation of IP address in specified network block", resourceIPAssociationIDString(d))
	match_client := "MAC_ADDRESS"
	ipAddr := d.Get("ip_addr").(string)
	vmID := d.Get("vm_id").(string)
	vmName := d.Get("vm_name").(string)
	tenantID := d.Get("tenant_id").(string)
	zone := d.Get("zone").(string)
	dnsView := d.Get("dns_view").(string)

	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)

	if (zone != "" || len(zone) != 0) && (dnsView != "" || len(dnsView) != 0) {
		_, err := objMgr.UpdateHostRecord(d.Id(), ipAddr, "00:00:00:00:00:00", vmID, vmName)
		if err != nil {
			return fmt.Errorf("Error Releasing IP from network block having reference (%s): %s", d.Id(), err)
		}
		d.SetId("")
	} else {
		_, err := objMgr.UpdateFixedAddress(d.Id(), match_client, "00:00:00:00:00:00", "", "")
		if err != nil {
			return fmt.Errorf("Error Releasing IP from network block having reference (%s): %s", d.Id(), err)
		}
		d.SetId("")
	}
	log.Printf("[DEBUG] %s: Finishing Release of allocated IP in specified network block", resourceIPAssociationIDString(d))

	return nil
}

type resourceIPAssociationIDStringInterface interface {
	Id() string
}

func resourceIPAssociationIDString(d resourceIPAssociationIDStringInterface) string {
	id := d.Id()
	if id == "" {
		id = "<new resource>"
	}
	return fmt.Sprintf("infoblox_mac_allocation (ID = %s)", id)
}

func Resource(d *schema.ResourceData, m interface{}) error {

	match_client := "MAC_ADDRESS"
	networkViewName := d.Get("network_view_name").(string)
	Name := d.Get("vm_name").(string)
	ipAddr := d.Get("ip_addr").(string)
	cidr := d.Get("cidr").(string)
	macAddr := d.Get("mac_addr").(string)
	tenantID := d.Get("tenant_id").(string)
	vmID := d.Get("vm_id").(string)
	zone := d.Get("zone").(string)
	dnsView := d.Get("dns_view").(string)

	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)
	//conversion from bit reversed EUI-48 format to hexadecimal EUI-48 format
	macAddr = strings.Replace(macAddr, "-", ":", -1)

	if (zone != "" || len(zone) != 0) && (dnsView != "" || len(dnsView) != 0) {
		name := Name + "." + zone
		hostRecordObj, err := objMgr.GetHostRecord(name, networkViewName, cidr, ipAddr)
		if err != nil {
			return fmt.Errorf("GetHostRecord failed from network block(%s):%s", cidr, err)
		}
		_, err = objMgr.UpdateHostRecord(hostRecordObj.Ref, ipAddr, macAddr, vmID, Name)
		if err != nil {
			return fmt.Errorf("UpdateFixedAddress error from network block(%s):%s", cidr, err)
		}
		d.SetId(hostRecordObj.Ref)
	} else {
		fixedAddressObj, err := objMgr.GetFixedAddress(networkViewName, cidr, ipAddr, "")
		if err != nil {
			return fmt.Errorf("GetFixedAddress error from network block(%s):%s", cidr, err)
		}

		_, err = objMgr.UpdateFixedAddress(fixedAddressObj.Ref, match_client, macAddr, vmID, Name)
		if err != nil {
			return fmt.Errorf("UpdateFixedAddress error from network block(%s):%s", cidr, err)
		}
		d.SetId(fixedAddressObj.Ref)
	}
	return nil
}
