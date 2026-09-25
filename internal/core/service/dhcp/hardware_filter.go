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

type HardwareFilterService interface {
	Create(ctx context.Context, obj *dhcp.HardwareFilter, opts *core.Options) (*dhcp.HardwareFilter, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dhcp.HardwareFilter, *http.Response, error)
	Update(ctx context.Context, id string, obj *dhcp.HardwareFilter, opts *core.Options) (*dhcp.HardwareFilter, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.HardwareFilter, *http.Response, string, error)
}

type hardwareFilterService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewHardwareFilterService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) HardwareFilterService {
	return &hardwareFilterService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new HardwareFilter and returns the created object
func (s *hardwareFilterService) Create(ctx context.Context, obj *dhcp.HardwareFilter, opts *core.Options) (*dhcp.HardwareFilter, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *hardwareFilterService) createUDDI(ctx context.Context, obj *dhcp.HardwareFilter, opts *core.Options) (*dhcp.HardwareFilter, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.HardwareFilter](obj, mapper.HardwareFilterUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAddressManagementAPI.HardwareFilterAPI.
		Create(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIHardwareFilterToResponse(&result), httpResp, nil
}

// Read retrieves a HardwareFilter by ID
func (s *hardwareFilterService) Read(ctx context.Context, id string, opts *core.Options) (*dhcp.HardwareFilter, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *hardwareFilterService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dhcp.HardwareFilter, *http.Response, error) {
	req := s.uddiClient.IPAddressManagementAPI.HardwareFilterAPI.
		Read(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIHardwareFilterToResponse(&result), httpResp, nil
}

// Update modifies an existing HardwareFilter and returns the updated object
func (s *hardwareFilterService) Update(ctx context.Context, id string, obj *dhcp.HardwareFilter, opts *core.Options) (*dhcp.HardwareFilter, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *hardwareFilterService) updateUDDI(ctx context.Context, id string, obj *dhcp.HardwareFilter, opts *core.Options) (*dhcp.HardwareFilter, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.HardwareFilter](obj, mapper.HardwareFilterUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAddressManagementAPI.HardwareFilterAPI.
		Update(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIHardwareFilterToResponse(&result), httpResp, nil
}

// Delete removes a HardwareFilter by ID
func (s *hardwareFilterService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *hardwareFilterService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.IPAddressManagementAPI.HardwareFilterAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves HardwareFilter objects based on filter options
func (s *hardwareFilterService) List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.HardwareFilter, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *hardwareFilterService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dhcp.HardwareFilter, *http.Response, string, error) {
	req := s.uddiClient.IPAddressManagementAPI.HardwareFilterAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.HardwareFilterFilterFieldMap[core.BackendUDDI])
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
	items := make([]*dhcp.HardwareFilter, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIHardwareFilterToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDIHardwareFilterToResponse(r *uddiipam.HardwareFilter) *dhcp.HardwareFilter {
	resp := &dhcp.HardwareFilter{
		Id: r.Id,
	}
	resp.UDDI = &dhcp.UDDIHardwareFilterExt{
		Addresses:                       r.Addresses,
		Comment:                         r.Comment,
		CreatedAt:                       r.CreatedAt,
		DhcpOptions:                     r.DhcpOptions,
		HeaderOptionFilename:            r.HeaderOptionFilename,
		HeaderOptionServerAddress:       r.HeaderOptionServerAddress,
		HeaderOptionServerName:          r.HeaderOptionServerName,
		LeaseTime:                       r.LeaseTime,
		Name:                            r.Name,
		Role:                            r.Role,
		UpdatedAt:                       r.UpdatedAt,
		VendorSpecificOptionOptionSpace: r.VendorSpecificOptionOptionSpace,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
