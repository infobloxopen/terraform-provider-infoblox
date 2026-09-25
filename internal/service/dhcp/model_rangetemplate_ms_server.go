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

// RangetemplateMsServerModel is the Terraform model for RangetemplateMsServer
type RangetemplateMsServerModel struct {
	Ipv4addr types.String `tfsdk:"ipv4addr"`
}

// RangetemplateMsServerAttrTypes contains the attribute types for RangetemplateMsServerModel
var RangetemplateMsServerAttrTypes = map[string]attr.Type{
	"ipv4addr": types.StringType,
}

// RangetemplateMsServerResourceSchemaAttributes contains the schema attributes for RangetemplateMsServerModel
var RangetemplateMsServerResourceSchemaAttributes = map[string]schema.Attribute{
	"ipv4addr": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidIPv4OrFQDN(),
		},
		MarkdownDescription: "The IPv4 Address or FQDN of the Microsoft server.",
	},
}

// ExpandRangetemplateMsServer converts a Terraform Object to SDK type
func ExpandRangetemplateMsServer(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.RangetemplateMsServer {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RangetemplateMsServerModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RangetemplateMsServerModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.RangetemplateMsServer {
	if m == nil {
		return nil
	}
	to := &niosdhcp.RangetemplateMsServer{
		Ipv4addr: flex.ExpandStringPointerNullAsEmpty(m.Ipv4addr),
	}
	return to
}

// FlattenRangetemplateMsServer converts an SDK type to Terraform Object
func FlattenRangetemplateMsServer(ctx context.Context, from *niosdhcp.RangetemplateMsServer, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RangetemplateMsServerAttrTypes)
	}
	m := &RangetemplateMsServerModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RangetemplateMsServerAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RangetemplateMsServerModel) Flatten(ctx context.Context, from *niosdhcp.RangetemplateMsServer, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Ipv4addr = flex.FlattenStringPointerEmptyAsNull(from.Ipv4addr)
}
