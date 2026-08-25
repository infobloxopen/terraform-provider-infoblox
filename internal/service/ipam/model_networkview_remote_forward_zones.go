package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// NetworkviewRemoteForwardZonesModel is the Terraform model for NetworkviewRemoteForwardZones
type NetworkviewRemoteForwardZonesModel struct {
	Fqdn                types.String      `tfsdk:"fqdn"`
	ServerAddress       iptypes.IPAddress `tfsdk:"server_address"`
	GssTsigDnsPrincipal types.String      `tfsdk:"gss_tsig_dns_principal"`
	GssTsigDomain       types.String      `tfsdk:"gss_tsig_domain"`
	TsigKey             types.String      `tfsdk:"tsig_key"`
	TsigKeyAlg          types.String      `tfsdk:"tsig_key_alg"`
	TsigKeyName         types.String      `tfsdk:"tsig_key_name"`
	KeyType             types.String      `tfsdk:"key_type"`
}

// NetworkviewRemoteForwardZonesAttrTypes contains the attribute types for NetworkviewRemoteForwardZonesModel
var NetworkviewRemoteForwardZonesAttrTypes = map[string]attr.Type{
	"fqdn":                   types.StringType,
	"server_address":         iptypes.IPAddressType{},
	"gss_tsig_dns_principal": types.StringType,
	"gss_tsig_domain":        types.StringType,
	"tsig_key":               types.StringType,
	"tsig_key_alg":           types.StringType,
	"tsig_key_name":          types.StringType,
	"key_type":               types.StringType,
}

// NetworkviewRemoteForwardZonesResourceSchemaAttributes contains the schema attributes for NetworkviewRemoteForwardZonesModel
var NetworkviewRemoteForwardZonesResourceSchemaAttributes = map[string]schema.Attribute{
	"fqdn": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidNIOSDomainName(),
		},
		MarkdownDescription: "The FQDN of the remote server.",
	},
	"server_address": schema.StringAttribute{
		Required:   true,
		CustomType: iptypes.IPAddressType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The remote server IP address.",
	},
	"gss_tsig_dns_principal": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The principal name in which GSS-TSIG for dynamic updates is enabled.",
	},
	"gss_tsig_domain": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The domain in which GSS-TSIG for dynamic updates is enabled.",
	},
	"tsig_key": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The TSIG key value.",
	},
	"tsig_key_alg": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("HMAC-MD5", "HMAC-SHA256"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The TSIG key alorithm name.",
	},
	"tsig_key_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the TSIG key. The key name entered here must match the TSIG key name on the external name server.",
	},
	"key_type": schema.StringAttribute{
		Default: stringdefault.StaticString("NONE"),
		Validators: []validator.String{
			stringvalidator.OneOf("GSS-TSIG", "NONE", "TSIG"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The key type to be used.",
	},
}

// ExpandNetworkviewRemoteForwardZones converts a Terraform Object to SDK type
func ExpandNetworkviewRemoteForwardZones(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.NetworkviewRemoteForwardZones {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkviewRemoteForwardZonesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NetworkviewRemoteForwardZonesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.NetworkviewRemoteForwardZones {
	if m == nil {
		return nil
	}
	to := &niosipam.NetworkviewRemoteForwardZones{
		Fqdn:                flex.ExpandStringPointerNullAsEmpty(m.Fqdn),
		ServerAddress:       flex.ExpandIPAddress(m.ServerAddress),
		GssTsigDnsPrincipal: flex.ExpandStringPointerNullAsEmpty(m.GssTsigDnsPrincipal),
		GssTsigDomain:       flex.ExpandStringPointerNullAsEmpty(m.GssTsigDomain),
		TsigKey:             flex.ExpandStringPointerNullAsEmpty(m.TsigKey),
		TsigKeyAlg:          flex.ExpandStringPointer(m.TsigKeyAlg),
		TsigKeyName:         flex.ExpandStringPointerNullAsEmpty(m.TsigKeyName),
		KeyType:             flex.ExpandStringPointerNullAsEmpty(m.KeyType),
	}
	return to
}

// FlattenNetworkviewRemoteForwardZones converts an SDK type to Terraform Object
func FlattenNetworkviewRemoteForwardZones(ctx context.Context, from *niosipam.NetworkviewRemoteForwardZones, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkviewRemoteForwardZonesAttrTypes)
	}
	m := &NetworkviewRemoteForwardZonesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkviewRemoteForwardZonesAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NetworkviewRemoteForwardZonesModel) Flatten(ctx context.Context, from *niosipam.NetworkviewRemoteForwardZones, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Fqdn = flex.FlattenStringPointerEmptyAsNull(from.Fqdn)
	m.ServerAddress = flex.FlattenIPAddress(from.ServerAddress)
	m.GssTsigDnsPrincipal = flex.FlattenStringPointerEmptyAsNull(from.GssTsigDnsPrincipal)
	m.GssTsigDomain = flex.FlattenStringPointerEmptyAsNull(from.GssTsigDomain)
	m.TsigKey = flex.FlattenStringPointerEmptyAsNull(from.TsigKey)
	m.TsigKeyAlg = flex.FlattenStringPointerEmptyAsNull(from.TsigKeyAlg)
	m.TsigKeyName = flex.FlattenStringPointerEmptyAsNull(from.TsigKeyName)
	m.KeyType = flex.FlattenStringPointerEmptyAsNull(from.KeyType)
}
