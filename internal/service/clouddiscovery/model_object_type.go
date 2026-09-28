package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
	uddiclouddiscovery "github.com/infobloxopen/universal-ddi-go-client/clouddiscovery"
)

// ObjectTypeModel is the Terraform model for ObjectType
type ObjectTypeModel struct {
	DiscoverNew types.Bool    `tfsdk:"discover_new"`
	Objects     types.List    `tfsdk:"objects"`
	Version     types.Float64 `tfsdk:"version"`
}

// ObjectTypeAttrTypes contains the attribute types for ObjectTypeModel
var ObjectTypeAttrTypes = map[string]attr.Type{
	"discover_new": types.BoolType,
	"objects":      types.ListType{ElemType: types.ObjectType{AttrTypes: ObjectAttrTypes}},
	"version":      types.Float64Type,
}

// ObjectTypeResourceSchemaAttributes contains the schema attributes for ObjectTypeModel
var ObjectTypeResourceSchemaAttributes = map[string]schema.Attribute{
	"discover_new": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Discover new objects.",
	},
	"objects": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ObjectResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "List of objects to discover.",
	},
	"version": schema.Float64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Version of the object type.",
	},
}

// ExpandObjectType converts a Terraform Object to SDK type
func ExpandObjectType(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiclouddiscovery.ObjectType {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ObjectTypeModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ObjectTypeModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiclouddiscovery.ObjectType {
	if m == nil {
		return nil
	}
	to := &uddiclouddiscovery.ObjectType{
		DiscoverNew: flex.ExpandBoolPointer(m.DiscoverNew),
		Objects:     flex.ExpandFrameworkListNestedBlock(ctx, m.Objects, diags, ExpandObject),
		Version:     flex.ExpandFloat32Pointer(m.Version),
	}
	return to
}

// FlattenObjectType converts an SDK type to Terraform Object
func FlattenObjectType(ctx context.Context, from *uddiclouddiscovery.ObjectType, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ObjectTypeAttrTypes)
	}
	m := &ObjectTypeModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ObjectTypeAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ObjectTypeModel) Flatten(ctx context.Context, from *uddiclouddiscovery.ObjectType, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.DiscoverNew = flex.FlattenBoolPointer(from.DiscoverNew)
	m.Objects = flex.FlattenFrameworkListNestedBlock(ctx, from.Objects, ObjectAttrTypes, diags, FlattenObject)
	m.Version = flex.FlattenFloat32Pointer(from.Version)
}
