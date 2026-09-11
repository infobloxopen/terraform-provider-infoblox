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

type SharedrecordTxtService interface {
	Create(ctx context.Context, obj *dns.SharedrecordTxt, opts *core.Options) (*dns.SharedrecordTxt, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dns.SharedrecordTxt, *http.Response, error)
	Update(ctx context.Context, id string, obj *dns.SharedrecordTxt, opts *core.Options) (*dns.SharedrecordTxt, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dns.SharedrecordTxt, *http.Response, string, error)
}

type sharedrecordTxtService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewSharedrecordTxtService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) SharedrecordTxtService {
	return &sharedrecordTxtService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new SharedrecordTxt and returns the created object
func (s *sharedrecordTxtService) Create(ctx context.Context, obj *dns.SharedrecordTxt, opts *core.Options) (*dns.SharedrecordTxt, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordTxtService) createNIOS(ctx context.Context, obj *dns.SharedrecordTxt, opts *core.Options) (*dns.SharedrecordTxt, *http.Response, error) {
	payload, err := common.MapTo[niosdns.SharedrecordTxt](obj, mapper.SharedrecordTxtNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.SharedrecordTxtAPI.
		Create(ctx).
		SharedrecordTxt(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateSharedrecordTxtResponseAsObject.GetResult()

	return mapNIOSSharedrecordTxtToResponse(&result), httpResp, nil
}

// Read retrieves a SharedrecordTxt by ID
func (s *sharedrecordTxtService) Read(ctx context.Context, id string, opts *core.Options) (*dns.SharedrecordTxt, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordTxtService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dns.SharedrecordTxt, *http.Response, error) {
	req := s.niosClient.DNSAPI.SharedrecordTxtAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetSharedrecordTxtResponseObjectAsResult.GetResult()

	return mapNIOSSharedrecordTxtToResponse(&result), httpResp, nil
}

// Update modifies an existing SharedrecordTxt and returns the updated object
func (s *sharedrecordTxtService) Update(ctx context.Context, id string, obj *dns.SharedrecordTxt, opts *core.Options) (*dns.SharedrecordTxt, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordTxtService) updateNIOS(ctx context.Context, id string, obj *dns.SharedrecordTxt, opts *core.Options) (*dns.SharedrecordTxt, *http.Response, error) {
	payload, err := common.MapTo[niosdns.SharedrecordTxt](obj, mapper.SharedrecordTxtNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.SharedrecordTxtAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		SharedrecordTxt(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateSharedrecordTxtResponseAsObject.GetResult()

	return mapNIOSSharedrecordTxtToResponse(&result), httpResp, nil
}

// Delete removes a SharedrecordTxt by ID
func (s *sharedrecordTxtService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordTxtService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DNSAPI.SharedrecordTxtAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves SharedrecordTxt objects based on filter options
func (s *sharedrecordTxtService) List(ctx context.Context, opts *core.ListOptions) ([]*dns.SharedrecordTxt, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordTxtService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dns.SharedrecordTxt, *http.Response, string, error) {
	req := s.niosClient.DNSAPI.SharedrecordTxtAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.SharedrecordTxtFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListSharedrecordTxtResponseObject.GetResult()
	items := make([]*dns.SharedrecordTxt, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSSharedrecordTxtToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListSharedrecordTxtResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSSharedrecordTxtToResponse(r *niosdns.SharedrecordTxt) *dns.SharedrecordTxt {
	resp := &dns.SharedrecordTxt{
		Id: r.Ref,
	}
	resp.NIOS = &dns.NIOSSharedrecordTxtExt{
		Comment:           r.Comment,
		Disable:           r.Disable,
		Name:              r.Name,
		SharedRecordGroup: r.SharedRecordGroup,
		Text:              r.Text,
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
