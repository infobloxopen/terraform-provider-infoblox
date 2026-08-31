package misc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/misc"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type RulesetModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var RulesetAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSRulesetAttrTypes},
}

type NIOSRulesetModel struct {
	Comment       types.String `tfsdk:"comment"`
	Disabled      types.Bool   `tfsdk:"disabled"`
	Name          types.String `tfsdk:"name"`
	NxdomainRules types.List   `tfsdk:"nxdomain_rules"`
	Type          types.String `tfsdk:"type"`
}

var NIOSRulesetAttrTypes = map[string]attr.Type{
	"comment":        types.StringType,
	"disabled":       types.BoolType,
	"name":           types.StringType,
	"nxdomain_rules": types.ListType{ElemType: types.ObjectType{AttrTypes: RulesetNxdomainRulesAttrTypes}},
	"type":           types.StringType,
}

const (
	RulesetReturnFields = "comment,disabled,name,nxdomain_rules,type"
)

var RulesetResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          RulesetResourceNiosSchemaAttributes,
	},
}

var RulesetResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Descriptive comment about the Ruleset object.",
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that indicates if the Ruleset object is disabled.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of this Ruleset object.",
	},
	"nxdomain_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RulesetNxdomainRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of Rules assigned to this Ruleset object. Rules can be set only when the Ruleset type is set to \"NXDOMAIN\".",
	},
	"type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("NXDOMAIN", "BLACKLIST"),
		},
		Required:            true,
		MarkdownDescription: "The type of this Ruleset object.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *RulesetModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Ruleset {
	if m == nil {
		return nil
	}

	obj := &coremodel.Ruleset{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSRulesetModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSRulesetModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSRulesetExt {
	return &coremodel.NIOSRulesetExt{
		Comment:       flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Disabled:      flex.ExpandBoolPointer(m.Disabled),
		Name:          flex.ExpandStringPointerNullAsEmpty(m.Name),
		NxdomainRules: flex.ExpandFrameworkListNestedBlock(ctx, m.NxdomainRules, diags, ExpandRulesetNxdomainRules),
		Type:          flex.ExpandStringPointerNullAsEmpty(m.Type),
	}
}

// Flatten populates the TF model from a core response.
func (m *RulesetModel) Flatten(ctx context.Context, resp *coremodel.Ruleset, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSRulesetModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSRulesetModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSRulesetAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSRulesetAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSRulesetModel) Flatten(ctx context.Context, from *coremodel.NIOSRulesetExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Disabled = flex.FlattenBoolPointer(from.Disabled)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.NxdomainRules = flex.FlattenFrameworkListNestedBlock(ctx, from.NxdomainRules, RulesetNxdomainRulesAttrTypes, diags, FlattenRulesetNxdomainRules)
	m.Type = flex.FlattenStringPointerEmptyAsNull(from.Type)
}
