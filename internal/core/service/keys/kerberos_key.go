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

type KerberosKeyService interface {
	Read(ctx context.Context, id string, opts *core.Options) (*keys.KerberosKey, *http.Response, error)
	Update(ctx context.Context, id string, obj *keys.KerberosKey, opts *core.Options) (*keys.KerberosKey, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*keys.KerberosKey, *http.Response, string, error)
}

type kerberosKeyService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewKerberosKeyService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) KerberosKeyService {
	return &kerberosKeyService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Read retrieves a KerberosKey by ID
func (s *kerberosKeyService) Read(ctx context.Context, id string, opts *core.Options) (*keys.KerberosKey, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *kerberosKeyService) readUDDI(ctx context.Context, id string, opts *core.Options) (*keys.KerberosKey, *http.Response, error) {
	req := s.uddiClient.KeysAPI.KerberosAPI.
		Read(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIKerberosKeyToResponse(&result), httpResp, nil
}

// Update modifies an existing KerberosKey and returns the updated object
func (s *kerberosKeyService) Update(ctx context.Context, id string, obj *keys.KerberosKey, opts *core.Options) (*keys.KerberosKey, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *kerberosKeyService) updateUDDI(ctx context.Context, id string, obj *keys.KerberosKey, opts *core.Options) (*keys.KerberosKey, *http.Response, error) {
	payload, err := common.MapTo[uddikeys.KerberosKey](obj, mapper.KerberosKeyUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.KeysAPI.KerberosAPI.
		Update(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIKerberosKeyToResponse(&result), httpResp, nil
}

// Delete removes a KerberosKey by ID
func (s *kerberosKeyService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *kerberosKeyService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.KeysAPI.KerberosAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves KerberosKey objects based on filter options
func (s *kerberosKeyService) List(ctx context.Context, opts *core.ListOptions) ([]*keys.KerberosKey, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *kerberosKeyService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*keys.KerberosKey, *http.Response, string, error) {
	req := s.uddiClient.KeysAPI.KerberosAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.KerberosKeyFilterFieldMap[core.BackendUDDI])
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
	items := make([]*keys.KerberosKey, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIKerberosKeyToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDIKerberosKeyToResponse(r *uddikeys.KerberosKey) *keys.KerberosKey {
	resp := &keys.KerberosKey{
		Id: r.Id,
	}
	resp.UDDI = &keys.UDDIKerberosKeyExt{
		Algorithm:  r.Algorithm,
		Comment:    r.Comment,
		Domain:     r.Domain,
		Principal:  r.Principal,
		UploadedAt: r.UploadedAt,
		Version:    r.Version,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
