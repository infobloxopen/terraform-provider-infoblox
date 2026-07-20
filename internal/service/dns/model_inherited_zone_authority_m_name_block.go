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

// InheritedZoneAuthorityMNameBlockModel is the Terraform model for InheritedZoneAuthorityMNameBlock
type InheritedZoneAuthorityMNameBlockModel struct {
	Action types.String `tfsdk:"action"`
}

// InheritedZoneAuthorityMNameBlockAttrTypes contains the attribute types for InheritedZoneAuthorityMNameBlockModel
var InheritedZoneAuthorityMNameBlockAttrTypes = map[string]attr.Type{
	"action": types.StringType,
}

// InheritedZoneAuthorityMNameBlockResourceSchemaAttributes contains the schema attributes for InheritedZoneAuthorityMNameBlockModel
var InheritedZoneAuthorityMNameBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Defaults to _inherit_.",
	},
}

// ExpandInheritedZoneAuthorityMNameBlock converts a Terraform Object to SDK type
func ExpandInheritedZoneAuthorityMNameBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.InheritedZoneAuthorityMNameBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritedZoneAuthorityMNameBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *InheritedZoneAuthorityMNameBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.InheritedZoneAuthorityMNameBlock {
	if m == nil {
		return nil
	}
	to := &uddidns.InheritedZoneAuthorityMNameBlock{
		Action: flex.ExpandStringPointer(m.Action),
	}
	return to
}

// FlattenInheritedZoneAuthorityMNameBlock converts an SDK type to Terraform Object
func FlattenInheritedZoneAuthorityMNameBlock(ctx context.Context, from *uddidns.InheritedZoneAuthorityMNameBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritedZoneAuthorityMNameBlockAttrTypes)
	}
	m := &InheritedZoneAuthorityMNameBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritedZoneAuthorityMNameBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *InheritedZoneAuthorityMNameBlockModel) Flatten(ctx context.Context, from *uddidns.InheritedZoneAuthorityMNameBlock, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
}
