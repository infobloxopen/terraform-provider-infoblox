package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/dynamicallocation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateFixedaddress validates the Fixedaddress configuration.
func ValidateFixedaddress(ctx context.Context, data FixedaddressModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSFixedaddressModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateFixedaddressNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIFixedaddressModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateFixedaddressUDDIConfig(ctx, uddi, resp)
	}
}

func validateFixedaddressNIOSConfig(ctx context.Context, m *NIOSFixedaddressModel, resp *resource.ValidateConfigResponse) {
	niosPath := path.Root("nios")

	if m.MatchClient.ValueString() == "MAC_ADDRESS" {
		if m.Mac.IsNull() {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("mac"),
				"Invalid configuration",
				"The 'mac' attribute must be set when 'match_client' is set to 'MAC_ADDRESS'.",
			)
		}
		if (!m.AgentCircuitId.IsNull() && !m.AgentCircuitId.IsUnknown()) ||
			(!m.AgentRemoteId.IsNull() && !m.AgentRemoteId.IsUnknown()) ||
			(!m.DhcpClientIdentifier.IsNull() && !m.DhcpClientIdentifier.IsUnknown()) {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("match_client"),
				"Invalid configuration",
				"When 'match_client' is set to 'MAC_ADDRESS', the 'agent_circuit_id', 'agent_remote_id', and 'dhcp_client_identifier' attributes must not be set.",
			)
		}

	} else if m.MatchClient.ValueString() == "CLIENT_ID" {
		if !m.DhcpClientIdentifier.IsUnknown() && (m.DhcpClientIdentifier.IsNull() || m.DhcpClientIdentifier.ValueString() == "") {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("dhcp_client_identifier"),
				"Invalid configuration",
				"The 'dhcp_client_identifier' attribute must be set and cannot be empty when 'match_client' is set to 'CLIENT_ID'.",
			)
		}
		if (!m.AgentCircuitId.IsNull() && !m.AgentCircuitId.IsUnknown()) ||
			(!m.AgentRemoteId.IsNull() && !m.AgentRemoteId.IsUnknown()) ||
			(!m.Mac.IsNull() && !m.Mac.IsUnknown()) {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("match_client"),
				"Invalid configuration",
				"When 'match_client' is set to 'CLIENT_ID', the 'agent_circuit_id', 'agent_remote_id', and 'mac' attributes must not be set.",
			)
		}
	} else if m.MatchClient.ValueString() == "CIRCUIT_ID" {
		if !m.AgentCircuitId.IsUnknown() && (m.AgentCircuitId.IsNull() || m.AgentCircuitId.ValueString() == "") {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("agent_circuit_id"),
				"Invalid configuration",
				"The 'agent_circuit_id' attribute must be set when 'match_client' is set to 'CIRCUIT_ID'.",
			)
		}
		if (!m.Mac.IsNull() && !m.Mac.IsUnknown()) ||
			(!m.DhcpClientIdentifier.IsNull() && !m.DhcpClientIdentifier.IsUnknown()) {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("match_client"),
				"Invalid configuration",
				"When 'match_client' is set to 'CIRCUIT_ID', the 'mac' and 'dhcp_client_identifier' attributes must not be set.",
			)
		}
	} else if m.MatchClient.ValueString() == "REMOTE_ID" {
		if !m.AgentRemoteId.IsUnknown() && (m.AgentRemoteId.IsNull() || m.AgentRemoteId.ValueString() == "") {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("agent_remote_id"),
				"Invalid configuration",
				"The 'agent_remote_id' attribute must be set when 'match_client' is set to 'REMOTE_ID'.",
			)
		}
		if (!m.Mac.IsNull() && !m.Mac.IsUnknown()) ||
			(!m.DhcpClientIdentifier.IsNull() && !m.DhcpClientIdentifier.IsUnknown()) {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("match_client"),
				"Invalid configuration",
				"When 'match_client' is set to 'REMOTE_ID', the 'mac' and 'dhcp_client_identifier' attributes must not be set.",
			)
		}
	} else if m.MatchClient.ValueString() == "RESERVED" {
		if !m.Mac.IsNull() && !m.Mac.IsUnknown() && m.Mac.ValueString() != "00:00:00:00:00:00" {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("mac"),
				"Invalid configuration",
				"When 'match_client' is set to 'RESERVED', the 'mac' attribute must be set to '00:00:00:00:00:00' or left unset.",
			)
		}
		if (!m.AgentCircuitId.IsNull() && !m.AgentCircuitId.IsUnknown()) ||
			(!m.AgentRemoteId.IsNull() && !m.AgentRemoteId.IsUnknown()) ||
			(!m.DhcpClientIdentifier.IsNull() && !m.DhcpClientIdentifier.IsUnknown()) {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("match_client"),
				"Invalid configuration",
				"When 'match_client' is set to 'RESERVED', the 'agent_circuit_id', 'agent_remote_id', and 'dhcp_client_identifier' attributes must not be set.",
			)
		}
	}

	// DHCP options validation
	utils.ValidateDHCPOptionsConfig(ctx, m.Options, niosPath.AtName("options"), &resp.Diagnostics)

	// Check if allow_telnet is true, then cli_credentials must contain at least one element with credential_type set to "TELNET"
	if !m.AllowTelnet.IsUnknown() && !m.AllowTelnet.IsNull() && m.AllowTelnet.ValueBool() {
		isTelnet := false
		isSSH := false
		if !m.CliCredentials.IsNull() && !m.CliCredentials.IsUnknown() {
			// Iterate through cli_credentials to check if an element has credential_type set to "TELNET"
			var cliCredentials []FixedaddressCliCredentialsModel
			diags := m.CliCredentials.ElementsAs(ctx, &cliCredentials, false)
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}

			for _, credentials := range cliCredentials {
				if credentials.CredentialType.IsUnknown() || credentials.CredentialType.IsNull() {
					continue
				}
				credentialsType := credentials.CredentialType.ValueString()
				if credentialsType == "SSH" {
					isSSH = true
				}
				if credentialsType == "TELNET" {
					isTelnet = true
				}
			}
		}
		if !isSSH {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("allow_telnet"),
				"Invalid configuration",
				"The 'cli_credentials' must contain credentials with 'credential_type' set to 'SSH'.",
			)
		}
		if !isTelnet {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("allow_telnet"),
				"Invalid configuration",
				"The 'allow_telnet' attribute must be set to false when 'cli_credentials' is not set or does not contain any credentials with 'credential_type' set to 'TELNET'.",
			)
		}
	}

	// Check if SNMP , then the corresponding use_snmp_credential attribute must be set to true
	if !m.SnmpCredential.IsUnknown() && !m.SnmpCredential.IsNull() {
		// Validate that community_string is provided when snmp_credential block is set
		var snmpCredential FixedaddressSnmpCredentialModel
		diags := m.SnmpCredential.As(ctx, &snmpCredential, basetypes.ObjectAsOptions{})
		resp.Diagnostics.Append(diags...)
		if !resp.Diagnostics.HasError() {
			if snmpCredential.CommunityString.IsNull() || (!snmpCredential.CommunityString.IsUnknown() && snmpCredential.CommunityString.ValueString() == "") {
				resp.Diagnostics.AddAttributeError(
					path.Root("snmp_credential").AtName("community_string"),
					"Invalid configuration",
					"The 'community_string' attribute must be set when 'snmp_credential' is configured.",
				)
			}
		}
	}
}

func validateFixedaddressUDDIConfig(ctx context.Context, m *UDDIFixedaddressModel, resp *resource.ValidateConfigResponse) {
}

func BuildFixedaddressAllocation(ctx context.Context, allocObj types.Object, diags *diag.Diagnostics) *string {
	if allocObj.IsNull() || allocObj.IsUnknown() {
		return nil
	}

	var m dynamicallocation.NextAvailableAddressModel
	diags.Append(allocObj.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	if m.NextAvailableId.IsNull() || m.NextAvailableId.IsUnknown() {
		return nil
	}

	allocated := m.Suffixed("/nextavailableip")
	return &allocated
}

// LockFixedaddressAllocation serializes concurrent next-available allocations
// that target the same parent scope by acquiring a per-scope mutex keyed on the next_available_id
func LockFixedaddressAllocation(ctx context.Context, uddiBlock types.Object, diags *diag.Diagnostics) func() {
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

	var m dynamicallocation.NextAvailableAddressModel
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

func BuildFixedaddressFuncCall(ctx context.Context, data types.Object, diags *diag.Diagnostics) *niosdhcp.FuncCall {
	if data.IsNull() || data.IsUnknown() {
		return nil
	}

	var m dynamicallocation.NextAvailableIpModel
	diags.Append(data.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.FuncCallDHCP(ctx, "Ipv4addr", "network", diags)
}

func PostFlattenFixedaddressNIOS(ctx context.Context, planned, flattened *NIOSFixedaddressModel, diags *diag.Diagnostics) {
	if planned == nil || flattened == nil {
		return
	}

	if !planned.CliCredentials.IsUnknown() {
		if reordered, d := utils.ReorderAndFilterNestedListResponse(ctx, planned.CliCredentials, flattened.CliCredentials, "credential_type"); !d.HasError() {
			if reorderedList, ok := reordered.(basetypes.ListValue); ok {
				flattened.CliCredentials = reorderedList
			}
		}
	}

	if !planned.Options.IsUnknown() {
		if reordered, d := utils.ReorderAndFilterDHCPOptions(ctx, planned.Options, flattened.Options); !d.HasError() {
			if reorderedList, ok := reordered.(basetypes.ListValue); ok {
				flattened.Options = reorderedList
			}
		}
	}

	if flattened.Template.IsNull() && !planned.Template.IsNull() && !planned.Template.IsUnknown() {
		flattened.Template = planned.Template
	}

	if result, d := utils.CopyFieldFromPlanToRespList(ctx, planned.CliCredentials, flattened.CliCredentials, "password"); !d.HasError() {
		if resultList, ok := result.(basetypes.ListValue); ok {
			flattened.CliCredentials = resultList
		}
	}

	for _, field := range []string{"authentication_password", "privacy_password"} {
		if result, d := utils.CopyFieldFromPlanToRespObject(ctx, planned.Snmp3Credential, flattened.Snmp3Credential, field); !d.HasError() {
			if resultObj, ok := result.(basetypes.ObjectValue); ok {
				flattened.Snmp3Credential = resultObj
			}
		}
	}
}
