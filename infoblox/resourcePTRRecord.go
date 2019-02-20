package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/infobloxopen/infoblox-go-client"
	"log"
)

func resourcePTRRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourcePTRRecordCreate,
		Read:   resourcePTRRecordGet,
		Update: resourcePTRRecordUpdate,
		Delete: resourcePTRRecordDelete,

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

func resourcePTRRecordCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning to create PTR record from  required network block", resourcePTRRecordIDString(d))

	networkViewName := d.Get("network_view_name").(string)
	//This is for record Name
	Name := d.Get("vm_name").(string)
	ipAddr := d.Get("ip_addr").(string)
	cidr := d.Get("cidr").(string)
	vmID := d.Get("vm_id").(string)
	vmName := d.Get("vm_name").(string)
	zone := d.Get("zone").(string)
	dnsView := d.Get("dns_view").(string)
	tenantID := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)
	name := Name + "." + zone
	recordPTR, err := objMgr.CreatePTRRecord(networkViewName, dnsView, name, cidr, ipAddr, vmID, vmName)
	if err != nil {
		return fmt.Errorf("Error creating PTR Record from network block(%s): %s", cidr, err)
	}

	d.Set("recordName", name)
	d.SetId(recordPTR.Ref)

	log.Printf("[DEBUG] %s: Creation of PTR Record complete", resourceARecordIDString(d))
	return nil
}

func resourcePTRRecordGet(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Begining to Get PTR Record", resourcePTRRecordIDString(d))

	tenantID := d.Get("tenant_id").(string)
	dnsView := d.Get("dns_view").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)

	_, err := objMgr.GetPTRRecordByRef(d.Id())
	if err != nil {
		return fmt.Errorf("Getting PTR Record from dns view (%s) failed : %s", dnsView, err)
	}

	log.Printf("[DEBUG] %s: Completed reading required PTR Record ", resourcePTRRecordIDString(d))
	return nil
}

func resourcePTRRecordUpdate(d *schema.ResourceData, m interface{}) error {
	//not supported by Infoblox Go Client for now
	return nil
}

func resourcePTRRecordDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Deletion of PTR Record", resourcePTRRecordIDString(d))

	tenantID := d.Get("tenant_id").(string)
	dnsView := d.Get("dns_view").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)

	_, err := objMgr.DeletePTRRecord(d.Id())
	if err != nil {
		return fmt.Errorf("Deletion of PTR Record failed from dns view(%s) : %s", dnsView, err)
	}
	d.SetId("")

	log.Printf("[DEBUG] %s: Deletion of PTR Record complete", resourcePTRRecordIDString(d))
	return nil
}

type resourcePTRRecordIDStringInterface interface {
	Id() string
}

func resourcePTRRecordIDString(d resourcePTRRecordIDStringInterface) string {
	id := d.Id()
	if id == "" {
		id = "<new resource>"
	}
	return fmt.Sprintf("infoblox_ptr_record (ID = %s)", id)
}
