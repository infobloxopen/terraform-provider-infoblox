package redirect

import (
	"context"
	"fmt"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/redirect"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/redirect"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	uddiredirect "github.com/infobloxopen/universal-ddi-go-client/redirect"
)

type CustomRedirectService interface {
	Create(ctx context.Context, obj *redirect.CustomRedirect, opts *core.Options) (*redirect.CustomRedirect, *http.Response, error)
	Read(ctx context.Context, id int32, opts *core.Options) (*redirect.CustomRedirect, *http.Response, error)
	Update(ctx context.Context, id int32, obj *redirect.CustomRedirect, opts *core.Options) (*redirect.CustomRedirect, *http.Response, error)
	Delete(ctx context.Context, id int32) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*redirect.CustomRedirect, *http.Response, string, error)
}

type customRedirectService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewCustomRedirectService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) CustomRedirectService {
	return &customRedirectService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new CustomRedirect and returns the created object
func (s *customRedirectService) Create(ctx context.Context, obj *redirect.CustomRedirect, opts *core.Options) (*redirect.CustomRedirect, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *customRedirectService) createUDDI(ctx context.Context, obj *redirect.CustomRedirect, opts *core.Options) (*redirect.CustomRedirect, *http.Response, error) {
	payload, err := common.MapTo[uddiredirect.CustomRedirect](obj, mapper.CustomRedirectUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.RedirectAPI.CustomRedirectsAPI.
		CreateCustomRedirect(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResults()

	return mapUDDICustomRedirectToResponse(&result), httpResp, nil
}

// Read retrieves a CustomRedirect by ID
func (s *customRedirectService) Read(ctx context.Context, id int32, opts *core.Options) (*redirect.CustomRedirect, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *customRedirectService) readUDDI(ctx context.Context, id int32, opts *core.Options) (*redirect.CustomRedirect, *http.Response, error) {
	req := s.uddiClient.RedirectAPI.CustomRedirectsAPI.
		ReadCustomRedirect(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResults()

	return mapUDDICustomRedirectToResponse(&result), httpResp, nil
}

// Update modifies an existing CustomRedirect and returns the updated object
func (s *customRedirectService) Update(ctx context.Context, id int32, obj *redirect.CustomRedirect, opts *core.Options) (*redirect.CustomRedirect, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *customRedirectService) updateUDDI(ctx context.Context, id int32, obj *redirect.CustomRedirect, opts *core.Options) (*redirect.CustomRedirect, *http.Response, error) {
	payload, err := common.MapTo[uddiredirect.CustomRedirect](obj, mapper.CustomRedirectUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.RedirectAPI.CustomRedirectsAPI.
		UpdateCustomRedirect(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResults()

	return mapUDDICustomRedirectToResponse(&result), httpResp, nil
}

// Delete removes a CustomRedirect by ID
func (s *customRedirectService) Delete(ctx context.Context, id int32) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *customRedirectService) deleteUDDI(ctx context.Context, id int32) (*http.Response, error) {
	httpResp, err := s.uddiClient.RedirectAPI.CustomRedirectsAPI.
		DeleteSingleCustomRedirect(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves CustomRedirect objects based on filter options
func (s *customRedirectService) List(ctx context.Context, opts *core.ListOptions) ([]*redirect.CustomRedirect, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *customRedirectService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*redirect.CustomRedirect, *http.Response, string, error) {
	req := s.uddiClient.RedirectAPI.CustomRedirectsAPI.ListCustomRedirect(ctx)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.CustomRedirectFilterFieldMap[core.BackendUDDI])
		for k, v := range translatedFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		if len(filters) > 0 {
			req = req.Filter(core.JoinFilters(filters))
		}

	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, "", err
	}

	results := resp.GetResults()
	items := make([]*redirect.CustomRedirect, 0, len(results))
	for i := range results {
		items = append(items, mapUDDICustomRedirectToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDICustomRedirectToResponse(r *uddiredirect.CustomRedirect) *redirect.CustomRedirect {
	resp := &redirect.CustomRedirect{
		Id: r.Id,
	}
	resp.UDDI = &redirect.UDDICustomRedirectExt{
		Data: r.Data,
		Name: r.Name,
	}
	return resp
}
