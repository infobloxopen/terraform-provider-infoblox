package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	uddidtc "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// TTLInheritanceDnsconfigModel is the Terraform model for TTLInheritance
type TTLInheritanceDnsconfigModel struct {
	Ttl types.Object `tfsdk:"ttl"`
}

// TTLInheritanceDnsconfigAttrTypes contains the attribute types for TTLInheritanceDnsconfigModel
var TTLInheritanceDnsconfigAttrTypes = map[string]attr.Type{
	"ttl": types.ObjectType{AttrTypes: Inheritance2InheritedUInt32DnsconfigAttrTypes},
}

// TTLInheritanceDnsconfigResourceSchemaAttributes contains the schema attributes for TTLInheritanceDnsconfigModel
var TTLInheritanceDnsconfigResourceSchemaAttributes = map[string]schema.Attribute{
	"ttl": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedUInt32DnsconfigResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
}

// ExpandTTLInheritanceDnsconfig converts a Terraform Object to SDK type
func ExpandTTLInheritanceDnsconfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidtc.TTLInheritance {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m TTLInheritanceDnsconfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *TTLInheritanceDnsconfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidtc.TTLInheritance {
	if m == nil {
		return nil
	}
	to := &uddidtc.TTLInheritance{
		Ttl: ExpandInheritance2InheritedUInt32Dnsconfig(ctx, m.Ttl, diags),
	}
	return to
}

// FlattenTTLInheritanceDnsconfig converts an SDK type to Terraform Object
func FlattenTTLInheritanceDnsconfig(ctx context.Context, from *uddidtc.TTLInheritance, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(TTLInheritanceDnsconfigAttrTypes)
	}
	m := &TTLInheritanceDnsconfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, TTLInheritanceDnsconfigAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *TTLInheritanceDnsconfigModel) Flatten(ctx context.Context, from *uddidtc.TTLInheritance, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Ttl = FlattenInheritance2InheritedUInt32Dnsconfig(ctx, from.Ttl, diags)
}
