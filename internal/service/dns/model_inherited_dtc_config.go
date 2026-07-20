package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	uddidns "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// InheritedDtcConfigModel is the Terraform model for InheritedDtcConfig
type InheritedDtcConfigModel struct {
	DefaultTtl types.Object `tfsdk:"default_ttl"`
}

// InheritedDtcConfigAttrTypes contains the attribute types for InheritedDtcConfigModel
var InheritedDtcConfigAttrTypes = map[string]attr.Type{
	"default_ttl": types.ObjectType{AttrTypes: Inheritance2InheritedUInt32DnsconfigAttrTypes},
}

// InheritedDtcConfigResourceSchemaAttributes contains the schema attributes for InheritedDtcConfigModel
var InheritedDtcConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"default_ttl": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedUInt32DnsconfigResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _default_ttl_ field from _DTCConfig_ object.",
	},
}

// ExpandInheritedDtcConfig converts a Terraform Object to SDK type
func ExpandInheritedDtcConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.InheritedDtcConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritedDtcConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *InheritedDtcConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.InheritedDtcConfig {
	if m == nil {
		return nil
	}
	to := &uddidns.InheritedDtcConfig{
		DefaultTtl: ExpandInheritance2InheritedUInt32Dnsconfig(ctx, m.DefaultTtl, diags),
	}
	return to
}

// FlattenInheritedDtcConfig converts an SDK type to Terraform Object
func FlattenInheritedDtcConfig(ctx context.Context, from *uddidns.InheritedDtcConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritedDtcConfigAttrTypes)
	}
	m := &InheritedDtcConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritedDtcConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *InheritedDtcConfigModel) Flatten(ctx context.Context, from *uddidns.InheritedDtcConfig, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.DefaultTtl = FlattenInheritance2InheritedUInt32Dnsconfig(ctx, from.DefaultTtl, diags)
}
