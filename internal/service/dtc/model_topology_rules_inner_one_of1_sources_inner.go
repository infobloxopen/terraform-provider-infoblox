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

	niosdtc "github.com/infobloxopen/infoblox-nios-go-client/dtc"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// TopologyRulesInnerOneOf1SourcesInnerModel is the Terraform model for TopologyRulesInnerOneOf1SourcesInner
type TopologyRulesInnerOneOf1SourcesInnerModel struct {
	SourceOp    types.String `tfsdk:"source_op"`
	SourceType  types.String `tfsdk:"source_type"`
	SourceValue types.String `tfsdk:"source_value"`
}

// TopologyRulesInnerOneOf1SourcesInnerAttrTypes contains the attribute types for TopologyRulesInnerOneOf1SourcesInnerModel
var TopologyRulesInnerOneOf1SourcesInnerAttrTypes = map[string]attr.Type{
	"source_op":    types.StringType,
	"source_type":  types.StringType,
	"source_value": types.StringType,
}

// TopologyRulesInnerOneOf1SourcesInnerResourceSchemaAttributes contains the schema attributes for TopologyRulesInnerOneOf1SourcesInnerModel
var TopologyRulesInnerOneOf1SourcesInnerResourceSchemaAttributes = map[string]schema.Attribute{
	"source_op": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("IS", "IS_NOT"),
		},
		Optional:            true,
		MarkdownDescription: "Operation for matching the source.",
	},
	"source_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("CITY", "CONTINENT", "COUNTRY", "EA0", "EA1", "EA2", "EA3", "SUBDIVISION", "SUBNET"),
		},
		Required:            true,
		MarkdownDescription: "Type of the source.",
	},
	"source_value": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Value of the source.",
	},
}

// ExpandTopologyRulesInnerOneOf1SourcesInner converts a Terraform Object to SDK type
func ExpandTopologyRulesInnerOneOf1SourcesInner(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdtc.DtcTopologyRulesInnerOneOf1SourcesInner {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m TopologyRulesInnerOneOf1SourcesInnerModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *TopologyRulesInnerOneOf1SourcesInnerModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdtc.DtcTopologyRulesInnerOneOf1SourcesInner {
	if m == nil {
		return nil
	}
	to := &niosdtc.DtcTopologyRulesInnerOneOf1SourcesInner{
		SourceOp:    flex.ExpandStringPointerNullAsEmpty(m.SourceOp),
		SourceType:  flex.ExpandStringPointerNullAsEmpty(m.SourceType),
		SourceValue: flex.ExpandStringPointerNullAsEmpty(m.SourceValue),
	}
	return to
}

// FlattenTopologyRulesInnerOneOf1SourcesInner converts an SDK type to Terraform Object
func FlattenTopologyRulesInnerOneOf1SourcesInner(ctx context.Context, from *niosdtc.DtcTopologyRulesInnerOneOf1SourcesInner, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(TopologyRulesInnerOneOf1SourcesInnerAttrTypes)
	}
	m := &TopologyRulesInnerOneOf1SourcesInnerModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, TopologyRulesInnerOneOf1SourcesInnerAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *TopologyRulesInnerOneOf1SourcesInnerModel) Flatten(ctx context.Context, from *niosdtc.DtcTopologyRulesInnerOneOf1SourcesInner, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.SourceOp = flex.FlattenStringPointerEmptyAsNull(from.SourceOp)
	m.SourceType = flex.FlattenStringPointerEmptyAsNull(from.SourceType)
	m.SourceValue = flex.FlattenStringPointerEmptyAsNull(from.SourceValue)
}
