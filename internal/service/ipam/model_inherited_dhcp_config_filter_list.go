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

// InheritedDHCPConfigFilterListModel is the Terraform model for InheritedDHCPConfigFilterList
type InheritedDHCPConfigFilterListModel struct {
	Action types.String `tfsdk:"action"`
}

// InheritedDHCPConfigFilterListAttrTypes contains the attribute types for InheritedDHCPConfigFilterListModel
var InheritedDHCPConfigFilterListAttrTypes = map[string]attr.Type{
	"action": types.StringType,
}

// InheritedDHCPConfigFilterListResourceSchemaAttributes contains the schema attributes for InheritedDHCPConfigFilterListModel
var InheritedDHCPConfigFilterListResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance setting.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.",
	},
}

// ExpandInheritedDHCPConfigFilterList converts a Terraform Object to SDK type
func ExpandInheritedDHCPConfigFilterList(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiipam.InheritedDHCPConfigFilterList {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritedDHCPConfigFilterListModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *InheritedDHCPConfigFilterListModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiipam.InheritedDHCPConfigFilterList {
	if m == nil {
		return nil
	}
	to := &uddiipam.InheritedDHCPConfigFilterList{
		Action: flex.ExpandStringPointer(m.Action),
	}
	return to
}

// FlattenInheritedDHCPConfigFilterList converts an SDK type to Terraform Object
func FlattenInheritedDHCPConfigFilterList(ctx context.Context, from *uddiipam.InheritedDHCPConfigFilterList, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritedDHCPConfigFilterListAttrTypes)
	}
	m := &InheritedDHCPConfigFilterListModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritedDHCPConfigFilterListAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *InheritedDHCPConfigFilterListModel) Flatten(ctx context.Context, from *uddiipam.InheritedDHCPConfigFilterList, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
}
