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

type DtcMonitorHttpService interface {
	Create(ctx context.Context, obj *dtc.DtcMonitorHttp, opts *core.Options) (*dtc.DtcMonitorHttp, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dtc.DtcMonitorHttp, *http.Response, error)
	Update(ctx context.Context, id string, obj *dtc.DtcMonitorHttp, opts *core.Options) (*dtc.DtcMonitorHttp, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcMonitorHttp, *http.Response, string, error)
}

type dtcMonitorHttpService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
	uddiClient *uddiclient.APIClient
}

func NewDtcMonitorHttpService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) DtcMonitorHttpService {
	return &dtcMonitorHttpService{
		backend:    backend,
		niosClient: nios,
		uddiClient: uddi,
	}
}

// Create creates a new DtcMonitorHttp and returns the created object
func (s *dtcMonitorHttpService) Create(ctx context.Context, obj *dtc.DtcMonitorHttp, opts *core.Options) (*dtc.DtcMonitorHttp, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcMonitorHttpService) createNIOS(ctx context.Context, obj *dtc.DtcMonitorHttp, opts *core.Options) (*dtc.DtcMonitorHttp, *http.Response, error) {
	payload, err := common.MapTo[niosdtc.DtcMonitorHttp](obj, mapper.DtcMonitorHttpNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DTCAPI.DtcMonitorHttpAPI.
		Create(ctx).
		DtcMonitorHttp(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateDtcMonitorHttpResponseAsObject.GetResult()

	return mapNIOSDtcMonitorHttpToResponse(&result), httpResp, nil
}

func (s *dtcMonitorHttpService) createUDDI(ctx context.Context, obj *dtc.DtcMonitorHttp, opts *core.Options) (*dtc.DtcMonitorHttp, *http.Response, error) {
	payload, err := common.MapTo[uddidtc.HTTPHealthCheck](obj, mapper.DtcMonitorHttpUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSTrafficControlAPI.HealthCheckHttpAPI.
		Create(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIDtcMonitorHttpToResponse(&result), httpResp, nil
}

// Read retrieves a DtcMonitorHttp by ID
func (s *dtcMonitorHttpService) Read(ctx context.Context, id string, opts *core.Options) (*dtc.DtcMonitorHttp, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcMonitorHttpService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dtc.DtcMonitorHttp, *http.Response, error) {
	req := s.niosClient.DTCAPI.DtcMonitorHttpAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetDtcMonitorHttpResponseObjectAsResult.GetResult()

	return mapNIOSDtcMonitorHttpToResponse(&result), httpResp, nil
}

func (s *dtcMonitorHttpService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dtc.DtcMonitorHttp, *http.Response, error) {
	req := s.uddiClient.DNSTrafficControlAPI.HealthCheckHttpAPI.
		Read(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIDtcMonitorHttpToResponse(&result), httpResp, nil
}

// Update modifies an existing DtcMonitorHttp and returns the updated object
func (s *dtcMonitorHttpService) Update(ctx context.Context, id string, obj *dtc.DtcMonitorHttp, opts *core.Options) (*dtc.DtcMonitorHttp, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcMonitorHttpService) updateNIOS(ctx context.Context, id string, obj *dtc.DtcMonitorHttp, opts *core.Options) (*dtc.DtcMonitorHttp, *http.Response, error) {
	payload, err := common.MapTo[niosdtc.DtcMonitorHttp](obj, mapper.DtcMonitorHttpNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DTCAPI.DtcMonitorHttpAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		DtcMonitorHttp(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateDtcMonitorHttpResponseAsObject.GetResult()

	return mapNIOSDtcMonitorHttpToResponse(&result), httpResp, nil
}

func (s *dtcMonitorHttpService) updateUDDI(ctx context.Context, id string, obj *dtc.DtcMonitorHttp, opts *core.Options) (*dtc.DtcMonitorHttp, *http.Response, error) {
	payload, err := common.MapTo[uddidtc.HTTPHealthCheck](obj, mapper.DtcMonitorHttpUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSTrafficControlAPI.HealthCheckHttpAPI.
		Update(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIDtcMonitorHttpToResponse(&result), httpResp, nil
}

// Delete removes a DtcMonitorHttp by ID
func (s *dtcMonitorHttpService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcMonitorHttpService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DTCAPI.DtcMonitorHttpAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

func (s *dtcMonitorHttpService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.DNSTrafficControlAPI.HealthCheckHttpAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves DtcMonitorHttp objects based on filter options
func (s *dtcMonitorHttpService) List(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcMonitorHttp, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcMonitorHttpService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcMonitorHttp, *http.Response, string, error) {
	req := s.niosClient.DTCAPI.DtcMonitorHttpAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.DtcMonitorHttpFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListDtcMonitorHttpResponseObject.GetResult()
	items := make([]*dtc.DtcMonitorHttp, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSDtcMonitorHttpToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListDtcMonitorHttpResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func (s *dtcMonitorHttpService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcMonitorHttp, *http.Response, string, error) {
	req := s.uddiClient.DNSTrafficControlAPI.HealthCheckHttpAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.DtcMonitorHttpFilterFieldMap[core.BackendUDDI])
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
	items := make([]*dtc.DtcMonitorHttp, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIDtcMonitorHttpToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapNIOSDtcMonitorHttpToResponse(r *niosdtc.DtcMonitorHttp) *dtc.DtcMonitorHttp {
	resp := &dtc.DtcMonitorHttp{
		Id: r.Ref,
	}
	resp.NIOS = &dtc.NIOSDtcMonitorHttpExt{
		Ciphers:             r.Ciphers,
		ClientCert:          r.ClientCert,
		Comment:             r.Comment,
		ContentCheck:        r.ContentCheck,
		ContentCheckInput:   r.ContentCheckInput,
		ContentCheckOp:      r.ContentCheckOp,
		ContentCheckRegex:   r.ContentCheckRegex,
		ContentExtractGroup: r.ContentExtractGroup,
		ContentExtractType:  r.ContentExtractType,
		ContentExtractValue: r.ContentExtractValue,
		EnableSni:           r.EnableSni,
		Interval:            r.Interval,
		Name:                r.Name,
		Port:                r.Port,
		Request:             r.Request,
		Result:              r.Result,
		ResultCode:          r.ResultCode,
		RetryDown:           r.RetryDown,
		RetryUp:             r.RetryUp,
		Secure:              r.Secure,
		Timeout:             r.Timeout,
		ValidateCert:        r.ValidateCert,
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

func mapUDDIDtcMonitorHttpToResponse(r *uddidtc.HTTPHealthCheck) *dtc.DtcMonitorHttp {
	resp := &dtc.DtcMonitorHttp{
		Id: r.Id,
	}
	resp.UDDI = &dtc.UDDIDtcMonitorHttpExt{
		CheckResponseBody:           r.CheckResponseBody,
		CheckResponseBodyNegative:   r.CheckResponseBodyNegative,
		CheckResponseBodyRegex:      r.CheckResponseBodyRegex,
		CheckResponseHeader:         r.CheckResponseHeader,
		CheckResponseHeaderNegative: r.CheckResponseHeaderNegative,
		CheckResponseHeaderRegexes:  r.CheckResponseHeaderRegexes,
		Codes:                       r.Codes,
		Comment:                     r.Comment,
		Disabled:                    r.Disabled,
		Https:                       r.Https,
		Interval:                    r.Interval,
		Metadata:                    r.Metadata,
		Name:                        r.Name,
		Port:                        r.Port,
		Request:                     r.Request,
		RetryDown:                   r.RetryDown,
		RetryUp:                     r.RetryUp,
		Timeout:                     r.Timeout,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
