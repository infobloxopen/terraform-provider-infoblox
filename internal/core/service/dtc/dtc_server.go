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

type DtcServerService interface {
	Create(ctx context.Context, obj *dtc.DtcServer, opts *core.Options) (*dtc.DtcServer, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dtc.DtcServer, *http.Response, error)
	Update(ctx context.Context, id string, obj *dtc.DtcServer, opts *core.Options) (*dtc.DtcServer, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcServer, *http.Response, string, error)
}

type dtcServerService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
	uddiClient *uddiclient.APIClient
}

func NewDtcServerService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) DtcServerService {
	return &dtcServerService{
		backend:    backend,
		niosClient: nios,
		uddiClient: uddi,
	}
}

// Create creates a new DtcServer and returns the created object
func (s *dtcServerService) Create(ctx context.Context, obj *dtc.DtcServer, opts *core.Options) (*dtc.DtcServer, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcServerService) createNIOS(ctx context.Context, obj *dtc.DtcServer, opts *core.Options) (*dtc.DtcServer, *http.Response, error) {
	payload, err := common.MapTo[niosdtc.DtcServer](obj, mapper.DtcServerNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DTCAPI.DtcServerAPI.
		Create(ctx).
		DtcServer(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateDtcServerResponseAsObject.GetResult()

	return mapNIOSDtcServerToResponse(&result), httpResp, nil
}

func (s *dtcServerService) createUDDI(ctx context.Context, obj *dtc.DtcServer, opts *core.Options) (*dtc.DtcServer, *http.Response, error) {
	payload, err := common.MapTo[uddidtc.Server](obj, mapper.DtcServerUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSTrafficControlAPI.ServerAPI.
		Create(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIDtcServerToResponse(&result), httpResp, nil
}

// Read retrieves a DtcServer by ID
func (s *dtcServerService) Read(ctx context.Context, id string, opts *core.Options) (*dtc.DtcServer, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcServerService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dtc.DtcServer, *http.Response, error) {
	req := s.niosClient.DTCAPI.DtcServerAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetDtcServerResponseObjectAsResult.GetResult()

	return mapNIOSDtcServerToResponse(&result), httpResp, nil
}

func (s *dtcServerService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dtc.DtcServer, *http.Response, error) {
	req := s.uddiClient.DNSTrafficControlAPI.ServerAPI.
		Read(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIDtcServerToResponse(&result), httpResp, nil
}

// Update modifies an existing DtcServer and returns the updated object
func (s *dtcServerService) Update(ctx context.Context, id string, obj *dtc.DtcServer, opts *core.Options) (*dtc.DtcServer, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcServerService) updateNIOS(ctx context.Context, id string, obj *dtc.DtcServer, opts *core.Options) (*dtc.DtcServer, *http.Response, error) {
	payload, err := common.MapTo[niosdtc.DtcServer](obj, mapper.DtcServerNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DTCAPI.DtcServerAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		DtcServer(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateDtcServerResponseAsObject.GetResult()

	return mapNIOSDtcServerToResponse(&result), httpResp, nil
}

func (s *dtcServerService) updateUDDI(ctx context.Context, id string, obj *dtc.DtcServer, opts *core.Options) (*dtc.DtcServer, *http.Response, error) {
	payload, err := common.MapTo[uddidtc.Server](obj, mapper.DtcServerUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSTrafficControlAPI.ServerAPI.
		Update(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIDtcServerToResponse(&result), httpResp, nil
}

// Delete removes a DtcServer by ID
func (s *dtcServerService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcServerService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DTCAPI.DtcServerAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

func (s *dtcServerService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.DNSTrafficControlAPI.ServerAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves DtcServer objects based on filter options
func (s *dtcServerService) List(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcServer, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcServerService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcServer, *http.Response, string, error) {
	req := s.niosClient.DTCAPI.DtcServerAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.DtcServerFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListDtcServerResponseObject.GetResult()
	items := make([]*dtc.DtcServer, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSDtcServerToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListDtcServerResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func (s *dtcServerService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcServer, *http.Response, string, error) {
	req := s.uddiClient.DNSTrafficControlAPI.ServerAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.DtcServerFilterFieldMap[core.BackendUDDI])
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
	items := make([]*dtc.DtcServer, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIDtcServerToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapNIOSDtcServerToResponse(r *niosdtc.DtcServer) *dtc.DtcServer {
	resp := &dtc.DtcServer{
		Id: r.Ref,
	}
	resp.NIOS = &dtc.NIOSDtcServerExt{
		AutoCreateHostRecord: r.AutoCreateHostRecord,
		Comment:              r.Comment,
		Disable:              r.Disable,
		Health:               r.Health,
		Host:                 r.Host,
		Monitors:             r.Monitors,
		Name:                 r.Name,
		SniHostname:          r.SniHostname,
		UseSniHostname:       r.UseSniHostname,
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

func mapUDDIDtcServerToResponse(r *uddidtc.Server) *dtc.DtcServer {
	resp := &dtc.DtcServer{
		Id: r.Id,
	}
	resp.UDDI = &dtc.UDDIDtcServerExt{
		Address:                   r.Address,
		AutoCreateResponseRecords: r.AutoCreateResponseRecords,
		Comment:                   r.Comment,
		Disabled:                  r.Disabled,
		EndpointType:              r.EndpointType,
		Fqdn:                      r.Fqdn,
		Metadata:                  r.Metadata,
		Name:                      r.Name,
		Records:                   r.Records,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
