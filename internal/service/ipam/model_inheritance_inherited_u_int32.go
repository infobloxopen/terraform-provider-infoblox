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

// InheritanceInheritedUInt32Model is the Terraform model for InheritanceInheritedUInt32
type InheritanceInheritedUInt32Model struct {
	Action types.String `tfsdk:"action"`
}

// InheritanceInheritedUInt32AttrTypes contains the attribute types for InheritanceInheritedUInt32Model
var InheritanceInheritedUInt32AttrTypes = map[string]attr.Type{
	"action": types.StringType,
}

// InheritanceInheritedUInt32ResourceSchemaAttributes contains the schema attributes for InheritanceInheritedUInt32Model
var InheritanceInheritedUInt32ResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance setting for a field.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.",
	},
}

// ExpandInheritanceInheritedUInt32 converts a Terraform Object to SDK type
func ExpandInheritanceInheritedUInt32(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiipam.InheritanceInheritedUInt32 {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritanceInheritedUInt32Model
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *InheritanceInheritedUInt32Model) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiipam.InheritanceInheritedUInt32 {
	if m == nil {
		return nil
	}
	to := &uddiipam.InheritanceInheritedUInt32{
		Action: flex.ExpandStringPointer(m.Action),
	}
	return to
}

// FlattenInheritanceInheritedUInt32 converts an SDK type to Terraform Object
func FlattenInheritanceInheritedUInt32(ctx context.Context, from *uddiipam.InheritanceInheritedUInt32, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritanceInheritedUInt32AttrTypes)
	}
	m := &InheritanceInheritedUInt32Model{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritanceInheritedUInt32AttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *InheritanceInheritedUInt32Model) Flatten(ctx context.Context, from *uddiipam.InheritanceInheritedUInt32, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
}
