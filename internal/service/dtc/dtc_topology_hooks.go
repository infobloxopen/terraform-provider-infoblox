package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateDtcTopology validates the DtcTopology configuration.
func ValidateDtcTopology(ctx context.Context, data DtcTopologyModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSDtcTopologyModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateDtcTopologyNIOSConfig(ctx, nios, resp)
	}
}

func validateDtcTopologyNIOSConfig(ctx context.Context, m *NIOSDtcTopologyModel, resp *resource.ValidateConfigResponse) {
}
