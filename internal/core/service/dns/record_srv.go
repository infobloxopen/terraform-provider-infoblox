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

type RecordSrvService interface {
	Create(ctx context.Context, obj *dns.RecordSrv, opts *core.Options) (*dns.RecordSrv, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dns.RecordSrv, *http.Response, error)
	Update(ctx context.Context, id string, obj *dns.RecordSrv, opts *core.Options) (*dns.RecordSrv, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dns.RecordSrv, *http.Response, string, error)
}

type recordSrvService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
	uddiClient *uddiclient.APIClient
}

func NewRecordSrvService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) RecordSrvService {
	return &recordSrvService{
		backend:    backend,
		niosClient: nios,
		uddiClient: uddi,
	}
}

// Create creates a new RecordSrv and returns the created object
func (s *recordSrvService) Create(ctx context.Context, obj *dns.RecordSrv, opts *core.Options) (*dns.RecordSrv, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordSrvService) createNIOS(ctx context.Context, obj *dns.RecordSrv, opts *core.Options) (*dns.RecordSrv, *http.Response, error) {
	payload, err := common.MapTo[niosdns.RecordSrv](obj, mapper.RecordSrvNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.RecordSrvAPI.
		Create(ctx).
		RecordSrv(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateRecordSrvResponseAsObject.GetResult()

	return mapNIOSRecordSrvToResponse(&result), httpResp, nil
}

func (s *recordSrvService) createUDDI(ctx context.Context, obj *dns.RecordSrv, opts *core.Options) (*dns.RecordSrv, *http.Response, error) {
	payload, err := common.MapTo[uddidnsdata.Record](obj, mapper.RecordSrvUDDIFieldMap)
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

	return mapUDDIRecordSrvToResponse(&result), httpResp, nil
}

// Read retrieves a RecordSrv by ID
func (s *recordSrvService) Read(ctx context.Context, id string, opts *core.Options) (*dns.RecordSrv, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordSrvService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dns.RecordSrv, *http.Response, error) {
	req := s.niosClient.DNSAPI.RecordSrvAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetRecordSrvResponseObjectAsResult.GetResult()

	return mapNIOSRecordSrvToResponse(&result), httpResp, nil
}

func (s *recordSrvService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dns.RecordSrv, *http.Response, error) {
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

	return mapUDDIRecordSrvToResponse(&result), httpResp, nil
}

// Update modifies an existing RecordSrv and returns the updated object
func (s *recordSrvService) Update(ctx context.Context, id string, obj *dns.RecordSrv, opts *core.Options) (*dns.RecordSrv, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordSrvService) updateNIOS(ctx context.Context, id string, obj *dns.RecordSrv, opts *core.Options) (*dns.RecordSrv, *http.Response, error) {
	payload, err := common.MapTo[niosdns.RecordSrv](obj, mapper.RecordSrvNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.RecordSrvAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		RecordSrv(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateRecordSrvResponseAsObject.GetResult()

	return mapNIOSRecordSrvToResponse(&result), httpResp, nil
}

func (s *recordSrvService) updateUDDI(ctx context.Context, id string, obj *dns.RecordSrv, opts *core.Options) (*dns.RecordSrv, *http.Response, error) {
	payload, err := common.MapTo[uddidnsdata.Record](obj, mapper.RecordSrvUDDIFieldMap)
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

	return mapUDDIRecordSrvToResponse(&result), httpResp, nil
}

// Delete removes a RecordSrv by ID
func (s *recordSrvService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordSrvService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DNSAPI.RecordSrvAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

func (s *recordSrvService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.DNSDataAPI.RecordAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves RecordSrv objects based on filter options
func (s *recordSrvService) List(ctx context.Context, opts *core.ListOptions) ([]*dns.RecordSrv, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordSrvService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dns.RecordSrv, *http.Response, string, error) {
	req := s.niosClient.DNSAPI.RecordSrvAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.RecordSrvFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListRecordSrvResponseObject.GetResult()
	items := make([]*dns.RecordSrv, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSRecordSrvToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListRecordSrvResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func (s *recordSrvService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dns.RecordSrv, *http.Response, string, error) {
	req := s.uddiClient.DNSDataAPI.RecordAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.RecordSrvFilterFieldMap[core.BackendUDDI])
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
	items := make([]*dns.RecordSrv, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIRecordSrvToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapNIOSRecordSrvToResponse(r *niosdns.RecordSrv) *dns.RecordSrv {
	resp := &dns.RecordSrv{
		Id: r.Ref,
	}
	resp.NIOS = &dns.NIOSRecordSrvExt{
		CloudInfo:         r.CloudInfo,
		Comment:           r.Comment,
		Creator:           r.Creator,
		DdnsPrincipal:     r.DdnsPrincipal,
		DdnsProtected:     r.DdnsProtected,
		Disable:           r.Disable,
		ForbidReclamation: r.ForbidReclamation,
		Name:              r.Name,
		Port:              r.Port,
		Priority:          r.Priority,
		Target:            r.Target,
		Ttl:               r.Ttl,
		UseTtl:            r.UseTtl,
		View:              r.View,
		Weight:            r.Weight,
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

func mapUDDIRecordSrvToResponse(r *uddidnsdata.Record) *dns.RecordSrv {
	resp := &dns.RecordSrv{
		Id: r.Id,
	}
	resp.UDDI = &dns.UDDIRecordSrvExt{
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
