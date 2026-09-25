package dhcp

import (
	"context"
	"fmt"
	"maps"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dhcp"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

type RangeService interface {
	Create(ctx context.Context, obj *dhcp.Range, opts *core.Options) (*dhcp.Range, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dhcp.Range, *http.Response, error)
	Update(ctx context.Context, id string, obj *dhcp.Range, opts *core.Options) (*dhcp.Range, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Range, *http.Response, string, error)
}

type rangeService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
	uddiClient *uddiclient.APIClient
}

func NewRangeService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) RangeService {
	return &rangeService{
		backend:    backend,
		niosClient: nios,
		uddiClient: uddi,
	}
}

// Create creates a new Range and returns the created object
func (s *rangeService) Create(ctx context.Context, obj *dhcp.Range, opts *core.Options) (*dhcp.Range, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *rangeService) createNIOS(ctx context.Context, obj *dhcp.Range, opts *core.Options) (*dhcp.Range, *http.Response, error) {
	payload, err := common.MapTo[niosdhcp.Range](obj, mapper.RangeNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DHCPAPI.RangeAPI.
		Create(ctx).
		Range_(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateRangeResponseAsObject.GetResult()

	return mapNIOSRangeToResponse(&result), httpResp, nil
}

func (s *rangeService) createUDDI(ctx context.Context, obj *dhcp.Range, opts *core.Options) (*dhcp.Range, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.Range](obj, mapper.RangeUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAddressManagementAPI.RangeAPI.
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

	return mapUDDIRangeToResponse(&result), httpResp, nil
}

// Read retrieves a Range by ID
func (s *rangeService) Read(ctx context.Context, id string, opts *core.Options) (*dhcp.Range, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *rangeService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dhcp.Range, *http.Response, error) {
	req := s.niosClient.DHCPAPI.RangeAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetRangeResponseObjectAsResult.GetResult()

	return mapNIOSRangeToResponse(&result), httpResp, nil
}

func (s *rangeService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dhcp.Range, *http.Response, error) {
	req := s.uddiClient.IPAddressManagementAPI.RangeAPI.
		Read(ctx, id)

	if opts != nil && opts.Inherit != "" {
		req = req.Inherit(opts.Inherit)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIRangeToResponse(&result), httpResp, nil
}

// Update modifies an existing Range and returns the updated object
func (s *rangeService) Update(ctx context.Context, id string, obj *dhcp.Range, opts *core.Options) (*dhcp.Range, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *rangeService) updateNIOS(ctx context.Context, id string, obj *dhcp.Range, opts *core.Options) (*dhcp.Range, *http.Response, error) {
	payload, err := common.MapTo[niosdhcp.Range](obj, mapper.RangeNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DHCPAPI.RangeAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Range_(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateRangeResponseAsObject.GetResult()

	return mapNIOSRangeToResponse(&result), httpResp, nil
}

func (s *rangeService) updateUDDI(ctx context.Context, id string, obj *dhcp.Range, opts *core.Options) (*dhcp.Range, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.Range](obj, mapper.RangeUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAddressManagementAPI.RangeAPI.
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

	return mapUDDIRangeToResponse(&result), httpResp, nil
}

// Delete removes a Range by ID
func (s *rangeService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *rangeService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DHCPAPI.RangeAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

func (s *rangeService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.IPAddressManagementAPI.RangeAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves Range objects based on filter options
func (s *rangeService) List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Range, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *rangeService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Range, *http.Response, string, error) {
	req := s.niosClient.DHCPAPI.RangeAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.RangeFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListRangeResponseObject.GetResult()
	items := make([]*dhcp.Range, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSRangeToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListRangeResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func (s *rangeService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Range, *http.Response, string, error) {
	req := s.uddiClient.IPAddressManagementAPI.RangeAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.RangeFilterFieldMap[core.BackendUDDI])
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
	items := make([]*dhcp.Range, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIRangeToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapNIOSRangeToResponse(r *niosdhcp.Range) *dhcp.Range {
	resp := &dhcp.Range{
		Id: r.Ref,
	}
	resp.NIOS = &dhcp.NIOSRangeExt{
		AlwaysUpdateDns:                  r.AlwaysUpdateDns,
		Bootfile:                         r.Bootfile,
		Bootserver:                       r.Bootserver,
		CloudInfo:                        r.CloudInfo,
		Comment:                          r.Comment,
		DdnsDomainname:                   r.DdnsDomainname,
		DdnsGenerateHostname:             r.DdnsGenerateHostname,
		DenyAllClients:                   r.DenyAllClients,
		DenyBootp:                        r.DenyBootp,
		Disable:                          r.Disable,
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
		EndAddr:                          r.EndAddr,
		Exclude:                          r.Exclude,
		FailoverAssociation:              r.FailoverAssociation,
		FingerprintFilterRules:           r.FingerprintFilterRules,
		HighWaterMark:                    r.HighWaterMark,
		HighWaterMarkReset:               r.HighWaterMarkReset,
		IgnoreDhcpOptionListRequest:      r.IgnoreDhcpOptionListRequest,
		IgnoreId:                         r.IgnoreId,
		IgnoreMacAddresses:               r.IgnoreMacAddresses,
		KnownClients:                     r.KnownClients,
		LeaseScavengeTime:                r.LeaseScavengeTime,
		LogicFilterRules:                 r.LogicFilterRules,
		LowWaterMark:                     r.LowWaterMark,
		LowWaterMarkReset:                r.LowWaterMarkReset,
		MacFilterRules:                   r.MacFilterRules,
		Member:                           r.Member,
		MsOptions:                        r.MsOptions,
		MsServer:                         r.MsServer,
		NacFilterRules:                   r.NacFilterRules,
		Name:                             r.Name,
		Network:                          r.Network,
		NetworkView:                      r.NetworkView,
		Nextserver:                       r.Nextserver,
		OptionFilterRules:                r.OptionFilterRules,
		Options:                          r.Options,
		PortControlBlackoutSetting:       r.PortControlBlackoutSetting,
		PxeLeaseTime:                     r.PxeLeaseTime,
		RecycleLeases:                    r.RecycleLeases,
		RelayAgentFilterRules:            r.RelayAgentFilterRules,
		RestartIfNeeded:                  r.RestartIfNeeded,
		SamePortControlDiscoveryBlackout: r.SamePortControlDiscoveryBlackout,
		ServerAssociationType:            r.ServerAssociationType,
		SplitMember:                      r.SplitMember,
		SplitScopeExclusionPercent:       r.SplitScopeExclusionPercent,
		StartAddr:                        r.StartAddr,
		SubscribeSettings:                r.SubscribeSettings,
		Template:                         r.Template,
		UnknownClients:                   r.UnknownClients,
		UpdateDnsOnLeaseRenewal:          r.UpdateDnsOnLeaseRenewal,
		UseBlackoutSetting:               r.UseBlackoutSetting,
		UseBootfile:                      r.UseBootfile,
		UseBootserver:                    r.UseBootserver,
		UseDdnsDomainname:                r.UseDdnsDomainname,
		UseDdnsGenerateHostname:          r.UseDdnsGenerateHostname,
		UseDenyBootp:                     r.UseDenyBootp,
		UseDiscoveryBasicPollingSettings: r.UseDiscoveryBasicPollingSettings,
		UseEmailList:                     r.UseEmailList,
		UseEnableDdns:                    r.UseEnableDdns,
		UseEnableDhcpThresholds:          r.UseEnableDhcpThresholds,
		UseEnableDiscovery:               r.UseEnableDiscovery,
		UseEnableIfmapPublishing:         r.UseEnableIfmapPublishing,
		UseIgnoreDhcpOptionListRequest:   r.UseIgnoreDhcpOptionListRequest,
		UseIgnoreId:                      r.UseIgnoreId,
		UseKnownClients:                  r.UseKnownClients,
		UseLeaseScavengeTime:             r.UseLeaseScavengeTime,
		UseLogicFilterRules:              r.UseLogicFilterRules,
		UseMsOptions:                     r.UseMsOptions,
		UseNextserver:                    r.UseNextserver,
		UseOptions:                       r.UseOptions,
		UsePxeLeaseTime:                  r.UsePxeLeaseTime,
		UseRecycleLeases:                 r.UseRecycleLeases,
		UseSubscribeSettings:             r.UseSubscribeSettings,
		UseUnknownClients:                r.UseUnknownClients,
		UseUpdateDnsOnLeaseRenewal:       r.UseUpdateDnsOnLeaseRenewal,
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

func mapUDDIRangeToResponse(r *uddiipam.Range) *dhcp.Range {
	resp := &dhcp.Range{
		Id: r.Id,
	}
	resp.UDDI = &dhcp.UDDIRangeExt{
		Comment:            r.Comment,
		DhcpHost:           r.DhcpHost,
		DhcpOptions:        r.DhcpOptions,
		DisableDhcp:        r.DisableDhcp,
		End:                r.End,
		ExclusionRanges:    r.ExclusionRanges,
		Filters:            r.Filters,
		InheritanceParent:  r.InheritanceParent,
		InheritanceSources: r.InheritanceSources,
		Name:               r.Name,
		Parent:             r.Parent,
		Space:              r.Space,
		Start:              r.Start,
		Threshold:          r.Threshold,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
