package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddidtc "github.com/infobloxopen/universal-ddi-go-client/dtc"
)

// PoolServerModel is the Terraform model for PoolServer
type PoolServerModel struct {
	Name     types.String `tfsdk:"name"`
	ServerId types.String `tfsdk:"server_id"`
	Weight   types.Int64  `tfsdk:"weight"`
}

// PoolServerAttrTypes contains the attribute types for PoolServerModel
var PoolServerAttrTypes = map[string]attr.Type{
	"name":      types.StringType,
	"server_id": types.StringType,
	"weight":    types.Int64Type,
}

// PoolServerResourceSchemaAttributes contains the schema attributes for PoolServerModel
var PoolServerResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Display name of __Server__.",
	},
	"server_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"weight": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Weight of __Server__ to be used for load balancing. Unsigned integer, min 1; max 65535.",
	},
}

// ExpandPoolServer converts a Terraform Object to SDK type
func ExpandPoolServer(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidtc.PoolServer {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m PoolServerModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *PoolServerModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidtc.PoolServer {
	if m == nil {
		return nil
	}
	to := &uddidtc.PoolServer{
		Name:     flex.ExpandStringPointer(m.Name),
		ServerId: flex.ExpandString(m.ServerId),
		Weight:   flex.ExpandInt64Pointer(m.Weight),
	}
	return to
}

// FlattenPoolServer converts an SDK type to Terraform Object
func FlattenPoolServer(ctx context.Context, from *uddidtc.PoolServer, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(PoolServerAttrTypes)
	}
	m := &PoolServerModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, PoolServerAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *PoolServerModel) Flatten(ctx context.Context, from *uddidtc.PoolServer, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.ServerId = flex.FlattenString(from.ServerId)
	m.Weight = flex.FlattenInt64Pointer(from.Weight)
}
