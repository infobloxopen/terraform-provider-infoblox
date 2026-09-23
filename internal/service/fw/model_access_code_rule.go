package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddifw "github.com/infobloxopen/universal-ddi-go-client/fw"
)

// AccessCodeRuleModel is the Terraform model for AccessCodeRule
type AccessCodeRuleModel struct {
	Action       types.String `tfsdk:"action"`
	Data         types.String `tfsdk:"data"`
	Description  types.String `tfsdk:"description"`
	RedirectName types.String `tfsdk:"redirect_name"`
	Type         types.String `tfsdk:"type"`
}

// AccessCodeRuleAttrTypes contains the attribute types for AccessCodeRuleModel
var AccessCodeRuleAttrTypes = map[string]attr.Type{
	"action":        types.StringType,
	"data":          types.StringType,
	"description":   types.StringType,
	"redirect_name": types.StringType,
	"type":          types.StringType,
}

// AccessCodeRuleResourceSchemaAttributes contains the schema attributes for AccessCodeRuleModel
var AccessCodeRuleResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"data": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "",
	},
	"description": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"redirect_name": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "",
	},
}

// ExpandAccessCodeRule converts a Terraform Object to SDK type
func ExpandAccessCodeRule(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddifw.AccessCodeRule {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AccessCodeRuleModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *AccessCodeRuleModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddifw.AccessCodeRule {
	if m == nil {
		return nil
	}
	to := &uddifw.AccessCodeRule{
		Action:       flex.ExpandStringPointer(m.Action),
		Data:         flex.ExpandStringPointer(m.Data),
		Description:  flex.ExpandStringPointer(m.Description),
		RedirectName: flex.ExpandStringPointer(m.RedirectName),
		Type:         flex.ExpandStringPointer(m.Type),
	}
	return to
}

// FlattenAccessCodeRule converts an SDK type to Terraform Object
func FlattenAccessCodeRule(ctx context.Context, from *uddifw.AccessCodeRule, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AccessCodeRuleAttrTypes)
	}
	m := &AccessCodeRuleModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AccessCodeRuleAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *AccessCodeRuleModel) Flatten(ctx context.Context, from *uddifw.AccessCodeRule, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.Data = flex.FlattenStringPointer(from.Data)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.RedirectName = flex.FlattenStringPointer(from.RedirectName)
	m.Type = flex.FlattenStringPointer(from.Type)
}
