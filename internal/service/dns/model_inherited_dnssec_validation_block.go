package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddidns "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// InheritedDNSSECValidationBlockModel is the Terraform model for InheritedDNSSECValidationBlock
type InheritedDNSSECValidationBlockModel struct {
	Action types.String `tfsdk:"action"`
}

// InheritedDNSSECValidationBlockAttrTypes contains the attribute types for InheritedDNSSECValidationBlockModel
var InheritedDNSSECValidationBlockAttrTypes = map[string]attr.Type{
	"action": types.StringType,
}

// InheritedDNSSECValidationBlockResourceSchemaAttributes contains the schema attributes for InheritedDNSSECValidationBlockModel
var InheritedDNSSECValidationBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Defaults to _inherit_.",
	},
}

// ExpandInheritedDNSSECValidationBlock converts a Terraform Object to SDK type
func ExpandInheritedDNSSECValidationBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.InheritedDNSSECValidationBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritedDNSSECValidationBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *InheritedDNSSECValidationBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.InheritedDNSSECValidationBlock {
	if m == nil {
		return nil
	}
	to := &uddidns.InheritedDNSSECValidationBlock{
		Action: flex.ExpandStringPointer(m.Action),
	}
	return to
}

// FlattenInheritedDNSSECValidationBlock converts an SDK type to Terraform Object
func FlattenInheritedDNSSECValidationBlock(ctx context.Context, from *uddidns.InheritedDNSSECValidationBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritedDNSSECValidationBlockAttrTypes)
	}
	m := &InheritedDNSSECValidationBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritedDNSSECValidationBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *InheritedDNSSECValidationBlockModel) Flatten(ctx context.Context, from *uddidns.InheritedDNSSECValidationBlock, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
}
