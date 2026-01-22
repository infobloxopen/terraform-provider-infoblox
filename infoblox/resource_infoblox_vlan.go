package infoblox

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

var (
	vlanRegExp = regexp.MustCompile("^vlan/.+")
)

// VlanCreate is a custom struct for creating VLANs that supports both
// numeric vlan_id and "func:nextavailablevlanid" function syntax.
// This is needed because the standard Vlan struct has Id as *uint32.
type VlanCreate struct {
	ibclient.IBBase `json:"-"`

	Ref         string          `json:"_ref,omitempty"`
	Name        *string         `json:"name,omitempty"`
	Id          interface{}     `json:"id,omitempty"` // Can be uint32 or string for func:nextavailablevlanid
	Parent      *string         `json:"parent,omitempty"`
	Comment     *string         `json:"comment,omitempty"`
	Description *string         `json:"description,omitempty"`
	Department  *string         `json:"department,omitempty"`
	Contact     *string         `json:"contact,omitempty"`
	Ea          ibclient.EA     `json:"extattrs,omitempty"`

	returnFields []string
}

func (v VlanCreate) ObjectType() string {
	return "vlan"
}

func (v VlanCreate) ReturnFields() []string {
	return v.returnFields
}

func (v *VlanCreate) SetReturnFields(fields []string) {
	v.returnFields = fields
}

func resourceVlan() *schema.Resource {
	return &schema.Resource{
		Create: resourceVlanCreate,
		Read:   resourceVlanRead,
		Update: resourceVlanUpdate,
		Delete: resourceVlanDelete,
		Importer: &schema.ResourceImporter{
			State: resourceVlanImport,
		},
		CustomizeDiff: func(context context.Context, d *schema.ResourceDiff, meta interface{}) error {
			if internalID := d.Get("internal_id"); internalID == "" || internalID == nil {
				err := d.SetNewComputed("internal_id")
				if err != nil {
					return err
				}
			}
			// Mark vlan_id as computed if not explicitly set (next available will be used)
			if vlanId, ok := d.GetOk("vlan_id"); !ok || vlanId == 0 {
				if d.Id() == "" { // Only for new resources
					err := d.SetNewComputed("vlan_id")
					if err != nil {
						return err
					}
				}
			}
			return nil
		},

		Schema: map[string]*schema.Schema{
			"parent": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The VLAN View or VLAN Range to which this VLAN belongs (reference string).",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the VLAN.",
			},
			"vlan_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "VLAN ID value (typically 1-4094). If not specified, the next available VLAN ID will be allocated from the parent range.",
			},
			"comment": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "A descriptive comment for this VLAN.",
			},
			"description": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "Description for the VLAN object, may be potentially used for longer VLAN names.",
			},
			"department": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "Department where VLAN is used.",
			},
			"contact": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "Contact information for person/team managing or using VLAN.",
			},
			"ext_attrs": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "The Extensible attributes of the VLAN to be added/updated, as a map in JSON format",
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

func resourceVlanCreate(d *schema.ResourceData, m interface{}) error {
	if intId := d.Get("internal_id"); intId.(string) != "" {
		return fmt.Errorf("the value of 'internal_id' field must not be set manually")
	}

	parent := d.Get("parent").(string)
	name := d.Get("name").(string)
	comment := d.Get("comment").(string)
	description := d.Get("description").(string)
	department := d.Get("department").(string)
	contact := d.Get("contact").(string)
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}

	connector := m.(ibclient.IBConnector)

	// Generate UUID for internal_id and add to the EA
	internalId := generateInternalId()
	extAttrs[eaNameForInternalId] = internalId.String()

	// Determine VLAN ID - use provided value or next available
	var vlanIdValue interface{}
	if v, ok := d.GetOk("vlan_id"); ok {
		vlanIdValue = uint32(v.(int))
	} else {
		// Use next available VLAN ID function
		vlanIdValue = fmt.Sprintf("func:nextavailablevlanid:%s", parent)
	}

	// Create VLAN object using custom struct that supports both numeric and function ID
	vlan := &VlanCreate{
		Parent:      &parent,
		Name:        &name,
		Id:          vlanIdValue,
		Comment:     &comment,
		Description: &description,
		Department:  &department,
		Contact:     &contact,
		Ea:          extAttrs,
	}
	vlan.SetReturnFields([]string{"id", "name", "comment", "description", "department", "contact", "extattrs"})

	ref, err := connector.CreateObject(vlan)
	if err != nil {
		return fmt.Errorf("failed to create VLAN: %s", err)
	}

	d.SetId(ref)
	if err = d.Set("internal_id", internalId.String()); err != nil {
		return err
	}
	if err = d.Set("ref", ref); err != nil {
		return err
	}

	// Read back the VLAN to get the actual vlan_id (especially important for next available)
	return resourceVlanRead(d, m)
}

func resourceVlanRead(d *schema.ResourceData, m interface{}) error {
	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return err
	}

	connector := m.(ibclient.IBConnector)

	// Get VLAN object by ref
	// Note: parent field is not included in return fields because the API returns it as an object
	// but the Go client expects a string. Parent is already stored in state from creation.
	vlan := &ibclient.Vlan{}
	vlan.SetReturnFields([]string{"id", "name", "comment", "description", "department", "contact", "extattrs"})
	err = connector.GetObject(vlan, d.Id(), ibclient.NewQueryParams(false, nil), vlan)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); ok {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("failed to get VLAN: %s", err.Error())
	}

	if !vlanRegExp.MatchString(vlan.Ref) {
		return fmt.Errorf("reference '%s' for 'vlan' object has an invalid format", vlan.Ref)
	}

	delete(vlan.Ea, eaNameForInternalId)
	omittedEAs := omitEAs(vlan.Ea, extAttrs)

	if omittedEAs != nil && len(omittedEAs) > 0 {
		eaJSON, err := terraformSerializeEAs(omittedEAs)
		if err != nil {
			return err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return err
		}
	}

	d.SetId(vlan.Ref)
	if err = d.Set("ref", vlan.Ref); err != nil {
		return err
	}
	if vlan.Name != nil {
		if err = d.Set("name", *vlan.Name); err != nil {
			return err
		}
	}
	if vlan.Id != nil {
		if err = d.Set("vlan_id", int(*vlan.Id)); err != nil {
			return err
		}
	}
	if vlan.Comment != nil {
		if err = d.Set("comment", *vlan.Comment); err != nil {
			return err
		}
	}
	if vlan.Description != nil {
		if err = d.Set("description", *vlan.Description); err != nil {
			return err
		}
	}
	if vlan.Department != nil {
		if err = d.Set("department", *vlan.Department); err != nil {
			return err
		}
	}
	if vlan.Contact != nil {
		if err = d.Set("contact", *vlan.Contact); err != nil {
			return err
		}
	}
	// Note: parent is not read back from API due to type mismatch in Go client.
	// It's already stored in state and is ForceNew, so it can't change.

	return nil
}

func resourceVlanUpdate(d *schema.ResourceData, m interface{}) error {
	var updateSuccessful bool
	defer func() {
		// Reverting the state back, in case of a failure,
		// otherwise Terraform will keep the values, which leaded to the failure,
		// in the state file.
		if !updateSuccessful {
			prevName, _ := d.GetChange("name")
			prevVlanId, _ := d.GetChange("vlan_id")
			prevComment, _ := d.GetChange("comment")
			prevDescription, _ := d.GetChange("description")
			prevDepartment, _ := d.GetChange("department")
			prevContact, _ := d.GetChange("contact")
			prevEa, _ := d.GetChange("ext_attrs")

			_ = d.Set("name", prevName.(string))
			_ = d.Set("vlan_id", prevVlanId.(int))
			_ = d.Set("comment", prevComment.(string))
			_ = d.Set("description", prevDescription.(string))
			_ = d.Set("department", prevDepartment.(string))
			_ = d.Set("contact", prevContact.(string))
			_ = d.Set("ext_attrs", prevEa.(string))
		}
	}()

	if d.HasChange("internal_id") {
		return fmt.Errorf("changing the value of 'internal_id' field is not allowed")
	}

	name := d.Get("name").(string)
	vlanId := uint32(d.Get("vlan_id").(int))
	comment := d.Get("comment").(string)
	description := d.Get("description").(string)
	department := d.Get("department").(string)
	contact := d.Get("contact").(string)

	oldExtAttrJSON, newExtAttrJSON := d.GetChange("ext_attrs")

	newExtAttrs, err := terraformDeserializeEAs(newExtAttrJSON.(string))
	if err != nil {
		return err
	}

	oldExtAttrs, err := terraformDeserializeEAs(oldExtAttrJSON.(string))
	if err != nil {
		return err
	}

	connector := m.(ibclient.IBConnector)

	// Get current VLAN object
	// Note: parent field is not included due to type mismatch in Go client
	vlan := &ibclient.Vlan{}
	vlan.SetReturnFields([]string{"id", "name", "comment", "description", "department", "contact", "extattrs"})
	err = connector.GetObject(vlan, d.Id(), ibclient.NewQueryParams(false, nil), vlan)
	if err != nil {
		return fmt.Errorf("failed to read VLAN for update operation: %w", err)
	}

	internalId := d.Get("internal_id").(string)

	if internalId == "" {
		internalId = generateInternalId().String()
	}

	newInternalId := newInternalResourceIdFromString(internalId)
	newExtAttrs[eaNameForInternalId] = newInternalId.String()

	updExtAttrs, err := mergeEAs(vlan.Ea, newExtAttrs, oldExtAttrs, connector)
	if err != nil {
		return err
	}

	// Update VLAN object
	vlan.Name = &name
	vlan.Id = &vlanId
	vlan.Comment = &comment
	vlan.Description = &description
	vlan.Department = &department
	vlan.Contact = &contact
	vlan.Ea = updExtAttrs

	updatedRef, err := connector.UpdateObject(vlan, d.Id())
	if err != nil {
		return fmt.Errorf("failed to update VLAN: %s", err.Error())
	}
	updateSuccessful = true

	if err = d.Set("internal_id", newInternalId.String()); err != nil {
		return err
	}
	if err = d.Set("ref", updatedRef); err != nil {
		return err
	}
	d.SetId(updatedRef)

	return nil
}

func resourceVlanDelete(d *schema.ResourceData, m interface{}) error {
	connector := m.(ibclient.IBConnector)

	// Delete VLAN directly by reference
	_, err := connector.DeleteObject(d.Id())
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); ok {
			// Already deleted, clear state
			d.SetId("")
			return nil
		}
		return fmt.Errorf("deletion of VLAN failed: %s", err.Error())
	}

	d.SetId("")
	return nil
}

func resourceVlanImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	connector := m.(ibclient.IBConnector)

	// Get VLAN object by ref
	// Note: parent field is not included due to type mismatch in Go client.
	// User must specify parent in their Terraform config after import.
	vlan := &ibclient.Vlan{}
	vlan.SetReturnFields([]string{"id", "name", "comment", "description", "department", "contact", "extattrs"})
	err := connector.GetObject(vlan, d.Id(), ibclient.NewQueryParams(false, nil), vlan)
	if err != nil {
		return nil, fmt.Errorf("failed to get VLAN: %s", err.Error())
	}

	if !vlanRegExp.MatchString(d.Id()) {
		return nil, fmt.Errorf("reference '%s' for 'vlan' object has an invalid format", d.Id())
	}

	if vlan.Ea != nil && len(vlan.Ea) > 0 {
		eaJSON, err := terraformSerializeEAs(vlan.Ea)
		if err != nil {
			return nil, err
		}
		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return nil, err
		}
	}

	d.SetId(vlan.Ref)
	if vlan.Name != nil {
		if err = d.Set("name", *vlan.Name); err != nil {
			return nil, err
		}
	}
	if vlan.Id != nil {
		if err = d.Set("vlan_id", int(*vlan.Id)); err != nil {
			return nil, err
		}
	}
	if vlan.Comment != nil {
		if err = d.Set("comment", *vlan.Comment); err != nil {
			return nil, err
		}
	}
	if vlan.Description != nil {
		if err = d.Set("description", *vlan.Description); err != nil {
			return nil, err
		}
	}
	if vlan.Department != nil {
		if err = d.Set("department", *vlan.Department); err != nil {
			return nil, err
		}
	}
	if vlan.Contact != nil {
		if err = d.Set("contact", *vlan.Contact); err != nil {
			return nil, err
		}
	}
	// Note: parent is not read from API due to type mismatch. User must specify it in config.

	// Update the resource with the EA Terraform Internal ID
	err = resourceVlanUpdate(d, m)
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
