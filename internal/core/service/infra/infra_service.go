package infra

import (
	"context"
	"fmt"
	"maps"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/infra"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/infra"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	uddiinframgmt "github.com/infobloxopen/universal-ddi-go-client/inframgmt"
)

type InfraServiceService interface {
	Create(ctx context.Context, obj *infra.InfraService, opts *core.Options) (*infra.InfraService, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*infra.InfraService, *http.Response, error)
	Update(ctx context.Context, id string, obj *infra.InfraService, opts *core.Options) (*infra.InfraService, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*infra.InfraService, *http.Response, string, error)
}

type infraServiceService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewInfraServiceService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) InfraServiceService {
	return &infraServiceService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new InfraService and returns the created object
func (s *infraServiceService) Create(ctx context.Context, obj *infra.InfraService, opts *core.Options) (*infra.InfraService, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *infraServiceService) createUDDI(ctx context.Context, obj *infra.InfraService, opts *core.Options) (*infra.InfraService, *http.Response, error) {
	payload, err := common.MapTo[uddiinframgmt.Service](obj, mapper.InfraServiceUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.InfraManagementAPI.ServicesAPI.
		Create(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIInfraServiceToResponse(&result), httpResp, nil
}

// Read retrieves a InfraService by ID
func (s *infraServiceService) Read(ctx context.Context, id string, opts *core.Options) (*infra.InfraService, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *infraServiceService) readUDDI(ctx context.Context, id string, opts *core.Options) (*infra.InfraService, *http.Response, error) {
	req := s.uddiClient.InfraManagementAPI.ServicesAPI.
		Read(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIInfraServiceToResponse(&result), httpResp, nil
}

// Update modifies an existing InfraService and returns the updated object
func (s *infraServiceService) Update(ctx context.Context, id string, obj *infra.InfraService, opts *core.Options) (*infra.InfraService, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *infraServiceService) updateUDDI(ctx context.Context, id string, obj *infra.InfraService, opts *core.Options) (*infra.InfraService, *http.Response, error) {
	payload, err := common.MapTo[uddiinframgmt.Service](obj, mapper.InfraServiceUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.InfraManagementAPI.ServicesAPI.
		Update(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIInfraServiceToResponse(&result), httpResp, nil
}

// Delete removes a InfraService by ID
func (s *infraServiceService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *infraServiceService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.InfraManagementAPI.ServicesAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves InfraService objects based on filter options
func (s *infraServiceService) List(ctx context.Context, opts *core.ListOptions) ([]*infra.InfraService, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *infraServiceService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*infra.InfraService, *http.Response, string, error) {
	req := s.uddiClient.InfraManagementAPI.ServicesAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.InfraServiceFilterFieldMap[core.BackendUDDI])
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
	items := make([]*infra.InfraService, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIInfraServiceToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDIInfraServiceToResponse(r *uddiinframgmt.Service) *infra.InfraService {
	resp := &infra.InfraService{
		Id: r.Id,
	}
	resp.UDDI = &infra.UDDIInfraServiceExt{
		Configs:         r.Configs,
		CreatedAt:       r.CreatedAt,
		Description:     r.Description,
		DesiredState:    r.DesiredState,
		DesiredVersion:  r.DesiredVersion,
		InterfaceLabels: r.InterfaceLabels,
		Name:            r.Name,
		PoolId:          r.PoolId,
		ServiceType:     r.ServiceType,
		UpdatedAt:       r.UpdatedAt,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
