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

// ValidateRecordNs validates the RecordNs configuration.
func ValidateRecordNs(ctx context.Context, data RecordNsModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordNsModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordNsNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIRecordNsModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateRecordNsUDDIConfig(ctx, uddi, resp)
	}
}

func validateRecordNsNIOSConfig(ctx context.Context, m *NIOSRecordNsModel, resp *resource.ValidateConfigResponse) {
}

func validateRecordNsUDDIConfig(ctx context.Context, m *UDDIRecordNsModel, resp *resource.ValidateConfigResponse) {
}

type UDDIRecordNsRdataModel struct {
	Dname types.String `tfsdk:"dname"`
}

var UDDIRecordNsRdataAttrTypes = map[string]attr.Type{
	"dname": types.StringType,
}

var UDDIRecordNsRdataResourceSchemaAttributes = map[string]schema.Attribute{
	"dname": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The domain name of the authoritative name server for the zone.",
	},
}

func ExpandUDDIRecordNsRdata(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]any {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UDDIRecordNsRdataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	rdata := make(map[string]any)
	if dname := flex.ExpandString(m.Dname); dname != "" {
		rdata["dname"] = dname
	}
	return rdata
}

func FlattenUDDIRecordNsRdata(ctx context.Context, from map[string]any, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UDDIRecordNsRdataAttrTypes)
	}
	m := UDDIRecordNsRdataModel{
		Dname: flex.FlattenStringPointer(flex.RDataStringPtr(from["dname"])),
	}
	obj, d := types.ObjectValueFrom(ctx, UDDIRecordNsRdataAttrTypes, m)
	diags.Append(d...)
	return obj
}
