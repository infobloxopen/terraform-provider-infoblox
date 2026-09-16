package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// ZoneDelegatedDelegateToModel is the Terraform model for ZoneDelegatedDelegateTo
type ZoneDelegatedDelegateToModel struct {
	Address                      iptypes.IPAddress `tfsdk:"address"`
	Name                         types.String      `tfsdk:"name"`
	SharedWithMsParentDelegation types.Bool        `tfsdk:"shared_with_ms_parent_delegation"`
	Stealth                      types.Bool        `tfsdk:"stealth"`
	TsigKey                      types.String      `tfsdk:"tsig_key"`
	TsigKeyAlg                   types.String      `tfsdk:"tsig_key_alg"`
	TsigKeyName                  types.String      `tfsdk:"tsig_key_name"`
	UseTsigKeyName               types.Bool        `tfsdk:"use_tsig_key_name"`
}

// ZoneDelegatedDelegateToAttrTypes contains the attribute types for ZoneDelegatedDelegateToModel
var ZoneDelegatedDelegateToAttrTypes = map[string]attr.Type{
	"address":                          iptypes.IPAddressType{},
	"name":                             types.StringType,
	"shared_with_ms_parent_delegation": types.BoolType,
	"stealth":                          types.BoolType,
	"tsig_key":                         types.StringType,
	"tsig_key_alg":                     types.StringType,
	"tsig_key_name":                    types.StringType,
	"use_tsig_key_name":                types.BoolType,
}

// ZoneDelegatedDelegateToResourceSchemaAttributes contains the schema attributes for ZoneDelegatedDelegateToModel
var ZoneDelegatedDelegateToResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required:   true,
		CustomType: iptypes.IPAddressType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv4 Address or IPv6 Address of the server.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "A resolvable domain name for the external DNS server.",
	},
	"shared_with_ms_parent_delegation": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "This flag represents whether the name server is shared with the parent Microsoft primary zone's delegation server.",
	},
	"stealth": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Set this flag to hide the NS record for the primary name server from DNS queries.",
	},
	"tsig_key": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "A generated TSIG key.",
	},
	"tsig_key_alg": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("HMAC-MD5", "HMAC-SHA256"),
		},
		Optional:            true,
		MarkdownDescription: "The TSIG key algorithm.",
	},
	"tsig_key_name": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The TSIG key name.",
	},
	"use_tsig_key_name": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: tsig_key_name",
	},
}

// ExpandZoneDelegatedDelegateTo converts a Terraform Object to SDK type
func ExpandZoneDelegatedDelegateTo(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.ZoneDelegatedDelegateTo {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneDelegatedDelegateToModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ZoneDelegatedDelegateToModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.ZoneDelegatedDelegateTo {
	if m == nil {
		return nil
	}
	to := &niosdns.ZoneDelegatedDelegateTo{
		Address:                      flex.ExpandIPAddress(m.Address),
		Name:                         flex.ExpandStringPointerNullAsEmpty(m.Name),
		SharedWithMsParentDelegation: flex.ExpandBoolPointer(m.SharedWithMsParentDelegation),
		Stealth:                      flex.ExpandBoolPointer(m.Stealth),
		TsigKey:                      flex.ExpandStringPointerNullAsEmpty(m.TsigKey),
		TsigKeyAlg:                   flex.ExpandStringPointer(m.TsigKeyAlg),
		TsigKeyName:                  flex.ExpandStringPointer(m.TsigKeyName),
		UseTsigKeyName:               flex.ExpandBoolPointer(m.UseTsigKeyName),
	}
	return to
}

// FlattenZoneDelegatedDelegateTo converts an SDK type to Terraform Object
func FlattenZoneDelegatedDelegateTo(ctx context.Context, from *niosdns.ZoneDelegatedDelegateTo, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneDelegatedDelegateToAttrTypes)
	}
	m := &ZoneDelegatedDelegateToModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneDelegatedDelegateToAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ZoneDelegatedDelegateToModel) Flatten(ctx context.Context, from *niosdns.ZoneDelegatedDelegateTo, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenIPAddress(from.Address)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.SharedWithMsParentDelegation = flex.FlattenBoolPointer(from.SharedWithMsParentDelegation)
	m.Stealth = flex.FlattenBoolPointer(from.Stealth)
	m.TsigKey = flex.FlattenStringPointerEmptyAsNull(from.TsigKey)
	m.TsigKeyAlg = flex.FlattenStringPointerEmptyAsNull(from.TsigKeyAlg)
	m.TsigKeyName = flex.FlattenStringPointerEmptyAsNull(from.TsigKeyName)
	m.UseTsigKeyName = flex.FlattenBoolPointer(from.UseTsigKeyName)
}
