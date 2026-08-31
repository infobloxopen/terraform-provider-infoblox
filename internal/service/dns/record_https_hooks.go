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

// ValidateRecordHttps validates the RecordHttps configuration.
func ValidateRecordHttps(ctx context.Context, data RecordHttpsModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIRecordHttpsModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateRecordHttpsUDDIConfig(ctx, uddi, resp)
	}
}

func validateRecordHttpsUDDIConfig(ctx context.Context, m *UDDIRecordHttpsModel, resp *resource.ValidateConfigResponse) {
}

type UDDIRecordHttpsRdataModel struct {
	Priority   types.Int64  `tfsdk:"priority"`
	TargetName types.String `tfsdk:"target_name"`
}

var UDDIRecordHttpsRdataAttrTypes = map[string]attr.Type{
	"priority":    types.Int64Type,
	"target_name": types.StringType,
}

var UDDIRecordHttpsRdataResourceSchemaAttributes = map[string]schema.Attribute{
	"priority": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(0),
		MarkdownDescription: "An unsigned 16-bit integer in the range 0 to 65535 that indicates the priority of the HTTPS record. Lower values are preferred.",
	},
	"target_name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The domain name of the HTTPS target. Use \".\" to indicate the service is located at the owner name itself.",
	},
}

func ExpandUDDIRecordHttpsRdata(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]any {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UDDIRecordHttpsRdataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return map[string]any{
		"priority":    flex.ExpandInt64(m.Priority),
		"target_name": flex.ExpandString(m.TargetName),
	}
}

func FlattenUDDIRecordHttpsRdata(ctx context.Context, from map[string]any, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UDDIRecordHttpsRdataAttrTypes)
	}
	m := UDDIRecordHttpsRdataModel{
		Priority:   flex.FlattenInt64Pointer(flex.RDataInt64Ptr(from["priority"])),
		TargetName: flex.FlattenStringPointer(flex.RDataStringPtr(from["target_name"])),
	}
	obj, d := types.ObjectValueFrom(ctx, UDDIRecordHttpsRdataAttrTypes, m)
	diags.Append(d...)
	return obj
}
