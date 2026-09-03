package dns

import (
	"context"
	"fmt"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
)

type ZoneStubService interface {
	Create(ctx context.Context, obj *dns.ZoneStub, opts *core.Options) (*dns.ZoneStub, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dns.ZoneStub, *http.Response, error)
	Update(ctx context.Context, id string, obj *dns.ZoneStub, opts *core.Options) (*dns.ZoneStub, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dns.ZoneStub, *http.Response, string, error)
}

type zoneStubService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewZoneStubService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) ZoneStubService {
	return &zoneStubService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new ZoneStub and returns the created object
func (s *zoneStubService) Create(ctx context.Context, obj *dns.ZoneStub, opts *core.Options) (*dns.ZoneStub, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *zoneStubService) createNIOS(ctx context.Context, obj *dns.ZoneStub, opts *core.Options) (*dns.ZoneStub, *http.Response, error) {
	payload, err := common.MapTo[niosdns.ZoneStub](obj, mapper.ZoneStubNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.ZoneStubAPI.
		Create(ctx).
		ZoneStub(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateZoneStubResponseAsObject.GetResult()

	return mapNIOSZoneStubToResponse(&result), httpResp, nil
}

// Read retrieves a ZoneStub by ID
func (s *zoneStubService) Read(ctx context.Context, id string, opts *core.Options) (*dns.ZoneStub, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *zoneStubService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dns.ZoneStub, *http.Response, error) {
	req := s.niosClient.DNSAPI.ZoneStubAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetZoneStubResponseObjectAsResult.GetResult()

	return mapNIOSZoneStubToResponse(&result), httpResp, nil
}

// Update modifies an existing ZoneStub and returns the updated object
func (s *zoneStubService) Update(ctx context.Context, id string, obj *dns.ZoneStub, opts *core.Options) (*dns.ZoneStub, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *zoneStubService) updateNIOS(ctx context.Context, id string, obj *dns.ZoneStub, opts *core.Options) (*dns.ZoneStub, *http.Response, error) {
	payload, err := common.MapTo[niosdns.ZoneStub](obj, mapper.ZoneStubNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.ZoneStubAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		ZoneStub(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateZoneStubResponseAsObject.GetResult()

	return mapNIOSZoneStubToResponse(&result), httpResp, nil
}

// Delete removes a ZoneStub by ID
func (s *zoneStubService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *zoneStubService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DNSAPI.ZoneStubAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves ZoneStub objects based on filter options
func (s *zoneStubService) List(ctx context.Context, opts *core.ListOptions) ([]*dns.ZoneStub, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *zoneStubService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dns.ZoneStub, *http.Response, string, error) {
	req := s.niosClient.DNSAPI.ZoneStubAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.ZoneStubFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListZoneStubResponseObject.GetResult()
	items := make([]*dns.ZoneStub, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSZoneStubToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListZoneStubResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSZoneStubToResponse(r *niosdns.ZoneStub) *dns.ZoneStub {
	resp := &dns.ZoneStub{
		Id: r.Ref,
	}
	resp.NIOS = &dns.NIOSZoneStubExt{
		Comment:           r.Comment,
		Disable:           r.Disable,
		DisableForwarding: r.DisableForwarding,
		ExternalNsGroup:   r.ExternalNsGroup,
		Fqdn:              r.Fqdn,
		Locked:            r.Locked,
		MsAdIntegrated:    r.MsAdIntegrated,
		MsDdnsMode:        r.MsDdnsMode,
		NsGroup:           r.NsGroup,
		Prefix:            r.Prefix,
		StubFrom:          r.StubFrom,
		StubMembers:       r.StubMembers,
		StubMsservers:     r.StubMsservers,
		View:              r.View,
		ZoneFormat:        r.ZoneFormat,
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
