package misc

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

	niosmisc "github.com/infobloxopen/infoblox-nios-go-client/misc"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// RulesetNxdomainRulesModel is the Terraform model for RulesetNxdomainRules
type RulesetNxdomainRulesModel struct {
	Action  types.String `tfsdk:"action"`
	Pattern types.String `tfsdk:"pattern"`
}

// RulesetNxdomainRulesAttrTypes contains the attribute types for RulesetNxdomainRulesModel
var RulesetNxdomainRulesAttrTypes = map[string]attr.Type{
	"action":  types.StringType,
	"pattern": types.StringType,
}

// RulesetNxdomainRulesResourceSchemaAttributes contains the schema attributes for RulesetNxdomainRulesModel
var RulesetNxdomainRulesResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Default: stringdefault.StaticString("PASS"),
		Validators: []validator.String{
			stringvalidator.OneOf("PASS", "REDIRECT", "MODIFY"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The action to perform when a domain name matches the pattern defined in this Ruleset.",
	},
	"pattern": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The pattern that is used to match the domain name.",
	},
}

// ExpandRulesetNxdomainRules converts a Terraform Object to SDK type
func ExpandRulesetNxdomainRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosmisc.RulesetNxdomainRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RulesetNxdomainRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RulesetNxdomainRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosmisc.RulesetNxdomainRules {
	if m == nil {
		return nil
	}
	to := &niosmisc.RulesetNxdomainRules{
		Action:  flex.ExpandStringPointerNullAsEmpty(m.Action),
		Pattern: flex.ExpandStringPointerNullAsEmpty(m.Pattern),
	}
	return to
}

// FlattenRulesetNxdomainRules converts an SDK type to Terraform Object
func FlattenRulesetNxdomainRules(ctx context.Context, from *niosmisc.RulesetNxdomainRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RulesetNxdomainRulesAttrTypes)
	}
	m := &RulesetNxdomainRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RulesetNxdomainRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RulesetNxdomainRulesModel) Flatten(ctx context.Context, from *niosmisc.RulesetNxdomainRules, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointerEmptyAsNull(from.Action)
	m.Pattern = flex.FlattenStringPointerEmptyAsNull(from.Pattern)
}
