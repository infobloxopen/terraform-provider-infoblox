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

type DtcMonitorTcpService interface {
	Create(ctx context.Context, obj *dtc.DtcMonitorTcp, opts *core.Options) (*dtc.DtcMonitorTcp, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dtc.DtcMonitorTcp, *http.Response, error)
	Update(ctx context.Context, id string, obj *dtc.DtcMonitorTcp, opts *core.Options) (*dtc.DtcMonitorTcp, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcMonitorTcp, *http.Response, string, error)
}

type dtcMonitorTcpService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
	uddiClient *uddiclient.APIClient
}

func NewDtcMonitorTcpService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) DtcMonitorTcpService {
	return &dtcMonitorTcpService{
		backend:    backend,
		niosClient: nios,
		uddiClient: uddi,
	}
}

// Create creates a new DtcMonitorTcp and returns the created object
func (s *dtcMonitorTcpService) Create(ctx context.Context, obj *dtc.DtcMonitorTcp, opts *core.Options) (*dtc.DtcMonitorTcp, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcMonitorTcpService) createNIOS(ctx context.Context, obj *dtc.DtcMonitorTcp, opts *core.Options) (*dtc.DtcMonitorTcp, *http.Response, error) {
	payload, err := common.MapTo[niosdtc.DtcMonitorTcp](obj, mapper.DtcMonitorTcpNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DTCAPI.DtcMonitorTcpAPI.
		Create(ctx).
		DtcMonitorTcp(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateDtcMonitorTcpResponseAsObject.GetResult()

	return mapNIOSDtcMonitorTcpToResponse(&result), httpResp, nil
}

func (s *dtcMonitorTcpService) createUDDI(ctx context.Context, obj *dtc.DtcMonitorTcp, opts *core.Options) (*dtc.DtcMonitorTcp, *http.Response, error) {
	payload, err := common.MapTo[uddidtc.TCPHealthCheck](obj, mapper.DtcMonitorTcpUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSTrafficControlAPI.HealthCheckTcpAPI.
		Create(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIDtcMonitorTcpToResponse(&result), httpResp, nil
}

// Read retrieves a DtcMonitorTcp by ID
func (s *dtcMonitorTcpService) Read(ctx context.Context, id string, opts *core.Options) (*dtc.DtcMonitorTcp, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcMonitorTcpService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dtc.DtcMonitorTcp, *http.Response, error) {
	req := s.niosClient.DTCAPI.DtcMonitorTcpAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetDtcMonitorTcpResponseObjectAsResult.GetResult()

	return mapNIOSDtcMonitorTcpToResponse(&result), httpResp, nil
}

func (s *dtcMonitorTcpService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dtc.DtcMonitorTcp, *http.Response, error) {
	req := s.uddiClient.DNSTrafficControlAPI.HealthCheckTcpAPI.
		Read(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIDtcMonitorTcpToResponse(&result), httpResp, nil
}

// Update modifies an existing DtcMonitorTcp and returns the updated object
func (s *dtcMonitorTcpService) Update(ctx context.Context, id string, obj *dtc.DtcMonitorTcp, opts *core.Options) (*dtc.DtcMonitorTcp, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcMonitorTcpService) updateNIOS(ctx context.Context, id string, obj *dtc.DtcMonitorTcp, opts *core.Options) (*dtc.DtcMonitorTcp, *http.Response, error) {
	payload, err := common.MapTo[niosdtc.DtcMonitorTcp](obj, mapper.DtcMonitorTcpNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DTCAPI.DtcMonitorTcpAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		DtcMonitorTcp(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateDtcMonitorTcpResponseAsObject.GetResult()

	return mapNIOSDtcMonitorTcpToResponse(&result), httpResp, nil
}

func (s *dtcMonitorTcpService) updateUDDI(ctx context.Context, id string, obj *dtc.DtcMonitorTcp, opts *core.Options) (*dtc.DtcMonitorTcp, *http.Response, error) {
	payload, err := common.MapTo[uddidtc.TCPHealthCheck](obj, mapper.DtcMonitorTcpUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSTrafficControlAPI.HealthCheckTcpAPI.
		Update(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIDtcMonitorTcpToResponse(&result), httpResp, nil
}

// Delete removes a DtcMonitorTcp by ID
func (s *dtcMonitorTcpService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcMonitorTcpService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DTCAPI.DtcMonitorTcpAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

func (s *dtcMonitorTcpService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.DNSTrafficControlAPI.HealthCheckTcpAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves DtcMonitorTcp objects based on filter options
func (s *dtcMonitorTcpService) List(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcMonitorTcp, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcMonitorTcpService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcMonitorTcp, *http.Response, string, error) {
	req := s.niosClient.DTCAPI.DtcMonitorTcpAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.DtcMonitorTcpFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListDtcMonitorTcpResponseObject.GetResult()
	items := make([]*dtc.DtcMonitorTcp, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSDtcMonitorTcpToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListDtcMonitorTcpResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func (s *dtcMonitorTcpService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcMonitorTcp, *http.Response, string, error) {
	req := s.uddiClient.DNSTrafficControlAPI.HealthCheckTcpAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.DtcMonitorTcpFilterFieldMap[core.BackendUDDI])
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
	items := make([]*dtc.DtcMonitorTcp, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIDtcMonitorTcpToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapNIOSDtcMonitorTcpToResponse(r *niosdtc.DtcMonitorTcp) *dtc.DtcMonitorTcp {
	resp := &dtc.DtcMonitorTcp{
		Id: r.Ref,
	}
	resp.NIOS = &dtc.NIOSDtcMonitorTcpExt{
		Comment:   r.Comment,
		Interval:  r.Interval,
		Name:      r.Name,
		Port:      r.Port,
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

func mapUDDIDtcMonitorTcpToResponse(r *uddidtc.TCPHealthCheck) *dtc.DtcMonitorTcp {
	resp := &dtc.DtcMonitorTcp{
		Id: r.Id,
	}
	resp.UDDI = &dtc.UDDIDtcMonitorTcpExt{
		Comment:   r.Comment,
		Disabled:  r.Disabled,
		Interval:  r.Interval,
		Metadata:  r.Metadata,
		Name:      r.Name,
		Port:      r.Port,
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
