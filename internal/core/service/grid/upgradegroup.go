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

type UpgradegroupService interface {
	Create(ctx context.Context, obj *grid.Upgradegroup, opts *core.Options) (*grid.Upgradegroup, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*grid.Upgradegroup, *http.Response, error)
	Update(ctx context.Context, id string, obj *grid.Upgradegroup, opts *core.Options) (*grid.Upgradegroup, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*grid.Upgradegroup, *http.Response, string, error)
}

type upgradegroupService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewUpgradegroupService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) UpgradegroupService {
	return &upgradegroupService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new Upgradegroup and returns the created object
func (s *upgradegroupService) Create(ctx context.Context, obj *grid.Upgradegroup, opts *core.Options) (*grid.Upgradegroup, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *upgradegroupService) createNIOS(ctx context.Context, obj *grid.Upgradegroup, opts *core.Options) (*grid.Upgradegroup, *http.Response, error) {
	payload, err := common.MapTo[niosgrid.Upgradegroup](obj, mapper.UpgradegroupNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.GridAPI.UpgradegroupAPI.
		Create(ctx).
		Upgradegroup(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateUpgradegroupResponseAsObject.GetResult()

	return mapNIOSUpgradegroupToResponse(&result), httpResp, nil
}

// Read retrieves a Upgradegroup by ID
func (s *upgradegroupService) Read(ctx context.Context, id string, opts *core.Options) (*grid.Upgradegroup, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *upgradegroupService) readNIOS(ctx context.Context, id string, opts *core.Options) (*grid.Upgradegroup, *http.Response, error) {
	req := s.niosClient.GridAPI.UpgradegroupAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetUpgradegroupResponseObjectAsResult.GetResult()

	return mapNIOSUpgradegroupToResponse(&result), httpResp, nil
}

// Update modifies an existing Upgradegroup and returns the updated object
func (s *upgradegroupService) Update(ctx context.Context, id string, obj *grid.Upgradegroup, opts *core.Options) (*grid.Upgradegroup, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *upgradegroupService) updateNIOS(ctx context.Context, id string, obj *grid.Upgradegroup, opts *core.Options) (*grid.Upgradegroup, *http.Response, error) {
	payload, err := common.MapTo[niosgrid.Upgradegroup](obj, mapper.UpgradegroupNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.GridAPI.UpgradegroupAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Upgradegroup(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateUpgradegroupResponseAsObject.GetResult()

	return mapNIOSUpgradegroupToResponse(&result), httpResp, nil
}

// Delete removes a Upgradegroup by ID
func (s *upgradegroupService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *upgradegroupService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.GridAPI.UpgradegroupAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves Upgradegroup objects based on filter options
func (s *upgradegroupService) List(ctx context.Context, opts *core.ListOptions) ([]*grid.Upgradegroup, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *upgradegroupService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*grid.Upgradegroup, *http.Response, string, error) {
	req := s.niosClient.GridAPI.UpgradegroupAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.UpgradegroupFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListUpgradegroupResponseObject.GetResult()
	items := make([]*grid.Upgradegroup, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSUpgradegroupToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListUpgradegroupResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSUpgradegroupToResponse(r *niosgrid.Upgradegroup) *grid.Upgradegroup {
	resp := &grid.Upgradegroup{
		Id: r.Ref,
	}
	resp.NIOS = &grid.NIOSUpgradegroupExt{
		Comment:                    r.Comment,
		DistributionDependentGroup: r.DistributionDependentGroup,
		DistributionPolicy:         r.DistributionPolicy,
		DistributionTime:           r.DistributionTime,
		Members:                    r.Members,
		Name:                       r.Name,
		UpgradeDependentGroup:      r.UpgradeDependentGroup,
		UpgradePolicy:              r.UpgradePolicy,
		UpgradeTime:                r.UpgradeTime,
	}
	return resp
}
