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

type RecordRpzAIpaddressService interface {
	Create(ctx context.Context, obj *rpz.RecordRpzAIpaddress, opts *core.Options) (*rpz.RecordRpzAIpaddress, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*rpz.RecordRpzAIpaddress, *http.Response, error)
	Update(ctx context.Context, id string, obj *rpz.RecordRpzAIpaddress, opts *core.Options) (*rpz.RecordRpzAIpaddress, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*rpz.RecordRpzAIpaddress, *http.Response, string, error)
}

type recordRpzAIpaddressService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewRecordRpzAIpaddressService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) RecordRpzAIpaddressService {
	return &recordRpzAIpaddressService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new RecordRpzAIpaddress and returns the created object
func (s *recordRpzAIpaddressService) Create(ctx context.Context, obj *rpz.RecordRpzAIpaddress, opts *core.Options) (*rpz.RecordRpzAIpaddress, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzAIpaddressService) createNIOS(ctx context.Context, obj *rpz.RecordRpzAIpaddress, opts *core.Options) (*rpz.RecordRpzAIpaddress, *http.Response, error) {
	payload, err := common.MapTo[niosrpz.RecordRpzAIpaddress](obj, mapper.RecordRpzAIpaddressNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.RPZAPI.RecordRpzAIpaddressAPI.
		Create(ctx).
		RecordRpzAIpaddress(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateRecordRpzAIpaddressResponseAsObject.GetResult()

	return mapNIOSRecordRpzAIpaddressToResponse(&result), httpResp, nil
}

// Read retrieves a RecordRpzAIpaddress by ID
func (s *recordRpzAIpaddressService) Read(ctx context.Context, id string, opts *core.Options) (*rpz.RecordRpzAIpaddress, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzAIpaddressService) readNIOS(ctx context.Context, id string, opts *core.Options) (*rpz.RecordRpzAIpaddress, *http.Response, error) {
	req := s.niosClient.RPZAPI.RecordRpzAIpaddressAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetRecordRpzAIpaddressResponseObjectAsResult.GetResult()

	return mapNIOSRecordRpzAIpaddressToResponse(&result), httpResp, nil
}

// Update modifies an existing RecordRpzAIpaddress and returns the updated object
func (s *recordRpzAIpaddressService) Update(ctx context.Context, id string, obj *rpz.RecordRpzAIpaddress, opts *core.Options) (*rpz.RecordRpzAIpaddress, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzAIpaddressService) updateNIOS(ctx context.Context, id string, obj *rpz.RecordRpzAIpaddress, opts *core.Options) (*rpz.RecordRpzAIpaddress, *http.Response, error) {
	payload, err := common.MapTo[niosrpz.RecordRpzAIpaddress](obj, mapper.RecordRpzAIpaddressNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.RPZAPI.RecordRpzAIpaddressAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		RecordRpzAIpaddress(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateRecordRpzAIpaddressResponseAsObject.GetResult()

	return mapNIOSRecordRpzAIpaddressToResponse(&result), httpResp, nil
}

// Delete removes a RecordRpzAIpaddress by ID
func (s *recordRpzAIpaddressService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzAIpaddressService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.RPZAPI.RecordRpzAIpaddressAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves RecordRpzAIpaddress objects based on filter options
func (s *recordRpzAIpaddressService) List(ctx context.Context, opts *core.ListOptions) ([]*rpz.RecordRpzAIpaddress, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzAIpaddressService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*rpz.RecordRpzAIpaddress, *http.Response, string, error) {
	req := s.niosClient.RPZAPI.RecordRpzAIpaddressAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.RecordRpzAIpaddressFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListRecordRpzAIpaddressResponseObject.GetResult()
	items := make([]*rpz.RecordRpzAIpaddress, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSRecordRpzAIpaddressToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListRecordRpzAIpaddressResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSRecordRpzAIpaddressToResponse(r *niosrpz.RecordRpzAIpaddress) *rpz.RecordRpzAIpaddress {
	resp := &rpz.RecordRpzAIpaddress{
		Id: r.Ref,
	}
	resp.NIOS = &rpz.NIOSRecordRpzAIpaddressExt{
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
