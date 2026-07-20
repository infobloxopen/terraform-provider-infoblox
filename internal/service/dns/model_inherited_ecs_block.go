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

// InheritedECSBlockModel is the Terraform model for InheritedECSBlock
type InheritedECSBlockModel struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.Object `tfsdk:"value"`
}

// InheritedECSBlockAttrTypes contains the attribute types for InheritedECSBlockModel
var InheritedECSBlockAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.ObjectType{AttrTypes: ECSBlockAttrTypes},
}

// InheritedECSBlockResourceSchemaAttributes contains the schema attributes for InheritedECSBlockModel
var InheritedECSBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Defaults to _inherit_.",
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
	"value": schema.SingleNestedAttribute{
		Attributes:          ECSBlockResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Inherited value.",
	},
}

// ExpandInheritedECSBlock converts a Terraform Object to SDK type
func ExpandInheritedECSBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.InheritedECSBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritedECSBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *InheritedECSBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.InheritedECSBlock {
	if m == nil {
		return nil
	}
	to := &uddidns.InheritedECSBlock{
		Action:      flex.ExpandStringPointer(m.Action),
		DisplayName: flex.ExpandStringPointer(m.DisplayName),
		Source:      flex.ExpandStringPointer(m.Source),
		Value:       ExpandECSBlock(ctx, m.Value, diags),
	}
	return to
}

// FlattenInheritedECSBlock converts an SDK type to Terraform Object
func FlattenInheritedECSBlock(ctx context.Context, from *uddidns.InheritedECSBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritedECSBlockAttrTypes)
	}
	m := &InheritedECSBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritedECSBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *InheritedECSBlockModel) Flatten(ctx context.Context, from *uddidns.InheritedECSBlock, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.DisplayName = flex.FlattenStringPointer(from.DisplayName)
	m.Source = flex.FlattenStringPointer(from.Source)
	m.Value = FlattenECSBlock(ctx, from.Value, diags)
}
