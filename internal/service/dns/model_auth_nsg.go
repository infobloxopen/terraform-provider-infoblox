package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

type AuthNsgModel struct {
	Id   types.String `tfsdk:"id"`
	UDDI types.Object `tfsdk:"uddi"`
}

var AuthNsgAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"uddi": types.ObjectType{AttrTypes: UDDIAuthNsgAttrTypes},
}

type UDDIAuthNsgModel struct {
	Comment             types.String `tfsdk:"comment"`
	ExternalPrimaries   types.List   `tfsdk:"external_primaries"`
	ExternalSecondaries types.List   `tfsdk:"external_secondaries"`
	InternalSecondaries types.List   `tfsdk:"internal_secondaries"`
	Name                types.String `tfsdk:"name"`
	Nsgs                types.List   `tfsdk:"nsgs"`
	Tags                types.Map    `tfsdk:"tags"`
	TagsAll             types.Map    `tfsdk:"tags_all"`
}

var UDDIAuthNsgAttrTypes = map[string]attr.Type{
	"comment":              types.StringType,
	"external_primaries":   types.ListType{ElemType: types.ObjectType{AttrTypes: ExternalPrimaryAttrTypes}},
	"external_secondaries": types.ListType{ElemType: types.ObjectType{AttrTypes: ExternalSecondaryAttrTypes}},
	"internal_secondaries": types.ListType{ElemType: types.ObjectType{AttrTypes: InternalSecondaryAttrTypes}},
	"name":                 types.StringType,
	"nsgs":                 types.ListType{ElemType: types.StringType},
	"tags":                 types.MapType{ElemType: types.StringType},
	"tags_all":             types.MapType{ElemType: types.StringType},
}

const (
	AuthNsgReturnFields = ""
)

var AuthNsgResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          AuthNsgResourceUddiSchemaAttributes,
	},
}

var AuthNsgResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Comment for the object.",
	},
	"external_primaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ExternalPrimaryResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "Optional. DNS primaries external to Universal DDI. Order is not significant.",
	},
	"external_secondaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ExternalSecondaryResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "DNS secondaries external to Universal DDI. Order is not significant.",
	},
	"internal_secondaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: InternalSecondaryResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "Optional. Universal DDI hosts acting as internal secondaries. Order is not significant.",
	},
	"name": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
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
func (m *AuthNsgModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.AuthNsg {
	if m == nil {
		return nil
	}

	obj := &coremodel.AuthNsg{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIAuthNsgModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIAuthNsgModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIAuthNsgExt {
	return &coremodel.UDDIAuthNsgExt{
		Comment:             flex.ExpandStringPointer(m.Comment),
		ExternalPrimaries:   flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalPrimaries, diags, ExpandExternalPrimary),
		ExternalSecondaries: flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalSecondaries, diags, ExpandExternalSecondary),
		InternalSecondaries: flex.ExpandFrameworkListNestedBlock(ctx, m.InternalSecondaries, diags, ExpandInternalSecondary),
		Name:                flex.ExpandString(m.Name),
		Nsgs:                flex.ExpandFrameworkListString(ctx, m.Nsgs, diags),
		Tags:                flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
}

// Flatten populates the TF model from a core response.
func (m *AuthNsgModel) Flatten(ctx context.Context, resp *coremodel.AuthNsg, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIAuthNsgModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIAuthNsgModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIAuthNsgAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIAuthNsgAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIAuthNsgModel) Flatten(ctx context.Context, from *coremodel.UDDIAuthNsgExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.ExternalPrimaries = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalPrimaries, ExternalPrimaryAttrTypes, diags, FlattenExternalPrimary)
	m.ExternalSecondaries = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalSecondaries, ExternalSecondaryAttrTypes, diags, FlattenExternalSecondary)
	m.InternalSecondaries = flex.FlattenFrameworkListNestedBlock(ctx, from.InternalSecondaries, InternalSecondaryAttrTypes, diags, FlattenInternalSecondary)
	m.Name = flex.FlattenString(from.Name)
	m.Nsgs = flex.FlattenFrameworkListString(ctx, from.Nsgs, diags)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
}
