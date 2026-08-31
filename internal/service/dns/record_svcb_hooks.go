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

// ValidateRecordSvcb validates the RecordSvcb configuration.
func ValidateRecordSvcb(ctx context.Context, data RecordSvcbModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIRecordSvcbModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateRecordSvcbUDDIConfig(ctx, uddi, resp)
	}
}

func validateRecordSvcbUDDIConfig(ctx context.Context, m *UDDIRecordSvcbModel, resp *resource.ValidateConfigResponse) {
}

func PostFlattenRecordSvcbUDDI(ctx context.Context, planned, flattened *UDDIRecordSvcbModel, diags *diag.Diagnostics) {
}

type UDDIRecordSvcbRdataModel struct {
}

var UDDIRecordSvcbRdataAttrTypes = map[string]attr.Type{}

var UDDIRecordSvcbRdataResourceSchemaAttributes = map[string]schema.Attribute{}

func ExpandUDDIRecordSvcbRdata(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]any {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UDDIRecordSvcbRdataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	to := map[string]any{}
	// TODO: populate `to` from m's subfields.
	return to
}

func FlattenUDDIRecordSvcbRdata(ctx context.Context, from map[string]any, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UDDIRecordSvcbRdataAttrTypes)
	}
	m := UDDIRecordSvcbRdataModel{}
	// TODO: populate m from `from` using the flex map-boundary adapters, e.g.
	//   flex.FlattenStringPointer(flex.RDataStringPtr(from["<key>"]))
	//   flex.FlattenBoolPointer(flex.RDataBoolPtr(from["<key>"]))
	//   flex.FlattenInt64Pointer(flex.RDataInt64Ptr(from["<key>"]))
	//   flex.FlattenIPv4Address(flex.RDataStringPtr(from["<key>"]))
	obj, d := types.ObjectValueFrom(ctx, UDDIRecordSvcbRdataAttrTypes, m)
	diags.Append(d...)
	return obj
}
