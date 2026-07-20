package dns

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
	uddidns "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// TSIGKeyModel is the Terraform model for TSIGKey
type TSIGKeyModel struct {
	Algorithm    types.String `tfsdk:"algorithm"`
	Comment      types.String `tfsdk:"comment"`
	Key          types.String `tfsdk:"key"`
	Name         types.String `tfsdk:"name"`
	ProtocolName types.String `tfsdk:"protocol_name"`
	Secret       types.String `tfsdk:"secret"`
}

// TSIGKeyAttrTypes contains the attribute types for TSIGKeyModel
var TSIGKeyAttrTypes = map[string]attr.Type{
	"algorithm":     types.StringType,
	"comment":       types.StringType,
	"key":           types.StringType,
	"name":          types.StringType,
	"protocol_name": types.StringType,
	"secret":        types.StringType,
}

// TSIGKeyResourceSchemaAttributes contains the schema attributes for TSIGKeyModel
var TSIGKeyResourceSchemaAttributes = map[string]schema.Attribute{
	"algorithm": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "TSIG key algorithm.  Possible values:  * _hmac_sha256_,  * _hmac_sha1_,  * _hmac_sha224_,  * _hmac_sha384_,  * _hmac_sha512_.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Comment for TSIG key.",
	},
	"key": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotNull(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "TSIG key name, FQDN.",
	},
	"protocol_name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "TSIG key name in punycode.",
	},
	"secret": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "TSIG key secret, base64 string.",
	},
}

// ExpandTSIGKey converts a Terraform Object to SDK type
func ExpandTSIGKey(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.TSIGKey {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m TSIGKeyModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *TSIGKeyModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.TSIGKey {
	if m == nil {
		return nil
	}
	to := &uddidns.TSIGKey{
		Algorithm:    flex.ExpandStringPointer(m.Algorithm),
		Comment:      flex.ExpandStringPointer(m.Comment),
		Key:          flex.ExpandStringPointer(m.Key),
		Name:         flex.ExpandStringPointer(m.Name),
		ProtocolName: flex.ExpandStringPointer(m.ProtocolName),
		Secret:       flex.ExpandStringPointer(m.Secret),
	}
	return to
}

// FlattenTSIGKey converts an SDK type to Terraform Object
func FlattenTSIGKey(ctx context.Context, from *uddidns.TSIGKey, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(TSIGKeyAttrTypes)
	}
	m := &TSIGKeyModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, TSIGKeyAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *TSIGKeyModel) Flatten(ctx context.Context, from *uddidns.TSIGKey, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Algorithm = flex.FlattenStringPointer(from.Algorithm)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Key = flex.FlattenStringPointer(from.Key)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.ProtocolName = flex.FlattenStringPointer(from.ProtocolName)
	m.Secret = flex.FlattenStringPointer(from.Secret)
}
