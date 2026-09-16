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

type SharedrecordAService interface {
	Create(ctx context.Context, obj *dns.SharedrecordA, opts *core.Options) (*dns.SharedrecordA, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dns.SharedrecordA, *http.Response, error)
	Update(ctx context.Context, id string, obj *dns.SharedrecordA, opts *core.Options) (*dns.SharedrecordA, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dns.SharedrecordA, *http.Response, string, error)
}

type sharedrecordAService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewSharedrecordAService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) SharedrecordAService {
	return &sharedrecordAService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new SharedrecordA and returns the created object
func (s *sharedrecordAService) Create(ctx context.Context, obj *dns.SharedrecordA, opts *core.Options) (*dns.SharedrecordA, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordAService) createNIOS(ctx context.Context, obj *dns.SharedrecordA, opts *core.Options) (*dns.SharedrecordA, *http.Response, error) {
	payload, err := common.MapTo[niosdns.SharedrecordA](obj, mapper.SharedrecordANIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.SharedrecordAAPI.
		Create(ctx).
		SharedrecordA(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateSharedrecordAResponseAsObject.GetResult()

	return mapNIOSSharedrecordAToResponse(&result), httpResp, nil
}

// Read retrieves a SharedrecordA by ID
func (s *sharedrecordAService) Read(ctx context.Context, id string, opts *core.Options) (*dns.SharedrecordA, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordAService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dns.SharedrecordA, *http.Response, error) {
	req := s.niosClient.DNSAPI.SharedrecordAAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetSharedrecordAResponseObjectAsResult.GetResult()

	return mapNIOSSharedrecordAToResponse(&result), httpResp, nil
}

// Update modifies an existing SharedrecordA and returns the updated object
func (s *sharedrecordAService) Update(ctx context.Context, id string, obj *dns.SharedrecordA, opts *core.Options) (*dns.SharedrecordA, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordAService) updateNIOS(ctx context.Context, id string, obj *dns.SharedrecordA, opts *core.Options) (*dns.SharedrecordA, *http.Response, error) {
	payload, err := common.MapTo[niosdns.SharedrecordA](obj, mapper.SharedrecordANIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.SharedrecordAAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		SharedrecordA(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateSharedrecordAResponseAsObject.GetResult()

	return mapNIOSSharedrecordAToResponse(&result), httpResp, nil
}

// Delete removes a SharedrecordA by ID
func (s *sharedrecordAService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordAService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DNSAPI.SharedrecordAAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves SharedrecordA objects based on filter options
func (s *sharedrecordAService) List(ctx context.Context, opts *core.ListOptions) ([]*dns.SharedrecordA, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordAService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dns.SharedrecordA, *http.Response, string, error) {
	req := s.niosClient.DNSAPI.SharedrecordAAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.SharedrecordAFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListSharedrecordAResponseObject.GetResult()
	items := make([]*dns.SharedrecordA, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSSharedrecordAToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListSharedrecordAResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSSharedrecordAToResponse(r *niosdns.SharedrecordA) *dns.SharedrecordA {
	resp := &dns.SharedrecordA{
		Id: r.Ref,
	}
	resp.NIOS = &dns.NIOSSharedrecordAExt{
		Comment:           r.Comment,
		Disable:           r.Disable,
		Ipv4addr:          r.Ipv4addr,
		Name:              r.Name,
		SharedRecordGroup: r.SharedRecordGroup,
		Ttl:               r.Ttl,
		UseTtl:            r.UseTtl,
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
