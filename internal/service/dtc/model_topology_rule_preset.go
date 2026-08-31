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
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
	uddidtc "github.com/infobloxopen/universal-ddi-go-client/dtc"
)

// TopologyRulePresetModel is the Terraform model for TopologyRulePreset
type TopologyRulePresetModel struct {
	Name    types.String `tfsdk:"name"`
	Source  types.String `tfsdk:"source"`
	Subnets types.List   `tfsdk:"subnets"`
	Tags    types.List   `tfsdk:"tags"`
}

// TopologyRulePresetAttrTypes contains the attribute types for TopologyRulePresetModel
var TopologyRulePresetAttrTypes = map[string]attr.Type{
	"name":    types.StringType,
	"source":  types.StringType,
	"subnets": types.ListType{ElemType: types.StringType},
	"tags":    types.ListType{ElemType: types.ObjectType{AttrTypes: TagRuleAttrTypes}},
}

// TopologyRulePresetResourceSchemaAttributes contains the schema attributes for TopologyRulePresetModel
var TopologyRulePresetResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Required. Display name of __TopologyRulePreset__. Must be unique within __Topology__.",
	},
	"source": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("subnet", "tags", "default"),
		},
		Optional:            true,
		MarkdownDescription: "Type of source.  Allowed values: - subnet - tags - default  Defaults to _default_.",
	},
	"subnets": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. List of subnets in CIDR format.  Must be set if _source_ is _subnet_, otherwise must be empty.",
	},
	"tags": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: TagRuleResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. List of tag rules to match against a source object's effective tags. Effective tags = direct tags plus tags inherited from the IPAM parent chain (IPSpace → Address Block → Subnet); the closer level wins on key conflicts. All rules use AND semantics: an object must satisfy every __TagRule__ to match.  Must be set if _source_ is _tags_, otherwise must be empty.",
	},
}

// ExpandTopologyRulePreset converts a Terraform Object to SDK type
func ExpandTopologyRulePreset(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidtc.TopologyRulePreset {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m TopologyRulePresetModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *TopologyRulePresetModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidtc.TopologyRulePreset {
	if m == nil {
		return nil
	}
	to := &uddidtc.TopologyRulePreset{
		Name:    flex.ExpandString(m.Name),
		Source:  flex.ExpandStringPointer(m.Source),
		Subnets: flex.ExpandFrameworkListString(ctx, m.Subnets, diags),
		Tags:    flex.ExpandFrameworkListNestedBlock(ctx, m.Tags, diags, ExpandTagRule),
	}
	return to
}

// FlattenTopologyRulePreset converts an SDK type to Terraform Object
func FlattenTopologyRulePreset(ctx context.Context, from *uddidtc.TopologyRulePreset, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(TopologyRulePresetAttrTypes)
	}
	m := &TopologyRulePresetModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, TopologyRulePresetAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *TopologyRulePresetModel) Flatten(ctx context.Context, from *uddidtc.TopologyRulePreset, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Name = flex.FlattenString(from.Name)
	m.Source = flex.FlattenStringPointer(from.Source)
	m.Subnets = flex.FlattenFrameworkListString(ctx, from.Subnets, diags)
	m.Tags = flex.FlattenFrameworkListNestedBlock(ctx, from.Tags, TagRuleAttrTypes, diags, FlattenTagRule)
}
