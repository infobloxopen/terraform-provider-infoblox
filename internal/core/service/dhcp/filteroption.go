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

type FilteroptionService interface {
	Create(ctx context.Context, obj *dhcp.Filteroption, opts *core.Options) (*dhcp.Filteroption, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dhcp.Filteroption, *http.Response, error)
	Update(ctx context.Context, id string, obj *dhcp.Filteroption, opts *core.Options) (*dhcp.Filteroption, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Filteroption, *http.Response, string, error)
}

type filteroptionService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
	uddiClient *uddiclient.APIClient
}

func NewFilteroptionService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) FilteroptionService {
	return &filteroptionService{
		backend:    backend,
		niosClient: nios,
		uddiClient: uddi,
	}
}

// Create creates a new Filteroption and returns the created object
func (s *filteroptionService) Create(ctx context.Context, obj *dhcp.Filteroption, opts *core.Options) (*dhcp.Filteroption, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *filteroptionService) createNIOS(ctx context.Context, obj *dhcp.Filteroption, opts *core.Options) (*dhcp.Filteroption, *http.Response, error) {
	payload, err := common.MapTo[niosdhcp.Filteroption](obj, mapper.FilteroptionNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DHCPAPI.FilteroptionAPI.
		Create(ctx).
		Filteroption(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateFilteroptionResponseAsObject.GetResult()

	return mapNIOSFilteroptionToResponse(&result), httpResp, nil
}

func (s *filteroptionService) createUDDI(ctx context.Context, obj *dhcp.Filteroption, opts *core.Options) (*dhcp.Filteroption, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.OptionFilter](obj, mapper.FilteroptionUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAddressManagementAPI.OptionFilterAPI.
		Create(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIFilteroptionToResponse(&result), httpResp, nil
}

// Read retrieves a Filteroption by ID
func (s *filteroptionService) Read(ctx context.Context, id string, opts *core.Options) (*dhcp.Filteroption, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *filteroptionService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dhcp.Filteroption, *http.Response, error) {
	req := s.niosClient.DHCPAPI.FilteroptionAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetFilteroptionResponseObjectAsResult.GetResult()

	return mapNIOSFilteroptionToResponse(&result), httpResp, nil
}

func (s *filteroptionService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dhcp.Filteroption, *http.Response, error) {
	req := s.uddiClient.IPAddressManagementAPI.OptionFilterAPI.
		Read(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIFilteroptionToResponse(&result), httpResp, nil
}

// Update modifies an existing Filteroption and returns the updated object
func (s *filteroptionService) Update(ctx context.Context, id string, obj *dhcp.Filteroption, opts *core.Options) (*dhcp.Filteroption, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *filteroptionService) updateNIOS(ctx context.Context, id string, obj *dhcp.Filteroption, opts *core.Options) (*dhcp.Filteroption, *http.Response, error) {
	payload, err := common.MapTo[niosdhcp.Filteroption](obj, mapper.FilteroptionNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DHCPAPI.FilteroptionAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Filteroption(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateFilteroptionResponseAsObject.GetResult()

	return mapNIOSFilteroptionToResponse(&result), httpResp, nil
}

func (s *filteroptionService) updateUDDI(ctx context.Context, id string, obj *dhcp.Filteroption, opts *core.Options) (*dhcp.Filteroption, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.OptionFilter](obj, mapper.FilteroptionUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAddressManagementAPI.OptionFilterAPI.
		Update(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIFilteroptionToResponse(&result), httpResp, nil
}

// Delete removes a Filteroption by ID
func (s *filteroptionService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *filteroptionService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DHCPAPI.FilteroptionAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

func (s *filteroptionService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.IPAddressManagementAPI.OptionFilterAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves Filteroption objects based on filter options
func (s *filteroptionService) List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Filteroption, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *filteroptionService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Filteroption, *http.Response, string, error) {
	req := s.niosClient.DHCPAPI.FilteroptionAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.FilteroptionFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListFilteroptionResponseObject.GetResult()
	items := make([]*dhcp.Filteroption, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSFilteroptionToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListFilteroptionResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func (s *filteroptionService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Filteroption, *http.Response, string, error) {
	req := s.uddiClient.IPAddressManagementAPI.OptionFilterAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.FilteroptionFilterFieldMap[core.BackendUDDI])
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
	items := make([]*dhcp.Filteroption, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIFilteroptionToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapNIOSFilteroptionToResponse(r *niosdhcp.Filteroption) *dhcp.Filteroption {
	resp := &dhcp.Filteroption{
		Id: r.Ref,
	}
	resp.NIOS = &dhcp.NIOSFilteroptionExt{
		ApplyAsClass: r.ApplyAsClass,
		Bootfile:     r.Bootfile,
		Bootserver:   r.Bootserver,
		Comment:      r.Comment,
		Expression:   r.Expression,
		LeaseTime:    r.LeaseTime,
		Name:         r.Name,
		NextServer:   r.NextServer,
		OptionList:   r.OptionList,
		OptionSpace:  r.OptionSpace,
		PxeLeaseTime: r.PxeLeaseTime,
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

func mapUDDIFilteroptionToResponse(r *uddiipam.OptionFilter) *dhcp.Filteroption {
	resp := &dhcp.Filteroption{
		Id: r.Id,
	}
	resp.UDDI = &dhcp.UDDIFilteroptionExt{
		Comment:                         r.Comment,
		DhcpOptions:                     r.DhcpOptions,
		HeaderOptionFilename:            r.HeaderOptionFilename,
		HeaderOptionServerAddress:       r.HeaderOptionServerAddress,
		HeaderOptionServerName:          r.HeaderOptionServerName,
		LeaseTime:                       r.LeaseTime,
		Name:                            r.Name,
		Protocol:                        r.Protocol,
		Role:                            r.Role,
		Rules:                           &r.Rules,
		VendorSpecificOptionOptionSpace: r.VendorSpecificOptionOptionSpace,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
