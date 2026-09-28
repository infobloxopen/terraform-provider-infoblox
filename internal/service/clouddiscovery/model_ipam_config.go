package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddiclouddiscovery "github.com/infobloxopen/universal-ddi-go-client/clouddiscovery"
)

// IPAMConfigModel is the Terraform model for IPAMConfig
type IPAMConfigModel struct {
	DhcpServer            types.String `tfsdk:"dhcp_server"`
	DisableIpamProjection types.Bool   `tfsdk:"disable_ipam_projection"`
	IpSpace               types.String `tfsdk:"ip_space"`
}

// IPAMConfigAttrTypes contains the attribute types for IPAMConfigModel
var IPAMConfigAttrTypes = map[string]attr.Type{
	"dhcp_server":             types.StringType,
	"disable_ipam_projection": types.BoolType,
	"ip_space":                types.StringType,
}

// IPAMConfigResourceSchemaAttributes contains the schema attributes for IPAMConfigModel
var IPAMConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"dhcp_server": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Address of the DHCP Server.",
	},
	"disable_ipam_projection": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag controls the IPAM Sync/Reconciliation for the provider",
	},
	"ip_space": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "",
	},
}

// ExpandIPAMConfig converts a Terraform Object to SDK type
func ExpandIPAMConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiclouddiscovery.IPAMConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IPAMConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *IPAMConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiclouddiscovery.IPAMConfig {
	if m == nil {
		return nil
	}
	to := &uddiclouddiscovery.IPAMConfig{
		DhcpServer:            flex.ExpandStringPointer(m.DhcpServer),
		DisableIpamProjection: flex.ExpandBoolPointer(m.DisableIpamProjection),
		IpSpace:               flex.ExpandStringPointer(m.IpSpace),
	}
	return to
}

// FlattenIPAMConfig converts an SDK type to Terraform Object
func FlattenIPAMConfig(ctx context.Context, from *uddiclouddiscovery.IPAMConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IPAMConfigAttrTypes)
	}
	m := &IPAMConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IPAMConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *IPAMConfigModel) Flatten(ctx context.Context, from *uddiclouddiscovery.IPAMConfig, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.DhcpServer = flex.FlattenStringPointer(from.DhcpServer)
	m.DisableIpamProjection = flex.FlattenBoolPointer(from.DisableIpamProjection)
	m.IpSpace = flex.FlattenStringPointer(from.IpSpace)
}
