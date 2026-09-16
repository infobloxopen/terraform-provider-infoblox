package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddidtc "github.com/infobloxopen/universal-ddi-go-client/dtc"
)

// Inheritance2InheritedUInt32Model is the Terraform model for Inheritance2InheritedUInt32
type Inheritance2InheritedUInt32Model struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.Int64  `tfsdk:"value"`
}

// Inheritance2InheritedUInt32AttrTypes contains the attribute types for Inheritance2InheritedUInt32Model
var Inheritance2InheritedUInt32AttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.Int64Type,
}

// Inheritance2InheritedUInt32ResourceSchemaAttributes contains the schema attributes for Inheritance2InheritedUInt32Model
var Inheritance2InheritedUInt32ResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance setting for a field.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.",
	},
	"display_name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The human-readable display name for the object referred to by _source_.",
	},
	"source": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"value": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inherited value.",
	},
}

// ExpandInheritance2InheritedUInt32 converts a Terraform Object to SDK type
func ExpandInheritance2InheritedUInt32(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidtc.Inheritance2InheritedUInt32 {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Inheritance2InheritedUInt32Model
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Inheritance2InheritedUInt32Model) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidtc.Inheritance2InheritedUInt32 {
	if m == nil {
		return nil
	}
	to := &uddidtc.Inheritance2InheritedUInt32{
		Action:      flex.ExpandStringPointer(m.Action),
		DisplayName: flex.ExpandStringPointer(m.DisplayName),
		Source:      flex.ExpandStringPointer(m.Source),
		Value:       flex.ExpandInt64Pointer(m.Value),
	}
	return to
}

// FlattenInheritance2InheritedUInt32 converts an SDK type to Terraform Object
func FlattenInheritance2InheritedUInt32(ctx context.Context, from *uddidtc.Inheritance2InheritedUInt32, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Inheritance2InheritedUInt32AttrTypes)
	}
	m := &Inheritance2InheritedUInt32Model{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Inheritance2InheritedUInt32AttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Inheritance2InheritedUInt32Model) Flatten(ctx context.Context, from *uddidtc.Inheritance2InheritedUInt32, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.DisplayName = flex.FlattenStringPointer(from.DisplayName)
	m.Source = flex.FlattenStringPointer(from.Source)
	m.Value = flex.FlattenInt64Pointer(from.Value)
}
