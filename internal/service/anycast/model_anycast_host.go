package anycast

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/anycast"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type AnycastHostModel struct {
	Id            types.Int64  `tfsdk:"id"`
	UpdateTrigger types.String `tfsdk:"update_trigger"`
	UDDI          types.Object `tfsdk:"uddi"`
}

var AnycastHostAttrTypes = map[string]attr.Type{
	"id":             types.Int64Type,
	"update_trigger": types.StringType,
	"uddi":           types.ObjectType{AttrTypes: UDDIAnycastHostAttrTypes},
}

type UDDIAnycastHostModel struct {
	AnycastConfigRefs types.List   `tfsdk:"anycast_config_refs"`
	ConfigBgp         types.Object `tfsdk:"config_bgp"`
	ConfigOspf        types.Object `tfsdk:"config_ospf"`
	ConfigOspfv3      types.Object `tfsdk:"config_ospfv3"`
	IpAddress         types.String `tfsdk:"ip_address"`
	Ipv6Address       types.String `tfsdk:"ipv6_address"`
	Name              types.String `tfsdk:"name"`
}

var UDDIAnycastHostAttrTypes = map[string]attr.Type{
	"anycast_config_refs": types.ListType{ElemType: types.ObjectType{AttrTypes: AnycastConfigRefAttrTypes}},
	"config_bgp":          types.ObjectType{AttrTypes: BgpConfigAttrTypes},
	"config_ospf":         types.ObjectType{AttrTypes: OspfConfigAttrTypes},
	"config_ospfv3":       types.ObjectType{AttrTypes: Ospfv3ConfigAttrTypes},
	"ip_address":          types.StringType,
	"ipv6_address":        types.StringType,
	"name":                types.StringType,
}

const (
	AnycastHostReturnFields = ""
)

var AnycastHostResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "",
	},
	"update_trigger": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "An arbitrary value used to trigger an update. Not sent to the API. Change it when Terraform reports no infrastructure changes.",
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          AnycastHostResourceUddiSchemaAttributes,
	},
}

var AnycastHostResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"anycast_config_refs": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: AnycastConfigRefResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Array of AnycastConfigRef structures, identifying the anycast configurations that this host is a member of.",
	},
	"config_bgp": schema.SingleNestedAttribute{
		Attributes:          BgpConfigResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Struct BGP configuration; defines BGP configuration for one anycast-enabled on-prem host.",
	},
	"config_ospf": schema.SingleNestedAttribute{
		Attributes:          OspfConfigResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Struct OSPF configuration; defines OSPF configuration for one anycast-enabled on-prem host.",
	},
	"config_ospfv3": schema.SingleNestedAttribute{
		Attributes:          Ospfv3ConfigResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Struct OSPFv3 configuration; defines OSPFv3 configuration for one anycast-enabled on-prem host.",
	},
	"ip_address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "IPv4 address of the on-prem host",
	},
	"ipv6_address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "IPv6 address of the on-prem host",
	},
	"name": schema.StringAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "User-friendly name of the host @example \"dns-host-1\", \"Central Office Server\".",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *AnycastHostModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.AnycastHost {
	if m == nil {
		return nil
	}

	obj := &coremodel.AnycastHost{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIAnycastHostModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIAnycastHostModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIAnycastHostExt {
	return &coremodel.UDDIAnycastHostExt{
		AnycastConfigRefs: flex.ExpandFrameworkListNestedBlock(ctx, m.AnycastConfigRefs, diags, ExpandAnycastConfigRef),
		ConfigBgp:         ExpandBgpConfig(ctx, m.ConfigBgp, diags),
		ConfigOspf:        ExpandOspfConfig(ctx, m.ConfigOspf, diags),
		ConfigOspfv3:      ExpandOspfv3Config(ctx, m.ConfigOspfv3, diags),
		IpAddress:         flex.ExpandStringPointer(m.IpAddress),
		Ipv6Address:       flex.ExpandStringPointer(m.Ipv6Address),
		Name:              flex.ExpandStringPointer(m.Name),
	}
}

// Flatten populates the TF model from a core response.
func (m *AnycastHostModel) Flatten(ctx context.Context, resp *coremodel.AnycastHost, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenInt64Pointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIAnycastHostModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIAnycastHostModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIAnycastHostAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIAnycastHostAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIAnycastHostModel) Flatten(ctx context.Context, from *coremodel.UDDIAnycastHostExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AnycastConfigRefs = flex.FlattenFrameworkListNestedBlock(ctx, from.AnycastConfigRefs, AnycastConfigRefAttrTypes, diags, FlattenAnycastConfigRef)
	m.ConfigBgp = FlattenBgpConfig(ctx, from.ConfigBgp, diags)
	m.ConfigOspf = FlattenOspfConfig(ctx, from.ConfigOspf, diags)
	m.ConfigOspfv3 = FlattenOspfv3Config(ctx, from.ConfigOspfv3, diags)
	m.IpAddress = flex.FlattenStringPointer(from.IpAddress)
	m.Ipv6Address = flex.FlattenStringPointer(from.Ipv6Address)
	m.Name = flex.FlattenStringPointer(from.Name)
}
