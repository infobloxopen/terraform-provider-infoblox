package rpz

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateRecordRpzNaptr validates the RecordRpzNaptr configuration.
func ValidateRecordRpzNaptr(ctx context.Context, data RecordRpzNaptrModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordRpzNaptrModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordRpzNaptrNIOSConfig(ctx, nios, resp)
	}
}

func validateRecordRpzNaptrNIOSConfig(ctx context.Context, m *NIOSRecordRpzNaptrModel, resp *resource.ValidateConfigResponse) {
}

func PostFlattenRecordRpzNaptrNIOS(ctx context.Context, planned, flattened *NIOSRecordRpzNaptrModel, diags *diag.Diagnostics) {
	if planned != nil && planned.Ttl.IsNull() {
		flattened.Ttl = types.Int64Null()
	}
}
