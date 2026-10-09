package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// ZoneAuthLastQueriedAclModel is the Terraform model for ZoneAuthLastQueriedAcl
type ZoneAuthLastQueriedAclModel struct {
	Address    types.String `tfsdk:"address"`
	Permission types.String `tfsdk:"permission"`
}

// ZoneAuthLastQueriedAclAttrTypes contains the attribute types for ZoneAuthLastQueriedAclModel
var ZoneAuthLastQueriedAclAttrTypes = map[string]attr.Type{
	"address":    types.StringType,
	"permission": types.StringType,
}

// ZoneAuthLastQueriedAclResourceSchemaAttributes contains the schema attributes for ZoneAuthLastQueriedAclModel
var ZoneAuthLastQueriedAclResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The address this rule applies to or \"Any\".",
	},
	"permission": schema.StringAttribute{
		Default: stringdefault.StaticString("ALLOW"),
		Validators: []validator.String{
			stringvalidator.OneOf("ALLOW", "DENY"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The permission to use for this address.",
	},
}

// ExpandZoneAuthLastQueriedAcl converts a Terraform Object to SDK type
func ExpandZoneAuthLastQueriedAcl(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.ZoneAuthLastQueriedAcl {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneAuthLastQueriedAclModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ZoneAuthLastQueriedAclModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.ZoneAuthLastQueriedAcl {
	if m == nil {
		return nil
	}
	to := &niosdns.ZoneAuthLastQueriedAcl{
		Address:    flex.ExpandStringPointerNullAsEmpty(m.Address),
		Permission: flex.ExpandStringPointerNullAsEmpty(m.Permission),
	}
	return to
}

// FlattenZoneAuthLastQueriedAcl converts an SDK type to Terraform Object
func FlattenZoneAuthLastQueriedAcl(ctx context.Context, from *niosdns.ZoneAuthLastQueriedAcl, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneAuthLastQueriedAclAttrTypes)
	}
	m := &ZoneAuthLastQueriedAclModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneAuthLastQueriedAclAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ZoneAuthLastQueriedAclModel) Flatten(ctx context.Context, from *niosdns.ZoneAuthLastQueriedAcl, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenStringPointerEmptyAsNull(from.Address)
	m.Permission = flex.FlattenStringPointerEmptyAsNull(from.Permission)
}
