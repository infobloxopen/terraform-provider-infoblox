package dtc

import (
	"context"
	"fmt"
	"maps"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niosdtc "github.com/infobloxopen/infoblox-nios-go-client/dtc"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/dtc"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dtc"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	uddidnsconfig "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

type DtcLbdnService interface {
	Create(ctx context.Context, obj *dtc.DtcLbdn, opts *core.Options) (*dtc.DtcLbdn, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dtc.DtcLbdn, *http.Response, error)
	Update(ctx context.Context, id string, obj *dtc.DtcLbdn, opts *core.Options) (*dtc.DtcLbdn, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcLbdn, *http.Response, string, error)
}

type dtcLbdnService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
	uddiClient *uddiclient.APIClient
}

func NewDtcLbdnService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) DtcLbdnService {
	return &dtcLbdnService{
		backend:    backend,
		niosClient: nios,
		uddiClient: uddi,
	}
}

// Create creates a new DtcLbdn and returns the created object
func (s *dtcLbdnService) Create(ctx context.Context, obj *dtc.DtcLbdn, opts *core.Options) (*dtc.DtcLbdn, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcLbdnService) createNIOS(ctx context.Context, obj *dtc.DtcLbdn, opts *core.Options) (*dtc.DtcLbdn, *http.Response, error) {
	payload, err := common.MapTo[niosdtc.DtcLbdn](obj, mapper.DtcLbdnNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DTCAPI.DtcLbdnAPI.
		Create(ctx).
		DtcLbdn(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateDtcLbdnResponseAsObject.GetResult()

	return mapNIOSDtcLbdnToResponse(&result), httpResp, nil
}

func (s *dtcLbdnService) createUDDI(ctx context.Context, obj *dtc.DtcLbdn, opts *core.Options) (*dtc.DtcLbdn, *http.Response, error) {
	payload, err := common.MapTo[uddidnsconfig.LBDN](obj, mapper.DtcLbdnUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSConfigurationAPI.LbdnAPI.
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

	return mapUDDIDtcLbdnToResponse(&result), httpResp, nil
}

// Read retrieves a DtcLbdn by ID
func (s *dtcLbdnService) Read(ctx context.Context, id string, opts *core.Options) (*dtc.DtcLbdn, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcLbdnService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dtc.DtcLbdn, *http.Response, error) {
	req := s.niosClient.DTCAPI.DtcLbdnAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetDtcLbdnResponseObjectAsResult.GetResult()

	return mapNIOSDtcLbdnToResponse(&result), httpResp, nil
}

func (s *dtcLbdnService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dtc.DtcLbdn, *http.Response, error) {
	req := s.uddiClient.DNSConfigurationAPI.LbdnAPI.
		Read(ctx, id)

	if opts != nil && opts.Inherit != "" {
		req = req.Inherit(opts.Inherit)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIDtcLbdnToResponse(&result), httpResp, nil
}

// Update modifies an existing DtcLbdn and returns the updated object
func (s *dtcLbdnService) Update(ctx context.Context, id string, obj *dtc.DtcLbdn, opts *core.Options) (*dtc.DtcLbdn, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcLbdnService) updateNIOS(ctx context.Context, id string, obj *dtc.DtcLbdn, opts *core.Options) (*dtc.DtcLbdn, *http.Response, error) {
	payload, err := common.MapTo[niosdtc.DtcLbdn](obj, mapper.DtcLbdnNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DTCAPI.DtcLbdnAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		DtcLbdn(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateDtcLbdnResponseAsObject.GetResult()

	return mapNIOSDtcLbdnToResponse(&result), httpResp, nil
}

func (s *dtcLbdnService) updateUDDI(ctx context.Context, id string, obj *dtc.DtcLbdn, opts *core.Options) (*dtc.DtcLbdn, *http.Response, error) {
	payload, err := common.MapTo[uddidnsconfig.LBDN](obj, mapper.DtcLbdnUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSConfigurationAPI.LbdnAPI.
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

	return mapUDDIDtcLbdnToResponse(&result), httpResp, nil
}

// Delete removes a DtcLbdn by ID
func (s *dtcLbdnService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcLbdnService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DTCAPI.DtcLbdnAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

func (s *dtcLbdnService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.DNSConfigurationAPI.LbdnAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves DtcLbdn objects based on filter options
func (s *dtcLbdnService) List(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcLbdn, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcLbdnService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcLbdn, *http.Response, string, error) {
	req := s.niosClient.DTCAPI.DtcLbdnAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.DtcLbdnFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListDtcLbdnResponseObject.GetResult()
	items := make([]*dtc.DtcLbdn, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSDtcLbdnToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListDtcLbdnResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func (s *dtcLbdnService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcLbdn, *http.Response, string, error) {
	req := s.uddiClient.DNSConfigurationAPI.LbdnAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.DtcLbdnFilterFieldMap[core.BackendUDDI])
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
	items := make([]*dtc.DtcLbdn, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIDtcLbdnToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapNIOSDtcLbdnToResponse(r *niosdtc.DtcLbdn) *dtc.DtcLbdn {
	resp := &dtc.DtcLbdn{
		Id: r.Ref,
	}
	resp.NIOS = &dtc.NIOSDtcLbdnExt{
		AuthZones:                r.AuthZones,
		AutoConsolidatedMonitors: r.AutoConsolidatedMonitors,
		Comment:                  r.Comment,
		Disable:                  r.Disable,
		Health:                   r.Health,
		LbMethod:                 r.LbMethod,
		Name:                     r.Name,
		Patterns:                 r.Patterns,
		Persistence:              r.Persistence,
		Pools:                    r.Pools,
		Priority:                 r.Priority,
		Topology:                 r.Topology,
		Ttl:                      r.Ttl,
		Types:                    r.Types,
		UseTtl:                   r.UseTtl,
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

func mapUDDIDtcLbdnToResponse(r *uddidnsconfig.LBDN) *dtc.DtcLbdn {
	resp := &dtc.DtcLbdn{
		Id: r.Id,
	}
	resp.UDDI = &dtc.UDDIDtcLbdnExt{
		Comment:            r.Comment,
		Disabled:           r.Disabled,
		DtcPolicy:          r.DtcPolicy,
		InheritanceSources: r.InheritanceSources,
		Name:               r.Name,
		Precedence:         r.Precedence,
		Ttl:                r.Ttl,
		View:               r.View,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
