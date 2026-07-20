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

type NextAvailableNetworkModel struct {
	Network      basetypes.StringValue `tfsdk:"network"`
	NetworkView  basetypes.StringValue `tfsdk:"network_view"`
	Cidr         basetypes.Int64Value  `tfsdk:"cidr"`
	Exclude      basetypes.ListValue   `tfsdk:"exclude"`
	FilterParams basetypes.MapValue    `tfsdk:"filter_params"`
}

var NextAvailableNetworkAttrTypes = map[string]attr.Type{
	"network":       basetypes.StringType{},
	"network_view":  basetypes.StringType{},
	"cidr":          basetypes.Int64Type{},
	"exclude":       basetypes.ListType{ElemType: basetypes.StringType{}},
	"filter_params": basetypes.MapType{ElemType: basetypes.StringType{}},
}

var NextAvailableNetworkResourceSchemaAttributes = map[string]schema.Attribute{
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

func (m NextAvailableNetworkModel) FuncCall(ctx context.Context, attributeName string, object string, diags *diag.Diagnostics) *niosipam.FuncCall {
	fc := &niosipam.FuncCall{}
	fc.SetAttributeName(attributeName)
	fc.SetObject(object)
	fc.SetObjectFunction("next_available_network")
	fc.SetResultField("networks")

	objectParams := map[string]any{}
	if !m.Network.IsNull() && !m.Network.IsUnknown() {
		objectParams["network"] = m.Network.ValueString()
	}
	if !m.NetworkView.IsNull() && !m.NetworkView.IsUnknown() {
		objectParams["network_view"] = m.NetworkView.ValueString()
	}
	if !m.FilterParams.IsNull() && !m.FilterParams.IsUnknown() {
		var filter map[string]string
		diags.Append(m.FilterParams.ElementsAs(ctx, &filter, false)...)
		for k, v := range filter {
			objectParams[k] = v
		}
	}
	fc.SetObjectParameters(objectParams)

	params := map[string]any{}
	if !m.Cidr.IsNull() && !m.Cidr.IsUnknown() {
		params["cidr"] = m.Cidr.ValueInt64()
	}
	if !m.Exclude.IsNull() && !m.Exclude.IsUnknown() {
		var exclude []string
		diags.Append(m.Exclude.ElementsAs(ctx, &exclude, false)...)
		if len(exclude) > 0 {
			excludeAny := make([]any, len(exclude))
			for i, v := range exclude {
				excludeAny[i] = v
			}
			params["exclude"] = excludeAny
		}
	}
	if len(params) > 0 {
		fc.SetParameters(params)
	}

	return fc
}
