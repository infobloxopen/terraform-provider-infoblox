package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
	uddidtc "github.com/infobloxopen/universal-ddi-go-client/dtc"
)

// ConsolidatedHealthCheckModel is the Terraform model for ConsolidatedHealthCheck
type ConsolidatedHealthCheckModel struct {
	Designators types.List `tfsdk:"designators"`
}

// ConsolidatedHealthCheckAttrTypes contains the attribute types for ConsolidatedHealthCheckModel
var ConsolidatedHealthCheckAttrTypes = map[string]attr.Type{
	"designators": types.ListType{ElemType: types.ObjectType{AttrTypes: DesignatorServiceAttrTypes}},
}

// ConsolidatedHealthCheckResourceSchemaAttributes contains the schema attributes for ConsolidatedHealthCheckModel
var ConsolidatedHealthCheckResourceSchemaAttributes = map[string]schema.Attribute{
	"designators": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: DesignatorServiceResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Designator __DNS Service__ references where the corresponding health checks will be associated to. Must contain at least one entry when set.  On request: only _dns_service_id_ is honoured. On response: _dns_service_name_ is echoed alongside, resolved from inventory.",
	},
}

// ExpandConsolidatedHealthCheck converts a Terraform Object to SDK type
func ExpandConsolidatedHealthCheck(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidtc.ConsolidatedHealthCheck {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConsolidatedHealthCheckModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ConsolidatedHealthCheckModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidtc.ConsolidatedHealthCheck {
	if m == nil {
		return nil
	}
	to := &uddidtc.ConsolidatedHealthCheck{
		Designators: flex.ExpandFrameworkListNestedBlock(ctx, m.Designators, diags, ExpandDesignatorService),
	}
	return to
}

// FlattenConsolidatedHealthCheck converts an SDK type to Terraform Object
func FlattenConsolidatedHealthCheck(ctx context.Context, from *uddidtc.ConsolidatedHealthCheck, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConsolidatedHealthCheckAttrTypes)
	}
	m := &ConsolidatedHealthCheckModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConsolidatedHealthCheckAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ConsolidatedHealthCheckModel) Flatten(ctx context.Context, from *uddidtc.ConsolidatedHealthCheck, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Designators = flex.FlattenFrameworkListNestedBlock(ctx, from.Designators, DesignatorServiceAttrTypes, diags, FlattenDesignatorService)
}
