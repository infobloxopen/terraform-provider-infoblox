package dns

import (
	"context"
	"fmt"
	"maps"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	uddidnsconfig "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

type AuthNsgService interface {
	Create(ctx context.Context, obj *dns.AuthNsg, opts *core.Options) (*dns.AuthNsg, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dns.AuthNsg, *http.Response, error)
	Update(ctx context.Context, id string, obj *dns.AuthNsg, opts *core.Options) (*dns.AuthNsg, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dns.AuthNsg, *http.Response, string, error)
}

type authNsgService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewAuthNsgService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) AuthNsgService {
	return &authNsgService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new AuthNsg and returns the created object
func (s *authNsgService) Create(ctx context.Context, obj *dns.AuthNsg, opts *core.Options) (*dns.AuthNsg, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *authNsgService) createUDDI(ctx context.Context, obj *dns.AuthNsg, opts *core.Options) (*dns.AuthNsg, *http.Response, error) {
	payload, err := common.MapTo[uddidnsconfig.AuthNSG](obj, mapper.AuthNsgUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSConfigurationAPI.AuthNsgAPI.
		Create(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIAuthNsgToResponse(&result), httpResp, nil
}

// Read retrieves a AuthNsg by ID
func (s *authNsgService) Read(ctx context.Context, id string, opts *core.Options) (*dns.AuthNsg, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *authNsgService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dns.AuthNsg, *http.Response, error) {
	req := s.uddiClient.DNSConfigurationAPI.AuthNsgAPI.
		Read(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIAuthNsgToResponse(&result), httpResp, nil
}

// Update modifies an existing AuthNsg and returns the updated object
func (s *authNsgService) Update(ctx context.Context, id string, obj *dns.AuthNsg, opts *core.Options) (*dns.AuthNsg, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *authNsgService) updateUDDI(ctx context.Context, id string, obj *dns.AuthNsg, opts *core.Options) (*dns.AuthNsg, *http.Response, error) {
	payload, err := common.MapTo[uddidnsconfig.AuthNSG](obj, mapper.AuthNsgUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSConfigurationAPI.AuthNsgAPI.
		Update(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIAuthNsgToResponse(&result), httpResp, nil
}

// Delete removes a AuthNsg by ID
func (s *authNsgService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *authNsgService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.DNSConfigurationAPI.AuthNsgAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves AuthNsg objects based on filter options
func (s *authNsgService) List(ctx context.Context, opts *core.ListOptions) ([]*dns.AuthNsg, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *authNsgService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dns.AuthNsg, *http.Response, string, error) {
	req := s.uddiClient.DNSConfigurationAPI.AuthNsgAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.AuthNsgFilterFieldMap[core.BackendUDDI])
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
	items := make([]*dns.AuthNsg, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIAuthNsgToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDIAuthNsgToResponse(r *uddidnsconfig.AuthNSG) *dns.AuthNsg {
	resp := &dns.AuthNsg{
		Id: r.Id,
	}
	resp.UDDI = &dns.UDDIAuthNsgExt{
		Comment:             r.Comment,
		ExternalPrimaries:   r.ExternalPrimaries,
		ExternalSecondaries: r.ExternalSecondaries,
		InternalSecondaries: r.InternalSecondaries,
		Name:                r.Name,
		Nsgs:                r.Nsgs,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
