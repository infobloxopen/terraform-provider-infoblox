package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	internaltypes "github.com/infobloxopen/terraform-provider-infoblox/internal/types"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateIpv6sharednetwork validates the Ipv6sharednetwork configuration.
func ValidateIpv6sharednetwork(ctx context.Context, data Ipv6sharednetworkModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSIpv6sharednetworkModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateIpv6sharednetworkNIOSConfig(ctx, nios, resp)
	}
}

func validateIpv6sharednetworkNIOSConfig(ctx context.Context, m *NIOSIpv6sharednetworkModel, resp *resource.ValidateConfigResponse) {
	utils.ValidateDHCPOptionsConfig(ctx, m.Options, path.Root("nios").AtName("options"), &resp.Diagnostics)
}

func PostFlattenIpv6sharednetworkNIOS(ctx context.Context, planned, flattened *NIOSIpv6sharednetworkModel, diags *diag.Diagnostics) {
	if planned == nil || flattened == nil {
		return
	}

	if !planned.Options.IsUnknown() {
		if reordered, d := utils.ReorderAndFilterDHCPOptions(ctx, planned.Options, flattened.Options); !d.HasError() {
			if reorderedList, ok := reordered.(basetypes.ListValue); ok {
				flattened.Options = reorderedList
			}
		}
	}
}

func ExpandNetworks(ctx context.Context, networks internaltypes.UnorderedListValue, diags *diag.Diagnostics) []niosdhcp.Ipv6sharednetworkNetworks {
	if networks.IsNull() || networks.IsUnknown() {
		return nil
	}

	var networkRefs []string
	diags.Append(networks.ElementsAs(ctx, &networkRefs, false)...)
	if diags.HasError() {
		return nil
	}

	result := make([]niosdhcp.Ipv6sharednetworkNetworks, len(networkRefs))
	for i, ref := range networkRefs {
		result[i] = niosdhcp.Ipv6sharednetworkNetworks{
			Ref: &ref,
		}
	}
	return result
}

func FlattenNetworks(ctx context.Context, networks []niosdhcp.Ipv6sharednetworkNetworks, diags *diag.Diagnostics) internaltypes.UnorderedListValue {
	if networks == nil {
		return internaltypes.NewUnorderedListValueNull(types.StringType)
	}

	networkRefs := make([]string, len(networks))
	for i, network := range networks {
		if network.Ref != nil {
			networkRefs[i] = *network.Ref
		}
	}

	listValue, d := internaltypes.NewUnorderedListValueFrom(ctx, types.StringType, networkRefs)
	diags.Append(d...)
	return listValue
}
