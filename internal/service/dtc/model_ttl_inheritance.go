package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	uddidtc "github.com/infobloxopen/universal-ddi-go-client/dtc"
)

// TTLInheritanceModel is the Terraform model for TTLInheritance
type TTLInheritanceModel struct {
	Ttl types.Object `tfsdk:"ttl"`
}

// TTLInheritanceAttrTypes contains the attribute types for TTLInheritanceModel
var TTLInheritanceAttrTypes = map[string]attr.Type{
	"ttl": types.ObjectType{AttrTypes: Inheritance2InheritedUInt32AttrTypes},
}

// TTLInheritanceResourceSchemaAttributes contains the schema attributes for TTLInheritanceModel
var TTLInheritanceResourceSchemaAttributes = map[string]schema.Attribute{
	"ttl": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedUInt32ResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
}

// ExpandTTLInheritance converts a Terraform Object to SDK type
func ExpandTTLInheritance(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidtc.TTLInheritance {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m TTLInheritanceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *TTLInheritanceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidtc.TTLInheritance {
	if m == nil {
		return nil
	}
	to := &uddidtc.TTLInheritance{
		Ttl: ExpandInheritance2InheritedUInt32(ctx, m.Ttl, diags),
	}
	return to
}

// FlattenTTLInheritance converts an SDK type to Terraform Object
func FlattenTTLInheritance(ctx context.Context, from *uddidtc.TTLInheritance, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(TTLInheritanceAttrTypes)
	}
	m := &TTLInheritanceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, TTLInheritanceAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *TTLInheritanceModel) Flatten(ctx context.Context, from *uddidtc.TTLInheritance, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Ttl = FlattenInheritance2InheritedUInt32(ctx, from.Ttl, diags)
}
