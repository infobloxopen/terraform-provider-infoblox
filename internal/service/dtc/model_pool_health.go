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

// PoolHealthModel is the Terraform model for PoolHealth
type PoolHealthModel struct {
	Availability types.String `tfsdk:"availability"`
	EnabledState types.String `tfsdk:"enabled_state"`
	Description  types.String `tfsdk:"description"`
}

// PoolHealthAttrTypes contains the attribute types for PoolHealthModel
var PoolHealthAttrTypes = map[string]attr.Type{
	"availability":  types.StringType,
	"enabled_state": types.StringType,
	"description":   types.StringType,
}

// PoolHealthResourceSchemaAttributes contains the schema attributes for PoolHealthModel
var PoolHealthResourceSchemaAttributes = map[string]schema.Attribute{
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

// ExpandPoolHealth converts a Terraform Object to SDK type
func ExpandPoolHealth(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdtc.DtcPoolHealth {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m PoolHealthModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *PoolHealthModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdtc.DtcPoolHealth {
	if m == nil {
		return nil
	}
	to := &niosdtc.DtcPoolHealth{
		Availability: flex.ExpandStringPointerNullAsEmpty(m.Availability),
		EnabledState: flex.ExpandStringPointerNullAsEmpty(m.EnabledState),
		Description:  flex.ExpandStringPointerNullAsEmpty(m.Description),
	}
	return to
}

// FlattenPoolHealth converts an SDK type to Terraform Object
func FlattenPoolHealth(ctx context.Context, from *niosdtc.DtcPoolHealth, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(PoolHealthAttrTypes)
	}
	m := &PoolHealthModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, PoolHealthAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *PoolHealthModel) Flatten(ctx context.Context, from *niosdtc.DtcPoolHealth, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Availability = flex.FlattenStringPointerEmptyAsNull(from.Availability)
	m.EnabledState = flex.FlattenStringPointerEmptyAsNull(from.EnabledState)
	m.Description = flex.FlattenStringPointerEmptyAsNull(from.Description)
}
