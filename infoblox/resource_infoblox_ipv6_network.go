package infoblox

import (
	"fmt"
	"log"

	ibclient "github.com/anagha-infoblox/infoblox-go-client"
	"github.com/hashicorp/terraform/helper/schema"
	//ibclient "github.com/infobloxopen/infoblox-go-client"
)

func resourceIPv6Network() *schema.Resource {
	return &schema.Resource{
		Create: resourceIPv6NetworkCreate,
		Read:   resourceIPv6NetworkRead,
		Update: resourceIPv6NetworkUpdate,
		Delete: resourceIPv6NetworkDelete,

		Schema: map[string]*schema.Schema{
			"network_view_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
				Description: "Network view name available in NIOS Server.",
			},
			"network_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of your network block.",
			},
			"cidr": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The network block in cidr format.",
			},
			"tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique identifier of your tenant in cloud.",
			},
			"allocate_prefix_len": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Set parameter value>0 to allocate next available network with prefix=value from network container defined by parent_cidr.",
			},
			"parent_cidr": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The parent network container block in cidr format to allocate from.",
			},
			"extensible_attributes": &schema.Schema{
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "The Extensible attributes to be updated on the Network",
			},
		},
	}

}

func resourceIPv6NetworkCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning IPv6 network block creation", resourceIPv6NetworkIDString(d))

	networkViewName := d.Get("network_view_name").(string)
	cidr := d.Get("cidr").(string)
	parentCidr := d.Get("parent_cidr").(string)
	networkName := d.Get("network_name").(string)
	tenantID := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)
	prefixLen := d.Get("allocate_prefix_len").(int)

	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)
	//ea := make(ibclient.EA)

	var IPv6Network *ibclient.IPv6Network
	var err error

	if cidr == "" && parentCidr != "" && prefixLen > 1 {
		IPv6Network, err = objMgr.AllocateIPv6Network(networkViewName, parentCidr, uint(prefixLen), networkName)
		if err != nil {
			return fmt.Errorf("Allocation of IPv6 network block failed in network view (%s) : %s", networkViewName, err)
		}
		d.Set("cidr", IPv6Network.Cidr)
	} else if cidr != "" {
		IPv6Network, err = objMgr.CreateIPv6Network(networkViewName, cidr, networkName)
		if err != nil {
			return fmt.Errorf("Creation of IPv6 network block failed in network view (%s) : %s", networkViewName, err)
		}
	} else {
		return fmt.Errorf("Creation of IPv6 network block failed: neither cidr nor parent_cidr with allocate_prefix_len was specified")
	}

	d.SetId(IPv6Network.Ref)

	log.Printf("[DEBUG] %s: Creation on network block complete", resourceIPv6NetworkIDString(d))
	return resourceIPv6NetworkRead(d, m)
}

func resourceIPv6NetworkRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Reading the required network block", resourceIPv6NetworkIDString(d))

	networkViewName := d.Get("network_view_name").(string)
	tenantID := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)
	obj, err := objMgr.GetIPv6NetworkWithRef(d.Id())
	if err != nil {
		return fmt.Errorf("Getting Network block from network view (%s) failed : %s", networkViewName, err)
	}
	d.SetId(obj.Ref)
	log.Printf("[DEBUG] %s: Completed reading network block", resourceNetworkIDString(d))
	return nil
}

func resourceIPv6NetworkUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Updating IPv6 network block EAs", resourceIPv6NetworkIDString(d))

	networkViewName := d.Get("network_view_name").(string)
	tenantID := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)
	extensibleAttributes := d.Get("extensible_attributes").(map[string]interface{})

	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	addEA := make(map[string]interface{})
	removeEA := make(map[string]interface{})

	var err error
	var IPv6Network *ibclient.IPv6Network
	for key, value := range extensibleAttributes {
		if value == "nil" {
			removeEA[key] = value
			IPv6Network, err = objMgr.UpdateIPv6NetworkEA(d.Id(), addEA, removeEA)
			if err != nil {
				return fmt.Errorf("Deletion of IPv6 network EA failed from network view(%s): %s", networkViewName, err)
			}
		} else {
			addEA[key] = value
			IPv6Network, err = objMgr.UpdateIPv6NetworkEA(d.Id(), addEA, removeEA)
			if err != nil {
				return fmt.Errorf("Addition of IPv6 network EA failed from network view(%s): %s", networkViewName, err)
			}

		}
	}
	d.SetId(IPv6Network.Ref)

	log.Printf("[DEBUG] %s: Updation on IPv6 network block complete", resourceIPv6NetworkIDString(d))
	return resourceIPv6NetworkRead(d, m)
}

func resourceIPv6NetworkDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Deletion of IPv6 network block", resourceIPv6NetworkIDString(d))

	networkViewName := d.Get("network_view_name").(string)
	tenantID := d.Get("tenant_id").(string)
	connector := m.(*ibclient.Connector)

	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)
	_, err := objMgr.DeleteIPv6Network(d.Id(), d.Get("network_view_name").(string))
	if err != nil {
		return fmt.Errorf("Deletion of IPv6 Network block failed from network view(%s): %s", networkViewName, err)
	}
	d.SetId("")

	log.Printf("[DEBUG] %s: Deletion of IPv6 network block complete", resourceIPv6NetworkIDString(d))
	return nil
}

type resourceIPv6NetworkIDStringInterface interface {
	Id() string
}

func resourceIPv6NetworkIDString(d resourceIPv6NetworkIDStringInterface) string {
	id := d.Id()
	if id == "" {
		id = "<new resource>"
	}
	return fmt.Sprintf("infobloxIPv6Network (ID = %s)", id)
}
