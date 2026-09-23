package anycast

import (
	"context"
	"fmt"
	"maps"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/anycast"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/anycast"
	uddianycast "github.com/infobloxopen/universal-ddi-go-client/anycast"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
)

type AnycastConfigService interface {
	Create(ctx context.Context, obj *anycast.AnycastConfig, opts *core.Options) (*anycast.AnycastConfig, *http.Response, error)
	Read(ctx context.Context, id int64, opts *core.Options) (*anycast.AnycastConfig, *http.Response, error)
	Update(ctx context.Context, id int64, obj *anycast.AnycastConfig, opts *core.Options) (*anycast.AnycastConfig, *http.Response, error)
	Delete(ctx context.Context, id int64) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*anycast.AnycastConfig, *http.Response, string, error)
}

type anycastConfigService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewAnycastConfigService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) AnycastConfigService {
	return &anycastConfigService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new AnycastConfig and returns the created object
func (s *anycastConfigService) Create(ctx context.Context, obj *anycast.AnycastConfig, opts *core.Options) (*anycast.AnycastConfig, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *anycastConfigService) createUDDI(ctx context.Context, obj *anycast.AnycastConfig, opts *core.Options) (*anycast.AnycastConfig, *http.Response, error) {
	payload, err := common.MapTo[uddianycast.AnycastConfig](obj, mapper.AnycastConfigUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.AnycastAPI.OnPremAnycastManagerAPI.
		CreateAnycastConfig(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResults()

	return mapUDDIAnycastConfigToResponse(&result), httpResp, nil
}

// Read retrieves a AnycastConfig by ID
func (s *anycastConfigService) Read(ctx context.Context, id int64, opts *core.Options) (*anycast.AnycastConfig, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *anycastConfigService) readUDDI(ctx context.Context, id int64, opts *core.Options) (*anycast.AnycastConfig, *http.Response, error) {
	req := s.uddiClient.AnycastAPI.OnPremAnycastManagerAPI.
		GetAnycastConfig(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResults()

	return mapUDDIAnycastConfigToResponse(&result), httpResp, nil
}

// Update modifies an existing AnycastConfig and returns the updated object
func (s *anycastConfigService) Update(ctx context.Context, id int64, obj *anycast.AnycastConfig, opts *core.Options) (*anycast.AnycastConfig, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *anycastConfigService) updateUDDI(ctx context.Context, id int64, obj *anycast.AnycastConfig, opts *core.Options) (*anycast.AnycastConfig, *http.Response, error) {
	payload, err := common.MapTo[uddianycast.AnycastConfig](obj, mapper.AnycastConfigUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.AnycastAPI.OnPremAnycastManagerAPI.
		UpdateAnycastConfig(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResults()

	return mapUDDIAnycastConfigToResponse(&result), httpResp, nil
}

// Delete removes a AnycastConfig by ID
func (s *anycastConfigService) Delete(ctx context.Context, id int64) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *anycastConfigService) deleteUDDI(ctx context.Context, id int64) (*http.Response, error) {
	// This endpoint declares a delete response body, so Execute returns it too.
	_, httpResp, err := s.uddiClient.AnycastAPI.OnPremAnycastManagerAPI.
		DeleteAnycastConfig(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves AnycastConfig objects based on filter options
func (s *anycastConfigService) List(ctx context.Context, opts *core.ListOptions) ([]*anycast.AnycastConfig, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *anycastConfigService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*anycast.AnycastConfig, *http.Response, string, error) {
	req := s.uddiClient.AnycastAPI.OnPremAnycastManagerAPI.GetAnycastConfigList(ctx)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.AnycastConfigFilterFieldMap[core.BackendUDDI])
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

	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, "", err
	}

	results := resp.GetResults()
	items := make([]*anycast.AnycastConfig, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIAnycastConfigToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDIAnycastConfigToResponse(r *uddianycast.AnycastConfig) *anycast.AnycastConfig {
	resp := &anycast.AnycastConfig{
		Id: r.Id,
	}
	resp.UDDI = &anycast.UDDIAnycastConfigExt{
		AccountId:          r.AccountId,
		AnycastIpAddress:   r.AnycastIpAddress,
		AnycastIpv6Address: r.AnycastIpv6Address,
		CreatedAt:          r.CreatedAt,
		Description:        r.Description,
		Fields:             r.Fields,
		IsConfigured:       r.IsConfigured,
		Name:               r.Name,
		OnpremHosts:        r.OnpremHosts,
		RuntimeStatus:      r.RuntimeStatus,
		Service:            r.Service,
		UpdatedAt:          r.UpdatedAt,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
