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
	uddidtc "github.com/infobloxopen/universal-ddi-go-client/dtc"
)

type DtcPoolService interface {
	Create(ctx context.Context, obj *dtc.DtcPool, opts *core.Options) (*dtc.DtcPool, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dtc.DtcPool, *http.Response, error)
	Update(ctx context.Context, id string, obj *dtc.DtcPool, opts *core.Options) (*dtc.DtcPool, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcPool, *http.Response, string, error)
}

type dtcPoolService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
	uddiClient *uddiclient.APIClient
}

func NewDtcPoolService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) DtcPoolService {
	return &dtcPoolService{
		backend:    backend,
		niosClient: nios,
		uddiClient: uddi,
	}
}

// Create creates a new DtcPool and returns the created object
func (s *dtcPoolService) Create(ctx context.Context, obj *dtc.DtcPool, opts *core.Options) (*dtc.DtcPool, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcPoolService) createNIOS(ctx context.Context, obj *dtc.DtcPool, opts *core.Options) (*dtc.DtcPool, *http.Response, error) {
	payload, err := common.MapTo[niosdtc.DtcPool](obj, mapper.DtcPoolNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DTCAPI.DtcPoolAPI.
		Create(ctx).
		DtcPool(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateDtcPoolResponseAsObject.GetResult()

	return mapNIOSDtcPoolToResponse(&result), httpResp, nil
}

func (s *dtcPoolService) createUDDI(ctx context.Context, obj *dtc.DtcPool, opts *core.Options) (*dtc.DtcPool, *http.Response, error) {
	payload, err := common.MapTo[uddidtc.Pool](obj, mapper.DtcPoolUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSTrafficControlAPI.PoolAPI.
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

	return mapUDDIDtcPoolToResponse(&result), httpResp, nil
}

// Read retrieves a DtcPool by ID
func (s *dtcPoolService) Read(ctx context.Context, id string, opts *core.Options) (*dtc.DtcPool, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcPoolService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dtc.DtcPool, *http.Response, error) {
	req := s.niosClient.DTCAPI.DtcPoolAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetDtcPoolResponseObjectAsResult.GetResult()

	return mapNIOSDtcPoolToResponse(&result), httpResp, nil
}

func (s *dtcPoolService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dtc.DtcPool, *http.Response, error) {
	req := s.uddiClient.DNSTrafficControlAPI.PoolAPI.
		Read(ctx, id)

	if opts != nil && opts.Inherit != "" {
		req = req.Inherit(opts.Inherit)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIDtcPoolToResponse(&result), httpResp, nil
}

// Update modifies an existing DtcPool and returns the updated object
func (s *dtcPoolService) Update(ctx context.Context, id string, obj *dtc.DtcPool, opts *core.Options) (*dtc.DtcPool, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcPoolService) updateNIOS(ctx context.Context, id string, obj *dtc.DtcPool, opts *core.Options) (*dtc.DtcPool, *http.Response, error) {
	payload, err := common.MapTo[niosdtc.DtcPool](obj, mapper.DtcPoolNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DTCAPI.DtcPoolAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		DtcPool(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateDtcPoolResponseAsObject.GetResult()

	return mapNIOSDtcPoolToResponse(&result), httpResp, nil
}

func (s *dtcPoolService) updateUDDI(ctx context.Context, id string, obj *dtc.DtcPool, opts *core.Options) (*dtc.DtcPool, *http.Response, error) {
	payload, err := common.MapTo[uddidtc.Pool](obj, mapper.DtcPoolUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSTrafficControlAPI.PoolAPI.
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

	return mapUDDIDtcPoolToResponse(&result), httpResp, nil
}

// Delete removes a DtcPool by ID
func (s *dtcPoolService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcPoolService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DTCAPI.DtcPoolAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

func (s *dtcPoolService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.DNSTrafficControlAPI.PoolAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves DtcPool objects based on filter options
func (s *dtcPoolService) List(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcPool, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcPoolService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcPool, *http.Response, string, error) {
	req := s.niosClient.DTCAPI.DtcPoolAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.DtcPoolFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListDtcPoolResponseObject.GetResult()
	items := make([]*dtc.DtcPool, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSDtcPoolToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListDtcPoolResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func (s *dtcPoolService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcPool, *http.Response, string, error) {
	req := s.uddiClient.DNSTrafficControlAPI.PoolAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.DtcPoolFilterFieldMap[core.BackendUDDI])
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
	items := make([]*dtc.DtcPool, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIDtcPoolToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapNIOSDtcPoolToResponse(r *niosdtc.DtcPool) *dtc.DtcPool {
	resp := &dtc.DtcPool{
		Id: r.Ref,
	}
	resp.NIOS = &dtc.NIOSDtcPoolExt{
		AutoConsolidatedMonitors: r.AutoConsolidatedMonitors,
		Availability:             r.Availability,
		Comment:                  r.Comment,
		ConsolidatedMonitors:     r.ConsolidatedMonitors,
		Disable:                  r.Disable,
		Health:                   r.Health,
		LbAlternateMethod:        r.LbAlternateMethod,
		LbAlternateTopology:      r.LbAlternateTopology,
		LbDynamicRatioAlternate:  r.LbDynamicRatioAlternate,
		LbDynamicRatioPreferred:  r.LbDynamicRatioPreferred,
		LbPreferredMethod:        r.LbPreferredMethod,
		LbPreferredTopology:      r.LbPreferredTopology,
		Monitors:                 r.Monitors,
		Name:                     r.Name,
		Quorum:                   r.Quorum,
		Servers:                  r.Servers,
		Ttl:                      r.Ttl,
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

func mapUDDIDtcPoolToResponse(r *uddidtc.Pool) *dtc.DtcPool {
	resp := &dtc.DtcPool{
		Id: r.Id,
	}
	resp.UDDI = &dtc.UDDIDtcPoolExt{
		Comment:                   r.Comment,
		ConsolidatedHealthEnabled: r.ConsolidatedHealthEnabled,
		Disabled:                  r.Disabled,
		HealthChecks:              r.HealthChecks,
		InheritanceSources:        r.InheritanceSources,
		Metadata:                  r.Metadata,
		Method:                    r.Method,
		Name:                      r.Name,
		PoolAvailability:          r.PoolAvailability,
		PoolServersQuorum:         r.PoolServersQuorum,
		ServerAvailability:        r.ServerAvailability,
		ServerHealthChecksQuorum:  r.ServerHealthChecksQuorum,
		Servers:                   r.Servers,
		Ttl:                       r.Ttl,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
