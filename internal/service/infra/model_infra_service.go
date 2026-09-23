package infra

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/infra"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

type InfraServiceModel struct {
	Id   types.String `tfsdk:"id"`
	UDDI types.Object `tfsdk:"uddi"`
}

var InfraServiceAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"uddi": types.ObjectType{AttrTypes: UDDIInfraServiceAttrTypes},
}

type UDDIInfraServiceModel struct {
	Configs         types.List        `tfsdk:"configs"`
	CreatedAt       timetypes.RFC3339 `tfsdk:"created_at"`
	Description     types.String      `tfsdk:"description"`
	DesiredState    types.String      `tfsdk:"desired_state"`
	DesiredVersion  types.String      `tfsdk:"desired_version"`
	InterfaceLabels types.List        `tfsdk:"interface_labels"`
	Name            types.String      `tfsdk:"name"`
	PoolId          types.String      `tfsdk:"pool_id"`
	ServiceType     types.String      `tfsdk:"service_type"`
	Tags            types.Map         `tfsdk:"tags"`
	TagsAll         types.Map         `tfsdk:"tags_all"`
	UpdatedAt       timetypes.RFC3339 `tfsdk:"updated_at"`
}

var UDDIInfraServiceAttrTypes = map[string]attr.Type{
	"configs":          types.ListType{ElemType: types.ObjectType{AttrTypes: ServiceHostConfigAttrTypes}},
	"created_at":       timetypes.RFC3339Type{},
	"description":      types.StringType,
	"desired_state":    types.StringType,
	"desired_version":  types.StringType,
	"interface_labels": types.ListType{ElemType: types.StringType},
	"name":             types.StringType,
	"pool_id":          types.StringType,
	"service_type":     types.StringType,
	"tags":             types.MapType{ElemType: types.StringType},
	"tags_all":         types.MapType{ElemType: types.StringType},
	"updated_at":       timetypes.RFC3339Type{},
}

const (
	InfraServiceReturnFields = ""
)

var InfraServiceResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          InfraServiceResourceUddiSchemaAttributes,
	},
}

var InfraServiceResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"configs": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ServiceHostConfigResourceSchemaAttributes,
		},
		Computed:            true,
		MarkdownDescription: "List of Host-specific configurations of this Service.",
	},
	"created_at": schema.StringAttribute{
		Computed:            true,
		CustomType:          timetypes.RFC3339Type{},
		MarkdownDescription: "Timestamp of creation of Service.",
	},
	"description": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The description of the Service (optional).",
	},
	"desired_state": schema.StringAttribute{
		Default: stringdefault.StaticString("stop"),
		Validators: []validator.String{
			stringvalidator.OneOf("start", "stop"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The desired state of the Service. Should either be `\"start\"` or `\"stop\"`.",
	},
	"desired_version": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The desired version of the Service.",
	},
	"interface_labels": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "List of interfaces on which this Service can operate. Note: The list can contain custom interface labels (Example: `[\"WAN\",\"LAN\",\"label1\",\"label2\"]`)",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the Service (unique).",
	},
	"pool_id": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"service_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("authn", "anycast", "cdc", "dhcp", "dns", "dfp", "orpheus", "msad", "ntp", "bgp", "rip", "ospf"),
		},
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The type of the Service deployed on the Host (`dns`, `cdc`, etc.).",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Tags associated with this Service.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
	"updated_at": schema.StringAttribute{
		Computed:            true,
		CustomType:          timetypes.RFC3339Type{},
		MarkdownDescription: "Timestamp of the latest update on Service.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *InfraServiceModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.InfraService {
	if m == nil {
		return nil
	}

	obj := &coremodel.InfraService{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIInfraServiceModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIInfraServiceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIInfraServiceExt {
	return &coremodel.UDDIInfraServiceExt{
		Description:     flex.ExpandStringPointer(m.Description),
		DesiredState:    flex.ExpandStringPointer(m.DesiredState),
		DesiredVersion:  flex.ExpandStringPointer(m.DesiredVersion),
		InterfaceLabels: flex.ExpandFrameworkListString(ctx, m.InterfaceLabels, diags),
		Name:            flex.ExpandString(m.Name),
		PoolId:          flex.ExpandString(m.PoolId),
		ServiceType:     flex.ExpandString(m.ServiceType),
		Tags:            flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
}

// Flatten populates the TF model from a core response.
func (m *InfraServiceModel) Flatten(ctx context.Context, resp *coremodel.InfraService, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIInfraServiceModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIInfraServiceModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIInfraServiceAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIInfraServiceAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIInfraServiceModel) Flatten(ctx context.Context, from *coremodel.UDDIInfraServiceExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Configs = flex.FlattenFrameworkListNestedBlock(ctx, from.Configs, ServiceHostConfigAttrTypes, diags, FlattenServiceHostConfig)
	m.CreatedAt = flex.FlattenRFC3339(from.CreatedAt)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.DesiredState = flex.FlattenStringPointer(from.DesiredState)
	m.DesiredVersion = flex.FlattenStringPointer(from.DesiredVersion)
	m.InterfaceLabels = flex.FlattenFrameworkListString(ctx, from.InterfaceLabels, diags)
	m.Name = flex.FlattenString(from.Name)
	m.PoolId = flex.FlattenString(from.PoolId)
	m.ServiceType = flex.FlattenString(from.ServiceType)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
	m.UpdatedAt = flex.FlattenRFC3339(from.UpdatedAt)
}
