package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	uddiipam "github.com/infobloxopen/bloxone-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// NameModel is the Terraform model for Name
type NameModel struct {
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

// NameAttrTypes contains the attribute types for NameModel
var NameAttrTypes = map[string]attr.Type{
	"name": types.StringType,
	"type": types.StringType,
}

// NameResourceSchemaAttributes contains the schema attributes for NameModel
var NameResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name expressed as a single label or FQDN.",
	},
	"type": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The origin of the name.",
	},
}

// ExpandName converts a Terraform Object to SDK type
func ExpandName(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiipam.Name {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NameModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NameModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiipam.Name {
	if m == nil {
		return nil
	}
	to := &uddiipam.Name{
		Name: flex.ExpandString(m.Name),
		Type: flex.ExpandString(m.Type),
	}
	return to
}

// FlattenName converts an SDK type to Terraform Object
func FlattenName(ctx context.Context, from *uddiipam.Name, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NameAttrTypes)
	}
	m := &NameModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NameAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NameModel) Flatten(ctx context.Context, from *uddiipam.Name, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Name = flex.FlattenString(from.Name)
	m.Type = flex.FlattenString(from.Type)
}
