package ipamfederation

import (
	"context"
	"fmt"
	"maps"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/ipamfederation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/ipamfederation"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	uddiipamfederation "github.com/infobloxopen/universal-ddi-go-client/ipamfederation"
)

type FederatedRealmService interface {
	Create(ctx context.Context, obj *ipamfederation.FederatedRealm, opts *core.Options) (*ipamfederation.FederatedRealm, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*ipamfederation.FederatedRealm, *http.Response, error)
	Update(ctx context.Context, id string, obj *ipamfederation.FederatedRealm, opts *core.Options) (*ipamfederation.FederatedRealm, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*ipamfederation.FederatedRealm, *http.Response, string, error)
}

type federatedRealmService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewFederatedRealmService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) FederatedRealmService {
	return &federatedRealmService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new FederatedRealm and returns the created object
func (s *federatedRealmService) Create(ctx context.Context, obj *ipamfederation.FederatedRealm, opts *core.Options) (*ipamfederation.FederatedRealm, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *federatedRealmService) createUDDI(ctx context.Context, obj *ipamfederation.FederatedRealm, opts *core.Options) (*ipamfederation.FederatedRealm, *http.Response, error) {
	payload, err := common.MapTo[uddiipamfederation.FederatedRealm](obj, mapper.FederatedRealmUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAMFederationAPI.FederatedRealmAPI.
		Create(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIFederatedRealmToResponse(&result), httpResp, nil
}

// Read retrieves a FederatedRealm by ID
func (s *federatedRealmService) Read(ctx context.Context, id string, opts *core.Options) (*ipamfederation.FederatedRealm, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *federatedRealmService) readUDDI(ctx context.Context, id string, opts *core.Options) (*ipamfederation.FederatedRealm, *http.Response, error) {
	req := s.uddiClient.IPAMFederationAPI.FederatedRealmAPI.
		Read(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIFederatedRealmToResponse(&result), httpResp, nil
}

// Update modifies an existing FederatedRealm and returns the updated object
func (s *federatedRealmService) Update(ctx context.Context, id string, obj *ipamfederation.FederatedRealm, opts *core.Options) (*ipamfederation.FederatedRealm, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *federatedRealmService) updateUDDI(ctx context.Context, id string, obj *ipamfederation.FederatedRealm, opts *core.Options) (*ipamfederation.FederatedRealm, *http.Response, error) {
	payload, err := common.MapTo[uddiipamfederation.FederatedRealm](obj, mapper.FederatedRealmUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAMFederationAPI.FederatedRealmAPI.
		Update(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIFederatedRealmToResponse(&result), httpResp, nil
}

// Delete removes a FederatedRealm by ID
func (s *federatedRealmService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *federatedRealmService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.IPAMFederationAPI.FederatedRealmAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves FederatedRealm objects based on filter options
func (s *federatedRealmService) List(ctx context.Context, opts *core.ListOptions) ([]*ipamfederation.FederatedRealm, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *federatedRealmService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*ipamfederation.FederatedRealm, *http.Response, string, error) {
	req := s.uddiClient.IPAMFederationAPI.FederatedRealmAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.FederatedRealmFilterFieldMap[core.BackendUDDI])
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
	items := make([]*ipamfederation.FederatedRealm, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIFederatedRealmToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDIFederatedRealmToResponse(r *uddiipamfederation.FederatedRealm) *ipamfederation.FederatedRealm {
	resp := &ipamfederation.FederatedRealm{
		Id: r.Id,
	}
	resp.UDDI = &ipamfederation.UDDIFederatedRealmExt{
		Comment: r.Comment,
		Name:    r.Name,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
