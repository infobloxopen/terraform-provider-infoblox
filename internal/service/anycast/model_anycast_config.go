package anycast

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/anycast"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type AnycastConfigModel struct {
	Id            types.Int64  `tfsdk:"id"`
	UpdateTrigger types.String `tfsdk:"update_trigger"`
	UDDI          types.Object `tfsdk:"uddi"`
}

var AnycastConfigAttrTypes = map[string]attr.Type{
	"id":             types.Int64Type,
	"update_trigger": types.StringType,
	"uddi":           types.ObjectType{AttrTypes: UDDIAnycastConfigAttrTypes},
}

type UDDIAnycastConfigModel struct {
	AnycastIpAddress   iptypes.IPv4Address `tfsdk:"anycast_ip_address"`
	AnycastIpv6Address iptypes.IPv6Address `tfsdk:"anycast_ipv6_address"`
	Description        types.String        `tfsdk:"description"`
	Name               types.String        `tfsdk:"name"`
	OnpremHosts        types.List          `tfsdk:"onprem_hosts"`
	Service            types.String        `tfsdk:"service"`
	Tags               types.Map           `tfsdk:"tags"`
	TagsAll            types.Map           `tfsdk:"tags_all"`
}

var UDDIAnycastConfigAttrTypes = map[string]attr.Type{
	"anycast_ip_address":   iptypes.IPv4AddressType{},
	"anycast_ipv6_address": iptypes.IPv6AddressType{},
	"description":          types.StringType,
	"name":                 types.StringType,
	"onprem_hosts":         types.ListType{ElemType: types.ObjectType{AttrTypes: OnpremHostRefAttrTypes}},
	"service":              types.StringType,
	"tags":                 types.MapType{ElemType: types.StringType},
	"tags_all":             types.MapType{ElemType: types.StringType},
}

const (
	AnycastConfigReturnFields = ""
)

var AnycastConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"update_trigger": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "An arbitrary value used to trigger an update. Not sent to the API. Change it when Terraform reports no infrastructure changes.",
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          AnycastConfigResourceUddiSchemaAttributes,
	},
}

var AnycastConfigResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"anycast_ip_address": schema.StringAttribute{
		Required:            true,
		CustomType:          iptypes.IPv4AddressType{},
		MarkdownDescription: "IPv4 address of the host in string format.",
	},
	"anycast_ipv6_address": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		CustomType:          iptypes.IPv6AddressType{},
		MarkdownDescription: "IPv6 address of the host in string format",
	},
	"description": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The description for the address object. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the anycast configuration.",
	},
	"onprem_hosts": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: OnpremHostRefResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Struct on-prem host reference.",
	},
	"service": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("DNS", "DFP", "NTP"),
		},
		Required:            true,
		MarkdownDescription: "The type of the Service used in anycast configuration, supports (`dns`, `ntp`, `dfp`).",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The tags for the anycast configuration object.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *AnycastConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.AnycastConfig {
	if m == nil {
		return nil
	}

	obj := &coremodel.AnycastConfig{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIAnycastConfigModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIAnycastConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIAnycastConfigExt {
	return &coremodel.UDDIAnycastConfigExt{
		AnycastIpAddress:   flex.ExpandIPv4Address(m.AnycastIpAddress),
		AnycastIpv6Address: flex.ExpandIPv6Address(m.AnycastIpv6Address),
		Description:        flex.ExpandStringPointer(m.Description),
		Name:               flex.ExpandStringPointer(m.Name),
		OnpremHosts:        flex.ExpandFrameworkListNestedBlock(ctx, m.OnpremHosts, diags, ExpandOnpremHostRef),
		Service:            flex.ExpandStringPointer(m.Service),
		Tags:               flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
}

// Flatten populates the TF model from a core response.
func (m *AnycastConfigModel) Flatten(ctx context.Context, resp *coremodel.AnycastConfig, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenInt64Pointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIAnycastConfigModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIAnycastConfigModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIAnycastConfigAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIAnycastConfigAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIAnycastConfigModel) Flatten(ctx context.Context, from *coremodel.UDDIAnycastConfigExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AnycastIpAddress = flex.FlattenIPv4Address(from.AnycastIpAddress)
	m.AnycastIpv6Address = flex.FlattenIPv6Address(from.AnycastIpv6Address)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.OnpremHosts = flex.FlattenFrameworkListNestedBlock(ctx, from.OnpremHosts, OnpremHostRefAttrTypes, diags, FlattenOnpremHostRef)
	m.Service = flex.FlattenStringPointer(from.Service)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
}
