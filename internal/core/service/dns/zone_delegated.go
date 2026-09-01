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

type ZoneDelegatedService interface {
	Create(ctx context.Context, obj *dns.ZoneDelegated, opts *core.Options) (*dns.ZoneDelegated, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dns.ZoneDelegated, *http.Response, error)
	Update(ctx context.Context, id string, obj *dns.ZoneDelegated, opts *core.Options) (*dns.ZoneDelegated, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dns.ZoneDelegated, *http.Response, string, error)
}

type zoneDelegatedService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
	uddiClient *uddiclient.APIClient
}

func NewZoneDelegatedService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) ZoneDelegatedService {
	return &zoneDelegatedService{
		backend:    backend,
		niosClient: nios,
		uddiClient: uddi,
	}
}

// Create creates a new ZoneDelegated and returns the created object
func (s *zoneDelegatedService) Create(ctx context.Context, obj *dns.ZoneDelegated, opts *core.Options) (*dns.ZoneDelegated, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *zoneDelegatedService) createNIOS(ctx context.Context, obj *dns.ZoneDelegated, opts *core.Options) (*dns.ZoneDelegated, *http.Response, error) {
	payload, err := common.MapTo[niosdns.ZoneDelegated](obj, mapper.ZoneDelegatedNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.ZoneDelegatedAPI.
		Create(ctx).
		ZoneDelegated(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateZoneDelegatedResponseAsObject.GetResult()

	return mapNIOSZoneDelegatedToResponse(&result), httpResp, nil
}

func (s *zoneDelegatedService) createUDDI(ctx context.Context, obj *dns.ZoneDelegated, opts *core.Options) (*dns.ZoneDelegated, *http.Response, error) {
	payload, err := common.MapTo[uddidnsconfig.Delegation](obj, mapper.ZoneDelegatedUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSConfigurationAPI.DelegationAPI.
		Create(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIZoneDelegatedToResponse(&result), httpResp, nil
}

// Read retrieves a ZoneDelegated by ID
func (s *zoneDelegatedService) Read(ctx context.Context, id string, opts *core.Options) (*dns.ZoneDelegated, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *zoneDelegatedService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dns.ZoneDelegated, *http.Response, error) {
	req := s.niosClient.DNSAPI.ZoneDelegatedAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetZoneDelegatedResponseObjectAsResult.GetResult()

	return mapNIOSZoneDelegatedToResponse(&result), httpResp, nil
}

func (s *zoneDelegatedService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dns.ZoneDelegated, *http.Response, error) {
	req := s.uddiClient.DNSConfigurationAPI.DelegationAPI.
		Read(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIZoneDelegatedToResponse(&result), httpResp, nil
}

// Update modifies an existing ZoneDelegated and returns the updated object
func (s *zoneDelegatedService) Update(ctx context.Context, id string, obj *dns.ZoneDelegated, opts *core.Options) (*dns.ZoneDelegated, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *zoneDelegatedService) updateNIOS(ctx context.Context, id string, obj *dns.ZoneDelegated, opts *core.Options) (*dns.ZoneDelegated, *http.Response, error) {
	payload, err := common.MapTo[niosdns.ZoneDelegated](obj, mapper.ZoneDelegatedNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.ZoneDelegatedAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		ZoneDelegated(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateZoneDelegatedResponseAsObject.GetResult()

	return mapNIOSZoneDelegatedToResponse(&result), httpResp, nil
}

func (s *zoneDelegatedService) updateUDDI(ctx context.Context, id string, obj *dns.ZoneDelegated, opts *core.Options) (*dns.ZoneDelegated, *http.Response, error) {
	payload, err := common.MapTo[uddidnsconfig.Delegation](obj, mapper.ZoneDelegatedUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSConfigurationAPI.DelegationAPI.
		Update(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIZoneDelegatedToResponse(&result), httpResp, nil
}

// Delete removes a ZoneDelegated by ID
func (s *zoneDelegatedService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *zoneDelegatedService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DNSAPI.ZoneDelegatedAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

func (s *zoneDelegatedService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.DNSConfigurationAPI.DelegationAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves ZoneDelegated objects based on filter options
func (s *zoneDelegatedService) List(ctx context.Context, opts *core.ListOptions) ([]*dns.ZoneDelegated, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *zoneDelegatedService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dns.ZoneDelegated, *http.Response, string, error) {
	req := s.niosClient.DNSAPI.ZoneDelegatedAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.ZoneDelegatedFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListZoneDelegatedResponseObject.GetResult()
	items := make([]*dns.ZoneDelegated, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSZoneDelegatedToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListZoneDelegatedResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func (s *zoneDelegatedService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dns.ZoneDelegated, *http.Response, string, error) {
	req := s.uddiClient.DNSConfigurationAPI.DelegationAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.ZoneDelegatedFilterFieldMap[core.BackendUDDI])
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
	items := make([]*dns.ZoneDelegated, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIZoneDelegatedToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapNIOSZoneDelegatedToResponse(r *niosdns.ZoneDelegated) *dns.ZoneDelegated {
	resp := &dns.ZoneDelegated{
		Id: r.Ref,
	}
	resp.NIOS = &dns.NIOSZoneDelegatedExt{
		Comment:                r.Comment,
		DelegateTo:             r.DelegateTo,
		DelegatedTtl:           r.DelegatedTtl,
		Disable:                r.Disable,
		EnableRfc2317Exclusion: r.EnableRfc2317Exclusion,
		Fqdn:                   r.Fqdn,
		Locked:                 r.Locked,
		MsAdIntegrated:         r.MsAdIntegrated,
		MsDdnsMode:             r.MsDdnsMode,
		NsGroup:                r.NsGroup,
		Prefix:                 r.Prefix,
		UseDelegatedTtl:        r.UseDelegatedTtl,
		View:                   r.View,
		ZoneFormat:             r.ZoneFormat,
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

func mapUDDIZoneDelegatedToResponse(r *uddidnsconfig.Delegation) *dns.ZoneDelegated {
	resp := &dns.ZoneDelegated{
		Id: r.Id,
	}
	resp.UDDI = &dns.UDDIZoneDelegatedExt{
		Comment:           r.Comment,
		CompartmentId:     r.CompartmentId,
		DelegationServers: r.DelegationServers,
		Disabled:          r.Disabled,
		Fqdn:              r.Fqdn,
		Parent:            r.Parent,
		View:              r.View,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
