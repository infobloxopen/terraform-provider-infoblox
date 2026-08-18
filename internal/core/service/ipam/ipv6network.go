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

type Ipv6networkService interface {
	Create(ctx context.Context, obj *ipam.Ipv6network, opts *core.Options) (*ipam.Ipv6network, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*ipam.Ipv6network, *http.Response, error)
	Update(ctx context.Context, id string, obj *ipam.Ipv6network, opts *core.Options) (*ipam.Ipv6network, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*ipam.Ipv6network, *http.Response, string, error)
}

type ipv6networkService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
	uddiClient *uddiclient.APIClient
}

func NewIpv6networkService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) Ipv6networkService {
	return &ipv6networkService{
		backend:    backend,
		niosClient: nios,
		uddiClient: uddi,
	}
}

// Create creates a new Ipv6network and returns the created object
func (s *ipv6networkService) Create(ctx context.Context, obj *ipam.Ipv6network, opts *core.Options) (*ipam.Ipv6network, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6networkService) createNIOS(ctx context.Context, obj *ipam.Ipv6network, opts *core.Options) (*ipam.Ipv6network, *http.Response, error) {
	payload, err := common.MapTo[niosipam.Ipv6network](obj, mapper.Ipv6networkNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if payload.FuncCall != nil && payload.Network == nil {
		payload.Network = &niosipam.Ipv6networkNetwork{}
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.IPAMAPI.Ipv6networkAPI.
		Create(ctx).
		Ipv6network(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateIpv6networkResponseAsObject.GetResult()

	return mapNIOSIpv6networkToResponse(&result), httpResp, nil
}

func (s *ipv6networkService) createUDDI(ctx context.Context, obj *ipam.Ipv6network, opts *core.Options) (*ipam.Ipv6network, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.Subnet](obj, mapper.Ipv6networkUDDIFieldMap)
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

	return mapUDDIIpv6networkToResponse(&result), httpResp, nil
}

// Read retrieves a Ipv6network by ID
func (s *ipv6networkService) Read(ctx context.Context, id string, opts *core.Options) (*ipam.Ipv6network, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6networkService) readNIOS(ctx context.Context, id string, opts *core.Options) (*ipam.Ipv6network, *http.Response, error) {
	req := s.niosClient.IPAMAPI.Ipv6networkAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetIpv6networkResponseObjectAsResult.GetResult()

	return mapNIOSIpv6networkToResponse(&result), httpResp, nil
}

func (s *ipv6networkService) readUDDI(ctx context.Context, id string, opts *core.Options) (*ipam.Ipv6network, *http.Response, error) {
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

	return mapUDDIIpv6networkToResponse(&result), httpResp, nil
}

// Update modifies an existing Ipv6network and returns the updated object
func (s *ipv6networkService) Update(ctx context.Context, id string, obj *ipam.Ipv6network, opts *core.Options) (*ipam.Ipv6network, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6networkService) updateNIOS(ctx context.Context, id string, obj *ipam.Ipv6network, opts *core.Options) (*ipam.Ipv6network, *http.Response, error) {
	payload, err := common.MapTo[niosipam.Ipv6network](obj, mapper.Ipv6networkNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.IPAMAPI.Ipv6networkAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Ipv6network(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateIpv6networkResponseAsObject.GetResult()

	return mapNIOSIpv6networkToResponse(&result), httpResp, nil
}

func (s *ipv6networkService) updateUDDI(ctx context.Context, id string, obj *ipam.Ipv6network, opts *core.Options) (*ipam.Ipv6network, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.Subnet](obj, mapper.Ipv6networkUDDIFieldMap)
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

	return mapUDDIIpv6networkToResponse(&result), httpResp, nil
}

// Delete removes a Ipv6network by ID
func (s *ipv6networkService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6networkService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.IPAMAPI.Ipv6networkAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

func (s *ipv6networkService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.IPAddressManagementAPI.SubnetAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves Ipv6network objects based on filter options
func (s *ipv6networkService) List(ctx context.Context, opts *core.ListOptions) ([]*ipam.Ipv6network, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6networkService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*ipam.Ipv6network, *http.Response, string, error) {
	req := s.niosClient.IPAMAPI.Ipv6networkAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.Ipv6networkFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListIpv6networkResponseObject.GetResult()
	items := make([]*ipam.Ipv6network, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSIpv6networkToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListIpv6networkResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func (s *ipv6networkService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*ipam.Ipv6network, *http.Response, string, error) {
	req := s.uddiClient.IPAddressManagementAPI.SubnetAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.Ipv6networkFilterFieldMap[core.BackendUDDI])
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
	items := make([]*ipam.Ipv6network, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIIpv6networkToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapNIOSIpv6networkToResponse(r *niosipam.Ipv6network) *ipam.Ipv6network {
	resp := &ipam.Ipv6network{
		Id: r.Ref,
	}
	resp.NIOS = &ipam.NIOSIpv6networkExt{
		AutoCreateReversezone:            r.AutoCreateReversezone,
		CloudInfo:                        r.CloudInfo,
		Comment:                          r.Comment,
		DdnsDomainname:                   r.DdnsDomainname,
		DdnsEnableOptionFqdn:             r.DdnsEnableOptionFqdn,
		DdnsGenerateHostname:             r.DdnsGenerateHostname,
		DdnsServerAlwaysUpdates:          r.DdnsServerAlwaysUpdates,
		DdnsTtl:                          r.DdnsTtl,
		DeleteReason:                     r.DeleteReason,
		Disable:                          r.Disable,
		DiscoveredBridgeDomain:           r.DiscoveredBridgeDomain,
		DiscoveredTenant:                 r.DiscoveredTenant,
		DiscoveryBasicPollSettings:       r.DiscoveryBasicPollSettings,
		DiscoveryBlackoutSetting:         r.DiscoveryBlackoutSetting,
		DiscoveryMember:                  r.DiscoveryMember,
		DomainName:                       r.DomainName,
		DomainNameServers:                r.DomainNameServers,
		EnableDdns:                       r.EnableDdns,
		EnableDiscovery:                  r.EnableDiscovery,
		EnableIfmapPublishing:            r.EnableIfmapPublishing,
		EnableImmediateDiscovery:         r.EnableImmediateDiscovery,
		FederatedRealms:                  r.FederatedRealms,
		LogicFilterRules:                 r.LogicFilterRules,
		Members:                          r.Members,
		MgmPrivate:                       r.MgmPrivate,
		NetworkView:                      r.NetworkView,
		Options:                          r.Options,
		PortControlBlackoutSetting:       r.PortControlBlackoutSetting,
		PreferredLifetime:                r.PreferredLifetime,
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
		UseBlackoutSetting:               r.UseBlackoutSetting,
		UseDdnsDomainname:                r.UseDdnsDomainname,
		UseDdnsEnableOptionFqdn:          r.UseDdnsEnableOptionFqdn,
		UseDdnsGenerateHostname:          r.UseDdnsGenerateHostname,
		UseDdnsTtl:                       r.UseDdnsTtl,
		UseDiscoveryBasicPollingSettings: r.UseDiscoveryBasicPollingSettings,
		UseDomainName:                    r.UseDomainName,
		UseDomainNameServers:             r.UseDomainNameServers,
		UseEnableDdns:                    r.UseEnableDdns,
		UseEnableDiscovery:               r.UseEnableDiscovery,
		UseEnableIfmapPublishing:         r.UseEnableIfmapPublishing,
		UseLogicFilterRules:              r.UseLogicFilterRules,
		UseMgmPrivate:                    r.UseMgmPrivate,
		UseOptions:                       r.UseOptions,
		UsePreferredLifetime:             r.UsePreferredLifetime,
		UseRecycleLeases:                 r.UseRecycleLeases,
		UseSubscribeSettings:             r.UseSubscribeSettings,
		UseUpdateDnsOnLeaseRenewal:       r.UseUpdateDnsOnLeaseRenewal,
		UseValidLifetime:                 r.UseValidLifetime,
		UseZoneAssociations:              r.UseZoneAssociations,
		ValidLifetime:                    r.ValidLifetime,
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

func mapUDDIIpv6networkToResponse(r *uddiipam.Subnet) *ipam.Ipv6network {
	resp := &ipam.Ipv6network{
		Id: r.Id,
	}
	resp.UDDI = &ipam.UDDIIpv6networkExt{
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
