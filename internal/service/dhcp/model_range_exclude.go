package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// RangeExcludeModel is the Terraform model for RangeExclude
type RangeExcludeModel struct {
	StartAddress iptypes.IPv4Address `tfsdk:"start_address"`
	EndAddress   iptypes.IPv4Address `tfsdk:"end_address"`
	Comment      types.String        `tfsdk:"comment"`
}

// RangeExcludeAttrTypes contains the attribute types for RangeExcludeModel
var RangeExcludeAttrTypes = map[string]attr.Type{
	"start_address": iptypes.IPv4AddressType{},
	"end_address":   iptypes.IPv4AddressType{},
	"comment":       types.StringType,
}

// RangeExcludeResourceSchemaAttributes contains the schema attributes for RangeExcludeModel
var RangeExcludeResourceSchemaAttributes = map[string]schema.Attribute{
	"start_address": schema.StringAttribute{
		Required:   true,
		CustomType: iptypes.IPv4AddressType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv4 Address starting address of the exclusion range.",
	},
	"end_address": schema.StringAttribute{
		Required:   true,
		CustomType: iptypes.IPv4AddressType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv4 Address ending address of the exclusion range.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for the exclusion range; maximum 256 characters.",
	},
}

// ExpandRangeExclude converts a Terraform Object to SDK type
func ExpandRangeExclude(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.RangeExclude {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RangeExcludeModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RangeExcludeModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.RangeExclude {
	if m == nil {
		return nil
	}
	to := &niosdhcp.RangeExclude{
		StartAddress: flex.ExpandIPv4Address(m.StartAddress),
		EndAddress:   flex.ExpandIPv4Address(m.EndAddress),
		Comment:      flex.ExpandStringPointerNullAsEmpty(m.Comment),
	}
	return to
}

// FlattenRangeExclude converts an SDK type to Terraform Object
func FlattenRangeExclude(ctx context.Context, from *niosdhcp.RangeExclude, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RangeExcludeAttrTypes)
	}
	m := &RangeExcludeModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RangeExcludeAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RangeExcludeModel) Flatten(ctx context.Context, from *niosdhcp.RangeExclude, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.StartAddress = flex.FlattenIPv4Address(from.StartAddress)
	m.EndAddress = flex.FlattenIPv4Address(from.EndAddress)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
}
