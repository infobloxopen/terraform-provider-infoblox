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

// FixedaddressMsServerModel is the Terraform model for FixedaddressMsServer
type FixedaddressMsServerModel struct {
	Ipv4addr types.String `tfsdk:"ipv4addr"`
}

// FixedaddressMsServerAttrTypes contains the attribute types for FixedaddressMsServerModel
var FixedaddressMsServerAttrTypes = map[string]attr.Type{
	"ipv4addr": types.StringType,
}

// FixedaddressMsServerResourceSchemaAttributes contains the schema attributes for FixedaddressMsServerModel
var FixedaddressMsServerResourceSchemaAttributes = map[string]schema.Attribute{
	"ipv4addr": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv4 Address or FQDN of the Microsoft server.",
	},
}

// ExpandFixedaddressMsServer converts a Terraform Object to SDK type
func ExpandFixedaddressMsServer(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.FixedaddressMsServer {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m FixedaddressMsServerModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *FixedaddressMsServerModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.FixedaddressMsServer {
	if m == nil {
		return nil
	}
	to := &niosdhcp.FixedaddressMsServer{
		Ipv4addr: flex.ExpandStringPointerNullAsEmpty(m.Ipv4addr),
	}
	return to
}

// FlattenFixedaddressMsServer converts an SDK type to Terraform Object
func FlattenFixedaddressMsServer(ctx context.Context, from *niosdhcp.FixedaddressMsServer, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FixedaddressMsServerAttrTypes)
	}
	m := &FixedaddressMsServerModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, FixedaddressMsServerAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *FixedaddressMsServerModel) Flatten(ctx context.Context, from *niosdhcp.FixedaddressMsServer, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Ipv4addr = flex.FlattenStringPointerEmptyAsNull(from.Ipv4addr)
}
