package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// InheritedDDNSUpdateBlockModel is the Terraform model for InheritedDDNSUpdateBlock
type InheritedDDNSUpdateBlockModel struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.Object `tfsdk:"value"`
}

// InheritedDDNSUpdateBlockAttrTypes contains the attribute types for InheritedDDNSUpdateBlockModel
var InheritedDDNSUpdateBlockAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.ObjectType{AttrTypes: DDNSUpdateBlockAttrTypes},
}

// InheritedDDNSUpdateBlockResourceSchemaAttributes contains the schema attributes for InheritedDDNSUpdateBlockModel
var InheritedDDNSUpdateBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The inheritance setting.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.",
	},
	"display_name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The human-readable display name for the object referred to by _source_.",
	},
	"source": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"value": schema.SingleNestedAttribute{
		Attributes:          DDNSUpdateBlockResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inherited value.",
	},
}

// ExpandInheritedDDNSUpdateBlock converts a Terraform Object to SDK type
func ExpandInheritedDDNSUpdateBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiipam.InheritedDDNSUpdateBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritedDDNSUpdateBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *InheritedDDNSUpdateBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiipam.InheritedDDNSUpdateBlock {
	if m == nil {
		return nil
	}
	to := &uddiipam.InheritedDDNSUpdateBlock{
		Action:      flex.ExpandStringPointer(m.Action),
		DisplayName: flex.ExpandStringPointer(m.DisplayName),
		Source:      flex.ExpandStringPointer(m.Source),
		Value:       ExpandDDNSUpdateBlock(ctx, m.Value, diags),
	}
	return to
}

// FlattenInheritedDDNSUpdateBlock converts an SDK type to Terraform Object
func FlattenInheritedDDNSUpdateBlock(ctx context.Context, from *uddiipam.InheritedDDNSUpdateBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritedDDNSUpdateBlockAttrTypes)
	}
	m := &InheritedDDNSUpdateBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritedDDNSUpdateBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *InheritedDDNSUpdateBlockModel) Flatten(ctx context.Context, from *uddiipam.InheritedDDNSUpdateBlock, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.DisplayName = flex.FlattenStringPointer(from.DisplayName)
	m.Source = flex.FlattenStringPointer(from.Source)
	m.Value = FlattenDDNSUpdateBlock(ctx, from.Value, diags)
}
