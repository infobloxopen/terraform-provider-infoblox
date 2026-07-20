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

// Inheritance2InheritedUInt32DnsconfigModel is the Terraform model for Inheritance2InheritedUInt32
type Inheritance2InheritedUInt32DnsconfigModel struct {
	Action types.String `tfsdk:"action"`
}

// Inheritance2InheritedUInt32DnsconfigAttrTypes contains the attribute types for Inheritance2InheritedUInt32DnsconfigModel
var Inheritance2InheritedUInt32DnsconfigAttrTypes = map[string]attr.Type{
	"action": types.StringType,
}

// Inheritance2InheritedUInt32DnsconfigResourceSchemaAttributes contains the schema attributes for Inheritance2InheritedUInt32DnsconfigModel
var Inheritance2InheritedUInt32DnsconfigResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance setting for a field.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.",
	},
}

// ExpandInheritance2InheritedUInt32Dnsconfig converts a Terraform Object to SDK type
func ExpandInheritance2InheritedUInt32Dnsconfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.Inheritance2InheritedUInt32 {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Inheritance2InheritedUInt32DnsconfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Inheritance2InheritedUInt32DnsconfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.Inheritance2InheritedUInt32 {
	if m == nil {
		return nil
	}
	to := &uddidns.Inheritance2InheritedUInt32{
		Action: flex.ExpandStringPointer(m.Action),
	}
	return to
}

// FlattenInheritance2InheritedUInt32Dnsconfig converts an SDK type to Terraform Object
func FlattenInheritance2InheritedUInt32Dnsconfig(ctx context.Context, from *uddidns.Inheritance2InheritedUInt32, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Inheritance2InheritedUInt32DnsconfigAttrTypes)
	}
	m := &Inheritance2InheritedUInt32DnsconfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Inheritance2InheritedUInt32DnsconfigAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Inheritance2InheritedUInt32DnsconfigModel) Flatten(ctx context.Context, from *uddidns.Inheritance2InheritedUInt32, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
}
