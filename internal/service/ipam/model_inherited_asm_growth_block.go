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

// InheritedAsmGrowthBlockModel is the Terraform model for InheritedAsmGrowthBlock
type InheritedAsmGrowthBlockModel struct {
	Action types.String `tfsdk:"action"`
}

// InheritedAsmGrowthBlockAttrTypes contains the attribute types for InheritedAsmGrowthBlockModel
var InheritedAsmGrowthBlockAttrTypes = map[string]attr.Type{
	"action": types.StringType,
}

// InheritedAsmGrowthBlockResourceSchemaAttributes contains the schema attributes for InheritedAsmGrowthBlockModel
var InheritedAsmGrowthBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance setting.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.",
	},
}

// ExpandInheritedAsmGrowthBlock converts a Terraform Object to SDK type
func ExpandInheritedAsmGrowthBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiipam.InheritedAsmGrowthBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritedAsmGrowthBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *InheritedAsmGrowthBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiipam.InheritedAsmGrowthBlock {
	if m == nil {
		return nil
	}
	to := &uddiipam.InheritedAsmGrowthBlock{
		Action: flex.ExpandStringPointer(m.Action),
	}
	return to
}

// FlattenInheritedAsmGrowthBlock converts an SDK type to Terraform Object
func FlattenInheritedAsmGrowthBlock(ctx context.Context, from *uddiipam.InheritedAsmGrowthBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritedAsmGrowthBlockAttrTypes)
	}
	m := &InheritedAsmGrowthBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritedAsmGrowthBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *InheritedAsmGrowthBlockModel) Flatten(ctx context.Context, from *uddiipam.InheritedAsmGrowthBlock, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
}
