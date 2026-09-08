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

// LbdnPoolsModel is the Terraform model for LbdnPools
type LbdnPoolsModel struct {
	Pool  types.String `tfsdk:"pool"`
	Ratio types.Int64  `tfsdk:"ratio"`
}

// LbdnPoolsAttrTypes contains the attribute types for LbdnPoolsModel
var LbdnPoolsAttrTypes = map[string]attr.Type{
	"pool":  types.StringType,
	"ratio": types.Int64Type,
}

// LbdnPoolsResourceSchemaAttributes contains the schema attributes for LbdnPoolsModel
var LbdnPoolsResourceSchemaAttributes = map[string]schema.Attribute{
	"pool": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The pool to link with.",
	},
	"ratio": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The weight of pool.",
	},
}

// ExpandLbdnPools converts a Terraform Object to SDK type
func ExpandLbdnPools(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdtc.DtcLbdnPools {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m LbdnPoolsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *LbdnPoolsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdtc.DtcLbdnPools {
	if m == nil {
		return nil
	}
	to := &niosdtc.DtcLbdnPools{
		Pool:  flex.ExpandStringPointerNullAsEmpty(m.Pool),
		Ratio: flex.ExpandInt64Pointer(m.Ratio),
	}
	return to
}

// FlattenLbdnPools converts an SDK type to Terraform Object
func FlattenLbdnPools(ctx context.Context, from *niosdtc.DtcLbdnPools, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(LbdnPoolsAttrTypes)
	}
	m := &LbdnPoolsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, LbdnPoolsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *LbdnPoolsModel) Flatten(ctx context.Context, from *niosdtc.DtcLbdnPools, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Pool = flex.FlattenStringPointerEmptyAsNull(from.Pool)
	m.Ratio = flex.FlattenInt64Pointer(from.Ratio)
}
