package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddidns "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// Inheritance2InheritedBoolModel is the Terraform model for Inheritance2InheritedBool
type Inheritance2InheritedBoolModel struct {
	Action types.String `tfsdk:"action"`
}

// Inheritance2InheritedBoolAttrTypes contains the attribute types for Inheritance2InheritedBoolModel
var Inheritance2InheritedBoolAttrTypes = map[string]attr.Type{
	"action": types.StringType,
}

// Inheritance2InheritedBoolResourceSchemaAttributes contains the schema attributes for Inheritance2InheritedBoolModel
var Inheritance2InheritedBoolResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance setting for a field.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.",
	},
}

// ExpandInheritance2InheritedBool converts a Terraform Object to SDK type
func ExpandInheritance2InheritedBool(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.Inheritance2InheritedBool {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Inheritance2InheritedBoolModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Inheritance2InheritedBoolModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.Inheritance2InheritedBool {
	if m == nil {
		return nil
	}
	to := &uddidns.Inheritance2InheritedBool{
		Action: flex.ExpandStringPointer(m.Action),
	}
	return to
}

// FlattenInheritance2InheritedBool converts an SDK type to Terraform Object
func FlattenInheritance2InheritedBool(ctx context.Context, from *uddidns.Inheritance2InheritedBool, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Inheritance2InheritedBoolAttrTypes)
	}
	m := &Inheritance2InheritedBoolModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Inheritance2InheritedBoolAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Inheritance2InheritedBoolModel) Flatten(ctx context.Context, from *uddidns.Inheritance2InheritedBool, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
}
