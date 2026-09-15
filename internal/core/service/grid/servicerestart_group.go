package grid

import (
	"context"
	"fmt"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niosgrid "github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/grid"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/grid"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
)

type ServicerestartGroupService interface {
	Create(ctx context.Context, obj *grid.ServicerestartGroup, opts *core.Options) (*grid.ServicerestartGroup, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*grid.ServicerestartGroup, *http.Response, error)
	Update(ctx context.Context, id string, obj *grid.ServicerestartGroup, opts *core.Options) (*grid.ServicerestartGroup, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*grid.ServicerestartGroup, *http.Response, string, error)
}

type servicerestartGroupService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewServicerestartGroupService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) ServicerestartGroupService {
	return &servicerestartGroupService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new ServicerestartGroup and returns the created object
func (s *servicerestartGroupService) Create(ctx context.Context, obj *grid.ServicerestartGroup, opts *core.Options) (*grid.ServicerestartGroup, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *servicerestartGroupService) createNIOS(ctx context.Context, obj *grid.ServicerestartGroup, opts *core.Options) (*grid.ServicerestartGroup, *http.Response, error) {
	payload, err := common.MapTo[niosgrid.GridServicerestartGroup](obj, mapper.ServicerestartGroupNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.GridAPI.GridServicerestartGroupAPI.
		Create(ctx).
		GridServicerestartGroup(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateGridServicerestartGroupResponseAsObject.GetResult()

	return mapNIOSServicerestartGroupToResponse(&result), httpResp, nil
}

// Read retrieves a ServicerestartGroup by ID
func (s *servicerestartGroupService) Read(ctx context.Context, id string, opts *core.Options) (*grid.ServicerestartGroup, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *servicerestartGroupService) readNIOS(ctx context.Context, id string, opts *core.Options) (*grid.ServicerestartGroup, *http.Response, error) {
	req := s.niosClient.GridAPI.GridServicerestartGroupAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetGridServicerestartGroupResponseObjectAsResult.GetResult()

	return mapNIOSServicerestartGroupToResponse(&result), httpResp, nil
}

// Update modifies an existing ServicerestartGroup and returns the updated object
func (s *servicerestartGroupService) Update(ctx context.Context, id string, obj *grid.ServicerestartGroup, opts *core.Options) (*grid.ServicerestartGroup, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *servicerestartGroupService) updateNIOS(ctx context.Context, id string, obj *grid.ServicerestartGroup, opts *core.Options) (*grid.ServicerestartGroup, *http.Response, error) {
	payload, err := common.MapTo[niosgrid.GridServicerestartGroup](obj, mapper.ServicerestartGroupNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.GridAPI.GridServicerestartGroupAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		GridServicerestartGroup(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateGridServicerestartGroupResponseAsObject.GetResult()

	return mapNIOSServicerestartGroupToResponse(&result), httpResp, nil
}

// Delete removes a ServicerestartGroup by ID
func (s *servicerestartGroupService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *servicerestartGroupService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.GridAPI.GridServicerestartGroupAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves ServicerestartGroup objects based on filter options
func (s *servicerestartGroupService) List(ctx context.Context, opts *core.ListOptions) ([]*grid.ServicerestartGroup, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *servicerestartGroupService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*grid.ServicerestartGroup, *http.Response, string, error) {
	req := s.niosClient.GridAPI.GridServicerestartGroupAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.ServicerestartGroupFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListGridServicerestartGroupResponseObject.GetResult()
	items := make([]*grid.ServicerestartGroup, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSServicerestartGroupToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListGridServicerestartGroupResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSServicerestartGroupToResponse(r *niosgrid.GridServicerestartGroup) *grid.ServicerestartGroup {
	resp := &grid.ServicerestartGroup{
		Id: r.Ref,
	}
	resp.NIOS = &grid.NIOSServicerestartGroupExt{
		Comment:           r.Comment,
		Members:           r.Members,
		Mode:              r.Mode,
		Name:              r.Name,
		RecurringSchedule: r.RecurringSchedule,
		Service:           r.Service,
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
