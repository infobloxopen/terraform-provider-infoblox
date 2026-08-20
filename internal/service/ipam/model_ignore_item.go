package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// IgnoreItemModel is the Terraform model for IgnoreItem
type IgnoreItemModel struct {
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

// IgnoreItemAttrTypes contains the attribute types for IgnoreItemModel
var IgnoreItemAttrTypes = map[string]attr.Type{
	"type":  types.StringType,
	"value": types.StringType,
}

// IgnoreItemResourceSchemaAttributes contains the schema attributes for IgnoreItemModel
var IgnoreItemResourceSchemaAttributes = map[string]schema.Attribute{
	"type": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Type of ignore matching: client to ignore by client identifier (client hex or client text) or hardware to ignore by hardware identifier (MAC address). It can have one of the following values:  * _client_hex_,  * _client_text_,  * _hardware_.",
	},
	"value": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Value to match.",
	},
}

// ExpandIgnoreItem converts a Terraform Object to SDK type
func ExpandIgnoreItem(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiipam.IgnoreItem {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IgnoreItemModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *IgnoreItemModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiipam.IgnoreItem {
	if m == nil {
		return nil
	}
	to := &uddiipam.IgnoreItem{
		Type:  flex.ExpandString(m.Type),
		Value: flex.ExpandString(m.Value),
	}
	return to
}

// FlattenIgnoreItem converts an SDK type to Terraform Object
func FlattenIgnoreItem(ctx context.Context, from *uddiipam.IgnoreItem, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IgnoreItemAttrTypes)
	}
	m := &IgnoreItemModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IgnoreItemAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *IgnoreItemModel) Flatten(ctx context.Context, from *uddiipam.IgnoreItem, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Type = flex.FlattenString(from.Type)
	m.Value = flex.FlattenString(from.Value)
}
