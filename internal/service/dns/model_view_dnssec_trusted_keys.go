package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// ViewDnssecTrustedKeysModel is the Terraform model for ViewDnssecTrustedKeys
type ViewDnssecTrustedKeysModel struct {
	Fqdn               types.String `tfsdk:"fqdn"`
	Algorithm          types.String `tfsdk:"algorithm"`
	Key                types.String `tfsdk:"key"`
	SecureEntryPoint   types.Bool   `tfsdk:"secure_entry_point"`
	DnssecMustBeSecure types.Bool   `tfsdk:"dnssec_must_be_secure"`
}

// ViewDnssecTrustedKeysAttrTypes contains the attribute types for ViewDnssecTrustedKeysModel
var ViewDnssecTrustedKeysAttrTypes = map[string]attr.Type{
	"fqdn":                  types.StringType,
	"algorithm":             types.StringType,
	"key":                   types.StringType,
	"secure_entry_point":    types.BoolType,
	"dnssec_must_be_secure": types.BoolType,
}

// ViewDnssecTrustedKeysResourceSchemaAttributes contains the schema attributes for ViewDnssecTrustedKeysModel
var ViewDnssecTrustedKeysResourceSchemaAttributes = map[string]schema.Attribute{
	"fqdn": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The FQDN of the domain for which the member validates responses to recursive queries.",
	},
	"algorithm": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The DNSSEC algorithm used to generate the key.",
	},
	"key": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The DNSSEC key.",
	},
	"secure_entry_point": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The secure entry point flag, if set it means this is a KSK configuration.",
	},
	"dnssec_must_be_secure": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Responses must be DNSSEC secure for this hierarchy/domain.",
	},
}

// ExpandViewDnssecTrustedKeys converts a Terraform Object to SDK type
func ExpandViewDnssecTrustedKeys(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.ViewDnssecTrustedKeys {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ViewDnssecTrustedKeysModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ViewDnssecTrustedKeysModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.ViewDnssecTrustedKeys {
	if m == nil {
		return nil
	}
	to := &niosdns.ViewDnssecTrustedKeys{
		Fqdn:               flex.ExpandStringPointerNullAsEmpty(m.Fqdn),
		Algorithm:          flex.ExpandStringPointerNullAsEmpty(m.Algorithm),
		Key:                flex.ExpandStringPointerNullAsEmpty(m.Key),
		SecureEntryPoint:   flex.ExpandBoolPointer(m.SecureEntryPoint),
		DnssecMustBeSecure: flex.ExpandBoolPointer(m.DnssecMustBeSecure),
	}
	return to
}

// FlattenViewDnssecTrustedKeys converts an SDK type to Terraform Object
func FlattenViewDnssecTrustedKeys(ctx context.Context, from *niosdns.ViewDnssecTrustedKeys, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ViewDnssecTrustedKeysAttrTypes)
	}
	m := &ViewDnssecTrustedKeysModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ViewDnssecTrustedKeysAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ViewDnssecTrustedKeysModel) Flatten(ctx context.Context, from *niosdns.ViewDnssecTrustedKeys, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Fqdn = flex.FlattenStringPointerEmptyAsNull(from.Fqdn)
	m.Algorithm = flex.FlattenStringPointerEmptyAsNull(from.Algorithm)
	m.Key = flex.FlattenStringPointerEmptyAsNull(from.Key)
	m.SecureEntryPoint = flex.FlattenBoolPointer(from.SecureEntryPoint)
	m.DnssecMustBeSecure = flex.FlattenBoolPointer(from.DnssecMustBeSecure)
}
