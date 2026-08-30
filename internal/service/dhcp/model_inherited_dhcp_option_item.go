package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddidhcp "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// InheritedDHCPOptionItemModel is the Terraform model for InheritedDHCPOptionItem
type InheritedDHCPOptionItemModel struct {
	Option          types.Object `tfsdk:"option"`
	OverridingGroup types.String `tfsdk:"overriding_group"`
}

// InheritedDHCPOptionItemAttrTypes contains the attribute types for InheritedDHCPOptionItemModel
var InheritedDHCPOptionItemAttrTypes = map[string]attr.Type{
	"option":           types.ObjectType{AttrTypes: OptionItemAttrTypes},
	"overriding_group": types.StringType,
}

// InheritedDHCPOptionItemResourceSchemaAttributes contains the schema attributes for InheritedDHCPOptionItemModel
var InheritedDHCPOptionItemResourceSchemaAttributes = map[string]schema.Attribute{
	"option": schema.SingleNestedAttribute{
		Attributes:          OptionItemResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Option inherited from the ancestor.",
	},
	"overriding_group": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
}

// ExpandInheritedDHCPOptionItem converts a Terraform Object to SDK type
func ExpandInheritedDHCPOptionItem(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidhcp.InheritedDHCPOptionItem {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritedDHCPOptionItemModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *InheritedDHCPOptionItemModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidhcp.InheritedDHCPOptionItem {
	if m == nil {
		return nil
	}
	to := &uddidhcp.InheritedDHCPOptionItem{
		Option:          ExpandOptionItem(ctx, m.Option, diags),
		OverridingGroup: flex.ExpandStringPointer(m.OverridingGroup),
	}
	return to
}

// FlattenInheritedDHCPOptionItem converts an SDK type to Terraform Object
func FlattenInheritedDHCPOptionItem(ctx context.Context, from *uddidhcp.InheritedDHCPOptionItem, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritedDHCPOptionItemAttrTypes)
	}
	m := &InheritedDHCPOptionItemModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritedDHCPOptionItemAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *InheritedDHCPOptionItemModel) Flatten(ctx context.Context, from *uddidhcp.InheritedDHCPOptionItem, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Option = FlattenOptionItem(ctx, from.Option, diags)
	m.OverridingGroup = flex.FlattenStringPointer(from.OverridingGroup)
}
