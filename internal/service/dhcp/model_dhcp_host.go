package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

type DhcpHostModel struct {
	Id   types.String `tfsdk:"id"`
	UDDI types.Object `tfsdk:"uddi"`
}

var DhcpHostAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"uddi": types.ObjectType{AttrTypes: UDDIDhcpHostAttrTypes},
}

type UDDIDhcpHostModel struct {
	Address          types.String `tfsdk:"address"`
	AnycastAddresses types.List   `tfsdk:"anycast_addresses"`
	Comment          types.String `tfsdk:"comment"`
	CurrentVersion   types.String `tfsdk:"current_version"`
	IpSpace          types.String `tfsdk:"ip_space"`
	Name             types.String `tfsdk:"name"`
	Ophid            types.String `tfsdk:"ophid"`
	ProviderId       types.String `tfsdk:"provider_id"`
	Server           types.String `tfsdk:"server"`
	Tags             types.Map    `tfsdk:"tags"`
	TagsAll          types.Map    `tfsdk:"tags_all"`
	Type             types.String `tfsdk:"type"`
}

var UDDIDhcpHostAttrTypes = map[string]attr.Type{
	"address":           types.StringType,
	"anycast_addresses": types.ListType{ElemType: types.StringType},
	"comment":           types.StringType,
	"current_version":   types.StringType,
	"ip_space":          types.StringType,
	"name":              types.StringType,
	"ophid":             types.StringType,
	"provider_id":       types.StringType,
	"server":            types.StringType,
	"tags":              types.MapType{ElemType: types.StringType},
	"tags_all":          types.MapType{ElemType: types.StringType},
	"type":              types.StringType,
}

const (
	DhcpHostReturnFields = ""
)

var DhcpHostResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The resource identifier of the on-prem host. Must be provided because DHCP hosts are system-managed objects that pre-exist and cannot be created via API.",
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          DhcpHostResourceUddiSchemaAttributes,
	},
}

var DhcpHostResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The primary IP address of the on-prem host.",
	},
	"anycast_addresses": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "Anycast address configured to the host. Order is not significant.",
	},
	"comment": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The description of the on-prem host.",
	},
	"current_version": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Current DHCP application version of the host.",
	},
	"ip_space": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier of the IP space associated with the host.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The display name of the on-prem host.",
	},
	"ophid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The on-prem host ID.",
	},
	"provider_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "External provider identifier.",
	},
	"server": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The resource identifier of the DHCP Config Profile to associate with this host.",
	},
	"tags": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "The tags of the on-prem host in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
	"type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Defines the type of host. Allowed values:  * _bloxone_ddi_: host type is BloxOne DDI,  * _microsoft_azure_: host type is Microsoft Azure,  * _amazon_web_service_: host type is Amazon Web Services.  * _microsoft_active_directory_: host type is Microsoft Active Directory.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *DhcpHostModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.DhcpHost {
	if m == nil {
		return nil
	}

	obj := &coremodel.DhcpHost{
		Id: flex.ExpandStringPointer(m.Id),
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIDhcpHostModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIDhcpHostModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIDhcpHostExt {
	return &coremodel.UDDIDhcpHostExt{
		IpSpace: flex.ExpandStringPointer(m.IpSpace),
		Server:  flex.ExpandStringPointer(m.Server),
		Tags:    flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
}

// Flatten populates the TF model from a core response.
func (m *DhcpHostModel) Flatten(ctx context.Context, resp *coremodel.DhcpHost, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIDhcpHostModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIDhcpHostModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIDhcpHostAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIDhcpHostAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIDhcpHostModel) Flatten(ctx context.Context, from *coremodel.UDDIDhcpHostExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.AnycastAddresses = flex.FlattenFrameworkListString(ctx, from.AnycastAddresses, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CurrentVersion = flex.FlattenStringPointer(from.CurrentVersion)
	m.IpSpace = flex.FlattenStringPointer(from.IpSpace)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Ophid = flex.FlattenStringPointer(from.Ophid)
	m.ProviderId = flex.FlattenStringPointer(from.ProviderId)
	m.Server = flex.FlattenStringPointer(from.Server)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
	m.Type = flex.FlattenStringPointer(from.Type)
}
