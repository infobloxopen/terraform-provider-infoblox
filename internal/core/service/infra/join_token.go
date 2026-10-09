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
	uddiinfraprovision "github.com/infobloxopen/universal-ddi-go-client/infraprovision"
)

type JoinTokenService interface {
	Create(ctx context.Context, obj *infra.JoinToken, opts *core.Options) (*infra.JoinToken, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*infra.JoinToken, *http.Response, error)
	Update(ctx context.Context, id string, obj *infra.JoinToken, opts *core.Options) (*infra.JoinToken, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*infra.JoinToken, *http.Response, string, error)
}

type joinTokenService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewJoinTokenService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) JoinTokenService {
	return &joinTokenService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new JoinToken and returns the created object
func (s *joinTokenService) Create(ctx context.Context, obj *infra.JoinToken, opts *core.Options) (*infra.JoinToken, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *joinTokenService) createUDDI(ctx context.Context, obj *infra.JoinToken, opts *core.Options) (*infra.JoinToken, *http.Response, error) {
	payload, err := common.MapTo[uddiinfraprovision.JoinToken](obj, mapper.JoinTokenUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.HostActivationAPI.UIJoinTokenAPI.
		Create(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()
	created := mapUDDIJoinTokenToResponse(&result)
	// These live on the create response envelope, not on the object, and are
	// never returned again. Carry them across before the result is handed back.
	if created.UDDI != nil {
		created.UDDI.JoinToken = resp.JoinToken
	}

	return created, httpResp, nil
}

// Read retrieves a JoinToken by ID
func (s *joinTokenService) Read(ctx context.Context, id string, opts *core.Options) (*infra.JoinToken, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *joinTokenService) readUDDI(ctx context.Context, id string, opts *core.Options) (*infra.JoinToken, *http.Response, error) {
	req := s.uddiClient.HostActivationAPI.UIJoinTokenAPI.
		Read(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIJoinTokenToResponse(&result), httpResp, nil
}

// Update modifies an existing JoinToken and returns the updated object
func (s *joinTokenService) Update(ctx context.Context, id string, obj *infra.JoinToken, opts *core.Options) (*infra.JoinToken, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *joinTokenService) updateUDDI(ctx context.Context, id string, obj *infra.JoinToken, opts *core.Options) (*infra.JoinToken, *http.Response, error) {
	payload, err := common.MapTo[uddiinfraprovision.JoinToken](obj, mapper.JoinTokenUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.HostActivationAPI.UIJoinTokenAPI.
		Update(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIJoinTokenToResponse(&result), httpResp, nil
}

// Delete removes a JoinToken by ID
func (s *joinTokenService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *joinTokenService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.HostActivationAPI.UIJoinTokenAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves JoinToken objects based on filter options
func (s *joinTokenService) List(ctx context.Context, opts *core.ListOptions) ([]*infra.JoinToken, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *joinTokenService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*infra.JoinToken, *http.Response, string, error) {
	req := s.uddiClient.HostActivationAPI.UIJoinTokenAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.JoinTokenFilterFieldMap[core.BackendUDDI])
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
	items := make([]*infra.JoinToken, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIJoinTokenToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDIJoinTokenToResponse(r *uddiinfraprovision.JoinToken) *infra.JoinToken {
	resp := &infra.JoinToken{
		Id: r.Id,
	}
	resp.UDDI = &infra.UDDIJoinTokenExt{
		DeletedAt:   r.DeletedAt,
		Description: r.Description,
		ExpiresAt:   r.ExpiresAt,
		LastUsedAt:  r.LastUsedAt,
		Name:        r.Name,
		Status:      r.Status,
		TokenId:     r.TokenId,
		UseCounter:  r.UseCounter,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
