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

// DNSSECValidationBlockModel is the Terraform model for DNSSECValidationBlock
type DNSSECValidationBlockModel struct {
	DnssecEnableValidation types.Bool `tfsdk:"dnssec_enable_validation"`
	DnssecEnabled          types.Bool `tfsdk:"dnssec_enabled"`
	DnssecTrustAnchors     types.List `tfsdk:"dnssec_trust_anchors"`
	DnssecValidateExpiry   types.Bool `tfsdk:"dnssec_validate_expiry"`
}

// DNSSECValidationBlockAttrTypes contains the attribute types for DNSSECValidationBlockModel
var DNSSECValidationBlockAttrTypes = map[string]attr.Type{
	"dnssec_enable_validation": types.BoolType,
	"dnssec_enabled":           types.BoolType,
	"dnssec_trust_anchors":     types.ListType{ElemType: types.ObjectType{AttrTypes: TrustAnchorAttrTypes}},
	"dnssec_validate_expiry":   types.BoolType,
}

// DNSSECValidationBlockResourceSchemaAttributes contains the schema attributes for DNSSECValidationBlockModel
var DNSSECValidationBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"dnssec_enable_validation": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _dnssec_enable_validation_ field.",
	},
	"dnssec_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _dnssec_enabled_ field.",
	},
	"dnssec_trust_anchors": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: TrustAnchorResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _dnssec_trust_anchors_ field.",
	},
	"dnssec_validate_expiry": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _dnssec_validate_expiry_ field.",
	},
}

// ExpandDNSSECValidationBlock converts a Terraform Object to SDK type
func ExpandDNSSECValidationBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.DNSSECValidationBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DNSSECValidationBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *DNSSECValidationBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.DNSSECValidationBlock {
	if m == nil {
		return nil
	}
	to := &uddidns.DNSSECValidationBlock{
		DnssecEnableValidation: flex.ExpandBoolPointer(m.DnssecEnableValidation),
		DnssecEnabled:          flex.ExpandBoolPointer(m.DnssecEnabled),
		DnssecTrustAnchors:     flex.ExpandFrameworkListNestedBlock(ctx, m.DnssecTrustAnchors, diags, ExpandTrustAnchor),
		DnssecValidateExpiry:   flex.ExpandBoolPointer(m.DnssecValidateExpiry),
	}
	return to
}

// FlattenDNSSECValidationBlock converts an SDK type to Terraform Object
func FlattenDNSSECValidationBlock(ctx context.Context, from *uddidns.DNSSECValidationBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DNSSECValidationBlockAttrTypes)
	}
	m := &DNSSECValidationBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DNSSECValidationBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *DNSSECValidationBlockModel) Flatten(ctx context.Context, from *uddidns.DNSSECValidationBlock, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.DnssecEnableValidation = flex.FlattenBoolPointer(from.DnssecEnableValidation)
	m.DnssecEnabled = flex.FlattenBoolPointer(from.DnssecEnabled)
	m.DnssecTrustAnchors = flex.FlattenFrameworkListNestedBlock(ctx, from.DnssecTrustAnchors, TrustAnchorAttrTypes, diags, FlattenTrustAnchor)
	m.DnssecValidateExpiry = flex.FlattenBoolPointer(from.DnssecValidateExpiry)
}
