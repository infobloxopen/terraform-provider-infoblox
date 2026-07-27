package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
}

var UDDIRecordAaaaOptionsAttrTypes = map[string]attr.Type{}

var UDDIRecordAaaaOptionsResourceSchemaAttributes = map[string]schema.Attribute{}

func ExpandUDDIRecordAaaaOptions(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]any {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UDDIRecordAaaaOptionsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	to := map[string]any{}
	// TODO: populate `to` from m's subfields.
	return to
}

func FlattenUDDIRecordAaaaOptions(ctx context.Context, from map[string]any, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UDDIRecordAaaaOptionsAttrTypes)
	}
	m := UDDIRecordAaaaOptionsModel{}
	// TODO: populate m from `from` using the flex map-boundary adapters, e.g.
	//   flex.FlattenStringPointer(flex.RDataStringPtr(from["<key>"]))
	//   flex.FlattenBoolPointer(flex.RDataBoolPtr(from["<key>"]))
	//   flex.FlattenInt64Pointer(flex.RDataInt64Ptr(from["<key>"]))
	//   flex.FlattenIPv4Address(flex.RDataStringPtr(from["<key>"]))
	obj, d := types.ObjectValueFrom(ctx, UDDIRecordAaaaOptionsAttrTypes, m)
	diags.Append(d...)
	return obj
}

type UDDIRecordAaaaRdataModel struct {
}

var UDDIRecordAaaaRdataAttrTypes = map[string]attr.Type{}

var UDDIRecordAaaaRdataResourceSchemaAttributes = map[string]schema.Attribute{}

func ExpandUDDIRecordAaaaRdata(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]any {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UDDIRecordAaaaRdataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	to := map[string]any{}
	// TODO: populate `to` from m's subfields.
	return to
}

func FlattenUDDIRecordAaaaRdata(ctx context.Context, from map[string]any, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UDDIRecordAaaaRdataAttrTypes)
	}
	m := UDDIRecordAaaaRdataModel{}
	// TODO: populate m from `from` using the flex map-boundary adapters, e.g.
	//   flex.FlattenStringPointer(flex.RDataStringPtr(from["<key>"]))
	//   flex.FlattenBoolPointer(flex.RDataBoolPtr(from["<key>"]))
	//   flex.FlattenInt64Pointer(flex.RDataInt64Ptr(from["<key>"]))
	//   flex.FlattenIPv4Address(flex.RDataStringPtr(from["<key>"]))
	obj, d := types.ObjectValueFrom(ctx, UDDIRecordAaaaRdataAttrTypes, m)
	diags.Append(d...)
	return obj
}
