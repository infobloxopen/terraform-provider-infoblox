package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddidtc "github.com/infobloxopen/universal-ddi-go-client/dtc"
)

// HeaderRegexModel is the Terraform model for HeaderRegex
type HeaderRegexModel struct {
	Header types.String `tfsdk:"header"`
	Regex  types.String `tfsdk:"regex"`
}

// HeaderRegexAttrTypes contains the attribute types for HeaderRegexModel
var HeaderRegexAttrTypes = map[string]attr.Type{
	"header": types.StringType,
	"regex":  types.StringType,
}

// HeaderRegexResourceSchemaAttributes contains the schema attributes for HeaderRegexModel
var HeaderRegexResourceSchemaAttributes = map[string]schema.Attribute{
	"header": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "HTTP header name.",
	},
	"regex": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Regular expression to match against HTTP header value.",
	},
}

// ExpandHeaderRegex converts a Terraform Object to SDK type
func ExpandHeaderRegex(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidtc.HeaderRegex {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m HeaderRegexModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *HeaderRegexModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidtc.HeaderRegex {
	if m == nil {
		return nil
	}
	to := &uddidtc.HeaderRegex{
		Header: flex.ExpandString(m.Header),
		Regex:  flex.ExpandString(m.Regex),
	}
	return to
}

// FlattenHeaderRegex converts an SDK type to Terraform Object
func FlattenHeaderRegex(ctx context.Context, from *uddidtc.HeaderRegex, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(HeaderRegexAttrTypes)
	}
	m := &HeaderRegexModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, HeaderRegexAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *HeaderRegexModel) Flatten(ctx context.Context, from *uddidtc.HeaderRegex, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Header = flex.FlattenString(from.Header)
	m.Regex = flex.FlattenString(from.Regex)
}
