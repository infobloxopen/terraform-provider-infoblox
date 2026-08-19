package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// OptionItemModel is the Terraform model for OptionItem
type OptionItemModel struct {
	Group       types.String `tfsdk:"group"`
	OptionCode  types.String `tfsdk:"option_code"`
	OptionValue types.String `tfsdk:"option_value"`
	Type        types.String `tfsdk:"type"`
}

// OptionItemAttrTypes contains the attribute types for OptionItemModel
var OptionItemAttrTypes = map[string]attr.Type{
	"group":        types.StringType,
	"option_code":  types.StringType,
	"option_value": types.StringType,
	"type":         types.StringType,
}

// OptionItemResourceSchemaAttributes contains the schema attributes for OptionItemModel
var OptionItemResourceSchemaAttributes = map[string]schema.Attribute{
	"group": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(
				path.MatchRelative().AtParent().AtName("option_code"),
				path.MatchRelative().AtParent().AtName("group"),
			),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"option_code": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(
				path.MatchRelative().AtParent().AtName("option_code"),
				path.MatchRelative().AtParent().AtName("group"),
			),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"option_value": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("option_code")),
		},
		MarkdownDescription: "The option value.",
	},
	"type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("group", "option"),
		},
		Optional:            true,
		MarkdownDescription: "The type of item.  Valid values are: * _group_ * _option_",
	},
}

// ExpandOptionItem converts a Terraform Object to SDK type
func ExpandOptionItem(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiipam.OptionItem {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m OptionItemModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *OptionItemModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiipam.OptionItem {
	if m == nil {
		return nil
	}
	to := &uddiipam.OptionItem{
		Group:       flex.ExpandStringPointer(m.Group),
		OptionCode:  flex.ExpandStringPointer(m.OptionCode),
		OptionValue: flex.ExpandStringPointer(m.OptionValue),
		Type:        flex.ExpandStringPointer(m.Type),
	}
	return to
}

// FlattenOptionItem converts an SDK type to Terraform Object
func FlattenOptionItem(ctx context.Context, from *uddiipam.OptionItem, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(OptionItemAttrTypes)
	}
	m := &OptionItemModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, OptionItemAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *OptionItemModel) Flatten(ctx context.Context, from *uddiipam.OptionItem, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Group = flex.FlattenStringPointer(from.Group)
	m.OptionCode = flex.FlattenStringPointer(from.OptionCode)
	m.OptionValue = flex.FlattenStringPointer(from.OptionValue)
	m.Type = flex.FlattenStringPointer(from.Type)
}
