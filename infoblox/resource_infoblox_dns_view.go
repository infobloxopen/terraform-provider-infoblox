package infoblox

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"github.com/infobloxopen/infoblox-go-client/v2/utils"
	"regexp"
)

var (
	dnsViewRegExp = regexp.MustCompile("^view/.+")
)

func resourceDNSView() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDNSViewCreate,
		ReadContext:   resourceDNSViewRead,
		UpdateContext: resourceDNSViewUpdate,
		DeleteContext: resourceDNSViewDelete,
		Importer: &schema.ResourceImporter{
			State: resourceDNSViewImport,
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
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comment for the DNS View object; maximum 256 characters.",
			},

			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the DNS View to be specified",
			},

			"network_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
				Description: "The name of the Network View in which DNS View exists.",
			},

			"ext_attrs": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "The Extensible attributes of the DNS view to be added/updated, as a map in JSON format",
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

func resourceDNSViewCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if intId := d.Get("internal_id"); intId.(string) != "" {
		return diag.FromErr(fmt.Errorf("the value of 'internal_id' field must not be set manually"))
	}

	if ref := d.Get("ref"); ref.(string) != "" {
		return diag.FromErr(fmt.Errorf("the value of 'ref' field must not be set manually"))
	}

	conn := m.(ibclient.IBConnector)

	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return diag.FromErr(err)
	}

	// Generate internal ID and add it to the extensible attributes
	internalId := generateInternalId()
	extAttrs[eaNameForInternalId] = internalId.String()

	v := &ibclient.View{
		Name: utils.StringPtr(d.Get("name").(string)),
		Ea:   extAttrs,
	}

	if d.HasChange("comment") {
		v.Comment = utils.StringPtr(d.Get("comment").(string))
	}

	if d.HasChange("network_view") {
		v.NetworkView = utils.StringPtr(d.Get("network_view").(string))
	}

	viewRef, err := conn.CreateObject(v)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(viewRef)

	if err = d.Set("ref", viewRef); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("internal_id", internalId.String()); err != nil {
		return diag.FromErr(err)
	}

	return resourceDNSViewRead(ctx, d, m)
}

func resourceDNSViewRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	v := &ibclient.View{}
	v.SetReturnFields([]string{"name", "comment", "network_view", "extattrs"})

	if !dnsViewRegExp.MatchString(d.Id()) {
		return diag.FromErr(fmt.Errorf("reference '%s' for 'view' object has an invalid format", d.Id()))
	}

	rec, err := searchObjectByRefOrInternalId("DNSView", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return diag.FromErr(ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err)))
		} else {
			d.SetId("")
			return nil
		}
	}

	// Assertion of object type and error handling
	var vResult *ibclient.View
	recJson, _ := json.Marshal(rec)
	err = json.Unmarshal(recJson, &vResult)

	if err != nil && vResult.Ref != "" {
		return diag.FromErr(fmt.Errorf("getting DNS View with ID: %s failed: %w", d.Id(), err))
	}

	if err = d.Set("name", vResult.Name); err != nil {
		return diag.FromErr(err)
	}

	if vResult.Comment != nil {
		if err = d.Set("comment", vResult.Comment); err != nil {
			return diag.FromErr(err)
		}
	}

	if vResult.NetworkView != nil {
		if err = d.Set("network_view", vResult.NetworkView); err != nil {
			return diag.FromErr(err)
		}
	}

	extAttrsJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrsJSON)
	if err != nil {
		return diag.FromErr(err)
	}

	delete(vResult.Ea, eaNameForInternalId)
	omittedEAs := omitEAs(vResult.Ea, extAttrs)

	if omittedEAs != nil && len(omittedEAs) > 0 {
		eaJSON, err := terraformSerializeEAs(omittedEAs)
		if err != nil {
			return diag.FromErr(err)
		}

		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return diag.FromErr(err)
		}
	}
	if err = d.Set("ref", vResult.Ref); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(vResult.Ref)

	return nil
}

func resourceDNSViewUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if d.HasChange("internal_id") {
		return diag.FromErr(fmt.Errorf("changing the value of 'internal_id' field is not allowed"))
	}

	if d.HasChange("ref") {
		return diag.FromErr(fmt.Errorf("changing the value of 'ref' field is not allowed"))
	}

	conn := m.(ibclient.IBConnector)

	oldExtAttrsJSON, newExtAttrsJSON := d.GetChange("ext_attrs")

	newExtAttrs, err := terraformDeserializeEAs(newExtAttrsJSON.(string))
	if err != nil {
		return diag.FromErr(err)
	}

	oldExtAttrs, err := terraformDeserializeEAs(oldExtAttrsJSON.(string))
	if err != nil {
		return diag.FromErr(err)
	}

	v := &ibclient.View{}
	v.SetReturnFields([]string{"extattrs"})

	vResult := ibclient.View{}

	err = conn.GetObject(v, d.Id(), nil, &vResult)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to read DNS View for update operation: %w", err))
	}

	internalId := d.Get("internal_id").(string)

	if internalId == "" {
		internalId = generateInternalId().String()
	}

	newInternalId := newInternalResourceIdFromString(internalId)
	newExtAttrs[eaNameForInternalId] = newInternalId.String()

	mergedExtAttrs, err := mergeEAs(vResult.Ea, newExtAttrs, oldExtAttrs, conn)
	if err != nil {
		return diag.FromErr(err)
	}
	vUpd := &ibclient.View{
		Name: utils.StringPtr(d.Get("name").(string)),
	}

	if d.HasChange("comment") {
		vUpd.Comment = utils.StringPtr(d.Get("comment").(string))
	}

	if d.HasChange("network_view") {
		vUpd.NetworkView = utils.StringPtr(d.Get("network_view").(string))
	}

	vUpd.Ea = mergedExtAttrs

	viewRef, err := conn.UpdateObject(vUpd, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(viewRef)

	if err = d.Set("ref", viewRef); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("internal_id", newInternalId.String()); err != nil {
		return diag.FromErr(err)
	}

	return resourceDNSViewRead(ctx, d, m)
}

func resourceDNSViewDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	conn := m.(ibclient.IBConnector)

	rec, err := searchObjectByRefOrInternalId("DNSView", d, m)
	if err != nil {
		if _, ok := err.(*ibclient.NotFoundError); !ok {
			return diag.FromErr(ibclient.NewNotFoundError(fmt.Sprintf(
				"cannot find appropriate object on NIOS side for resource with ID '%s': %s;", d.Id(), err)))
		} else {
			d.SetId("")
			return nil
		}
	}

	// Assertion of object type and error handling
	var vResult *ibclient.View
	recJson, _ := json.Marshal(rec)
	err = json.Unmarshal(recJson, &vResult)

	if err != nil && vResult.Ref != "" {
		return diag.FromErr(fmt.Errorf("getting DNS View with ID: %s failed: %w", d.Id(), err))
	}

	if _, err := conn.DeleteObject(vResult.Ref); err != nil {
		return diag.FromErr(fmt.Errorf("deletion of DNS View failed: %w", err))
	}

	return nil
}

func resourceDNSViewImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	conn := m.(ibclient.IBConnector)

	viewRef := d.Id()

	v := &ibclient.View{}
	v.SetReturnFields([]string{"name", "comment", "network_view", "extattrs"})

	vResult := ibclient.View{}

	if !dnsViewRegExp.MatchString(d.Id()) {
		return nil, fmt.Errorf("reference '%s' for 'view' object has an invalid format", d.Id())
	}

	err := conn.GetObject(v, viewRef, nil, &vResult)
	if err != nil {
		return nil, fmt.Errorf("failed to read DNS View: %w", err)
	}

	if err = d.Set("name", vResult.Name); err != nil {
		return nil, err
	}

	if vResult.Comment != nil {
		if err = d.Set("comment", vResult.Comment); err != nil {
			return nil, err
		}
	}

	if vResult.NetworkView != nil {
		if err = d.Set("network_view", vResult.NetworkView); err != nil {
			return nil, err
		}
	}

	extAttrsJSON := d.Get("ext_attrs").(string)
	_, err = terraformDeserializeEAs(extAttrsJSON)
	if err != nil {
		return nil, err
	}

	if vResult.Ea != nil && len(vResult.Ea) > 0 {
		eaJSON, err := terraformSerializeEAs(vResult.Ea)
		if err != nil {
			return nil, err
		}

		if err = d.Set("ext_attrs", eaJSON); err != nil {
			return nil, err
		}
	}

	d.SetId(vResult.Ref)

	var dErr diag.Diagnostics
	var ctx context.Context
	dErr = resourceDNSViewUpdate(ctx, d, m)
	if dErr != nil {
		return nil, fmt.Errorf("failed to import DNS View: %v", dErr)
	}

	return []*schema.ResourceData{d}, nil
}
