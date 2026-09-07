package dhcp

import (
	"context"
	"fmt"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dhcp"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
)

type RangetemplateService interface {
	Create(ctx context.Context, obj *dhcp.Rangetemplate, opts *core.Options) (*dhcp.Rangetemplate, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dhcp.Rangetemplate, *http.Response, error)
	Update(ctx context.Context, id string, obj *dhcp.Rangetemplate, opts *core.Options) (*dhcp.Rangetemplate, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Rangetemplate, *http.Response, string, error)
}

type rangetemplateService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewRangetemplateService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) RangetemplateService {
	return &rangetemplateService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new Rangetemplate and returns the created object
func (s *rangetemplateService) Create(ctx context.Context, obj *dhcp.Rangetemplate, opts *core.Options) (*dhcp.Rangetemplate, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *rangetemplateService) createNIOS(ctx context.Context, obj *dhcp.Rangetemplate, opts *core.Options) (*dhcp.Rangetemplate, *http.Response, error) {
	payload, err := common.MapTo[niosdhcp.Rangetemplate](obj, mapper.RangetemplateNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DHCPAPI.RangetemplateAPI.
		Create(ctx).
		Rangetemplate(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateRangetemplateResponseAsObject.GetResult()

	return mapNIOSRangetemplateToResponse(&result), httpResp, nil
}

// Read retrieves a Rangetemplate by ID
func (s *rangetemplateService) Read(ctx context.Context, id string, opts *core.Options) (*dhcp.Rangetemplate, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *rangetemplateService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dhcp.Rangetemplate, *http.Response, error) {
	req := s.niosClient.DHCPAPI.RangetemplateAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetRangetemplateResponseObjectAsResult.GetResult()

	return mapNIOSRangetemplateToResponse(&result), httpResp, nil
}

// Update modifies an existing Rangetemplate and returns the updated object
func (s *rangetemplateService) Update(ctx context.Context, id string, obj *dhcp.Rangetemplate, opts *core.Options) (*dhcp.Rangetemplate, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *rangetemplateService) updateNIOS(ctx context.Context, id string, obj *dhcp.Rangetemplate, opts *core.Options) (*dhcp.Rangetemplate, *http.Response, error) {
	payload, err := common.MapTo[niosdhcp.Rangetemplate](obj, mapper.RangetemplateNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DHCPAPI.RangetemplateAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Rangetemplate(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateRangetemplateResponseAsObject.GetResult()

	return mapNIOSRangetemplateToResponse(&result), httpResp, nil
}

// Delete removes a Rangetemplate by ID
func (s *rangetemplateService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *rangetemplateService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DHCPAPI.RangetemplateAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves Rangetemplate objects based on filter options
func (s *rangetemplateService) List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Rangetemplate, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *rangetemplateService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Rangetemplate, *http.Response, string, error) {
	req := s.niosClient.DHCPAPI.RangetemplateAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.RangetemplateFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListRangetemplateResponseObject.GetResult()
	items := make([]*dhcp.Rangetemplate, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSRangetemplateToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListRangetemplateResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSRangetemplateToResponse(r *niosdhcp.Rangetemplate) *dhcp.Rangetemplate {
	resp := &dhcp.Rangetemplate{
		Id: r.Ref,
	}
	resp.NIOS = &dhcp.NIOSRangetemplateExt{
		Bootfile:                       r.Bootfile,
		Bootserver:                     r.Bootserver,
		CloudApiCompatible:             r.CloudApiCompatible,
		Comment:                        r.Comment,
		DdnsDomainname:                 r.DdnsDomainname,
		DdnsGenerateHostname:           r.DdnsGenerateHostname,
		DelegatedMember:                r.DelegatedMember,
		DenyAllClients:                 r.DenyAllClients,
		DenyBootp:                      r.DenyBootp,
		EmailList:                      r.EmailList,
		EnableDdns:                     r.EnableDdns,
		EnableDhcpThresholds:           r.EnableDhcpThresholds,
		EnableEmailWarnings:            r.EnableEmailWarnings,
		EnablePxeLeaseTime:             r.EnablePxeLeaseTime,
		EnableSnmpWarnings:             r.EnableSnmpWarnings,
		Exclude:                        r.Exclude,
		FailoverAssociation:            r.FailoverAssociation,
		FingerprintFilterRules:         r.FingerprintFilterRules,
		HighWaterMark:                  r.HighWaterMark,
		HighWaterMarkReset:             r.HighWaterMarkReset,
		IgnoreDhcpOptionListRequest:    r.IgnoreDhcpOptionListRequest,
		KnownClients:                   r.KnownClients,
		LeaseScavengeTime:              r.LeaseScavengeTime,
		LogicFilterRules:               r.LogicFilterRules,
		LowWaterMark:                   r.LowWaterMark,
		LowWaterMarkReset:              r.LowWaterMarkReset,
		MacFilterRules:                 r.MacFilterRules,
		Member:                         r.Member,
		MsOptions:                      r.MsOptions,
		MsServer:                       r.MsServer,
		NacFilterRules:                 r.NacFilterRules,
		Name:                           r.Name,
		Nextserver:                     r.Nextserver,
		NumberOfAddresses:              r.NumberOfAddresses,
		Offset:                         r.Offset,
		OptionFilterRules:              r.OptionFilterRules,
		Options:                        r.Options,
		PxeLeaseTime:                   r.PxeLeaseTime,
		RecycleLeases:                  r.RecycleLeases,
		RelayAgentFilterRules:          r.RelayAgentFilterRules,
		ServerAssociationType:          r.ServerAssociationType,
		UnknownClients:                 r.UnknownClients,
		UpdateDnsOnLeaseRenewal:        r.UpdateDnsOnLeaseRenewal,
		UseBootfile:                    r.UseBootfile,
		UseBootserver:                  r.UseBootserver,
		UseDdnsDomainname:              r.UseDdnsDomainname,
		UseDdnsGenerateHostname:        r.UseDdnsGenerateHostname,
		UseDenyBootp:                   r.UseDenyBootp,
		UseEmailList:                   r.UseEmailList,
		UseEnableDdns:                  r.UseEnableDdns,
		UseEnableDhcpThresholds:        r.UseEnableDhcpThresholds,
		UseIgnoreDhcpOptionListRequest: r.UseIgnoreDhcpOptionListRequest,
		UseKnownClients:                r.UseKnownClients,
		UseLeaseScavengeTime:           r.UseLeaseScavengeTime,
		UseLogicFilterRules:            r.UseLogicFilterRules,
		UseMsOptions:                   r.UseMsOptions,
		UseNextserver:                  r.UseNextserver,
		UseOptions:                     r.UseOptions,
		UsePxeLeaseTime:                r.UsePxeLeaseTime,
		UseRecycleLeases:               r.UseRecycleLeases,
		UseUnknownClients:              r.UseUnknownClients,
		UseUpdateDnsOnLeaseRenewal:     r.UseUpdateDnsOnLeaseRenewal,
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
