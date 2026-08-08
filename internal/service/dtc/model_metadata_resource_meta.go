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

// MetadataResourceMetaModel is the Terraform model for MetadataResourceMeta
type MetadataResourceMetaModel struct {
	DisplayName types.String `tfsdk:"display_name"`
	Id          types.String `tfsdk:"id"`
}

// MetadataResourceMetaAttrTypes contains the attribute types for MetadataResourceMetaModel
var MetadataResourceMetaAttrTypes = map[string]attr.Type{
	"display_name": types.StringType,
	"id":           types.StringType,
}

// MetadataResourceMetaResourceSchemaAttributes contains the schema attributes for MetadataResourceMetaModel
var MetadataResourceMetaResourceSchemaAttributes = map[string]schema.Attribute{
	"display_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Display name of the configuration resource.",
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
}

// ExpandMetadataResourceMeta converts a Terraform Object to SDK type
func ExpandMetadataResourceMeta(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidtc.MetadataResourceMeta {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MetadataResourceMetaModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *MetadataResourceMetaModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidtc.MetadataResourceMeta {
	if m == nil {
		return nil
	}
	to := &uddidtc.MetadataResourceMeta{
		DisplayName: flex.ExpandStringPointer(m.DisplayName),
		Id:          flex.ExpandStringPointer(m.Id),
	}
	return to
}

// FlattenMetadataResourceMeta converts an SDK type to Terraform Object
func FlattenMetadataResourceMeta(ctx context.Context, from *uddidtc.MetadataResourceMeta, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MetadataResourceMetaAttrTypes)
	}
	m := &MetadataResourceMetaModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MetadataResourceMetaAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *MetadataResourceMetaModel) Flatten(ctx context.Context, from *uddidtc.MetadataResourceMeta, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.DisplayName = flex.FlattenStringPointer(from.DisplayName)
	m.Id = flex.FlattenStringPointer(from.Id)
}
