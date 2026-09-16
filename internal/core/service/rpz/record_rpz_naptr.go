package rpz

import (
	"context"
	"fmt"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niosrpz "github.com/infobloxopen/infoblox-nios-go-client/rpz"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/rpz"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/rpz"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
)

type RecordRpzNaptrService interface {
	Create(ctx context.Context, obj *rpz.RecordRpzNaptr, opts *core.Options) (*rpz.RecordRpzNaptr, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*rpz.RecordRpzNaptr, *http.Response, error)
	Update(ctx context.Context, id string, obj *rpz.RecordRpzNaptr, opts *core.Options) (*rpz.RecordRpzNaptr, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*rpz.RecordRpzNaptr, *http.Response, string, error)
}

type recordRpzNaptrService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewRecordRpzNaptrService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) RecordRpzNaptrService {
	return &recordRpzNaptrService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new RecordRpzNaptr and returns the created object
func (s *recordRpzNaptrService) Create(ctx context.Context, obj *rpz.RecordRpzNaptr, opts *core.Options) (*rpz.RecordRpzNaptr, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzNaptrService) createNIOS(ctx context.Context, obj *rpz.RecordRpzNaptr, opts *core.Options) (*rpz.RecordRpzNaptr, *http.Response, error) {
	payload, err := common.MapTo[niosrpz.RecordRpzNaptr](obj, mapper.RecordRpzNaptrNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.RPZAPI.RecordRpzNaptrAPI.
		Create(ctx).
		RecordRpzNaptr(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateRecordRpzNaptrResponseAsObject.GetResult()

	return mapNIOSRecordRpzNaptrToResponse(&result), httpResp, nil
}

// Read retrieves a RecordRpzNaptr by ID
func (s *recordRpzNaptrService) Read(ctx context.Context, id string, opts *core.Options) (*rpz.RecordRpzNaptr, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzNaptrService) readNIOS(ctx context.Context, id string, opts *core.Options) (*rpz.RecordRpzNaptr, *http.Response, error) {
	req := s.niosClient.RPZAPI.RecordRpzNaptrAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetRecordRpzNaptrResponseObjectAsResult.GetResult()

	return mapNIOSRecordRpzNaptrToResponse(&result), httpResp, nil
}

// Update modifies an existing RecordRpzNaptr and returns the updated object
func (s *recordRpzNaptrService) Update(ctx context.Context, id string, obj *rpz.RecordRpzNaptr, opts *core.Options) (*rpz.RecordRpzNaptr, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzNaptrService) updateNIOS(ctx context.Context, id string, obj *rpz.RecordRpzNaptr, opts *core.Options) (*rpz.RecordRpzNaptr, *http.Response, error) {
	payload, err := common.MapTo[niosrpz.RecordRpzNaptr](obj, mapper.RecordRpzNaptrNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.RPZAPI.RecordRpzNaptrAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		RecordRpzNaptr(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateRecordRpzNaptrResponseAsObject.GetResult()

	return mapNIOSRecordRpzNaptrToResponse(&result), httpResp, nil
}

// Delete removes a RecordRpzNaptr by ID
func (s *recordRpzNaptrService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzNaptrService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.RPZAPI.RecordRpzNaptrAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves RecordRpzNaptr objects based on filter options
func (s *recordRpzNaptrService) List(ctx context.Context, opts *core.ListOptions) ([]*rpz.RecordRpzNaptr, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzNaptrService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*rpz.RecordRpzNaptr, *http.Response, string, error) {
	req := s.niosClient.RPZAPI.RecordRpzNaptrAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.RecordRpzNaptrFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListRecordRpzNaptrResponseObject.GetResult()
	items := make([]*rpz.RecordRpzNaptr, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSRecordRpzNaptrToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListRecordRpzNaptrResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSRecordRpzNaptrToResponse(r *niosrpz.RecordRpzNaptr) *rpz.RecordRpzNaptr {
	resp := &rpz.RecordRpzNaptr{
		Id: r.Ref,
	}
	resp.NIOS = &rpz.NIOSRecordRpzNaptrExt{
		Comment:     r.Comment,
		Disable:     r.Disable,
		Flags:       r.Flags,
		Name:        r.Name,
		Order:       r.Order,
		Preference:  r.Preference,
		Regexp:      r.Regexp,
		Replacement: r.Replacement,
		RpZone:      r.RpZone,
		Services:    r.Services,
		Ttl:         r.Ttl,
		UseTtl:      r.UseTtl,
		View:        r.View,
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
