package infoblox

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func resourceNetworkView() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkViewCreate,
		Read:   resourceNetworkViewRead,
		Update: resourceNetworkViewUpdate,
		Delete: resourceNetworkViewDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Desired name of the view shown in NIOS appliance.",
			},
			"comment": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "A description of the IP allocation.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "The Extensible attributes of the network container to be added/updated, as a map in JSON format",
			},
		},
	}
}

func resourceNetworkViewCreate(d *schema.ResourceData, m interface{}) error {

	networkView := d.Get("name").(string)
	comment := d.Get("comment").(string)
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
		}
	}
	var tenantID string
	if tempVal, ok := extAttrs["Tenant ID"]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	nv, err := objMgr.CreateNetworkView(networkView, comment, extAttrs)
	if err != nil {
		return fmt.Errorf("Failed to create Network View : %s", err)
	}
	d.SetId(nv.Ref)
	return nil
}

func resourceNetworkViewRead(d *schema.ResourceData, m interface{}) error {

	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
		}
	}
	var tenantID string
	if tempVal, ok := extAttrs["Tenant ID"]; ok {
		tenantID = tempVal.(string)
	}

	Connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(Connector, "Terraform", tenantID)

	obj, err := objMgr.GetNetworkViewByRef(d.Id())
	if err != nil {
		return fmt.Errorf("Failed to get Network View : %s", err.Error())
	}

	d.SetId(obj.Ref)
	return nil
}

func resourceNetworkViewUpdate(d *schema.ResourceData, m interface{}) error {
	networkView := d.Get("name").(string)
	comment := d.Get("comment").(string)
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
		}
	}
	var tenantID string
	if tempVal, ok := extAttrs["Tenant ID"]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)
	nv, err := objMgr.UpdateNetworkView(d.Id(), networkView, comment, extAttrs)
	if err != nil {
		return fmt.Errorf("Failed to update Network View : %s", err.Error())
	}
	d.SetId(nv.Ref)
	return nil
}

func resourceNetworkViewDelete(d *schema.ResourceData, m interface{}) error {
	if d.HasChange("name") {
		return fmt.Errorf("changing the value of 'networkView' field is not recommended")
	}
	networkView := d.Get("name").(string)
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
		}
	}
	var tenantID string
	if tempVal, ok := extAttrs["Tenant ID"]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	_, err := objMgr.DeleteNetworkView(d.Id())
	if err != nil {
		return fmt.Errorf("Deletion of Network view %s failed: %s", networkView, err.Error())
	}

	return nil
}
