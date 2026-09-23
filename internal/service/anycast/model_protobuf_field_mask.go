package anycast

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
	uddianycast "github.com/infobloxopen/universal-ddi-go-client/anycast"
)

// ProtobufFieldMaskModel is the Terraform model for ProtobufFieldMask
type ProtobufFieldMaskModel struct {
	Paths types.List `tfsdk:"paths"`
}

// ProtobufFieldMaskAttrTypes contains the attribute types for ProtobufFieldMaskModel
var ProtobufFieldMaskAttrTypes = map[string]attr.Type{
	"paths": types.ListType{ElemType: types.StringType},
}

// ProtobufFieldMaskResourceSchemaAttributes contains the schema attributes for ProtobufFieldMaskModel
var ProtobufFieldMaskResourceSchemaAttributes = map[string]schema.Attribute{
	"paths": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The set of field mask paths.",
	},
}

// ExpandProtobufFieldMask converts a Terraform Object to SDK type
func ExpandProtobufFieldMask(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddianycast.ProtobufFieldMask {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ProtobufFieldMaskModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ProtobufFieldMaskModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddianycast.ProtobufFieldMask {
	if m == nil {
		return nil
	}
	to := &uddianycast.ProtobufFieldMask{
		Paths: flex.ExpandFrameworkListString(ctx, m.Paths, diags),
	}
	return to
}

// FlattenProtobufFieldMask converts an SDK type to Terraform Object
func FlattenProtobufFieldMask(ctx context.Context, from *uddianycast.ProtobufFieldMask, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ProtobufFieldMaskAttrTypes)
	}
	m := &ProtobufFieldMaskModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ProtobufFieldMaskAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ProtobufFieldMaskModel) Flatten(ctx context.Context, from *uddianycast.ProtobufFieldMask, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Paths = flex.FlattenFrameworkListString(ctx, from.Paths, diags)
}
