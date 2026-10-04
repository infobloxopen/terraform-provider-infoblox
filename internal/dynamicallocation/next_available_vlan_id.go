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
	VlanRange    basetypes.StringValue `tfsdk:"vlan_range"`
	FilterObject basetypes.StringValue `tfsdk:"filter_object"`
	FilterParams basetypes.MapValue    `tfsdk:"filter_params"`
}

var NextAvailableVlanIdAttrTypes = map[string]attr.Type{
	"vlan_view":     basetypes.StringType{},
	"vlan_range":    basetypes.StringType{},
	"filter_object": basetypes.StringType{},
	"filter_params": basetypes.MapType{ElemType: basetypes.StringType{}},
}

var NextAvailableVlanIdResourceSchemaAttributes = map[string]schema.Attribute{
	"vlan_view": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			stringvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("vlan_range"),
			),
		},
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplace(),
		},
		MarkdownDescription: "The name of the VLAN View to allocate the next available VLAN ID from.",
	},
	"vlan_range": schema.StringAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplace(),
		},
		MarkdownDescription: "The name of the VLAN Range to allocate the next available VLAN ID from.",
	},
	"filter_object": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("vlanview"),
		Validators: []validator.String{
			stringvalidator.OneOf("vlanview", "vlanrange"),
		},
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplace(),
		},
		MarkdownDescription: "Whether \"filter_params\" searches VLAN Views or VLAN Ranges. Valid values are \"vlanview\" and \"vlanrange\". Defaults to \"vlanview\".",
	},
	"filter_params": schema.MapAttribute{
		Optional:    true,
		ElementType: basetypes.StringType{},
		PlanModifiers: []planmodifier.Map{
			mapplanmodifier.RequiresReplace(),
		},
		MarkdownDescription: "Extensible Attribute filters used to select the VLAN View or VLAN Range to allocate from (e.g. {\"*Site\" = \"location-1\"}). The object type searched is set by \"filter_object\".",
	},
}

func (m NextAvailableVlanIdModel) FuncCall(ctx context.Context, attributeName string, object string, diags *diag.Diagnostics) *niosipam.FuncCall {
	fc := &niosipam.FuncCall{}
	fc.SetAttributeName(attributeName)
	fc.SetObjectFunction("next_available_vlan_id")
	fc.SetResultField("vlan_ids")

	filtering := !m.FilterParams.IsNull() && !m.FilterParams.IsUnknown()

	allocFrom := object
	switch {
	case filtering && !m.FilterObject.IsNull() && !m.FilterObject.IsUnknown():
		allocFrom = m.FilterObject.ValueString()
	case !m.VlanRange.IsNull() && !m.VlanRange.IsUnknown():
		allocFrom = "vlanrange"
	}

	fc.SetObject(allocFrom)

	objectParams := map[string]any{}
	if filtering {
		var filter map[string]string
		diags.Append(m.FilterParams.ElementsAs(ctx, &filter, false)...)
		for k, v := range filter {
			objectParams[k] = v
		}
	} else if allocFrom == "vlanrange" {
		objectParams["name"] = m.VlanRange.ValueString()
	} else if !m.VlanView.IsNull() && !m.VlanView.IsUnknown() {
		objectParams["name"] = m.VlanView.ValueString()
	}
	fc.SetObjectParameters(objectParams)

	return fc
}
