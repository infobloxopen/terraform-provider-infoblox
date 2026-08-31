package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosgrid "github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// ExtensibleattributedefListValuesModel is the Terraform model for ExtensibleattributedefListValues
type ExtensibleattributedefListValuesModel struct {
	Value types.String `tfsdk:"value"`
}

// ExtensibleattributedefListValuesAttrTypes contains the attribute types for ExtensibleattributedefListValuesModel
var ExtensibleattributedefListValuesAttrTypes = map[string]attr.Type{
	"value": types.StringType,
}

// ExtensibleattributedefListValuesResourceSchemaAttributes contains the schema attributes for ExtensibleattributedefListValuesModel
var ExtensibleattributedefListValuesResourceSchemaAttributes = map[string]schema.Attribute{
	"value": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Enum value",
	},
}

// ExpandExtensibleattributedefListValues converts a Terraform Object to SDK type
func ExpandExtensibleattributedefListValues(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosgrid.ExtensibleattributedefListValues {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ExtensibleattributedefListValuesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ExtensibleattributedefListValuesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosgrid.ExtensibleattributedefListValues {
	if m == nil {
		return nil
	}
	to := &niosgrid.ExtensibleattributedefListValues{
		Value: flex.ExpandStringPointerNullAsEmpty(m.Value),
	}
	return to
}

// FlattenExtensibleattributedefListValues converts an SDK type to Terraform Object
func FlattenExtensibleattributedefListValues(ctx context.Context, from *niosgrid.ExtensibleattributedefListValues, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ExtensibleattributedefListValuesAttrTypes)
	}
	m := &ExtensibleattributedefListValuesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ExtensibleattributedefListValuesAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ExtensibleattributedefListValuesModel) Flatten(ctx context.Context, from *niosgrid.ExtensibleattributedefListValues, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Value = flex.FlattenStringPointerEmptyAsNull(from.Value)
}
