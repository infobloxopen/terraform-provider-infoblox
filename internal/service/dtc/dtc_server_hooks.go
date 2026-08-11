package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateDtcServer validates the DtcServer configuration.
func ValidateDtcServer(ctx context.Context, data DtcServerModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSDtcServerModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateDtcServerNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIDtcServerModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateDtcServerUDDIConfig(ctx, uddi, resp)
	}
}

func validateDtcServerNIOSConfig(ctx context.Context, m *NIOSDtcServerModel, resp *resource.ValidateConfigResponse) {
}

func validateDtcServerUDDIConfig(ctx context.Context, m *UDDIDtcServerModel, resp *resource.ValidateConfigResponse) {
	endpointTypePath := path.Root("uddi").AtName("endpoint_type")

	if !m.Address.IsNull() && !m.Address.IsUnknown() {
		if m.EndpointType.IsNull() || m.EndpointType.IsUnknown() || m.EndpointType.ValueString() != "address" {
			resp.Diagnostics.AddAttributeError(
				endpointTypePath,
				"Conflicting endpoint_type",
				`endpoint_type must be set to "address" when address is provided.`,
			)
		}
	}

	if !m.Fqdn.IsNull() && !m.Fqdn.IsUnknown() {
		if m.EndpointType.IsNull() || m.EndpointType.IsUnknown() || m.EndpointType.ValueString() != "fqdn" {
			resp.Diagnostics.AddAttributeError(
				endpointTypePath,
				"Conflicting endpoint_type",
				`endpoint_type must be set to "fqdn" when fqdn is provided.`,
			)
		}
	}
}
