package dynamicallocation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
)

type NextAvailableNetworkContainerModel struct {
	Network      basetypes.StringValue `tfsdk:"network"`
	NetworkView  basetypes.StringValue `tfsdk:"network_view"`
	Cidr         basetypes.Int64Value  `tfsdk:"cidr"`
	Exclude      basetypes.ListValue   `tfsdk:"exclude"`
	FilterParams basetypes.MapValue    `tfsdk:"filter_params"`
}

var NextAvailableNetworkContainerAttrTypes = map[string]attr.Type{
	"network":       basetypes.StringType{},
	"network_view":  basetypes.StringType{},
	"cidr":          basetypes.Int64Type{},
	"exclude":       basetypes.ListType{ElemType: basetypes.StringType{}},
	"filter_params": basetypes.MapType{ElemType: basetypes.StringType{}},
}

var NextAvailableNetworkContainerResourceSchemaAttributes = map[string]schema.Attribute{
	"network": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(
				path.MatchRelative().AtParent().AtName("filter_params"),
			),
		},
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplace(),
		},
		MarkdownDescription: "The network container to allocate the next available network from, in CIDR notation (e.g. \"10.0.0.0/16\"). Mutually exclusive with \"filter_params\".",
	},
	"network_view": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString("default"),
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplace(),
		},
		MarkdownDescription: "The network view of the network container. Defaults to the default network view when omitted.",
	},
	"cidr": schema.Int64Attribute{
		Required: true,
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.RequiresReplace(),
		},
		MarkdownDescription: "The prefix length (CIDR) of the network to allocate (e.g. 24 for a /24).",
	},
	"exclude": schema.ListAttribute{
		Optional:    true,
		ElementType: basetypes.StringType{},
		PlanModifiers: []planmodifier.List{
			listplanmodifier.RequiresReplace(),
		},
		MarkdownDescription: "A list of networks or ranges to exclude from allocation.",
	},
	"filter_params": schema.MapAttribute{
		Optional:    true,
		ElementType: basetypes.StringType{},
		PlanModifiers: []planmodifier.Map{
			mapplanmodifier.RequiresReplace(),
		},
		MarkdownDescription: "Extensible-attribute filters used to select the network container to allocate from (e.g. {\"*Site\" = \"location-1\"}). Mutually exclusive with \"network\".",
	},
}

func (m NextAvailableNetworkContainerModel) FuncCall(ctx context.Context, attributeName string, object string, diags *diag.Diagnostics) *niosipam.FuncCall {
	return NextAvailableNetworkModel{
		Network:      m.Network,
		NetworkView:  m.NetworkView,
		Cidr:         m.Cidr,
		Exclude:      m.Exclude,
		FilterParams: m.FilterParams,
	}.FuncCall(ctx, attributeName, object, diags)
}
