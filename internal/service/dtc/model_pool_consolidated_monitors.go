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

// PoolConsolidatedMonitorsModel is the Terraform model for PoolConsolidatedMonitors
type PoolConsolidatedMonitorsModel struct {
	Members                 types.List   `tfsdk:"members"`
	Monitor                 types.String `tfsdk:"monitor"`
	Availability            types.String `tfsdk:"availability"`
	FullHealthCommunication types.Bool   `tfsdk:"full_health_communication"`
}

// PoolConsolidatedMonitorsAttrTypes contains the attribute types for PoolConsolidatedMonitorsModel
var PoolConsolidatedMonitorsAttrTypes = map[string]attr.Type{
	"members":                   types.ListType{ElemType: types.StringType},
	"monitor":                   types.StringType,
	"availability":              types.StringType,
	"full_health_communication": types.BoolType,
}

// PoolConsolidatedMonitorsResourceSchemaAttributes contains the schema attributes for PoolConsolidatedMonitorsModel
var PoolConsolidatedMonitorsResourceSchemaAttributes = map[string]schema.Attribute{
	"members": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Members whose monitor statuses are shared across other members in a pool.",
	},
	"monitor": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Monitor whose statuses are shared across other members in a pool.",
	},
	"availability": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("ALL", "ANY"),
		},
		Optional:            true,
		MarkdownDescription: "Servers assigned to a pool with monitor defined are healthy if ANY or ALL members report healthy status.",
	},
	"full_health_communication": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Flag for switching health performing and sharing behavior to perform health checks on each DTC grid member that serves related LBDN(s) and send them across all DTC grid members from both selected and non-selected lists.",
	},
}

// ExpandPoolConsolidatedMonitors converts a Terraform Object to SDK type
func ExpandPoolConsolidatedMonitors(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdtc.DtcPoolConsolidatedMonitors {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m PoolConsolidatedMonitorsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *PoolConsolidatedMonitorsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdtc.DtcPoolConsolidatedMonitors {
	if m == nil {
		return nil
	}
	to := &niosdtc.DtcPoolConsolidatedMonitors{
		Members:                 flex.ExpandFrameworkListString(ctx, m.Members, diags),
		Monitor:                 flex.ExpandStringPointerNullAsEmpty(m.Monitor),
		Availability:            flex.ExpandStringPointerNullAsEmpty(m.Availability),
		FullHealthCommunication: flex.ExpandBoolPointer(m.FullHealthCommunication),
	}
	return to
}

// FlattenPoolConsolidatedMonitors converts an SDK type to Terraform Object
func FlattenPoolConsolidatedMonitors(ctx context.Context, from *niosdtc.DtcPoolConsolidatedMonitors, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(PoolConsolidatedMonitorsAttrTypes)
	}
	m := &PoolConsolidatedMonitorsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, PoolConsolidatedMonitorsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *PoolConsolidatedMonitorsModel) Flatten(ctx context.Context, from *niosdtc.DtcPoolConsolidatedMonitors, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Members = flex.FlattenFrameworkListString(ctx, from.Members, diags)
	m.Monitor = flex.FlattenStringPointerEmptyAsNull(from.Monitor)
	m.Availability = flex.FlattenStringPointerEmptyAsNull(from.Availability)
	m.FullHealthCommunication = flex.FlattenBoolPointer(from.FullHealthCommunication)
}
