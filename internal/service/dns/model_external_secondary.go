package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
	uddidns "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// ExternalSecondaryModel is the Terraform model for ExternalSecondary
type ExternalSecondaryModel struct {
	Address      types.String `tfsdk:"address"`
	Fqdn         types.String `tfsdk:"fqdn"`
	ProtocolFqdn types.String `tfsdk:"protocol_fqdn"`
	Stealth      types.Bool   `tfsdk:"stealth"`
	TsigEnabled  types.Bool   `tfsdk:"tsig_enabled"`
	TsigKey      types.Object `tfsdk:"tsig_key"`
}

// ExternalSecondaryAttrTypes contains the attribute types for ExternalSecondaryModel
var ExternalSecondaryAttrTypes = map[string]attr.Type{
	"address":       types.StringType,
	"fqdn":          types.StringType,
	"protocol_fqdn": types.StringType,
	"stealth":       types.BoolType,
	"tsig_enabled":  types.BoolType,
	"tsig_key":      types.ObjectType{AttrTypes: TSIGKeyAttrTypes},
}

// ExternalSecondaryResourceSchemaAttributes contains the schema attributes for ExternalSecondaryModel
var ExternalSecondaryResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "IP Address of nameserver.",
	},
	"fqdn": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.IsValidUDDIDomainName(),
		},
		MarkdownDescription: "FQDN of nameserver.",
	},
	"protocol_fqdn": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "FQDN of nameserver in punycode.",
	},
	"stealth": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If enabled, the NS record and glue record will NOT be automatically generated according to secondaries nameserver assignment.  Default: _false_",
	},
	"tsig_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If enabled, secondaries will use the configured TSIG key when requesting a zone transfer.  Default: _false_",
	},
	"tsig_key": schema.SingleNestedAttribute{
		Attributes:          TSIGKeyResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "TSIG key.  Error if empty while _tsig_enabled_ is _true_.",
	},
}

// ExpandExternalSecondary converts a Terraform Object to SDK type
func ExpandExternalSecondary(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.ExternalSecondary {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ExternalSecondaryModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ExternalSecondaryModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.ExternalSecondary {
	if m == nil {
		return nil
	}
	to := &uddidns.ExternalSecondary{
		Address:     flex.ExpandString(m.Address),
		Fqdn:        flex.ExpandString(m.Fqdn),
		Stealth:     flex.ExpandBoolPointer(m.Stealth),
		TsigEnabled: flex.ExpandBoolPointer(m.TsigEnabled),
		TsigKey:     ExpandTSIGKey(ctx, m.TsigKey, diags),
	}
	return to
}

// FlattenExternalSecondary converts an SDK type to Terraform Object
func FlattenExternalSecondary(ctx context.Context, from *uddidns.ExternalSecondary, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ExternalSecondaryAttrTypes)
	}
	m := &ExternalSecondaryModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ExternalSecondaryAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ExternalSecondaryModel) Flatten(ctx context.Context, from *uddidns.ExternalSecondary, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenString(from.Address)
	m.Fqdn = flex.FlattenString(from.Fqdn)
	m.ProtocolFqdn = flex.FlattenStringPointer(from.ProtocolFqdn)
	m.Stealth = flex.FlattenBoolPointer(from.Stealth)
	m.TsigEnabled = flex.FlattenBoolPointer(from.TsigEnabled)
	m.TsigKey = FlattenTSIGKey(ctx, from.TsigKey, diags)
}
