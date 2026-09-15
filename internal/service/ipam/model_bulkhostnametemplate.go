package ipam

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type BulkhostnametemplateModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var BulkhostnametemplateAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSBulkhostnametemplateAttrTypes},
}

type NIOSBulkhostnametemplateModel struct {
	TemplateFormat types.String `tfsdk:"template_format"`
	TemplateName   types.String `tfsdk:"template_name"`
}

var NIOSBulkhostnametemplateAttrTypes = map[string]attr.Type{
	"template_format": types.StringType,
	"template_name":   types.StringType,
}

const (
	BulkhostnametemplateReturnFields = "is_grid_default,pre_defined,template_format,template_name"
)

var BulkhostnametemplateResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          BulkhostnametemplateResourceNiosSchemaAttributes,
	},
}

var BulkhostnametemplateResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"template_format": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.RegexMatches(regexp.MustCompile(`.*(\$4|#4).*`), "Template format must contain at least one of $4 or #4 placeholders"),
		},
		MarkdownDescription: "The format of bulk host name template. It should follow certain rules (please use Administration Guide as reference).",
	},
	"template_name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of bulk host name template.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *BulkhostnametemplateModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Bulkhostnametemplate {
	if m == nil {
		return nil
	}

	obj := &coremodel.Bulkhostnametemplate{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSBulkhostnametemplateModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSBulkhostnametemplateModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSBulkhostnametemplateExt {
	return &coremodel.NIOSBulkhostnametemplateExt{
		TemplateFormat: flex.ExpandStringPointerNullAsEmpty(m.TemplateFormat),
		TemplateName:   flex.ExpandStringPointerNullAsEmpty(m.TemplateName),
	}
}

// Flatten populates the TF model from a core response.
func (m *BulkhostnametemplateModel) Flatten(ctx context.Context, resp *coremodel.Bulkhostnametemplate, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSBulkhostnametemplateModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSBulkhostnametemplateModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSBulkhostnametemplateAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSBulkhostnametemplateAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSBulkhostnametemplateModel) Flatten(ctx context.Context, from *coremodel.NIOSBulkhostnametemplateExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.TemplateFormat = flex.FlattenStringPointerEmptyAsNull(from.TemplateFormat)
	m.TemplateName = flex.FlattenStringPointerEmptyAsNull(from.TemplateName)
}
