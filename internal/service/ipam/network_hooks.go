package ipam

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/dynamicallocation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateNetwork validates the Network configuration.
func ValidateNetwork(ctx context.Context, data NetworkModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSNetworkModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateNetworkNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDINetworkModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateNetworkUDDIConfig(ctx, uddi, resp)
	}
}

func validateNetworkNIOSConfig(ctx context.Context, m *NIOSNetworkModel, resp *resource.ValidateConfigResponse) {
	niosPath := path.Root("nios")

	// DHCP options validation
	utils.ValidateDHCPOptionsConfig(ctx, m.Options, niosPath.AtName("options"), &resp.Diagnostics)

	// Members validation
	if !m.Members.IsNull() && !m.Members.IsUnknown() {
		var members []NetworkMembersModel
		resp.Diagnostics.Append(m.Members.ElementsAs(ctx, &members, false)...)
		if resp.Diagnostics.HasError() {
			return
		}

		membersPath := niosPath.AtName("members")
		for i, member := range members {
			if member.Struct.ValueString() == "msdhcpserver" {
				if !member.Ipv6addr.IsNull() && !member.Ipv6addr.IsUnknown() {
					resp.Diagnostics.AddAttributeError(
						membersPath.AtListIndex(i).AtName("ipv6addr"),
						"Invalid Configuration",
						"ipv6addr cannot be set when struct is 'msdhcpserver'. Only ipv4addr is supported for msdhcpserver.",
					)
				}

				if !member.Name.IsNull() && !member.Name.IsUnknown() {
					resp.Diagnostics.AddAttributeError(
						membersPath.AtListIndex(i).AtName("name"),
						"Invalid Configuration",
						"name cannot be set when struct is 'msdhcpserver'. Only ipv4addr is supported for msdhcpserver.",
					)
				}
			}
		}
	}

	// Validate discovery_blackout_setting blackout_schedule
	if !m.DiscoveryBlackoutSetting.IsNull() && !m.DiscoveryBlackoutSetting.IsUnknown() {
		utils.ValidateScheduleConfig(
			m.DiscoveryBlackoutSetting,
			"blackout_schedule",
			niosPath.AtName("discovery_blackout_setting"),
			&resp.Diagnostics,
		)
	}

	// Validate port_control_blackout_setting blackout_schedule
	if !m.PortControlBlackoutSetting.IsNull() && !m.PortControlBlackoutSetting.IsUnknown() {
		utils.ValidateScheduleConfig(
			m.PortControlBlackoutSetting,
			"blackout_schedule",
			niosPath.AtName("port_control_blackout_setting"),
			&resp.Diagnostics,
		)
	}

	// enabled_attributes is required when subscribe_settings is configured
	if !m.SubscribeSettings.IsNull() && !m.SubscribeSettings.IsUnknown() {
		attrs := m.SubscribeSettings.Attributes()
		enabledAttrs, exists := attrs["enabled_attributes"]
		if !exists || enabledAttrs.IsNull() {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("subscribe_settings").AtName("enabled_attributes"),
				"Missing Required Attribute",
				"The 'enabled_attributes' attribute is required when 'subscribe_settings' is configured.",
			)
		}
	}
}

func validateNetworkUDDIConfig(ctx context.Context, m *UDDINetworkModel, resp *resource.ValidateConfigResponse) {
}

func BuildNetworkFuncCall(ctx context.Context, data types.Object, diags *diag.Diagnostics) *niosipam.FuncCall {
	if data.IsNull() || data.IsUnknown() {
		return nil
	}

	var m dynamicallocation.NextAvailableNetworkModel
	diags.Append(data.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.FuncCall(ctx, "Network", "network", diags)
}

func PostFlattenNetworkNIOS(ctx context.Context, planned, flattened *NIOSNetworkModel, diags *diag.Diagnostics) {
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

func BuildNetworkAllocation(ctx context.Context, allocObj types.Object, diags *diag.Diagnostics) *string {
	if allocObj.IsNull() || allocObj.IsUnknown() {
		return nil
	}

	var m dynamicallocation.NextAvailableSubnetModel
	diags.Append(allocObj.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	if m.NextAvailableId.IsNull() || m.NextAvailableId.IsUnknown() {
		return nil
	}

	allocated := m.Suffixed("/nextavailablesubnet")
	return &allocated
}

// LockNetworkAllocation serializes concurrent next-available allocations
// that target the same parent scope by acquiring a per-scope mutex keyed on the next_available_id
func LockNetworkAllocation(ctx context.Context, uddiBlock types.Object, diags *diag.Diagnostics) func() {
	noop := func() {}
	if uddiBlock.IsNull() || uddiBlock.IsUnknown() {
		return noop
	}

	allocVal, ok := uddiBlock.Attributes()["dynamic_allocation"]
	if !ok {
		return noop
	}
	allocObj, ok := allocVal.(types.Object)
	if !ok || allocObj.IsNull() || allocObj.IsUnknown() {
		return noop
	}

	var m dynamicallocation.NextAvailableSubnetModel
	diags.Append(allocObj.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return noop
	}

	key := m.NextAvailableId.ValueString()
	if key == "" {
		return noop
	}

	utils.GlobalMutexStore.Lock(key)
	return func() { utils.GlobalMutexStore.Unlock(key) }
}

func (r *NetworkResource) isNetworkContainerConversionError(err error) bool {
	errVal := err.Error()
	return (strings.Contains(errVal, "The search parameters") &&
		strings.Contains(errVal, "for object network did not return any result")) ||
		strings.Contains(errVal, "will overlap an existing network")
}

func (r *NetworkResource) isNetworkConvertedToContainer(ctx context.Context, data *NetworkModel) bool {
	if r.backend != core.BackendNIOS || r.containerService == nil {
		return false
	}

	var diags diag.Diagnostics
	nios := flex.ExpandNestedObject[NIOSNetworkModel](ctx, data.NIOS, &diags)
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
