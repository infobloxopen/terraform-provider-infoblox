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

// InheritedDHCPConfigIgnoreItemListModel is the Terraform model for InheritedDHCPConfigIgnoreItemList
type InheritedDHCPConfigIgnoreItemListModel struct {
	Action types.String `tfsdk:"action"`
}

// InheritedDHCPConfigIgnoreItemListAttrTypes contains the attribute types for InheritedDHCPConfigIgnoreItemListModel
var InheritedDHCPConfigIgnoreItemListAttrTypes = map[string]attr.Type{
	"action": types.StringType,
}

// InheritedDHCPConfigIgnoreItemListResourceSchemaAttributes contains the schema attributes for InheritedDHCPConfigIgnoreItemListModel
var InheritedDHCPConfigIgnoreItemListResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance setting.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.",
	},
}

// ExpandInheritedDHCPConfigIgnoreItemList converts a Terraform Object to SDK type
func ExpandInheritedDHCPConfigIgnoreItemList(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiipam.InheritedDHCPConfigIgnoreItemList {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritedDHCPConfigIgnoreItemListModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *InheritedDHCPConfigIgnoreItemListModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiipam.InheritedDHCPConfigIgnoreItemList {
	if m == nil {
		return nil
	}
	to := &uddiipam.InheritedDHCPConfigIgnoreItemList{
		Action: flex.ExpandStringPointer(m.Action),
	}
	return to
}

// FlattenInheritedDHCPConfigIgnoreItemList converts an SDK type to Terraform Object
func FlattenInheritedDHCPConfigIgnoreItemList(ctx context.Context, from *uddiipam.InheritedDHCPConfigIgnoreItemList, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritedDHCPConfigIgnoreItemListAttrTypes)
	}
	m := &InheritedDHCPConfigIgnoreItemListModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritedDHCPConfigIgnoreItemListAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *InheritedDHCPConfigIgnoreItemListModel) Flatten(ctx context.Context, from *uddiipam.InheritedDHCPConfigIgnoreItemList, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
}
