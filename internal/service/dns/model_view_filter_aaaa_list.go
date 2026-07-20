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

	"github.com/hashicorp/terraform-plugin-framework/path"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// ViewFilterAaaaListModel is the Terraform model for ViewFilterAaaaList
type ViewFilterAaaaListModel struct {
	Ref        types.String `tfsdk:"ref"`
	Address    types.String `tfsdk:"address"`
	Permission types.String `tfsdk:"permission"`
}

// ViewFilterAaaaListAttrTypes contains the attribute types for ViewFilterAaaaListModel
var ViewFilterAaaaListAttrTypes = map[string]attr.Type{
	"ref":        types.StringType,
	"address":    types.StringType,
	"permission": types.StringType,
}

// ViewFilterAaaaListResourceSchemaAttributes contains the schema attributes for ViewFilterAaaaListModel
var ViewFilterAaaaListResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("address")),
			stringvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("permission")),
		},
		MarkdownDescription: "The reference to the object.",
	},
	"address": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The address this rule applies to or \"Any\".",
	},
	"permission": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("ALLOW", "DENY"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The permission to use for this address.",
	},
}

// ExpandViewFilterAaaaList converts a Terraform Object to SDK type
func ExpandViewFilterAaaaList(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.ViewFilterAaaaList {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ViewFilterAaaaListModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ViewFilterAaaaListModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.ViewFilterAaaaList {
	if m == nil {
		return nil
	}
	to := &niosdns.ViewFilterAaaaList{
		Ref:        flex.ExpandStringPointerNullAsEmpty(m.Ref),
		Address:    flex.ExpandStringPointerNullAsEmpty(m.Address),
		Permission: flex.ExpandStringPointerNullAsEmpty(m.Permission),
	}
	return to
}

// FlattenViewFilterAaaaList converts an SDK type to Terraform Object
func FlattenViewFilterAaaaList(ctx context.Context, from *niosdns.ViewFilterAaaaList, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ViewFilterAaaaListAttrTypes)
	}
	m := &ViewFilterAaaaListModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ViewFilterAaaaListAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ViewFilterAaaaListModel) Flatten(ctx context.Context, from *niosdns.ViewFilterAaaaList, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Ref = flex.FlattenStringPointerEmptyAsNull(from.Ref)
	m.Address = flex.FlattenStringPointerEmptyAsNull(from.Address)
	m.Permission = flex.FlattenStringPointerEmptyAsNull(from.Permission)
}
