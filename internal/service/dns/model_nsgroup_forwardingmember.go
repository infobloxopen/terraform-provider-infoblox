package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
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

type NsgroupForwardingmemberModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var NsgroupForwardingmemberAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSNsgroupForwardingmemberAttrTypes},
}

type NIOSNsgroupForwardingmemberModel struct {
	Comment           types.String `tfsdk:"comment"`
	ExtAttrs          types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll       types.Map    `tfsdk:"ext_attrs_all"`
	ForwardingServers types.List   `tfsdk:"forwarding_servers"`
	Name              types.String `tfsdk:"name"`
}

var NIOSNsgroupForwardingmemberAttrTypes = map[string]attr.Type{
	"comment":            types.StringType,
	"ext_attrs":          types.MapType{ElemType: types.StringType},
	"ext_attrs_all":      types.MapType{ElemType: types.StringType},
	"forwarding_servers": types.ListType{ElemType: types.ObjectType{AttrTypes: NsgroupForwardingmemberForwardingServersAttrTypes}},
	"name":               types.StringType,
}

const (
	NsgroupForwardingmemberReturnFields = "comment,extattrs,forwarding_servers,name"
)

var NsgroupForwardingmemberResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          NsgroupForwardingmemberResourceNiosSchemaAttributes,
	},
}

var NsgroupForwardingmemberResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for the Forwarding Member Name Server Group; maximum 256 characters.",
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
	"forwarding_servers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NsgroupForwardingmemberForwardingServersResourceSchemaAttributes,
		},
		Required: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of forwarding member servers.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of the Forwarding Member Name Server Group.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *NsgroupForwardingmemberModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.NsgroupForwardingmember {
	if m == nil {
		return nil
	}

	obj := &coremodel.NsgroupForwardingmember{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSNsgroupForwardingmemberModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSNsgroupForwardingmemberModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSNsgroupForwardingmemberExt {
	return &coremodel.NIOSNsgroupForwardingmemberExt{
		Comment:           flex.ExpandStringPointerNullAsEmpty(m.Comment),
		ExtAttrs:          flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		ForwardingServers: flex.ExpandFrameworkListNestedBlock(ctx, m.ForwardingServers, diags, ExpandNsgroupForwardingmemberForwardingServers),
		Name:              flex.ExpandStringPointerNullAsEmpty(m.Name),
	}
}

// Flatten populates the TF model from a core response.
func (m *NsgroupForwardingmemberModel) Flatten(ctx context.Context, resp *coremodel.NsgroupForwardingmember, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSNsgroupForwardingmemberModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSNsgroupForwardingmemberModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSNsgroupForwardingmemberAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSNsgroupForwardingmemberAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSNsgroupForwardingmemberModel) Flatten(ctx context.Context, from *coremodel.NIOSNsgroupForwardingmemberExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.ForwardingServers = flex.FlattenFrameworkListNestedBlock(ctx, from.ForwardingServers, NsgroupForwardingmemberForwardingServersAttrTypes, diags, FlattenNsgroupForwardingmemberForwardingServers)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
}
