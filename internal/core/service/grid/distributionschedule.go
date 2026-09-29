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

type DistributionscheduleService interface {
	Create(ctx context.Context, obj *grid.Distributionschedule, opts *core.Options) (*grid.Distributionschedule, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*grid.Distributionschedule, *http.Response, error)
	Update(ctx context.Context, id string, obj *grid.Distributionschedule, opts *core.Options) (*grid.Distributionschedule, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*grid.Distributionschedule, *http.Response, string, error)
}

type distributionscheduleService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewDistributionscheduleService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) DistributionscheduleService {
	return &distributionscheduleService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new Distributionschedule and returns the created object
func (s *distributionscheduleService) Create(ctx context.Context, obj *grid.Distributionschedule, opts *core.Options) (*grid.Distributionschedule, *http.Response, error) {
	switch s.backend {
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

// Read retrieves a Distributionschedule by ID
func (s *distributionscheduleService) Read(ctx context.Context, id string, opts *core.Options) (*grid.Distributionschedule, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *distributionscheduleService) readNIOS(ctx context.Context, id string, opts *core.Options) (*grid.Distributionschedule, *http.Response, error) {
	req := s.niosClient.GridAPI.DistributionscheduleAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetDistributionscheduleResponseObjectAsResult.GetResult()

	return mapNIOSDistributionscheduleToResponse(&result), httpResp, nil
}

// Update modifies an existing Distributionschedule and returns the updated object
func (s *distributionscheduleService) Update(ctx context.Context, id string, obj *grid.Distributionschedule, opts *core.Options) (*grid.Distributionschedule, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *distributionscheduleService) updateNIOS(ctx context.Context, id string, obj *grid.Distributionschedule, opts *core.Options) (*grid.Distributionschedule, *http.Response, error) {
	payload, err := common.MapTo[niosgrid.Distributionschedule](obj, mapper.DistributionscheduleNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.GridAPI.DistributionscheduleAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Distributionschedule(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateDistributionscheduleResponseAsObject.GetResult()

	return mapNIOSDistributionscheduleToResponse(&result), httpResp, nil
}

// Delete removes a Distributionschedule by ID
func (s *distributionscheduleService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

// List retrieves Distributionschedule objects based on filter options
func (s *distributionscheduleService) List(ctx context.Context, opts *core.ListOptions) ([]*grid.Distributionschedule, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *distributionscheduleService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*grid.Distributionschedule, *http.Response, string, error) {
	req := s.niosClient.GridAPI.DistributionscheduleAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.DistributionscheduleFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListDistributionscheduleResponseObject.GetResult()
	items := make([]*grid.Distributionschedule, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSDistributionscheduleToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListDistributionscheduleResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSDistributionscheduleToResponse(r *niosgrid.Distributionschedule) *grid.Distributionschedule {
	resp := &grid.Distributionschedule{
		Id: r.Ref,
	}
	resp.NIOS = &grid.NIOSDistributionscheduleExt{
		Active:        r.Active,
		StartTime:     r.StartTime,
		TimeZone:      r.TimeZone,
		UpgradeGroups: r.UpgradeGroups,
	}
	return resp
}
