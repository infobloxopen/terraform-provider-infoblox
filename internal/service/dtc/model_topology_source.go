package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
	uddidtc "github.com/infobloxopen/universal-ddi-go-client/dtc"
)

// TopologySourceModel is the Terraform model for TopologySource
type TopologySourceModel struct {
	Name     types.String `tfsdk:"name"`
	Source   types.String `tfsdk:"source"`
	Subnets  types.List   `tfsdk:"subnets"`
	TagRules types.List   `tfsdk:"tag_rules"`
}

// TopologySourceAttrTypes contains the attribute types for TopologySourceModel
var TopologySourceAttrTypes = map[string]attr.Type{
	"name":      types.StringType,
	"source":    types.StringType,
	"subnets":   types.ListType{ElemType: types.StringType},
	"tag_rules": types.ListType{ElemType: types.ObjectType{AttrTypes: TagRuleAttrTypes}},
}

// TopologySourceResourceSchemaAttributes contains the schema attributes for TopologySourceModel
var TopologySourceResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Required. Display name of __TopologySource__.",
	},
	"source": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Type of source.  Allowed values: - subnet - tag_rule  Required.",
	},
	"subnets": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. List of subnets in CIDR format.  Must be set if _source_ is set to _subnet_, otherwise must be empty.",
	},
	"tag_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: TagRuleResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. List of tag rules to match against infrastructure source objects effective tags.  Must be set if _source_ is set to _tag_rule_, otherwise must be empty.",
	},
}

// ExpandTopologySource converts a Terraform Object to SDK type
func ExpandTopologySource(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidtc.TopologySource {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m TopologySourceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *TopologySourceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidtc.TopologySource {
	if m == nil {
		return nil
	}
	to := &uddidtc.TopologySource{
		Name:     flex.ExpandString(m.Name),
		Source:   flex.ExpandString(m.Source),
		Subnets:  flex.ExpandFrameworkListString(ctx, m.Subnets, diags),
		TagRules: flex.ExpandFrameworkListNestedBlock(ctx, m.TagRules, diags, ExpandTagRule),
	}
	return to
}

// FlattenTopologySource converts an SDK type to Terraform Object
func FlattenTopologySource(ctx context.Context, from *uddidtc.TopologySource, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(TopologySourceAttrTypes)
	}
	m := &TopologySourceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, TopologySourceAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *TopologySourceModel) Flatten(ctx context.Context, from *uddidtc.TopologySource, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Name = flex.FlattenString(from.Name)
	m.Source = flex.FlattenString(from.Source)
	m.Subnets = flex.FlattenFrameworkListString(ctx, from.Subnets, diags)
	m.TagRules = flex.FlattenFrameworkListNestedBlock(ctx, from.TagRules, TagRuleAttrTypes, diags, FlattenTagRule)
}
