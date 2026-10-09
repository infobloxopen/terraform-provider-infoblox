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

// InheritanceInheritedBoolModel is the Terraform model for InheritanceInheritedBool
type InheritanceInheritedBoolModel struct {
	Action types.String `tfsdk:"action"`
}

// InheritanceInheritedBoolAttrTypes contains the attribute types for InheritanceInheritedBoolModel
var InheritanceInheritedBoolAttrTypes = map[string]attr.Type{
	"action": types.StringType,
}

// InheritanceInheritedBoolResourceSchemaAttributes contains the schema attributes for InheritanceInheritedBoolModel
var InheritanceInheritedBoolResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance setting for a field.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.",
	},
}

// ExpandInheritanceInheritedBool converts a Terraform Object to SDK type
func ExpandInheritanceInheritedBool(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiipam.InheritanceInheritedBool {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritanceInheritedBoolModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *InheritanceInheritedBoolModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiipam.InheritanceInheritedBool {
	if m == nil {
		return nil
	}
	to := &uddiipam.InheritanceInheritedBool{
		Action: flex.ExpandStringPointer(m.Action),
	}
	return to
}

// FlattenInheritanceInheritedBool converts an SDK type to Terraform Object
func FlattenInheritanceInheritedBool(ctx context.Context, from *uddiipam.InheritanceInheritedBool, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritanceInheritedBoolAttrTypes)
	}
	m := &InheritanceInheritedBoolModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritanceInheritedBoolAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *InheritanceInheritedBoolModel) Flatten(ctx context.Context, from *uddiipam.InheritanceInheritedBool, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
}
