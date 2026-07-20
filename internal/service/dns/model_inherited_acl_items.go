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

// InheritedACLItemsModel is the Terraform model for InheritedACLItems
type InheritedACLItemsModel struct {
	Action types.String `tfsdk:"action"`
}

// InheritedACLItemsAttrTypes contains the attribute types for InheritedACLItemsModel
var InheritedACLItemsAttrTypes = map[string]attr.Type{
	"action": types.StringType,
}

// InheritedACLItemsResourceSchemaAttributes contains the schema attributes for InheritedACLItemsModel
var InheritedACLItemsResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Inheritance setting for a field. Defaults to _inherit_.",
	},
}

// ExpandInheritedACLItems converts a Terraform Object to SDK type
func ExpandInheritedACLItems(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.InheritedACLItems {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritedACLItemsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *InheritedACLItemsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.InheritedACLItems {
	if m == nil {
		return nil
	}
	to := &uddidns.InheritedACLItems{
		Action: flex.ExpandStringPointer(m.Action),
	}
	return to
}

// FlattenInheritedACLItems converts an SDK type to Terraform Object
func FlattenInheritedACLItems(ctx context.Context, from *uddidns.InheritedACLItems, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritedACLItemsAttrTypes)
	}
	m := &InheritedACLItemsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritedACLItemsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *InheritedACLItemsModel) Flatten(ctx context.Context, from *uddidns.InheritedACLItems, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
}
