package clouddiscovery

import (
	"context"
	"fmt"
	"maps"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/clouddiscovery"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/clouddiscovery"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	uddiclouddiscovery "github.com/infobloxopen/universal-ddi-go-client/clouddiscovery"
)

type CloudDiscoveryProviderService interface {
	Create(ctx context.Context, obj *clouddiscovery.CloudDiscoveryProvider, opts *core.Options) (*clouddiscovery.CloudDiscoveryProvider, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*clouddiscovery.CloudDiscoveryProvider, *http.Response, error)
	Update(ctx context.Context, id string, obj *clouddiscovery.CloudDiscoveryProvider, opts *core.Options) (*clouddiscovery.CloudDiscoveryProvider, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*clouddiscovery.CloudDiscoveryProvider, *http.Response, string, error)
}

type cloudDiscoveryProviderService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewCloudDiscoveryProviderService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) CloudDiscoveryProviderService {
	return &cloudDiscoveryProviderService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new CloudDiscoveryProvider and returns the created object
func (s *cloudDiscoveryProviderService) Create(ctx context.Context, obj *clouddiscovery.CloudDiscoveryProvider, opts *core.Options) (*clouddiscovery.CloudDiscoveryProvider, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *cloudDiscoveryProviderService) createUDDI(ctx context.Context, obj *clouddiscovery.CloudDiscoveryProvider, opts *core.Options) (*clouddiscovery.CloudDiscoveryProvider, *http.Response, error) {
	payload, err := common.MapTo[uddiclouddiscovery.DiscoveryConfig](obj, mapper.CloudDiscoveryProviderUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DiscoveryConfigurationAPIV2.ProvidersAPI.
		Create(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDICloudDiscoveryProviderToResponse(&result), httpResp, nil
}

// Read retrieves a CloudDiscoveryProvider by ID
func (s *cloudDiscoveryProviderService) Read(ctx context.Context, id string, opts *core.Options) (*clouddiscovery.CloudDiscoveryProvider, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *cloudDiscoveryProviderService) readUDDI(ctx context.Context, id string, opts *core.Options) (*clouddiscovery.CloudDiscoveryProvider, *http.Response, error) {
	req := s.uddiClient.DiscoveryConfigurationAPIV2.ProvidersAPI.
		Read(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDICloudDiscoveryProviderToResponse(&result), httpResp, nil
}

// Update modifies an existing CloudDiscoveryProvider and returns the updated object
func (s *cloudDiscoveryProviderService) Update(ctx context.Context, id string, obj *clouddiscovery.CloudDiscoveryProvider, opts *core.Options) (*clouddiscovery.CloudDiscoveryProvider, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *cloudDiscoveryProviderService) updateUDDI(ctx context.Context, id string, obj *clouddiscovery.CloudDiscoveryProvider, opts *core.Options) (*clouddiscovery.CloudDiscoveryProvider, *http.Response, error) {
	payload, err := common.MapTo[uddiclouddiscovery.DiscoveryConfig](obj, mapper.CloudDiscoveryProviderUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DiscoveryConfigurationAPIV2.ProvidersAPI.
		Update(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDICloudDiscoveryProviderToResponse(&result), httpResp, nil
}

// Delete removes a CloudDiscoveryProvider by ID
func (s *cloudDiscoveryProviderService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *cloudDiscoveryProviderService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.DiscoveryConfigurationAPIV2.ProvidersAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves CloudDiscoveryProvider objects based on filter options
func (s *cloudDiscoveryProviderService) List(ctx context.Context, opts *core.ListOptions) ([]*clouddiscovery.CloudDiscoveryProvider, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *cloudDiscoveryProviderService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*clouddiscovery.CloudDiscoveryProvider, *http.Response, string, error) {
	req := s.uddiClient.DiscoveryConfigurationAPIV2.ProvidersAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.CloudDiscoveryProviderFilterFieldMap[core.BackendUDDI])
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
	items := make([]*clouddiscovery.CloudDiscoveryProvider, 0, len(results))
	for i := range results {
		items = append(items, mapUDDICloudDiscoveryProviderToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDICloudDiscoveryProviderToResponse(r *uddiclouddiscovery.DiscoveryConfig) *clouddiscovery.CloudDiscoveryProvider {
	resp := &clouddiscovery.CloudDiscoveryProvider{
		Id: r.Id,
	}
	resp.UDDI = &clouddiscovery.UDDICloudDiscoveryProviderExt{
		AccountPreference:       r.AccountPreference,
		AdditionalConfig:        r.AdditionalConfig,
		CredentialPreference:    r.CredentialPreference,
		Description:             r.Description,
		DesiredState:            r.DesiredState,
		DestinationTypesEnabled: r.DestinationTypesEnabled,
		Destinations:            r.Destinations,
		IsDisabled:              r.IsDisabled,
		LabsProvider:            r.LabsProvider,
		Name:                    r.Name,
		ProviderType:            r.ProviderType,
		SourceConfigs:           r.SourceConfigs,
		SyncInterval:            r.SyncInterval,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
