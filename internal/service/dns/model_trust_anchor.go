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

// TrustAnchorModel is the Terraform model for TrustAnchor
type TrustAnchorModel struct {
	Algorithm types.Int64  `tfsdk:"algorithm"`
	PublicKey types.String `tfsdk:"public_key"`
	Sep       types.Bool   `tfsdk:"sep"`
	Zone      types.String `tfsdk:"zone"`
}

// TrustAnchorAttrTypes contains the attribute types for TrustAnchorModel
var TrustAnchorAttrTypes = map[string]attr.Type{
	"algorithm":  types.Int64Type,
	"public_key": types.StringType,
	"sep":        types.BoolType,
	"zone":       types.StringType,
}

// TrustAnchorResourceSchemaAttributes contains the schema attributes for TrustAnchorModel
var TrustAnchorResourceSchemaAttributes = map[string]schema.Attribute{
	"algorithm": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "",
	},
	"public_key": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "DNSSEC key data. Non-empty, valid base64 string.",
	},
	"sep": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Optional. Secure Entry Point flag.  Defaults to _true_.",
	},
	"zone": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Zone FQDN.",
	},
}

// ExpandTrustAnchor converts a Terraform Object to SDK type
func ExpandTrustAnchor(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.TrustAnchor {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m TrustAnchorModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *TrustAnchorModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.TrustAnchor {
	if m == nil {
		return nil
	}
	to := &uddidns.TrustAnchor{
		Algorithm: flex.ExpandInt64(m.Algorithm),
		PublicKey: flex.ExpandString(m.PublicKey),
		Sep:       flex.ExpandBoolPointer(m.Sep),
		Zone:      flex.ExpandString(m.Zone),
	}
	return to
}

// FlattenTrustAnchor converts an SDK type to Terraform Object
func FlattenTrustAnchor(ctx context.Context, from *uddidns.TrustAnchor, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(TrustAnchorAttrTypes)
	}
	m := &TrustAnchorModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, TrustAnchorAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *TrustAnchorModel) Flatten(ctx context.Context, from *uddidns.TrustAnchor, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Algorithm = flex.FlattenInt64(from.Algorithm)
	m.PublicKey = flex.FlattenString(from.PublicKey)
	m.Sep = flex.FlattenBoolPointer(from.Sep)
	m.Zone = flex.FlattenString(from.Zone)
}
