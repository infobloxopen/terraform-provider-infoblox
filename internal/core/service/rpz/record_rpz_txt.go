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

type RecordRpzTxtService interface {
	Create(ctx context.Context, obj *rpz.RecordRpzTxt, opts *core.Options) (*rpz.RecordRpzTxt, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*rpz.RecordRpzTxt, *http.Response, error)
	Update(ctx context.Context, id string, obj *rpz.RecordRpzTxt, opts *core.Options) (*rpz.RecordRpzTxt, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*rpz.RecordRpzTxt, *http.Response, string, error)
}

type recordRpzTxtService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewRecordRpzTxtService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) RecordRpzTxtService {
	return &recordRpzTxtService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new RecordRpzTxt and returns the created object
func (s *recordRpzTxtService) Create(ctx context.Context, obj *rpz.RecordRpzTxt, opts *core.Options) (*rpz.RecordRpzTxt, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzTxtService) createNIOS(ctx context.Context, obj *rpz.RecordRpzTxt, opts *core.Options) (*rpz.RecordRpzTxt, *http.Response, error) {
	payload, err := common.MapTo[niosrpz.RecordRpzTxt](obj, mapper.RecordRpzTxtNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.RPZAPI.RecordRpzTxtAPI.
		Create(ctx).
		RecordRpzTxt(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateRecordRpzTxtResponseAsObject.GetResult()

	return mapNIOSRecordRpzTxtToResponse(&result), httpResp, nil
}

// Read retrieves a RecordRpzTxt by ID
func (s *recordRpzTxtService) Read(ctx context.Context, id string, opts *core.Options) (*rpz.RecordRpzTxt, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzTxtService) readNIOS(ctx context.Context, id string, opts *core.Options) (*rpz.RecordRpzTxt, *http.Response, error) {
	req := s.niosClient.RPZAPI.RecordRpzTxtAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetRecordRpzTxtResponseObjectAsResult.GetResult()

	return mapNIOSRecordRpzTxtToResponse(&result), httpResp, nil
}

// Update modifies an existing RecordRpzTxt and returns the updated object
func (s *recordRpzTxtService) Update(ctx context.Context, id string, obj *rpz.RecordRpzTxt, opts *core.Options) (*rpz.RecordRpzTxt, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzTxtService) updateNIOS(ctx context.Context, id string, obj *rpz.RecordRpzTxt, opts *core.Options) (*rpz.RecordRpzTxt, *http.Response, error) {
	payload, err := common.MapTo[niosrpz.RecordRpzTxt](obj, mapper.RecordRpzTxtNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.RPZAPI.RecordRpzTxtAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		RecordRpzTxt(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateRecordRpzTxtResponseAsObject.GetResult()

	return mapNIOSRecordRpzTxtToResponse(&result), httpResp, nil
}

// Delete removes a RecordRpzTxt by ID
func (s *recordRpzTxtService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzTxtService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.RPZAPI.RecordRpzTxtAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves RecordRpzTxt objects based on filter options
func (s *recordRpzTxtService) List(ctx context.Context, opts *core.ListOptions) ([]*rpz.RecordRpzTxt, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzTxtService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*rpz.RecordRpzTxt, *http.Response, string, error) {
	req := s.niosClient.RPZAPI.RecordRpzTxtAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.RecordRpzTxtFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListRecordRpzTxtResponseObject.GetResult()
	items := make([]*rpz.RecordRpzTxt, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSRecordRpzTxtToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListRecordRpzTxtResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSRecordRpzTxtToResponse(r *niosrpz.RecordRpzTxt) *rpz.RecordRpzTxt {
	resp := &rpz.RecordRpzTxt{
		Id: r.Ref,
	}
	resp.NIOS = &rpz.NIOSRecordRpzTxtExt{
		Comment: r.Comment,
		Disable: r.Disable,
		Name:    r.Name,
		RpZone:  r.RpZone,
		Text:    r.Text,
		Ttl:     r.Ttl,
		UseTtl:  r.UseTtl,
		View:    r.View,
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
