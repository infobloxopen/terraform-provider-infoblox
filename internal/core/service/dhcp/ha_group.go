package dhcp

import (
	"context"
	"fmt"
	"maps"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dhcp"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

type HaGroupService interface {
	Create(ctx context.Context, obj *dhcp.HaGroup, opts *core.Options) (*dhcp.HaGroup, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dhcp.HaGroup, *http.Response, error)
	Update(ctx context.Context, id string, obj *dhcp.HaGroup, opts *core.Options) (*dhcp.HaGroup, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.HaGroup, *http.Response, string, error)
}

type haGroupService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewHaGroupService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) HaGroupService {
	return &haGroupService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new HaGroup and returns the created object
func (s *haGroupService) Create(ctx context.Context, obj *dhcp.HaGroup, opts *core.Options) (*dhcp.HaGroup, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *haGroupService) createUDDI(ctx context.Context, obj *dhcp.HaGroup, opts *core.Options) (*dhcp.HaGroup, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.HAGroup](obj, mapper.HaGroupUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAddressManagementAPI.HaGroupAPI.
		Create(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIHaGroupToResponse(&result), httpResp, nil
}

// Read retrieves a HaGroup by ID
func (s *haGroupService) Read(ctx context.Context, id string, opts *core.Options) (*dhcp.HaGroup, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *haGroupService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dhcp.HaGroup, *http.Response, error) {
	req := s.uddiClient.IPAddressManagementAPI.HaGroupAPI.
		Read(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIHaGroupToResponse(&result), httpResp, nil
}

// Update modifies an existing HaGroup and returns the updated object
func (s *haGroupService) Update(ctx context.Context, id string, obj *dhcp.HaGroup, opts *core.Options) (*dhcp.HaGroup, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *haGroupService) updateUDDI(ctx context.Context, id string, obj *dhcp.HaGroup, opts *core.Options) (*dhcp.HaGroup, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.HAGroup](obj, mapper.HaGroupUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAddressManagementAPI.HaGroupAPI.
		Update(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIHaGroupToResponse(&result), httpResp, nil
}

// Delete removes a HaGroup by ID
func (s *haGroupService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *haGroupService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.IPAddressManagementAPI.HaGroupAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves HaGroup objects based on filter options
func (s *haGroupService) List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.HaGroup, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *haGroupService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dhcp.HaGroup, *http.Response, string, error) {
	req := s.uddiClient.IPAddressManagementAPI.HaGroupAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.HaGroupFilterFieldMap[core.BackendUDDI])
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
	items := make([]*dhcp.HaGroup, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIHaGroupToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDIHaGroupToResponse(r *uddiipam.HAGroup) *dhcp.HaGroup {
	resp := &dhcp.HaGroup{
		Id: r.Id,
	}
	resp.UDDI = &dhcp.UDDIHaGroupExt{
		AnycastConfigId: r.AnycastConfigId,
		Comment:         r.Comment,
		Hosts:           r.Hosts,
		IpSpace:         r.IpSpace,
		Mode:            r.Mode,
		Name:            r.Name,
		Status:          r.Status,
		StatusV6:        r.StatusV6,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
