package infoblox

import (
	"context"
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
		},
	}
}

func resourceDNSViewCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	conn := m.(ibclient.IBConnector)

	extAttrJSON := d.Get("ext_attrs").(string)
	extAttrs, err := terraformDeserializeEAs(extAttrJSON)
	if err != nil {
		return diag.FromErr(err)
	}

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

	return resourceDNSViewRead(ctx, d, m)
}

func resourceDNSViewRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	conn := m.(ibclient.IBConnector)

	viewRef := d.Id()

	v := &ibclient.View{}
	v.SetReturnFields([]string{"name", "comment", "network_view", "extattrs"})

	vResult := ibclient.View{}

	if !dnsViewRegExp.MatchString(d.Id()) {
		return diag.FromErr(fmt.Errorf("reference '%s' for 'view' object has an invalid format", d.Id()))
	}

	err := conn.GetObject(v, viewRef, nil, &vResult)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to read DNS View: %w", err))
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

	d.SetId(vResult.Ref)

	return nil
}

func resourceDNSViewUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	return resourceDNSViewRead(ctx, d, m)
}

func resourceDNSViewDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	conn := m.(ibclient.IBConnector)

	if _, err := conn.DeleteObject(d.Id()); err != nil {
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

	return []*schema.ResourceData{d}, nil
}
