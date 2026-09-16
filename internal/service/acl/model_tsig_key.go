package acl

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
	uddiacl "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// TSIGKeyModel is the Terraform model for TSIGKey
type TSIGKeyModel struct {
	Key types.String `tfsdk:"key"`
}

// TSIGKeyAttrTypes contains the attribute types for TSIGKeyModel
var TSIGKeyAttrTypes = map[string]attr.Type{
	"key": types.StringType,
}

// TSIGKeyResourceSchemaAttributes contains the schema attributes for TSIGKeyModel
var TSIGKeyResourceSchemaAttributes = map[string]schema.Attribute{
	"key": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotNull(),
		},
		MarkdownDescription: "The resource identifier.",
	},
}

// ExpandTSIGKey converts a Terraform Object to SDK type
func ExpandTSIGKey(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiacl.TSIGKey {
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
func (m *TSIGKeyModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiacl.TSIGKey {
	if m == nil {
		return nil
	}
	to := &uddiacl.TSIGKey{
		Key: flex.ExpandStringPointer(m.Key),
	}
	return to
}

// FlattenTSIGKey converts an SDK type to Terraform Object
func FlattenTSIGKey(ctx context.Context, from *uddiacl.TSIGKey, diags *diag.Diagnostics) types.Object {
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
func (m *TSIGKeyModel) Flatten(ctx context.Context, from *uddiacl.TSIGKey, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Key = flex.FlattenStringPointer(from.Key)
}
