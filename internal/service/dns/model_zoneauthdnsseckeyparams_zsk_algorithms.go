package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ZoneauthdnsseckeyparamsZskAlgorithmsModel is the Terraform model for ZoneauthdnsseckeyparamsZskAlgorithms
type ZoneauthdnsseckeyparamsZskAlgorithmsModel struct {
	Algorithm types.String `tfsdk:"algorithm"`
	Size      types.Int64  `tfsdk:"size"`
}

// ZoneauthdnsseckeyparamsZskAlgorithmsAttrTypes contains the attribute types for ZoneauthdnsseckeyparamsZskAlgorithmsModel
var ZoneauthdnsseckeyparamsZskAlgorithmsAttrTypes = map[string]attr.Type{
	"algorithm": types.StringType,
	"size":      types.Int64Type,
}

// ZoneauthdnsseckeyparamsZskAlgorithmsResourceSchemaAttributes contains the schema attributes for ZoneauthdnsseckeyparamsZskAlgorithmsModel
var ZoneauthdnsseckeyparamsZskAlgorithmsResourceSchemaAttributes = map[string]schema.Attribute{
	"algorithm": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("RSASHA1", "RSASHA256", "RSASHA512", "ECDSAP256SHA256", "ECDSAP384SHA384"),
		},
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("RSASHA256"),
		MarkdownDescription: "The signing key algorithm.",
	},
	"size": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1024),
		MarkdownDescription: "The signing key size, in bits.",
	},
}

// ExpandZoneauthdnsseckeyparamsZskAlgorithms converts a Terraform Object to SDK type
func ExpandZoneauthdnsseckeyparamsZskAlgorithms(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.ZoneauthdnsseckeyparamsZskAlgorithms {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneauthdnsseckeyparamsZskAlgorithmsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ZoneauthdnsseckeyparamsZskAlgorithmsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.ZoneauthdnsseckeyparamsZskAlgorithms {
	if m == nil {
		return nil
	}
	to := &niosdns.ZoneauthdnsseckeyparamsZskAlgorithms{
		Algorithm: flex.ExpandStringPointer(m.Algorithm),
		Size:      flex.ExpandInt64Pointer(m.Size),
	}
	return to
}

// FlattenZoneauthdnsseckeyparamsZskAlgorithms converts an SDK type to Terraform Object
func FlattenZoneauthdnsseckeyparamsZskAlgorithms(ctx context.Context, from *niosdns.ZoneauthdnsseckeyparamsZskAlgorithms, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneauthdnsseckeyparamsZskAlgorithmsAttrTypes)
	}
	m := &ZoneauthdnsseckeyparamsZskAlgorithmsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneauthdnsseckeyparamsZskAlgorithmsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ZoneauthdnsseckeyparamsZskAlgorithmsModel) Flatten(ctx context.Context, from *niosdns.ZoneauthdnsseckeyparamsZskAlgorithms, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Algorithm = flex.FlattenStringPointerEmptyAsNull(from.Algorithm)
	m.Size = flex.FlattenInt64Pointer(from.Size)
}
