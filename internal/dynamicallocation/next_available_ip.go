package dynamicallocation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
)

type NextAvailableIpModel struct {
	Network      basetypes.StringValue `tfsdk:"network"`
	NetworkView  basetypes.StringValue `tfsdk:"network_view"`
	Exclude      basetypes.ListValue   `tfsdk:"exclude"`
	FilterParams basetypes.MapValue    `tfsdk:"filter_params"`
}

var NextAvailableIpAttrTypes = map[string]attr.Type{
	"network":       basetypes.StringType{},
	"network_view":  basetypes.StringType{},
	"exclude":       basetypes.ListType{ElemType: basetypes.StringType{}},
	"filter_params": basetypes.MapType{ElemType: basetypes.StringType{}},
}

var NextAvailableIpResourceSchemaAttributes = map[string]schema.Attribute{
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
		MarkdownDescription: "The network to allocate the next available address from, in CIDR notation (e.g. \"10.0.0.0/24\"). Mutually exclusive with \"filter_params\".",
	},
	"network_view": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString("default"),
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplace(),
		},
		MarkdownDescription: "The network view of the network. Defaults to the default network view when omitted.",
	},
	"exclude": schema.ListAttribute{
		Optional:    true,
		ElementType: basetypes.StringType{},
		PlanModifiers: []planmodifier.List{
			listplanmodifier.RequiresReplace(),
		},
		MarkdownDescription: "A list of IP addresses or ranges to exclude from allocation.",
	},
	"filter_params": schema.MapAttribute{
		Optional:    true,
		ElementType: basetypes.StringType{},
		PlanModifiers: []planmodifier.Map{
			mapplanmodifier.RequiresReplace(),
		},
		MarkdownDescription: "Extensible-attribute filters used to select the network to allocate from (e.g. {\"*Site\" = \"location-1\"}). Mutually exclusive with \"network\".",
	},
}

func (m NextAvailableIpModel) params(ctx context.Context, diags *diag.Diagnostics) (objectParams map[string]any, parameters map[string]any) {
	objectParams = map[string]any{}
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

	if !m.Exclude.IsNull() && !m.Exclude.IsUnknown() {
		var exclude []string
		diags.Append(m.Exclude.ElementsAs(ctx, &exclude, false)...)
		if len(exclude) > 0 {
			excludeAny := make([]any, len(exclude))
			for i, v := range exclude {
				excludeAny[i] = v
			}
			parameters = map[string]any{"exclude": excludeAny}
		}
	}

	return objectParams, parameters
}

func (m NextAvailableIpModel) FuncCall(ctx context.Context, attributeName string, object string, diags *diag.Diagnostics) *niosdns.FuncCall {
	fc := &niosdns.FuncCall{}
	fc.SetAttributeName(attributeName)
	fc.SetObject(object)
	fc.SetObjectFunction("next_available_ip")
	fc.SetResultField("ips")

	objectParams, parameters := m.params(ctx, diags)
	fc.SetObjectParameters(objectParams)
	if parameters != nil {
		fc.SetParameters(parameters)
	}

	return fc
}

func (m NextAvailableIpModel) FuncCallDHCP(ctx context.Context, attributeName string, object string, diags *diag.Diagnostics) *niosdhcp.FuncCall {
	fc := &niosdhcp.FuncCall{}
	fc.SetAttributeName(attributeName)
	fc.SetObject(object)
	fc.SetObjectFunction("next_available_ip")
	fc.SetResultField("ips")

	objectParams, parameters := m.params(ctx, diags)
	fc.SetObjectParameters(objectParams)
	if parameters != nil {
		fc.SetParameters(parameters)
	}

	return fc
}
