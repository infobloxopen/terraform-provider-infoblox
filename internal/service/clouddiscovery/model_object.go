package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
	uddiclouddiscovery "github.com/infobloxopen/universal-ddi-go-client/clouddiscovery"
)

// ObjectModel is the Terraform model for Object
type ObjectModel struct {
	Category    types.Object `tfsdk:"category"`
	ResourceSet types.List   `tfsdk:"resource_set"`
}

// ObjectAttrTypes contains the attribute types for ObjectModel
var ObjectAttrTypes = map[string]attr.Type{
	"category":     types.ObjectType{AttrTypes: CategoryAttrTypes},
	"resource_set": types.ListType{ElemType: types.ObjectType{AttrTypes: ResourceAttrTypes}},
}

// ObjectResourceSchemaAttributes contains the schema attributes for ObjectModel
var ObjectResourceSchemaAttributes = map[string]schema.Attribute{
	"category": schema.SingleNestedAttribute{
		Attributes:          CategoryResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "",
	},
	"resource_set": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ResourceResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "",
	},
}

// ExpandObject converts a Terraform Object to SDK type
func ExpandObject(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiclouddiscovery.Object {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ObjectModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ObjectModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiclouddiscovery.Object {
	if m == nil {
		return nil
	}
	to := &uddiclouddiscovery.Object{
		Category:    ExpandCategory(ctx, m.Category, diags),
		ResourceSet: flex.ExpandFrameworkListNestedBlock(ctx, m.ResourceSet, diags, ExpandResource),
	}
	return to
}

// FlattenObject converts an SDK type to Terraform Object
func FlattenObject(ctx context.Context, from *uddiclouddiscovery.Object, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ObjectAttrTypes)
	}
	m := &ObjectModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ObjectAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ObjectModel) Flatten(ctx context.Context, from *uddiclouddiscovery.Object, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Category = FlattenCategory(ctx, from.Category, diags)
	m.ResourceSet = flex.FlattenFrameworkListNestedBlock(ctx, from.ResourceSet, ResourceAttrTypes, diags, FlattenResource)
}
