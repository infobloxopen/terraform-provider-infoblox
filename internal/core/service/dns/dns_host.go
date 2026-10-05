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

type DnsHostService interface {
	Create(ctx context.Context, obj *dns.DnsHost, opts *core.Options) (*dns.DnsHost, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dns.DnsHost, *http.Response, error)
	Update(ctx context.Context, id string, obj *dns.DnsHost, opts *core.Options) (*dns.DnsHost, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dns.DnsHost, *http.Response, string, error)
}

type dnsHostService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewDnsHostService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) DnsHostService {
	return &dnsHostService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new DnsHost and returns the created object
func (s *dnsHostService) Create(ctx context.Context, obj *dns.DnsHost, opts *core.Options) (*dns.DnsHost, *http.Response, error) {
	switch s.backend {
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

// Read retrieves a DnsHost by ID
func (s *dnsHostService) Read(ctx context.Context, id string, opts *core.Options) (*dns.DnsHost, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dnsHostService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dns.DnsHost, *http.Response, error) {
	req := s.uddiClient.DNSConfigurationAPI.HostAPI.
		Read(ctx, id)

	if opts != nil && opts.Inherit != "" {
		req = req.Inherit(opts.Inherit)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIDnsHostToResponse(&result), httpResp, nil
}

// Update modifies an existing DnsHost and returns the updated object
func (s *dnsHostService) Update(ctx context.Context, id string, obj *dns.DnsHost, opts *core.Options) (*dns.DnsHost, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dnsHostService) updateUDDI(ctx context.Context, id string, obj *dns.DnsHost, opts *core.Options) (*dns.DnsHost, *http.Response, error) {
	payload, err := common.MapTo[uddidnsconfig.Host](obj, mapper.DnsHostUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSConfigurationAPI.HostAPI.
		Update(ctx, id).
		Body(payload)

	if opts != nil && opts.Inherit != "" {
		req = req.Inherit(opts.Inherit)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIDnsHostToResponse(&result), httpResp, nil
}

// Delete removes a DnsHost by ID
func (s *dnsHostService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

// List retrieves DnsHost objects based on filter options
func (s *dnsHostService) List(ctx context.Context, opts *core.ListOptions) ([]*dns.DnsHost, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dnsHostService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dns.DnsHost, *http.Response, string, error) {
	req := s.uddiClient.DNSConfigurationAPI.HostAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.DnsHostFilterFieldMap[core.BackendUDDI])
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
	items := make([]*dns.DnsHost, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIDnsHostToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDIDnsHostToResponse(r *uddidnsconfig.Host) *dns.DnsHost {
	resp := &dns.DnsHost{
		Id: r.Id,
	}
	resp.UDDI = &dns.UDDIDnsHostExt{
		AbsoluteName:       r.AbsoluteName,
		Address:            r.Address,
		AnycastAddresses:   r.AnycastAddresses,
		AssociatedServer:   r.AssociatedServer,
		DfpService:         r.DfpService,
		InheritanceSources: r.InheritanceSources,
		KerberosKeys:       r.KerberosKeys,
		Name:               r.Name,
		Ophid:              r.Ophid,
		ProviderId:         r.ProviderId,
		Server:             r.Server,
		Type:               r.Type,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
