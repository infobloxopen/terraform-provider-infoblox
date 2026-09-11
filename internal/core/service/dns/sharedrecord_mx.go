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

type SharedrecordMxService interface {
	Create(ctx context.Context, obj *dns.SharedrecordMx, opts *core.Options) (*dns.SharedrecordMx, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dns.SharedrecordMx, *http.Response, error)
	Update(ctx context.Context, id string, obj *dns.SharedrecordMx, opts *core.Options) (*dns.SharedrecordMx, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dns.SharedrecordMx, *http.Response, string, error)
}

type sharedrecordMxService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewSharedrecordMxService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) SharedrecordMxService {
	return &sharedrecordMxService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new SharedrecordMx and returns the created object
func (s *sharedrecordMxService) Create(ctx context.Context, obj *dns.SharedrecordMx, opts *core.Options) (*dns.SharedrecordMx, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordMxService) createNIOS(ctx context.Context, obj *dns.SharedrecordMx, opts *core.Options) (*dns.SharedrecordMx, *http.Response, error) {
	payload, err := common.MapTo[niosdns.SharedrecordMx](obj, mapper.SharedrecordMxNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.SharedrecordMxAPI.
		Create(ctx).
		SharedrecordMx(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateSharedrecordMxResponseAsObject.GetResult()

	return mapNIOSSharedrecordMxToResponse(&result), httpResp, nil
}

// Read retrieves a SharedrecordMx by ID
func (s *sharedrecordMxService) Read(ctx context.Context, id string, opts *core.Options) (*dns.SharedrecordMx, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordMxService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dns.SharedrecordMx, *http.Response, error) {
	req := s.niosClient.DNSAPI.SharedrecordMxAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetSharedrecordMxResponseObjectAsResult.GetResult()

	return mapNIOSSharedrecordMxToResponse(&result), httpResp, nil
}

// Update modifies an existing SharedrecordMx and returns the updated object
func (s *sharedrecordMxService) Update(ctx context.Context, id string, obj *dns.SharedrecordMx, opts *core.Options) (*dns.SharedrecordMx, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordMxService) updateNIOS(ctx context.Context, id string, obj *dns.SharedrecordMx, opts *core.Options) (*dns.SharedrecordMx, *http.Response, error) {
	payload, err := common.MapTo[niosdns.SharedrecordMx](obj, mapper.SharedrecordMxNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.SharedrecordMxAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		SharedrecordMx(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateSharedrecordMxResponseAsObject.GetResult()

	return mapNIOSSharedrecordMxToResponse(&result), httpResp, nil
}

// Delete removes a SharedrecordMx by ID
func (s *sharedrecordMxService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordMxService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DNSAPI.SharedrecordMxAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves SharedrecordMx objects based on filter options
func (s *sharedrecordMxService) List(ctx context.Context, opts *core.ListOptions) ([]*dns.SharedrecordMx, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordMxService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dns.SharedrecordMx, *http.Response, string, error) {
	req := s.niosClient.DNSAPI.SharedrecordMxAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.SharedrecordMxFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListSharedrecordMxResponseObject.GetResult()
	items := make([]*dns.SharedrecordMx, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSSharedrecordMxToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListSharedrecordMxResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSSharedrecordMxToResponse(r *niosdns.SharedrecordMx) *dns.SharedrecordMx {
	resp := &dns.SharedrecordMx{
		Id: r.Ref,
	}
	resp.NIOS = &dns.NIOSSharedrecordMxExt{
		Comment:           r.Comment,
		Disable:           r.Disable,
		MailExchanger:     r.MailExchanger,
		Name:              r.Name,
		Preference:        r.Preference,
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
