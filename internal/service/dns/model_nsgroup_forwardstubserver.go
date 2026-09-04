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

type NsgroupForwardstubserverModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var NsgroupForwardstubserverAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSNsgroupForwardstubserverAttrTypes},
}

type NIOSNsgroupForwardstubserverModel struct {
	Comment         types.String `tfsdk:"comment"`
	ExtAttrs        types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll     types.Map    `tfsdk:"ext_attrs_all"`
	ExternalServers types.List   `tfsdk:"external_servers"`
	Name            types.String `tfsdk:"name"`
}

var NIOSNsgroupForwardstubserverAttrTypes = map[string]attr.Type{
	"comment":          types.StringType,
	"ext_attrs":        types.MapType{ElemType: types.StringType},
	"ext_attrs_all":    types.MapType{ElemType: types.StringType},
	"external_servers": types.ListType{ElemType: types.ObjectType{AttrTypes: NsgroupForwardstubserverExternalServersAttrTypes}},
	"name":             types.StringType,
}

const (
	NsgroupForwardstubserverReturnFields = "comment,extattrs,external_servers,name"
)

var NsgroupForwardstubserverResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          NsgroupForwardstubserverResourceNiosSchemaAttributes,
	},
}

var NsgroupForwardstubserverResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for the Forward Stub Server Name Server Group; maximum 256 characters.",
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
	"external_servers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NsgroupForwardstubserverExternalServersResourceSchemaAttributes,
		},
		Required: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of external servers.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of this Forward Stub Server Name Server Group.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *NsgroupForwardstubserverModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.NsgroupForwardstubserver {
	if m == nil {
		return nil
	}

	obj := &coremodel.NsgroupForwardstubserver{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSNsgroupForwardstubserverModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSNsgroupForwardstubserverModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSNsgroupForwardstubserverExt {
	return &coremodel.NIOSNsgroupForwardstubserverExt{
		Comment:         flex.ExpandStringPointerNullAsEmpty(m.Comment),
		ExtAttrs:        flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		ExternalServers: flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalServers, diags, ExpandNsgroupForwardstubserverExternalServers),
		Name:            flex.ExpandStringPointerNullAsEmpty(m.Name),
	}
}

// Flatten populates the TF model from a core response.
func (m *NsgroupForwardstubserverModel) Flatten(ctx context.Context, resp *coremodel.NsgroupForwardstubserver, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSNsgroupForwardstubserverModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSNsgroupForwardstubserverModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSNsgroupForwardstubserverAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSNsgroupForwardstubserverAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSNsgroupForwardstubserverModel) Flatten(ctx context.Context, from *coremodel.NIOSNsgroupForwardstubserverExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.ExternalServers = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalServers, NsgroupForwardstubserverExternalServersAttrTypes, diags, FlattenNsgroupForwardstubserverExternalServers)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
}
