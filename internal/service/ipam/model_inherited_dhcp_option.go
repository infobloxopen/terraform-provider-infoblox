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

// InheritedDHCPOptionModel is the Terraform model for InheritedDHCPOption
type InheritedDHCPOptionModel struct {
	Action types.String `tfsdk:"action"`
}

// InheritedDHCPOptionAttrTypes contains the attribute types for InheritedDHCPOptionModel
var InheritedDHCPOptionAttrTypes = map[string]attr.Type{
	"action": types.StringType,
}

// InheritedDHCPOptionResourceSchemaAttributes contains the schema attributes for InheritedDHCPOptionModel
var InheritedDHCPOptionResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance setting.  Valid values are: * _inherit_: Use the inherited value. * _block_: Don't use the inherited value.  Defaults to _inherit_.",
	},
}

// ExpandInheritedDHCPOption converts a Terraform Object to SDK type
func ExpandInheritedDHCPOption(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiipam.InheritedDHCPOption {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritedDHCPOptionModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *InheritedDHCPOptionModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiipam.InheritedDHCPOption {
	if m == nil {
		return nil
	}
	to := &uddiipam.InheritedDHCPOption{
		Action: flex.ExpandStringPointer(m.Action),
	}
	return to
}

// FlattenInheritedDHCPOption converts an SDK type to Terraform Object
func FlattenInheritedDHCPOption(ctx context.Context, from *uddiipam.InheritedDHCPOption, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritedDHCPOptionAttrTypes)
	}
	m := &InheritedDHCPOptionModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritedDHCPOptionAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *InheritedDHCPOptionModel) Flatten(ctx context.Context, from *uddiipam.InheritedDHCPOption, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
}
