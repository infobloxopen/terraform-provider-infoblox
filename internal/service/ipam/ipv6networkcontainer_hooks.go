package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/dynamicallocation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateIpv6networkcontainer validates the Ipv6networkcontainer configuration.
func ValidateIpv6networkcontainer(ctx context.Context, data Ipv6networkcontainerModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSIpv6networkcontainerModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateIpv6networkcontainerNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIIpv6networkcontainerModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateIpv6networkcontainerUDDIConfig(ctx, uddi, resp)
	}
}

func validateIpv6networkcontainerNIOSConfig(ctx context.Context, m *NIOSIpv6networkcontainerModel, resp *resource.ValidateConfigResponse) {
	// DHCP option rules are identical across every NIOS object carrying options.
	utils.ValidateDHCPOptionsConfig(ctx, m.Options, path.Root("nios").AtName("options"), &resp.Diagnostics)

	// Validate discovery_blackout_setting blackout_schedule
	if !m.DiscoveryBlackoutSetting.IsNull() && !m.DiscoveryBlackoutSetting.IsUnknown() {
		utils.ValidateScheduleConfig(
			m.DiscoveryBlackoutSetting,
			"blackout_schedule",
			path.Root("nios").AtName("discovery_blackout_setting"),
			&resp.Diagnostics,
		)
	}

	// Validate port_control_blackout_setting blackout_schedule
	if !m.PortControlBlackoutSetting.IsNull() && !m.PortControlBlackoutSetting.IsUnknown() {
		utils.ValidateScheduleConfig(
			m.PortControlBlackoutSetting,
			"blackout_schedule",
			path.Root("nios").AtName("port_control_blackout_setting"),
			&resp.Diagnostics,
		)
	}

	// Validate subscribe_settings: enabled_attributes is required, and each
	// mapped_ea_attributes item requires name and mapped_ea.
	if !m.SubscribeSettings.IsNull() && !m.SubscribeSettings.IsUnknown() {
		var subscribeSettings Ipv6networkcontainerSubscribeSettingsModel
		resp.Diagnostics.Append(m.SubscribeSettings.As(ctx, &subscribeSettings, basetypes.ObjectAsOptions{})...)
		if !resp.Diagnostics.HasError() {
			// enabled_attributes is required when subscribe_settings is configured
			if subscribeSettings.EnabledAttributes.IsNull() {
				resp.Diagnostics.AddAttributeError(
					path.Root("nios").AtName("subscribe_settings").AtName("enabled_attributes"),
					"Missing Required Attribute",
					"The 'enabled_attributes' attribute is required when 'subscribe_settings' is configured.",
				)
			}

			if !subscribeSettings.MappedEaAttributes.IsNull() && !subscribeSettings.MappedEaAttributes.IsUnknown() {
				var mappedEaAttrs []Ipv6networkcontainersubscribesettingsMappedEaAttributesModel
				resp.Diagnostics.Append(subscribeSettings.MappedEaAttributes.ElementsAs(ctx, &mappedEaAttrs, false)...)
				for i, item := range mappedEaAttrs {
					if !item.Name.IsUnknown() && (item.Name.IsNull() || item.Name.ValueString() == "") {
						resp.Diagnostics.AddAttributeError(
							path.Root("nios").AtName("subscribe_settings").AtName("mapped_ea_attributes").AtListIndex(i).AtName("name"),
							"Missing Required Attribute",
							"The 'name' attribute is required for each item in 'mapped_ea_attributes'.",
						)
					}
					if !item.MappedEa.IsUnknown() && (item.MappedEa.IsNull() || item.MappedEa.ValueString() == "") {
						resp.Diagnostics.AddAttributeError(
							path.Root("nios").AtName("subscribe_settings").AtName("mapped_ea_attributes").AtListIndex(i).AtName("mapped_ea"),
							"Missing Required Attribute",
							"The 'mapped_ea' attribute is required for each item in 'mapped_ea_attributes'.",
						)
					}
				}
			}
		}
	}
}

func validateIpv6networkcontainerUDDIConfig(ctx context.Context, m *UDDIIpv6networkcontainerModel, resp *resource.ValidateConfigResponse) {
}

func PostFlattenIpv6networkcontainerNIOS(ctx context.Context, planned, flattened *NIOSIpv6networkcontainerModel, diags *diag.Diagnostics) {
	if planned == nil || flattened == nil {
		return
	}
	if flattened.RirRegistrationAction.IsNull() || flattened.RirRegistrationAction.IsUnknown() {
		flattened.RirRegistrationAction = planned.RirRegistrationAction
	}
	if !planned.Options.IsNull() && !planned.Options.IsUnknown() {
		if reordered, d := utils.ReorderAndFilterDHCPOptions(ctx, planned.Options, flattened.Options); !d.HasError() {
			flattened.Options = reordered.(basetypes.ListValue)
		}
	}
}

func BuildIpv6networkcontainerFuncCall(ctx context.Context, data types.Object, diags *diag.Diagnostics) *niosipam.FuncCall {
	if data.IsNull() || data.IsUnknown() {
		return nil
	}

	var m dynamicallocation.NextAvailableNetworkModel
	diags.Append(data.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.FuncCall(ctx, "Network", "ipv6networkcontainer", diags)

}
