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

type ForwardNsgService interface {
	Create(ctx context.Context, obj *dns.ForwardNsg, opts *core.Options) (*dns.ForwardNsg, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dns.ForwardNsg, *http.Response, error)
	Update(ctx context.Context, id string, obj *dns.ForwardNsg, opts *core.Options) (*dns.ForwardNsg, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dns.ForwardNsg, *http.Response, string, error)
}

type forwardNsgService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewForwardNsgService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) ForwardNsgService {
	return &forwardNsgService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new ForwardNsg and returns the created object
func (s *forwardNsgService) Create(ctx context.Context, obj *dns.ForwardNsg, opts *core.Options) (*dns.ForwardNsg, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *forwardNsgService) createUDDI(ctx context.Context, obj *dns.ForwardNsg, opts *core.Options) (*dns.ForwardNsg, *http.Response, error) {
	payload, err := common.MapTo[uddidnsconfig.ForwardNSG](obj, mapper.ForwardNsgUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSConfigurationAPI.ForwardNsgAPI.
		Create(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIForwardNsgToResponse(&result), httpResp, nil
}

// Read retrieves a ForwardNsg by ID
func (s *forwardNsgService) Read(ctx context.Context, id string, opts *core.Options) (*dns.ForwardNsg, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *forwardNsgService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dns.ForwardNsg, *http.Response, error) {
	req := s.uddiClient.DNSConfigurationAPI.ForwardNsgAPI.
		Read(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIForwardNsgToResponse(&result), httpResp, nil
}

// Update modifies an existing ForwardNsg and returns the updated object
func (s *forwardNsgService) Update(ctx context.Context, id string, obj *dns.ForwardNsg, opts *core.Options) (*dns.ForwardNsg, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *forwardNsgService) updateUDDI(ctx context.Context, id string, obj *dns.ForwardNsg, opts *core.Options) (*dns.ForwardNsg, *http.Response, error) {
	payload, err := common.MapTo[uddidnsconfig.ForwardNSG](obj, mapper.ForwardNsgUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSConfigurationAPI.ForwardNsgAPI.
		Update(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIForwardNsgToResponse(&result), httpResp, nil
}

// Delete removes a ForwardNsg by ID
func (s *forwardNsgService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *forwardNsgService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.DNSConfigurationAPI.ForwardNsgAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves ForwardNsg objects based on filter options
func (s *forwardNsgService) List(ctx context.Context, opts *core.ListOptions) ([]*dns.ForwardNsg, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *forwardNsgService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dns.ForwardNsg, *http.Response, string, error) {
	req := s.uddiClient.DNSConfigurationAPI.ForwardNsgAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.ForwardNsgFilterFieldMap[core.BackendUDDI])
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
	items := make([]*dns.ForwardNsg, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIForwardNsgToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDIForwardNsgToResponse(r *uddidnsconfig.ForwardNSG) *dns.ForwardNsg {
	resp := &dns.ForwardNsg{
		Id: r.Id,
	}
	resp.UDDI = &dns.UDDIForwardNsgExt{
		Comment:            r.Comment,
		ExternalForwarders: r.ExternalForwarders,
		ForwardersOnly:     r.ForwardersOnly,
		Hosts:              r.Hosts,
		InternalForwarders: r.InternalForwarders,
		Name:               r.Name,
		Nsgs:               r.Nsgs,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
