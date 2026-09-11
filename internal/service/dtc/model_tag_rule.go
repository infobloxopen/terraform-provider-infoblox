package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddidtc "github.com/infobloxopen/universal-ddi-go-client/dtc"
)

// TagRuleModel is the Terraform model for TagRule
type TagRuleModel struct {
	Key   types.String `tfsdk:"key"`
	Op    types.String `tfsdk:"op"`
	Value types.String `tfsdk:"value"`
}

// TagRuleAttrTypes contains the attribute types for TagRuleModel
var TagRuleAttrTypes = map[string]attr.Type{
	"key":   types.StringType,
	"op":    types.StringType,
	"value": types.StringType,
}

// TagRuleResourceSchemaAttributes contains the schema attributes for TagRuleModel
var TagRuleResourceSchemaAttributes = map[string]schema.Attribute{
	"key": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Required. Tag key to match against a source object's effective tags.",
	},
	"op": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("EQUALS", "NOT_EQUALS"),
		},
		Optional:            true,
		MarkdownDescription: "Optional. Match operator.  Supported values: - EQUALS: matches when the key exists and its value equals the configured value. - NOT_EQUALS: matches when the key exists and all values for that key differ   from the configured value.  A missing key does not satisfy either operator.  Defaults to _EQUALS_.",
	},
	"value": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Required. Tag value to match against a source object's effective tags.",
	},
}

// ExpandTagRule converts a Terraform Object to SDK type
func ExpandTagRule(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidtc.TagRule {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m TagRuleModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *TagRuleModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidtc.TagRule {
	if m == nil {
		return nil
	}
	to := &uddidtc.TagRule{
		Key:   flex.ExpandString(m.Key),
		Op:    flex.ExpandStringPointer(m.Op),
		Value: flex.ExpandString(m.Value),
	}
	return to
}

// FlattenTagRule converts an SDK type to Terraform Object
func FlattenTagRule(ctx context.Context, from *uddidtc.TagRule, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(TagRuleAttrTypes)
	}
	m := &TagRuleModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, TagRuleAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *TagRuleModel) Flatten(ctx context.Context, from *uddidtc.TagRule, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Key = flex.FlattenString(from.Key)
	m.Op = flex.FlattenStringPointer(from.Op)
	m.Value = flex.FlattenString(from.Value)
}
