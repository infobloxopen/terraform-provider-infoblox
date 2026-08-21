package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
	IpSpace types.String `tfsdk:"ip_space"`
	Server  types.String `tfsdk:"server"`
	Tags    types.Map    `tfsdk:"tags"`
	TagsAll types.Map    `tfsdk:"tags_all"`
}

var UDDIDhcpHostAttrTypes = map[string]attr.Type{
	"ip_space": types.StringType,
	"server":   types.StringType,
	"tags":     types.MapType{ElemType: types.StringType},
	"tags_all": types.MapType{ElemType: types.StringType},
}

const (
	DhcpHostReturnFields = ""
)

var DhcpHostResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The resource identifier of the on-prem host. Must be provided because DHCP hosts are system-managed objects that pre-exist and cannot be created via API.",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          DhcpHostResourceUddiSchemaAttributes,
	},
}

var DhcpHostResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"ip_space": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier of the IP space associated with the host.",
	},
	"server": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier of the DHCP Config Profile to associate with this host.",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
	},
	"tags": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "The tags of the on-prem host. Read-only — the DhcpHost Update API does not accept tag modifications.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
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
		// Tags are read-only in the DhcpHost API; never send them in Create/Update.
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
	m.IpSpace = flex.FlattenStringPointer(from.IpSpace)
	m.Server = flex.FlattenStringPointer(from.Server)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	m.Tags = tagsAll
	m.TagsAll = tagsAll
}
