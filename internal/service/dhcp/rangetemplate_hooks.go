package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateRangetemplate validates the Rangetemplate configuration.
func ValidateRangetemplate(ctx context.Context, data RangetemplateModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRangetemplateModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRangetemplateNIOSConfig(ctx, nios, resp)
	}
}

func validateRangetemplateNIOSConfig(ctx context.Context, m *NIOSRangetemplateModel, resp *resource.ValidateConfigResponse) {
	// For Configuration object, any attributes not defined by the user appear as null, unless derived from another instance.
	// We perform IsUnknown() check to handle variables from .tfvars that are resolved
	// during the plan phase rather than validation phase, preventing false validation errors.

	var serverAssociationType string
	niosPath := path.Root("nios")

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

	// If server_association_type is FAILOVER, failover_association field must be set
	if serverAssociationType == "FAILOVER" {
		if m.FailoverAssociation.IsNull() {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("failover_association"),
				"Invalid Configuration",
				"The 'failover_association' field must be set when 'server_association_type' is set to 'FAILOVER'.",
			)
		}
	}

	// If server_association_type is MS_SERVER, ms_server field must be set
	if serverAssociationType == "MS_SERVER" {
		if m.MsServer.IsNull() {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("ms_server"),
				"Invalid Configuration",
				"The 'ms_server' field must be set when 'server_association_type' is set to 'MS_SERVER'.",
			)
		}
	}

	// DHCP options validation
	utils.ValidateDHCPOptionsConfig(ctx, m.Options, niosPath.AtName("options"), &resp.Diagnostics)
}

func PostFlattenRangetemplateNIOS(ctx context.Context, planned, flattened *NIOSRangetemplateModel, diags *diag.Diagnostics) {
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
