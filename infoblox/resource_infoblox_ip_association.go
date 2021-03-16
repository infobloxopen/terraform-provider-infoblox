package infoblox

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client"
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
				Default:     "default",
				Description: "Network view name available in Nios server.",
			},
			"vm_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the vm.",
			},
			"cidr": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The address in cidr format.",
			},
			"ip_addr": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "IP address your instance in cloud.",
			},
			"mac_addr": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "mac address of your instance in cloud.",
			},
			"dns_view": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "view in which record has to be created.",
			},
			"zone": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "zone under which record has been created.",
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

//This method has an update call for the reason that,we are creating
//a reservation which doesnt have the details of the mac address
//at the beginig and we are using this update call to update the mac address
//of the record after the VM has been provisined.It is in the create method
//because for this resource we are doing association instead of allocation.
func resourceIPAssociationCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Association of IP address in specified network block", resourceIPAssociationIDString(d))

	if err := Resource(d, m); err != nil {
		return err
	}

	log.Printf("[DEBUG] %s:completing Association of IP address in specified network block", resourceIPAssociationIDString(d))
	return resourceIPAssociationRead(d, m)
}

func resourceIPAssociationUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s:update operation on Association of IP address in specified network block", resourceIPAssociationIDString(d))

	if err := Resource(d, m); err != nil {
		return err
	}

	log.Printf("[DEBUG] %s:completing updation on Association of IP address in specified network block", resourceIPAssociationIDString(d))
	return resourceIPAssociationRead(d, m)
}

func resourceIPAssociationRead(d *schema.ResourceData, m interface{}) error {
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

//we are updating the record with an empty mac address after the vm has been
//destroyed because if we implement the delete hostrecord method here then there
//will be a conflict of resources
func resourceIPAssociationDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Reassociation of IP address in specified network block", resourceIPAssociationIDString(d))
	matchClient := "MAC_ADDRESS"
	ipAddr := d.Get("ip_addr").(string)
	vmID := d.Get("vm_id").(string)
	vmName := d.Get("vm_name").(string)
	tenantID := d.Get("tenant_id").(string)
	zone := d.Get("zone").(string)
	dnsView := d.Get("dns_view").(string)

	connector := m.(*ibclient.Connector)

	ZeroMacAddr := "00:00:00:00:00:00"
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	if (zone != "" || len(zone) != 0) && (dnsView != "" || len(dnsView) != 0) {
		_, err := objMgr.UpdateHostRecord(d.Id(), ipAddr, ZeroMacAddr, vmID, vmName)
		if err != nil {
			return fmt.Errorf("Error Releasing IP from network block having reference (%s): %s", d.Id(), err)
		}
		d.SetId("")
	} else {
		_, err := objMgr.UpdateFixedAddress(d.Id(), matchClient, ZeroMacAddr, "", "")
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

	matchClient := "MAC_ADDRESS"
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

	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)
	//conversion from bit reversed EUI-48 format to hexadecimal EUI-48 format
	macAddr = strings.Replace(macAddr, "-", ":", -1)
	name := Name + "." + zone

	if (zone != "" || len(zone) != 0) && (dnsView != "" || len(dnsView) != 0) {
		hostRecordObj, err := objMgr.GetHostRecord(name)
		if err != nil {
			return fmt.Errorf("GetHostRecord failed from network block(%s):%s", cidr, err)
		}
		if hostRecordObj == nil {
			return fmt.Errorf("HostRecord %s not found.", name)
		}
		_, err = objMgr.UpdateHostRecord(hostRecordObj.Ref, ipAddr, macAddr, vmID, Name)
		if err != nil {
			return fmt.Errorf("UpdateHost Record error from network block(%s):%s", cidr, err)
		}
		d.SetId(hostRecordObj.Ref)
	} else {
		fixedAddressObj, err := objMgr.GetFixedAddress(networkViewName, cidr, ipAddr, "")
		if err != nil {
			return fmt.Errorf("GetFixedAddress error from network block(%s):%s", cidr, err)
		}
		if fixedAddressObj == nil {
			return fmt.Errorf("FixedAddress %s not found in network %s.", ipAddr, cidr)
		}

		_, err = objMgr.UpdateFixedAddress(fixedAddressObj.Ref, matchClient, macAddr, vmID, Name)
		if err != nil {
			return fmt.Errorf("UpdateFixedAddress error from network block(%s):%s", cidr, err)
		}
		d.SetId(fixedAddressObj.Ref)
	}
	return nil
}
