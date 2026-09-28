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

// RangetemplateExcludeModel is the Terraform model for RangetemplateExclude
type RangetemplateExcludeModel struct {
	Offset            types.Int64  `tfsdk:"offset"`
	NumberOfAddresses types.Int64  `tfsdk:"number_of_addresses"`
	Comment           types.String `tfsdk:"comment"`
}

// RangetemplateExcludeAttrTypes contains the attribute types for RangetemplateExcludeModel
var RangetemplateExcludeAttrTypes = map[string]attr.Type{
	"offset":              types.Int64Type,
	"number_of_addresses": types.Int64Type,
	"comment":             types.StringType,
}

// RangetemplateExcludeResourceSchemaAttributes contains the schema attributes for RangetemplateExcludeModel
var RangetemplateExcludeResourceSchemaAttributes = map[string]schema.Attribute{
	"offset": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "The address offset of the DHCP exclusion range template.",
	},
	"number_of_addresses": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "The number of addresses in the DHCP exclusion range template.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "A descriptive comment of a DHCP exclusion range template.",
	},
}

// ExpandRangetemplateExclude converts a Terraform Object to SDK type
func ExpandRangetemplateExclude(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.RangetemplateExclude {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RangetemplateExcludeModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RangetemplateExcludeModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.RangetemplateExclude {
	if m == nil {
		return nil
	}
	to := &niosdhcp.RangetemplateExclude{
		Offset:            flex.ExpandInt64Pointer(m.Offset),
		NumberOfAddresses: flex.ExpandInt64Pointer(m.NumberOfAddresses),
		Comment:           flex.ExpandStringPointerNullAsEmpty(m.Comment),
	}
	return to
}

// FlattenRangetemplateExclude converts an SDK type to Terraform Object
func FlattenRangetemplateExclude(ctx context.Context, from *niosdhcp.RangetemplateExclude, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RangetemplateExcludeAttrTypes)
	}
	m := &RangetemplateExcludeModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RangetemplateExcludeAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RangetemplateExcludeModel) Flatten(ctx context.Context, from *niosdhcp.RangetemplateExclude, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Offset = flex.FlattenInt64Pointer(from.Offset)
	m.NumberOfAddresses = flex.FlattenInt64Pointer(from.NumberOfAddresses)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
}
