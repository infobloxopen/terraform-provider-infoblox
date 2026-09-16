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

type RecordRpzAService interface {
	Create(ctx context.Context, obj *rpz.RecordRpzA, opts *core.Options) (*rpz.RecordRpzA, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*rpz.RecordRpzA, *http.Response, error)
	Update(ctx context.Context, id string, obj *rpz.RecordRpzA, opts *core.Options) (*rpz.RecordRpzA, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*rpz.RecordRpzA, *http.Response, string, error)
}

type recordRpzAService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewRecordRpzAService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) RecordRpzAService {
	return &recordRpzAService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new RecordRpzA and returns the created object
func (s *recordRpzAService) Create(ctx context.Context, obj *rpz.RecordRpzA, opts *core.Options) (*rpz.RecordRpzA, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzAService) createNIOS(ctx context.Context, obj *rpz.RecordRpzA, opts *core.Options) (*rpz.RecordRpzA, *http.Response, error) {
	payload, err := common.MapTo[niosrpz.RecordRpzA](obj, mapper.RecordRpzANIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.RPZAPI.RecordRpzAAPI.
		Create(ctx).
		RecordRpzA(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateRecordRpzAResponseAsObject.GetResult()

	return mapNIOSRecordRpzAToResponse(&result), httpResp, nil
}

// Read retrieves a RecordRpzA by ID
func (s *recordRpzAService) Read(ctx context.Context, id string, opts *core.Options) (*rpz.RecordRpzA, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzAService) readNIOS(ctx context.Context, id string, opts *core.Options) (*rpz.RecordRpzA, *http.Response, error) {
	req := s.niosClient.RPZAPI.RecordRpzAAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetRecordRpzAResponseObjectAsResult.GetResult()

	return mapNIOSRecordRpzAToResponse(&result), httpResp, nil
}

// Update modifies an existing RecordRpzA and returns the updated object
func (s *recordRpzAService) Update(ctx context.Context, id string, obj *rpz.RecordRpzA, opts *core.Options) (*rpz.RecordRpzA, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzAService) updateNIOS(ctx context.Context, id string, obj *rpz.RecordRpzA, opts *core.Options) (*rpz.RecordRpzA, *http.Response, error) {
	payload, err := common.MapTo[niosrpz.RecordRpzA](obj, mapper.RecordRpzANIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.RPZAPI.RecordRpzAAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		RecordRpzA(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateRecordRpzAResponseAsObject.GetResult()

	return mapNIOSRecordRpzAToResponse(&result), httpResp, nil
}

// Delete removes a RecordRpzA by ID
func (s *recordRpzAService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzAService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.RPZAPI.RecordRpzAAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves RecordRpzA objects based on filter options
func (s *recordRpzAService) List(ctx context.Context, opts *core.ListOptions) ([]*rpz.RecordRpzA, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzAService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*rpz.RecordRpzA, *http.Response, string, error) {
	req := s.niosClient.RPZAPI.RecordRpzAAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.RecordRpzAFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListRecordRpzAResponseObject.GetResult()
	items := make([]*rpz.RecordRpzA, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSRecordRpzAToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListRecordRpzAResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSRecordRpzAToResponse(r *niosrpz.RecordRpzA) *rpz.RecordRpzA {
	resp := &rpz.RecordRpzA{
		Id: r.Ref,
	}
	resp.NIOS = &rpz.NIOSRecordRpzAExt{
		Comment:  r.Comment,
		Disable:  r.Disable,
		Ipv4addr: r.Ipv4addr,
		Name:     r.Name,
		RpZone:   r.RpZone,
		Ttl:      r.Ttl,
		UseTtl:   r.UseTtl,
		View:     r.View,
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
