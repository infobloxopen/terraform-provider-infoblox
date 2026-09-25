package dynamicallocation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
)

type NextAvailableVlanIdModel struct {
	VlanView     basetypes.StringValue `tfsdk:"vlan_view"`
	FilterParams basetypes.MapValue    `tfsdk:"filter_params"`
}

var NextAvailableVlanIdAttrTypes = map[string]attr.Type{
	"vlan_view":     basetypes.StringType{},
	"filter_params": basetypes.MapType{ElemType: basetypes.StringType{}},
}

var NextAvailableVlanIdResourceSchemaAttributes = map[string]schema.Attribute{
	"vlan_view": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString("default"),
		Validators: []validator.String{
			stringvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("filter_params"),
			),
		},
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplace(),
		},
		MarkdownDescription: "The name of the VLAN View to allocate the next available VLAN ID from. Defaults to the default VLAN View when omitted. Mutually exclusive with \"filter_params\".",
	},
	"filter_params": schema.MapAttribute{
		Optional:    true,
		ElementType: basetypes.StringType{},
		PlanModifiers: []planmodifier.Map{
			mapplanmodifier.RequiresReplace(),
		},
		MarkdownDescription: "Extensible-attribute filters used to select the VLAN View to allocate from (e.g. {\"*Site\" = \"location-1\"}). Mutually exclusive with \"vlan_view\".",
	},
}

func (m NextAvailableVlanIdModel) FuncCall(ctx context.Context, attributeName string, object string, diags *diag.Diagnostics) *niosipam.FuncCall {
	fc := &niosipam.FuncCall{}
	fc.SetAttributeName(attributeName)
	fc.SetObject(object)
	fc.SetObjectFunction("next_available_vlan_id")
	fc.SetResultField("vlan_ids")

	objectParams := map[string]any{}
	if !m.FilterParams.IsNull() && !m.FilterParams.IsUnknown() {
		var filter map[string]string
		diags.Append(m.FilterParams.ElementsAs(ctx, &filter, false)...)
		for k, v := range filter {
			objectParams[k] = v
		}
	} else if !m.VlanView.IsNull() && !m.VlanView.IsUnknown() {
		objectParams["name"] = m.VlanView.ValueString()
	}
	fc.SetObjectParameters(objectParams)
	return fc
}
