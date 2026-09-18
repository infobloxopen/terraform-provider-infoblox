package infra

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

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/infra"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

type InfraHostModel struct {
	Id   types.String `tfsdk:"id"`
	UDDI types.Object `tfsdk:"uddi"`
}

var InfraHostAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"uddi": types.ObjectType{AttrTypes: UDDIInfraHostAttrTypes},
}

type UDDIInfraHostModel struct {
	Description     types.String `tfsdk:"description"`
	DisplayName     types.String `tfsdk:"display_name"`
	IpSpace         types.String `tfsdk:"ip_space"`
	LocationId      types.String `tfsdk:"location_id"`
	MaintenanceMode types.String `tfsdk:"maintenance_mode"`
	PoolId          types.String `tfsdk:"pool_id"`
	SerialNumber    types.String `tfsdk:"serial_number"`
	Tags            types.Map    `tfsdk:"tags"`
	TagsAll         types.Map    `tfsdk:"tags_all"`
}

var UDDIInfraHostAttrTypes = map[string]attr.Type{
	"description":      types.StringType,
	"display_name":     types.StringType,
	"ip_space":         types.StringType,
	"location_id":      types.StringType,
	"maintenance_mode": types.StringType,
	"pool_id":          types.StringType,
	"serial_number":    types.StringType,
	"tags":             types.MapType{ElemType: types.StringType},
	"tags_all":         types.MapType{ElemType: types.StringType},
}

const (
	InfraHostReturnFields = ""
)

var InfraHostResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          InfraHostResourceUddiSchemaAttributes,
	},
}

var InfraHostResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"description": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The description of the Host (optional).",
	},
	"display_name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the Host (unique).",
	},
	"ip_space": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IP Space of the Host.",
	},
	"location_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"maintenance_mode": schema.StringAttribute{
		Default:             stringdefault.StaticString("disabled"),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"pool_id": schema.StringAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"serial_number": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The unique serial number of the Host.",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Tags associated with this Host.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *InfraHostModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.InfraHost {
	if m == nil {
		return nil
	}

	obj := &coremodel.InfraHost{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIInfraHostModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIInfraHostModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIInfraHostExt {
	return &coremodel.UDDIInfraHostExt{
		Description:     flex.ExpandStringPointer(m.Description),
		DisplayName:     flex.ExpandString(m.DisplayName),
		IpSpace:         flex.ExpandStringPointer(m.IpSpace),
		LocationId:      flex.ExpandStringPointer(m.LocationId),
		MaintenanceMode: flex.ExpandStringPointer(m.MaintenanceMode),
		PoolId:          flex.ExpandStringPointer(m.PoolId),
		SerialNumber:    flex.ExpandStringPointer(m.SerialNumber),
		Tags:            flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
}

// Flatten populates the TF model from a core response.
func (m *InfraHostModel) Flatten(ctx context.Context, resp *coremodel.InfraHost, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIInfraHostModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIInfraHostModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIInfraHostAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIInfraHostAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIInfraHostModel) Flatten(ctx context.Context, from *coremodel.UDDIInfraHostExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Description = flex.FlattenStringPointer(from.Description)
	m.DisplayName = flex.FlattenString(from.DisplayName)
	m.IpSpace = flex.FlattenStringPointer(from.IpSpace)
	m.LocationId = flex.FlattenStringPointer(from.LocationId)
	m.MaintenanceMode = flex.FlattenStringPointer(from.MaintenanceMode)
	m.PoolId = flex.FlattenStringPointer(from.PoolId)
	m.SerialNumber = flex.FlattenStringPointer(from.SerialNumber)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
}
