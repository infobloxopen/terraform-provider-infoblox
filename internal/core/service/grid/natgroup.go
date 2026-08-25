package grid

import (
	"context"
	"fmt"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niosgrid "github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/grid"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/grid"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
)

type NatgroupService interface {
	Create(ctx context.Context, obj *grid.Natgroup, opts *core.Options) (*grid.Natgroup, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*grid.Natgroup, *http.Response, error)
	Update(ctx context.Context, id string, obj *grid.Natgroup, opts *core.Options) (*grid.Natgroup, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*grid.Natgroup, *http.Response, string, error)
}

type natgroupService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewNatgroupService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) NatgroupService {
	return &natgroupService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new Natgroup and returns the created object
func (s *natgroupService) Create(ctx context.Context, obj *grid.Natgroup, opts *core.Options) (*grid.Natgroup, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *natgroupService) createNIOS(ctx context.Context, obj *grid.Natgroup, opts *core.Options) (*grid.Natgroup, *http.Response, error) {
	payload, err := common.MapTo[niosgrid.Natgroup](obj, mapper.NatgroupNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.GridAPI.NatgroupAPI.
		Create(ctx).
		Natgroup(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateNatgroupResponseAsObject.GetResult()

	return mapNIOSNatgroupToResponse(&result), httpResp, nil
}

// Read retrieves a Natgroup by ID
func (s *natgroupService) Read(ctx context.Context, id string, opts *core.Options) (*grid.Natgroup, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *natgroupService) readNIOS(ctx context.Context, id string, opts *core.Options) (*grid.Natgroup, *http.Response, error) {
	req := s.niosClient.GridAPI.NatgroupAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetNatgroupResponseObjectAsResult.GetResult()

	return mapNIOSNatgroupToResponse(&result), httpResp, nil
}

// Update modifies an existing Natgroup and returns the updated object
func (s *natgroupService) Update(ctx context.Context, id string, obj *grid.Natgroup, opts *core.Options) (*grid.Natgroup, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *natgroupService) updateNIOS(ctx context.Context, id string, obj *grid.Natgroup, opts *core.Options) (*grid.Natgroup, *http.Response, error) {
	payload, err := common.MapTo[niosgrid.Natgroup](obj, mapper.NatgroupNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.GridAPI.NatgroupAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Natgroup(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateNatgroupResponseAsObject.GetResult()

	return mapNIOSNatgroupToResponse(&result), httpResp, nil
}

// Delete removes a Natgroup by ID
func (s *natgroupService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *natgroupService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.GridAPI.NatgroupAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves Natgroup objects based on filter options
func (s *natgroupService) List(ctx context.Context, opts *core.ListOptions) ([]*grid.Natgroup, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *natgroupService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*grid.Natgroup, *http.Response, string, error) {
	req := s.niosClient.GridAPI.NatgroupAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.NatgroupFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListNatgroupResponseObject.GetResult()
	items := make([]*grid.Natgroup, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSNatgroupToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListNatgroupResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSNatgroupToResponse(r *niosgrid.Natgroup) *grid.Natgroup {
	resp := &grid.Natgroup{
		Id: r.Ref,
	}
	resp.NIOS = &grid.NIOSNatgroupExt{
		Comment: r.Comment,
		Name:    r.Name,
	}
	return resp
}
