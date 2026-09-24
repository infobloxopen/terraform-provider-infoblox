package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddiclouddiscovery "github.com/infobloxopen/universal-ddi-go-client/clouddiscovery"
)

// CategoryModel is the Terraform model for Category
type CategoryModel struct {
	Excluded types.Bool   `tfsdk:"excluded"`
	Id       types.String `tfsdk:"id"`
}

// CategoryAttrTypes contains the attribute types for CategoryModel
var CategoryAttrTypes = map[string]attr.Type{
	"excluded": types.BoolType,
	"id":       types.StringType,
}

// CategoryResourceSchemaAttributes contains the schema attributes for CategoryModel
var CategoryResourceSchemaAttributes = map[string]schema.Attribute{
	"excluded": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "",
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "",
	},
}

// ExpandCategory converts a Terraform Object to SDK type
func ExpandCategory(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiclouddiscovery.Category {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m CategoryModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *CategoryModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiclouddiscovery.Category {
	if m == nil {
		return nil
	}
	to := &uddiclouddiscovery.Category{
		Excluded: flex.ExpandBoolPointer(m.Excluded),
		Id:       flex.ExpandStringPointer(m.Id),
	}
	return to
}

// FlattenCategory converts an SDK type to Terraform Object
func FlattenCategory(ctx context.Context, from *uddiclouddiscovery.Category, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(CategoryAttrTypes)
	}
	m := &CategoryModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, CategoryAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *CategoryModel) Flatten(ctx context.Context, from *uddiclouddiscovery.Category, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Excluded = flex.FlattenBoolPointer(from.Excluded)
	m.Id = flex.FlattenStringPointer(from.Id)
}
