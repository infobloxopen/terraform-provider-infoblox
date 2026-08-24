package dhcp

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateDhcpOptiondefinition validates the DhcpOptiondefinition configuration.
func ValidateDhcpOptiondefinition(ctx context.Context, data DhcpOptiondefinitionModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSDhcpOptiondefinitionModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateDhcpOptiondefinitionNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIDhcpOptiondefinitionModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateDhcpOptiondefinitionUDDIConfig(ctx, uddi, resp)
	}
}

func validateDhcpOptiondefinitionNIOSConfig(ctx context.Context, m *NIOSDhcpOptiondefinitionModel, resp *resource.ValidateConfigResponse) {
}

func validateDhcpOptiondefinitionUDDIConfig(ctx context.Context, m *UDDIDhcpOptiondefinitionModel, resp *resource.ValidateConfigResponse) {
}

func (r *DhcpOptiondefinitionResource) refreshDhcpOptiondefinitionId(ctx context.Context, resp *resource.UpdateResponse, data, stateData *DhcpOptiondefinitionModel) {
	if r.backend != core.BackendNIOS {
		return
	}

	planNIOS := flex.ExpandNestedObject[NIOSDhcpOptiondefinitionModel](ctx, data.NIOS, &resp.Diagnostics)
	stateNIOS := flex.ExpandNestedObject[NIOSDhcpOptiondefinitionModel](ctx, stateData.NIOS, &resp.Diagnostics)
	if resp.Diagnostics.HasError() || planNIOS == nil || stateNIOS == nil {
		return
	}

	if planNIOS.Space.IsUnknown() || planNIOS.Space.Equal(stateNIOS.Space) {
		return
	}

	results, _, _, err := r.service.List(ctx, &core.ListOptions{
		Filters: map[string]string{
			"nios.name":  stateNIOS.Name.ValueString(),
			"nios.space": planNIOS.Space.ValueString(),
			"nios.code":  strconv.FormatInt(stateNIOS.Code.ValueInt64(), 10),
			"nios.type":  stateNIOS.Type.ValueString(),
		},
		ReturnFields: DhcpOptiondefinitionReturnFields,
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to list DhcpOptiondefinition to refresh its reference after the option space changed: %s", err),
		)
		return
	}

	if len(results) == 0 || results[0].Id == nil {
		return
	}

	data.Id = types.StringValue(*results[0].Id)
}
