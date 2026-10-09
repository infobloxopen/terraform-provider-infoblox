package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	planmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
	uddidns "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// ExternalPrimaryModel is the Terraform model for ExternalPrimary
type ExternalPrimaryModel struct {
	Address      types.String `tfsdk:"address"`
	Fqdn         types.String `tfsdk:"fqdn"`
	Nsg          types.String `tfsdk:"nsg"`
	ProtocolFqdn types.String `tfsdk:"protocol_fqdn"`
	TsigEnabled  types.Bool   `tfsdk:"tsig_enabled"`
	TsigKey      types.Object `tfsdk:"tsig_key"`
	Type         types.String `tfsdk:"type"`
}

// ExternalPrimaryAttrTypes contains the attribute types for ExternalPrimaryModel
var ExternalPrimaryAttrTypes = map[string]attr.Type{
	"address":       types.StringType,
	"fqdn":          types.StringType,
	"nsg":           types.StringType,
	"protocol_fqdn": types.StringType,
	"tsig_enabled":  types.BoolType,
	"tsig_key":      types.ObjectType{AttrTypes: TSIGKeyAttrTypes},
	"type":          types.StringType,
}

// ExternalPrimaryResourceSchemaAttributes contains the schema attributes for ExternalPrimaryModel
var ExternalPrimaryResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			planmod.UseEmptyStringForNull(),
		},
		MarkdownDescription: "Optional. Required only if _type_ is _server_. IP Address of nameserver.",
	},
	"fqdn": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			planmod.UseEmptyStringForNull(),
		},
		Validators: []validator.String{
			customvalidator.IsValidUDDIDomainName(),
		},
		MarkdownDescription: "Optional. Required only if _type_ is _server_. FQDN of nameserver.",
	},
	"nsg": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"protocol_fqdn": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "FQDN of nameserver in punycode.",
	},
	"tsig_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. If enabled, secondaries will use the configured TSIG key when requesting a zone transfer from this primary.",
	},
	"tsig_key": schema.SingleNestedAttribute{
		Attributes:          TSIGKeyResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Optional. TSIG key.  Error if empty while _tsig_enabled_ is _true_.",
	},
	"type": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Allowed values: * _nsg_, * _primary_.",
	},
}

// ExpandExternalPrimary converts a Terraform Object to SDK type
func ExpandExternalPrimary(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.ExternalPrimary {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ExternalPrimaryModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ExternalPrimaryModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.ExternalPrimary {
	if m == nil {
		return nil
	}
	to := &uddidns.ExternalPrimary{
		Address:     flex.ExpandStringPointer(m.Address),
		Fqdn:        flex.ExpandStringPointer(m.Fqdn),
		Nsg:         flex.ExpandStringPointer(m.Nsg),
		TsigEnabled: flex.ExpandBoolPointer(m.TsigEnabled),
		TsigKey:     ExpandTSIGKey(ctx, m.TsigKey, diags),
		Type:        flex.ExpandString(m.Type),
	}
	return to
}

// FlattenExternalPrimary converts an SDK type to Terraform Object
func FlattenExternalPrimary(ctx context.Context, from *uddidns.ExternalPrimary, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ExternalPrimaryAttrTypes)
	}
	m := &ExternalPrimaryModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ExternalPrimaryAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ExternalPrimaryModel) Flatten(ctx context.Context, from *uddidns.ExternalPrimary, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Fqdn = flex.FlattenStringPointer(from.Fqdn)
	m.Nsg = flex.FlattenStringPointer(from.Nsg)
	m.ProtocolFqdn = flex.FlattenStringPointer(from.ProtocolFqdn)
	m.TsigEnabled = flex.FlattenBoolPointer(from.TsigEnabled)
	m.TsigKey = FlattenTSIGKey(ctx, from.TsigKey, diags)
	m.Type = flex.FlattenString(from.Type)
}
