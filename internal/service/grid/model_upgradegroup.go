package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/grid"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type UpgradegroupModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var UpgradegroupAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSUpgradegroupAttrTypes},
}

type NIOSUpgradegroupModel struct {
	Comment                    types.String `tfsdk:"comment"`
	DistributionDependentGroup types.String `tfsdk:"distribution_dependent_group"`
	DistributionPolicy         types.String `tfsdk:"distribution_policy"`
	DistributionTime           types.String `tfsdk:"distribution_time"`
	Members                    types.List   `tfsdk:"members"`
	Name                       types.String `tfsdk:"name"`
	UpgradeDependentGroup      types.String `tfsdk:"upgrade_dependent_group"`
	UpgradePolicy              types.String `tfsdk:"upgrade_policy"`
	UpgradeTime                types.String `tfsdk:"upgrade_time"`
}

var NIOSUpgradegroupAttrTypes = map[string]attr.Type{
	"comment":                      types.StringType,
	"distribution_dependent_group": types.StringType,
	"distribution_policy":          types.StringType,
	"distribution_time":            types.StringType,
	"members":                      types.ListType{ElemType: types.ObjectType{AttrTypes: UpgradegroupMembersAttrTypes}},
	"name":                         types.StringType,
	"upgrade_dependent_group":      types.StringType,
	"upgrade_policy":               types.StringType,
	"upgrade_time":                 types.StringType,
}

const (
	UpgradegroupReturnFields = "comment,distribution_dependent_group,distribution_policy,distribution_time,members,name,time_zone,upgrade_dependent_group,upgrade_policy,upgrade_time"
)

var UpgradegroupResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          UpgradegroupResourceNiosSchemaAttributes,
	},
}

var UpgradegroupResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The upgrade group descriptive comment.",
	},
	"distribution_dependent_group": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The distribution dependent group name.",
	},
	"distribution_policy": schema.StringAttribute{
		Default: stringdefault.StaticString("SIMULTANEOUSLY"),
		Validators: []validator.String{
			stringvalidator.OneOf("SIMULTANEOUSLY", "SEQUENTIALLY"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The distribution scheduling policy.",
	},
	"distribution_time": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTimeFormat(),
		},
		MarkdownDescription: "The time of the next scheduled distribution.",
	},
	"members": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: UpgradegroupMembersResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The upgrade group members.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The upgrade group name.",
	},
	"upgrade_dependent_group": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The upgrade dependent group name.",
	},
	"upgrade_policy": schema.StringAttribute{
		Default: stringdefault.StaticString("SEQUENTIALLY"),
		Validators: []validator.String{
			stringvalidator.OneOf("SIMULTANEOUSLY", "SEQUENTIALLY"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The upgrade scheduling policy.",
	},
	"upgrade_time": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTimeFormat(),
		},
		MarkdownDescription: "The time of the next scheduled upgrade.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *UpgradegroupModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Upgradegroup {
	if m == nil {
		return nil
	}

	obj := &coremodel.Upgradegroup{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSUpgradegroupModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSUpgradegroupModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSUpgradegroupExt {
	return &coremodel.NIOSUpgradegroupExt{
		Comment:                    flex.ExpandStringPointerNullAsEmpty(m.Comment),
		DistributionDependentGroup: flex.ExpandStringPointer(m.DistributionDependentGroup),
		DistributionPolicy:         flex.ExpandStringPointerNullAsEmpty(m.DistributionPolicy),
		DistributionTime:           flex.ExpandTimeToUnix(m.DistributionTime, diags),
		Members:                    flex.ExpandFrameworkListNestedBlock(ctx, m.Members, diags, ExpandUpgradegroupMembers),
		Name:                       flex.ExpandStringPointerNullAsEmpty(m.Name),
		UpgradeDependentGroup:      flex.ExpandStringPointer(m.UpgradeDependentGroup),
		UpgradePolicy:              flex.ExpandStringPointerNullAsEmpty(m.UpgradePolicy),
		UpgradeTime:                flex.ExpandTimeToUnix(m.UpgradeTime, diags),
	}
}

// Flatten populates the TF model from a core response.
func (m *UpgradegroupModel) Flatten(ctx context.Context, resp *coremodel.Upgradegroup, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSUpgradegroupModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSUpgradegroupModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSUpgradegroupAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSUpgradegroupAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSUpgradegroupModel) Flatten(ctx context.Context, from *coremodel.NIOSUpgradegroupExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.DistributionDependentGroup = flex.FlattenStringPointerEmptyAsNull(from.DistributionDependentGroup)
	m.DistributionPolicy = flex.FlattenStringPointerEmptyAsNull(from.DistributionPolicy)
	m.DistributionTime = flex.FlattenUnixTime(from.DistributionTime, diags)
	m.Members = flex.FlattenFrameworkListNestedBlock(ctx, from.Members, UpgradegroupMembersAttrTypes, diags, FlattenUpgradegroupMembers)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.UpgradeDependentGroup = flex.FlattenStringPointerEmptyAsNull(from.UpgradeDependentGroup)
	m.UpgradePolicy = flex.FlattenStringPointerEmptyAsNull(from.UpgradePolicy)
	m.UpgradeTime = flex.FlattenUnixTime(from.UpgradeTime, diags)
}
