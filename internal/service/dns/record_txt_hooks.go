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

// ValidateRecordTxt validates the RecordTxt configuration.
func ValidateRecordTxt(ctx context.Context, data RecordTxtModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordTxtModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordTxtNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIRecordTxtModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateRecordTxtUDDIConfig(ctx, uddi, resp)
	}
}

func validateRecordTxtNIOSConfig(ctx context.Context, m *NIOSRecordTxtModel, resp *resource.ValidateConfigResponse) {
}

func validateRecordTxtUDDIConfig(ctx context.Context, m *UDDIRecordTxtModel, resp *resource.ValidateConfigResponse) {
}

type UDDIRecordTxtRdataModel struct {
	Text types.String `tfsdk:"text"`
}

var UDDIRecordTxtRdataAttrTypes = map[string]attr.Type{
	"text": types.StringType,
}

var UDDIRecordTxtRdataResourceSchemaAttributes = map[string]schema.Attribute{
	"text": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The semantics of the text depends on the domain where it is found.",
	},
}

func ExpandUDDIRecordTxtRdata(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]any {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UDDIRecordTxtRdataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	rdata := make(map[string]any)
	rdata["text"] = flex.ExpandString(m.Text)
	return rdata
}

func FlattenUDDIRecordTxtRdata(ctx context.Context, from map[string]any, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UDDIRecordTxtRdataAttrTypes)
	}
	m := UDDIRecordTxtRdataModel{
		Text: flex.FlattenString(flex.RDataString(from["text"])),
	}
	obj, d := types.ObjectValueFrom(ctx, UDDIRecordTxtRdataAttrTypes, m)
	diags.Append(d...)
	return obj
}
