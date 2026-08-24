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

// ValidateRecordMx validates the RecordMx configuration.
func ValidateRecordMx(ctx context.Context, data RecordMxModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordMxModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordMxNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIRecordMxModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateRecordMxUDDIConfig(ctx, uddi, resp)
	}
}

func validateRecordMxNIOSConfig(ctx context.Context, m *NIOSRecordMxModel, resp *resource.ValidateConfigResponse) {
}

func validateRecordMxUDDIConfig(ctx context.Context, m *UDDIRecordMxModel, resp *resource.ValidateConfigResponse) {
}

type UDDIRecordMxRdataModel struct {
	Exchange   types.String `tfsdk:"exchange"`
	Preference types.Int64  `tfsdk:"preference"`
}

var UDDIRecordMxRdataAttrTypes = map[string]attr.Type{
	"exchange":   types.StringType,
	"preference": types.Int64Type,
}

var UDDIRecordMxRdataResourceSchemaAttributes = map[string]schema.Attribute{
	"exchange": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "A domain name which specifies a host willing to act as a mail exchange for the owner name.",
	},
	"preference": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "An unsigned 16-bit integer which specifies the preference given to this RR among others at the same owner. Lower values are preferred. The range of the value is 0 to 65535.",
	},
}

func ExpandUDDIRecordMxRdata(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]any {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UDDIRecordMxRdataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	rdata := map[string]any{
		"exchange":   flex.ExpandString(m.Exchange),
		"preference": flex.ExpandInt64(m.Preference),
	}
	return rdata
}

func FlattenUDDIRecordMxRdata(ctx context.Context, from map[string]any, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UDDIRecordMxRdataAttrTypes)
	}
	m := UDDIRecordMxRdataModel{
		Exchange:   flex.FlattenStringPointer(flex.RDataStringPtr(from["exchange"])),
		Preference: flex.FlattenInt64Pointer(flex.RDataInt64Ptr(from["preference"])),
	}
	obj, d := types.ObjectValueFrom(ctx, UDDIRecordMxRdataAttrTypes, m)
	diags.Append(d...)
	return obj
}
