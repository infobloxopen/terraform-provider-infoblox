package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// InheritedDHCPOptionListModel is the Terraform model for InheritedDHCPOptionList
type InheritedDHCPOptionListModel struct {
	Action types.String `tfsdk:"action"`
	Value  types.List   `tfsdk:"value"`
}

// InheritedDHCPOptionListAttrTypes contains the attribute types for InheritedDHCPOptionListModel
var InheritedDHCPOptionListAttrTypes = map[string]attr.Type{
	"action": types.StringType,
	"value":  types.ListType{ElemType: types.ObjectType{AttrTypes: InheritedDHCPOptionAttrTypes}},
}

// InheritedDHCPOptionListResourceSchemaAttributes contains the schema attributes for InheritedDHCPOptionListModel
var InheritedDHCPOptionListResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance setting.  Valid values are: * _inherit_: Use the inherited value. * _block_: Don't use the inherited value.  Defaults to _inherit_.",
	},
	"value": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: InheritedDHCPOptionResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The inherited DHCP option values.",
	},
}

// ExpandInheritedDHCPOptionList converts a Terraform Object to SDK type
func ExpandInheritedDHCPOptionList(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiipam.InheritedDHCPOptionList {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritedDHCPOptionListModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *InheritedDHCPOptionListModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiipam.InheritedDHCPOptionList {
	if m == nil {
		return nil
	}
	to := &uddiipam.InheritedDHCPOptionList{
		Action: flex.ExpandStringPointer(m.Action),
		Value:  flex.ExpandFrameworkListNestedBlock(ctx, m.Value, diags, ExpandInheritedDHCPOption),
	}
	return to
}

// FlattenInheritedDHCPOptionList converts an SDK type to Terraform Object
func FlattenInheritedDHCPOptionList(ctx context.Context, from *uddiipam.InheritedDHCPOptionList, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritedDHCPOptionListAttrTypes)
	}
	m := &InheritedDHCPOptionListModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritedDHCPOptionListAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *InheritedDHCPOptionListModel) Flatten(ctx context.Context, from *uddiipam.InheritedDHCPOptionList, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.Value = flex.FlattenFrameworkListNestedBlock(ctx, from.Value, InheritedDHCPOptionAttrTypes, diags, FlattenInheritedDHCPOption)
}
