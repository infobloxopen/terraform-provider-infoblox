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

type Ipv6rangetemplateService interface {
	Create(ctx context.Context, obj *dhcp.Ipv6rangetemplate, opts *core.Options) (*dhcp.Ipv6rangetemplate, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dhcp.Ipv6rangetemplate, *http.Response, error)
	Update(ctx context.Context, id string, obj *dhcp.Ipv6rangetemplate, opts *core.Options) (*dhcp.Ipv6rangetemplate, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Ipv6rangetemplate, *http.Response, string, error)
}

type ipv6rangetemplateService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewIpv6rangetemplateService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) Ipv6rangetemplateService {
	return &ipv6rangetemplateService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new Ipv6rangetemplate and returns the created object
func (s *ipv6rangetemplateService) Create(ctx context.Context, obj *dhcp.Ipv6rangetemplate, opts *core.Options) (*dhcp.Ipv6rangetemplate, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6rangetemplateService) createNIOS(ctx context.Context, obj *dhcp.Ipv6rangetemplate, opts *core.Options) (*dhcp.Ipv6rangetemplate, *http.Response, error) {
	payload, err := common.MapTo[niosdhcp.Ipv6rangetemplate](obj, mapper.Ipv6rangetemplateNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.DHCPAPI.Ipv6rangetemplateAPI.
		Create(ctx).
		Ipv6rangetemplate(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateIpv6rangetemplateResponseAsObject.GetResult()

	return mapNIOSIpv6rangetemplateToResponse(&result), httpResp, nil
}

// Read retrieves a Ipv6rangetemplate by ID
func (s *ipv6rangetemplateService) Read(ctx context.Context, id string, opts *core.Options) (*dhcp.Ipv6rangetemplate, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6rangetemplateService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dhcp.Ipv6rangetemplate, *http.Response, error) {
	req := s.niosClient.DHCPAPI.Ipv6rangetemplateAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetIpv6rangetemplateResponseObjectAsResult.GetResult()

	return mapNIOSIpv6rangetemplateToResponse(&result), httpResp, nil
}

// Update modifies an existing Ipv6rangetemplate and returns the updated object
func (s *ipv6rangetemplateService) Update(ctx context.Context, id string, obj *dhcp.Ipv6rangetemplate, opts *core.Options) (*dhcp.Ipv6rangetemplate, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6rangetemplateService) updateNIOS(ctx context.Context, id string, obj *dhcp.Ipv6rangetemplate, opts *core.Options) (*dhcp.Ipv6rangetemplate, *http.Response, error) {
	payload, err := common.MapTo[niosdhcp.Ipv6rangetemplate](obj, mapper.Ipv6rangetemplateNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.DHCPAPI.Ipv6rangetemplateAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Ipv6rangetemplate(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateIpv6rangetemplateResponseAsObject.GetResult()

	return mapNIOSIpv6rangetemplateToResponse(&result), httpResp, nil
}

// Delete removes a Ipv6rangetemplate by ID
func (s *ipv6rangetemplateService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6rangetemplateService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DHCPAPI.Ipv6rangetemplateAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves Ipv6rangetemplate objects based on filter options
func (s *ipv6rangetemplateService) List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Ipv6rangetemplate, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6rangetemplateService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Ipv6rangetemplate, *http.Response, string, error) {
	req := s.niosClient.DHCPAPI.Ipv6rangetemplateAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.Ipv6rangetemplateFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListIpv6rangetemplateResponseObject.GetResult()
	items := make([]*dhcp.Ipv6rangetemplate, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSIpv6rangetemplateToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListIpv6rangetemplateResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSIpv6rangetemplateToResponse(r *niosdhcp.Ipv6rangetemplate) *dhcp.Ipv6rangetemplate {
	resp := &dhcp.Ipv6rangetemplate{
		Id: r.Ref,
	}
	resp.NIOS = &dhcp.NIOSIpv6rangetemplateExt{
		CloudApiCompatible:    r.CloudApiCompatible,
		Comment:               r.Comment,
		DelegatedMember:       r.DelegatedMember,
		Exclude:               r.Exclude,
		LogicFilterRules:      r.LogicFilterRules,
		Member:                r.Member,
		Name:                  r.Name,
		NumberOfAddresses:     r.NumberOfAddresses,
		Offset:                r.Offset,
		OptionFilterRules:     r.OptionFilterRules,
		RecycleLeases:         r.RecycleLeases,
		ServerAssociationType: r.ServerAssociationType,
		UseLogicFilterRules:   r.UseLogicFilterRules,
		UseRecycleLeases:      r.UseRecycleLeases,
	}
	return resp
}
