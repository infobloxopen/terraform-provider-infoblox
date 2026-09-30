package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	objectplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	uddidhcp "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// DHCPOptionsInheritanceModel is the Terraform model for DHCPOptionsInheritance
type DHCPOptionsInheritanceModel struct {
	DhcpOptions types.Object `tfsdk:"dhcp_options"`
}

// DHCPOptionsInheritanceAttrTypes contains the attribute types for DHCPOptionsInheritanceModel
var DHCPOptionsInheritanceAttrTypes = map[string]attr.Type{
	"dhcp_options": types.ObjectType{AttrTypes: InheritedDHCPOptionListAttrTypes},
}

// DHCPOptionsInheritanceResourceSchemaAttributes contains the schema attributes for DHCPOptionsInheritanceModel
var DHCPOptionsInheritanceResourceSchemaAttributes = map[string]schema.Attribute{
	"dhcp_options": schema.SingleNestedAttribute{
		Attributes: InheritedDHCPOptionListResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		PlanModifiers: []planmodifier.Object{
			objectplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The inheritance configuration for the _dhcp_options_ field.",
	},
}

// ExpandDHCPOptionsInheritance converts a Terraform Object to SDK type
func ExpandDHCPOptionsInheritance(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidhcp.DHCPOptionsInheritance {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DHCPOptionsInheritanceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *DHCPOptionsInheritanceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidhcp.DHCPOptionsInheritance {
	if m == nil {
		return nil
	}
	to := &uddidhcp.DHCPOptionsInheritance{
		DhcpOptions: ExpandInheritedDHCPOptionList(ctx, m.DhcpOptions, diags),
	}
	return to
}

// FlattenDHCPOptionsInheritance converts an SDK type to Terraform Object
func FlattenDHCPOptionsInheritance(ctx context.Context, from *uddidhcp.DHCPOptionsInheritance, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DHCPOptionsInheritanceAttrTypes)
	}
	m := &DHCPOptionsInheritanceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DHCPOptionsInheritanceAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *DHCPOptionsInheritanceModel) Flatten(ctx context.Context, from *uddidhcp.DHCPOptionsInheritance, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.DhcpOptions = FlattenInheritedDHCPOptionList(ctx, from.DhcpOptions, diags)
}
