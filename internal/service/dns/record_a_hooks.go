package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/dynamicallocation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// UDDIRecordARdataModel is the typed rdata for an A (Address) record.
type UDDIRecordARdataModel struct {
	Address iptypes.IPv4Address `tfsdk:"address"`
}

var UDDIRecordARdataAttrTypes = map[string]attr.Type{
	"address": iptypes.IPv4AddressType{},
}

var UDDIRecordARdataResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required:            true,
		CustomType:          iptypes.IPv4AddressType{},
		MarkdownDescription: "The IPv4 address of the host.",
	},
}

func ExpandUDDIRecordARdata(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]interface{} {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UDDIRecordARdataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	rdata := make(map[string]interface{})
	if addr := flex.ExpandIPv4Address(m.Address); addr != nil {
		rdata["address"] = *addr
	}
	return rdata
}

func FlattenUDDIRecordARdata(ctx context.Context, from map[string]interface{}, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UDDIRecordARdataAttrTypes)
	}
	m := UDDIRecordARdataModel{
		Address: flex.FlattenIPv4Address(flex.RDataStringPtr(from["address"])),
	}
	obj, d := types.ObjectValueFrom(ctx, UDDIRecordARdataAttrTypes, m)
	diags.Append(d...)
	return obj
}

// UDDIRecordAOptionsModel is the typed options block for an A record.
type UDDIRecordAOptionsModel struct {
	CreatePtr types.Bool `tfsdk:"create_ptr"`
	CheckRmz  types.Bool `tfsdk:"check_rmz"`
}

var UDDIRecordAOptionsAttrTypes = map[string]attr.Type{
	"create_ptr": types.BoolType,
	"check_rmz":  types.BoolType,
}

var UDDIRecordAOptionsResourceSchemaAttributes = map[string]schema.Attribute{
	"create_ptr": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "A boolean flag which can be set to true to automatically create the corresponding PTR record.",
	},
	"check_rmz": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "A boolean flag which can be set to true to check the existence of the reverse zone for creating the corresponding PTR record. Only applicable if create_ptr is true.",
	},
}

func ExpandUDDIRecordAOptions(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]interface{} {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UDDIRecordAOptionsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	opts := make(map[string]interface{})
	if !m.CreatePtr.IsNull() && !m.CreatePtr.IsUnknown() {
		opts["create_ptr"] = m.CreatePtr.ValueBool()
	}
	if !m.CheckRmz.IsNull() && !m.CheckRmz.IsUnknown() {
		opts["check_rmz"] = m.CheckRmz.ValueBool()
	}
	return opts
}

func FlattenUDDIRecordAOptions(ctx context.Context, from map[string]interface{}, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UDDIRecordAOptionsAttrTypes)
	}
	m := UDDIRecordAOptionsModel{
		CreatePtr: flex.FlattenBoolPointer(flex.RDataBoolPtr(from["create_ptr"])),
		CheckRmz:  flex.FlattenBoolPointer(flex.RDataBoolPtr(from["check_rmz"])),
	}
	obj, d := types.ObjectValueFrom(ctx, UDDIRecordAOptionsAttrTypes, m)
	diags.Append(d...)
	return obj
}

// ValidateRecordA validates the RecordA configuration.
func ValidateRecordA(ctx context.Context, data RecordAModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordAModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordANIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIRecordAModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateRecordAUDDIConfig(ctx, uddi, resp)
	}
}

func validateRecordANIOSConfig(ctx context.Context, m *NIOSRecordAModel, resp *resource.ValidateConfigResponse) {
}

func validateRecordAUDDIConfig(ctx context.Context, m *UDDIRecordAModel, resp *resource.ValidateConfigResponse) {
}

func BuildRecordAFuncCall(ctx context.Context, data types.Object, diags *diag.Diagnostics) *niosdns.FuncCall {
	if data.IsNull() || data.IsUnknown() {
		return nil
	}

	var m dynamicallocation.NextAvailableIpModel
	diags.Append(data.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.FuncCall(ctx, "Ipv4addr", "network", diags)
}

func PostExpandRecordAUDDI(ctx context.Context, ext *coremodel.UDDIRecordAExt, diags *diag.Diagnostics) *coremodel.UDDIRecordAExt {
	return ext
}

func PostFlattenRecordAUDDI(ctx context.Context, planned, flattened *UDDIRecordAModel, diags *diag.Diagnostics) {
	if flattened == nil {
		return
	}
	if planned != nil {
		flattened.Options = planned.Options
	} else {
		flattened.Options = types.ObjectNull(UDDIRecordAOptionsAttrTypes)
	}
}
