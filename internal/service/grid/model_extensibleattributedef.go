package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/grid"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	immutable "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/immutable"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type ExtensibleattributedefModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var ExtensibleattributedefAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSExtensibleattributedefAttrTypes},
}

type NIOSExtensibleattributedefModel struct {
	AllowedObjectTypes types.List   `tfsdk:"allowed_object_types"`
	Comment            types.String `tfsdk:"comment"`
	DefaultValue       types.String `tfsdk:"default_value"`
	Flags              types.String `tfsdk:"flags"`
	ListValues         types.List   `tfsdk:"list_values"`
	Max                types.Int64  `tfsdk:"max"`
	Min                types.Int64  `tfsdk:"min"`
	Name               types.String `tfsdk:"name"`
	Type               types.String `tfsdk:"type"`
}

var NIOSExtensibleattributedefAttrTypes = map[string]attr.Type{
	"allowed_object_types": types.ListType{ElemType: types.StringType},
	"comment":              types.StringType,
	"default_value":        types.StringType,
	"flags":                types.StringType,
	"list_values":          types.ListType{ElemType: types.ObjectType{AttrTypes: ExtensibleattributedefListValuesAttrTypes}},
	"max":                  types.Int64Type,
	"min":                  types.Int64Type,
	"name":                 types.StringType,
	"type":                 types.StringType,
}

const (
	ExtensibleattributedefReturnFields = "allowed_object_types,comment,default_value,flags,list_values,max,min,name,namespace,type"
)

var ExtensibleattributedefResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          ExtensibleattributedefResourceNiosSchemaAttributes,
	},
}

var ExtensibleattributedefResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"allowed_object_types": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		Default:     listdefault.StaticValue(types.ListNull(types.StringType)),
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The object types this extensible attribute is allowed to associate with.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "Comment for the Extensible Attribute Definition; maximum 256 characters.",
	},
	"default_value": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Default value used to pre-populate the attribute value in the GUI. For email, URL, and string types, the value is a string with a maximum of 256 characters. For an integer, the value is an integer from -2147483648 through 2147483647. For a date, the value is the number of seconds that have elapsed since January 1st, 1970 UTC.",
	},
	"flags": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "This field contains extensible attribute flags. Possible values: (A)udited, (C)loud API, Cloud (G)master, (I)nheritable, (L)isted, (M)andatory value, MGM (P)rivate, (R)ead Only, (S)ort enum values, Multiple (V)alues If there are two or more flags in the field, you must list them according to the order they are listed above. For example, 'CR' is a valid value for the 'flags' field because C = Cloud API is listed before R = Read only. However, the value 'RC' is invalid because the order for the 'flags' field is broken.",
	},
	"list_values": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ExtensibleattributedefListValuesResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Default:  listdefault.StaticValue(types.ListNull(types.ObjectType{AttrTypes: ExtensibleattributedefListValuesAttrTypes})),
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "List of Values. Applicable if the extensible attribute type is ENUM.",
	},
	"max": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Maximum allowed value of extensible attribute. Applicable if the extensible attribute type is INTEGER.",
	},
	"min": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Minimum allowed value of extensible attribute. Applicable if the extensible attribute type is INTEGER.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the Extensible Attribute Definition.",
	},
	"type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("STRING", "INTEGER", "EMAIL", "DATE", "ENUM", "URL"),
		},
		Required: true,
		PlanModifiers: []planmodifier.String{
			immutable.ImmutableString(),
		},
		MarkdownDescription: "Type for the Extensible Attribute Definition.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *ExtensibleattributedefModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Extensibleattributedef {
	if m == nil {
		return nil
	}

	obj := &coremodel.Extensibleattributedef{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSExtensibleattributedefModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSExtensibleattributedefModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSExtensibleattributedefExt {
	return &coremodel.NIOSExtensibleattributedefExt{
		AllowedObjectTypes: flex.ExpandFrameworkListString(ctx, m.AllowedObjectTypes, diags),
		Comment:            flex.ExpandStringPointerNullAsEmpty(m.Comment),
		DefaultValue:       flex.ExpandExtensibleAttributeDefDefaultValue(ctx, m.DefaultValue, m.Type, diags),
		Flags:              flex.ExpandStringPointerNullAsEmpty(m.Flags),
		ListValues:         flex.ExpandFrameworkListNestedBlock(ctx, m.ListValues, diags, ExpandExtensibleattributedefListValues),
		Max:                flex.ExpandInt64Pointer(m.Max),
		Min:                flex.ExpandInt64Pointer(m.Min),
		Name:               flex.ExpandStringPointerNullAsEmpty(m.Name),
		Type:               flex.ExpandStringPointerNullAsEmpty(m.Type),
	}
}

// Flatten populates the TF model from a core response.
func (m *ExtensibleattributedefModel) Flatten(ctx context.Context, resp *coremodel.Extensibleattributedef, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSExtensibleattributedefModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSExtensibleattributedefModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSExtensibleattributedefAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSExtensibleattributedefAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSExtensibleattributedefModel) Flatten(ctx context.Context, from *coremodel.NIOSExtensibleattributedefExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AllowedObjectTypes = flex.FlattenFrameworkListString(ctx, from.AllowedObjectTypes, diags)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.DefaultValue = flex.FlattenExtensibleAttributeDefDefaultValue(ctx, from.DefaultValue, diags)
	m.Flags = flex.FlattenStringPointerEmptyAsNull(from.Flags)
	m.ListValues = flex.FlattenFrameworkListNestedBlock(ctx, from.ListValues, ExtensibleattributedefListValuesAttrTypes, diags, FlattenExtensibleattributedefListValues)
	m.Max = flex.FlattenInt64Pointer(from.Max)
	m.Min = flex.FlattenInt64Pointer(from.Min)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Type = flex.FlattenStringPointerEmptyAsNull(from.Type)
}
