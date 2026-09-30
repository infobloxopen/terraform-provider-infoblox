package fw

import (
	"context"
	"fmt"
	"maps"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/fw"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/fw"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	uddifw "github.com/infobloxopen/universal-ddi-go-client/fw"
)

type CategoryFilterService interface {
	Create(ctx context.Context, obj *fw.CategoryFilter, opts *core.Options) (*fw.CategoryFilter, *http.Response, error)
	Read(ctx context.Context, id int32, opts *core.Options) (*fw.CategoryFilter, *http.Response, error)
	Update(ctx context.Context, id int32, obj *fw.CategoryFilter, opts *core.Options) (*fw.CategoryFilter, *http.Response, error)
	Delete(ctx context.Context, id int32) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*fw.CategoryFilter, *http.Response, string, error)
}

type categoryFilterService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewCategoryFilterService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) CategoryFilterService {
	return &categoryFilterService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new CategoryFilter and returns the created object
func (s *categoryFilterService) Create(ctx context.Context, obj *fw.CategoryFilter, opts *core.Options) (*fw.CategoryFilter, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *categoryFilterService) createUDDI(ctx context.Context, obj *fw.CategoryFilter, opts *core.Options) (*fw.CategoryFilter, *http.Response, error) {
	payload, err := common.MapTo[uddifw.CategoryFilter](obj, mapper.CategoryFilterUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.FWAPI.CategoryFiltersAPI.
		CreateCategoryFilter(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResults()

	return mapUDDICategoryFilterToResponse(&result), httpResp, nil
}

// Read retrieves a CategoryFilter by ID
func (s *categoryFilterService) Read(ctx context.Context, id int32, opts *core.Options) (*fw.CategoryFilter, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *categoryFilterService) readUDDI(ctx context.Context, id int32, opts *core.Options) (*fw.CategoryFilter, *http.Response, error) {
	req := s.uddiClient.FWAPI.CategoryFiltersAPI.
		ReadCategoryFilter(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResults()

	return mapUDDICategoryFilterToResponse(&result), httpResp, nil
}

// Update modifies an existing CategoryFilter and returns the updated object
func (s *categoryFilterService) Update(ctx context.Context, id int32, obj *fw.CategoryFilter, opts *core.Options) (*fw.CategoryFilter, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *categoryFilterService) updateUDDI(ctx context.Context, id int32, obj *fw.CategoryFilter, opts *core.Options) (*fw.CategoryFilter, *http.Response, error) {
	payload, err := common.MapTo[uddifw.CategoryFilter](obj, mapper.CategoryFilterUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.FWAPI.CategoryFiltersAPI.
		UpdateCategoryFilter(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResults()

	return mapUDDICategoryFilterToResponse(&result), httpResp, nil
}

// Delete removes a CategoryFilter by ID
func (s *categoryFilterService) Delete(ctx context.Context, id int32) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *categoryFilterService) deleteUDDI(ctx context.Context, id int32) (*http.Response, error) {
	httpResp, err := s.uddiClient.FWAPI.CategoryFiltersAPI.
		DeleteSingleCategoryFilters(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves CategoryFilter objects based on filter options
func (s *categoryFilterService) List(ctx context.Context, opts *core.ListOptions) ([]*fw.CategoryFilter, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *categoryFilterService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*fw.CategoryFilter, *http.Response, string, error) {
	req := s.uddiClient.FWAPI.CategoryFiltersAPI.ListCategoryFilters(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.CategoryFilterFilterFieldMap[core.BackendUDDI])
		for k, v := range translatedFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		if len(filters) > 0 {
			req = req.Filter(core.JoinFilters(filters))
		}

		if len(opts.TagFilter) > 0 {
			var tfilters []string
			for k, v := range opts.TagFilter {
				tfilters = append(tfilters, "'"+k+"'=='"+v+"'")
			}
			req = req.Tfilter(core.JoinFilters(tfilters))
		}

		if opts.Offset > 0 {
			req = req.Offset(opts.Offset)
		}

		if opts.Limit > 0 {
			req = req.Limit(opts.Limit)
		}
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, "", err
	}

	results := resp.GetResults()
	items := make([]*fw.CategoryFilter, 0, len(results))
	for i := range results {
		items = append(items, mapUDDICategoryFilterToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDICategoryFilterToResponse(r *uddifw.CategoryFilter) *fw.CategoryFilter {
	resp := &fw.CategoryFilter{
		Id: r.Id,
	}
	resp.UDDI = &fw.UDDICategoryFilterExt{
		Categories:  r.Categories,
		Description: r.Description,
		Name:        r.Name,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
