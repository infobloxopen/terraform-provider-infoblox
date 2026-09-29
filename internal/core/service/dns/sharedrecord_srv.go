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

type SharedrecordSrvService interface {
	Create(ctx context.Context, obj *dns.SharedrecordSrv, opts *core.Options) (*dns.SharedrecordSrv, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dns.SharedrecordSrv, *http.Response, error)
	Update(ctx context.Context, id string, obj *dns.SharedrecordSrv, opts *core.Options) (*dns.SharedrecordSrv, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dns.SharedrecordSrv, *http.Response, string, error)
}

type sharedrecordSrvService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewSharedrecordSrvService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) SharedrecordSrvService {
	return &sharedrecordSrvService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new SharedrecordSrv and returns the created object
func (s *sharedrecordSrvService) Create(ctx context.Context, obj *dns.SharedrecordSrv, opts *core.Options) (*dns.SharedrecordSrv, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordSrvService) createNIOS(ctx context.Context, obj *dns.SharedrecordSrv, opts *core.Options) (*dns.SharedrecordSrv, *http.Response, error) {
	payload, err := common.MapTo[niosdns.SharedrecordSrv](obj, mapper.SharedrecordSrvNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.SharedrecordSrvAPI.
		Create(ctx).
		SharedrecordSrv(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateSharedrecordSrvResponseAsObject.GetResult()

	return mapNIOSSharedrecordSrvToResponse(&result), httpResp, nil
}

// Read retrieves a SharedrecordSrv by ID
func (s *sharedrecordSrvService) Read(ctx context.Context, id string, opts *core.Options) (*dns.SharedrecordSrv, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordSrvService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dns.SharedrecordSrv, *http.Response, error) {
	req := s.niosClient.DNSAPI.SharedrecordSrvAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetSharedrecordSrvResponseObjectAsResult.GetResult()

	return mapNIOSSharedrecordSrvToResponse(&result), httpResp, nil
}

// Update modifies an existing SharedrecordSrv and returns the updated object
func (s *sharedrecordSrvService) Update(ctx context.Context, id string, obj *dns.SharedrecordSrv, opts *core.Options) (*dns.SharedrecordSrv, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordSrvService) updateNIOS(ctx context.Context, id string, obj *dns.SharedrecordSrv, opts *core.Options) (*dns.SharedrecordSrv, *http.Response, error) {
	payload, err := common.MapTo[niosdns.SharedrecordSrv](obj, mapper.SharedrecordSrvNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.SharedrecordSrvAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		SharedrecordSrv(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateSharedrecordSrvResponseAsObject.GetResult()

	return mapNIOSSharedrecordSrvToResponse(&result), httpResp, nil
}

// Delete removes a SharedrecordSrv by ID
func (s *sharedrecordSrvService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordSrvService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DNSAPI.SharedrecordSrvAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves SharedrecordSrv objects based on filter options
func (s *sharedrecordSrvService) List(ctx context.Context, opts *core.ListOptions) ([]*dns.SharedrecordSrv, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordSrvService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dns.SharedrecordSrv, *http.Response, string, error) {
	req := s.niosClient.DNSAPI.SharedrecordSrvAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.SharedrecordSrvFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListSharedrecordSrvResponseObject.GetResult()
	items := make([]*dns.SharedrecordSrv, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSSharedrecordSrvToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListSharedrecordSrvResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSSharedrecordSrvToResponse(r *niosdns.SharedrecordSrv) *dns.SharedrecordSrv {
	resp := &dns.SharedrecordSrv{
		Id: r.Ref,
	}
	resp.NIOS = &dns.NIOSSharedrecordSrvExt{
		Comment:           r.Comment,
		Disable:           r.Disable,
		Name:              r.Name,
		Port:              r.Port,
		Priority:          r.Priority,
		SharedRecordGroup: r.SharedRecordGroup,
		Target:            r.Target,
		Ttl:               r.Ttl,
		UseTtl:            r.UseTtl,
		Weight:            r.Weight,
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
