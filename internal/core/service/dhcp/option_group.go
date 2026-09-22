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

type OptionGroupService interface {
	Create(ctx context.Context, obj *dhcp.OptionGroup, opts *core.Options) (*dhcp.OptionGroup, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dhcp.OptionGroup, *http.Response, error)
	Update(ctx context.Context, id string, obj *dhcp.OptionGroup, opts *core.Options) (*dhcp.OptionGroup, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.OptionGroup, *http.Response, string, error)
}

type optionGroupService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewOptionGroupService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) OptionGroupService {
	return &optionGroupService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new OptionGroup and returns the created object
func (s *optionGroupService) Create(ctx context.Context, obj *dhcp.OptionGroup, opts *core.Options) (*dhcp.OptionGroup, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *optionGroupService) createUDDI(ctx context.Context, obj *dhcp.OptionGroup, opts *core.Options) (*dhcp.OptionGroup, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.OptionGroup](obj, mapper.OptionGroupUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAddressManagementAPI.OptionGroupAPI.
		Create(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIOptionGroupToResponse(&result), httpResp, nil
}

// Read retrieves a OptionGroup by ID
func (s *optionGroupService) Read(ctx context.Context, id string, opts *core.Options) (*dhcp.OptionGroup, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *optionGroupService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dhcp.OptionGroup, *http.Response, error) {
	req := s.uddiClient.IPAddressManagementAPI.OptionGroupAPI.
		Read(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIOptionGroupToResponse(&result), httpResp, nil
}

// Update modifies an existing OptionGroup and returns the updated object
func (s *optionGroupService) Update(ctx context.Context, id string, obj *dhcp.OptionGroup, opts *core.Options) (*dhcp.OptionGroup, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *optionGroupService) updateUDDI(ctx context.Context, id string, obj *dhcp.OptionGroup, opts *core.Options) (*dhcp.OptionGroup, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.OptionGroup](obj, mapper.OptionGroupUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAddressManagementAPI.OptionGroupAPI.
		Update(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIOptionGroupToResponse(&result), httpResp, nil
}

// Delete removes a OptionGroup by ID
func (s *optionGroupService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *optionGroupService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.IPAddressManagementAPI.OptionGroupAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves OptionGroup objects based on filter options
func (s *optionGroupService) List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.OptionGroup, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *optionGroupService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dhcp.OptionGroup, *http.Response, string, error) {
	req := s.uddiClient.IPAddressManagementAPI.OptionGroupAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.OptionGroupFilterFieldMap[core.BackendUDDI])
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
	items := make([]*dhcp.OptionGroup, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIOptionGroupToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDIOptionGroupToResponse(r *uddiipam.OptionGroup) *dhcp.OptionGroup {
	resp := &dhcp.OptionGroup{
		Id: r.Id,
	}
	resp.UDDI = &dhcp.UDDIOptionGroupExt{
		Comment:     r.Comment,
		CreatedAt:   r.CreatedAt,
		DhcpOptions: r.DhcpOptions,
		Name:        r.Name,
		Protocol:    r.Protocol,
		UpdatedAt:   r.UpdatedAt,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
