package ipam

import (
	"context"
	"fmt"
	"maps"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/ipam"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

type NetworkService interface {
	Create(ctx context.Context, obj *ipam.Network, opts *core.Options) (*ipam.Network, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*ipam.Network, *http.Response, error)
	Update(ctx context.Context, id string, obj *ipam.Network, opts *core.Options) (*ipam.Network, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*ipam.Network, *http.Response, string, error)
}

type networkService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
	uddiClient *uddiclient.APIClient
}

func NewNetworkService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) NetworkService {
	return &networkService{
		backend:    backend,
		niosClient: nios,
		uddiClient: uddi,
	}
}

// Create creates a new Network and returns the created object
func (s *networkService) Create(ctx context.Context, obj *ipam.Network, opts *core.Options) (*ipam.Network, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *networkService) createNIOS(ctx context.Context, obj *ipam.Network, opts *core.Options) (*ipam.Network, *http.Response, error) {
	payload, err := common.MapTo[niosipam.Network](obj, mapper.NetworkNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if payload.FuncCall != nil && payload.Network == nil {
		payload.Network = &niosipam.NetworkNetwork{}
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.IPAMAPI.NetworkAPI.
		Create(ctx).
		Network(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateNetworkResponseAsObject.GetResult()

	return mapNIOSNetworkToResponse(&result), httpResp, nil
}

func (s *networkService) createUDDI(ctx context.Context, obj *ipam.Network, opts *core.Options) (*ipam.Network, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.Subnet](obj, mapper.NetworkUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAddressManagementAPI.SubnetAPI.
		Create(ctx).
		Body(payload)

	if opts != nil && opts.Inherit != "" {
		req = req.Inherit(opts.Inherit)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDINetworkToResponse(&result), httpResp, nil
}

// Read retrieves a Network by ID
func (s *networkService) Read(ctx context.Context, id string, opts *core.Options) (*ipam.Network, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *networkService) readNIOS(ctx context.Context, id string, opts *core.Options) (*ipam.Network, *http.Response, error) {
	req := s.niosClient.IPAMAPI.NetworkAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetNetworkResponseObjectAsResult.GetResult()

	return mapNIOSNetworkToResponse(&result), httpResp, nil
}

func (s *networkService) readUDDI(ctx context.Context, id string, opts *core.Options) (*ipam.Network, *http.Response, error) {
	req := s.uddiClient.IPAddressManagementAPI.SubnetAPI.
		Read(ctx, id)

	if opts != nil && opts.Inherit != "" {
		req = req.Inherit(opts.Inherit)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDINetworkToResponse(&result), httpResp, nil
}

// Update modifies an existing Network and returns the updated object
func (s *networkService) Update(ctx context.Context, id string, obj *ipam.Network, opts *core.Options) (*ipam.Network, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *networkService) updateNIOS(ctx context.Context, id string, obj *ipam.Network, opts *core.Options) (*ipam.Network, *http.Response, error) {
	payload, err := common.MapTo[niosipam.Network](obj, mapper.NetworkNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.IPAMAPI.NetworkAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Network(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateNetworkResponseAsObject.GetResult()

	return mapNIOSNetworkToResponse(&result), httpResp, nil
}

func (s *networkService) updateUDDI(ctx context.Context, id string, obj *ipam.Network, opts *core.Options) (*ipam.Network, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.Subnet](obj, mapper.NetworkUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAddressManagementAPI.SubnetAPI.
		Update(ctx, id).
		Body(payload)

	if opts != nil && opts.Inherit != "" {
		req = req.Inherit(opts.Inherit)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDINetworkToResponse(&result), httpResp, nil
}

// Delete removes a Network by ID
func (s *networkService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *networkService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.IPAMAPI.NetworkAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

func (s *networkService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.IPAddressManagementAPI.SubnetAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves Network objects based on filter options
func (s *networkService) List(ctx context.Context, opts *core.ListOptions) ([]*ipam.Network, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *networkService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*ipam.Network, *http.Response, string, error) {
	req := s.niosClient.IPAMAPI.NetworkAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.NetworkFilterFieldMap[core.BackendNIOS])
			filters := make(map[string]any, len(translatedFilters))
			for k, v := range translatedFilters {
				filters[k] = v
			}
			req = req.Filters(filters)
		}
		if len(opts.ExtAttrFilter) > 0 {
			extAttrFilters := make(map[string]any, len(opts.ExtAttrFilter))
			for k, v := range opts.ExtAttrFilter {
				extAttrFilters[k] = v
			}
			req = req.Extattrfilter(extAttrFilters)
		}
		if opts.PageID != "" {
			req = req.PageId(opts.PageID)
		}
		req = req.Paging(opts.Paging)
		maxResults := opts.MaxResults
		if maxResults <= 0 {
			maxResults = core.DefaultListLimit
		}
		req = req.MaxResults(maxResults)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, "", err
	}

	results := resp.ListNetworkResponseObject.GetResult()
	items := make([]*ipam.Network, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSNetworkToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListNetworkResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func (s *networkService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*ipam.Network, *http.Response, string, error) {
	req := s.uddiClient.IPAddressManagementAPI.SubnetAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.NetworkFilterFieldMap[core.BackendUDDI])
		for k, v := range translatedFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		if len(filters) > 0 {
			req = req.Filter(core.JoinFilters(filters))
		}

		if len(opts.TagFilter) > 0 {
			var tfilters []string
			for k, v := range opts.TagFilter {
				tfilters = append(tfilters, "'"+k+"'=='"+v+"'")
			}
			req = req.Tfilter(core.JoinFilters(tfilters))
		}

		if opts.Offset > 0 {
			req = req.Offset(opts.Offset)
		}

		if opts.Limit > 0 {
			req = req.Limit(opts.Limit)
		}
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, "", err
	}

	results := resp.GetResults()
	items := make([]*ipam.Network, 0, len(results))
	for i := range results {
		items = append(items, mapUDDINetworkToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapNIOSNetworkToResponse(r *niosipam.Network) *ipam.Network {
	resp := &ipam.Network{
		Id: r.Ref,
	}
	resp.NIOS = &ipam.NIOSNetworkExt{
		Authority:                        r.Authority,
		AutoCreateReversezone:            r.AutoCreateReversezone,
		Bootfile:                         r.Bootfile,
		Bootserver:                       r.Bootserver,
		CloudInfo:                        r.CloudInfo,
		CloudShared:                      r.CloudShared,
		Comment:                          r.Comment,
		DdnsDomainname:                   r.DdnsDomainname,
		DdnsGenerateHostname:             r.DdnsGenerateHostname,
		DdnsServerAlwaysUpdates:          r.DdnsServerAlwaysUpdates,
		DdnsTtl:                          r.DdnsTtl,
		DdnsUpdateFixedAddresses:         r.DdnsUpdateFixedAddresses,
		DdnsUseOption81:                  r.DdnsUseOption81,
		DeleteReason:                     r.DeleteReason,
		DenyBootp:                        r.DenyBootp,
		Disable:                          r.Disable,
		DiscoveredBridgeDomain:           r.DiscoveredBridgeDomain,
		DiscoveredTenant:                 r.DiscoveredTenant,
		DiscoveryBasicPollSettings:       r.DiscoveryBasicPollSettings,
		DiscoveryBlackoutSetting:         r.DiscoveryBlackoutSetting,
		DiscoveryMember:                  r.DiscoveryMember,
		EmailList:                        r.EmailList,
		EnableDdns:                       r.EnableDdns,
		EnableDhcpThresholds:             r.EnableDhcpThresholds,
		EnableDiscovery:                  r.EnableDiscovery,
		EnableEmailWarnings:              r.EnableEmailWarnings,
		EnableIfmapPublishing:            r.EnableIfmapPublishing,
		EnableImmediateDiscovery:         r.EnableImmediateDiscovery,
		EnablePxeLeaseTime:               r.EnablePxeLeaseTime,
		EnableSnmpWarnings:               r.EnableSnmpWarnings,
		FederatedRealms:                  r.FederatedRealms,
		HighWaterMark:                    r.HighWaterMark,
		HighWaterMarkReset:               r.HighWaterMarkReset,
		IgnoreDhcpOptionListRequest:      r.IgnoreDhcpOptionListRequest,
		IgnoreId:                         r.IgnoreId,
		IgnoreMacAddresses:               r.IgnoreMacAddresses,
		IpamEmailAddresses:               r.IpamEmailAddresses,
		IpamThresholdSettings:            r.IpamThresholdSettings,
		IpamTrapSettings:                 r.IpamTrapSettings,
		Ipv4addr:                         r.Ipv4addr,
		LeaseScavengeTime:                r.LeaseScavengeTime,
		LogicFilterRules:                 r.LogicFilterRules,
		LowWaterMark:                     r.LowWaterMark,
		LowWaterMarkReset:                r.LowWaterMarkReset,
		Members:                          r.Members,
		MgmPrivate:                       r.MgmPrivate,
		Netmask:                          r.Netmask,
		NetworkView:                      r.NetworkView,
		Nextserver:                       r.Nextserver,
		Options:                          r.Options,
		PortControlBlackoutSetting:       r.PortControlBlackoutSetting,
		PxeLeaseTime:                     r.PxeLeaseTime,
		RecycleLeases:                    r.RecycleLeases,
		RestartIfNeeded:                  r.RestartIfNeeded,
		RirOrganization:                  r.RirOrganization,
		RirRegistrationAction:            r.RirRegistrationAction,
		RirRegistrationStatus:            r.RirRegistrationStatus,
		SamePortControlDiscoveryBlackout: r.SamePortControlDiscoveryBlackout,
		SendRirRequest:                   r.SendRirRequest,
		SubscribeSettings:                r.SubscribeSettings,
		Template:                         r.Template,
		Unmanaged:                        r.Unmanaged,
		UpdateDnsOnLeaseRenewal:          r.UpdateDnsOnLeaseRenewal,
		UseAuthority:                     r.UseAuthority,
		UseBlackoutSetting:               r.UseBlackoutSetting,
		UseBootfile:                      r.UseBootfile,
		UseBootserver:                    r.UseBootserver,
		UseDdnsDomainname:                r.UseDdnsDomainname,
		UseDdnsGenerateHostname:          r.UseDdnsGenerateHostname,
		UseDdnsTtl:                       r.UseDdnsTtl,
		UseDdnsUpdateFixedAddresses:      r.UseDdnsUpdateFixedAddresses,
		UseDdnsUseOption81:               r.UseDdnsUseOption81,
		UseDenyBootp:                     r.UseDenyBootp,
		UseDiscoveryBasicPollingSettings: r.UseDiscoveryBasicPollingSettings,
		UseEmailList:                     r.UseEmailList,
		UseEnableDdns:                    r.UseEnableDdns,
		UseEnableDhcpThresholds:          r.UseEnableDhcpThresholds,
		UseEnableDiscovery:               r.UseEnableDiscovery,
		UseEnableIfmapPublishing:         r.UseEnableIfmapPublishing,
		UseIgnoreDhcpOptionListRequest:   r.UseIgnoreDhcpOptionListRequest,
		UseIgnoreId:                      r.UseIgnoreId,
		UseIpamEmailAddresses:            r.UseIpamEmailAddresses,
		UseIpamThresholdSettings:         r.UseIpamThresholdSettings,
		UseIpamTrapSettings:              r.UseIpamTrapSettings,
		UseLeaseScavengeTime:             r.UseLeaseScavengeTime,
		UseLogicFilterRules:              r.UseLogicFilterRules,
		UseMgmPrivate:                    r.UseMgmPrivate,
		UseNextserver:                    r.UseNextserver,
		UseOptions:                       r.UseOptions,
		UsePxeLeaseTime:                  r.UsePxeLeaseTime,
		UseRecycleLeases:                 r.UseRecycleLeases,
		UseSubscribeSettings:             r.UseSubscribeSettings,
		UseUpdateDnsOnLeaseRenewal:       r.UseUpdateDnsOnLeaseRenewal,
		UseZoneAssociations:              r.UseZoneAssociations,
		Vlans:                            r.Vlans,
		ZoneAssociations:                 r.ZoneAssociations,
	}
	if r.Network != nil {
		resp.NIOS.Network = r.Network.String
	}
	if r.ExtAttrs != nil {
		attrs := make(map[string]any, len(*r.ExtAttrs))
		for k, v := range *r.ExtAttrs {
			attrs[k] = core.StringifyEAValue(v.Value)
		}
		resp.NIOS.ExtAttrs = attrs
	}
	return resp
}

func mapUDDINetworkToResponse(r *uddiipam.Subnet) *ipam.Network {
	resp := &ipam.Network{
		Id: r.Id,
	}
	resp.UDDI = &ipam.UDDINetworkExt{
		Address:                    r.Address,
		AsmConfig:                  r.AsmConfig,
		Cidr:                       r.Cidr,
		Comment:                    r.Comment,
		ConfigProfiles:             r.ConfigProfiles,
		DdnsClientUpdate:           r.DdnsClientUpdate,
		DdnsConflictResolutionMode: r.DdnsConflictResolutionMode,
		DdnsDomain:                 r.DdnsDomain,
		DdnsGenerateName:           r.DdnsGenerateName,
		DdnsGeneratedPrefix:        r.DdnsGeneratedPrefix,
		DdnsSendUpdates:            r.DdnsSendUpdates,
		DdnsTtlPercent:             r.DdnsTtlPercent,
		DdnsUpdateOnRenew:          r.DdnsUpdateOnRenew,
		DdnsUseConflictResolution:  r.DdnsUseConflictResolution,
		DhcpConfig:                 r.DhcpConfig,
		DhcpHost:                   r.DhcpHost,
		DhcpOptions:                r.DhcpOptions,
		DisableDhcp:                r.DisableDhcp,
		ExternalKeys:               r.ExternalKeys,
		FederatedRealms:            r.FederatedRealms,
		HeaderOptionFilename:       r.HeaderOptionFilename,
		HeaderOptionServerAddress:  r.HeaderOptionServerAddress,
		HeaderOptionServerName:     r.HeaderOptionServerName,
		HostnameRewriteChar:        r.HostnameRewriteChar,
		HostnameRewriteEnabled:     r.HostnameRewriteEnabled,
		HostnameRewriteRegex:       r.HostnameRewriteRegex,
		InheritanceParent:          r.InheritanceParent,
		InheritanceSources:         r.InheritanceSources,
		Name:                       r.Name,
		Parent:                     r.Parent,
		RebindTime:                 r.RebindTime,
		RenewTime:                  r.RenewTime,
		Space:                      r.Space,
		Threshold:                  r.Threshold,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
