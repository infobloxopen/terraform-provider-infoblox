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

type NetworkcontainerService interface {
	Create(ctx context.Context, obj *ipam.Networkcontainer, opts *core.Options) (*ipam.Networkcontainer, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*ipam.Networkcontainer, *http.Response, error)
	Update(ctx context.Context, id string, obj *ipam.Networkcontainer, opts *core.Options) (*ipam.Networkcontainer, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*ipam.Networkcontainer, *http.Response, string, error)
}

type networkcontainerService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
	uddiClient *uddiclient.APIClient
}

func NewNetworkcontainerService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) NetworkcontainerService {
	return &networkcontainerService{
		backend:    backend,
		niosClient: nios,
		uddiClient: uddi,
	}
}

// Create creates a new Networkcontainer and returns the created object
func (s *networkcontainerService) Create(ctx context.Context, obj *ipam.Networkcontainer, opts *core.Options) (*ipam.Networkcontainer, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *networkcontainerService) createNIOS(ctx context.Context, obj *ipam.Networkcontainer, opts *core.Options) (*ipam.Networkcontainer, *http.Response, error) {
	payload, err := common.MapTo[niosipam.Networkcontainer](obj, mapper.NetworkcontainerNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if payload.FuncCall != nil && payload.Network == nil {
		payload.Network = &niosipam.NetworkcontainerNetwork{}
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.IPAMAPI.NetworkcontainerAPI.
		Create(ctx).
		Networkcontainer(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateNetworkcontainerResponseAsObject.GetResult()

	return mapNIOSNetworkcontainerToResponse(&result), httpResp, nil
}

func (s *networkcontainerService) createUDDI(ctx context.Context, obj *ipam.Networkcontainer, opts *core.Options) (*ipam.Networkcontainer, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.AddressBlock](obj, mapper.NetworkcontainerUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAddressManagementAPI.AddressBlockAPI.
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

	return mapUDDINetworkcontainerToResponse(&result), httpResp, nil
}

// Read retrieves a Networkcontainer by ID
func (s *networkcontainerService) Read(ctx context.Context, id string, opts *core.Options) (*ipam.Networkcontainer, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *networkcontainerService) readNIOS(ctx context.Context, id string, opts *core.Options) (*ipam.Networkcontainer, *http.Response, error) {
	req := s.niosClient.IPAMAPI.NetworkcontainerAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetNetworkcontainerResponseObjectAsResult.GetResult()

	return mapNIOSNetworkcontainerToResponse(&result), httpResp, nil
}

func (s *networkcontainerService) readUDDI(ctx context.Context, id string, opts *core.Options) (*ipam.Networkcontainer, *http.Response, error) {
	req := s.uddiClient.IPAddressManagementAPI.AddressBlockAPI.
		Read(ctx, id)

	if opts != nil && opts.Inherit != "" {
		req = req.Inherit(opts.Inherit)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDINetworkcontainerToResponse(&result), httpResp, nil
}

// Update modifies an existing Networkcontainer and returns the updated object
func (s *networkcontainerService) Update(ctx context.Context, id string, obj *ipam.Networkcontainer, opts *core.Options) (*ipam.Networkcontainer, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *networkcontainerService) updateNIOS(ctx context.Context, id string, obj *ipam.Networkcontainer, opts *core.Options) (*ipam.Networkcontainer, *http.Response, error) {
	payload, err := common.MapTo[niosipam.Networkcontainer](obj, mapper.NetworkcontainerNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.IPAMAPI.NetworkcontainerAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Networkcontainer(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateNetworkcontainerResponseAsObject.GetResult()

	return mapNIOSNetworkcontainerToResponse(&result), httpResp, nil
}

func (s *networkcontainerService) updateUDDI(ctx context.Context, id string, obj *ipam.Networkcontainer, opts *core.Options) (*ipam.Networkcontainer, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.AddressBlock](obj, mapper.NetworkcontainerUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAddressManagementAPI.AddressBlockAPI.
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

	return mapUDDINetworkcontainerToResponse(&result), httpResp, nil
}

// Delete removes a Networkcontainer by ID
func (s *networkcontainerService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *networkcontainerService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.IPAMAPI.NetworkcontainerAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

func (s *networkcontainerService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.IPAddressManagementAPI.AddressBlockAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves Networkcontainer objects based on filter options
func (s *networkcontainerService) List(ctx context.Context, opts *core.ListOptions) ([]*ipam.Networkcontainer, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *networkcontainerService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*ipam.Networkcontainer, *http.Response, string, error) {
	req := s.niosClient.IPAMAPI.NetworkcontainerAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.NetworkcontainerFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListNetworkcontainerResponseObject.GetResult()
	items := make([]*ipam.Networkcontainer, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSNetworkcontainerToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListNetworkcontainerResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func (s *networkcontainerService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*ipam.Networkcontainer, *http.Response, string, error) {
	req := s.uddiClient.IPAddressManagementAPI.AddressBlockAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.NetworkcontainerFilterFieldMap[core.BackendUDDI])
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
	items := make([]*ipam.Networkcontainer, 0, len(results))
	for i := range results {
		items = append(items, mapUDDINetworkcontainerToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapNIOSNetworkcontainerToResponse(r *niosipam.Networkcontainer) *ipam.Networkcontainer {
	resp := &ipam.Networkcontainer{
		Id: r.Ref,
	}
	resp.NIOS = &ipam.NIOSNetworkcontainerExt{
		Authority:                        r.Authority,
		AutoCreateReversezone:            r.AutoCreateReversezone,
		Bootfile:                         r.Bootfile,
		Bootserver:                       r.Bootserver,
		CloudInfo:                        r.CloudInfo,
		Comment:                          r.Comment,
		DdnsDomainname:                   r.DdnsDomainname,
		DdnsGenerateHostname:             r.DdnsGenerateHostname,
		DdnsServerAlwaysUpdates:          r.DdnsServerAlwaysUpdates,
		DdnsTtl:                          r.DdnsTtl,
		DdnsUpdateFixedAddresses:         r.DdnsUpdateFixedAddresses,
		DdnsUseOption81:                  r.DdnsUseOption81,
		DeleteReason:                     r.DeleteReason,
		DenyBootp:                        r.DenyBootp,
		DiscoveryBasicPollSettings:       r.DiscoveryBasicPollSettings,
		DiscoveryBlackoutSetting:         r.DiscoveryBlackoutSetting,
		DiscoveryMember:                  r.DiscoveryMember,
		EmailList:                        r.EmailList,
		EnableDdns:                       r.EnableDdns,
		EnableDhcpThresholds:             r.EnableDhcpThresholds,
		EnableDiscovery:                  r.EnableDiscovery,
		EnableEmailWarnings:              r.EnableEmailWarnings,
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
		LeaseScavengeTime:                r.LeaseScavengeTime,
		LogicFilterRules:                 r.LogicFilterRules,
		LowWaterMark:                     r.LowWaterMark,
		LowWaterMarkReset:                r.LowWaterMarkReset,
		MgmPrivate:                       r.MgmPrivate,
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

func mapUDDINetworkcontainerToResponse(r *uddiipam.AddressBlock) *ipam.Networkcontainer {
	resp := &ipam.Networkcontainer{
		Id: r.Id,
	}
	resp.UDDI = &ipam.UDDINetworkcontainerExt{
		Address:                    r.Address,
		AsmConfig:                  r.AsmConfig,
		Cidr:                       r.Cidr,
		Comment:                    r.Comment,
		CompartmentId:              r.CompartmentId,
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
		DhcpOptions:                r.DhcpOptions,
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
