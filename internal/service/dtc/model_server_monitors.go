package dtc

import (
	"context"

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

// ServerMonitorsModel is the Terraform model for ServerMonitors
type ServerMonitorsModel struct {
	Monitor types.String `tfsdk:"monitor"`
	Host    types.String `tfsdk:"host"`
}

// ServerMonitorsAttrTypes contains the attribute types for ServerMonitorsModel
var ServerMonitorsAttrTypes = map[string]attr.Type{
	"monitor": types.StringType,
	"host":    types.StringType,
}

// ServerMonitorsResourceSchemaAttributes contains the schema attributes for ServerMonitorsModel
var ServerMonitorsResourceSchemaAttributes = map[string]schema.Attribute{
	"monitor": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The monitor related to server.",
	},
	"host": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "IP address or FQDN of the server used for monitoring.",
	},
}

// ExpandServerMonitors converts a Terraform Object to SDK type
func ExpandServerMonitors(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdtc.DtcServerMonitors {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ServerMonitorsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ServerMonitorsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdtc.DtcServerMonitors {
	if m == nil {
		return nil
	}
	to := &niosdtc.DtcServerMonitors{
		Monitor: flex.ExpandStringPointerNullAsEmpty(m.Monitor),
		Host:    flex.ExpandStringPointerNullAsEmpty(m.Host),
	}
	return to
}

// FlattenServerMonitors converts an SDK type to Terraform Object
func FlattenServerMonitors(ctx context.Context, from *niosdtc.DtcServerMonitors, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ServerMonitorsAttrTypes)
	}
	m := &ServerMonitorsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ServerMonitorsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ServerMonitorsModel) Flatten(ctx context.Context, from *niosdtc.DtcServerMonitors, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Monitor = flex.FlattenStringPointerEmptyAsNull(from.Monitor)
	m.Host = flex.FlattenStringPointerEmptyAsNull(from.Host)
}
