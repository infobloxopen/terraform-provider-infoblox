package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// InheritanceInheritedFloatModel is the Terraform model for InheritanceInheritedFloat
type InheritanceInheritedFloatModel struct {
	Action types.String `tfsdk:"action"`
}

// InheritanceInheritedFloatAttrTypes contains the attribute types for InheritanceInheritedFloatModel
var InheritanceInheritedFloatAttrTypes = map[string]attr.Type{
	"action": types.StringType,
}

// InheritanceInheritedFloatResourceSchemaAttributes contains the schema attributes for InheritanceInheritedFloatModel
var InheritanceInheritedFloatResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance setting for a field.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.",
	},
}

// ExpandInheritanceInheritedFloat converts a Terraform Object to SDK type
func ExpandInheritanceInheritedFloat(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiipam.InheritanceInheritedFloat {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritanceInheritedFloatModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *InheritanceInheritedFloatModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiipam.InheritanceInheritedFloat {
	if m == nil {
		return nil
	}
	to := &uddiipam.InheritanceInheritedFloat{
		Action: flex.ExpandStringPointer(m.Action),
	}
	return to
}

// FlattenInheritanceInheritedFloat converts an SDK type to Terraform Object
func FlattenInheritanceInheritedFloat(ctx context.Context, from *uddiipam.InheritanceInheritedFloat, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritanceInheritedFloatAttrTypes)
	}
	m := &InheritanceInheritedFloatModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritanceInheritedFloatAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *InheritanceInheritedFloatModel) Flatten(ctx context.Context, from *uddiipam.InheritanceInheritedFloat, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
}
