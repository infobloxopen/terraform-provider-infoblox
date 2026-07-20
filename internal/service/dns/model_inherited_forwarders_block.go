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

// InheritedForwardersBlockModel is the Terraform model for InheritedForwardersBlock
type InheritedForwardersBlockModel struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.Object `tfsdk:"value"`
}

// InheritedForwardersBlockAttrTypes contains the attribute types for InheritedForwardersBlockModel
var InheritedForwardersBlockAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.ObjectType{AttrTypes: ForwardersBlockAttrTypes},
}

// InheritedForwardersBlockResourceSchemaAttributes contains the schema attributes for InheritedForwardersBlockModel
var InheritedForwardersBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Defaults to _inherit_.",
	},
	"display_name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Human-readable display name for the object referred to by _source_.",
	},
	"source": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"value": schema.SingleNestedAttribute{
		Attributes:          ForwardersBlockResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Inherited value.",
	},
}

// ExpandInheritedForwardersBlock converts a Terraform Object to SDK type
func ExpandInheritedForwardersBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.InheritedForwardersBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritedForwardersBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *InheritedForwardersBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.InheritedForwardersBlock {
	if m == nil {
		return nil
	}
	to := &uddidns.InheritedForwardersBlock{
		Action:      flex.ExpandStringPointer(m.Action),
		DisplayName: flex.ExpandStringPointer(m.DisplayName),
		Source:      flex.ExpandStringPointer(m.Source),
		Value:       ExpandForwardersBlock(ctx, m.Value, diags),
	}
	return to
}

// FlattenInheritedForwardersBlock converts an SDK type to Terraform Object
func FlattenInheritedForwardersBlock(ctx context.Context, from *uddidns.InheritedForwardersBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritedForwardersBlockAttrTypes)
	}
	m := &InheritedForwardersBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritedForwardersBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *InheritedForwardersBlockModel) Flatten(ctx context.Context, from *uddidns.InheritedForwardersBlock, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.DisplayName = flex.FlattenStringPointer(from.DisplayName)
	m.Source = flex.FlattenStringPointer(from.Source)
	m.Value = FlattenForwardersBlock(ctx, from.Value, diags)
}
