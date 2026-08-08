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

// ServerHealthModel is the Terraform model for ServerHealth
type ServerHealthModel struct {
	Availability types.String `tfsdk:"availability"`
	EnabledState types.String `tfsdk:"enabled_state"`
	Description  types.String `tfsdk:"description"`
}

// ServerHealthAttrTypes contains the attribute types for ServerHealthModel
var ServerHealthAttrTypes = map[string]attr.Type{
	"availability":  types.StringType,
	"enabled_state": types.StringType,
	"description":   types.StringType,
}

// ServerHealthResourceSchemaAttributes contains the schema attributes for ServerHealthModel
var ServerHealthResourceSchemaAttributes = map[string]schema.Attribute{
	"availability": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("NONE", "GREEN", "YELLOW", "RED", "BLUE", "GRAY"),
		},
		Optional:            true,
		MarkdownDescription: "The availability color status.",
	},
	"enabled_state": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("NONE", "ENABLED", "DISABLED", "DISABLED_BY_PARENT"),
		},
		Optional:            true,
		MarkdownDescription: "The enabled state of the object.",
	},
	"description": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The textual description of the object's status.",
	},
}

// ExpandServerHealth converts a Terraform Object to SDK type
func ExpandServerHealth(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdtc.DtcServerHealth {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ServerHealthModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ServerHealthModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdtc.DtcServerHealth {
	if m == nil {
		return nil
	}
	to := &niosdtc.DtcServerHealth{
		Availability: flex.ExpandStringPointerNullAsEmpty(m.Availability),
		EnabledState: flex.ExpandStringPointerNullAsEmpty(m.EnabledState),
		Description:  flex.ExpandStringPointerNullAsEmpty(m.Description),
	}
	return to
}

// FlattenServerHealth converts an SDK type to Terraform Object
func FlattenServerHealth(ctx context.Context, from *niosdtc.DtcServerHealth, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ServerHealthAttrTypes)
	}
	m := &ServerHealthModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ServerHealthAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ServerHealthModel) Flatten(ctx context.Context, from *niosdtc.DtcServerHealth, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Availability = flex.FlattenStringPointerEmptyAsNull(from.Availability)
	m.EnabledState = flex.FlattenStringPointerEmptyAsNull(from.EnabledState)
	m.Description = flex.FlattenStringPointerEmptyAsNull(from.Description)
}
