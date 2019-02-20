package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/infobloxopen/infoblox-go-client"
	"log"
)

func resourceARecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceARecordCreate,
		Read:   resourceARecordGet,
		Update: resourceARecordUpdate,
		Delete: resourceARecordDelete,

		Schema: map[string]*schema.Schema{
			"network_view_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("networkViewName", "default"),
				Description: "Network view name available in Nios server.",
			},
			"vm_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("hostName", nil),
				Description: "The name of the VM.",
			},
			"cidr": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("net_address", nil),
				Description: "Give the address in cidr format.",
			},
			"zone": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("zone", nil),
				Description: "Zone under which record has to be created .",
			},
			"dns_view": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("dns_view", nil),
				Description: "Dns View under which the zone has been created .",
			},
			"ip_addr": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ipaddr", nil),
				Description: "IP address you want to allocate yourinstance in cloud.If field is not specified , it akes next avaliable ip address",
				Computed:    true,
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
				DefaultFunc: schema.EnvDefaultFunc("tenantID", nil),
				Description: "Unique identifier of your tenant in cloud.",
			},
		},
	}
}

func resourceARecordCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning to create A record from  required network block", resourceARecordIDString(d))

	networkViewName := d.Get("network_view_name").(string)
	//This is for record Name
	Name := d.Get("vm_name").(string)
	ipAddr := d.Get("ip_addr").(string)
	cidr := d.Get("cidr").(string)
	vmID := d.Get("vm_id").(string)
	//This is for vm name
	vmName := d.Get("vm_name").(string)
	zone := d.Get("zone").(string)
	dnsView := d.Get("dns_view").(string)
	tenantID := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)
	name := Name + "." + zone
	recordA, err := objMgr.CreateARecord(networkViewName, dnsView, name, cidr, ipAddr, vmID, vmName)
	if err != nil {
		return fmt.Errorf("Error creating A Record from network block(%s): %s", cidr, err)
	}

	d.Set("recordName", name)
	d.SetId(recordA.Ref)

	log.Printf("[DEBUG] %s: Creation of A Record complete", resourceARecordIDString(d))
	return nil
}

func resourceARecordGet(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Begining to Get A Record", resourceARecordIDString(d))

	dnsView := d.Get("dns_view").(string)
	tenantID := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)

	_, err := objMgr.GetARecordByRef(d.Id())
	if err != nil {
		return fmt.Errorf("Getting A record failed from dns view (%s) : %s", dnsView, err)
	}

	log.Printf("[DEBUG] %s: Completed reading required A Record ", resourceARecordIDString(d))
	return nil
}

func resourceARecordUpdate(d *schema.ResourceData, m interface{}) error {
	//not supported by Infoblox Go Client for now
	return nil
}

func resourceARecordDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Deletion of A Record", resourceARecordIDString(d))

	dnsView := d.Get("dns_view").(string)
	tenantID := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)

	_, err := objMgr.DeleteARecord(d.Id())
	if err != nil {
		return fmt.Errorf("Deletion of A Record failed from dns view(%s) : %s", dnsView, err)
	}
	d.SetId("")

	log.Printf("[DEBUG] %s: Deletion of A Record complete", resourceARecordIDString(d))
	return nil
}

type resourceARecordIDStringInterface interface {
	Id() string
}

func resourceARecordIDString(d resourceARecordIDStringInterface) string {
	id := d.Id()
	if id == "" {
		id = "<new resource>"
	}
	return fmt.Sprintf("infoblox_a_record (ID = %s)", id)
}
