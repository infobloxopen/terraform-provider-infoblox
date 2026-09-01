package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	internaltypes "github.com/infobloxopen/terraform-provider-infoblox/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type SuperhostModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var SuperhostAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSSuperhostAttrTypes},
}

type NIOSSuperhostModel struct {
	Comment                 types.String                     `tfsdk:"comment"`
	DeleteAssociatedObjects types.Bool                       `tfsdk:"delete_associated_objects"`
	DhcpAssociatedObjects   internaltypes.UnorderedListValue `tfsdk:"dhcp_associated_objects"`
	Disabled                types.Bool                       `tfsdk:"disabled"`
	DnsAssociatedObjects    internaltypes.UnorderedListValue `tfsdk:"dns_associated_objects"`
	ExtAttrs                types.Map                        `tfsdk:"ext_attrs"`
	ExtAttrsAll             types.Map                        `tfsdk:"ext_attrs_all"`
	Name                    types.String                     `tfsdk:"name"`
}

var NIOSSuperhostAttrTypes = map[string]attr.Type{
	"comment":                   types.StringType,
	"delete_associated_objects": types.BoolType,
	"dhcp_associated_objects":   internaltypes.UnorderedListOfStringType,
	"disabled":                  types.BoolType,
	"dns_associated_objects":    internaltypes.UnorderedListOfStringType,
	"ext_attrs":                 types.MapType{ElemType: types.StringType},
	"ext_attrs_all":             types.MapType{ElemType: types.StringType},
	"name":                      types.StringType,
}

const (
	SuperhostReturnFields = "comment,dhcp_associated_objects,disabled,dns_associated_objects,extattrs,name"
)

var SuperhostResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          SuperhostResourceNiosSchemaAttributes,
	},
}

var SuperhostResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The comment for Super Host.",
	},
	"delete_associated_objects": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "True if we have to delete all DNS/DHCP associated objects with Super Host, false by default.",
	},
	"dhcp_associated_objects": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		CustomType:  internaltypes.UnorderedListOfStringType,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "A list of DHCP objects refs which are associated with Super Host.",
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Disable all DNS/DHCP associated objects with Super Host if True, False by default.",
	},
	"dns_associated_objects": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		CustomType:  internaltypes.UnorderedListOfStringType,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "A list of object refs of the DNS resource records which are associated with Super Host.",
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
		MarkdownDescription: "Name of the Superhost.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *SuperhostModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Superhost {
	if m == nil {
		return nil
	}

	obj := &coremodel.Superhost{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSSuperhostModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSSuperhostModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSSuperhostExt {
	return &coremodel.NIOSSuperhostExt{
		Comment:                 flex.ExpandStringPointerNullAsEmpty(m.Comment),
		DeleteAssociatedObjects: flex.ExpandBoolPointer(m.DeleteAssociatedObjects),
		DhcpAssociatedObjects:   flex.ExpandFrameworkListString(ctx, m.DhcpAssociatedObjects, diags),
		Disabled:                flex.ExpandBoolPointer(m.Disabled),
		DnsAssociatedObjects:    flex.ExpandFrameworkListString(ctx, m.DnsAssociatedObjects, diags),
		ExtAttrs:                flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		Name:                    flex.ExpandStringPointerNullAsEmpty(m.Name),
	}
}

// Flatten populates the TF model from a core response.
func (m *SuperhostModel) Flatten(ctx context.Context, resp *coremodel.Superhost, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSSuperhostModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSSuperhostModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSSuperhostAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSSuperhostAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSSuperhostModel) Flatten(ctx context.Context, from *coremodel.NIOSSuperhostExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.DhcpAssociatedObjects = flex.FlattenFrameworkUnorderedListString(ctx, from.DhcpAssociatedObjects, diags)
	m.Disabled = flex.FlattenBoolPointer(from.Disabled)
	m.DnsAssociatedObjects = flex.FlattenFrameworkUnorderedListString(ctx, from.DnsAssociatedObjects, diags)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
}
