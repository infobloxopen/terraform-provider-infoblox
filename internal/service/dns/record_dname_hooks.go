package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

type UDDIRecordDnameRdataModel struct {
	Target types.String `tfsdk:"target"`
}

var UDDIRecordDnameRdataAttrTypes = map[string]attr.Type{
	"target": types.StringType,
}

var UDDIRecordDnameRdataResourceSchemaAttributes = map[string]schema.Attribute{
	"target": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The target domain name to which the zone will be mapped. Can be empty.",
	},
}

func ExpandUDDIRecordDnameRdata(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]any {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UDDIRecordDnameRdataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	rdata := make(map[string]any)
	rdata["target"] = flex.ExpandString(m.Target)
	return rdata
}

func FlattenUDDIRecordDnameRdata(ctx context.Context, from map[string]any, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UDDIRecordDnameRdataAttrTypes)
	}
	m := UDDIRecordDnameRdataModel{
		Target: flex.FlattenStringPointer(flex.RDataStringPtr(from["target"])),
	}
	obj, d := types.ObjectValueFrom(ctx, UDDIRecordDnameRdataAttrTypes, m)
	diags.Append(d...)
	return obj
}

// ValidateRecordDname validates the RecordDname configuration.
func ValidateRecordDname(ctx context.Context, data RecordDnameModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordDnameModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordDnameNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIRecordDnameModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateRecordDnameUDDIConfig(ctx, uddi, resp)
	}
}

func validateRecordDnameNIOSConfig(ctx context.Context, m *NIOSRecordDnameModel, resp *resource.ValidateConfigResponse) {
}

func validateRecordDnameUDDIConfig(ctx context.Context, m *UDDIRecordDnameModel, resp *resource.ValidateConfigResponse) {
}
