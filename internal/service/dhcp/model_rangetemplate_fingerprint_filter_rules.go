package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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

// RangetemplateFingerprintFilterRulesModel is the Terraform model for RangetemplateFingerprintFilterRules
type RangetemplateFingerprintFilterRulesModel struct {
	Filter     types.String `tfsdk:"filter"`
	Permission types.String `tfsdk:"permission"`
}

// RangetemplateFingerprintFilterRulesAttrTypes contains the attribute types for RangetemplateFingerprintFilterRulesModel
var RangetemplateFingerprintFilterRulesAttrTypes = map[string]attr.Type{
	"filter":     types.StringType,
	"permission": types.StringType,
}

// RangetemplateFingerprintFilterRulesResourceSchemaAttributes contains the schema attributes for RangetemplateFingerprintFilterRulesModel
var RangetemplateFingerprintFilterRulesResourceSchemaAttributes = map[string]schema.Attribute{
	"filter": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the DHCP filter.",
	},
	"permission": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("Allow", "Deny"),
		},
		Required:            true,
		MarkdownDescription: "The permission to be applied.",
	},
}

// ExpandRangetemplateFingerprintFilterRules converts a Terraform Object to SDK type
func ExpandRangetemplateFingerprintFilterRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.RangetemplateFingerprintFilterRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RangetemplateFingerprintFilterRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RangetemplateFingerprintFilterRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.RangetemplateFingerprintFilterRules {
	if m == nil {
		return nil
	}
	to := &niosdhcp.RangetemplateFingerprintFilterRules{
		Filter:     flex.ExpandStringPointerNullAsEmpty(m.Filter),
		Permission: flex.ExpandStringPointerNullAsEmpty(m.Permission),
	}
	return to
}

// FlattenRangetemplateFingerprintFilterRules converts an SDK type to Terraform Object
func FlattenRangetemplateFingerprintFilterRules(ctx context.Context, from *niosdhcp.RangetemplateFingerprintFilterRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RangetemplateFingerprintFilterRulesAttrTypes)
	}
	m := &RangetemplateFingerprintFilterRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RangetemplateFingerprintFilterRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RangetemplateFingerprintFilterRulesModel) Flatten(ctx context.Context, from *niosdhcp.RangetemplateFingerprintFilterRules, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Filter = flex.FlattenStringPointerEmptyAsNull(from.Filter)
	m.Permission = flex.FlattenStringPointerEmptyAsNull(from.Permission)
}
