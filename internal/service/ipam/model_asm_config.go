package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// ASMConfigModel is the Terraform model for ASMConfig
type ASMConfigModel struct {
	AsmThreshold       types.Int64       `tfsdk:"asm_threshold"`
	Enable             types.Bool        `tfsdk:"enable"`
	EnableNotification types.Bool        `tfsdk:"enable_notification"`
	ForecastPeriod     types.Int64       `tfsdk:"forecast_period"`
	GrowthFactor       types.Int64       `tfsdk:"growth_factor"`
	GrowthType         types.String      `tfsdk:"growth_type"`
	History            types.Int64       `tfsdk:"history"`
	MinTotal           types.Int64       `tfsdk:"min_total"`
	MinUnused          types.Int64       `tfsdk:"min_unused"`
	ReenableDate       timetypes.RFC3339 `tfsdk:"reenable_date"`
}

// ASMConfigAttrTypes contains the attribute types for ASMConfigModel
var ASMConfigAttrTypes = map[string]attr.Type{
	"asm_threshold":       types.Int64Type,
	"enable":              types.BoolType,
	"enable_notification": types.BoolType,
	"forecast_period":     types.Int64Type,
	"growth_factor":       types.Int64Type,
	"growth_type":         types.StringType,
	"history":             types.Int64Type,
	"min_total":           types.Int64Type,
	"min_unused":          types.Int64Type,
	"reenable_date":       timetypes.RFC3339Type{},
}

// ASMConfigResourceSchemaAttributes contains the schema attributes for ASMConfigModel
var ASMConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"asm_threshold": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(90),
		MarkdownDescription: "ASM shows the number of addresses forecast to be used _forecast_period_ days in the future, if it is greater than _asm_threshold_ percent * _dhcp_total_ (see _dhcp_utilization_) then the subnet is flagged.",
	},
	"enable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Indicates if Automated Scope Management is enabled.",
	},
	"enable_notification": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Indicates if ASM should send notifications to the user.",
	},
	"forecast_period": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(14),
		MarkdownDescription: "The forecast period in days.",
	},
	"growth_factor": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(20),
		MarkdownDescription: "Indicates the growth in the number or percentage of IP addresses.",
	},
	"growth_type": schema.StringAttribute{
		Default:             stringdefault.StaticString("percent"),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The type of factor to use: _percent_ or _count_.",
	},
	"history": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(30),
		MarkdownDescription: "The minimum amount of history needed before ASM can run on this subnet.",
	},
	"min_total": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(10),
		MarkdownDescription: "The minimum size of range needed for ASM to run on this subnet.",
	},
	"min_unused": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(10),
		MarkdownDescription: "The minimum percentage of addresses that must be available outside of the DHCP ranges and fixed addresses when making a suggested change..",
	},
	"reenable_date": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		CustomType:          timetypes.RFC3339Type{},
		Default:             stringdefault.StaticString("1970-01-01T00:00:00Z"),
		MarkdownDescription: "",
	},
}

// ExpandASMConfig converts a Terraform Object to SDK type
func ExpandASMConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiipam.ASMConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ASMConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ASMConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiipam.ASMConfig {
	if m == nil {
		return nil
	}
	to := &uddiipam.ASMConfig{
		AsmThreshold:       flex.ExpandInt64Pointer(m.AsmThreshold),
		Enable:             flex.ExpandBoolPointer(m.Enable),
		EnableNotification: flex.ExpandBoolPointer(m.EnableNotification),
		ForecastPeriod:     flex.ExpandInt64Pointer(m.ForecastPeriod),
		GrowthFactor:       flex.ExpandInt64Pointer(m.GrowthFactor),
		GrowthType:         flex.ExpandStringPointer(m.GrowthType),
		History:            flex.ExpandInt64Pointer(m.History),
		MinTotal:           flex.ExpandInt64Pointer(m.MinTotal),
		MinUnused:          flex.ExpandInt64Pointer(m.MinUnused),
		ReenableDate:       flex.ExpandRFC3339(m.ReenableDate, diags),
	}
	return to
}

// FlattenASMConfig converts an SDK type to Terraform Object
func FlattenASMConfig(ctx context.Context, from *uddiipam.ASMConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ASMConfigAttrTypes)
	}
	m := &ASMConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ASMConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ASMConfigModel) Flatten(ctx context.Context, from *uddiipam.ASMConfig, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AsmThreshold = flex.FlattenInt64Pointer(from.AsmThreshold)
	m.Enable = flex.FlattenBoolPointer(from.Enable)
	m.EnableNotification = flex.FlattenBoolPointer(from.EnableNotification)
	m.ForecastPeriod = flex.FlattenInt64Pointer(from.ForecastPeriod)
	m.GrowthFactor = flex.FlattenInt64Pointer(from.GrowthFactor)
	m.GrowthType = flex.FlattenStringPointer(from.GrowthType)
	m.History = flex.FlattenInt64Pointer(from.History)
	m.MinTotal = flex.FlattenInt64Pointer(from.MinTotal)
	m.MinUnused = flex.FlattenInt64Pointer(from.MinUnused)
	m.ReenableDate = flex.FlattenRFC3339(from.ReenableDate)
}
