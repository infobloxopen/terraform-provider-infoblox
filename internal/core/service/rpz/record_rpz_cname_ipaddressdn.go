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

type RecordRpzCnameIpaddressdnService interface {
	Create(ctx context.Context, obj *rpz.RecordRpzCnameIpaddressdn, opts *core.Options) (*rpz.RecordRpzCnameIpaddressdn, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*rpz.RecordRpzCnameIpaddressdn, *http.Response, error)
	Update(ctx context.Context, id string, obj *rpz.RecordRpzCnameIpaddressdn, opts *core.Options) (*rpz.RecordRpzCnameIpaddressdn, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*rpz.RecordRpzCnameIpaddressdn, *http.Response, string, error)
}

type recordRpzCnameIpaddressdnService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewRecordRpzCnameIpaddressdnService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) RecordRpzCnameIpaddressdnService {
	return &recordRpzCnameIpaddressdnService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new RecordRpzCnameIpaddressdn and returns the created object
func (s *recordRpzCnameIpaddressdnService) Create(ctx context.Context, obj *rpz.RecordRpzCnameIpaddressdn, opts *core.Options) (*rpz.RecordRpzCnameIpaddressdn, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzCnameIpaddressdnService) createNIOS(ctx context.Context, obj *rpz.RecordRpzCnameIpaddressdn, opts *core.Options) (*rpz.RecordRpzCnameIpaddressdn, *http.Response, error) {
	payload, err := common.MapTo[niosrpz.RecordRpzCnameIpaddressdn](obj, mapper.RecordRpzCnameIpaddressdnNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.RPZAPI.RecordRpzCnameIpaddressdnAPI.
		Create(ctx).
		RecordRpzCnameIpaddressdn(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateRecordRpzCnameIpaddressdnResponseAsObject.GetResult()

	return mapNIOSRecordRpzCnameIpaddressdnToResponse(&result), httpResp, nil
}

// Read retrieves a RecordRpzCnameIpaddressdn by ID
func (s *recordRpzCnameIpaddressdnService) Read(ctx context.Context, id string, opts *core.Options) (*rpz.RecordRpzCnameIpaddressdn, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzCnameIpaddressdnService) readNIOS(ctx context.Context, id string, opts *core.Options) (*rpz.RecordRpzCnameIpaddressdn, *http.Response, error) {
	req := s.niosClient.RPZAPI.RecordRpzCnameIpaddressdnAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetRecordRpzCnameIpaddressdnResponseObjectAsResult.GetResult()

	return mapNIOSRecordRpzCnameIpaddressdnToResponse(&result), httpResp, nil
}

// Update modifies an existing RecordRpzCnameIpaddressdn and returns the updated object
func (s *recordRpzCnameIpaddressdnService) Update(ctx context.Context, id string, obj *rpz.RecordRpzCnameIpaddressdn, opts *core.Options) (*rpz.RecordRpzCnameIpaddressdn, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzCnameIpaddressdnService) updateNIOS(ctx context.Context, id string, obj *rpz.RecordRpzCnameIpaddressdn, opts *core.Options) (*rpz.RecordRpzCnameIpaddressdn, *http.Response, error) {
	payload, err := common.MapTo[niosrpz.RecordRpzCnameIpaddressdn](obj, mapper.RecordRpzCnameIpaddressdnNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.RPZAPI.RecordRpzCnameIpaddressdnAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		RecordRpzCnameIpaddressdn(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateRecordRpzCnameIpaddressdnResponseAsObject.GetResult()

	return mapNIOSRecordRpzCnameIpaddressdnToResponse(&result), httpResp, nil
}

// Delete removes a RecordRpzCnameIpaddressdn by ID
func (s *recordRpzCnameIpaddressdnService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzCnameIpaddressdnService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.RPZAPI.RecordRpzCnameIpaddressdnAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves RecordRpzCnameIpaddressdn objects based on filter options
func (s *recordRpzCnameIpaddressdnService) List(ctx context.Context, opts *core.ListOptions) ([]*rpz.RecordRpzCnameIpaddressdn, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzCnameIpaddressdnService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*rpz.RecordRpzCnameIpaddressdn, *http.Response, string, error) {
	req := s.niosClient.RPZAPI.RecordRpzCnameIpaddressdnAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.RecordRpzCnameIpaddressdnFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListRecordRpzCnameIpaddressdnResponseObject.GetResult()
	items := make([]*rpz.RecordRpzCnameIpaddressdn, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSRecordRpzCnameIpaddressdnToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListRecordRpzCnameIpaddressdnResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSRecordRpzCnameIpaddressdnToResponse(r *niosrpz.RecordRpzCnameIpaddressdn) *rpz.RecordRpzCnameIpaddressdn {
	resp := &rpz.RecordRpzCnameIpaddressdn{
		Id: r.Ref,
	}
	resp.NIOS = &rpz.NIOSRecordRpzCnameIpaddressdnExt{
		Canonical: r.Canonical,
		Comment:   r.Comment,
		Disable:   r.Disable,
		Name:      r.Name,
		RpZone:    r.RpZone,
		Ttl:       r.Ttl,
		UseTtl:    r.UseTtl,
		View:      r.View,
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
