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

type RecordCaaService interface {
	Create(ctx context.Context, obj *dns.RecordCaa, opts *core.Options) (*dns.RecordCaa, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dns.RecordCaa, *http.Response, error)
	Update(ctx context.Context, id string, obj *dns.RecordCaa, opts *core.Options) (*dns.RecordCaa, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dns.RecordCaa, *http.Response, string, error)
}

type recordCaaService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
	uddiClient *uddiclient.APIClient
}

func NewRecordCaaService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) RecordCaaService {
	return &recordCaaService{
		backend:    backend,
		niosClient: nios,
		uddiClient: uddi,
	}
}

// Create creates a new RecordCaa and returns the created object
func (s *recordCaaService) Create(ctx context.Context, obj *dns.RecordCaa, opts *core.Options) (*dns.RecordCaa, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordCaaService) createNIOS(ctx context.Context, obj *dns.RecordCaa, opts *core.Options) (*dns.RecordCaa, *http.Response, error) {
	payload, err := common.MapTo[niosdns.RecordCaa](obj, mapper.RecordCaaNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.RecordCaaAPI.
		Create(ctx).
		RecordCaa(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateRecordCaaResponseAsObject.GetResult()

	return mapNIOSRecordCaaToResponse(&result), httpResp, nil
}

func (s *recordCaaService) createUDDI(ctx context.Context, obj *dns.RecordCaa, opts *core.Options) (*dns.RecordCaa, *http.Response, error) {
	payload, err := common.MapTo[uddidnsdata.Record](obj, mapper.RecordCaaUDDIFieldMap)
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

	return mapUDDIRecordCaaToResponse(&result), httpResp, nil
}

// Read retrieves a RecordCaa by ID
func (s *recordCaaService) Read(ctx context.Context, id string, opts *core.Options) (*dns.RecordCaa, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordCaaService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dns.RecordCaa, *http.Response, error) {
	req := s.niosClient.DNSAPI.RecordCaaAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetRecordCaaResponseObjectAsResult.GetResult()

	return mapNIOSRecordCaaToResponse(&result), httpResp, nil
}

func (s *recordCaaService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dns.RecordCaa, *http.Response, error) {
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

	return mapUDDIRecordCaaToResponse(&result), httpResp, nil
}

// Update modifies an existing RecordCaa and returns the updated object
func (s *recordCaaService) Update(ctx context.Context, id string, obj *dns.RecordCaa, opts *core.Options) (*dns.RecordCaa, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordCaaService) updateNIOS(ctx context.Context, id string, obj *dns.RecordCaa, opts *core.Options) (*dns.RecordCaa, *http.Response, error) {
	payload, err := common.MapTo[niosdns.RecordCaa](obj, mapper.RecordCaaNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.RecordCaaAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		RecordCaa(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateRecordCaaResponseAsObject.GetResult()

	return mapNIOSRecordCaaToResponse(&result), httpResp, nil
}

func (s *recordCaaService) updateUDDI(ctx context.Context, id string, obj *dns.RecordCaa, opts *core.Options) (*dns.RecordCaa, *http.Response, error) {
	payload, err := common.MapTo[uddidnsdata.Record](obj, mapper.RecordCaaUDDIFieldMap)
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

	return mapUDDIRecordCaaToResponse(&result), httpResp, nil
}

// Delete removes a RecordCaa by ID
func (s *recordCaaService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordCaaService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DNSAPI.RecordCaaAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

func (s *recordCaaService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.DNSDataAPI.RecordAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves RecordCaa objects based on filter options
func (s *recordCaaService) List(ctx context.Context, opts *core.ListOptions) ([]*dns.RecordCaa, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordCaaService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dns.RecordCaa, *http.Response, string, error) {
	req := s.niosClient.DNSAPI.RecordCaaAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.RecordCaaFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListRecordCaaResponseObject.GetResult()
	items := make([]*dns.RecordCaa, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSRecordCaaToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListRecordCaaResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func (s *recordCaaService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dns.RecordCaa, *http.Response, string, error) {
	req := s.uddiClient.DNSDataAPI.RecordAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.RecordCaaFilterFieldMap[core.BackendUDDI])
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
	items := make([]*dns.RecordCaa, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIRecordCaaToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapNIOSRecordCaaToResponse(r *niosdns.RecordCaa) *dns.RecordCaa {
	resp := &dns.RecordCaa{
		Id: r.Ref,
	}
	resp.NIOS = &dns.NIOSRecordCaaExt{
		CaFlag:            r.CaFlag,
		CaTag:             r.CaTag,
		CaValue:           r.CaValue,
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
	if r.ExtAttrs != nil {
		attrs := make(map[string]any, len(*r.ExtAttrs))
		for k, v := range *r.ExtAttrs {
			attrs[k] = core.StringifyEAValue(v.Value)
		}
		resp.NIOS.ExtAttrs = attrs
	}
	return resp
}

func mapUDDIRecordCaaToResponse(r *uddidnsdata.Record) *dns.RecordCaa {
	resp := &dns.RecordCaa{
		Id: r.Id,
	}
	resp.UDDI = &dns.UDDIRecordCaaExt{
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
