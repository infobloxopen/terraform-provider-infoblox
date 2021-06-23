package infoblox

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client"
)

func resourceNetworkView() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkViewCreate,
		Read:   resourceNetworkViewRead,
		Update: resourceNetworkViewUpdate,
		Delete: resourceNetworkViewDelete,

		Schema: map[string]*schema.Schema{
			"network_view_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Desired name of the view shown in NIOS appliance.",
			},
			"tenant_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of a tenant to be used when creating objects.",
			},

			// TODO: uncomment this once EA and comment support is added for NetView in go-client
			//"comment": {
			//	Type:        schema.TypeString,
			//	Default:     "",
			//	Optional:    true,
			//	Description: "A description of the IP allocation.",
			//},
			//"extensible_attributes": {
			//	Type:        schema.TypeString,
			//	Default:     "",
			//	Optional:    true,
			//	Description: "The Extensible attributes of the network container to be added/updated, as a map in JSON format",
			//},
		},
	}
}

func resourceNetworkViewCreate(d *schema.ResourceData, m interface{}) error {
	var tenantID string

	networkViewName := d.Get("network_view_name").(string)

	// TODO: uncomment this once EA and comment support is added for NetView in go-client
	//extAttrJSON := d.Get("extensible_attributes").(string)
	//extAttrs := make(map[string]interface{})
	//if extAttrJSON != "" {
	//	if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
	//		return fmt.Errorf("cannot process 'extensible_attributes' field: %s", err.Error())
	//	}
	//}
	//for attrName, attrValue := range extAttrs {
	//	if attrName == "Tenant ID" {
	//		tenantID = attrValue.(string)
	//		break
	//	}
	//}

	tenantIdTfProp, found := d.GetOk("tenant_id")
	if found {
		tenantID = tenantIdTfProp.(string)
	}

	Connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(Connector, "Terraform", tenantID)

	// TODO: Add comment and EAs while creation
	// comment := d.Get("comment").(string)
	//networkView, err := objMgr.CreateNetworkView(networkViewName, comment, extAttrs)

	obj, err := objMgr.CreateNetworkView(networkViewName)
	if err != nil {
		return fmt.Errorf("Failed to create Network View : %s", err)
	}
	d.SetId(obj.Ref)
	return nil
}

func resourceNetworkViewRead(d *schema.ResourceData, m interface{}) error {
	var tenantID string

	// TODO: uncomment this once EA and comment support is added for NetView in go-client
	//extAttrJSON := d.Get("extensible_attributes").(string)
	//extAttrs := make(map[string]interface{})
	//if extAttrJSON != "" {
	//	if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
	//		return fmt.Errorf("cannot process 'extensible_attributes' field: %s", err.Error())
	//	}
	//}
	//for attrName, attrValue := range extAttrs {
	//	if attrName == "Tenant ID" {
	//		tenantID = attrValue.(string)
	//		break
	//	}
	//}

	tenantIdTfProp, found := d.GetOk("tenant_id")
	if found {
		tenantID = tenantIdTfProp.(string)
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
	// TODO: Implement update at go client
	/*
		networkViewName := d.Get("network_view_name").(string)
		comment := d.Get("comment").(string)
		extAttrJSON := d.Get("extensible_attributes").(string)
		extAttrs := make(map[string]interface{})
		if extAttrJSON != "" {
			if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
				return fmt.Errorf("cannot process 'extensible_attributes' field: %s", err.Error())
			}
		}
		var tenantID string
		for attrName, attrValue := range extAttrs {
			if attrName == "Tenant ID" {
				tenantID = attrValue.(string)
				break
			}
		}
		Connector := m.(*ibclient.Connector)
		objMgr := ibclient.NewObjectManager(Connector, "Terraform", tenantID)
		obj, err := objMgr.UpdateNetworkView(networkViewName, comment, extAttrs)
		if err != nil {
			return fmt.Errorf("Failed to update Network View : %s", err.Error())
		}
		d.SetId(obj.Ref)
		return nil
	*/
	return fmt.Errorf("network view updation is not supported")
}

func resourceNetworkViewDelete(d *schema.ResourceData, m interface{}) error {
	var tenantID string

	networkViewName := d.Get("network_view_name").(string)

	// TODO: uncomment this once EA and comment support is added for NetView in go-client
	//extAttrJSON := d.Get("extensible_attributes").(string)
	//extAttrs := make(map[string]interface{})
	//if extAttrJSON != "" {
	//	if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
	//		return fmt.Errorf("cannot process 'extensible_attributes' field: %s", err.Error())
	//	}
	//}
	//for attrName, attrValue := range extAttrs {
	//	if attrName == "Tenant ID" {
	//		tenantID = attrValue.(string)
	//		break
	//	}
	//}

	tenantIdTfProp, found := d.GetOk("tenant_id")
	if found {
		tenantID = tenantIdTfProp.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	_, err := objMgr.DeleteNetworkView(d.Id())
	if err != nil {
		return fmt.Errorf("Deletion of Network view %s failed: %s", networkViewName, err.Error())
	}

	return nil
}
