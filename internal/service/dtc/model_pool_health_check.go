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

// PoolHealthCheckModel is the Terraform model for PoolHealthCheck
type PoolHealthCheckModel struct {
	ConsolidatedHealthCheck types.Object `tfsdk:"consolidated_health_check"`
	HealthCheckId           types.String `tfsdk:"health_check_id"`
	Name                    types.String `tfsdk:"name"`
}

// PoolHealthCheckAttrTypes contains the attribute types for PoolHealthCheckModel
var PoolHealthCheckAttrTypes = map[string]attr.Type{
	"consolidated_health_check": types.ObjectType{AttrTypes: ConsolidatedHealthCheckAttrTypes},
	"health_check_id":           types.StringType,
	"name":                      types.StringType,
}

// PoolHealthCheckResourceSchemaAttributes contains the schema attributes for PoolHealthCheckModel
var PoolHealthCheckResourceSchemaAttributes = map[string]schema.Attribute{
	"consolidated_health_check": schema.SingleNestedAttribute{
		Attributes:          ConsolidatedHealthCheckResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Optional. Consolidated health check configuration. When set, the DNS server running the named designator __DNS Service__ performs the health check on behalf of the __Pool__, and the result is shared with other DNS servers linked to the __Pool__ via __LBDN__ association. When unset, each DNS server performs the health check independently and no health status is shared.",
	},
	"health_check_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Display name of __HealthCheck__.",
	},
}

// ExpandPoolHealthCheck converts a Terraform Object to SDK type
func ExpandPoolHealthCheck(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidtc.PoolHealthCheck {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m PoolHealthCheckModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *PoolHealthCheckModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidtc.PoolHealthCheck {
	if m == nil {
		return nil
	}
	to := &uddidtc.PoolHealthCheck{
		ConsolidatedHealthCheck: ExpandConsolidatedHealthCheck(ctx, m.ConsolidatedHealthCheck, diags),
		HealthCheckId:           flex.ExpandString(m.HealthCheckId),
		Name:                    flex.ExpandStringPointer(m.Name),
	}
	return to
}

// FlattenPoolHealthCheck converts an SDK type to Terraform Object
func FlattenPoolHealthCheck(ctx context.Context, from *uddidtc.PoolHealthCheck, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(PoolHealthCheckAttrTypes)
	}
	m := &PoolHealthCheckModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, PoolHealthCheckAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *PoolHealthCheckModel) Flatten(ctx context.Context, from *uddidtc.PoolHealthCheck, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.ConsolidatedHealthCheck = FlattenConsolidatedHealthCheck(ctx, from.ConsolidatedHealthCheck, diags)
	m.HealthCheckId = flex.FlattenString(from.HealthCheckId)
	m.Name = flex.FlattenStringPointer(from.Name)
}
