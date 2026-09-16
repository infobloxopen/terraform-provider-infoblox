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

type HaGroupModel struct {
	Id   types.String `tfsdk:"id"`
	UDDI types.Object `tfsdk:"uddi"`
}

var HaGroupAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"uddi": types.ObjectType{AttrTypes: UDDIHaGroupAttrTypes},
}

type UDDIHaGroupModel struct {
	AnycastConfigId types.String `tfsdk:"anycast_config_id"`
	Comment         types.String `tfsdk:"comment"`
	Hosts           types.List   `tfsdk:"hosts"`
	IpSpace         types.String `tfsdk:"ip_space"`
	Mode            types.String `tfsdk:"mode"`
	Name            types.String `tfsdk:"name"`
	Tags            types.Map    `tfsdk:"tags"`
	TagsAll         types.Map    `tfsdk:"tags_all"`
}

var UDDIHaGroupAttrTypes = map[string]attr.Type{
	"anycast_config_id": types.StringType,
	"comment":           types.StringType,
	"hosts":             types.ListType{ElemType: types.ObjectType{AttrTypes: HAGroupHostAttrTypes}},
	"ip_space":          types.StringType,
	"mode":              types.StringType,
	"name":              types.StringType,
	"tags":              types.MapType{ElemType: types.StringType},
	"tags_all":          types.MapType{ElemType: types.StringType},
}

const (
	HaGroupReturnFields = ""
)

var HaGroupResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          HaGroupResourceUddiSchemaAttributes,
	},
}

var HaGroupResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"anycast_config_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"comment": schema.StringAttribute{
		Default:  stringdefault.StaticString(""),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 1024),
		},
		MarkdownDescription: "The description for the HA group. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"hosts": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: HAGroupHostResourceSchemaAttributes,
		},
		Required: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of hosts.",
	},
	"ip_space": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"mode": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("active-active", "active-passive", "advanced-active-passive", "split-ranges"),
		},
		Required:            true,
		MarkdownDescription: "The mode of the HA group.  Valid values are: * _active-active_: Both on-prem hosts remain active. * _active-passive_: One on-prem host remains active and one remains passive. When the active on-prem host is down, the passive on-prem host takes over. * _advanced-active-passive_: One on-prem host may be part of multiple HA groups. When the active on-prem host is down, the passive on-prem host takes over.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(1, 256),
		},
		MarkdownDescription: "The name of the HA group. Must contain 1 to 256 characters. Can include UTF-8.",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The tags for the HA group.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *HaGroupModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.HaGroup {
	if m == nil {
		return nil
	}

	obj := &coremodel.HaGroup{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIHaGroupModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIHaGroupModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIHaGroupExt {
	return &coremodel.UDDIHaGroupExt{
		AnycastConfigId: flex.ExpandStringPointer(m.AnycastConfigId),
		Comment:         flex.ExpandStringPointer(m.Comment),
		Hosts:           flex.ExpandFrameworkListNestedBlock(ctx, m.Hosts, diags, ExpandHAGroupHost),
		Mode:            flex.ExpandStringPointer(m.Mode),
		Name:            flex.ExpandString(m.Name),
		Tags:            flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
}

// Flatten populates the TF model from a core response.
func (m *HaGroupModel) Flatten(ctx context.Context, resp *coremodel.HaGroup, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIHaGroupModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIHaGroupModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIHaGroupAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIHaGroupAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIHaGroupModel) Flatten(ctx context.Context, from *coremodel.UDDIHaGroupExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AnycastConfigId = flex.FlattenStringPointer(from.AnycastConfigId)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Hosts = flex.FlattenFrameworkListNestedBlock(ctx, from.Hosts, HAGroupHostAttrTypes, diags, FlattenHAGroupHost)
	m.IpSpace = flex.FlattenStringPointer(from.IpSpace)
	m.Mode = flex.FlattenStringPointer(from.Mode)
	m.Name = flex.FlattenString(from.Name)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
}
