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

// RangeMsServerModel is the Terraform model for RangeMsServer
type RangeMsServerModel struct {
	Ipv4addr types.String `tfsdk:"ipv4addr"`
}

// RangeMsServerAttrTypes contains the attribute types for RangeMsServerModel
var RangeMsServerAttrTypes = map[string]attr.Type{
	"ipv4addr": types.StringType,
}

// RangeMsServerResourceSchemaAttributes contains the schema attributes for RangeMsServerModel
var RangeMsServerResourceSchemaAttributes = map[string]schema.Attribute{
	"ipv4addr": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidIPv4OrFQDN(),
		},
		MarkdownDescription: "The IPv4 Address or FQDN of the Microsoft server.",
	},
}

// ExpandRangeMsServer converts a Terraform Object to SDK type
func ExpandRangeMsServer(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.RangeMsServer {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RangeMsServerModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RangeMsServerModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.RangeMsServer {
	if m == nil {
		return nil
	}
	to := &niosdhcp.RangeMsServer{
		Ipv4addr: flex.ExpandStringPointerNullAsEmpty(m.Ipv4addr),
	}
	return to
}

// FlattenRangeMsServer converts an SDK type to Terraform Object
func FlattenRangeMsServer(ctx context.Context, from *niosdhcp.RangeMsServer, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RangeMsServerAttrTypes)
	}
	m := &RangeMsServerModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RangeMsServerAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RangeMsServerModel) Flatten(ctx context.Context, from *niosdhcp.RangeMsServer, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Ipv4addr = flex.FlattenStringPointerEmptyAsNull(from.Ipv4addr)
}
