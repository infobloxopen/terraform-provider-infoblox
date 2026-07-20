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

// DTCConfigModel is the Terraform model for DTCConfig
type DTCConfigModel struct {
	DefaultTtl types.Int64 `tfsdk:"default_ttl"`
}

// DTCConfigAttrTypes contains the attribute types for DTCConfigModel
var DTCConfigAttrTypes = map[string]attr.Type{
	"default_ttl": types.Int64Type,
}

// DTCConfigResourceSchemaAttributes contains the schema attributes for DTCConfigModel
var DTCConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"default_ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Optional. Default TTL for synthesized DTC records (value in seconds).  Defaults to 300.",
	},
}

// ExpandDTCConfig converts a Terraform Object to SDK type
func ExpandDTCConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.DTCConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DTCConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *DTCConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.DTCConfig {
	if m == nil {
		return nil
	}
	to := &uddidns.DTCConfig{
		DefaultTtl: flex.ExpandInt64Pointer(m.DefaultTtl),
	}
	return to
}

// FlattenDTCConfig converts an SDK type to Terraform Object
func FlattenDTCConfig(ctx context.Context, from *uddidns.DTCConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DTCConfigAttrTypes)
	}
	m := &DTCConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DTCConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *DTCConfigModel) Flatten(ctx context.Context, from *uddidns.DTCConfig, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.DefaultTtl = flex.FlattenInt64Pointer(from.DefaultTtl)
}
