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

type Ipv6fixedaddresstemplateService interface {
	Create(ctx context.Context, obj *dhcp.Ipv6fixedaddresstemplate, opts *core.Options) (*dhcp.Ipv6fixedaddresstemplate, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dhcp.Ipv6fixedaddresstemplate, *http.Response, error)
	Update(ctx context.Context, id string, obj *dhcp.Ipv6fixedaddresstemplate, opts *core.Options) (*dhcp.Ipv6fixedaddresstemplate, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Ipv6fixedaddresstemplate, *http.Response, string, error)
}

type ipv6fixedaddresstemplateService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewIpv6fixedaddresstemplateService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) Ipv6fixedaddresstemplateService {
	return &ipv6fixedaddresstemplateService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new Ipv6fixedaddresstemplate and returns the created object
func (s *ipv6fixedaddresstemplateService) Create(ctx context.Context, obj *dhcp.Ipv6fixedaddresstemplate, opts *core.Options) (*dhcp.Ipv6fixedaddresstemplate, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6fixedaddresstemplateService) createNIOS(ctx context.Context, obj *dhcp.Ipv6fixedaddresstemplate, opts *core.Options) (*dhcp.Ipv6fixedaddresstemplate, *http.Response, error) {
	payload, err := common.MapTo[niosdhcp.Ipv6fixedaddresstemplate](obj, mapper.Ipv6fixedaddresstemplateNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DHCPAPI.Ipv6fixedaddresstemplateAPI.
		Create(ctx).
		Ipv6fixedaddresstemplate(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateIpv6fixedaddresstemplateResponseAsObject.GetResult()

	return mapNIOSIpv6fixedaddresstemplateToResponse(&result), httpResp, nil
}

// Read retrieves a Ipv6fixedaddresstemplate by ID
func (s *ipv6fixedaddresstemplateService) Read(ctx context.Context, id string, opts *core.Options) (*dhcp.Ipv6fixedaddresstemplate, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6fixedaddresstemplateService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dhcp.Ipv6fixedaddresstemplate, *http.Response, error) {
	req := s.niosClient.DHCPAPI.Ipv6fixedaddresstemplateAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetIpv6fixedaddresstemplateResponseObjectAsResult.GetResult()

	return mapNIOSIpv6fixedaddresstemplateToResponse(&result), httpResp, nil
}

// Update modifies an existing Ipv6fixedaddresstemplate and returns the updated object
func (s *ipv6fixedaddresstemplateService) Update(ctx context.Context, id string, obj *dhcp.Ipv6fixedaddresstemplate, opts *core.Options) (*dhcp.Ipv6fixedaddresstemplate, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6fixedaddresstemplateService) updateNIOS(ctx context.Context, id string, obj *dhcp.Ipv6fixedaddresstemplate, opts *core.Options) (*dhcp.Ipv6fixedaddresstemplate, *http.Response, error) {
	payload, err := common.MapTo[niosdhcp.Ipv6fixedaddresstemplate](obj, mapper.Ipv6fixedaddresstemplateNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DHCPAPI.Ipv6fixedaddresstemplateAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Ipv6fixedaddresstemplate(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateIpv6fixedaddresstemplateResponseAsObject.GetResult()

	return mapNIOSIpv6fixedaddresstemplateToResponse(&result), httpResp, nil
}

// Delete removes a Ipv6fixedaddresstemplate by ID
func (s *ipv6fixedaddresstemplateService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6fixedaddresstemplateService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DHCPAPI.Ipv6fixedaddresstemplateAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves Ipv6fixedaddresstemplate objects based on filter options
func (s *ipv6fixedaddresstemplateService) List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Ipv6fixedaddresstemplate, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6fixedaddresstemplateService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Ipv6fixedaddresstemplate, *http.Response, string, error) {
	req := s.niosClient.DHCPAPI.Ipv6fixedaddresstemplateAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.Ipv6fixedaddresstemplateFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListIpv6fixedaddresstemplateResponseObject.GetResult()
	items := make([]*dhcp.Ipv6fixedaddresstemplate, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSIpv6fixedaddresstemplateToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListIpv6fixedaddresstemplateResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSIpv6fixedaddresstemplateToResponse(r *niosdhcp.Ipv6fixedaddresstemplate) *dhcp.Ipv6fixedaddresstemplate {
	resp := &dhcp.Ipv6fixedaddresstemplate{
		Id: r.Ref,
	}
	resp.NIOS = &dhcp.NIOSIpv6fixedaddresstemplateExt{
		Comment:              r.Comment,
		DomainName:           r.DomainName,
		DomainNameServers:    r.DomainNameServers,
		LogicFilterRules:     r.LogicFilterRules,
		Name:                 r.Name,
		NumberOfAddresses:    r.NumberOfAddresses,
		Offset:               r.Offset,
		Options:              r.Options,
		PreferredLifetime:    r.PreferredLifetime,
		UseDomainName:        r.UseDomainName,
		UseDomainNameServers: r.UseDomainNameServers,
		UseLogicFilterRules:  r.UseLogicFilterRules,
		UseOptions:           r.UseOptions,
		UsePreferredLifetime: r.UsePreferredLifetime,
		UseValidLifetime:     r.UseValidLifetime,
		ValidLifetime:        r.ValidLifetime,
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
