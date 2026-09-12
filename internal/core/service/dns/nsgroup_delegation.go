package dns

import (
	"context"
	"fmt"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
)

type NsgroupDelegationService interface {
	Create(ctx context.Context, obj *dns.NsgroupDelegation, opts *core.Options) (*dns.NsgroupDelegation, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dns.NsgroupDelegation, *http.Response, error)
	Update(ctx context.Context, id string, obj *dns.NsgroupDelegation, opts *core.Options) (*dns.NsgroupDelegation, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dns.NsgroupDelegation, *http.Response, string, error)
}

type nsgroupDelegationService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewNsgroupDelegationService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) NsgroupDelegationService {
	return &nsgroupDelegationService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new NsgroupDelegation and returns the created object
func (s *nsgroupDelegationService) Create(ctx context.Context, obj *dns.NsgroupDelegation, opts *core.Options) (*dns.NsgroupDelegation, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *nsgroupDelegationService) createNIOS(ctx context.Context, obj *dns.NsgroupDelegation, opts *core.Options) (*dns.NsgroupDelegation, *http.Response, error) {
	payload, err := common.MapTo[niosdns.NsgroupDelegation](obj, mapper.NsgroupDelegationNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.NsgroupDelegationAPI.
		Create(ctx).
		NsgroupDelegation(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateNsgroupDelegationResponseAsObject.GetResult()

	return mapNIOSNsgroupDelegationToResponse(&result), httpResp, nil
}

// Read retrieves a NsgroupDelegation by ID
func (s *nsgroupDelegationService) Read(ctx context.Context, id string, opts *core.Options) (*dns.NsgroupDelegation, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *nsgroupDelegationService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dns.NsgroupDelegation, *http.Response, error) {
	req := s.niosClient.DNSAPI.NsgroupDelegationAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetNsgroupDelegationResponseObjectAsResult.GetResult()

	return mapNIOSNsgroupDelegationToResponse(&result), httpResp, nil
}

// Update modifies an existing NsgroupDelegation and returns the updated object
func (s *nsgroupDelegationService) Update(ctx context.Context, id string, obj *dns.NsgroupDelegation, opts *core.Options) (*dns.NsgroupDelegation, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *nsgroupDelegationService) updateNIOS(ctx context.Context, id string, obj *dns.NsgroupDelegation, opts *core.Options) (*dns.NsgroupDelegation, *http.Response, error) {
	payload, err := common.MapTo[niosdns.NsgroupDelegation](obj, mapper.NsgroupDelegationNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.NsgroupDelegationAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		NsgroupDelegation(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateNsgroupDelegationResponseAsObject.GetResult()

	return mapNIOSNsgroupDelegationToResponse(&result), httpResp, nil
}

// Delete removes a NsgroupDelegation by ID
func (s *nsgroupDelegationService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *nsgroupDelegationService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DNSAPI.NsgroupDelegationAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves NsgroupDelegation objects based on filter options
func (s *nsgroupDelegationService) List(ctx context.Context, opts *core.ListOptions) ([]*dns.NsgroupDelegation, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *nsgroupDelegationService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dns.NsgroupDelegation, *http.Response, string, error) {
	req := s.niosClient.DNSAPI.NsgroupDelegationAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.NsgroupDelegationFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListNsgroupDelegationResponseObject.GetResult()
	items := make([]*dns.NsgroupDelegation, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSNsgroupDelegationToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListNsgroupDelegationResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSNsgroupDelegationToResponse(r *niosdns.NsgroupDelegation) *dns.NsgroupDelegation {
	resp := &dns.NsgroupDelegation{
		Id: r.Ref,
	}
	resp.NIOS = &dns.NIOSNsgroupDelegationExt{
		Comment:    r.Comment,
		DelegateTo: r.DelegateTo,
		Name:       r.Name,
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
