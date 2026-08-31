package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type DhcpOptionspaceModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
	UDDI types.Object `tfsdk:"uddi"`
}

var DhcpOptionspaceAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSDhcpOptionspaceAttrTypes},
	"uddi": types.ObjectType{AttrTypes: UDDIDhcpOptionspaceAttrTypes},
}

type NIOSDhcpOptionspaceModel struct {
	Comment types.String `tfsdk:"comment"`
	Name    types.String `tfsdk:"name"`
}

var NIOSDhcpOptionspaceAttrTypes = map[string]attr.Type{
	"comment": types.StringType,
	"name":    types.StringType,
}

type UDDIDhcpOptionspaceModel struct {
	Comment  types.String `tfsdk:"comment"`
	Name     types.String `tfsdk:"name"`
	Protocol types.String `tfsdk:"protocol"`
	Tags     types.Map    `tfsdk:"tags"`
	TagsAll  types.Map    `tfsdk:"tags_all"`
}

var UDDIDhcpOptionspaceAttrTypes = map[string]attr.Type{
	"comment":  types.StringType,
	"name":     types.StringType,
	"protocol": types.StringType,
	"tags":     types.MapType{ElemType: types.StringType},
	"tags_all": types.MapType{ElemType: types.StringType},
}

const (
	DhcpOptionspaceReturnFields = "comment,name,option_definitions,space_type"
)

var DhcpOptionspaceResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          DhcpOptionspaceResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          DhcpOptionspaceResourceUddiSchemaAttributes,
	},
}

var DhcpOptionspaceResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "A descriptive comment of a DHCP option space object.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of a DHCP option space object.",
	},
}

var DhcpOptionspaceResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Default:  stringdefault.StaticString(""),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 1024),
		},
		MarkdownDescription: "The description for the option space. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(1, 256),
		},
		MarkdownDescription: "The name of the option space. Must contain 1 to 256 characters. Can include UTF-8.",
	},
	"protocol": schema.StringAttribute{
		Default:             stringdefault.StaticString("ip4"),
		Computed:            true,
		MarkdownDescription: "The type of protocol for the option space (_ip4_ or _ip6_).",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The tags for the option space in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *DhcpOptionspaceModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.DhcpOptionspace {
	if m == nil {
		return nil
	}

	obj := &coremodel.DhcpOptionspace{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSDhcpOptionspaceModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIDhcpOptionspaceModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags, isCreate)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSDhcpOptionspaceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSDhcpOptionspaceExt {
	return &coremodel.NIOSDhcpOptionspaceExt{
		Comment: flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Name:    flex.ExpandStringPointerNullAsEmpty(m.Name),
	}
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIDhcpOptionspaceModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.UDDIDhcpOptionspaceExt {
	ext := &coremodel.UDDIDhcpOptionspaceExt{
		Comment: flex.ExpandStringPointer(m.Comment),
		Name:    flex.ExpandString(m.Name),
		Tags:    flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
	if isCreate {
		ext.Protocol = flex.ExpandStringPointer(m.Protocol)
	}
	return ext
}

// Flatten populates the TF model from a core response.
func (m *DhcpOptionspaceModel) Flatten(ctx context.Context, resp *coremodel.DhcpOptionspace, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSDhcpOptionspaceModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSDhcpOptionspaceModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSDhcpOptionspaceAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSDhcpOptionspaceAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIDhcpOptionspaceModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIDhcpOptionspaceModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIDhcpOptionspaceAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIDhcpOptionspaceAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSDhcpOptionspaceModel) Flatten(ctx context.Context, from *coremodel.NIOSDhcpOptionspaceExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIDhcpOptionspaceModel) Flatten(ctx context.Context, from *coremodel.UDDIDhcpOptionspaceExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Name = flex.FlattenString(from.Name)
	m.Protocol = flex.FlattenStringPointer(from.Protocol)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
}
