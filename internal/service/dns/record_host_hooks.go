package dns

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/dynamicallocation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateRecordHost validates the RecordHost configuration.
func ValidateRecordHost(ctx context.Context, data RecordHostModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordHostModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordHostNIOSConfig(ctx, nios, resp)
	}
}

func validateRecordHostNIOSConfig(ctx context.Context, m *NIOSRecordHostModel, resp *resource.ValidateConfigResponse) {
	// A host record needs an address, and NIOS gives it at most one of each family.
	v4, v6 := m.Ipv4addrs, m.Ipv6addrs
	if !v4.IsUnknown() && !v6.IsUnknown() && len(v4.Elements()) == 0 && len(v6.Elements()) == 0 {
		resp.Diagnostics.AddError(
			"Invalid Configuration",
			"At least one of 'ipv4addrs' or 'ipv6addrs' must be configured.",
		)
	}
	if !v4.IsUnknown() && len(v4.Elements()) > 1 {
		resp.Diagnostics.AddError("Invalid Configuration", "'ipv4addrs' can contain at most one element.")
	}
	if !v6.IsUnknown() && len(v6.Elements()) > 1 {
		resp.Diagnostics.AddError("Invalid Configuration", "'ipv6addrs' can contain at most one element.")
	}
}

func BuildRecordHostIpv4addrFuncCall(ctx context.Context, data types.Object, diags *diag.Diagnostics) *niosdns.FuncCall {
	m := nextAvailableIp(ctx, data, diags)
	if m == nil {
		return nil
	}
	return m.FuncCall(ctx, "Ipv4addr", "network", diags)
}

func BuildRecordHostIpv6addrFuncCall(ctx context.Context, data types.Object, diags *diag.Diagnostics) *niosdns.FuncCall {
	m := nextAvailableIp(ctx, data, diags)
	if m == nil {
		return nil
	}
	return m.FuncCall(ctx, "Ipv6addr", "ipv6network", diags)
}

func nextAvailableIp(ctx context.Context, data types.Object, diags *diag.Diagnostics) *dynamicallocation.NextAvailableIpModel {
	if data.IsNull() || data.IsUnknown() {
		return nil
	}
	var m dynamicallocation.NextAvailableIpModel
	diags.Append(data.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return &m
}

// DHCP address attributes that infoblox_ip_association writes. The host record is
// shared, so an update here must send back whatever is currently on the object
// instead of clearing it.
var (
	dhcpSettingsOwnedByAssociationV4 = []string{"mac", "configure_for_dhcp", "match_client"}
	dhcpSettingsOwnedByAssociationV6 = []string{"mac", "duid", "configure_for_dhcp", "match_client"}
)

// refreshRecordHostId re-finds the record when its ref has gone stale - the ref embeds
// the name, so any rename by the association resource sharing this host record invalidates it -
// and carries the association's DHCP settings into the plan before it is expanded.
func (r *RecordHostResource) refreshRecordHostId(ctx context.Context, resp *resource.UpdateResponse, data, stateData *RecordHostModel) {
	if r.backend != core.BackendNIOS {
		return
	}

	current, httpResp, err := r.service.Read(ctx, data.Id.ValueString(), &core.Options{
		ReturnFields: RecordHostReturnFields,
	})
	if err != nil {
		if httpResp == nil || httpResp.StatusCode != http.StatusNotFound {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read RecordHost before update: %s", err))
			return
		}
		if current = r.findRecordHostByInternalID(ctx, stateData, &resp.Diagnostics); current == nil {
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Not Found", "RecordHost was not found by ref or by internal ID.")
			}
			return
		}
		data.Id = types.StringValue(*current.Id)
	}

	if current.NIOS == nil {
		return
	}

	planNIOS := flex.ExpandNestedObject[NIOSRecordHostModel](ctx, data.NIOS, &resp.Diagnostics)
	if resp.Diagnostics.HasError() || planNIOS == nil {
		return
	}

	liveV4 := flex.FlattenFrameworkListNestedBlock(ctx, current.NIOS.Ipv4addrs, RecordHostIpv4addrAttrTypes, &resp.Diagnostics, FlattenRecordHostIpv4addr)
	liveV6 := flex.FlattenFrameworkListNestedBlock(ctx, current.NIOS.Ipv6addrs, RecordHostIpv6addrAttrTypes, &resp.Diagnostics, FlattenRecordHostIpv6addr)
	preserveDhcpSettings(ctx, &planNIOS.Ipv4addrs, liveV4, dhcpSettingsOwnedByAssociationV4)
	preserveDhcpSettings(ctx, &planNIOS.Ipv6addrs, liveV6, dhcpSettingsOwnedByAssociationV6)

	data.NIOS = flex.FlattenNestedObject(ctx, planNIOS, NIOSRecordHostAttrTypes, &resp.Diagnostics)
}

// preserveDhcpSettings copies the given attributes from the addresses as they currently
// exist onto the planned ones, so an update here does not clear what the association set.
func preserveDhcpSettings(ctx context.Context, plan *types.List, live types.List, fields []string) {
	for _, field := range fields {
		preserved, d := utils.CopyFieldFromPlanToRespList(ctx, live, *plan, field)
		if d.HasError() {
			return
		}
		if list, ok := preserved.(types.List); ok {
			*plan = list
		}
	}
}

// findRecordHostByInternalID locates the record by the Terraform Internal ID extensible attribute.
func (r *RecordHostResource) findRecordHostByInternalID(ctx context.Context, stateData *RecordHostModel, diags *diag.Diagnostics) *coremodel.RecordHost {
	stateNIOS := flex.ExpandNestedObject[NIOSRecordHostModel](ctx, stateData.NIOS, diags)
	if stateNIOS == nil || stateNIOS.ExtAttrsAll.IsNull() || stateNIOS.ExtAttrsAll.IsUnknown() {
		return nil
	}
	internalID, ok := stateNIOS.ExtAttrsAll.Elements()[flex.TerraformInternalID]
	if !ok {
		return nil
	}

	records, _, _, err := r.service.List(ctx, &core.ListOptions{
		ReturnFields:  RecordHostReturnFields,
		ExtAttrFilter: map[string]string{flex.TerraformInternalID: internalID.(types.String).ValueString()},
	})
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to search RecordHost by internal ID after its ref went stale: %s", err))
		return nil
	}
	if len(records) == 0 || records[0].Id == nil {
		return nil
	}
	return records[0]
}

// deleteRecordHostByInternalID runs when deleting by the stored ref returned 404.
func (r *RecordHostResource) deleteRecordHostByInternalID(ctx context.Context, resp *resource.DeleteResponse, data *RecordHostModel) {
	if r.backend != core.BackendNIOS {
		return
	}

	current := r.findRecordHostByInternalID(ctx, data, &resp.Diagnostics)
	if current == nil {
		// Either the search failed, which is already reported, or the record really
		// is gone and the 404 was the truth.
		return
	}

	if _, err := r.service.Delete(ctx, *current.Id); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete RecordHost found by internal ID: %s", err))
	}
}

// PostExpandRecordHostNIOS strips name and view from the WAPI payload when configure_for_dns is false
func PostExpandRecordHostNIOS(_ context.Context, obj *coremodel.NIOSRecordHostExt, _ *diag.Diagnostics) *coremodel.NIOSRecordHostExt {
	if obj == nil || obj.ConfigureForDns == nil || *obj.ConfigureForDns {
		return obj
	}
	obj.Name = nil
	obj.View = nil
	return obj
}

// PostFlattenRecordHostNIOS restores the values the backend never returns, which
// are lost during flattening because they sit inside a nested block - a top-level
// field would keep its planned value, but a nested one is rebuilt from scratch.
func PostFlattenRecordHostNIOS(ctx context.Context, planned, flattened *NIOSRecordHostModel, diags *diag.Diagnostics) {
	if planned == nil || flattened == nil {
		return
	}

	// When configure_for_dns is false NIOS doesn't return name and view.
	if !flattened.ConfigureForDns.IsNull() && !flattened.ConfigureForDns.IsUnknown() && !flattened.ConfigureForDns.ValueBool() {
		flattened.Name = planned.Name
		flattened.View = planned.View
	}

	for _, addrs := range []struct{ planned, flattened *types.List }{
		{&planned.Ipv4addrs, &flattened.Ipv4addrs},
		{&planned.Ipv6addrs, &flattened.Ipv6addrs},
	} {
		restored, d := utils.CopyFieldFromPlanToRespList(ctx, *addrs.planned, *addrs.flattened, "dynamic_allocation")
		diags.Append(*d...)
		if d.HasError() {
			return
		}
		if list, ok := restored.(types.List); ok {
			*addrs.flattened = list
		}
	}

	// Secrets are write-only - NIOS accepts them but never echoes them back.
	if restored, d := utils.CopyFieldFromPlanToRespList(ctx, planned.CliCredentials, flattened.CliCredentials, "password"); !d.HasError() {
		if list, ok := restored.(types.List); ok {
			flattened.CliCredentials = list
		}
	}

	// snmp3_credential - record:host omits the whole block from read responses, so
	// fall back to the full planned object rather than trying to inject fields.
	if (flattened.Snmp3Credential.IsNull() || flattened.Snmp3Credential.IsUnknown()) &&
		!planned.Snmp3Credential.IsNull() && !planned.Snmp3Credential.IsUnknown() {
		flattened.Snmp3Credential = planned.Snmp3Credential
	} else {
		for _, field := range []string{"authentication_password", "privacy_password"} {
			if restored, d := utils.CopyFieldFromPlanToRespObject(ctx, planned.Snmp3Credential, flattened.Snmp3Credential, field); !d.HasError() {
				if obj, ok := restored.(types.Object); ok {
					flattened.Snmp3Credential = obj
				}
			}
		}
	}
}
