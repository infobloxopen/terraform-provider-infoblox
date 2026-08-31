package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

type ForwardNsgModel struct {
	Id   types.String `tfsdk:"id"`
	UDDI types.Object `tfsdk:"uddi"`
}

var ForwardNsgAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"uddi": types.ObjectType{AttrTypes: UDDIForwardNsgAttrTypes},
}

type UDDIForwardNsgModel struct {
	Comment            types.String `tfsdk:"comment"`
	ExternalForwarders types.List   `tfsdk:"external_forwarders"`
	ForwardersOnly     types.Bool   `tfsdk:"forwarders_only"`
	Hosts              types.List   `tfsdk:"hosts"`
	InternalForwarders types.List   `tfsdk:"internal_forwarders"`
	Name               types.String `tfsdk:"name"`
	Nsgs               types.List   `tfsdk:"nsgs"`
	Tags               types.Map    `tfsdk:"tags"`
	TagsAll            types.Map    `tfsdk:"tags_all"`
}

var UDDIForwardNsgAttrTypes = map[string]attr.Type{
	"comment":             types.StringType,
	"external_forwarders": types.ListType{ElemType: types.ObjectType{AttrTypes: ForwarderAttrTypes}},
	"forwarders_only":     types.BoolType,
	"hosts":               types.ListType{ElemType: types.StringType},
	"internal_forwarders": types.ListType{ElemType: types.StringType},
	"name":                types.StringType,
	"nsgs":                types.ListType{ElemType: types.StringType},
	"tags":                types.MapType{ElemType: types.StringType},
	"tags_all":            types.MapType{ElemType: types.StringType},
}

const (
	ForwardNsgReturnFields = ""
)

var ForwardNsgResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          ForwardNsgResourceUddiSchemaAttributes,
	},
}

var ForwardNsgResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Comment for the object.",
	},
	"external_forwarders": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ForwarderResourceSchemaAttributes(true),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. External DNS servers to forward to. Order is not significant.",
	},
	"forwarders_only": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. _true_ to only forward.",
	},
	"hosts": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"internal_forwarders": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Name of the object.",
	},
	"nsgs": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
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
func (m *ForwardNsgModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.ForwardNsg {
	if m == nil {
		return nil
	}

	obj := &coremodel.ForwardNsg{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIForwardNsgModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIForwardNsgModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIForwardNsgExt {
	return &coremodel.UDDIForwardNsgExt{
		Comment:            flex.ExpandStringPointer(m.Comment),
		ExternalForwarders: flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalForwarders, diags, ExpandForwarder),
		ForwardersOnly:     flex.ExpandBoolPointer(m.ForwardersOnly),
		Hosts:              flex.ExpandFrameworkListString(ctx, m.Hosts, diags),
		InternalForwarders: flex.ExpandFrameworkListString(ctx, m.InternalForwarders, diags),
		Name:               flex.ExpandString(m.Name),
		Nsgs:               flex.ExpandFrameworkListString(ctx, m.Nsgs, diags),
		Tags:               flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
}

// Flatten populates the TF model from a core response.
func (m *ForwardNsgModel) Flatten(ctx context.Context, resp *coremodel.ForwardNsg, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIForwardNsgModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIForwardNsgModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIForwardNsgAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIForwardNsgAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIForwardNsgModel) Flatten(ctx context.Context, from *coremodel.UDDIForwardNsgExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.ExternalForwarders = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalForwarders, ForwarderAttrTypes, diags, FlattenForwarder)
	m.ForwardersOnly = flex.FlattenBoolPointer(from.ForwardersOnly)
	m.Hosts = flex.FlattenFrameworkListString(ctx, from.Hosts, diags)
	m.InternalForwarders = flex.FlattenFrameworkListString(ctx, from.InternalForwarders, diags)
	m.Name = flex.FlattenString(from.Name)
	m.Nsgs = flex.FlattenFrameworkListString(ctx, from.Nsgs, diags)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
}
