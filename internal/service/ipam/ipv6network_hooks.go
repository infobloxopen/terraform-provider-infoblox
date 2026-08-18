package ipam

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/dynamicallocation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateIpv6network validates the Ipv6network configuration.
func ValidateIpv6network(ctx context.Context, data Ipv6networkModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSIpv6networkModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateIpv6networkNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIIpv6networkModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateIpv6networkUDDIConfig(ctx, uddi, resp)
	}
}

func validateIpv6networkNIOSConfig(ctx context.Context, m *NIOSIpv6networkModel, resp *resource.ValidateConfigResponse) {
}

func validateIpv6networkUDDIConfig(ctx context.Context, m *UDDIIpv6networkModel, resp *resource.ValidateConfigResponse) {
}

func BuildIpv6networkFuncCall(ctx context.Context, data types.Object, diags *diag.Diagnostics) *niosipam.FuncCall {
	if data.IsNull() || data.IsUnknown() {
		return nil
	}

	var m dynamicallocation.NextAvailableNetworkModel
	diags.Append(data.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.FuncCall(ctx, "Network", "ipv6network", diags)
}

func PostFlattenIpv6networkNIOS(ctx context.Context, planned, flattened *NIOSIpv6networkModel, diags *diag.Diagnostics) {
	if planned != nil && !planned.Options.IsUnknown() {
		reordered, d := utils.ReorderAndFilterDHCPOptions(ctx, planned.Options, flattened.Options)
		diags.Append(*d...)
		if d.HasError() {
			return
		}
		if reorderedList, ok := reordered.(basetypes.ListValue); ok {
			flattened.Options = reorderedList
		}
	}
}

func (r *Ipv6networkResource) isIpv6networkContainerConversionError(err error) bool {
	errVal := err.Error()
	return (strings.Contains(errVal, "The search parameters") &&
		strings.Contains(errVal, "for object ipv6network did not return any result")) ||
		strings.Contains(errVal, "will overlap an existing network")
}

func (r *Ipv6networkResource) isIpv6networkConvertedToContainer(ctx context.Context, data *Ipv6networkModel) bool {
	if r.backend != core.BackendNIOS || r.containerService == nil {
		return false
	}

	var diags diag.Diagnostics
	nios := flex.ExpandNestedObject[NIOSIpv6networkModel](ctx, data.NIOS, &diags)
	if nios == nil || diags.HasError() {
		return false
	}

	// Try to fetch as Network container
	records, _, _, err := r.containerService.List(ctx, &core.ListOptions{
		Filters: map[string]string{
			"nios.network": nios.Network.ValueString(),
		},
	})
	return err == nil && len(records) > 0
}
