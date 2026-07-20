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

// InheritedCustomRootNSBlockModel is the Terraform model for InheritedCustomRootNSBlock
type InheritedCustomRootNSBlockModel struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.Object `tfsdk:"value"`
}

// InheritedCustomRootNSBlockAttrTypes contains the attribute types for InheritedCustomRootNSBlockModel
var InheritedCustomRootNSBlockAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.ObjectType{AttrTypes: CustomRootNSBlockAttrTypes},
}

// InheritedCustomRootNSBlockResourceSchemaAttributes contains the schema attributes for InheritedCustomRootNSBlockModel
var InheritedCustomRootNSBlockResourceSchemaAttributes = map[string]schema.Attribute{
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
		Attributes:          CustomRootNSBlockResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Inherited value.",
	},
}

// ExpandInheritedCustomRootNSBlock converts a Terraform Object to SDK type
func ExpandInheritedCustomRootNSBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.InheritedCustomRootNSBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritedCustomRootNSBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *InheritedCustomRootNSBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.InheritedCustomRootNSBlock {
	if m == nil {
		return nil
	}
	to := &uddidns.InheritedCustomRootNSBlock{
		Action:      flex.ExpandStringPointer(m.Action),
		DisplayName: flex.ExpandStringPointer(m.DisplayName),
		Source:      flex.ExpandStringPointer(m.Source),
		Value:       ExpandCustomRootNSBlock(ctx, m.Value, diags),
	}
	return to
}

// FlattenInheritedCustomRootNSBlock converts an SDK type to Terraform Object
func FlattenInheritedCustomRootNSBlock(ctx context.Context, from *uddidns.InheritedCustomRootNSBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritedCustomRootNSBlockAttrTypes)
	}
	m := &InheritedCustomRootNSBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritedCustomRootNSBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *InheritedCustomRootNSBlockModel) Flatten(ctx context.Context, from *uddidns.InheritedCustomRootNSBlock, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.DisplayName = flex.FlattenStringPointer(from.DisplayName)
	m.Source = flex.FlattenStringPointer(from.Source)
	m.Value = FlattenCustomRootNSBlock(ctx, from.Value, diags)
}
