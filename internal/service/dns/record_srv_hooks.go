package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateRecordSrv validates the RecordSrv configuration.
func ValidateRecordSrv(ctx context.Context, data RecordSrvModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordSrvModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordSrvNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIRecordSrvModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateRecordSrvUDDIConfig(ctx, uddi, resp)
	}
}

func validateRecordSrvNIOSConfig(ctx context.Context, m *NIOSRecordSrvModel, resp *resource.ValidateConfigResponse) {
}

func validateRecordSrvUDDIConfig(ctx context.Context, m *UDDIRecordSrvModel, resp *resource.ValidateConfigResponse) {
}

type UDDIRecordSRVRdataModel struct {
	Port     types.Int64  `tfsdk:"port"`
	Priority types.Int64  `tfsdk:"priority"`
	Target   types.String `tfsdk:"target"`
	Weight   types.Int64  `tfsdk:"weight"`
}

var UDDIRecordSRVRdataAttrTypes = map[string]attr.Type{
	"port":     types.Int64Type,
	"priority": types.Int64Type,
	"target":   types.StringType,
	"weight":   types.Int64Type,
}

var UDDIRecordSRVRdataResourceSchemaAttributes = map[string]schema.Attribute{
	"port": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "The port on this target host of the service. Valid values are 0 to 65535.",
	},
	"priority": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "The priority of the SRV record. Valid values are 0 to 65535.",
	},
	"target": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The target host domain name for the SRV record.",
	},
	"weight": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(0),
		MarkdownDescription: "The weight of the SRV record. Valid values are 0 to 65535.",
	},
}

func ExpandUDDIRecordSRVRdata(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]any {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UDDIRecordSRVRdataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	to := map[string]any{}
	to["port"] = flex.ExpandInt64(m.Port)
	to["priority"] = flex.ExpandInt64(m.Priority)
	to["target"] = flex.ExpandString(m.Target)
	to["weight"] = flex.ExpandInt64(m.Weight)
	return to
}

func FlattenUDDIRecordSRVRdata(ctx context.Context, from map[string]any, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UDDIRecordSRVRdataAttrTypes)
	}
	m := UDDIRecordSRVRdataModel{
		Port:     flex.FlattenInt64Pointer(flex.RDataInt64Ptr(from["port"])),
		Priority: flex.FlattenInt64Pointer(flex.RDataInt64Ptr(from["priority"])),
		Target:   flex.FlattenStringPointer(flex.RDataStringPtr(from["target"])),
		Weight:   flex.FlattenInt64Pointer(flex.RDataInt64Ptr(from["weight"])),
	}
	obj, d := types.ObjectValueFrom(ctx, UDDIRecordSRVRdataAttrTypes, m)
	diags.Append(d...)
	return obj
}
