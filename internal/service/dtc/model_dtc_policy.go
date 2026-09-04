package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddidtc "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// DTCPolicyModel is the Terraform model for DTCPolicy
type DTCPolicyModel struct {
	Name     types.String `tfsdk:"name"`
	PolicyId types.String `tfsdk:"policy_id"`
}

// DTCPolicyAttrTypes contains the attribute types for DTCPolicyModel
var DTCPolicyAttrTypes = map[string]attr.Type{
	"name":      types.StringType,
	"policy_id": types.StringType,
}

// DTCPolicyResourceSchemaAttributes contains the schema attributes for DTCPolicyModel
var DTCPolicyResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "__DTC Policy__ display name.",
	},
	"policy_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
}

// ExpandDTCPolicy converts a Terraform Object to SDK type
func ExpandDTCPolicy(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidtc.DTCPolicy {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DTCPolicyModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *DTCPolicyModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidtc.DTCPolicy {
	if m == nil {
		return nil
	}
	to := &uddidtc.DTCPolicy{
		Name:     flex.ExpandStringPointer(m.Name),
		PolicyId: flex.ExpandStringPointer(m.PolicyId),
	}
	return to
}

// FlattenDTCPolicy converts an SDK type to Terraform Object
func FlattenDTCPolicy(ctx context.Context, from *uddidtc.DTCPolicy, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DTCPolicyAttrTypes)
	}
	m := &DTCPolicyModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DTCPolicyAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *DTCPolicyModel) Flatten(ctx context.Context, from *uddidtc.DTCPolicy, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.PolicyId = flex.FlattenStringPointer(from.PolicyId)
}
