package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateDtcMonitorSnmp validates the DtcMonitorSnmp configuration.
func ValidateDtcMonitorSnmp(ctx context.Context, data DtcMonitorSnmpModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSDtcMonitorSnmpModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateDtcMonitorSnmpNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIDtcMonitorSnmpModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateDtcMonitorSnmpUDDIConfig(ctx, uddi, resp)
	}
}

func validateDtcMonitorSnmpNIOSConfig(ctx context.Context, m *NIOSDtcMonitorSnmpModel, resp *resource.ValidateConfigResponse) {
}

func validateDtcMonitorSnmpUDDIConfig(ctx context.Context, m *UDDIDtcMonitorSnmpModel, resp *resource.ValidateConfigResponse) {
}

func PostFlattenDtcMonitorSnmpNIOS(ctx context.Context, planned, flattened *NIOSDtcMonitorSnmpModel, diags *diag.Diagnostics) {
	if planned == nil || flattened == nil {
		return
	}

	if !planned.Oids.IsUnknown() {
		if reordered, d := utils.ReorderAndFilterNestedListResponse(ctx, planned.Oids, flattened.Oids, "oid"); !d.HasError() {
			if reorderedList, ok := reordered.(basetypes.ListValue); ok {
				flattened.Oids = reorderedList
			}
		}
	}
}
