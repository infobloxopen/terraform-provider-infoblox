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
	Action types.String `tfsdk:"action"`
}

// InheritedSortListItemsAttrTypes contains the attribute types for InheritedSortListItemsModel
var InheritedSortListItemsAttrTypes = map[string]attr.Type{
	"action": types.StringType,
}

// InheritedSortListItemsResourceSchemaAttributes contains the schema attributes for InheritedSortListItemsModel
var InheritedSortListItemsResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Inheritance setting for a field. Defaults to _inherit_.",
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
		Action: flex.ExpandStringPointer(m.Action),
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
}
