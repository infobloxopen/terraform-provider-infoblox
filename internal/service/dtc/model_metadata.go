package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddidtc "github.com/infobloxopen/universal-ddi-go-client/dtc"
)

// MetadataModel is the Terraform model for Metadata
type MetadataModel struct {
	UsedBy types.List `tfsdk:"used_by"`
}

// MetadataAttrTypes contains the attribute types for MetadataModel
var MetadataAttrTypes = map[string]attr.Type{
	"used_by": types.ListType{ElemType: types.ObjectType{AttrTypes: MetadataResourceMetaAttrTypes}},
}

// MetadataResourceSchemaAttributes contains the schema attributes for MetadataModel
var MetadataResourceSchemaAttributes = map[string]schema.Attribute{
	"used_by": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MetadataResourceMetaResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "List of structs representing a limited view on configuration objects that use a resource the metadata is provided for.",
	},
}

// ExpandMetadata converts a Terraform Object to SDK type
func ExpandMetadata(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidtc.Metadata {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MetadataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *MetadataModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidtc.Metadata {
	if m == nil {
		return nil
	}
	to := &uddidtc.Metadata{
		UsedBy: flex.ExpandFrameworkListNestedBlock(ctx, m.UsedBy, diags, ExpandMetadataResourceMeta),
	}
	return to
}

// FlattenMetadata converts an SDK type to Terraform Object
func FlattenMetadata(ctx context.Context, from *uddidtc.Metadata, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MetadataAttrTypes)
	}
	m := &MetadataModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MetadataAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *MetadataModel) Flatten(ctx context.Context, from *uddidtc.Metadata, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.UsedBy = flex.FlattenFrameworkListNestedBlock(ctx, from.UsedBy, MetadataResourceMetaAttrTypes, diags, FlattenMetadataResourceMeta)
}
