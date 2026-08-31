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

// TopologyRulesInnerModel is the Terraform model for TopologyRulesInner
type TopologyRulesInnerModel struct {
	DestType        types.String `tfsdk:"dest_type"`
	DestinationLink types.String `tfsdk:"destination_link"`
	ReturnType      types.String `tfsdk:"return_type"`
	Topology        types.String `tfsdk:"topology"`
	Valid           types.Bool   `tfsdk:"valid"`
	Sources         types.List   `tfsdk:"sources"`
}

// TopologyRulesInnerAttrTypes contains the attribute types for TopologyRulesInnerModel
var TopologyRulesInnerAttrTypes = map[string]attr.Type{
	"dest_type":        types.StringType,
	"destination_link": types.StringType,
	"return_type":      types.StringType,
	"topology":         types.StringType,
	"valid":            types.BoolType,
	"sources":          types.ListType{ElemType: types.ObjectType{AttrTypes: TopologyRulesInnerOneOf1SourcesInnerAttrTypes}},
}

// TopologyRulesInnerResourceSchemaAttributes contains the schema attributes for TopologyRulesInnerModel
var TopologyRulesInnerResourceSchemaAttributes = map[string]schema.Attribute{
	"dest_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("POOL", "SERVER"),
		},
		Required:            true,
		MarkdownDescription: "The type of the destination for this rule.",
	},
	"destination_link": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The reference to the destination object.",
	},
	"return_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("NOERR", "NXDOMAIN", "REGULAR"),
		},
		Optional:            true,
		MarkdownDescription: "The type of the return value for this source.",
	},
	"topology": schema.StringAttribute{
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The topology for this rule.",
	},
	"valid": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Indicates whether the rule is valid.",
	},
	"sources": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: TopologyRulesInnerOneOf1SourcesInnerResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Conditions for matching sources.",
	},
}

// ExpandTopologyRulesInner converts a Terraform Object to SDK type
func ExpandTopologyRulesInner(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdtc.DtcTopologyRulesInner {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m TopologyRulesInnerModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *TopologyRulesInnerModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdtc.DtcTopologyRulesInner {
	if m == nil {
		return nil
	}
	to := &niosdtc.DtcTopologyRulesInner{
		DtcTopologyRulesInnerOneOf1: &niosdtc.DtcTopologyRulesInnerOneOf1{
			DestType:        flex.ExpandStringPointerNullAsEmpty(m.DestType),
			DestinationLink: flex.ExpandStringPointerNullAsEmpty(m.DestinationLink),
			ReturnType:      flex.ExpandStringPointerNullAsEmpty(m.ReturnType),
			Topology:        flex.ExpandStringPointerNullAsEmpty(m.Topology),
			Valid:           flex.ExpandBoolPointer(m.Valid),
			Sources:         flex.ExpandFrameworkListNestedBlock(ctx, m.Sources, diags, ExpandTopologyRulesInnerOneOf1SourcesInner),
		},
	}
	return to
}

// FlattenTopologyRulesInner converts an SDK type to Terraform Object
func FlattenTopologyRulesInner(ctx context.Context, from *niosdtc.DtcTopologyRulesInner, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(TopologyRulesInnerAttrTypes)
	}
	m := &TopologyRulesInnerModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, TopologyRulesInnerAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *TopologyRulesInnerModel) Flatten(ctx context.Context, src *niosdtc.DtcTopologyRulesInner, diags *diag.Diagnostics) {
	if src == nil || m == nil {
		return
	}
	// The modelled fields live on one oneOf arm. A response that used a different
	// arm leaves the model empty rather than dereferencing a nil arm.
	if src.DtcTopologyRulesInnerOneOf1 == nil {
		return
	}
	from := src.DtcTopologyRulesInnerOneOf1
	m.DestType = flex.FlattenStringPointerEmptyAsNull(from.DestType)
	m.DestinationLink = flex.FlattenStringPointerEmptyAsNull(from.DestinationLink)
	m.ReturnType = flex.FlattenStringPointerEmptyAsNull(from.ReturnType)
	m.Topology = flex.FlattenStringPointerEmptyAsNull(from.Topology)
	m.Valid = flex.FlattenBoolPointer(from.Valid)
	m.Sources = flex.FlattenFrameworkListNestedBlock(ctx, from.Sources, TopologyRulesInnerOneOf1SourcesInnerAttrTypes, diags, FlattenTopologyRulesInnerOneOf1SourcesInner)
}
