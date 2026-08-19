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

type NamedListService interface {
	Create(ctx context.Context, obj *fw.NamedList, opts *core.Options) (*fw.NamedList, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*fw.NamedList, *http.Response, error)
	Update(ctx context.Context, id string, obj *fw.NamedList, opts *core.Options) (*fw.NamedList, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*fw.NamedList, *http.Response, string, error)
}

type namedListService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewNamedListService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) NamedListService {
	return &namedListService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new NamedList and returns the created object
func (s *namedListService) Create(ctx context.Context, obj *fw.NamedList, opts *core.Options) (*fw.NamedList, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *namedListService) createUDDI(ctx context.Context, obj *fw.NamedList, opts *core.Options) (*fw.NamedList, *http.Response, error) {
	payload, err := common.MapTo[uddifw.NamedList](obj, mapper.NamedListUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.FWAPI.NamedListsAPI.
		Create(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDINamedListToResponse(&result), httpResp, nil
}

// Read retrieves a NamedList by ID
func (s *namedListService) Read(ctx context.Context, id string, opts *core.Options) (*fw.NamedList, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *namedListService) readUDDI(ctx context.Context, id string, opts *core.Options) (*fw.NamedList, *http.Response, error) {
	req := s.uddiClient.FWAPI.NamedListsAPI.
		Read(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDINamedListToResponse(&result), httpResp, nil
}

// Update modifies an existing NamedList and returns the updated object
func (s *namedListService) Update(ctx context.Context, id string, obj *fw.NamedList, opts *core.Options) (*fw.NamedList, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *namedListService) updateUDDI(ctx context.Context, id string, obj *fw.NamedList, opts *core.Options) (*fw.NamedList, *http.Response, error) {
	payload, err := common.MapTo[uddifw.NamedList](obj, mapper.NamedListUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.FWAPI.NamedListAPI.
		Update(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDINamedListToResponse(&result), httpResp, nil
}

// Delete removes a NamedList by ID
func (s *namedListService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *namedListService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.FWAPI.NamedListsAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves NamedList objects based on filter options
func (s *namedListService) List(ctx context.Context, opts *core.ListOptions) ([]*fw.NamedList, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *namedListService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*fw.NamedList, *http.Response, string, error) {
	req := s.uddiClient.FWAPI.NamedListsAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.NamedListFilterFieldMap[core.BackendUDDI])
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
	items := make([]*fw.NamedList, 0, len(results))
	for i := range results {
		items = append(items, mapUDDINamedListToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDINamedListToResponse(r *uddifw.NamedList) *fw.NamedList {
	resp := &fw.NamedList{
		Id: r.Id,
	}
	resp.UDDI = &fw.UDDINamedListExt{
		ConfidenceLevel: r.ConfidenceLevel,
		Description:     r.Description,
		Items:           r.Items,
		ItemsDescribed:  r.ItemsDescribed,
		Name:            r.Name,
		Policies:        r.Policies,
		ThreatLevel:     r.ThreatLevel,
		Type:            r.Type,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
