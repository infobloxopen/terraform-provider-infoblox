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

type ApplicationFilterService interface {
	Create(ctx context.Context, obj *fw.ApplicationFilter, opts *core.Options) (*fw.ApplicationFilter, *http.Response, error)
	Read(ctx context.Context, id int32, opts *core.Options) (*fw.ApplicationFilter, *http.Response, error)
	Update(ctx context.Context, id int32, obj *fw.ApplicationFilter, opts *core.Options) (*fw.ApplicationFilter, *http.Response, error)
	Delete(ctx context.Context, id int32) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*fw.ApplicationFilter, *http.Response, string, error)
}

type applicationFilterService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewApplicationFilterService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) ApplicationFilterService {
	return &applicationFilterService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new ApplicationFilter and returns the created object
func (s *applicationFilterService) Create(ctx context.Context, obj *fw.ApplicationFilter, opts *core.Options) (*fw.ApplicationFilter, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *applicationFilterService) createUDDI(ctx context.Context, obj *fw.ApplicationFilter, opts *core.Options) (*fw.ApplicationFilter, *http.Response, error) {
	payload, err := common.MapTo[uddifw.ApplicationFilter](obj, mapper.ApplicationFilterUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.FWAPI.ApplicationFiltersAPI.
		CreateApplicationFilter(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResults()

	return mapUDDIApplicationFilterToResponse(&result), httpResp, nil
}

// Read retrieves a ApplicationFilter by ID
func (s *applicationFilterService) Read(ctx context.Context, id int32, opts *core.Options) (*fw.ApplicationFilter, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *applicationFilterService) readUDDI(ctx context.Context, id int32, opts *core.Options) (*fw.ApplicationFilter, *http.Response, error) {
	req := s.uddiClient.FWAPI.ApplicationFiltersAPI.
		ReadApplicationFilter(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResults()

	return mapUDDIApplicationFilterToResponse(&result), httpResp, nil
}

// Update modifies an existing ApplicationFilter and returns the updated object
func (s *applicationFilterService) Update(ctx context.Context, id int32, obj *fw.ApplicationFilter, opts *core.Options) (*fw.ApplicationFilter, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *applicationFilterService) updateUDDI(ctx context.Context, id int32, obj *fw.ApplicationFilter, opts *core.Options) (*fw.ApplicationFilter, *http.Response, error) {
	payload, err := common.MapTo[uddifw.ApplicationFilter](obj, mapper.ApplicationFilterUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.FWAPI.ApplicationFiltersAPI.
		UpdateApplicationFilter(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResults()

	return mapUDDIApplicationFilterToResponse(&result), httpResp, nil
}

// Delete removes a ApplicationFilter by ID
func (s *applicationFilterService) Delete(ctx context.Context, id int32) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *applicationFilterService) deleteUDDI(ctx context.Context, id int32) (*http.Response, error) {
	httpResp, err := s.uddiClient.FWAPI.ApplicationFiltersAPI.
		DeleteSingleApplicationFilters(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves ApplicationFilter objects based on filter options
func (s *applicationFilterService) List(ctx context.Context, opts *core.ListOptions) ([]*fw.ApplicationFilter, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *applicationFilterService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*fw.ApplicationFilter, *http.Response, string, error) {
	req := s.uddiClient.FWAPI.ApplicationFiltersAPI.ListApplicationFilters(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.ApplicationFilterFilterFieldMap[core.BackendUDDI])
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
	items := make([]*fw.ApplicationFilter, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIApplicationFilterToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDIApplicationFilterToResponse(r *uddifw.ApplicationFilter) *fw.ApplicationFilter {
	resp := &fw.ApplicationFilter{
		Id: r.Id,
	}
	resp.UDDI = &fw.UDDIApplicationFilterExt{
		Criteria:    r.Criteria,
		Description: r.Description,
		Name:        r.Name,
		Readonly:    r.Readonly,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
