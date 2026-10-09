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

// ValidateRecordCname validates the RecordCname configuration.
func ValidateRecordCname(ctx context.Context, data RecordCnameModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordCnameModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordCnameNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIRecordCnameModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateRecordCnameUDDIConfig(ctx, uddi, resp)
	}
}

func validateRecordCnameNIOSConfig(ctx context.Context, m *NIOSRecordCnameModel, resp *resource.ValidateConfigResponse) {
}

func validateRecordCnameUDDIConfig(ctx context.Context, m *UDDIRecordCnameModel, resp *resource.ValidateConfigResponse) {
}

type UDDIRecordCnameRdataModel struct {
	Cname types.String `tfsdk:"cname"`
}

var UDDIRecordCnameRdataAttrTypes = map[string]attr.Type{
	"cname": types.StringType,
}

var UDDIRecordCnameRdataResourceSchemaAttributes = map[string]schema.Attribute{
	"cname": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "A domain name which specifies the canonical or primary name for the owner. The owner name is an alias. Can be empty.",
	},
}

func ExpandUDDIRecordCnameRdata(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]any {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UDDIRecordCnameRdataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	rdata := make(map[string]any)
	rdata["cname"] = flex.ExpandString(m.Cname)
	return rdata
}

func FlattenUDDIRecordCnameRdata(ctx context.Context, from map[string]any, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UDDIRecordCnameRdataAttrTypes)
	}
	m := UDDIRecordCnameRdataModel{
		Cname: flex.FlattenStringPointer(flex.RDataStringPtr(from["cname"])),
	}
	obj, d := types.ObjectValueFrom(ctx, UDDIRecordCnameRdataAttrTypes, m)
	diags.Append(d...)
	return obj
}
