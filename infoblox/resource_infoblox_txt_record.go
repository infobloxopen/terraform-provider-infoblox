package infoblox

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client"
)

func resourceTXTRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceTXTRecordCreate,
		Read:   resourceTXTRecordGet,
		Update: resourceTXTRecordUpdate,
		Delete: resourceTXTRecordDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the TXT record.",
			},
			"text": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The text for the record to contain.",
			},
			"ttl": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The TTL for the record.",
				Default:     0,
			},
			"dns_view": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
				Description: "Dns View under which the zone has been created.",
			},
			"tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique identifier of your tenant in cloud.",
			},
		},
	}
}

func resourceTXTRecordCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning to create TXT record from  required network block", resourceTXTRecordIDString(d))

	//This is for record Name
	recordName := d.Get("name").(string)
	text := d.Get("text").(string)
	ttl := d.Get("ttl").(int)
	dnsView := d.Get("dns_view").(string)
	tenantID := d.Get("tenant_id").(string)

	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)
	// fqdn

	recordTXT, err := objMgr.CreateTXTRecord(recordName, text, ttl, dnsView)

	if err != nil {
		return fmt.Errorf("Error creating TXT record: %s", err)
	}

	d.SetId(recordTXT.Ref)

	log.Printf("[DEBUG] %s: Creation of TXT Record complete", resourceTXTRecordIDString(d))
	return resourceTXTRecordGet(d, m)
}

func resourceTXTRecordGet(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Begining to Get TXT Record", resourceTXTRecordIDString(d))

	dnsView := d.Get("dns_view").(string)
	tenantID := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	obj, err := objMgr.GetTXTRecordByRef(d.Id())
	if err != nil {
		return fmt.Errorf("Getting TXT record failed from dns view (%s) : %s", dnsView, err)
	}
	d.SetId(obj.Ref)
	log.Printf("[DEBUG] %s: Completed reading required TXT Record ", resourceTXTRecordIDString(d))
	return nil
}

func resourceTXTRecordUpdate(d *schema.ResourceData, m interface{}) error {

	return fmt.Errorf("updating TXT record is not supported")
}

func resourceTXTRecordDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Deletion of TXT Record", resourceTXTRecordIDString(d))

	dnsView := d.Get("dns_view").(string)
	tenantID := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	_, err := objMgr.DeleteTXTRecord(d.Id())
	if err != nil {
		return fmt.Errorf("Deletion of TXT Record failed from dns view(%s) : %s", dnsView, err)
	}
	d.SetId("")

	log.Printf("[DEBUG] %s: Deletion of TXT Record complete", resourceTXTRecordIDString(d))
	return nil
}

type resourceTXTRecordIDStringInterface interface {
	Id() string
}

func resourceTXTRecordIDString(d resourceTXTRecordIDStringInterface) string {
	id := d.Id()
	if id == "" {
		id = "<new resource>"
	}
	return fmt.Sprintf("infoblox_txt_record (ID = %s)", id)
}
