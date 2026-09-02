package acl

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/acl"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type NamedaclModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
	UDDI types.Object `tfsdk:"uddi"`
}

var NamedaclAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSNamedaclAttrTypes},
	"uddi": types.ObjectType{AttrTypes: UDDINamedaclAttrTypes},
}

type NIOSNamedaclModel struct {
	AccessList  types.List   `tfsdk:"access_list"`
	Comment     types.String `tfsdk:"comment"`
	ExtAttrs    types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll types.Map    `tfsdk:"ext_attrs_all"`
	Name        types.String `tfsdk:"name"`
}

var NIOSNamedaclAttrTypes = map[string]attr.Type{
	"access_list":   types.ListType{ElemType: types.ObjectType{AttrTypes: NamedaclAccessListAttrTypes}},
	"comment":       types.StringType,
	"ext_attrs":     types.MapType{ElemType: types.StringType},
	"ext_attrs_all": types.MapType{ElemType: types.StringType},
	"name":          types.StringType,
}

type UDDINamedaclModel struct {
	Comment       types.String `tfsdk:"comment"`
	CompartmentId types.String `tfsdk:"compartment_id"`
	List          types.List   `tfsdk:"list"`
	Name          types.String `tfsdk:"name"`
	Tags          types.Map    `tfsdk:"tags"`
	TagsAll       types.Map    `tfsdk:"tags_all"`
}

var UDDINamedaclAttrTypes = map[string]attr.Type{
	"comment":        types.StringType,
	"compartment_id": types.StringType,
	"list":           types.ListType{ElemType: types.ObjectType{AttrTypes: ACLItemAttrTypes}},
	"name":           types.StringType,
	"tags":           types.MapType{ElemType: types.StringType},
	"tags_all":       types.MapType{ElemType: types.StringType},
}

const (
	NamedaclReturnFields = "access_list,comment,exploded_access_list,extattrs,name"
)

var NamedaclResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          NamedaclResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          NamedaclResourceUddiSchemaAttributes,
	},
}

var NamedaclResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"access_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NamedaclAccessListResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Default:  listdefault.StaticValue(types.ListNull(types.ObjectType{AttrTypes: NamedaclAccessListAttrTypes})),
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The access control list of IPv4/IPv6 addresses, networks, TSIG-based anonymous access controls, and other named ACLs.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "Comment for the named ACL; maximum 256 characters.",
	},
	"ext_attrs": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"ext_attrs_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All ext_attrs including Terraform Internal ID and inherited attributes.",
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of the named ACL.",
	},
}

var NamedaclResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Comment for ACL.",
	},
	"compartment_id": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The access view associated with the object. If no access view is associated with the object, the value defaults to empty.",
	},
	"list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ACLItemResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. Ordered list of access control elements.  Elements are evaluated in order to determine access. If evaluation reaches the end of the list then access is denied.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "ACL object name.",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Tagging specifics.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *NamedaclModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Namedacl {
	if m == nil {
		return nil
	}

	obj := &coremodel.Namedacl{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSNamedaclModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDINamedaclModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSNamedaclModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSNamedaclExt {
	return &coremodel.NIOSNamedaclExt{
		AccessList: flex.ExpandFrameworkListNestedBlock(ctx, m.AccessList, diags, ExpandNamedaclAccessList),
		Comment:    flex.ExpandStringPointerNullAsEmpty(m.Comment),
		ExtAttrs:   flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		Name:       flex.ExpandStringPointerNullAsEmpty(m.Name),
	}
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDINamedaclModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDINamedaclExt {
	return &coremodel.UDDINamedaclExt{
		Comment:       flex.ExpandStringPointer(m.Comment),
		CompartmentId: flex.ExpandStringPointer(m.CompartmentId),
		List:          flex.ExpandFrameworkListNestedBlock(ctx, m.List, diags, ExpandACLItem),
		Name:          flex.ExpandString(m.Name),
		Tags:          flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
}

// Flatten populates the TF model from a core response.
func (m *NamedaclModel) Flatten(ctx context.Context, resp *coremodel.Namedacl, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSNamedaclModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSNamedaclModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSNamedaclModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenNamedaclNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSNamedaclAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSNamedaclAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDINamedaclModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDINamedaclModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDINamedaclAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDINamedaclAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSNamedaclModel) Flatten(ctx context.Context, from *coremodel.NIOSNamedaclExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.AccessList = flex.FlattenFrameworkListNestedBlock(ctx, from.AccessList, NamedaclAccessListAttrTypes, diags, FlattenNamedaclAccessList)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDINamedaclModel) Flatten(ctx context.Context, from *coremodel.UDDINamedaclExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CompartmentId = flex.FlattenStringPointer(from.CompartmentId)
	m.List = flex.FlattenFrameworkListNestedBlock(ctx, from.List, ACLItemAttrTypes, diags, FlattenACLItem)
	m.Name = flex.FlattenString(from.Name)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
}
