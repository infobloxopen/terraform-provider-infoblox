package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"regexp"

	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

var (
	networkViewRegExp = regexp.MustCompile("^networkview/.+")
)

func resourceNetworkView() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkViewCreate,
		Read:   resourceNetworkViewRead,
		Update: resourceNetworkViewUpdate,
		Delete: resourceNetworkViewDelete,
		Importer: &schema.ResourceImporter{
			State: resourceNetworkViewImport,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Description: "Specifies the desired name of the network view as shown in the NIOS appliance." +
					" The name has the same requirements as the corresponding parameter in WAPI.",
			},
			"comment": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "A description of the network view.",
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
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}

	var tenantID string
	if tempVal, ok := extAttrs[eaNameForTenantId]; ok {
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
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}

	var tenantID string
	if tempVal, ok := extAttrs[eaNameForTenantId]; ok {
		tenantID = tempVal.(string)
	}

	Connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(Connector, "Terraform", tenantID)

	obj, err := objMgr.GetNetworkViewByRef(d.Id())
	if err != nil {
		return fmt.Errorf("Failed to get Network View : %s", err.Error())
	}

	if !networkViewRegExp.MatchString(d.Id()) {
		return fmt.Errorf("reference '%s' for 'networkview' object has an invalid format", d.Id())
	}

	omittedEAs := omitEAs(obj.Ea, extAttrs)

	if omittedEAs != nil && len(omittedEAs) > 0 {
		eaJSON, err := terraformSerializeEAs(omittedEAs)
		if err != nil {
			return err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return err
		}
	}

	d.SetId(obj.Ref)
	if err = d.Set("name", obj.Name); err != nil {
		return err
	}
	if err = d.Set("comment", obj.Comment); err != nil {
		return err
	}

	return nil
}

func resourceNetworkViewUpdate(d *schema.ResourceData, m interface{}) error {
	var updateSuccessful bool
	defer func() {
		// Reverting the state back, in case of a failure,
		// otherwise Terraform will keep the values, which leaded to the failure,
		// in the state file.
		if !updateSuccessful {
			prevName, _ := d.GetChange("name")
			prevComment, _ := d.GetChange("comment")
			prevEa, _ := d.GetChange("ext_attrs")

			_ = d.Set("name", prevName.(string))
			_ = d.Set("comment", prevComment.(string))
			_ = d.Set("ext_attrs", prevEa.(string))
		}
	}()

	networkView := d.Get("name").(string)
	comment := d.Get("comment").(string)

	oldExtAttrJSON, newExtAttrJSON := d.GetChange("ext_attrs")

	newExtAttrs, err := terraformDeserializeEAs(newExtAttrJSON.(string))
	if err != nil {
		return err
	}

	oldExtAttrs, err := terraformDeserializeEAs(oldExtAttrJSON.(string))
	if err != nil {
		return err
	}

	var tenantID string
	if tempVal, ok := newExtAttrs[eaNameForTenantId]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	nv, err := objMgr.GetNetworkViewByRef(d.Id())
	if err != nil {
		return fmt.Errorf("failed to read network for update operation: %w", err)
	}

	updExtAttrs, err := mergeEAs(nv.Ea, newExtAttrs, oldExtAttrs, connector)
	if err != nil {
		return err
	}

	nv, err = objMgr.UpdateNetworkView(d.Id(), networkView, comment, updExtAttrs)
	if err != nil {
		return fmt.Errorf("Failed to update Network View : %s", err.Error())
	}
	updateSuccessful = true
	d.SetId(nv.Ref)

	return nil
}

func resourceNetworkViewDelete(d *schema.ResourceData, m interface{}) error {
	if d.HasChange("name") {
		return fmt.Errorf("changing the value of 'networkView' field is not recommended")
	}
	networkView := d.Get("name").(string)

	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}
	var tenantID string
	if tempVal, ok := extAttrs[eaNameForTenantId]; ok {
		tenantID = tempVal.(string)
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)

	_, err = objMgr.DeleteNetworkView(d.Id())
	if err != nil {
		return fmt.Errorf("Deletion of Network view %s failed: %s", networkView, err.Error())
	}

	return nil
}

func resourceNetworkViewImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return nil, err
	}

	var tenantID string
	if tempVal, ok := extAttrs[eaNameForTenantId]; ok {
		tenantID = tempVal.(string)
	}

	Connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(Connector, "Terraform", tenantID)

	obj, err := objMgr.GetNetworkViewByRef(d.Id())
	if err != nil {
		return nil, fmt.Errorf("Failed to get Network View : %s", err.Error())
	}

	if !networkViewRegExp.MatchString(d.Id()) {
		return nil, fmt.Errorf("reference '%s' for 'networkview' object has an invalid format", d.Id())
	}

	if obj.Ea != nil && len(obj.Ea) > 0 {
		eaJSON, err := terraformSerializeEAs(obj.Ea)
		if err != nil {
			return nil, err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return nil, err
		}
	}

	d.SetId(obj.Ref)
	if err = d.Set("name", obj.Name); err != nil {
		return nil, err
	}
	if err = d.Set("comment", obj.Comment); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
