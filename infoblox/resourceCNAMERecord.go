package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/infobloxopen/infoblox-go-client"
	"log"
)

func resourceCNAMERecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceCNAMERecordCreate,
		Read:   resourceCNAMERecordGet,
		Update: resourceCNAMERecordUpdate,
		Delete: resourceCNAMERecordDelete,

		Schema: map[string]*schema.Schema{
			"zone": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("zone", nil),
				Description: "Zone under which record has to be created.",
			},
			"dns_view": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("dns_view", nil),
				Description: "Dns View under which the zone has been created.",
			},
			"canonical": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("canonical", nil),
				Description: "The Canonical name for the record.",
			},
			"alias": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("hostName", nil),
				Description: "The alias name for the record.",
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
func resourceCNAMERecordCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning to create CNAME record ", resourceCNAMERecordIDString(d))

	Zone := d.Get("zone").(string)
	dnsView := d.Get("dns_view").(string)
	Canonical := d.Get("canonical").(string) + "." + Zone
	Alias := d.Get("alias").(string) + "." + Zone
	tenantID := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)
	recordCNAME, err := objMgr.CreateCNAMERecord(Canonical, Alias, dnsView)
	if err != nil {
		return fmt.Errorf("Error creating CNAME Record : %s", err)
	}

	d.Set("recordName", Alias)
	d.SetId(recordCNAME.Ref)

	log.Printf("[DEBUG] %s: Creation of CNAME Record complete", resourceCNAMERecordIDString(d))
	return nil
}

func resourceCNAMERecordGet(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Begining to Get CNAME Record", resourceCNAMERecordIDString(d))

	dnsView := d.Get("dns_view").(string)
	tenantID := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)

	_, err := objMgr.GetCNAMERecordByRef(d.Id())
	if err != nil {
		return fmt.Errorf("Getting CNAME RECORD failed from dns view(%s) : %s", dnsView, err)
	}

	log.Printf("[DEBUG] %s: Completed reading required CNAME Record ", resourceCNAMERecordIDString(d))
	return nil
}

func resourceCNAMERecordUpdate(d *schema.ResourceData, m interface{}) error {
	//not supported by Infoblox Go Client for now
	return nil
}

func resourceCNAMERecordDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Deletion of CNAME Record", resourceCNAMERecordIDString(d))

	dnsView := d.Get("dns_view").(string)
	tenantID := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "terraform", tenantID)

	_, err := objMgr.DeleteCNAMERecord(d.Id())
	if err != nil {
		return fmt.Errorf("Deletion of CNAME Record failed from dns view %s", dnsView, err)
	}
	d.SetId("")

	log.Printf("[DEBUG] %s: Deletion of CNAME Record complete", resourceCNAMERecordIDString(d))
	return nil
}

type resourceCNAMERecordIDStringInterface interface {
	Id() string
}

func resourceCNAMERecordIDString(d resourceCNAMERecordIDStringInterface) string {
	id := d.Id()
	if id == "" {
		id = "<new resource>"
	}
	return fmt.Sprintf("infoblox_cname_record (ID = %s)", id)
}
