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

// InheritedHostnameRewriteBlockModel is the Terraform model for InheritedHostnameRewriteBlock
type InheritedHostnameRewriteBlockModel struct {
	Action types.String `tfsdk:"action"`
}

// InheritedHostnameRewriteBlockAttrTypes contains the attribute types for InheritedHostnameRewriteBlockModel
var InheritedHostnameRewriteBlockAttrTypes = map[string]attr.Type{
	"action": types.StringType,
}

// InheritedHostnameRewriteBlockResourceSchemaAttributes contains the schema attributes for InheritedHostnameRewriteBlockModel
var InheritedHostnameRewriteBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance setting.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.",
	},
}

// ExpandInheritedHostnameRewriteBlock converts a Terraform Object to SDK type
func ExpandInheritedHostnameRewriteBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiipam.InheritedHostnameRewriteBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritedHostnameRewriteBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *InheritedHostnameRewriteBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiipam.InheritedHostnameRewriteBlock {
	if m == nil {
		return nil
	}
	to := &uddiipam.InheritedHostnameRewriteBlock{
		Action: flex.ExpandStringPointer(m.Action),
	}
	return to
}

// FlattenInheritedHostnameRewriteBlock converts an SDK type to Terraform Object
func FlattenInheritedHostnameRewriteBlock(ctx context.Context, from *uddiipam.InheritedHostnameRewriteBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritedHostnameRewriteBlockAttrTypes)
	}
	m := &InheritedHostnameRewriteBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritedHostnameRewriteBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *InheritedHostnameRewriteBlockModel) Flatten(ctx context.Context, from *uddiipam.InheritedHostnameRewriteBlock, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
}
