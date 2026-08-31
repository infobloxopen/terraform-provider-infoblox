package dhcp

import (
	"context"
	"fmt"
	"maps"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dhcp"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

type DhcpOptionspaceService interface {
	Create(ctx context.Context, obj *dhcp.DhcpOptionspace, opts *core.Options) (*dhcp.DhcpOptionspace, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dhcp.DhcpOptionspace, *http.Response, error)
	Update(ctx context.Context, id string, obj *dhcp.DhcpOptionspace, opts *core.Options) (*dhcp.DhcpOptionspace, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.DhcpOptionspace, *http.Response, string, error)
}

type dhcpOptionspaceService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
	uddiClient *uddiclient.APIClient
}

func NewDhcpOptionspaceService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) DhcpOptionspaceService {
	return &dhcpOptionspaceService{
		backend:    backend,
		niosClient: nios,
		uddiClient: uddi,
	}
}

// Create creates a new DhcpOptionspace and returns the created object
func (s *dhcpOptionspaceService) Create(ctx context.Context, obj *dhcp.DhcpOptionspace, opts *core.Options) (*dhcp.DhcpOptionspace, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dhcpOptionspaceService) createNIOS(ctx context.Context, obj *dhcp.DhcpOptionspace, opts *core.Options) (*dhcp.DhcpOptionspace, *http.Response, error) {
	payload, err := common.MapTo[niosdhcp.Dhcpoptionspace](obj, mapper.DhcpOptionspaceNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.DHCPAPI.DhcpoptionspaceAPI.
		Create(ctx).
		Dhcpoptionspace(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateDhcpoptionspaceResponseAsObject.GetResult()

	return mapNIOSDhcpOptionspaceToResponse(&result), httpResp, nil
}

func (s *dhcpOptionspaceService) createUDDI(ctx context.Context, obj *dhcp.DhcpOptionspace, opts *core.Options) (*dhcp.DhcpOptionspace, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.OptionSpace](obj, mapper.DhcpOptionspaceUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAddressManagementAPI.OptionSpaceAPI.
		Create(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIDhcpOptionspaceToResponse(&result), httpResp, nil
}

// Read retrieves a DhcpOptionspace by ID
func (s *dhcpOptionspaceService) Read(ctx context.Context, id string, opts *core.Options) (*dhcp.DhcpOptionspace, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dhcpOptionspaceService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dhcp.DhcpOptionspace, *http.Response, error) {
	req := s.niosClient.DHCPAPI.DhcpoptionspaceAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetDhcpoptionspaceResponseObjectAsResult.GetResult()

	return mapNIOSDhcpOptionspaceToResponse(&result), httpResp, nil
}

func (s *dhcpOptionspaceService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dhcp.DhcpOptionspace, *http.Response, error) {
	req := s.uddiClient.IPAddressManagementAPI.OptionSpaceAPI.
		Read(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIDhcpOptionspaceToResponse(&result), httpResp, nil
}

// Update modifies an existing DhcpOptionspace and returns the updated object
func (s *dhcpOptionspaceService) Update(ctx context.Context, id string, obj *dhcp.DhcpOptionspace, opts *core.Options) (*dhcp.DhcpOptionspace, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dhcpOptionspaceService) updateNIOS(ctx context.Context, id string, obj *dhcp.DhcpOptionspace, opts *core.Options) (*dhcp.DhcpOptionspace, *http.Response, error) {
	payload, err := common.MapTo[niosdhcp.Dhcpoptionspace](obj, mapper.DhcpOptionspaceNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.DHCPAPI.DhcpoptionspaceAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Dhcpoptionspace(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateDhcpoptionspaceResponseAsObject.GetResult()

	return mapNIOSDhcpOptionspaceToResponse(&result), httpResp, nil
}

func (s *dhcpOptionspaceService) updateUDDI(ctx context.Context, id string, obj *dhcp.DhcpOptionspace, opts *core.Options) (*dhcp.DhcpOptionspace, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.OptionSpace](obj, mapper.DhcpOptionspaceUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAddressManagementAPI.OptionSpaceAPI.
		Update(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIDhcpOptionspaceToResponse(&result), httpResp, nil
}

// Delete removes a DhcpOptionspace by ID
func (s *dhcpOptionspaceService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dhcpOptionspaceService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DHCPAPI.DhcpoptionspaceAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

func (s *dhcpOptionspaceService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.IPAddressManagementAPI.OptionSpaceAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves DhcpOptionspace objects based on filter options
func (s *dhcpOptionspaceService) List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.DhcpOptionspace, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dhcpOptionspaceService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dhcp.DhcpOptionspace, *http.Response, string, error) {
	req := s.niosClient.DHCPAPI.DhcpoptionspaceAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.DhcpOptionspaceFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListDhcpoptionspaceResponseObject.GetResult()
	items := make([]*dhcp.DhcpOptionspace, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSDhcpOptionspaceToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListDhcpoptionspaceResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func (s *dhcpOptionspaceService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dhcp.DhcpOptionspace, *http.Response, string, error) {
	req := s.uddiClient.IPAddressManagementAPI.OptionSpaceAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.DhcpOptionspaceFilterFieldMap[core.BackendUDDI])
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
	items := make([]*dhcp.DhcpOptionspace, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIDhcpOptionspaceToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapNIOSDhcpOptionspaceToResponse(r *niosdhcp.Dhcpoptionspace) *dhcp.DhcpOptionspace {
	resp := &dhcp.DhcpOptionspace{
		Id: r.Ref,
	}
	resp.NIOS = &dhcp.NIOSDhcpOptionspaceExt{
		Comment:           r.Comment,
		Name:              r.Name,
		OptionDefinitions: r.OptionDefinitions,
	}
	return resp
}

func mapUDDIDhcpOptionspaceToResponse(r *uddiipam.OptionSpace) *dhcp.DhcpOptionspace {
	resp := &dhcp.DhcpOptionspace{
		Id: r.Id,
	}
	resp.UDDI = &dhcp.UDDIDhcpOptionspaceExt{
		Comment:  r.Comment,
		Name:     r.Name,
		Protocol: r.Protocol,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
