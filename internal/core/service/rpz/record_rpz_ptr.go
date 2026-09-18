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

type RecordRpzPtrService interface {
	Create(ctx context.Context, obj *rpz.RecordRpzPtr, opts *core.Options) (*rpz.RecordRpzPtr, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*rpz.RecordRpzPtr, *http.Response, error)
	Update(ctx context.Context, id string, obj *rpz.RecordRpzPtr, opts *core.Options) (*rpz.RecordRpzPtr, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*rpz.RecordRpzPtr, *http.Response, string, error)
}

type recordRpzPtrService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewRecordRpzPtrService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) RecordRpzPtrService {
	return &recordRpzPtrService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new RecordRpzPtr and returns the created object
func (s *recordRpzPtrService) Create(ctx context.Context, obj *rpz.RecordRpzPtr, opts *core.Options) (*rpz.RecordRpzPtr, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzPtrService) createNIOS(ctx context.Context, obj *rpz.RecordRpzPtr, opts *core.Options) (*rpz.RecordRpzPtr, *http.Response, error) {
	payload, err := common.MapTo[niosrpz.RecordRpzPtr](obj, mapper.RecordRpzPtrNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.RPZAPI.RecordRpzPtrAPI.
		Create(ctx).
		RecordRpzPtr(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateRecordRpzPtrResponseAsObject.GetResult()

	return mapNIOSRecordRpzPtrToResponse(&result), httpResp, nil
}

// Read retrieves a RecordRpzPtr by ID
func (s *recordRpzPtrService) Read(ctx context.Context, id string, opts *core.Options) (*rpz.RecordRpzPtr, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzPtrService) readNIOS(ctx context.Context, id string, opts *core.Options) (*rpz.RecordRpzPtr, *http.Response, error) {
	req := s.niosClient.RPZAPI.RecordRpzPtrAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetRecordRpzPtrResponseObjectAsResult.GetResult()

	return mapNIOSRecordRpzPtrToResponse(&result), httpResp, nil
}

// Update modifies an existing RecordRpzPtr and returns the updated object
func (s *recordRpzPtrService) Update(ctx context.Context, id string, obj *rpz.RecordRpzPtr, opts *core.Options) (*rpz.RecordRpzPtr, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzPtrService) updateNIOS(ctx context.Context, id string, obj *rpz.RecordRpzPtr, opts *core.Options) (*rpz.RecordRpzPtr, *http.Response, error) {
	payload, err := common.MapTo[niosrpz.RecordRpzPtr](obj, mapper.RecordRpzPtrNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.RPZAPI.RecordRpzPtrAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		RecordRpzPtr(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateRecordRpzPtrResponseAsObject.GetResult()

	return mapNIOSRecordRpzPtrToResponse(&result), httpResp, nil
}

// Delete removes a RecordRpzPtr by ID
func (s *recordRpzPtrService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzPtrService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.RPZAPI.RecordRpzPtrAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves RecordRpzPtr objects based on filter options
func (s *recordRpzPtrService) List(ctx context.Context, opts *core.ListOptions) ([]*rpz.RecordRpzPtr, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzPtrService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*rpz.RecordRpzPtr, *http.Response, string, error) {
	req := s.niosClient.RPZAPI.RecordRpzPtrAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.RecordRpzPtrFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListRecordRpzPtrResponseObject.GetResult()
	items := make([]*rpz.RecordRpzPtr, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSRecordRpzPtrToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListRecordRpzPtrResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSRecordRpzPtrToResponse(r *niosrpz.RecordRpzPtr) *rpz.RecordRpzPtr {
	resp := &rpz.RecordRpzPtr{
		Id: r.Ref,
	}
	resp.NIOS = &rpz.NIOSRecordRpzPtrExt{
		Comment:  r.Comment,
		Disable:  r.Disable,
		Ipv4addr: r.Ipv4addr,
		Ipv6addr: r.Ipv6addr,
		Name:     r.Name,
		Ptrdname: r.Ptrdname,
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
