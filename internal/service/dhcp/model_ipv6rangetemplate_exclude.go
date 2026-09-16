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

// Ipv6rangetemplateExcludeModel is the Terraform model for Ipv6rangetemplateExclude
type Ipv6rangetemplateExcludeModel struct {
	Offset            types.Int64  `tfsdk:"offset"`
	NumberOfAddresses types.Int64  `tfsdk:"number_of_addresses"`
	Comment           types.String `tfsdk:"comment"`
}

// Ipv6rangetemplateExcludeAttrTypes contains the attribute types for Ipv6rangetemplateExcludeModel
var Ipv6rangetemplateExcludeAttrTypes = map[string]attr.Type{
	"offset":              types.Int64Type,
	"number_of_addresses": types.Int64Type,
	"comment":             types.StringType,
}

// Ipv6rangetemplateExcludeResourceSchemaAttributes contains the schema attributes for Ipv6rangetemplateExcludeModel
var Ipv6rangetemplateExcludeResourceSchemaAttributes = map[string]schema.Attribute{
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
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "A descriptive comment of a DHCP exclusion range template.",
	},
}

// ExpandIpv6rangetemplateExclude converts a Terraform Object to SDK type
func ExpandIpv6rangetemplateExclude(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.Ipv6rangetemplateExclude {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6rangetemplateExcludeModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Ipv6rangetemplateExcludeModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.Ipv6rangetemplateExclude {
	if m == nil {
		return nil
	}
	to := &niosdhcp.Ipv6rangetemplateExclude{
		Offset:            flex.ExpandInt64Pointer(m.Offset),
		NumberOfAddresses: flex.ExpandInt64Pointer(m.NumberOfAddresses),
		Comment:           flex.ExpandStringPointerNullAsEmpty(m.Comment),
	}
	return to
}

// FlattenIpv6rangetemplateExclude converts an SDK type to Terraform Object
func FlattenIpv6rangetemplateExclude(ctx context.Context, from *niosdhcp.Ipv6rangetemplateExclude, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6rangetemplateExcludeAttrTypes)
	}
	m := &Ipv6rangetemplateExcludeModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6rangetemplateExcludeAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Ipv6rangetemplateExcludeModel) Flatten(ctx context.Context, from *niosdhcp.Ipv6rangetemplateExclude, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Offset = flex.FlattenInt64Pointer(from.Offset)
	m.NumberOfAddresses = flex.FlattenInt64Pointer(from.NumberOfAddresses)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
}
