package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/grid"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type DistributionscheduleModel struct {
	Id            types.String `tfsdk:"id"`
	UpdateTrigger types.String `tfsdk:"update_trigger"`
	NIOS          types.Object `tfsdk:"nios"`
}

var DistributionscheduleAttrTypes = map[string]attr.Type{
	"id":             types.StringType,
	"update_trigger": types.StringType,
	"nios":           types.ObjectType{AttrTypes: NIOSDistributionscheduleAttrTypes},
}

type NIOSDistributionscheduleModel struct {
	Active        types.Bool   `tfsdk:"active"`
	StartTime     types.String `tfsdk:"start_time"`
	TimeZone      types.String `tfsdk:"time_zone"`
	UpgradeGroups types.List   `tfsdk:"upgrade_groups"`
}

var NIOSDistributionscheduleAttrTypes = map[string]attr.Type{
	"active":         types.BoolType,
	"start_time":     types.StringType,
	"time_zone":      types.StringType,
	"upgrade_groups": types.ListType{ElemType: types.ObjectType{AttrTypes: DistributionscheduleUpgradeGroupsAttrTypes}},
}

const (
	DistributionscheduleReturnFields = "active,start_time,time_zone,upgrade_groups"
)

var DistributionscheduleResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"update_trigger": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "An arbitrary value used to trigger an update. Not sent to the API. Change it when Terraform reports no infrastructure changes.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          DistributionscheduleResourceNiosSchemaAttributes,
	},
}

var DistributionscheduleResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"active": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Determines whether the distribution schedule is active.",
	},
	"start_time": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTimeFormat(),
		},
		MarkdownDescription: "The start time of the distribution.",
	},
	"time_zone": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Time zone of the distribution start time.",
	},
	"upgrade_groups": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: DistributionscheduleUpgradeGroupsResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The upgrade groups scheduling settings.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *DistributionscheduleModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Distributionschedule {
	if m == nil {
		return nil
	}

	obj := &coremodel.Distributionschedule{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSDistributionscheduleModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
		obj.NIOS = PostExpandDistributionscheduleNIOS(ctx, obj.NIOS, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSDistributionscheduleModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSDistributionscheduleExt {
	return &coremodel.NIOSDistributionscheduleExt{
		Active:        flex.ExpandBoolPointer(m.Active),
		StartTime:     flex.ExpandTimeToUnix(m.StartTime, diags),
		UpgradeGroups: flex.ExpandFrameworkListNestedBlock(ctx, m.UpgradeGroups, diags, ExpandDistributionscheduleUpgradeGroups),
	}
}

// Flatten populates the TF model from a core response.
func (m *DistributionscheduleModel) Flatten(ctx context.Context, resp *coremodel.Distributionschedule, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSDistributionscheduleModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSDistributionscheduleModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSDistributionscheduleModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenDistributionscheduleNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSDistributionscheduleAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSDistributionscheduleAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSDistributionscheduleModel) Flatten(ctx context.Context, from *coremodel.NIOSDistributionscheduleExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Active = flex.FlattenBoolPointer(from.Active)
	m.StartTime = flex.FlattenUnixTime(from.StartTime, diags)
	m.TimeZone = flex.FlattenStringPointerEmptyAsNull(from.TimeZone)
	m.UpgradeGroups = flex.FlattenFrameworkListNestedBlock(ctx, from.UpgradeGroups, DistributionscheduleUpgradeGroupsAttrTypes, diags, FlattenDistributionscheduleUpgradeGroups)
}
