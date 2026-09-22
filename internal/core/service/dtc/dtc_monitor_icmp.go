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

type DtcMonitorIcmpService interface {
	Create(ctx context.Context, obj *dtc.DtcMonitorIcmp, opts *core.Options) (*dtc.DtcMonitorIcmp, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dtc.DtcMonitorIcmp, *http.Response, error)
	Update(ctx context.Context, id string, obj *dtc.DtcMonitorIcmp, opts *core.Options) (*dtc.DtcMonitorIcmp, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcMonitorIcmp, *http.Response, string, error)
}

type dtcMonitorIcmpService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
	uddiClient *uddiclient.APIClient
}

func NewDtcMonitorIcmpService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) DtcMonitorIcmpService {
	return &dtcMonitorIcmpService{
		backend:    backend,
		niosClient: nios,
		uddiClient: uddi,
	}
}

// Create creates a new DtcMonitorIcmp and returns the created object
func (s *dtcMonitorIcmpService) Create(ctx context.Context, obj *dtc.DtcMonitorIcmp, opts *core.Options) (*dtc.DtcMonitorIcmp, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcMonitorIcmpService) createNIOS(ctx context.Context, obj *dtc.DtcMonitorIcmp, opts *core.Options) (*dtc.DtcMonitorIcmp, *http.Response, error) {
	payload, err := common.MapTo[niosdtc.DtcMonitorIcmp](obj, mapper.DtcMonitorIcmpNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DTCAPI.DtcMonitorIcmpAPI.
		Create(ctx).
		DtcMonitorIcmp(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateDtcMonitorIcmpResponseAsObject.GetResult()

	return mapNIOSDtcMonitorIcmpToResponse(&result), httpResp, nil
}

func (s *dtcMonitorIcmpService) createUDDI(ctx context.Context, obj *dtc.DtcMonitorIcmp, opts *core.Options) (*dtc.DtcMonitorIcmp, *http.Response, error) {
	payload, err := common.MapTo[uddidtc.ICMPHealthCheck](obj, mapper.DtcMonitorIcmpUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSTrafficControlAPI.HealthCheckIcmpAPI.
		Create(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIDtcMonitorIcmpToResponse(&result), httpResp, nil
}

// Read retrieves a DtcMonitorIcmp by ID
func (s *dtcMonitorIcmpService) Read(ctx context.Context, id string, opts *core.Options) (*dtc.DtcMonitorIcmp, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcMonitorIcmpService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dtc.DtcMonitorIcmp, *http.Response, error) {
	req := s.niosClient.DTCAPI.DtcMonitorIcmpAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetDtcMonitorIcmpResponseObjectAsResult.GetResult()

	return mapNIOSDtcMonitorIcmpToResponse(&result), httpResp, nil
}

func (s *dtcMonitorIcmpService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dtc.DtcMonitorIcmp, *http.Response, error) {
	req := s.uddiClient.DNSTrafficControlAPI.HealthCheckIcmpAPI.
		Read(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIDtcMonitorIcmpToResponse(&result), httpResp, nil
}

// Update modifies an existing DtcMonitorIcmp and returns the updated object
func (s *dtcMonitorIcmpService) Update(ctx context.Context, id string, obj *dtc.DtcMonitorIcmp, opts *core.Options) (*dtc.DtcMonitorIcmp, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcMonitorIcmpService) updateNIOS(ctx context.Context, id string, obj *dtc.DtcMonitorIcmp, opts *core.Options) (*dtc.DtcMonitorIcmp, *http.Response, error) {
	payload, err := common.MapTo[niosdtc.DtcMonitorIcmp](obj, mapper.DtcMonitorIcmpNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DTCAPI.DtcMonitorIcmpAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		DtcMonitorIcmp(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateDtcMonitorIcmpResponseAsObject.GetResult()

	return mapNIOSDtcMonitorIcmpToResponse(&result), httpResp, nil
}

func (s *dtcMonitorIcmpService) updateUDDI(ctx context.Context, id string, obj *dtc.DtcMonitorIcmp, opts *core.Options) (*dtc.DtcMonitorIcmp, *http.Response, error) {
	payload, err := common.MapTo[uddidtc.ICMPHealthCheck](obj, mapper.DtcMonitorIcmpUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSTrafficControlAPI.HealthCheckIcmpAPI.
		Update(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIDtcMonitorIcmpToResponse(&result), httpResp, nil
}

// Delete removes a DtcMonitorIcmp by ID
func (s *dtcMonitorIcmpService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcMonitorIcmpService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DTCAPI.DtcMonitorIcmpAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

func (s *dtcMonitorIcmpService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.DNSTrafficControlAPI.HealthCheckIcmpAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves DtcMonitorIcmp objects based on filter options
func (s *dtcMonitorIcmpService) List(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcMonitorIcmp, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcMonitorIcmpService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcMonitorIcmp, *http.Response, string, error) {
	req := s.niosClient.DTCAPI.DtcMonitorIcmpAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.DtcMonitorIcmpFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListDtcMonitorIcmpResponseObject.GetResult()
	items := make([]*dtc.DtcMonitorIcmp, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSDtcMonitorIcmpToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListDtcMonitorIcmpResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func (s *dtcMonitorIcmpService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcMonitorIcmp, *http.Response, string, error) {
	req := s.uddiClient.DNSTrafficControlAPI.HealthCheckIcmpAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.DtcMonitorIcmpFilterFieldMap[core.BackendUDDI])
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
	items := make([]*dtc.DtcMonitorIcmp, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIDtcMonitorIcmpToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapNIOSDtcMonitorIcmpToResponse(r *niosdtc.DtcMonitorIcmp) *dtc.DtcMonitorIcmp {
	resp := &dtc.DtcMonitorIcmp{
		Id: r.Ref,
	}
	resp.NIOS = &dtc.NIOSDtcMonitorIcmpExt{
		Comment:   r.Comment,
		Interval:  r.Interval,
		Name:      r.Name,
		RetryDown: r.RetryDown,
		RetryUp:   r.RetryUp,
		Timeout:   r.Timeout,
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

func mapUDDIDtcMonitorIcmpToResponse(r *uddidtc.ICMPHealthCheck) *dtc.DtcMonitorIcmp {
	resp := &dtc.DtcMonitorIcmp{
		Id: r.Id,
	}
	resp.UDDI = &dtc.UDDIDtcMonitorIcmpExt{
		Comment:   r.Comment,
		Disabled:  r.Disabled,
		Interval:  r.Interval,
		Metadata:  r.Metadata,
		Name:      r.Name,
		RetryDown: r.RetryDown,
		RetryUp:   r.RetryUp,
		Timeout:   r.Timeout,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
