package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddifw "github.com/infobloxopen/universal-ddi-go-client/fw"
)

// ItemStructsModel is the Terraform model for ItemStructs
type ItemStructsModel struct {
	Description types.String `tfsdk:"description"`
	Item        types.String `tfsdk:"item"`
}

// ItemStructsAttrTypes contains the attribute types for ItemStructsModel
var ItemStructsAttrTypes = map[string]attr.Type{
	"description": types.StringType,
	"item":        types.StringType,
}

// ItemStructsResourceSchemaAttributes contains the schema attributes for ItemStructsModel
var ItemStructsResourceSchemaAttributes = map[string]schema.Attribute{
	"description": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The description of the item",
	},
	"item": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The data of the Item",
	},
}

// ExpandItemStructs converts a Terraform Object to SDK type
func ExpandItemStructs(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddifw.ItemStructs {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ItemStructsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ItemStructsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddifw.ItemStructs {
	if m == nil {
		return nil
	}
	to := &uddifw.ItemStructs{
		Description: flex.ExpandStringPointer(m.Description),
		Item:        flex.ExpandStringPointer(m.Item),
	}
	return to
}

// FlattenItemStructs converts an SDK type to Terraform Object
func FlattenItemStructs(ctx context.Context, from *uddifw.ItemStructs, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ItemStructsAttrTypes)
	}
	m := &ItemStructsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ItemStructsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ItemStructsModel) Flatten(ctx context.Context, from *uddifw.ItemStructs, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Item = flex.FlattenStringPointer(from.Item)
}
