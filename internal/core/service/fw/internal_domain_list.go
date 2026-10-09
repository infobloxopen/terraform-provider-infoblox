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

type InternalDomainListService interface {
	Create(ctx context.Context, obj *fw.InternalDomainList, opts *core.Options) (*fw.InternalDomainList, *http.Response, error)
	Read(ctx context.Context, id int32, opts *core.Options) (*fw.InternalDomainList, *http.Response, error)
	Update(ctx context.Context, id int32, obj *fw.InternalDomainList, opts *core.Options) (*fw.InternalDomainList, *http.Response, error)
	Delete(ctx context.Context, id int32) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*fw.InternalDomainList, *http.Response, string, error)
}

type internalDomainListService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewInternalDomainListService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) InternalDomainListService {
	return &internalDomainListService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new InternalDomainList and returns the created object
func (s *internalDomainListService) Create(ctx context.Context, obj *fw.InternalDomainList, opts *core.Options) (*fw.InternalDomainList, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *internalDomainListService) createUDDI(ctx context.Context, obj *fw.InternalDomainList, opts *core.Options) (*fw.InternalDomainList, *http.Response, error) {
	payload, err := common.MapTo[uddifw.InternalDomains](obj, mapper.InternalDomainListUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.FWAPI.InternalDomainListsAPI.
		CreateInternalDomains(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResults()

	return mapUDDIInternalDomainListToResponse(&result), httpResp, nil
}

// Read retrieves a InternalDomainList by ID
func (s *internalDomainListService) Read(ctx context.Context, id int32, opts *core.Options) (*fw.InternalDomainList, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *internalDomainListService) readUDDI(ctx context.Context, id int32, opts *core.Options) (*fw.InternalDomainList, *http.Response, error) {
	req := s.uddiClient.FWAPI.InternalDomainListsAPI.
		ReadInternalDomains(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResults()

	return mapUDDIInternalDomainListToResponse(&result), httpResp, nil
}

// Update modifies an existing InternalDomainList and returns the updated object
func (s *internalDomainListService) Update(ctx context.Context, id int32, obj *fw.InternalDomainList, opts *core.Options) (*fw.InternalDomainList, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *internalDomainListService) updateUDDI(ctx context.Context, id int32, obj *fw.InternalDomainList, opts *core.Options) (*fw.InternalDomainList, *http.Response, error) {
	payload, err := common.MapTo[uddifw.InternalDomains](obj, mapper.InternalDomainListUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.FWAPI.InternalDomainListsAPI.
		UpdateInternalDomains(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResults()

	return mapUDDIInternalDomainListToResponse(&result), httpResp, nil
}

// Delete removes a InternalDomainList by ID
func (s *internalDomainListService) Delete(ctx context.Context, id int32) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *internalDomainListService) deleteUDDI(ctx context.Context, id int32) (*http.Response, error) {
	httpResp, err := s.uddiClient.FWAPI.InternalDomainListsAPI.
		DeleteSingleInternalDomains(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves InternalDomainList objects based on filter options
func (s *internalDomainListService) List(ctx context.Context, opts *core.ListOptions) ([]*fw.InternalDomainList, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *internalDomainListService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*fw.InternalDomainList, *http.Response, string, error) {
	req := s.uddiClient.FWAPI.InternalDomainListsAPI.ListInternalDomains(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.InternalDomainListFilterFieldMap[core.BackendUDDI])
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
	items := make([]*fw.InternalDomainList, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIInternalDomainListToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDIInternalDomainListToResponse(r *uddifw.InternalDomains) *fw.InternalDomainList {
	resp := &fw.InternalDomainList{
		Id: r.Id,
	}
	resp.UDDI = &fw.UDDIInternalDomainListExt{
		Description:     r.Description,
		InternalDomains: r.InternalDomains,
		IsDefault:       r.IsDefault,
		Name:            r.Name,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
