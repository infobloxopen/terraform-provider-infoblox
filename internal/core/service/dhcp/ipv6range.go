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

type Ipv6rangeService interface {
	Create(ctx context.Context, obj *dhcp.Ipv6range, opts *core.Options) (*dhcp.Ipv6range, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dhcp.Ipv6range, *http.Response, error)
	Update(ctx context.Context, id string, obj *dhcp.Ipv6range, opts *core.Options) (*dhcp.Ipv6range, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Ipv6range, *http.Response, string, error)
}

type ipv6rangeService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
	uddiClient *uddiclient.APIClient
}

func NewIpv6rangeService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) Ipv6rangeService {
	return &ipv6rangeService{
		backend:    backend,
		niosClient: nios,
		uddiClient: uddi,
	}
}

// Create creates a new Ipv6range and returns the created object
func (s *ipv6rangeService) Create(ctx context.Context, obj *dhcp.Ipv6range, opts *core.Options) (*dhcp.Ipv6range, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6rangeService) createNIOS(ctx context.Context, obj *dhcp.Ipv6range, opts *core.Options) (*dhcp.Ipv6range, *http.Response, error) {
	payload, err := common.MapTo[niosdhcp.Ipv6range](obj, mapper.Ipv6rangeNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DHCPAPI.Ipv6rangeAPI.
		Create(ctx).
		Ipv6range(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateIpv6rangeResponseAsObject.GetResult()

	return mapNIOSIpv6rangeToResponse(&result), httpResp, nil
}

func (s *ipv6rangeService) createUDDI(ctx context.Context, obj *dhcp.Ipv6range, opts *core.Options) (*dhcp.Ipv6range, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.Range](obj, mapper.Ipv6rangeUDDIFieldMap)
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

	return mapUDDIIpv6rangeToResponse(&result), httpResp, nil
}

// Read retrieves a Ipv6range by ID
func (s *ipv6rangeService) Read(ctx context.Context, id string, opts *core.Options) (*dhcp.Ipv6range, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6rangeService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dhcp.Ipv6range, *http.Response, error) {
	req := s.niosClient.DHCPAPI.Ipv6rangeAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	req = req.ProxySearch(core.GetProxySearch())
	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetIpv6rangeResponseObjectAsResult.GetResult()

	return mapNIOSIpv6rangeToResponse(&result), httpResp, nil
}

func (s *ipv6rangeService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dhcp.Ipv6range, *http.Response, error) {
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

	return mapUDDIIpv6rangeToResponse(&result), httpResp, nil
}

// Update modifies an existing Ipv6range and returns the updated object
func (s *ipv6rangeService) Update(ctx context.Context, id string, obj *dhcp.Ipv6range, opts *core.Options) (*dhcp.Ipv6range, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6rangeService) updateNIOS(ctx context.Context, id string, obj *dhcp.Ipv6range, opts *core.Options) (*dhcp.Ipv6range, *http.Response, error) {
	payload, err := common.MapTo[niosdhcp.Ipv6range](obj, mapper.Ipv6rangeNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DHCPAPI.Ipv6rangeAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Ipv6range(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateIpv6rangeResponseAsObject.GetResult()

	return mapNIOSIpv6rangeToResponse(&result), httpResp, nil
}

func (s *ipv6rangeService) updateUDDI(ctx context.Context, id string, obj *dhcp.Ipv6range, opts *core.Options) (*dhcp.Ipv6range, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.Range](obj, mapper.Ipv6rangeUDDIFieldMap)
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

	return mapUDDIIpv6rangeToResponse(&result), httpResp, nil
}

// Delete removes a Ipv6range by ID
func (s *ipv6rangeService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6rangeService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DHCPAPI.Ipv6rangeAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

func (s *ipv6rangeService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.IPAddressManagementAPI.RangeAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves Ipv6range objects based on filter options
func (s *ipv6rangeService) List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Ipv6range, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6rangeService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Ipv6range, *http.Response, string, error) {
	req := s.niosClient.DHCPAPI.Ipv6rangeAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.Ipv6rangeFilterFieldMap[core.BackendNIOS])
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

	req = req.ProxySearch(core.GetProxySearch())
	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, "", err
	}

	results := resp.ListIpv6rangeResponseObject.GetResult()
	items := make([]*dhcp.Ipv6range, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSIpv6rangeToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListIpv6rangeResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func (s *ipv6rangeService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Ipv6range, *http.Response, string, error) {
	req := s.uddiClient.IPAddressManagementAPI.RangeAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.Ipv6rangeFilterFieldMap[core.BackendUDDI])
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
	items := make([]*dhcp.Ipv6range, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIIpv6rangeToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapNIOSIpv6rangeToResponse(r *niosdhcp.Ipv6range) *dhcp.Ipv6range {
	resp := &dhcp.Ipv6range{
		Id: r.Ref,
	}
	resp.NIOS = &dhcp.NIOSIpv6rangeExt{
		AddressType:                      r.AddressType,
		CloudInfo:                        r.CloudInfo,
		Comment:                          r.Comment,
		Disable:                          r.Disable,
		DiscoveryBasicPollSettings:       r.DiscoveryBasicPollSettings,
		DiscoveryBlackoutSetting:         r.DiscoveryBlackoutSetting,
		DiscoveryMember:                  r.DiscoveryMember,
		EnableDiscovery:                  r.EnableDiscovery,
		EnableImmediateDiscovery:         r.EnableImmediateDiscovery,
		EndAddr:                          r.EndAddr,
		Exclude:                          r.Exclude,
		Ipv6EndPrefix:                    r.Ipv6EndPrefix,
		Ipv6PrefixBits:                   r.Ipv6PrefixBits,
		Ipv6StartPrefix:                  r.Ipv6StartPrefix,
		LogicFilterRules:                 r.LogicFilterRules,
		Member:                           r.Member,
		Name:                             r.Name,
		Network:                          r.Network,
		NetworkView:                      r.NetworkView,
		OptionFilterRules:                r.OptionFilterRules,
		PortControlBlackoutSetting:       r.PortControlBlackoutSetting,
		RecycleLeases:                    r.RecycleLeases,
		RestartIfNeeded:                  r.RestartIfNeeded,
		SamePortControlDiscoveryBlackout: r.SamePortControlDiscoveryBlackout,
		ServerAssociationType:            r.ServerAssociationType,
		StartAddr:                        r.StartAddr,
		SubscribeSettings:                r.SubscribeSettings,
		Template:                         r.Template,
		UseBlackoutSetting:               r.UseBlackoutSetting,
		UseDiscoveryBasicPollingSettings: r.UseDiscoveryBasicPollingSettings,
		UseEnableDiscovery:               r.UseEnableDiscovery,
		UseLogicFilterRules:              r.UseLogicFilterRules,
		UseRecycleLeases:                 r.UseRecycleLeases,
		UseSubscribeSettings:             r.UseSubscribeSettings,
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

func mapUDDIIpv6rangeToResponse(r *uddiipam.Range) *dhcp.Ipv6range {
	resp := &dhcp.Ipv6range{
		Id: r.Id,
	}
	resp.UDDI = &dhcp.UDDIIpv6rangeExt{
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
