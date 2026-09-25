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

// AccessFilterModel is the Terraform model for AccessFilter
type AccessFilterModel struct {
	Access           types.String `tfsdk:"access"`
	HardwareFilterId types.String `tfsdk:"hardware_filter_id"`
	OptionFilterId   types.String `tfsdk:"option_filter_id"`
}

// AccessFilterAttrTypes contains the attribute types for AccessFilterModel
var AccessFilterAttrTypes = map[string]attr.Type{
	"access":             types.StringType,
	"hardware_filter_id": types.StringType,
	"option_filter_id":   types.StringType,
}

// AccessFilterResourceSchemaAttributes contains the schema attributes for AccessFilterModel
var AccessFilterResourceSchemaAttributes = map[string]schema.Attribute{
	"access": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The access type of DHCP filter (_allow_ or _deny_).  Defaults to _allow_.",
	},
	"hardware_filter_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"option_filter_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
}

// ExpandAccessFilter converts a Terraform Object to SDK type
func ExpandAccessFilter(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidhcp.AccessFilter {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AccessFilterModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *AccessFilterModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidhcp.AccessFilter {
	if m == nil {
		return nil
	}
	to := &uddidhcp.AccessFilter{
		Access:           flex.ExpandString(m.Access),
		HardwareFilterId: flex.ExpandStringPointer(m.HardwareFilterId),
		OptionFilterId:   flex.ExpandStringPointer(m.OptionFilterId),
	}
	return to
}

// FlattenAccessFilter converts an SDK type to Terraform Object
func FlattenAccessFilter(ctx context.Context, from *uddidhcp.AccessFilter, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AccessFilterAttrTypes)
	}
	m := &AccessFilterModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AccessFilterAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *AccessFilterModel) Flatten(ctx context.Context, from *uddidhcp.AccessFilter, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Access = flex.FlattenString(from.Access)
	m.HardwareFilterId = flex.FlattenStringPointer(from.HardwareFilterId)
	m.OptionFilterId = flex.FlattenStringPointer(from.OptionFilterId)
}
