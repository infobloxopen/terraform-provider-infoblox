package infoblox

import (
	"context"
	"encoding/json"
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
		CustomizeDiff: func(context context.Context, d *schema.ResourceDiff, meta interface{}) error {
			if internalID := d.Get("internal_id"); internalID == "" || internalID == nil {
				err := d.SetNewComputed("internal_id")
				if err != nil {
					return err
				}
			}
			return nil
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
			"internal_id": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "Internal ID of an object at NIOS side," +
					" used by Infoblox Terraform plugin to search for a NIOS's object" +
					" which corresponds to the Terraform resource.",
			},
			"ref": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "NIOS object's reference, not to be set by a user.",
			},
		},
	}
}

func resourceNetworkViewCreate(d *schema.ResourceData, m interface{}) error {
	if intId := d.Get("internal_id"); intId.(string) != "" {
		return fmt.Errorf("the value of 'internal_id' field must not be set manually")
	}
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

	// Generate UUID for internal_id and add to the EA
	internalId := generateInternalId()
	extAttrs[eaNameForInternalId] = internalId.String()

	nv, err := objMgr.CreateNetworkView(networkView, comment, extAttrs)
	if err != nil {
		return fmt.Errorf("Failed to create Network View : %s", err)
	}
	d.SetId(nv.Ref)
	if err = d.Set("internal_id", internalId.String()); err != nil {
		return err
	}
	if err = d.Set("ref", nv.Ref); err != nil {
		return err
	}
	return nil
}

func resourceNetworkViewRead(d *schema.ResourceData, m interface{}) error {

	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}

	obj, err := searchObjectByRefOrInternalId("NetworkView", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		} else {
			d.SetId("")
			return nil
		}
	}

	var nv *ibclient.NetworkView
	recJson, _ := json.Marshal(obj)
	err = json.Unmarshal(recJson, &nv)

	if err != nil {
		return fmt.Errorf("Failed to get Network View : %s", err.Error())
	}

	if !networkViewRegExp.MatchString(nv.Ref) {
		return fmt.Errorf("reference '%s' for 'networkview' object has an invalid format", nv.Ref)
	}
	delete(nv.Ea, eaNameForTenantId)
	omittedEAs := omitEAs(nv.Ea, extAttrs)

	if omittedEAs != nil && len(omittedEAs) > 0 {
		eaJSON, err := terraformSerializeEAs(omittedEAs)
		if err != nil {
			return err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return err
		}
	}

	d.SetId(nv.Ref)
	if err = d.Set("ref", nv.Ref); err != nil {
		return err
	}
	if err = d.Set("name", nv.Name); err != nil {
		return err
	}
	if err = d.Set("comment", nv.Comment); err != nil {
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

	if d.HasChange("internal_id") {
		return fmt.Errorf("changing the value of 'internal_id' field is not allowed")
	}

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

	internalId := d.Get("internal_id").(string)

	if internalId == "" {
		internalId = generateInternalId().String()
	}

	newInternalId := newInternalResourceIdFromString(internalId)
	newExtAttrs[eaNameForInternalId] = newInternalId.String()

	updExtAttrs, err := mergeEAs(nv.Ea, newExtAttrs, oldExtAttrs, connector)
	if err != nil {
		return err
	}

	nv, err = objMgr.UpdateNetworkView(d.Id(), networkView, comment, updExtAttrs)
	if err != nil {
		return fmt.Errorf("Failed to update Network View : %s", err.Error())
	}
	updateSuccessful = true

	if err = d.Set("internal_id", newInternalId.String()); err != nil {
		return err
	}
	if err = d.Set("ref", nv.Ref); err != nil {
		return err
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
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}
	var tenantID string
	if tempVal, ok := extAttrs[eaNameForTenantId]; ok {
		tenantID = tempVal.(string)
	}

	obj, err := searchObjectByRefOrInternalId("NetworkView", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err))
		} else {
			d.SetId("")
			return nil
		}
	}

	var nv *ibclient.NetworkView
	recJson, _ := json.Marshal(obj)
	err = json.Unmarshal(recJson, &nv)
	if err != nil {
		return err
	}

	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", tenantID)
	_, err = objMgr.DeleteNetworkView(nv.Ref)
	if err != nil {
		return fmt.Errorf("Deletion of Network view %s failed: %s", networkView, err.Error())
	}
	d.SetId("")
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

	// Update the resource with the EA Terraform Internal ID
	err = resourceNetworkViewUpdate(d, m)
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
