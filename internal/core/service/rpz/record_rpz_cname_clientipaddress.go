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

type RecordRpzCnameClientipaddressService interface {
	Create(ctx context.Context, obj *rpz.RecordRpzCnameClientipaddress, opts *core.Options) (*rpz.RecordRpzCnameClientipaddress, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*rpz.RecordRpzCnameClientipaddress, *http.Response, error)
	Update(ctx context.Context, id string, obj *rpz.RecordRpzCnameClientipaddress, opts *core.Options) (*rpz.RecordRpzCnameClientipaddress, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*rpz.RecordRpzCnameClientipaddress, *http.Response, string, error)
}

type recordRpzCnameClientipaddressService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewRecordRpzCnameClientipaddressService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) RecordRpzCnameClientipaddressService {
	return &recordRpzCnameClientipaddressService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new RecordRpzCnameClientipaddress and returns the created object
func (s *recordRpzCnameClientipaddressService) Create(ctx context.Context, obj *rpz.RecordRpzCnameClientipaddress, opts *core.Options) (*rpz.RecordRpzCnameClientipaddress, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzCnameClientipaddressService) createNIOS(ctx context.Context, obj *rpz.RecordRpzCnameClientipaddress, opts *core.Options) (*rpz.RecordRpzCnameClientipaddress, *http.Response, error) {
	payload, err := common.MapTo[niosrpz.RecordRpzCnameClientipaddress](obj, mapper.RecordRpzCnameClientipaddressNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.RPZAPI.RecordRpzCnameClientipaddressAPI.
		Create(ctx).
		RecordRpzCnameClientipaddress(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateRecordRpzCnameClientipaddressResponseAsObject.GetResult()

	return mapNIOSRecordRpzCnameClientipaddressToResponse(&result), httpResp, nil
}

// Read retrieves a RecordRpzCnameClientipaddress by ID
func (s *recordRpzCnameClientipaddressService) Read(ctx context.Context, id string, opts *core.Options) (*rpz.RecordRpzCnameClientipaddress, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzCnameClientipaddressService) readNIOS(ctx context.Context, id string, opts *core.Options) (*rpz.RecordRpzCnameClientipaddress, *http.Response, error) {
	req := s.niosClient.RPZAPI.RecordRpzCnameClientipaddressAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetRecordRpzCnameClientipaddressResponseObjectAsResult.GetResult()

	return mapNIOSRecordRpzCnameClientipaddressToResponse(&result), httpResp, nil
}

// Update modifies an existing RecordRpzCnameClientipaddress and returns the updated object
func (s *recordRpzCnameClientipaddressService) Update(ctx context.Context, id string, obj *rpz.RecordRpzCnameClientipaddress, opts *core.Options) (*rpz.RecordRpzCnameClientipaddress, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzCnameClientipaddressService) updateNIOS(ctx context.Context, id string, obj *rpz.RecordRpzCnameClientipaddress, opts *core.Options) (*rpz.RecordRpzCnameClientipaddress, *http.Response, error) {
	payload, err := common.MapTo[niosrpz.RecordRpzCnameClientipaddress](obj, mapper.RecordRpzCnameClientipaddressNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.RPZAPI.RecordRpzCnameClientipaddressAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		RecordRpzCnameClientipaddress(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateRecordRpzCnameClientipaddressResponseAsObject.GetResult()

	return mapNIOSRecordRpzCnameClientipaddressToResponse(&result), httpResp, nil
}

// Delete removes a RecordRpzCnameClientipaddress by ID
func (s *recordRpzCnameClientipaddressService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzCnameClientipaddressService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.RPZAPI.RecordRpzCnameClientipaddressAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves RecordRpzCnameClientipaddress objects based on filter options
func (s *recordRpzCnameClientipaddressService) List(ctx context.Context, opts *core.ListOptions) ([]*rpz.RecordRpzCnameClientipaddress, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzCnameClientipaddressService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*rpz.RecordRpzCnameClientipaddress, *http.Response, string, error) {
	req := s.niosClient.RPZAPI.RecordRpzCnameClientipaddressAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.RecordRpzCnameClientipaddressFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListRecordRpzCnameClientipaddressResponseObject.GetResult()
	items := make([]*rpz.RecordRpzCnameClientipaddress, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSRecordRpzCnameClientipaddressToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListRecordRpzCnameClientipaddressResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSRecordRpzCnameClientipaddressToResponse(r *niosrpz.RecordRpzCnameClientipaddress) *rpz.RecordRpzCnameClientipaddress {
	resp := &rpz.RecordRpzCnameClientipaddress{
		Id: r.Ref,
	}
	resp.NIOS = &rpz.NIOSRecordRpzCnameClientipaddressExt{
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
