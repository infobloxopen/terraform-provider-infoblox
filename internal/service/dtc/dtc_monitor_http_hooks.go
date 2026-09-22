package dtc

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateDtcMonitorHttp validates the DtcMonitorHttp configuration.
func ValidateDtcMonitorHttp(ctx context.Context, data DtcMonitorHttpModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSDtcMonitorHttpModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateDtcMonitorHttpNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIDtcMonitorHttpModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateDtcMonitorHttpUDDIConfig(ctx, uddi, resp)
	}
}

func validateDtcMonitorHttpNIOSConfig(ctx context.Context, m *NIOSDtcMonitorHttpModel, resp *resource.ValidateConfigResponse) {
	if m.ContentCheck.IsNull() || m.ContentCheck.IsUnknown() {
		return
	}

	contentCheckValue := m.ContentCheck.ValueString()
	niosPath := path.Root("nios")

	if contentCheckValue == "EXTRACT" {
		if m.ContentCheckRegex.IsNull() || m.ContentExtractType.IsNull() || m.ContentExtractValue.IsNull() || m.ContentCheckOp.IsNull() {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("content_check"),
				"Invalid configuration for content check EXTRACT",
				"When 'content_check' is set to 'EXTRACT', the fields 'content_check_regex', 'content_extract_type', 'content_check_op' and 'content_extract_value' must be provided.",
			)
		}
	}

	if contentCheckValue == "MATCH" {
		if m.ContentCheckRegex.IsNull() || m.ContentCheckOp.IsNull() {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("content_check"),
				"Invalid configuration for content check MATCH",
				"When 'content_check' is set to 'MATCH', 'content_check_regex' and 'content_check_op' must be provided.",
			)
		}
	}
}

func validateDtcMonitorHttpUDDIConfig(ctx context.Context, m *UDDIDtcMonitorHttpModel, resp *resource.ValidateConfigResponse) {
}

// PostFlattenDtcMonitorHttpUDDI strips the trailing "\r\n" that the UDDI API
// automatically appends to request values on read-back.
func PostFlattenDtcMonitorHttpUDDI(ctx context.Context, planned, flattened *UDDIDtcMonitorHttpModel, diags *diag.Diagnostics) {
	if planned == nil || flattened == nil {
		return
	}
	if flattened.Request.IsNull() || flattened.Request.IsUnknown() {
		return
	}

	requestValue := flattened.Request.ValueString()
	planValue := planned.Request.ValueString()

	if strings.HasSuffix(requestValue, "\r\n\r\n") && !strings.HasSuffix(planValue, "\r\n\r\n") {
		requestValue = strings.TrimSuffix(requestValue, "\r\n")
	}

	flattened.Request = types.StringValue(requestValue)
}

// PostFlattenDtcMonitorHttpNIOS strips the trailing "\nConnection: close\n\n" that
// NIOS automatically appends to any request value on read-back, so the stored state
// matches what the user configured and no spurious diff is produced.
func PostFlattenDtcMonitorHttpNIOS(ctx context.Context, planned, flattened *NIOSDtcMonitorHttpModel, diags *diag.Diagnostics) {
	if planned == nil || flattened == nil {
		return
	}
	if flattened.Request.IsNull() || flattened.Request.IsUnknown() {
		return
	}

	requestValue := flattened.Request.ValueString()
	planValue := planned.Request.ValueString()

	if strings.HasPrefix(requestValue, "POST") || strings.HasPrefix(requestValue, "GET") || strings.HasPrefix(requestValue, "HEAD") {
		if strings.HasSuffix(requestValue, "\n\n") && !strings.HasSuffix(planValue, "\n\n") {
			requestValue = strings.TrimSuffix(requestValue, "\n\n")
		}
		if strings.Contains(requestValue, "HTTP/1.1") && !strings.Contains(planValue, "\nConnection: close") {
			requestValue = strings.ReplaceAll(requestValue, "\nConnection: close", "")
		}
	}

	flattened.Request = types.StringValue(requestValue)
}
