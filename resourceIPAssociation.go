package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/infobloxopen/infoblox-go-client"
	"log"
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
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("vmid", nil),
				Description: "Virtual Machine name.",
			},
			"tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("tenant_id", nil),
				Description: "Unique identifier of your instance in cloud.",
			},
		},
	}
}

func resourceIPAssociationCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Association of IP address", resourceIPAssociationIDString(d))

	networkViewName := d.Get("network_view_name").(string)
	recordName := d.Get("host_name").(string)
	ipAddr := d.Get("ip_addr").(string)
	cidr := d.Get("cidr").(string)
	macAddr := d.Get("mac_addr").(string)
	tenantID := d.Get("tenant_id").(string)

	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)

	hostRecordObj, err := objMgr.GetRecordHostWithoutDNS(recordName, networkViewName, cidr, ipAddr)
	if err != nil {
		return fmt.Errorf("GetHostAddress error from network block(%s):%s", cidr, err)
	}
	_, err = objMgr.UpdateRecordHostWithoutDNS(hostRecordObj.Ref, ipAddr, macAddr)
	if err != nil {
		return fmt.Errorf("UpdateHostAddress error from network block(%s):%s", cidr, err)
	}

	d.SetId(hostRecordObj.Ref)
	log.Printf("[DEBUG] %s:completing Association of IP address ", resourceIPAssociationIDString(d))
	return nil
}

func resourceIPAssociationUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceIPAssociationRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceIPAssociationDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Reassociation of IP address", resourceIPAssociationIDString(d))

	ipAddr := d.Get("ip_addr").(string)
	cidr := d.Get("cidr").(string)
	tenantID := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)

	_, err := objMgr.UpdateRecordHostWithoutDNS(d.Id(), ipAddr, "00:00:00:00:00:00")
	if err != nil {
		return fmt.Errorf("Error Releasing IP from network(%s): %s", cidr, err)
	}

	d.SetId("")

	log.Printf("[DEBUG] %s: Finishing Release of allocated IP", resourceIPAssociationIDString(d))

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
