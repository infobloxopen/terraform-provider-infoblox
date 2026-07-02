package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	uddiipam "github.com/infobloxopen/bloxone-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-unified/internal/flex"
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
		Optional:            true,
		MarkdownDescription: "The name expressed as a single label or FQDN.",
	},
	"type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The origin of the name.",
	},
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

// ExpandListName converts a Terraform List to SDK slice
func ExpandListName(ctx context.Context, l types.List, diags *diag.Diagnostics) []uddiipam.Name {
	if l.IsNull() || l.IsUnknown() {
		return nil
	}
	var models []NameModel
	diags.Append(l.ElementsAs(ctx, &models, false)...)
	if diags.HasError() {
		return nil
	}
	result := make([]uddiipam.Name, 0, len(models))
	for _, m := range models {
		expanded := m.Expand(ctx, diags)
		if expanded != nil {
			result = append(result, *expanded)
		}
	}
	return result
}

// FlattenListName converts an SDK slice to Terraform List
func FlattenListName(ctx context.Context, from []uddiipam.Name, diags *diag.Diagnostics) types.List {
	if len(from) == 0 {
		return types.ListNull(types.ObjectType{AttrTypes: NameAttrTypes})
	}
	var models []NameModel
	for i := range from {
		m := &NameModel{}
		m.Flatten(ctx, &from[i], diags)
		models = append(models, *m)
	}
	listVal, d := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: NameAttrTypes}, models)
	diags.Append(d...)
	return listVal
}

// Flatten populates the Terraform model from SDK type
func (m *NameModel) Flatten(ctx context.Context, from *uddiipam.Name, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Name = flex.FlattenString(from.Name)
	m.Type = flex.FlattenString(from.Type)
}
