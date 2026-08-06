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

// InheritedAsmEnableBlockModel is the Terraform model for InheritedAsmEnableBlock
type InheritedAsmEnableBlockModel struct {
	Action types.String `tfsdk:"action"`
}

// InheritedAsmEnableBlockAttrTypes contains the attribute types for InheritedAsmEnableBlockModel
var InheritedAsmEnableBlockAttrTypes = map[string]attr.Type{
	"action": types.StringType,
}

// InheritedAsmEnableBlockResourceSchemaAttributes contains the schema attributes for InheritedAsmEnableBlockModel
var InheritedAsmEnableBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance setting.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.",
	},
}

// ExpandInheritedAsmEnableBlock converts a Terraform Object to SDK type
func ExpandInheritedAsmEnableBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiipam.InheritedAsmEnableBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritedAsmEnableBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *InheritedAsmEnableBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiipam.InheritedAsmEnableBlock {
	if m == nil {
		return nil
	}
	to := &uddiipam.InheritedAsmEnableBlock{
		Action: flex.ExpandStringPointer(m.Action),
	}
	return to
}

// FlattenInheritedAsmEnableBlock converts an SDK type to Terraform Object
func FlattenInheritedAsmEnableBlock(ctx context.Context, from *uddiipam.InheritedAsmEnableBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritedAsmEnableBlockAttrTypes)
	}
	m := &InheritedAsmEnableBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritedAsmEnableBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *InheritedAsmEnableBlockModel) Flatten(ctx context.Context, from *uddiipam.InheritedAsmEnableBlock, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
}
