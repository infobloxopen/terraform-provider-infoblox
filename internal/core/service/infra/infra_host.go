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

type InfraHostService interface {
	Create(ctx context.Context, obj *infra.InfraHost, opts *core.Options) (*infra.InfraHost, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*infra.InfraHost, *http.Response, error)
	Update(ctx context.Context, id string, obj *infra.InfraHost, opts *core.Options) (*infra.InfraHost, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*infra.InfraHost, *http.Response, string, error)
}

type infraHostService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewInfraHostService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) InfraHostService {
	return &infraHostService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new InfraHost and returns the created object
func (s *infraHostService) Create(ctx context.Context, obj *infra.InfraHost, opts *core.Options) (*infra.InfraHost, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *infraHostService) createUDDI(ctx context.Context, obj *infra.InfraHost, opts *core.Options) (*infra.InfraHost, *http.Response, error) {
	payload, err := common.MapTo[uddiinframgmt.Host](obj, mapper.InfraHostUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.InfraManagementAPI.HostsAPI.
		Create(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIInfraHostToResponse(&result), httpResp, nil
}

// Read retrieves a InfraHost by ID
func (s *infraHostService) Read(ctx context.Context, id string, opts *core.Options) (*infra.InfraHost, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *infraHostService) readUDDI(ctx context.Context, id string, opts *core.Options) (*infra.InfraHost, *http.Response, error) {
	req := s.uddiClient.InfraManagementAPI.HostsAPI.
		Read(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIInfraHostToResponse(&result), httpResp, nil
}

// Update modifies an existing InfraHost and returns the updated object
func (s *infraHostService) Update(ctx context.Context, id string, obj *infra.InfraHost, opts *core.Options) (*infra.InfraHost, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *infraHostService) updateUDDI(ctx context.Context, id string, obj *infra.InfraHost, opts *core.Options) (*infra.InfraHost, *http.Response, error) {
	payload, err := common.MapTo[uddiinframgmt.Host](obj, mapper.InfraHostUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.InfraManagementAPI.HostsAPI.
		Update(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIInfraHostToResponse(&result), httpResp, nil
}

// Delete removes a InfraHost by ID
func (s *infraHostService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *infraHostService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.InfraManagementAPI.HostsAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves InfraHost objects based on filter options
func (s *infraHostService) List(ctx context.Context, opts *core.ListOptions) ([]*infra.InfraHost, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *infraHostService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*infra.InfraHost, *http.Response, string, error) {
	req := s.uddiClient.InfraManagementAPI.HostsAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.InfraHostFilterFieldMap[core.BackendUDDI])
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
	items := make([]*infra.InfraHost, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIInfraHostToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDIInfraHostToResponse(r *uddiinframgmt.Host) *infra.InfraHost {
	resp := &infra.InfraHost{
		Id: r.Id,
	}
	resp.UDDI = &infra.UDDIInfraHostExt{
		Description:     r.Description,
		DisplayName:     r.DisplayName,
		IpSpace:         r.IpSpace,
		LocationId:      r.LocationId,
		MaintenanceMode: r.MaintenanceMode,
		PoolId:          r.PoolId,
		SerialNumber:    r.SerialNumber,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
