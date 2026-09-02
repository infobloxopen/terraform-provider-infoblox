package keys

import (
	"context"
	"fmt"
	"maps"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/keys"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/keys"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	uddikeys "github.com/infobloxopen/universal-ddi-go-client/keys"
)

type TsigKeyService interface {
	Create(ctx context.Context, obj *keys.TsigKey, opts *core.Options) (*keys.TsigKey, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*keys.TsigKey, *http.Response, error)
	Update(ctx context.Context, id string, obj *keys.TsigKey, opts *core.Options) (*keys.TsigKey, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*keys.TsigKey, *http.Response, string, error)
}

type tsigKeyService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewTsigKeyService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) TsigKeyService {
	return &tsigKeyService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new TsigKey and returns the created object
func (s *tsigKeyService) Create(ctx context.Context, obj *keys.TsigKey, opts *core.Options) (*keys.TsigKey, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *tsigKeyService) createUDDI(ctx context.Context, obj *keys.TsigKey, opts *core.Options) (*keys.TsigKey, *http.Response, error) {
	payload, err := common.MapTo[uddikeys.TSIGKey](obj, mapper.TsigKeyUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.KeysAPI.TsigAPI.
		Create(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDITsigKeyToResponse(&result), httpResp, nil
}

// Read retrieves a TsigKey by ID
func (s *tsigKeyService) Read(ctx context.Context, id string, opts *core.Options) (*keys.TsigKey, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *tsigKeyService) readUDDI(ctx context.Context, id string, opts *core.Options) (*keys.TsigKey, *http.Response, error) {
	req := s.uddiClient.KeysAPI.TsigAPI.
		Read(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDITsigKeyToResponse(&result), httpResp, nil
}

// Update modifies an existing TsigKey and returns the updated object
func (s *tsigKeyService) Update(ctx context.Context, id string, obj *keys.TsigKey, opts *core.Options) (*keys.TsigKey, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *tsigKeyService) updateUDDI(ctx context.Context, id string, obj *keys.TsigKey, opts *core.Options) (*keys.TsigKey, *http.Response, error) {
	payload, err := common.MapTo[uddikeys.TSIGKey](obj, mapper.TsigKeyUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.KeysAPI.TsigAPI.
		Update(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDITsigKeyToResponse(&result), httpResp, nil
}

// Delete removes a TsigKey by ID
func (s *tsigKeyService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *tsigKeyService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.KeysAPI.TsigAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves TsigKey objects based on filter options
func (s *tsigKeyService) List(ctx context.Context, opts *core.ListOptions) ([]*keys.TsigKey, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *tsigKeyService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*keys.TsigKey, *http.Response, string, error) {
	req := s.uddiClient.KeysAPI.TsigAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.TsigKeyFilterFieldMap[core.BackendUDDI])
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
	items := make([]*keys.TsigKey, 0, len(results))
	for i := range results {
		items = append(items, mapUDDITsigKeyToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDITsigKeyToResponse(r *uddikeys.TSIGKey) *keys.TsigKey {
	resp := &keys.TsigKey{
		Id: r.Id,
	}
	resp.UDDI = &keys.UDDITsigKeyExt{
		Algorithm: r.Algorithm,
		Comment:   r.Comment,
		Name:      r.Name,
		Secret:    r.Secret,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
