package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// ViewSortlistModel is the Terraform model for ViewSortlist
type ViewSortlistModel struct {
	Address   types.String `tfsdk:"address"`
	MatchList types.List   `tfsdk:"match_list"`
}

// ViewSortlistAttrTypes contains the attribute types for ViewSortlistModel
var ViewSortlistAttrTypes = map[string]attr.Type{
	"address":    types.StringType,
	"match_list": types.ListType{ElemType: types.StringType},
}

// ViewSortlistResourceSchemaAttributes contains the schema attributes for ViewSortlistModel
var ViewSortlistResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidIPCIDR(),
		},
		MarkdownDescription: "The source address of a sortlist object.",
	},
	"match_list": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The match list of a sortlist.",
	},
}

// ExpandViewSortlist converts a Terraform Object to SDK type
func ExpandViewSortlist(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.ViewSortlist {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ViewSortlistModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ViewSortlistModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.ViewSortlist {
	if m == nil {
		return nil
	}
	to := &niosdns.ViewSortlist{
		Address:   flex.ExpandStringPointerNullAsEmpty(m.Address),
		MatchList: flex.ExpandFrameworkListString(ctx, m.MatchList, diags),
	}
	return to
}

// FlattenViewSortlist converts an SDK type to Terraform Object
func FlattenViewSortlist(ctx context.Context, from *niosdns.ViewSortlist, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ViewSortlistAttrTypes)
	}
	m := &ViewSortlistModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ViewSortlistAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ViewSortlistModel) Flatten(ctx context.Context, from *niosdns.ViewSortlist, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenStringPointerEmptyAsNull(from.Address)
	m.MatchList = flex.FlattenFrameworkListString(ctx, from.MatchList, diags)
}
