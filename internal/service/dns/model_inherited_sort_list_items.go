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

// InheritedSortListItemsModel is the Terraform model for InheritedSortListItems
type InheritedSortListItemsModel struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.List   `tfsdk:"value"`
}

// InheritedSortListItemsAttrTypes contains the attribute types for InheritedSortListItemsModel
var InheritedSortListItemsAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.ListType{ElemType: types.ObjectType{AttrTypes: SortListItemAttrTypes}},
}

// InheritedSortListItemsResourceSchemaAttributes contains the schema attributes for InheritedSortListItemsModel
var InheritedSortListItemsResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Optional. Inheritance setting for a field. Defaults to _inherit_.",
	},
	"display_name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Human-readable display name for the object referred to by _source_.",
	},
	"source": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"value": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: SortListItemResourceSchemaAttributes,
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Inherited value.",
	},
}

// ExpandInheritedSortListItems converts a Terraform Object to SDK type
func ExpandInheritedSortListItems(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.InheritedSortListItems {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritedSortListItemsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *InheritedSortListItemsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.InheritedSortListItems {
	if m == nil {
		return nil
	}
	to := &uddidns.InheritedSortListItems{
		Action:      flex.ExpandStringPointer(m.Action),
		DisplayName: flex.ExpandStringPointer(m.DisplayName),
		Source:      flex.ExpandStringPointer(m.Source),
		Value:       flex.ExpandFrameworkListNestedBlock(ctx, m.Value, diags, ExpandSortListItem),
	}
	return to
}

// FlattenInheritedSortListItems converts an SDK type to Terraform Object
func FlattenInheritedSortListItems(ctx context.Context, from *uddidns.InheritedSortListItems, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritedSortListItemsAttrTypes)
	}
	m := &InheritedSortListItemsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritedSortListItemsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *InheritedSortListItemsModel) Flatten(ctx context.Context, from *uddidns.InheritedSortListItems, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.DisplayName = flex.FlattenStringPointer(from.DisplayName)
	m.Source = flex.FlattenStringPointer(from.Source)
	m.Value = flex.FlattenFrameworkListNestedBlock(ctx, from.Value, SortListItemAttrTypes, diags, FlattenSortListItem)
}
