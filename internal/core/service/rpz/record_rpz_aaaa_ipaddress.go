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

type RecordRpzAaaaIpaddressService interface {
	Create(ctx context.Context, obj *rpz.RecordRpzAaaaIpaddress, opts *core.Options) (*rpz.RecordRpzAaaaIpaddress, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*rpz.RecordRpzAaaaIpaddress, *http.Response, error)
	Update(ctx context.Context, id string, obj *rpz.RecordRpzAaaaIpaddress, opts *core.Options) (*rpz.RecordRpzAaaaIpaddress, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*rpz.RecordRpzAaaaIpaddress, *http.Response, string, error)
}

type recordRpzAaaaIpaddressService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewRecordRpzAaaaIpaddressService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) RecordRpzAaaaIpaddressService {
	return &recordRpzAaaaIpaddressService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new RecordRpzAaaaIpaddress and returns the created object
func (s *recordRpzAaaaIpaddressService) Create(ctx context.Context, obj *rpz.RecordRpzAaaaIpaddress, opts *core.Options) (*rpz.RecordRpzAaaaIpaddress, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzAaaaIpaddressService) createNIOS(ctx context.Context, obj *rpz.RecordRpzAaaaIpaddress, opts *core.Options) (*rpz.RecordRpzAaaaIpaddress, *http.Response, error) {
	payload, err := common.MapTo[niosrpz.RecordRpzAaaaIpaddress](obj, mapper.RecordRpzAaaaIpaddressNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.RPZAPI.RecordRpzAaaaIpaddressAPI.
		Create(ctx).
		RecordRpzAaaaIpaddress(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateRecordRpzAaaaIpaddressResponseAsObject.GetResult()

	return mapNIOSRecordRpzAaaaIpaddressToResponse(&result), httpResp, nil
}

// Read retrieves a RecordRpzAaaaIpaddress by ID
func (s *recordRpzAaaaIpaddressService) Read(ctx context.Context, id string, opts *core.Options) (*rpz.RecordRpzAaaaIpaddress, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzAaaaIpaddressService) readNIOS(ctx context.Context, id string, opts *core.Options) (*rpz.RecordRpzAaaaIpaddress, *http.Response, error) {
	req := s.niosClient.RPZAPI.RecordRpzAaaaIpaddressAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetRecordRpzAaaaIpaddressResponseObjectAsResult.GetResult()

	return mapNIOSRecordRpzAaaaIpaddressToResponse(&result), httpResp, nil
}

// Update modifies an existing RecordRpzAaaaIpaddress and returns the updated object
func (s *recordRpzAaaaIpaddressService) Update(ctx context.Context, id string, obj *rpz.RecordRpzAaaaIpaddress, opts *core.Options) (*rpz.RecordRpzAaaaIpaddress, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzAaaaIpaddressService) updateNIOS(ctx context.Context, id string, obj *rpz.RecordRpzAaaaIpaddress, opts *core.Options) (*rpz.RecordRpzAaaaIpaddress, *http.Response, error) {
	payload, err := common.MapTo[niosrpz.RecordRpzAaaaIpaddress](obj, mapper.RecordRpzAaaaIpaddressNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.RPZAPI.RecordRpzAaaaIpaddressAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		RecordRpzAaaaIpaddress(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateRecordRpzAaaaIpaddressResponseAsObject.GetResult()

	return mapNIOSRecordRpzAaaaIpaddressToResponse(&result), httpResp, nil
}

// Delete removes a RecordRpzAaaaIpaddress by ID
func (s *recordRpzAaaaIpaddressService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzAaaaIpaddressService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.RPZAPI.RecordRpzAaaaIpaddressAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves RecordRpzAaaaIpaddress objects based on filter options
func (s *recordRpzAaaaIpaddressService) List(ctx context.Context, opts *core.ListOptions) ([]*rpz.RecordRpzAaaaIpaddress, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordRpzAaaaIpaddressService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*rpz.RecordRpzAaaaIpaddress, *http.Response, string, error) {
	req := s.niosClient.RPZAPI.RecordRpzAaaaIpaddressAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.RecordRpzAaaaIpaddressFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListRecordRpzAaaaIpaddressResponseObject.GetResult()
	items := make([]*rpz.RecordRpzAaaaIpaddress, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSRecordRpzAaaaIpaddressToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListRecordRpzAaaaIpaddressResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSRecordRpzAaaaIpaddressToResponse(r *niosrpz.RecordRpzAaaaIpaddress) *rpz.RecordRpzAaaaIpaddress {
	resp := &rpz.RecordRpzAaaaIpaddress{
		Id: r.Ref,
	}
	resp.NIOS = &rpz.NIOSRecordRpzAaaaIpaddressExt{
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
