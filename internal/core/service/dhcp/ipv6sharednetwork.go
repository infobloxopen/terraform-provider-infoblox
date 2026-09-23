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

type Ipv6sharednetworkService interface {
	Create(ctx context.Context, obj *dhcp.Ipv6sharednetwork, opts *core.Options) (*dhcp.Ipv6sharednetwork, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dhcp.Ipv6sharednetwork, *http.Response, error)
	Update(ctx context.Context, id string, obj *dhcp.Ipv6sharednetwork, opts *core.Options) (*dhcp.Ipv6sharednetwork, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Ipv6sharednetwork, *http.Response, string, error)
}

type ipv6sharednetworkService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewIpv6sharednetworkService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) Ipv6sharednetworkService {
	return &ipv6sharednetworkService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new Ipv6sharednetwork and returns the created object
func (s *ipv6sharednetworkService) Create(ctx context.Context, obj *dhcp.Ipv6sharednetwork, opts *core.Options) (*dhcp.Ipv6sharednetwork, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6sharednetworkService) createNIOS(ctx context.Context, obj *dhcp.Ipv6sharednetwork, opts *core.Options) (*dhcp.Ipv6sharednetwork, *http.Response, error) {
	payload, err := common.MapTo[niosdhcp.Ipv6sharednetwork](obj, mapper.Ipv6sharednetworkNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DHCPAPI.Ipv6sharednetworkAPI.
		Create(ctx).
		Ipv6sharednetwork(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateIpv6sharednetworkResponseAsObject.GetResult()

	return mapNIOSIpv6sharednetworkToResponse(&result), httpResp, nil
}

// Read retrieves a Ipv6sharednetwork by ID
func (s *ipv6sharednetworkService) Read(ctx context.Context, id string, opts *core.Options) (*dhcp.Ipv6sharednetwork, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6sharednetworkService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dhcp.Ipv6sharednetwork, *http.Response, error) {
	req := s.niosClient.DHCPAPI.Ipv6sharednetworkAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetIpv6sharednetworkResponseObjectAsResult.GetResult()

	return mapNIOSIpv6sharednetworkToResponse(&result), httpResp, nil
}

// Update modifies an existing Ipv6sharednetwork and returns the updated object
func (s *ipv6sharednetworkService) Update(ctx context.Context, id string, obj *dhcp.Ipv6sharednetwork, opts *core.Options) (*dhcp.Ipv6sharednetwork, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6sharednetworkService) updateNIOS(ctx context.Context, id string, obj *dhcp.Ipv6sharednetwork, opts *core.Options) (*dhcp.Ipv6sharednetwork, *http.Response, error) {
	payload, err := common.MapTo[niosdhcp.Ipv6sharednetwork](obj, mapper.Ipv6sharednetworkNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DHCPAPI.Ipv6sharednetworkAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Ipv6sharednetwork(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateIpv6sharednetworkResponseAsObject.GetResult()

	return mapNIOSIpv6sharednetworkToResponse(&result), httpResp, nil
}

// Delete removes a Ipv6sharednetwork by ID
func (s *ipv6sharednetworkService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6sharednetworkService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DHCPAPI.Ipv6sharednetworkAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves Ipv6sharednetwork objects based on filter options
func (s *ipv6sharednetworkService) List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Ipv6sharednetwork, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6sharednetworkService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Ipv6sharednetwork, *http.Response, string, error) {
	req := s.niosClient.DHCPAPI.Ipv6sharednetworkAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.Ipv6sharednetworkFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListIpv6sharednetworkResponseObject.GetResult()
	items := make([]*dhcp.Ipv6sharednetwork, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSIpv6sharednetworkToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListIpv6sharednetworkResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSIpv6sharednetworkToResponse(r *niosdhcp.Ipv6sharednetwork) *dhcp.Ipv6sharednetwork {
	resp := &dhcp.Ipv6sharednetwork{
		Id: r.Ref,
	}
	resp.NIOS = &dhcp.NIOSIpv6sharednetworkExt{
		Comment:                    r.Comment,
		DdnsDomainname:             r.DdnsDomainname,
		DdnsGenerateHostname:       r.DdnsGenerateHostname,
		DdnsServerAlwaysUpdates:    r.DdnsServerAlwaysUpdates,
		DdnsTtl:                    r.DdnsTtl,
		DdnsUseOption81:            r.DdnsUseOption81,
		Disable:                    r.Disable,
		DomainName:                 r.DomainName,
		DomainNameServers:          r.DomainNameServers,
		EnableDdns:                 r.EnableDdns,
		LogicFilterRules:           r.LogicFilterRules,
		Name:                       r.Name,
		NetworkView:                r.NetworkView,
		Networks:                   r.Networks,
		Options:                    r.Options,
		PreferredLifetime:          r.PreferredLifetime,
		UpdateDnsOnLeaseRenewal:    r.UpdateDnsOnLeaseRenewal,
		UseDdnsDomainname:          r.UseDdnsDomainname,
		UseDdnsGenerateHostname:    r.UseDdnsGenerateHostname,
		UseDdnsTtl:                 r.UseDdnsTtl,
		UseDdnsUseOption81:         r.UseDdnsUseOption81,
		UseDomainName:              r.UseDomainName,
		UseDomainNameServers:       r.UseDomainNameServers,
		UseEnableDdns:              r.UseEnableDdns,
		UseLogicFilterRules:        r.UseLogicFilterRules,
		UseOptions:                 r.UseOptions,
		UsePreferredLifetime:       r.UsePreferredLifetime,
		UseUpdateDnsOnLeaseRenewal: r.UseUpdateDnsOnLeaseRenewal,
		UseValidLifetime:           r.UseValidLifetime,
		ValidLifetime:              r.ValidLifetime,
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
