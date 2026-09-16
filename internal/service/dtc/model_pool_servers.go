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

// PoolServersModel is the Terraform model for PoolServers
type PoolServersModel struct {
	Server types.String `tfsdk:"server"`
	Ratio  types.Int64  `tfsdk:"ratio"`
}

// PoolServersAttrTypes contains the attribute types for PoolServersModel
var PoolServersAttrTypes = map[string]attr.Type{
	"server": types.StringType,
	"ratio":  types.Int64Type,
}

// PoolServersResourceSchemaAttributes contains the schema attributes for PoolServersModel
var PoolServersResourceSchemaAttributes = map[string]schema.Attribute{
	"server": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The server to link with.",
	},
	"ratio": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The weight of server.",
	},
}

// ExpandPoolServers converts a Terraform Object to SDK type
func ExpandPoolServers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdtc.DtcPoolServers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m PoolServersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *PoolServersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdtc.DtcPoolServers {
	if m == nil {
		return nil
	}
	to := &niosdtc.DtcPoolServers{
		Server: flex.ExpandStringPointerNullAsEmpty(m.Server),
		Ratio:  flex.ExpandInt64Pointer(m.Ratio),
	}
	return to
}

// FlattenPoolServers converts an SDK type to Terraform Object
func FlattenPoolServers(ctx context.Context, from *niosdtc.DtcPoolServers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(PoolServersAttrTypes)
	}
	m := &PoolServersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, PoolServersAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *PoolServersModel) Flatten(ctx context.Context, from *niosdtc.DtcPoolServers, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Server = flex.FlattenStringPointerEmptyAsNull(from.Server)
	m.Ratio = flex.FlattenInt64Pointer(from.Ratio)
}
