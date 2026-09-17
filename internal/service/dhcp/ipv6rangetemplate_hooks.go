package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateIpv6rangetemplate validates the Ipv6rangetemplate configuration.
func ValidateIpv6rangetemplate(ctx context.Context, data Ipv6rangetemplateModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSIpv6rangetemplateModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateIpv6rangetemplateNIOSConfig(ctx, nios, resp)
	}
}

func validateIpv6rangetemplateNIOSConfig(ctx context.Context, m *NIOSIpv6rangetemplateModel, resp *resource.ValidateConfigResponse) {
	niosPath := path.Root("nios")

	var serverAssociationType string

	if !m.ServerAssociationType.IsUnknown() {
		serverAssociationType = "NONE"
		if !m.ServerAssociationType.IsNull() {
			serverAssociationType = m.ServerAssociationType.ValueString()
		}
	}

	// If server_association_type is MEMBER, member field must be set
	if serverAssociationType == "MEMBER" {
		if m.Member.IsNull() {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("member"),
				"Invalid Configuration",
				"The 'member' field must be set when 'server_association_type' is set to 'MEMBER'.",
			)
		}
	}

	// If server_association_type is NONE, member field cannot be set
	if serverAssociationType == "NONE" {
		if !m.Member.IsNull() && !m.Member.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("member"),
				"Invalid Configuration",
				"The 'member' field cannot be set when 'server_association_type' is set to 'NONE' (default).",
			)
		}
	}
}
