package infoblox

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"github.com/infobloxopen/infoblox-go-client/v2/utils"
)

func resourceEADefinition() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceEADefinitionCreate,
		ReadContext:   resourceEADefinitionRead,
		UpdateContext: resourceEADefinitionUpdate,
		DeleteContext: resourceEADefinitionDelete,

		Schema: map[string]*schema.Schema{
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comment for the Extensible Attribute Definition; maximum 256 characters.",
			},

			"flags": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "This field contains extensible attribute flags. Possible values: (A)udited,\n" +
					"(C)loud API, Cloud (G)master, (I)nheritable, (L)isted, (M)andatory value,\n" +
					"MGM (P)rivate, (R)ead Only, (S)ort enum values, Multiple (V)alues If there\n" +
					"are two or more flags in the field, you must list them according to the\n" +
					"order they are listed above. For example, 'CR' is a valid value for the\n" +
					"'flags' field because C = Cloud API is listed before R = Read only. However,\n" +
					"the value 'RC' is invalid because the order for the 'flags' field is broken.",
			},

			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the Extensible Attribute Definition.",
			},

			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type for the Extensible Attribute Definition.",
			},
		},
	}
}

func resourceEADefinitionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	conn := m.(ibclient.IBConnector)

	eaDef := &ibclient.EADefinition{
		Name: utils.StringPtr(d.Get("name").(string)),
		Type: d.Get("type").(string),
	}

	if d.HasChange("comment") {
		eaDef.Comment = utils.StringPtr(d.Get("comment").(string))
	}

	if d.HasChange("flags") {
		eaDef.Flags = utils.StringPtr(d.Get("flags").(string))
	}

	eaDefRef, err := conn.CreateObject(eaDef)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(eaDefRef)

	return resourceEADefinitionRead(ctx, d, m)
}

func resourceEADefinitionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	conn := m.(ibclient.IBConnector)

	eaDefRef := d.Id()

	eaDef := &ibclient.EADefinition{}

	eaDefResult := ibclient.EADefinition{}

	err := conn.GetObject(eaDef, eaDefRef, nil, &eaDefResult)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to read EA definition: %w", err))
	}

	if err = d.Set("name", eaDefResult.Name); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("type", eaDefResult.Type); err != nil {
		return diag.FromErr(err)
	}

	if eaDefResult.Comment != nil {
		if err = d.Set("comment", eaDefResult.Comment); err != nil {
			return diag.FromErr(err)
		}
	}

	if eaDefResult.Flags != nil {
		if err = d.Set("flags", eaDefResult.Flags); err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(eaDefResult.Ref)

	return nil
}

func resourceEADefinitionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	conn := m.(ibclient.IBConnector)

	eaDef := &ibclient.EADefinition{
		Name: utils.StringPtr(d.Get("name").(string)),
		Type: d.Get("type").(string),
	}

	if d.HasChange("comment") {
		eaDef.Comment = utils.StringPtr(d.Get("comment").(string))
	}

	if d.HasChange("flags") {
		eaDef.Flags = utils.StringPtr(d.Get("flags").(string))
	}

	eaDefRef, err := conn.UpdateObject(eaDef, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(eaDefRef)

	return resourceEADefinitionRead(ctx, d, m)
}

func resourceEADefinitionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	conn := m.(ibclient.IBConnector)

	if _, err := conn.DeleteObject(d.Id()); err != nil {
		return diag.FromErr(fmt.Errorf("deletion of EA definition failed: %w", err))
	}

	return nil
}
