package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddifw "github.com/infobloxopen/universal-ddi-go-client/fw"
)

// SecurityPolicyRuleModel is the Terraform model for SecurityPolicyRule
type SecurityPolicyRuleModel struct {
	Action       types.String `tfsdk:"action"`
	Data         types.String `tfsdk:"data"`
	ListId       types.Int32  `tfsdk:"list_id"`
	PolicyId     types.Int32  `tfsdk:"policy_id"`
	PolicyName   types.String `tfsdk:"policy_name"`
	RedirectName types.String `tfsdk:"redirect_name"`
	Type         types.String `tfsdk:"type"`
}

// SecurityPolicyRuleAttrTypes contains the attribute types for SecurityPolicyRuleModel
var SecurityPolicyRuleAttrTypes = map[string]attr.Type{
	"action":        types.StringType,
	"data":          types.StringType,
	"list_id":       types.Int32Type,
	"policy_id":     types.Int32Type,
	"policy_name":   types.StringType,
	"redirect_name": types.StringType,
	"type":          types.StringType,
}

// SecurityPolicyRuleResourceSchemaAttributes contains the schema attributes for SecurityPolicyRuleModel
var SecurityPolicyRuleResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The action for the policy rule that can be either \"action_allow\" or \"action_log\" or \"action_redirect\" or \"action_block\" or \"action_allow_with_local_resolution\". \"action_allow_with_local_resolution\" only supported for application filter rule with enabled onprem_resolve flag on the Security policy.",
	},
	"data": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The data source for the policy rule, that can be either a name of the predefined feed for \"named_feed\", custom list name for \"custom_list\" type, category filter name for \"category_filter\" type and application filter name for \"application_filter\" type.",
	},
	"list_id": schema.Int32Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The Custom List object identifier with which the policy rule is associated. 0 value means no custom list is associated with this policy rule.",
	},
	"policy_id": schema.Int32Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The identifier of the Security Policy object with which the policy rule is associated.",
	},
	"policy_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the security policy with which the policy rule is associated.",
	},
	"redirect_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the redirect address for redirect actions that can be either IPv4 address or a domain name.",
	},
	"type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The policy rule type that can be either \"named_feed\" or \"custom_list\" or \"category_filter\" or \"application_filter\".",
	},
}

// ExpandSecurityPolicyRule converts a Terraform Object to SDK type
func ExpandSecurityPolicyRule(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddifw.SecurityPolicyRule {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SecurityPolicyRuleModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *SecurityPolicyRuleModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddifw.SecurityPolicyRule {
	if m == nil {
		return nil
	}
	to := &uddifw.SecurityPolicyRule{
		Action:       flex.ExpandStringPointer(m.Action),
		Data:         flex.ExpandStringPointer(m.Data),
		ListId:       flex.ExpandInt32Pointer(m.ListId),
		PolicyId:     flex.ExpandInt32Pointer(m.PolicyId),
		PolicyName:   flex.ExpandStringPointer(m.PolicyName),
		RedirectName: flex.ExpandStringPointer(m.RedirectName),
		Type:         flex.ExpandStringPointer(m.Type),
	}
	return to
}

// FlattenSecurityPolicyRule converts an SDK type to Terraform Object
func FlattenSecurityPolicyRule(ctx context.Context, from *uddifw.SecurityPolicyRule, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SecurityPolicyRuleAttrTypes)
	}
	m := &SecurityPolicyRuleModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SecurityPolicyRuleAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *SecurityPolicyRuleModel) Flatten(ctx context.Context, from *uddifw.SecurityPolicyRule, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.Data = flex.FlattenStringPointer(from.Data)
	m.ListId = flex.FlattenInt32Pointer(from.ListId)
	m.PolicyId = flex.FlattenInt32Pointer(from.PolicyId)
	m.PolicyName = flex.FlattenStringPointer(from.PolicyName)
	m.RedirectName = flex.FlattenStringPointer(from.RedirectName)
	m.Type = flex.FlattenStringPointer(from.Type)
}
