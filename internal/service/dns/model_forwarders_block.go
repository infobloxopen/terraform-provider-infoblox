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

// ForwardersBlockModel is the Terraform model for ForwardersBlock
type ForwardersBlockModel struct {
	Forwarders                                  types.List `tfsdk:"forwarders"`
	ForwardersOnly                              types.Bool `tfsdk:"forwarders_only"`
	UseRootForwardersForLocalResolutionWithB1td types.Bool `tfsdk:"use_root_forwarders_for_local_resolution_with_b1td"`
}

// ForwardersBlockAttrTypes contains the attribute types for ForwardersBlockModel
var ForwardersBlockAttrTypes = map[string]attr.Type{
	"forwarders":      types.ListType{ElemType: types.ObjectType{AttrTypes: ForwarderAttrTypes}},
	"forwarders_only": types.BoolType,
	"use_root_forwarders_for_local_resolution_with_b1td": types.BoolType,
}

// ForwardersBlockResourceSchemaAttributes contains the schema attributes for ForwardersBlockModel
var ForwardersBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"forwarders": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ForwarderResourceSchemaAttributes(false),
		},
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _forwarders_ field from.",
	},
	"forwarders_only": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _forwarders_only_ field.",
	},
	"use_root_forwarders_for_local_resolution_with_b1td": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _use_root_forwarders_for_local_resolution_with_b1td_ field.",
	},
}

// ExpandForwardersBlock converts a Terraform Object to SDK type
func ExpandForwardersBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.ForwardersBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ForwardersBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ForwardersBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.ForwardersBlock {
	if m == nil {
		return nil
	}
	to := &uddidns.ForwardersBlock{
		Forwarders:     flex.ExpandFrameworkListNestedBlock(ctx, m.Forwarders, diags, ExpandForwarder),
		ForwardersOnly: flex.ExpandBoolPointer(m.ForwardersOnly),
		UseRootForwardersForLocalResolutionWithB1td: flex.ExpandBoolPointer(m.UseRootForwardersForLocalResolutionWithB1td),
	}
	return to
}

// FlattenForwardersBlock converts an SDK type to Terraform Object
func FlattenForwardersBlock(ctx context.Context, from *uddidns.ForwardersBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ForwardersBlockAttrTypes)
	}
	m := &ForwardersBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ForwardersBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ForwardersBlockModel) Flatten(ctx context.Context, from *uddidns.ForwardersBlock, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Forwarders = flex.FlattenFrameworkListNestedBlock(ctx, from.Forwarders, ForwarderAttrTypes, diags, FlattenForwarder)
	m.ForwardersOnly = flex.FlattenBoolPointer(from.ForwardersOnly)
	m.UseRootForwardersForLocalResolutionWithB1td = flex.FlattenBoolPointer(from.UseRootForwardersForLocalResolutionWithB1td)
}
