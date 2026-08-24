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

// InheritedDDNSHostnameBlockModel is the Terraform model for InheritedDDNSHostnameBlock
type InheritedDDNSHostnameBlockModel struct {
	Action types.String `tfsdk:"action"`
}

// InheritedDDNSHostnameBlockAttrTypes contains the attribute types for InheritedDDNSHostnameBlockModel
var InheritedDDNSHostnameBlockAttrTypes = map[string]attr.Type{
	"action": types.StringType,
}

// InheritedDDNSHostnameBlockResourceSchemaAttributes contains the schema attributes for InheritedDDNSHostnameBlockModel
var InheritedDDNSHostnameBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance setting.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.",
	},
}

// ExpandInheritedDDNSHostnameBlock converts a Terraform Object to SDK type
func ExpandInheritedDDNSHostnameBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiipam.InheritedDDNSHostnameBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritedDDNSHostnameBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *InheritedDDNSHostnameBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiipam.InheritedDDNSHostnameBlock {
	if m == nil {
		return nil
	}
	to := &uddiipam.InheritedDDNSHostnameBlock{
		Action: flex.ExpandStringPointer(m.Action),
	}
	return to
}

// FlattenInheritedDDNSHostnameBlock converts an SDK type to Terraform Object
func FlattenInheritedDDNSHostnameBlock(ctx context.Context, from *uddiipam.InheritedDDNSHostnameBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritedDDNSHostnameBlockAttrTypes)
	}
	m := &InheritedDDNSHostnameBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritedDDNSHostnameBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *InheritedDDNSHostnameBlockModel) Flatten(ctx context.Context, from *uddiipam.InheritedDDNSHostnameBlock, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
}
