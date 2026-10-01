package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddiclouddiscovery "github.com/infobloxopen/universal-ddi-go-client/clouddiscovery"
)

// ResourceModel is the Terraform model for Resource
type ResourceModel struct {
	Excluded types.Bool   `tfsdk:"excluded"`
	Id       types.String `tfsdk:"id"`
}

// ResourceAttrTypes contains the attribute types for ResourceModel
var ResourceAttrTypes = map[string]attr.Type{
	"excluded": types.BoolType,
	"id":       types.StringType,
}

// ResourceResourceSchemaAttributes contains the schema attributes for ResourceModel
var ResourceResourceSchemaAttributes = map[string]schema.Attribute{
	"excluded": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If set true, the resource set of a particular category is excluded from discovery.",
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource set ID.",
	},
}

// ExpandResource converts a Terraform Object to SDK type
func ExpandResource(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiclouddiscovery.Resource {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ResourceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ResourceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiclouddiscovery.Resource {
	if m == nil {
		return nil
	}
	to := &uddiclouddiscovery.Resource{
		Excluded: flex.ExpandBoolPointer(m.Excluded),
		Id:       flex.ExpandStringPointer(m.Id),
	}
	return to
}

// FlattenResource converts an SDK type to Terraform Object
func FlattenResource(ctx context.Context, from *uddiclouddiscovery.Resource, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ResourceAttrTypes)
	}
	m := &ResourceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ResourceAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ResourceModel) Flatten(ctx context.Context, from *uddiclouddiscovery.Resource, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Excluded = flex.FlattenBoolPointer(from.Excluded)
	m.Id = flex.FlattenStringPointer(from.Id)
}
