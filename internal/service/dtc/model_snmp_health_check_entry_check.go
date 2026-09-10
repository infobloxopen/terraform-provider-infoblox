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

// SNMPHealthCheckEntryCheckModel is the Terraform model for SNMPHealthCheckEntryCheck
type SNMPHealthCheckEntryCheckModel struct {
	Comment  types.String `tfsdk:"comment"`
	MaxValue types.String `tfsdk:"max_value"`
	Name     types.String `tfsdk:"name"`
	Operator types.String `tfsdk:"operator"`
	Type     types.String `tfsdk:"type"`
	Value    types.String `tfsdk:"value"`
}

// SNMPHealthCheckEntryCheckAttrTypes contains the attribute types for SNMPHealthCheckEntryCheckModel
var SNMPHealthCheckEntryCheckAttrTypes = map[string]attr.Type{
	"comment":   types.StringType,
	"max_value": types.StringType,
	"name":      types.StringType,
	"operator":  types.StringType,
	"type":      types.StringType,
	"value":     types.StringType,
}

// SNMPHealthCheckEntryCheckResourceSchemaAttributes contains the schema attributes for SNMPHealthCheckEntryCheckModel
var SNMPHealthCheckEntryCheckResourceSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Optional. Comment for __EntryCheck__.",
	},
	"max_value": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Optional. Expected max value of an entry to check against. Used for __in__ operator only, otherwise ignored.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Name is a dotted-decimal number that defines the location of the entry in the universal MIB tree.",
	},
	"operator": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Operator defines operation to perform on an entry value.  Allowed values: * any - any value must be present * eq  - entry value must be equal to check's __value__. * leq - entry value must less or equal to check's __value__. * geq - entry value must be great or equal to check's __value__. * in  - entry value must be greater or equal than __value__ and less or equal than __max_value__.  Operator __in__ is supported only for __integer__ types.",
	},
	"type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Type defines type of an entry value.  Allowed values: * string * integer  String type does not support __in__ operator.",
	},
	"value": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Optional. Expected value of an entry to check against. Ignored for __any__ operator.",
	},
}

// ExpandSNMPHealthCheckEntryCheck converts a Terraform Object to SDK type
func ExpandSNMPHealthCheckEntryCheck(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidtc.SNMPHealthCheckEntryCheck {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SNMPHealthCheckEntryCheckModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *SNMPHealthCheckEntryCheckModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidtc.SNMPHealthCheckEntryCheck {
	if m == nil {
		return nil
	}
	to := &uddidtc.SNMPHealthCheckEntryCheck{
		Comment:  flex.ExpandStringPointer(m.Comment),
		MaxValue: flex.ExpandStringPointer(m.MaxValue),
		Name:     flex.ExpandString(m.Name),
		Operator: flex.ExpandString(m.Operator),
		Type:     flex.ExpandString(m.Type),
		Value:    flex.ExpandStringPointer(m.Value),
	}
	return to
}

// FlattenSNMPHealthCheckEntryCheck converts an SDK type to Terraform Object
func FlattenSNMPHealthCheckEntryCheck(ctx context.Context, from *uddidtc.SNMPHealthCheckEntryCheck, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SNMPHealthCheckEntryCheckAttrTypes)
	}
	m := &SNMPHealthCheckEntryCheckModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SNMPHealthCheckEntryCheckAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *SNMPHealthCheckEntryCheckModel) Flatten(ctx context.Context, from *uddidtc.SNMPHealthCheckEntryCheck, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.MaxValue = flex.FlattenStringPointer(from.MaxValue)
	m.Name = flex.FlattenString(from.Name)
	m.Operator = flex.FlattenString(from.Operator)
	m.Type = flex.FlattenString(from.Type)
	m.Value = flex.FlattenStringPointer(from.Value)
}
