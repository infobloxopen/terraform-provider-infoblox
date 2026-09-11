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

type RecordRpzAaaaService interface {
	Create(ctx context.Context, obj *rpz.RecordRpzAaaa, opts *core.Options) (*rpz.RecordRpzAaaa, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*rpz.RecordRpzAaaa, *http.Response, error)
	Update(ctx context.Context, id string, obj *rpz.RecordRpzAaaa, opts *core.Options) (*rpz.RecordRpzAaaa, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*rpz.RecordRpzAaaa, *http.Response, string, error)
}

type recordRpzAaaaService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewRecordRpzAaaaService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) RecordRpzAaaaService {
	return &recordRpzAaaaService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new RecordRpzAaaa and returns the created object
func (s *recordRpzAaaaService) Create(ctx context.Context, obj *rpz.RecordRpzAaaa, opts *core.Options) (*rpz.RecordRpzAaaa, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzAaaaService) createNIOS(ctx context.Context, obj *rpz.RecordRpzAaaa, opts *core.Options) (*rpz.RecordRpzAaaa, *http.Response, error) {
	payload, err := common.MapTo[niosrpz.RecordRpzAaaa](obj, mapper.RecordRpzAaaaNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.RPZAPI.RecordRpzAaaaAPI.
		Create(ctx).
		RecordRpzAaaa(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateRecordRpzAaaaResponseAsObject.GetResult()

	return mapNIOSRecordRpzAaaaToResponse(&result), httpResp, nil
}

// Read retrieves a RecordRpzAaaa by ID
func (s *recordRpzAaaaService) Read(ctx context.Context, id string, opts *core.Options) (*rpz.RecordRpzAaaa, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzAaaaService) readNIOS(ctx context.Context, id string, opts *core.Options) (*rpz.RecordRpzAaaa, *http.Response, error) {
	req := s.niosClient.RPZAPI.RecordRpzAaaaAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetRecordRpzAaaaResponseObjectAsResult.GetResult()

	return mapNIOSRecordRpzAaaaToResponse(&result), httpResp, nil
}

// Update modifies an existing RecordRpzAaaa and returns the updated object
func (s *recordRpzAaaaService) Update(ctx context.Context, id string, obj *rpz.RecordRpzAaaa, opts *core.Options) (*rpz.RecordRpzAaaa, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzAaaaService) updateNIOS(ctx context.Context, id string, obj *rpz.RecordRpzAaaa, opts *core.Options) (*rpz.RecordRpzAaaa, *http.Response, error) {
	payload, err := common.MapTo[niosrpz.RecordRpzAaaa](obj, mapper.RecordRpzAaaaNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.RPZAPI.RecordRpzAaaaAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		RecordRpzAaaa(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateRecordRpzAaaaResponseAsObject.GetResult()

	return mapNIOSRecordRpzAaaaToResponse(&result), httpResp, nil
}

// Delete removes a RecordRpzAaaa by ID
func (s *recordRpzAaaaService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzAaaaService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.RPZAPI.RecordRpzAaaaAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves RecordRpzAaaa objects based on filter options
func (s *recordRpzAaaaService) List(ctx context.Context, opts *core.ListOptions) ([]*rpz.RecordRpzAaaa, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzAaaaService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*rpz.RecordRpzAaaa, *http.Response, string, error) {
	req := s.niosClient.RPZAPI.RecordRpzAaaaAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.RecordRpzAaaaFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListRecordRpzAaaaResponseObject.GetResult()
	items := make([]*rpz.RecordRpzAaaa, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSRecordRpzAaaaToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListRecordRpzAaaaResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSRecordRpzAaaaToResponse(r *niosrpz.RecordRpzAaaa) *rpz.RecordRpzAaaa {
	resp := &rpz.RecordRpzAaaa{
		Id: r.Ref,
	}
	resp.NIOS = &rpz.NIOSRecordRpzAaaaExt{
		Comment:  r.Comment,
		Disable:  r.Disable,
		Ipv6addr: r.Ipv6addr,
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
