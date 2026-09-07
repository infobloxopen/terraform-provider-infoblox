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

type SharedrecordAaaaService interface {
	Create(ctx context.Context, obj *dns.SharedrecordAaaa, opts *core.Options) (*dns.SharedrecordAaaa, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dns.SharedrecordAaaa, *http.Response, error)
	Update(ctx context.Context, id string, obj *dns.SharedrecordAaaa, opts *core.Options) (*dns.SharedrecordAaaa, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dns.SharedrecordAaaa, *http.Response, string, error)
}

type sharedrecordAaaaService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewSharedrecordAaaaService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) SharedrecordAaaaService {
	return &sharedrecordAaaaService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new SharedrecordAaaa and returns the created object
func (s *sharedrecordAaaaService) Create(ctx context.Context, obj *dns.SharedrecordAaaa, opts *core.Options) (*dns.SharedrecordAaaa, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordAaaaService) createNIOS(ctx context.Context, obj *dns.SharedrecordAaaa, opts *core.Options) (*dns.SharedrecordAaaa, *http.Response, error) {
	payload, err := common.MapTo[niosdns.SharedrecordAaaa](obj, mapper.SharedrecordAaaaNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.SharedrecordAaaaAPI.
		Create(ctx).
		SharedrecordAaaa(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateSharedrecordAaaaResponseAsObject.GetResult()

	return mapNIOSSharedrecordAaaaToResponse(&result), httpResp, nil
}

// Read retrieves a SharedrecordAaaa by ID
func (s *sharedrecordAaaaService) Read(ctx context.Context, id string, opts *core.Options) (*dns.SharedrecordAaaa, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordAaaaService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dns.SharedrecordAaaa, *http.Response, error) {
	req := s.niosClient.DNSAPI.SharedrecordAaaaAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetSharedrecordAaaaResponseObjectAsResult.GetResult()

	return mapNIOSSharedrecordAaaaToResponse(&result), httpResp, nil
}

// Update modifies an existing SharedrecordAaaa and returns the updated object
func (s *sharedrecordAaaaService) Update(ctx context.Context, id string, obj *dns.SharedrecordAaaa, opts *core.Options) (*dns.SharedrecordAaaa, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordAaaaService) updateNIOS(ctx context.Context, id string, obj *dns.SharedrecordAaaa, opts *core.Options) (*dns.SharedrecordAaaa, *http.Response, error) {
	payload, err := common.MapTo[niosdns.SharedrecordAaaa](obj, mapper.SharedrecordAaaaNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.SharedrecordAaaaAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		SharedrecordAaaa(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateSharedrecordAaaaResponseAsObject.GetResult()

	return mapNIOSSharedrecordAaaaToResponse(&result), httpResp, nil
}

// Delete removes a SharedrecordAaaa by ID
func (s *sharedrecordAaaaService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordAaaaService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DNSAPI.SharedrecordAaaaAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves SharedrecordAaaa objects based on filter options
func (s *sharedrecordAaaaService) List(ctx context.Context, opts *core.ListOptions) ([]*dns.SharedrecordAaaa, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordAaaaService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dns.SharedrecordAaaa, *http.Response, string, error) {
	req := s.niosClient.DNSAPI.SharedrecordAaaaAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.SharedrecordAaaaFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListSharedrecordAaaaResponseObject.GetResult()
	items := make([]*dns.SharedrecordAaaa, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSSharedrecordAaaaToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListSharedrecordAaaaResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSSharedrecordAaaaToResponse(r *niosdns.SharedrecordAaaa) *dns.SharedrecordAaaa {
	resp := &dns.SharedrecordAaaa{
		Id: r.Ref,
	}
	resp.NIOS = &dns.NIOSSharedrecordAaaaExt{
		Comment:           r.Comment,
		Disable:           r.Disable,
		Ipv6addr:          r.Ipv6addr,
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
