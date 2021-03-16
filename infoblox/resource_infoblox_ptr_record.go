package infoblox

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client"
)

func resourcePTRRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourcePTRRecordCreate,
		Read:   resourcePTRRecordGet,
		Update: resourcePTRRecordUpdate,
		Delete: resourcePTRRecordDelete,

		Schema: map[string]*schema.Schema{
			"vm_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the VM.",
			},
			"cidr": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The network to allocate IP address when the ip_addr field is empty. Network address in cidr format.",
			},
			"zone": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Zone under which record has to be created.",
			},
			"dns_view": &schema.Schema{
				Type:        schema.TypeString,
				Default:     "default",
				Optional:    true,
				Description: "Dns View under which the zone has been created.",
			},
			"ip_addr": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "IP address your instance in cloud. For static allocation, set the field with valid IP. For dynamic allocation, leave this field empty and set the cidr field.",
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

func resourcePTRRecordCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning to create PTR record from  required network block", resourcePTRRecordIDString(d))

	//This is for record Name
	recordName := d.Get("vm_name").(string)
	ipAddr := d.Get("ip_addr").(string)
	cidr := d.Get("cidr").(string)
	vmID := d.Get("vm_id").(string)
	vmName := d.Get("vm_name").(string)
	zone := d.Get("zone").(string)
	dnsView := d.Get("dns_view").(string)
	tenantID := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)

	ea := make(ibclient.EA)

	ea["VM Name"] = vmName

	if vmID != "" {
		ea["VM ID"] = vmID
	}

	if ipAddr == "" && cidr == "" {
		return fmt.Errorf("Error creating PTR record: nether ip_addr nor cidr value provided.")
	}

	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)
	//fqdn
	name := recordName + "." + zone
	recordPTR, err := objMgr.CreatePTRRecord(dnsView, dnsView, name, cidr, ipAddr, ea)
	if err != nil {
		return fmt.Errorf("Error creating PTR Record from network block(%s): %s", cidr, err)
	}

	d.Set("recordName", name)
	d.SetId(recordPTR.Ref)

	log.Printf("[DEBUG] %s: Creation of PTR Record complete", resourcePTRRecordIDString(d))
	return resourcePTRRecordGet(d, m)
}

func resourcePTRRecordGet(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Begining to Get PTR Record", resourcePTRRecordIDString(d))

	tenantID := d.Get("tenant_id").(string)
	dnsView := d.Get("dns_view").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	obj, err := objMgr.GetPTRRecordByRef(d.Id())
	if err != nil {
		return fmt.Errorf("Getting PTR Record from dns view (%s) failed : %s", dnsView, err)
	}
	d.SetId(obj.Ref)
	log.Printf("[DEBUG] %s: Completed reading required PTR Record ", resourcePTRRecordIDString(d))
	return nil
}

func resourcePTRRecordUpdate(d *schema.ResourceData, m interface{}) error {

	return fmt.Errorf("updating a PTR record is not supported")
}

func resourcePTRRecordDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Deletion of PTR Record", resourcePTRRecordIDString(d))

	tenantID := d.Get("tenant_id").(string)
	dnsView := d.Get("dns_view").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

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
