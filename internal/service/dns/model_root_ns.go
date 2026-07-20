package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddidns "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// RootNSModel is the Terraform model for RootNS
type RootNSModel struct {
	Address      types.String `tfsdk:"address"`
	Fqdn         types.String `tfsdk:"fqdn"`
	ProtocolFqdn types.String `tfsdk:"protocol_fqdn"`
}

// RootNSAttrTypes contains the attribute types for RootNSModel
var RootNSAttrTypes = map[string]attr.Type{
	"address":       types.StringType,
	"fqdn":          types.StringType,
	"protocol_fqdn": types.StringType,
}

// RootNSResourceSchemaAttributes contains the schema attributes for RootNSModel
var RootNSResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "IPv4 address.",
	},
	"fqdn": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "FQDN.",
	},
	"protocol_fqdn": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "FQDN in punycode.",
	},
}

// ExpandRootNS converts a Terraform Object to SDK type
func ExpandRootNS(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.RootNS {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RootNSModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RootNSModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.RootNS {
	if m == nil {
		return nil
	}
	to := &uddidns.RootNS{
		Address:      flex.ExpandString(m.Address),
		Fqdn:         flex.ExpandString(m.Fqdn),
		ProtocolFqdn: flex.ExpandStringPointer(m.ProtocolFqdn),
	}
	return to
}

// FlattenRootNS converts an SDK type to Terraform Object
func FlattenRootNS(ctx context.Context, from *uddidns.RootNS, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RootNSAttrTypes)
	}
	m := &RootNSModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RootNSAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RootNSModel) Flatten(ctx context.Context, from *uddidns.RootNS, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenString(from.Address)
	m.Fqdn = flex.FlattenString(from.Fqdn)
	m.ProtocolFqdn = flex.FlattenStringPointer(from.ProtocolFqdn)
}
