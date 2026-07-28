package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/dynamicallocation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateRecordAaaa validates the RecordAaaa configuration.
func ValidateRecordAaaa(ctx context.Context, data RecordAaaaModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordAaaaModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordAaaaNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIRecordAaaaModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateRecordAaaaUDDIConfig(ctx, uddi, resp)
	}
}

func validateRecordAaaaNIOSConfig(ctx context.Context, m *NIOSRecordAaaaModel, resp *resource.ValidateConfigResponse) {
}

func validateRecordAaaaUDDIConfig(ctx context.Context, m *UDDIRecordAaaaModel, resp *resource.ValidateConfigResponse) {
}

func BuildRecordAaaaFuncCall(ctx context.Context, data types.Object, diags *diag.Diagnostics) *niosdns.FuncCall {
	if data.IsNull() || data.IsUnknown() {
		return nil
	}

	var m dynamicallocation.NextAvailableIpModel
	diags.Append(data.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.FuncCall(ctx, "Ipv6addr", "ipv6network", diags)
}

type UDDIRecordAaaaOptionsModel struct {
	CreatePtr types.Bool `tfsdk:"create_ptr"`
	CheckRmz  types.Bool `tfsdk:"check_rmz"`
}

var UDDIRecordAaaaOptionsAttrTypes = map[string]attr.Type{
	"create_ptr": types.BoolType,
	"check_rmz":  types.BoolType,
}

var UDDIRecordAaaaOptionsResourceSchemaAttributes = map[string]schema.Attribute{
	"create_ptr": schema.BoolAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.Bool{
			boolplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "A boolean flag which can be set to true to automatically create the corresponding PTR record.",
	},
	"check_rmz": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "A boolean flag which can be set to true to check the existence of the reverse zone for creating the corresponding PTR record. Only applicable if create_ptr is true.",
	},
}

func ExpandUDDIRecordAaaaOptions(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]any {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UDDIRecordAaaaOptionsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	opts := make(map[string]any)
	opts["create_ptr"] = flex.ExpandBool(m.CreatePtr)
	opts["check_rmz"] = flex.ExpandBool(m.CheckRmz)
	return opts
}

func FlattenUDDIRecordAaaaOptions(ctx context.Context, from map[string]any, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UDDIRecordAaaaOptionsAttrTypes)
	}
	m := UDDIRecordAaaaOptionsModel{
		CreatePtr: flex.FlattenBoolPointer(flex.RDataBoolPtr(from["create_ptr"])),
		CheckRmz:  flex.FlattenBoolPointer(flex.RDataBoolPtr(from["check_rmz"])),
	}
	obj, d := types.ObjectValueFrom(ctx, UDDIRecordAaaaOptionsAttrTypes, m)
	diags.Append(d...)
	return obj
}

type UDDIRecordAaaaRdataModel struct {
	Address iptypes.IPv6Address `tfsdk:"address"`
}

var UDDIRecordAaaaRdataAttrTypes = map[string]attr.Type{
	"address": iptypes.IPv6AddressType{},
}

var UDDIRecordAaaaRdataResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required:            true,
		CustomType:          iptypes.IPv6AddressType{},
		MarkdownDescription: "The IPv6 address of the host.",
	},
}

func ExpandUDDIRecordAaaaRdata(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]any {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UDDIRecordAaaaRdataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	rdata := make(map[string]any)
	if addr := flex.ExpandIPv6Address(m.Address); addr != nil {
		rdata["address"] = *addr
	}
	return rdata
}

func FlattenUDDIRecordAaaaRdata(ctx context.Context, from map[string]any, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UDDIRecordAaaaRdataAttrTypes)
	}
	m := UDDIRecordAaaaRdataModel{
		Address: flex.FlattenIPv6Address(flex.RDataStringPtr(from["address"])),
	}
	obj, d := types.ObjectValueFrom(ctx, UDDIRecordAaaaRdataAttrTypes, m)
	diags.Append(d...)
	return obj
}
