package dhcp

import (
	"context"
	"fmt"
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

type DhcpOptiondefinitionService interface {
	Create(ctx context.Context, obj *dhcp.DhcpOptiondefinition, opts *core.Options) (*dhcp.DhcpOptiondefinition, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dhcp.DhcpOptiondefinition, *http.Response, error)
	Update(ctx context.Context, id string, obj *dhcp.DhcpOptiondefinition, opts *core.Options) (*dhcp.DhcpOptiondefinition, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.DhcpOptiondefinition, *http.Response, string, error)
}

type dhcpOptiondefinitionService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
	uddiClient *uddiclient.APIClient
}

func NewDhcpOptiondefinitionService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) DhcpOptiondefinitionService {
	return &dhcpOptiondefinitionService{
		backend:    backend,
		niosClient: nios,
		uddiClient: uddi,
	}
}

// Create creates a new DhcpOptiondefinition and returns the created object
func (s *dhcpOptiondefinitionService) Create(ctx context.Context, obj *dhcp.DhcpOptiondefinition, opts *core.Options) (*dhcp.DhcpOptiondefinition, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dhcpOptiondefinitionService) createNIOS(ctx context.Context, obj *dhcp.DhcpOptiondefinition, opts *core.Options) (*dhcp.DhcpOptiondefinition, *http.Response, error) {
	payload, err := common.MapTo[niosdhcp.Dhcpoptiondefinition](obj, mapper.DhcpOptiondefinitionNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.DHCPAPI.DhcpoptiondefinitionAPI.
		Create(ctx).
		Dhcpoptiondefinition(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateDhcpoptiondefinitionResponseAsObject.GetResult()

	return mapNIOSDhcpOptiondefinitionToResponse(&result), httpResp, nil
}

func (s *dhcpOptiondefinitionService) createUDDI(ctx context.Context, obj *dhcp.DhcpOptiondefinition, opts *core.Options) (*dhcp.DhcpOptiondefinition, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.OptionCode](obj, mapper.DhcpOptiondefinitionUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAddressManagementAPI.OptionCodeAPI.
		Create(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIDhcpOptiondefinitionToResponse(&result), httpResp, nil
}

// Read retrieves a DhcpOptiondefinition by ID
func (s *dhcpOptiondefinitionService) Read(ctx context.Context, id string, opts *core.Options) (*dhcp.DhcpOptiondefinition, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dhcpOptiondefinitionService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dhcp.DhcpOptiondefinition, *http.Response, error) {
	req := s.niosClient.DHCPAPI.DhcpoptiondefinitionAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetDhcpoptiondefinitionResponseObjectAsResult.GetResult()

	return mapNIOSDhcpOptiondefinitionToResponse(&result), httpResp, nil
}

func (s *dhcpOptiondefinitionService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dhcp.DhcpOptiondefinition, *http.Response, error) {
	req := s.uddiClient.IPAddressManagementAPI.OptionCodeAPI.
		Read(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIDhcpOptiondefinitionToResponse(&result), httpResp, nil
}

// Update modifies an existing DhcpOptiondefinition and returns the updated object
func (s *dhcpOptiondefinitionService) Update(ctx context.Context, id string, obj *dhcp.DhcpOptiondefinition, opts *core.Options) (*dhcp.DhcpOptiondefinition, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dhcpOptiondefinitionService) updateNIOS(ctx context.Context, id string, obj *dhcp.DhcpOptiondefinition, opts *core.Options) (*dhcp.DhcpOptiondefinition, *http.Response, error) {
	payload, err := common.MapTo[niosdhcp.Dhcpoptiondefinition](obj, mapper.DhcpOptiondefinitionNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.DHCPAPI.DhcpoptiondefinitionAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Dhcpoptiondefinition(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateDhcpoptiondefinitionResponseAsObject.GetResult()

	return mapNIOSDhcpOptiondefinitionToResponse(&result), httpResp, nil
}

func (s *dhcpOptiondefinitionService) updateUDDI(ctx context.Context, id string, obj *dhcp.DhcpOptiondefinition, opts *core.Options) (*dhcp.DhcpOptiondefinition, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.OptionCode](obj, mapper.DhcpOptiondefinitionUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAddressManagementAPI.OptionCodeAPI.
		Update(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIDhcpOptiondefinitionToResponse(&result), httpResp, nil
}

// Delete removes a DhcpOptiondefinition by ID
func (s *dhcpOptiondefinitionService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dhcpOptiondefinitionService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DHCPAPI.DhcpoptiondefinitionAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

func (s *dhcpOptiondefinitionService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.IPAddressManagementAPI.OptionCodeAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves DhcpOptiondefinition objects based on filter options
func (s *dhcpOptiondefinitionService) List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.DhcpOptiondefinition, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dhcpOptiondefinitionService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dhcp.DhcpOptiondefinition, *http.Response, string, error) {
	req := s.niosClient.DHCPAPI.DhcpoptiondefinitionAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.DhcpOptiondefinitionFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListDhcpoptiondefinitionResponseObject.GetResult()
	items := make([]*dhcp.DhcpOptiondefinition, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSDhcpOptiondefinitionToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListDhcpoptiondefinitionResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func (s *dhcpOptiondefinitionService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dhcp.DhcpOptiondefinition, *http.Response, string, error) {
	req := s.uddiClient.IPAddressManagementAPI.OptionCodeAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.DhcpOptiondefinitionFilterFieldMap[core.BackendUDDI])
		for k, v := range translatedFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		if len(filters) > 0 {
			req = req.Filter(core.JoinFilters(filters))
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
	items := make([]*dhcp.DhcpOptiondefinition, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIDhcpOptiondefinitionToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapNIOSDhcpOptiondefinitionToResponse(r *niosdhcp.Dhcpoptiondefinition) *dhcp.DhcpOptiondefinition {
	resp := &dhcp.DhcpOptiondefinition{
		Id: r.Ref,
	}
	resp.NIOS = &dhcp.NIOSDhcpOptiondefinitionExt{
		Code:  r.Code,
		Name:  r.Name,
		Space: r.Space,
		Type:  r.Type,
	}
	return resp
}

func mapUDDIDhcpOptiondefinitionToResponse(r *uddiipam.OptionCode) *dhcp.DhcpOptiondefinition {
	resp := &dhcp.DhcpOptiondefinition{
		Id: r.Id,
	}
	resp.UDDI = &dhcp.UDDIDhcpOptiondefinitionExt{
		Array:       r.Array,
		Code:        r.Code,
		Comment:     r.Comment,
		Name:        r.Name,
		OptionSpace: r.OptionSpace,
		Type:        r.Type,
	}
	return resp
}
