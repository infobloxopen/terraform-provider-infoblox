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

// InheritanceInheritedIdentifierModel is the Terraform model for InheritanceInheritedIdentifier
type InheritanceInheritedIdentifierModel struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.String `tfsdk:"value"`
}

// InheritanceInheritedIdentifierAttrTypes contains the attribute types for InheritanceInheritedIdentifierModel
var InheritanceInheritedIdentifierAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.StringType,
}

// InheritanceInheritedIdentifierResourceSchemaAttributes contains the schema attributes for InheritanceInheritedIdentifierModel
var InheritanceInheritedIdentifierResourceSchemaAttributes = map[string]schema.Attribute{
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
	"value": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
}

// ExpandInheritanceInheritedIdentifier converts a Terraform Object to SDK type
func ExpandInheritanceInheritedIdentifier(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiipam.InheritanceInheritedIdentifier {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritanceInheritedIdentifierModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *InheritanceInheritedIdentifierModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiipam.InheritanceInheritedIdentifier {
	if m == nil {
		return nil
	}
	to := &uddiipam.InheritanceInheritedIdentifier{
		Action:      flex.ExpandStringPointer(m.Action),
		DisplayName: flex.ExpandStringPointer(m.DisplayName),
		Source:      flex.ExpandStringPointer(m.Source),
		Value:       flex.ExpandStringPointer(m.Value),
	}
	return to
}

// FlattenInheritanceInheritedIdentifier converts an SDK type to Terraform Object
func FlattenInheritanceInheritedIdentifier(ctx context.Context, from *uddiipam.InheritanceInheritedIdentifier, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritanceInheritedIdentifierAttrTypes)
	}
	m := &InheritanceInheritedIdentifierModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritanceInheritedIdentifierAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *InheritanceInheritedIdentifierModel) Flatten(ctx context.Context, from *uddiipam.InheritanceInheritedIdentifier, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.DisplayName = flex.FlattenStringPointer(from.DisplayName)
	m.Source = flex.FlattenStringPointer(from.Source)
	m.Value = flex.FlattenStringPointer(from.Value)
}
