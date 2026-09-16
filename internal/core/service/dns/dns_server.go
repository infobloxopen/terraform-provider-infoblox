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

type DnsServerService interface {
	Create(ctx context.Context, obj *dns.DnsServer, opts *core.Options) (*dns.DnsServer, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dns.DnsServer, *http.Response, error)
	Update(ctx context.Context, id string, obj *dns.DnsServer, opts *core.Options) (*dns.DnsServer, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dns.DnsServer, *http.Response, string, error)
}

type dnsServerService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewDnsServerService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) DnsServerService {
	return &dnsServerService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new DnsServer and returns the created object
func (s *dnsServerService) Create(ctx context.Context, obj *dns.DnsServer, opts *core.Options) (*dns.DnsServer, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dnsServerService) createUDDI(ctx context.Context, obj *dns.DnsServer, opts *core.Options) (*dns.DnsServer, *http.Response, error) {
	payload, err := common.MapTo[uddidnsconfig.Server](obj, mapper.DnsServerUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSConfigurationAPI.ServerAPI.
		Create(ctx).
		Body(payload)

	if opts != nil && opts.Inherit != "" {
		req = req.Inherit(opts.Inherit)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIDnsServerToResponse(&result), httpResp, nil
}

// Read retrieves a DnsServer by ID
func (s *dnsServerService) Read(ctx context.Context, id string, opts *core.Options) (*dns.DnsServer, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dnsServerService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dns.DnsServer, *http.Response, error) {
	req := s.uddiClient.DNSConfigurationAPI.ServerAPI.
		Read(ctx, id)

	if opts != nil && opts.Inherit != "" {
		req = req.Inherit(opts.Inherit)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIDnsServerToResponse(&result), httpResp, nil
}

// Update modifies an existing DnsServer and returns the updated object
func (s *dnsServerService) Update(ctx context.Context, id string, obj *dns.DnsServer, opts *core.Options) (*dns.DnsServer, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dnsServerService) updateUDDI(ctx context.Context, id string, obj *dns.DnsServer, opts *core.Options) (*dns.DnsServer, *http.Response, error) {
	payload, err := common.MapTo[uddidnsconfig.Server](obj, mapper.DnsServerUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSConfigurationAPI.ServerAPI.
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

	return mapUDDIDnsServerToResponse(&result), httpResp, nil
}

// Delete removes a DnsServer by ID
func (s *dnsServerService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dnsServerService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.DNSConfigurationAPI.ServerAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves DnsServer objects based on filter options
func (s *dnsServerService) List(ctx context.Context, opts *core.ListOptions) ([]*dns.DnsServer, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dnsServerService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dns.DnsServer, *http.Response, string, error) {
	req := s.uddiClient.DNSConfigurationAPI.ServerAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.DnsServerFilterFieldMap[core.BackendUDDI])
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
	items := make([]*dns.DnsServer, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIDnsServerToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDIDnsServerToResponse(r *uddidnsconfig.Server) *dns.DnsServer {
	resp := &dns.DnsServer{
		Id: r.Id,
	}
	resp.UDDI = &dns.UDDIDnsServerExt{
		AddEdnsOptionInOutgoingQuery:      r.AddEdnsOptionInOutgoingQuery,
		AutoSortViews:                     r.AutoSortViews,
		Comment:                           r.Comment,
		CustomRootNs:                      r.CustomRootNs,
		CustomRootNsEnabled:               r.CustomRootNsEnabled,
		DnssecEnableValidation:            r.DnssecEnableValidation,
		DnssecEnabled:                     r.DnssecEnabled,
		DnssecTrustAnchors:                r.DnssecTrustAnchors,
		DnssecValidateExpiry:              r.DnssecValidateExpiry,
		EcsEnabled:                        r.EcsEnabled,
		EcsForwarding:                     r.EcsForwarding,
		EcsPrefixV4:                       r.EcsPrefixV4,
		EcsPrefixV6:                       r.EcsPrefixV6,
		EcsZones:                          r.EcsZones,
		FilterAaaaAcl:                     r.FilterAaaaAcl,
		FilterAaaaOnV4:                    r.FilterAaaaOnV4,
		Forwarders:                        r.Forwarders,
		ForwardersOnly:                    r.ForwardersOnly,
		GssTsigEnabled:                    r.GssTsigEnabled,
		InheritanceSources:                r.InheritanceSources,
		KerberosKeys:                      r.KerberosKeys,
		LameTtl:                           r.LameTtl,
		LogQueryResponse:                  r.LogQueryResponse,
		MatchRecursiveOnly:                r.MatchRecursiveOnly,
		MaxCacheTtl:                       r.MaxCacheTtl,
		MaxNegativeTtl:                    r.MaxNegativeTtl,
		MinimalResponses:                  r.MinimalResponses,
		Name:                              r.Name,
		Notify:                            r.Notify,
		QueryAcl:                          r.QueryAcl,
		QueryPort:                         r.QueryPort,
		RecursionAcl:                      r.RecursionAcl,
		RecursionEnabled:                  r.RecursionEnabled,
		RecursiveClients:                  r.RecursiveClients,
		ResolverQueryTimeout:              r.ResolverQueryTimeout,
		SecondaryAxfrQueryLimit:           r.SecondaryAxfrQueryLimit,
		SecondarySoaQueryLimit:            r.SecondarySoaQueryLimit,
		SortList:                          r.SortList,
		SynthesizeAddressRecordsFromHttps: r.SynthesizeAddressRecordsFromHttps,
		TransferAcl:                       r.TransferAcl,
		UpdateAcl:                         r.UpdateAcl,
		UseForwardersForSubzones:          r.UseForwardersForSubzones,
		UseRootForwardersForLocalResolutionWithB1td: r.UseRootForwardersForLocalResolutionWithB1td,
		Views: r.Views,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
