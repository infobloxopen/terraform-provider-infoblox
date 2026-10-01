package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosgrid "github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// DistributionscheduleUpgradeGroupsModel is the Terraform model for DistributionscheduleUpgradeGroups
type DistributionscheduleUpgradeGroupsModel struct {
	Name                       types.String `tfsdk:"name"`
	TimeZone                   types.String `tfsdk:"time_zone"`
	DistributionDependentGroup types.String `tfsdk:"distribution_dependent_group"`
	UpgradeDependentGroup      types.String `tfsdk:"upgrade_dependent_group"`
	DistributionTime           types.String `tfsdk:"distribution_time"`
	UpgradeTime                types.Int64  `tfsdk:"upgrade_time"`
}

// DistributionscheduleUpgradeGroupsAttrTypes contains the attribute types for DistributionscheduleUpgradeGroupsModel
var DistributionscheduleUpgradeGroupsAttrTypes = map[string]attr.Type{
	"name":                         types.StringType,
	"time_zone":                    types.StringType,
	"distribution_dependent_group": types.StringType,
	"upgrade_dependent_group":      types.StringType,
	"distribution_time":            types.StringType,
	"upgrade_time":                 types.Int64Type,
}

// DistributionscheduleUpgradeGroupsResourceSchemaAttributes contains the schema attributes for DistributionscheduleUpgradeGroupsModel
var DistributionscheduleUpgradeGroupsResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The upgrade group name.",
	},
	"time_zone": schema.StringAttribute{
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The time zone for scheduling operations.",
	},
	"distribution_dependent_group": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The distribution dependent group name.",
	},
	"upgrade_dependent_group": schema.StringAttribute{
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The upgrade dependent group name.",
	},
	"distribution_time": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTimeFormat(),
		},
		MarkdownDescription: "The time of the next scheduled distribution.",
	},
	"upgrade_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The time of the next scheduled upgrade.",
	},
}

// ExpandDistributionscheduleUpgradeGroups converts a Terraform Object to SDK type
func ExpandDistributionscheduleUpgradeGroups(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosgrid.DistributionscheduleUpgradeGroups {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DistributionscheduleUpgradeGroupsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *DistributionscheduleUpgradeGroupsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosgrid.DistributionscheduleUpgradeGroups {
	if m == nil {
		return nil
	}
	to := &niosgrid.DistributionscheduleUpgradeGroups{
		Name:                       flex.ExpandStringPointerNullAsEmpty(m.Name),
		DistributionDependentGroup: flex.ExpandStringPointer(m.DistributionDependentGroup),
		UpgradeDependentGroup:      flex.ExpandStringPointer(m.UpgradeDependentGroup),
		DistributionTime:           flex.ExpandTimeToUnix(m.DistributionTime, diags),
		UpgradeTime:                flex.ExpandInt64Pointer(m.UpgradeTime),
	}
	return to
}

// FlattenDistributionscheduleUpgradeGroups converts an SDK type to Terraform Object
func FlattenDistributionscheduleUpgradeGroups(ctx context.Context, from *niosgrid.DistributionscheduleUpgradeGroups, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DistributionscheduleUpgradeGroupsAttrTypes)
	}
	m := &DistributionscheduleUpgradeGroupsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DistributionscheduleUpgradeGroupsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *DistributionscheduleUpgradeGroupsModel) Flatten(ctx context.Context, from *niosgrid.DistributionscheduleUpgradeGroups, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.TimeZone = flex.FlattenStringPointerEmptyAsNull(from.TimeZone)
	m.DistributionDependentGroup = flex.FlattenStringPointerEmptyAsNull(from.DistributionDependentGroup)
	m.UpgradeDependentGroup = flex.FlattenStringPointerEmptyAsNull(from.UpgradeDependentGroup)
	m.DistributionTime = flex.FlattenUnixTime(from.DistributionTime, diags)
	m.UpgradeTime = flex.FlattenInt64Pointer(from.UpgradeTime)
}
