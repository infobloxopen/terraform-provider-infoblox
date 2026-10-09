package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	niosdtc "github.com/infobloxopen/infoblox-nios-go-client/dtc"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// PoolLbDynamicRatioAlternateModel is the Terraform model for PoolLbDynamicRatioAlternate
type PoolLbDynamicRatioAlternateModel struct {
	Method              types.String `tfsdk:"method"`
	Monitor             types.String `tfsdk:"monitor"`
	MonitorMetric       types.String `tfsdk:"monitor_metric"`
	MonitorWeighing     types.String `tfsdk:"monitor_weighing"`
	InvertMonitorMetric types.Bool   `tfsdk:"invert_monitor_metric"`
}

// PoolLbDynamicRatioAlternateAttrTypes contains the attribute types for PoolLbDynamicRatioAlternateModel
var PoolLbDynamicRatioAlternateAttrTypes = map[string]attr.Type{
	"method":                types.StringType,
	"monitor":               types.StringType,
	"monitor_metric":        types.StringType,
	"monitor_weighing":      types.StringType,
	"invert_monitor_metric": types.BoolType,
}

// PoolLbDynamicRatioAlternateResourceSchemaAttributes contains the schema attributes for PoolLbDynamicRatioAlternateModel
var PoolLbDynamicRatioAlternateResourceSchemaAttributes = map[string]schema.Attribute{
	"method": schema.StringAttribute{
		Default: stringdefault.StaticString("MONITOR"),
		Validators: []validator.String{
			stringvalidator.OneOf("MONITOR", "ROUND_TRIP_DELAY"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The method of the DTC dynamic ratio load balancing.",
	},
	"monitor": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The DTC monitor output of which will be used for dynamic ratio load balancing.",
	},
	"monitor_metric": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The metric of the DTC SNMP monitor that will be used for dynamic weighing.",
	},
	"monitor_weighing": schema.StringAttribute{
		Default: stringdefault.StaticString("RATIO"),
		Validators: []validator.String{
			stringvalidator.OneOf("PRIORITY", "RATIO"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The DTC monitor weight. 'PRIORITY' means that all clients will be forwarded to the least loaded server. 'RATIO' means that distribution will be calculated based on dynamic weights.",
	},
	"invert_monitor_metric": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the inverted values of the DTC SNMP monitor metric will be used.",
	},
}

// ExpandPoolLbDynamicRatioAlternate converts a Terraform Object to SDK type
func ExpandPoolLbDynamicRatioAlternate(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdtc.DtcPoolLbDynamicRatioAlternate {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m PoolLbDynamicRatioAlternateModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *PoolLbDynamicRatioAlternateModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdtc.DtcPoolLbDynamicRatioAlternate {
	if m == nil {
		return nil
	}
	to := &niosdtc.DtcPoolLbDynamicRatioAlternate{
		Method:              flex.ExpandStringPointerNullAsEmpty(m.Method),
		Monitor:             flex.ExpandStringPointer(m.Monitor),
		MonitorMetric:       flex.ExpandStringPointerNullAsEmpty(m.MonitorMetric),
		MonitorWeighing:     flex.ExpandStringPointerNullAsEmpty(m.MonitorWeighing),
		InvertMonitorMetric: flex.ExpandBoolPointer(m.InvertMonitorMetric),
	}
	return to
}

// FlattenPoolLbDynamicRatioAlternate converts an SDK type to Terraform Object
func FlattenPoolLbDynamicRatioAlternate(ctx context.Context, from *niosdtc.DtcPoolLbDynamicRatioAlternate, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(PoolLbDynamicRatioAlternateAttrTypes)
	}
	m := &PoolLbDynamicRatioAlternateModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, PoolLbDynamicRatioAlternateAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *PoolLbDynamicRatioAlternateModel) Flatten(ctx context.Context, from *niosdtc.DtcPoolLbDynamicRatioAlternate, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Method = flex.FlattenStringPointerEmptyAsNull(from.Method)
	m.Monitor = flex.FlattenStringPointerEmptyAsNull(from.Monitor)
	m.MonitorMetric = flex.FlattenStringPointerEmptyAsNull(from.MonitorMetric)
	m.MonitorWeighing = flex.FlattenStringPointerEmptyAsNull(from.MonitorWeighing)
	m.InvertMonitorMetric = flex.FlattenBoolPointer(from.InvertMonitorMetric)
}
