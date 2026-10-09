package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// SharednetworkNetworksModel is the Terraform model for SharednetworkNetworks
type SharednetworkNetworksModel struct {
	Ref types.String `tfsdk:"ref"`
}

// SharednetworkNetworksAttrTypes contains the attribute types for SharednetworkNetworksModel
var SharednetworkNetworksAttrTypes = map[string]attr.Type{
	"ref": types.StringType,
}

// SharednetworkNetworksResourceSchemaAttributes contains the schema attributes for SharednetworkNetworksModel
var SharednetworkNetworksResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Reference to the Network.",
	},
}

// ExpandSharednetworkNetworks converts a Terraform Object to SDK type
func ExpandSharednetworkNetworks(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.SharednetworkNetworks {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SharednetworkNetworksModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *SharednetworkNetworksModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.SharednetworkNetworks {
	if m == nil {
		return nil
	}
	to := &niosdhcp.SharednetworkNetworks{
		Ref: flex.ExpandStringPointerNullAsEmpty(m.Ref),
	}
	return to
}

// FlattenSharednetworkNetworks converts an SDK type to Terraform Object
func FlattenSharednetworkNetworks(ctx context.Context, from *niosdhcp.SharednetworkNetworks, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SharednetworkNetworksAttrTypes)
	}
	m := &SharednetworkNetworksModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SharednetworkNetworksAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *SharednetworkNetworksModel) Flatten(ctx context.Context, from *niosdhcp.SharednetworkNetworks, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Ref = flex.FlattenStringPointerEmptyAsNull(from.Ref)
}
