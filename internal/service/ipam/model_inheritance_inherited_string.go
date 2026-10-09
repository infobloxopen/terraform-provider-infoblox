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

// InheritanceInheritedStringModel is the Terraform model for InheritanceInheritedString
type InheritanceInheritedStringModel struct {
	Action types.String `tfsdk:"action"`
}

// InheritanceInheritedStringAttrTypes contains the attribute types for InheritanceInheritedStringModel
var InheritanceInheritedStringAttrTypes = map[string]attr.Type{
	"action": types.StringType,
}

// InheritanceInheritedStringResourceSchemaAttributes contains the schema attributes for InheritanceInheritedStringModel
var InheritanceInheritedStringResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance setting for a field.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.",
	},
}

// ExpandInheritanceInheritedString converts a Terraform Object to SDK type
func ExpandInheritanceInheritedString(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiipam.InheritanceInheritedString {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritanceInheritedStringModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *InheritanceInheritedStringModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiipam.InheritanceInheritedString {
	if m == nil {
		return nil
	}
	to := &uddiipam.InheritanceInheritedString{
		Action: flex.ExpandStringPointer(m.Action),
	}
	return to
}

// FlattenInheritanceInheritedString converts an SDK type to Terraform Object
func FlattenInheritanceInheritedString(ctx context.Context, from *uddiipam.InheritanceInheritedString, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritanceInheritedStringAttrTypes)
	}
	m := &InheritanceInheritedStringModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritanceInheritedStringAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *InheritanceInheritedStringModel) Flatten(ctx context.Context, from *uddiipam.InheritanceInheritedString, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
}
