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
	uddidnsdata "github.com/infobloxopen/universal-ddi-go-client/dnsdata"
)

type RecordAaaaService interface {
	Create(ctx context.Context, obj *dns.RecordAaaa, opts *core.Options) (*dns.RecordAaaa, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dns.RecordAaaa, *http.Response, error)
	Update(ctx context.Context, id string, obj *dns.RecordAaaa, opts *core.Options) (*dns.RecordAaaa, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dns.RecordAaaa, *http.Response, string, error)
}

type recordAaaaService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
	uddiClient *uddiclient.APIClient
}

func NewRecordAaaaService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) RecordAaaaService {
	return &recordAaaaService{
		backend:    backend,
		niosClient: nios,
		uddiClient: uddi,
	}
}

// Create creates a new RecordAaaa and returns the created object
func (s *recordAaaaService) Create(ctx context.Context, obj *dns.RecordAaaa, opts *core.Options) (*dns.RecordAaaa, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordAaaaService) createNIOS(ctx context.Context, obj *dns.RecordAaaa, opts *core.Options) (*dns.RecordAaaa, *http.Response, error) {
	payload, err := common.MapTo[niosdns.RecordAaaa](obj, mapper.RecordAaaaNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if payload.FuncCall != nil && payload.Ipv6addr == nil {
		payload.Ipv6addr = &niosdns.RecordAaaaIpv6addr{}
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.RecordAaaaAPI.
		Create(ctx).
		RecordAaaa(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateRecordAaaaResponseAsObject.GetResult()

	return mapNIOSRecordAaaaToResponse(&result), httpResp, nil
}

func (s *recordAaaaService) createUDDI(ctx context.Context, obj *dns.RecordAaaa, opts *core.Options) (*dns.RecordAaaa, *http.Response, error) {
	payload, err := common.MapTo[uddidnsdata.Record](obj, mapper.RecordAaaaUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSDataAPI.RecordAPI.
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

	return mapUDDIRecordAaaaToResponse(&result), httpResp, nil
}

// Read retrieves a RecordAaaa by ID
func (s *recordAaaaService) Read(ctx context.Context, id string, opts *core.Options) (*dns.RecordAaaa, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordAaaaService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dns.RecordAaaa, *http.Response, error) {
	req := s.niosClient.DNSAPI.RecordAaaaAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetRecordAaaaResponseObjectAsResult.GetResult()

	return mapNIOSRecordAaaaToResponse(&result), httpResp, nil
}

func (s *recordAaaaService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dns.RecordAaaa, *http.Response, error) {
	req := s.uddiClient.DNSDataAPI.RecordAPI.
		Read(ctx, id)

	if opts != nil && opts.Inherit != "" {
		req = req.Inherit(opts.Inherit)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIRecordAaaaToResponse(&result), httpResp, nil
}

// Update modifies an existing RecordAaaa and returns the updated object
func (s *recordAaaaService) Update(ctx context.Context, id string, obj *dns.RecordAaaa, opts *core.Options) (*dns.RecordAaaa, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordAaaaService) updateNIOS(ctx context.Context, id string, obj *dns.RecordAaaa, opts *core.Options) (*dns.RecordAaaa, *http.Response, error) {
	payload, err := common.MapTo[niosdns.RecordAaaa](obj, mapper.RecordAaaaNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.RecordAaaaAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		RecordAaaa(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateRecordAaaaResponseAsObject.GetResult()

	return mapNIOSRecordAaaaToResponse(&result), httpResp, nil
}

func (s *recordAaaaService) updateUDDI(ctx context.Context, id string, obj *dns.RecordAaaa, opts *core.Options) (*dns.RecordAaaa, *http.Response, error) {
	payload, err := common.MapTo[uddidnsdata.Record](obj, mapper.RecordAaaaUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSDataAPI.RecordAPI.
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

	return mapUDDIRecordAaaaToResponse(&result), httpResp, nil
}

// Delete removes a RecordAaaa by ID
func (s *recordAaaaService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordAaaaService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DNSAPI.RecordAaaaAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

func (s *recordAaaaService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.DNSDataAPI.RecordAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves RecordAaaa objects based on filter options
func (s *recordAaaaService) List(ctx context.Context, opts *core.ListOptions) ([]*dns.RecordAaaa, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordAaaaService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dns.RecordAaaa, *http.Response, string, error) {
	req := s.niosClient.DNSAPI.RecordAaaaAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.RecordAaaaFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListRecordAaaaResponseObject.GetResult()
	items := make([]*dns.RecordAaaa, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSRecordAaaaToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListRecordAaaaResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func (s *recordAaaaService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dns.RecordAaaa, *http.Response, string, error) {
	req := s.uddiClient.DNSDataAPI.RecordAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.RecordAaaaFilterFieldMap[core.BackendUDDI])
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
	items := make([]*dns.RecordAaaa, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIRecordAaaaToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapNIOSRecordAaaaToResponse(r *niosdns.RecordAaaa) *dns.RecordAaaa {
	resp := &dns.RecordAaaa{
		Id: r.Ref,
	}
	resp.NIOS = &dns.NIOSRecordAaaaExt{
		CloudInfo:         r.CloudInfo,
		Comment:           r.Comment,
		Creator:           r.Creator,
		DdnsPrincipal:     r.DdnsPrincipal,
		DdnsProtected:     r.DdnsProtected,
		Disable:           r.Disable,
		ForbidReclamation: r.ForbidReclamation,
		Name:              r.Name,
		Ttl:               r.Ttl,
		UseTtl:            r.UseTtl,
		View:              r.View,
	}
	if r.Ipv6addr != nil {
		resp.NIOS.Ipv6addr = r.Ipv6addr.String
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

func mapUDDIRecordAaaaToResponse(r *uddidnsdata.Record) *dns.RecordAaaa {
	resp := &dns.RecordAaaa{
		Id: r.Id,
	}
	resp.UDDI = &dns.UDDIRecordAaaaExt{
		AbsoluteNameSpec:   r.AbsoluteNameSpec,
		Comment:            r.Comment,
		Disabled:           r.Disabled,
		InheritanceSources: r.InheritanceSources,
		NameInZone:         r.NameInZone,
		Options:            r.Options,
		Rdata:              r.Rdata,
		Ttl:                r.Ttl,
		Type:               r.Type,
		View:               r.View,
		Zone:               r.Zone,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
