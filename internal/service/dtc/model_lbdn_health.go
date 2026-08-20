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

// LbdnHealthModel is the Terraform model for LbdnHealth
type LbdnHealthModel struct {
	Availability types.String `tfsdk:"availability"`
	EnabledState types.String `tfsdk:"enabled_state"`
	Description  types.String `tfsdk:"description"`
}

// LbdnHealthAttrTypes contains the attribute types for LbdnHealthModel
var LbdnHealthAttrTypes = map[string]attr.Type{
	"availability":  types.StringType,
	"enabled_state": types.StringType,
	"description":   types.StringType,
}

// LbdnHealthResourceSchemaAttributes contains the schema attributes for LbdnHealthModel
var LbdnHealthResourceSchemaAttributes = map[string]schema.Attribute{
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

// ExpandLbdnHealth converts a Terraform Object to SDK type
func ExpandLbdnHealth(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdtc.DtcLbdnHealth {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m LbdnHealthModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *LbdnHealthModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdtc.DtcLbdnHealth {
	if m == nil {
		return nil
	}
	to := &niosdtc.DtcLbdnHealth{
		Availability: flex.ExpandStringPointerNullAsEmpty(m.Availability),
		EnabledState: flex.ExpandStringPointerNullAsEmpty(m.EnabledState),
		Description:  flex.ExpandStringPointerNullAsEmpty(m.Description),
	}
	return to
}

// FlattenLbdnHealth converts an SDK type to Terraform Object
func FlattenLbdnHealth(ctx context.Context, from *niosdtc.DtcLbdnHealth, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(LbdnHealthAttrTypes)
	}
	m := &LbdnHealthModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, LbdnHealthAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *LbdnHealthModel) Flatten(ctx context.Context, from *niosdtc.DtcLbdnHealth, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Availability = flex.FlattenStringPointerEmptyAsNull(from.Availability)
	m.EnabledState = flex.FlattenStringPointerEmptyAsNull(from.EnabledState)
	m.Description = flex.FlattenStringPointerEmptyAsNull(from.Description)
}
