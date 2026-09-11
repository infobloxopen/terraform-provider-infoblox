package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type NsgroupStubmemberModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var NsgroupStubmemberAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSNsgroupStubmemberAttrTypes},
}

type NIOSNsgroupStubmemberModel struct {
	Comment     types.String `tfsdk:"comment"`
	ExtAttrs    types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll types.Map    `tfsdk:"ext_attrs_all"`
	Name        types.String `tfsdk:"name"`
	StubMembers types.List   `tfsdk:"stub_members"`
}

var NIOSNsgroupStubmemberAttrTypes = map[string]attr.Type{
	"comment":       types.StringType,
	"ext_attrs":     types.MapType{ElemType: types.StringType},
	"ext_attrs_all": types.MapType{ElemType: types.StringType},
	"name":          types.StringType,
	"stub_members":  types.ListType{ElemType: types.ObjectType{AttrTypes: NsgroupStubmemberStubMembersAttrTypes}},
}

const (
	NsgroupStubmemberReturnFields = "comment,extattrs,name,stub_members"
)

var NsgroupStubmemberResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          NsgroupStubmemberResourceNiosSchemaAttributes,
	},
}

var NsgroupStubmemberResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthAtMost(256),
		},
		MarkdownDescription: "Comment for the Stub Member Name Server Group; maximum 256 characters.",
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
		MarkdownDescription: "The name of the Stub Member Name Server Group.",
	},
	"stub_members": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NsgroupStubmemberStubMembersResourceSchemaAttributes,
		},
		Required: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The Grid member servers of this stub zone. Note that the lead/stealth/grid_replicate/ preferred_primaries/override_preferred_primaries fields of the struct will be ignored when set in this field.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *NsgroupStubmemberModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.NsgroupStubmember {
	if m == nil {
		return nil
	}

	obj := &coremodel.NsgroupStubmember{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSNsgroupStubmemberModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSNsgroupStubmemberModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSNsgroupStubmemberExt {
	return &coremodel.NIOSNsgroupStubmemberExt{
		Comment:     flex.ExpandStringPointerNullAsEmpty(m.Comment),
		ExtAttrs:    flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		Name:        flex.ExpandStringPointerNullAsEmpty(m.Name),
		StubMembers: flex.ExpandFrameworkListNestedBlock(ctx, m.StubMembers, diags, ExpandNsgroupStubmemberStubMembers),
	}
}

// Flatten populates the TF model from a core response.
func (m *NsgroupStubmemberModel) Flatten(ctx context.Context, resp *coremodel.NsgroupStubmember, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSNsgroupStubmemberModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSNsgroupStubmemberModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSNsgroupStubmemberAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSNsgroupStubmemberAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSNsgroupStubmemberModel) Flatten(ctx context.Context, from *coremodel.NIOSNsgroupStubmemberExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.StubMembers = flex.FlattenFrameworkListNestedBlock(ctx, from.StubMembers, NsgroupStubmemberStubMembersAttrTypes, diags, FlattenNsgroupStubmemberStubMembers)
}
