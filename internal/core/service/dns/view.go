package dns

import (
	"context"
	"fmt"
	"maps"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	uddidnsconfig "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

type ViewService interface {
	Create(ctx context.Context, obj *dns.View, opts *core.Options) (*dns.View, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dns.View, *http.Response, error)
	Update(ctx context.Context, id string, obj *dns.View, opts *core.Options) (*dns.View, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dns.View, *http.Response, string, error)
}

type viewService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
	uddiClient *uddiclient.APIClient
}

func NewViewService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) ViewService {
	return &viewService{
		backend:    backend,
		niosClient: nios,
		uddiClient: uddi,
	}
}

// Create creates a new View and returns the created object
func (s *viewService) Create(ctx context.Context, obj *dns.View, opts *core.Options) (*dns.View, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *viewService) createNIOS(ctx context.Context, obj *dns.View, opts *core.Options) (*dns.View, *http.Response, error) {
	payload, err := common.MapTo[niosdns.View](obj, mapper.ViewNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.ViewAPI.
		Create(ctx).
		View(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateViewResponseAsObject.GetResult()

	return mapNIOSViewToResponse(&result), httpResp, nil
}

func (s *viewService) createUDDI(ctx context.Context, obj *dns.View, opts *core.Options) (*dns.View, *http.Response, error) {
	payload, err := common.MapTo[uddidnsconfig.View](obj, mapper.ViewUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSConfigurationAPI.ViewAPI.
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

	return mapUDDIViewToResponse(&result), httpResp, nil
}

// Read retrieves a View by ID
func (s *viewService) Read(ctx context.Context, id string, opts *core.Options) (*dns.View, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *viewService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dns.View, *http.Response, error) {
	req := s.niosClient.DNSAPI.ViewAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetViewResponseObjectAsResult.GetResult()

	return mapNIOSViewToResponse(&result), httpResp, nil
}

func (s *viewService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dns.View, *http.Response, error) {
	req := s.uddiClient.DNSConfigurationAPI.ViewAPI.
		Read(ctx, id)

	if opts != nil && opts.Inherit != "" {
		req = req.Inherit(opts.Inherit)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIViewToResponse(&result), httpResp, nil
}

// Update modifies an existing View and returns the updated object
func (s *viewService) Update(ctx context.Context, id string, obj *dns.View, opts *core.Options) (*dns.View, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *viewService) updateNIOS(ctx context.Context, id string, obj *dns.View, opts *core.Options) (*dns.View, *http.Response, error) {
	payload, err := common.MapTo[niosdns.View](obj, mapper.ViewNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.ViewAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		View(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateViewResponseAsObject.GetResult()

	return mapNIOSViewToResponse(&result), httpResp, nil
}

func (s *viewService) updateUDDI(ctx context.Context, id string, obj *dns.View, opts *core.Options) (*dns.View, *http.Response, error) {
	payload, err := common.MapTo[uddidnsconfig.View](obj, mapper.ViewUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSConfigurationAPI.ViewAPI.
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

	return mapUDDIViewToResponse(&result), httpResp, nil
}

// Delete removes a View by ID
func (s *viewService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *viewService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DNSAPI.ViewAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

func (s *viewService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.DNSConfigurationAPI.ViewAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves View objects based on filter options
func (s *viewService) List(ctx context.Context, opts *core.ListOptions) ([]*dns.View, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *viewService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dns.View, *http.Response, string, error) {
	req := s.niosClient.DNSAPI.ViewAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.ViewFilterFieldMap[core.BackendNIOS])
			filters := make(map[string]any, len(translatedFilters))
			for k, v := range translatedFilters {
				filters[k] = v
			}
			req = req.Filters(filters)
		}
		if len(opts.ExtAttrFilter) > 0 {
			extAttrFilters := make(map[string]any, len(opts.ExtAttrFilter))
			for k, v := range opts.ExtAttrFilter {
				extAttrFilters[k] = v
			}
			req = req.Extattrfilter(extAttrFilters)
		}
		if opts.PageID != "" {
			req = req.PageId(opts.PageID)
		}
		req = req.Paging(opts.Paging)
		maxResults := opts.MaxResults
		if maxResults <= 0 {
			maxResults = core.DefaultListLimit
		}
		req = req.MaxResults(maxResults)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, "", err
	}

	results := resp.ListViewResponseObject.GetResult()
	items := make([]*dns.View, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSViewToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListViewResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func (s *viewService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dns.View, *http.Response, string, error) {
	req := s.uddiClient.DNSConfigurationAPI.ViewAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.ViewFilterFieldMap[core.BackendUDDI])
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
	items := make([]*dns.View, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIViewToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapNIOSViewToResponse(r *niosdns.View) *dns.View {
	resp := &dns.View{
		Id: r.Ref,
	}
	resp.NIOS = &dns.NIOSViewExt{
		BlacklistAction:                     r.BlacklistAction,
		BlacklistLogQuery:                   r.BlacklistLogQuery,
		BlacklistRedirectAddresses:          r.BlacklistRedirectAddresses,
		BlacklistRedirectTtl:                r.BlacklistRedirectTtl,
		BlacklistRulesets:                   r.BlacklistRulesets,
		CloudInfo:                           r.CloudInfo,
		Comment:                             r.Comment,
		CustomRootNameServers:               r.CustomRootNameServers,
		DdnsForceCreationTimestampUpdate:    r.DdnsForceCreationTimestampUpdate,
		DdnsPrincipalGroup:                  r.DdnsPrincipalGroup,
		DdnsPrincipalTracking:               r.DdnsPrincipalTracking,
		DdnsRestrictPatterns:                r.DdnsRestrictPatterns,
		DdnsRestrictPatternsList:            r.DdnsRestrictPatternsList,
		DdnsRestrictProtected:               r.DdnsRestrictProtected,
		DdnsRestrictSecure:                  r.DdnsRestrictSecure,
		DdnsRestrictStatic:                  r.DdnsRestrictStatic,
		Disable:                             r.Disable,
		Dns64Enabled:                        r.Dns64Enabled,
		Dns64Groups:                         r.Dns64Groups,
		DnssecEnabled:                       r.DnssecEnabled,
		DnssecExpiredSignaturesEnabled:      r.DnssecExpiredSignaturesEnabled,
		DnssecNegativeTrustAnchors:          r.DnssecNegativeTrustAnchors,
		DnssecTrustedKeys:                   r.DnssecTrustedKeys,
		DnssecValidationEnabled:             r.DnssecValidationEnabled,
		EdnsUdpSize:                         r.EdnsUdpSize,
		EnableBlacklist:                     r.EnableBlacklist,
		EnableFixedRrsetOrderFqdns:          r.EnableFixedRrsetOrderFqdns,
		EnableMatchRecursiveOnly:            r.EnableMatchRecursiveOnly,
		FilterAaaa:                          r.FilterAaaa,
		FilterAaaaList:                      r.FilterAaaaList,
		FixedRrsetOrderFqdns:                r.FixedRrsetOrderFqdns,
		ForwardOnly:                         r.ForwardOnly,
		Forwarders:                          r.Forwarders,
		LastQueriedAcl:                      r.LastQueriedAcl,
		MatchClients:                        r.MatchClients,
		MatchDestinations:                   r.MatchDestinations,
		MaxCacheTtl:                         r.MaxCacheTtl,
		MaxNcacheTtl:                        r.MaxNcacheTtl,
		MaxUdpSize:                          r.MaxUdpSize,
		Name:                                r.Name,
		NetworkView:                         r.NetworkView,
		NotifyDelay:                         r.NotifyDelay,
		NxdomainLogQuery:                    r.NxdomainLogQuery,
		NxdomainRedirect:                    r.NxdomainRedirect,
		NxdomainRedirectAddresses:           r.NxdomainRedirectAddresses,
		NxdomainRedirectAddressesV6:         r.NxdomainRedirectAddressesV6,
		NxdomainRedirectTtl:                 r.NxdomainRedirectTtl,
		NxdomainRulesets:                    r.NxdomainRulesets,
		Recursion:                           r.Recursion,
		ResponseRateLimiting:                r.ResponseRateLimiting,
		RootNameServerType:                  r.RootNameServerType,
		RpzDropIpRuleEnabled:                r.RpzDropIpRuleEnabled,
		RpzDropIpRuleMinPrefixLengthIpv4:    r.RpzDropIpRuleMinPrefixLengthIpv4,
		RpzDropIpRuleMinPrefixLengthIpv6:    r.RpzDropIpRuleMinPrefixLengthIpv6,
		RpzQnameWaitRecurse:                 r.RpzQnameWaitRecurse,
		ScavengingSettings:                  r.ScavengingSettings,
		Sortlist:                            r.Sortlist,
		UseBlacklist:                        r.UseBlacklist,
		UseDdnsForceCreationTimestampUpdate: r.UseDdnsForceCreationTimestampUpdate,
		UseDdnsPatternsRestriction:          r.UseDdnsPatternsRestriction,
		UseDdnsPrincipalSecurity:            r.UseDdnsPrincipalSecurity,
		UseDdnsRestrictProtected:            r.UseDdnsRestrictProtected,
		UseDdnsRestrictStatic:               r.UseDdnsRestrictStatic,
		UseDns64:                            r.UseDns64,
		UseDnssec:                           r.UseDnssec,
		UseEdnsUdpSize:                      r.UseEdnsUdpSize,
		UseFilterAaaa:                       r.UseFilterAaaa,
		UseFixedRrsetOrderFqdns:             r.UseFixedRrsetOrderFqdns,
		UseForwarders:                       r.UseForwarders,
		UseMaxCacheTtl:                      r.UseMaxCacheTtl,
		UseMaxNcacheTtl:                     r.UseMaxNcacheTtl,
		UseMaxUdpSize:                       r.UseMaxUdpSize,
		UseNxdomainRedirect:                 r.UseNxdomainRedirect,
		UseRecursion:                        r.UseRecursion,
		UseResponseRateLimiting:             r.UseResponseRateLimiting,
		UseRootNameServer:                   r.UseRootNameServer,
		UseRpzDropIpRule:                    r.UseRpzDropIpRule,
		UseRpzQnameWaitRecurse:              r.UseRpzQnameWaitRecurse,
		UseScavengingSettings:               r.UseScavengingSettings,
		UseSortlist:                         r.UseSortlist,
	}
	if r.ExtAttrs != nil {
		attrs := make(map[string]any, len(*r.ExtAttrs))
		for k, v := range *r.ExtAttrs {
			attrs[k] = core.StringifyEAValue(v.Value)
		}
		resp.NIOS.ExtAttrs = attrs
	}
	return resp
}

func mapUDDIViewToResponse(r *uddidnsconfig.View) *dns.View {
	resp := &dns.View{
		Id: r.Id,
	}
	resp.UDDI = &dns.UDDIViewExt{
		AddEdnsOptionInOutgoingQuery:      r.AddEdnsOptionInOutgoingQuery,
		Comment:                           r.Comment,
		CompartmentId:                     r.CompartmentId,
		CustomRootNs:                      r.CustomRootNs,
		CustomRootNsEnabled:               r.CustomRootNsEnabled,
		Disabled:                          r.Disabled,
		DnssecEnableValidation:            r.DnssecEnableValidation,
		DnssecEnabled:                     r.DnssecEnabled,
		DnssecTrustAnchors:                r.DnssecTrustAnchors,
		DnssecValidateExpiry:              r.DnssecValidateExpiry,
		DtcConfig:                         r.DtcConfig,
		EcsEnabled:                        r.EcsEnabled,
		EcsForwarding:                     r.EcsForwarding,
		EcsPrefixV4:                       r.EcsPrefixV4,
		EcsPrefixV6:                       r.EcsPrefixV6,
		EcsZones:                          r.EcsZones,
		EdnsUdpSize:                       r.EdnsUdpSize,
		FilterAaaaAcl:                     r.FilterAaaaAcl,
		FilterAaaaOnV4:                    r.FilterAaaaOnV4,
		Forwarders:                        r.Forwarders,
		ForwardersOnly:                    r.ForwardersOnly,
		GssTsigEnabled:                    r.GssTsigEnabled,
		InheritanceSources:                r.InheritanceSources,
		IpSpaces:                          r.IpSpaces,
		LameTtl:                           r.LameTtl,
		MatchClientsAcl:                   r.MatchClientsAcl,
		MatchDestinationsAcl:              r.MatchDestinationsAcl,
		MatchRecursiveOnly:                r.MatchRecursiveOnly,
		MaxCacheTtl:                       r.MaxCacheTtl,
		MaxNegativeTtl:                    r.MaxNegativeTtl,
		MaxUdpSize:                        r.MaxUdpSize,
		MinimalResponses:                  r.MinimalResponses,
		Name:                              r.Name,
		Notify:                            r.Notify,
		QueryAcl:                          r.QueryAcl,
		RecursionAcl:                      r.RecursionAcl,
		RecursionEnabled:                  r.RecursionEnabled,
		SortList:                          r.SortList,
		SynthesizeAddressRecordsFromHttps: r.SynthesizeAddressRecordsFromHttps,
		TransferAcl:                       r.TransferAcl,
		UpdateAcl:                         r.UpdateAcl,
		UseForwardersForSubzones:          r.UseForwardersForSubzones,
		UseRootForwardersForLocalResolutionWithB1td: r.UseRootForwardersForLocalResolutionWithB1td,
		ZoneAuthority: r.ZoneAuthority,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
