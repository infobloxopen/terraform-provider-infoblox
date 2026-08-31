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

type ZoneForwardService interface {
	Create(ctx context.Context, obj *dns.ZoneForward, opts *core.Options) (*dns.ZoneForward, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dns.ZoneForward, *http.Response, error)
	Update(ctx context.Context, id string, obj *dns.ZoneForward, opts *core.Options) (*dns.ZoneForward, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dns.ZoneForward, *http.Response, string, error)
}

type zoneForwardService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
	uddiClient *uddiclient.APIClient
}

func NewZoneForwardService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) ZoneForwardService {
	return &zoneForwardService{
		backend:    backend,
		niosClient: nios,
		uddiClient: uddi,
	}
}

// Create creates a new ZoneForward and returns the created object
func (s *zoneForwardService) Create(ctx context.Context, obj *dns.ZoneForward, opts *core.Options) (*dns.ZoneForward, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *zoneForwardService) createNIOS(ctx context.Context, obj *dns.ZoneForward, opts *core.Options) (*dns.ZoneForward, *http.Response, error) {
	payload, err := common.MapTo[niosdns.ZoneForward](obj, mapper.ZoneForwardNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.ZoneForwardAPI.
		Create(ctx).
		ZoneForward(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateZoneForwardResponseAsObject.GetResult()

	return mapNIOSZoneForwardToResponse(&result), httpResp, nil
}

func (s *zoneForwardService) createUDDI(ctx context.Context, obj *dns.ZoneForward, opts *core.Options) (*dns.ZoneForward, *http.Response, error) {
	payload, err := common.MapTo[uddidnsconfig.ForwardZone](obj, mapper.ZoneForwardUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSConfigurationAPI.ForwardZoneAPI.
		Create(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIZoneForwardToResponse(&result), httpResp, nil
}

// Read retrieves a ZoneForward by ID
func (s *zoneForwardService) Read(ctx context.Context, id string, opts *core.Options) (*dns.ZoneForward, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *zoneForwardService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dns.ZoneForward, *http.Response, error) {
	req := s.niosClient.DNSAPI.ZoneForwardAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetZoneForwardResponseObjectAsResult.GetResult()

	return mapNIOSZoneForwardToResponse(&result), httpResp, nil
}

func (s *zoneForwardService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dns.ZoneForward, *http.Response, error) {
	req := s.uddiClient.DNSConfigurationAPI.ForwardZoneAPI.
		Read(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIZoneForwardToResponse(&result), httpResp, nil
}

// Update modifies an existing ZoneForward and returns the updated object
func (s *zoneForwardService) Update(ctx context.Context, id string, obj *dns.ZoneForward, opts *core.Options) (*dns.ZoneForward, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *zoneForwardService) updateNIOS(ctx context.Context, id string, obj *dns.ZoneForward, opts *core.Options) (*dns.ZoneForward, *http.Response, error) {
	payload, err := common.MapTo[niosdns.ZoneForward](obj, mapper.ZoneForwardNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.ZoneForwardAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		ZoneForward(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateZoneForwardResponseAsObject.GetResult()

	return mapNIOSZoneForwardToResponse(&result), httpResp, nil
}

func (s *zoneForwardService) updateUDDI(ctx context.Context, id string, obj *dns.ZoneForward, opts *core.Options) (*dns.ZoneForward, *http.Response, error) {
	payload, err := common.MapTo[uddidnsconfig.ForwardZone](obj, mapper.ZoneForwardUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSConfigurationAPI.ForwardZoneAPI.
		Update(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIZoneForwardToResponse(&result), httpResp, nil
}

// Delete removes a ZoneForward by ID
func (s *zoneForwardService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *zoneForwardService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DNSAPI.ZoneForwardAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

func (s *zoneForwardService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.DNSConfigurationAPI.ForwardZoneAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves ZoneForward objects based on filter options
func (s *zoneForwardService) List(ctx context.Context, opts *core.ListOptions) ([]*dns.ZoneForward, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *zoneForwardService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dns.ZoneForward, *http.Response, string, error) {
	req := s.niosClient.DNSAPI.ZoneForwardAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.ZoneForwardFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListZoneForwardResponseObject.GetResult()
	items := make([]*dns.ZoneForward, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSZoneForwardToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListZoneForwardResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func (s *zoneForwardService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dns.ZoneForward, *http.Response, string, error) {
	req := s.uddiClient.DNSConfigurationAPI.ForwardZoneAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.ZoneForwardFilterFieldMap[core.BackendUDDI])
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
	items := make([]*dns.ZoneForward, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIZoneForwardToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapNIOSZoneForwardToResponse(r *niosdns.ZoneForward) *dns.ZoneForward {
	resp := &dns.ZoneForward{
		Id: r.Ref,
	}
	resp.NIOS = &dns.NIOSZoneForwardExt{
		Comment:             r.Comment,
		Disable:             r.Disable,
		DisableNsGeneration: r.DisableNsGeneration,
		ExternalNsGroup:     r.ExternalNsGroup,
		ForwardTo:           r.ForwardTo,
		ForwardersOnly:      r.ForwardersOnly,
		ForwardingServers:   r.ForwardingServers,
		Fqdn:                r.Fqdn,
		Locked:              r.Locked,
		MsAdIntegrated:      r.MsAdIntegrated,
		MsDdnsMode:          r.MsDdnsMode,
		NsGroup:             r.NsGroup,
		Prefix:              r.Prefix,
		View:                r.View,
		ZoneFormat:          r.ZoneFormat,
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

func mapUDDIZoneForwardToResponse(r *uddidnsconfig.ForwardZone) *dns.ZoneForward {
	resp := &dns.ZoneForward{
		Id: r.Id,
	}
	resp.UDDI = &dns.UDDIZoneForwardExt{
		Comment:            r.Comment,
		CompartmentId:      r.CompartmentId,
		Disabled:           r.Disabled,
		ExternalForwarders: r.ExternalForwarders,
		ForwardOnly:        r.ForwardOnly,
		Fqdn:               r.Fqdn,
		Hosts:              r.Hosts,
		InternalForwarders: r.InternalForwarders,
		Nsgs:               r.Nsgs,
		Parent:             r.Parent,
		View:               r.View,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
