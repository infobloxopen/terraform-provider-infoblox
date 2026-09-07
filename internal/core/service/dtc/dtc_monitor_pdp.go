package dtc

import (
	"context"
	"fmt"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niosdtc "github.com/infobloxopen/infoblox-nios-go-client/dtc"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/dtc"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dtc"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
)

type DtcMonitorPdpService interface {
	Create(ctx context.Context, obj *dtc.DtcMonitorPdp, opts *core.Options) (*dtc.DtcMonitorPdp, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dtc.DtcMonitorPdp, *http.Response, error)
	Update(ctx context.Context, id string, obj *dtc.DtcMonitorPdp, opts *core.Options) (*dtc.DtcMonitorPdp, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcMonitorPdp, *http.Response, string, error)
}

type dtcMonitorPdpService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewDtcMonitorPdpService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) DtcMonitorPdpService {
	return &dtcMonitorPdpService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new DtcMonitorPdp and returns the created object
func (s *dtcMonitorPdpService) Create(ctx context.Context, obj *dtc.DtcMonitorPdp, opts *core.Options) (*dtc.DtcMonitorPdp, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcMonitorPdpService) createNIOS(ctx context.Context, obj *dtc.DtcMonitorPdp, opts *core.Options) (*dtc.DtcMonitorPdp, *http.Response, error) {
	payload, err := common.MapTo[niosdtc.DtcMonitorPdp](obj, mapper.DtcMonitorPdpNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DTCAPI.DtcMonitorPdpAPI.
		Create(ctx).
		DtcMonitorPdp(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateDtcMonitorPdpResponseAsObject.GetResult()

	return mapNIOSDtcMonitorPdpToResponse(&result), httpResp, nil
}

// Read retrieves a DtcMonitorPdp by ID
func (s *dtcMonitorPdpService) Read(ctx context.Context, id string, opts *core.Options) (*dtc.DtcMonitorPdp, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcMonitorPdpService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dtc.DtcMonitorPdp, *http.Response, error) {
	req := s.niosClient.DTCAPI.DtcMonitorPdpAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetDtcMonitorPdpResponseObjectAsResult.GetResult()

	return mapNIOSDtcMonitorPdpToResponse(&result), httpResp, nil
}

// Update modifies an existing DtcMonitorPdp and returns the updated object
func (s *dtcMonitorPdpService) Update(ctx context.Context, id string, obj *dtc.DtcMonitorPdp, opts *core.Options) (*dtc.DtcMonitorPdp, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcMonitorPdpService) updateNIOS(ctx context.Context, id string, obj *dtc.DtcMonitorPdp, opts *core.Options) (*dtc.DtcMonitorPdp, *http.Response, error) {
	payload, err := common.MapTo[niosdtc.DtcMonitorPdp](obj, mapper.DtcMonitorPdpNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DTCAPI.DtcMonitorPdpAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		DtcMonitorPdp(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateDtcMonitorPdpResponseAsObject.GetResult()

	return mapNIOSDtcMonitorPdpToResponse(&result), httpResp, nil
}

// Delete removes a DtcMonitorPdp by ID
func (s *dtcMonitorPdpService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcMonitorPdpService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DTCAPI.DtcMonitorPdpAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves DtcMonitorPdp objects based on filter options
func (s *dtcMonitorPdpService) List(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcMonitorPdp, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcMonitorPdpService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcMonitorPdp, *http.Response, string, error) {
	req := s.niosClient.DTCAPI.DtcMonitorPdpAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.DtcMonitorPdpFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListDtcMonitorPdpResponseObject.GetResult()
	items := make([]*dtc.DtcMonitorPdp, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSDtcMonitorPdpToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListDtcMonitorPdpResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSDtcMonitorPdpToResponse(r *niosdtc.DtcMonitorPdp) *dtc.DtcMonitorPdp {
	resp := &dtc.DtcMonitorPdp{
		Id: r.Ref,
	}
	resp.NIOS = &dtc.NIOSDtcMonitorPdpExt{
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
