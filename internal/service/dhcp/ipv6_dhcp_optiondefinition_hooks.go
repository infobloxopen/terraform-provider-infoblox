package dhcp

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateIpv6DhcpOptiondefinition validates the Ipv6DhcpOptiondefinition configuration.
func ValidateIpv6DhcpOptiondefinition(ctx context.Context, data Ipv6DhcpOptiondefinitionModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSIpv6DhcpOptiondefinitionModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateIpv6DhcpOptiondefinitionNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIIpv6DhcpOptiondefinitionModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateIpv6DhcpOptiondefinitionUDDIConfig(ctx, uddi, resp)
	}
}

const (
	niosDefaultIpv6OptionSpace = "DHCPv6"
	niosIpv6OptionNamePrefix   = "dhcp6."
)

func validateIpv6DhcpOptiondefinitionNIOSConfig(ctx context.Context, m *NIOSIpv6DhcpOptiondefinitionModel, resp *resource.ValidateConfigResponse) {
	if m.Space.IsUnknown() || m.Name.IsUnknown() || m.Name.IsNull() {
		return
	}

	space := niosDefaultIpv6OptionSpace
	if !m.Space.IsNull() {
		space = m.Space.ValueString()
	}
	hasPrefix := strings.HasPrefix(m.Name.ValueString(), niosIpv6OptionNamePrefix)

	switch {
	case space == niosDefaultIpv6OptionSpace && !hasPrefix:
		resp.Diagnostics.AddError(
			"Invalid Name for DHCPv6 Option Definition",
			"The name of a DHCP IPv6 option definition object in the default space (DHCPv6) must start with 'dhcp6.'.",
		)
	case space != niosDefaultIpv6OptionSpace && hasPrefix:
		resp.Diagnostics.AddError(
			"Invalid Name for Custom DHCPv6 Option Definition",
			"The name of a DHCP IPv6 option definition object in a custom space must not start with 'dhcp6.'.",
		)
	}
}

func validateIpv6DhcpOptiondefinitionUDDIConfig(ctx context.Context, m *UDDIIpv6DhcpOptiondefinitionModel, resp *resource.ValidateConfigResponse) {
}

func (r *Ipv6DhcpOptiondefinitionResource) refreshIpv6DhcpOptiondefinitionId(ctx context.Context, resp *resource.UpdateResponse, data, stateData *Ipv6DhcpOptiondefinitionModel) {
	if r.backend != core.BackendNIOS {
		return
	}

	planNIOS := flex.ExpandNestedObject[NIOSIpv6DhcpOptiondefinitionModel](ctx, data.NIOS, &resp.Diagnostics)
	stateNIOS := flex.ExpandNestedObject[NIOSIpv6DhcpOptiondefinitionModel](ctx, stateData.NIOS, &resp.Diagnostics)
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
		ReturnFields: Ipv6DhcpOptiondefinitionReturnFields,
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to list Ipv6DhcpOptiondefinition to refresh its reference after the option space changed: %s", err),
		)
		return
	}

	if len(results) == 0 || results[0].Id == nil {
		return
	}

	data.Id = types.StringValue(*results[0].Id)
}
