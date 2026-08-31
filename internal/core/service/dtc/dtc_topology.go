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

type DtcTopologyService interface {
	Create(ctx context.Context, obj *dtc.DtcTopology, opts *core.Options) (*dtc.DtcTopology, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dtc.DtcTopology, *http.Response, error)
	Update(ctx context.Context, id string, obj *dtc.DtcTopology, opts *core.Options) (*dtc.DtcTopology, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcTopology, *http.Response, string, error)
}

type dtcTopologyService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewDtcTopologyService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) DtcTopologyService {
	return &dtcTopologyService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new DtcTopology and returns the created object
func (s *dtcTopologyService) Create(ctx context.Context, obj *dtc.DtcTopology, opts *core.Options) (*dtc.DtcTopology, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcTopologyService) createNIOS(ctx context.Context, obj *dtc.DtcTopology, opts *core.Options) (*dtc.DtcTopology, *http.Response, error) {
	payload, err := common.MapTo[niosdtc.DtcTopology](obj, mapper.DtcTopologyNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DTCAPI.DtcTopologyAPI.
		Create(ctx).
		DtcTopology(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateDtcTopologyResponseAsObject.GetResult()

	return mapNIOSDtcTopologyToResponse(&result), httpResp, nil
}

// Read retrieves a DtcTopology by ID
func (s *dtcTopologyService) Read(ctx context.Context, id string, opts *core.Options) (*dtc.DtcTopology, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcTopologyService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dtc.DtcTopology, *http.Response, error) {
	req := s.niosClient.DTCAPI.DtcTopologyAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetDtcTopologyResponseObjectAsResult.GetResult()

	return mapNIOSDtcTopologyToResponse(&result), httpResp, nil
}

// Update modifies an existing DtcTopology and returns the updated object
func (s *dtcTopologyService) Update(ctx context.Context, id string, obj *dtc.DtcTopology, opts *core.Options) (*dtc.DtcTopology, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcTopologyService) updateNIOS(ctx context.Context, id string, obj *dtc.DtcTopology, opts *core.Options) (*dtc.DtcTopology, *http.Response, error) {
	payload, err := common.MapTo[niosdtc.DtcTopology](obj, mapper.DtcTopologyNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DTCAPI.DtcTopologyAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		DtcTopology(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateDtcTopologyResponseAsObject.GetResult()

	return mapNIOSDtcTopologyToResponse(&result), httpResp, nil
}

// Delete removes a DtcTopology by ID
func (s *dtcTopologyService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcTopologyService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DTCAPI.DtcTopologyAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves DtcTopology objects based on filter options
func (s *dtcTopologyService) List(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcTopology, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dtcTopologyService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dtc.DtcTopology, *http.Response, string, error) {
	req := s.niosClient.DTCAPI.DtcTopologyAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.DtcTopologyFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListDtcTopologyResponseObject.GetResult()
	items := make([]*dtc.DtcTopology, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSDtcTopologyToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListDtcTopologyResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSDtcTopologyToResponse(r *niosdtc.DtcTopology) *dtc.DtcTopology {
	resp := &dtc.DtcTopology{
		Id: r.Ref,
	}
	resp.NIOS = &dtc.NIOSDtcTopologyExt{
		Comment: r.Comment,
		Name:    r.Name,
		Rules:   r.Rules,
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
