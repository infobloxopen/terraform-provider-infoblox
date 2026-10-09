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

// InternalSecondaryModel is the Terraform model for InternalSecondary
type InternalSecondaryModel struct {
	Host types.String `tfsdk:"host"`
}

// InternalSecondaryAttrTypes contains the attribute types for InternalSecondaryModel
var InternalSecondaryAttrTypes = map[string]attr.Type{
	"host": types.StringType,
}

// InternalSecondaryResourceSchemaAttributes contains the schema attributes for InternalSecondaryModel
var InternalSecondaryResourceSchemaAttributes = map[string]schema.Attribute{
	"host": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The resource identifier.",
	},
}

// ExpandInternalSecondary converts a Terraform Object to SDK type
func ExpandInternalSecondary(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.InternalSecondary {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InternalSecondaryModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *InternalSecondaryModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.InternalSecondary {
	if m == nil {
		return nil
	}
	to := &uddidns.InternalSecondary{
		Host: flex.ExpandString(m.Host),
	}
	return to
}

// FlattenInternalSecondary converts an SDK type to Terraform Object
func FlattenInternalSecondary(ctx context.Context, from *uddidns.InternalSecondary, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InternalSecondaryAttrTypes)
	}
	m := &InternalSecondaryModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InternalSecondaryAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *InternalSecondaryModel) Flatten(ctx context.Context, from *uddidns.InternalSecondary, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Host = flex.FlattenString(from.Host)
}
