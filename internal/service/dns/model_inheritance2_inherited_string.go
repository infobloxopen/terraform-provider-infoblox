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

// Inheritance2InheritedStringModel is the Terraform model for Inheritance2InheritedString
type Inheritance2InheritedStringModel struct {
	Action types.String `tfsdk:"action"`
}

// Inheritance2InheritedStringAttrTypes contains the attribute types for Inheritance2InheritedStringModel
var Inheritance2InheritedStringAttrTypes = map[string]attr.Type{
	"action": types.StringType,
}

// Inheritance2InheritedStringResourceSchemaAttributes contains the schema attributes for Inheritance2InheritedStringModel
var Inheritance2InheritedStringResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance setting for a field.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.",
	},
}

// ExpandInheritance2InheritedString converts a Terraform Object to SDK type
func ExpandInheritance2InheritedString(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.Inheritance2InheritedString {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Inheritance2InheritedStringModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Inheritance2InheritedStringModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.Inheritance2InheritedString {
	if m == nil {
		return nil
	}
	to := &uddidns.Inheritance2InheritedString{
		Action: flex.ExpandStringPointer(m.Action),
	}
	return to
}

// FlattenInheritance2InheritedString converts an SDK type to Terraform Object
func FlattenInheritance2InheritedString(ctx context.Context, from *uddidns.Inheritance2InheritedString, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Inheritance2InheritedStringAttrTypes)
	}
	m := &Inheritance2InheritedStringModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Inheritance2InheritedStringAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Inheritance2InheritedStringModel) Flatten(ctx context.Context, from *uddidns.Inheritance2InheritedString, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
}
