package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
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
