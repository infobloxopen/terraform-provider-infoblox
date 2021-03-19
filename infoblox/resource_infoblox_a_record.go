package infoblox

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client"
)

func resourceARecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceARecordCreate,
		Read:   resourceARecordGet,
		Update: resourceARecordUpdate,
		Delete: resourceARecordDelete,

		Schema: map[string]*schema.Schema{
			"vm_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the VM.",
			},
			"cidr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The network to allocate IP address when the ip_addr field is empty. Network address in cidr format.",
			},
			"zone": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Zone under which record has to be created.",
			},
			"dns_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
				Description: "Dns View under which the zone has been created.",
			},
			"ip_addr": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "IP address your instance in cloud. For static allocation, set the field with valid IP. For dynamic allocation, leave this field empty and set the cidr field.",
			},
			"vm_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "instance id.",
			},
			"tenant_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique identifier of your tenant in cloud.",
			},
		},
	}
}

func resourceARecordCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning to create A record from  required network block", resourceARecordIDString(d))

	// This is for record Name
	recordName := d.Get("vm_name").(string)
	ipAddr := d.Get("ip_addr").(string)
	cidr := d.Get("cidr").(string)
	vmID := d.Get("vm_id").(string)
	// This is for vm name
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
		return fmt.Errorf("Error creating A record: nether ip_addr nor cidr value provided.")
	}

	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)
	// fqdn
	name := recordName + "." + zone
	recordA, err := objMgr.CreateARecord(dnsView, dnsView, name, cidr, ipAddr, ea)
	if err != nil {
		return fmt.Errorf("Error creating A Record from network block(%s): %s", cidr, err)
	}

	d.Set("recordName", name)
	d.SetId(recordA.Ref)

	log.Printf("[DEBUG] %s: Creation of A Record complete", resourceARecordIDString(d))
	return resourceARecordGet(d, m)
}

func resourceARecordGet(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning to Get A Record", resourceARecordIDString(d))

	dnsView := d.Get("dns_view").(string)
	tenantID := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	obj, err := objMgr.GetARecordByRef(d.Id())
	if err != nil {
		return fmt.Errorf("Getting A record failed from dns view (%s) : %s", dnsView, err)
	}
	d.SetId(obj.Ref)
	log.Printf("[DEBUG] %s: Completed reading required A Record ", resourceARecordIDString(d))
	return nil
}

func resourceARecordUpdate(d *schema.ResourceData, m interface{}) error {
	return fmt.Errorf("updating A record is not supported")
}

func resourceARecordDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Deletion of A Record", resourceARecordIDString(d))

	dnsView := d.Get("dns_view").(string)
	tenantID := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

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
