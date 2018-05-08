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
			"host_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("host_name", nil),
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
				DefaultFunc: schema.EnvDefaultFunc("tenant_id", nil),
				Description: "Unique identifier of your tenant in cloud.",
			},
		},
	}
}

//This method has an update call for the reason that,we are creating
//a host record which doesnt have the details of the mac address
//at the beginigand we are using this update call to update the mac address
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

	ipAddr := d.Get("ip_addr").(string)
	tenantID := d.Get("tenant_id").(string)

	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)

	_, err := objMgr.UpdateHostRecordWithoutDNS(d.Id(), ipAddr, "00:00:00:00:00:00", "")
	if err != nil {
		return fmt.Errorf("Error Releasing IP from network block having reference (%s): %s", d.Id(), err)
	}

	d.SetId("")

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

	networkViewName := d.Get("network_view_name").(string)
	recordName := d.Get("host_name").(string)
	ipAddr := d.Get("ip_addr").(string)
	cidr := d.Get("cidr").(string)
	macAddr := d.Get("mac_addr").(string)
	tenantID := d.Get("tenant_id").(string)
	vmID := d.Get("vm_id").(string)

	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)
	//conversion from bit reversed EUI-48 format to hexadecimal EUI-48 format
	macAddr = strings.Replace(macAddr, "-", ":", -1)

	hostRecordObj, err := objMgr.GetHostRecordWithoutDNS(recordName, networkViewName, cidr, ipAddr)
	if err != nil {
		return fmt.Errorf("GetHostAddress error from network block(%s):%s", cidr, err)
	}
	_, err = objMgr.UpdateHostRecordWithoutDNS(hostRecordObj.Ref, ipAddr, macAddr, vmID)
	if err != nil {
		return fmt.Errorf("UpdateHostAddress error from network block(%s):%s", cidr, err)
	}

	d.SetId(hostRecordObj.Ref)
	return nil
}
