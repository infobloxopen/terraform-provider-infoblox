package ipam

import (
	"context"
	"fmt"
	"maps"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/ipam"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

type NetworkviewService interface {
	Create(ctx context.Context, obj *ipam.Networkview, opts *core.Options) (*ipam.Networkview, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*ipam.Networkview, *http.Response, error)
	Update(ctx context.Context, id string, obj *ipam.Networkview, opts *core.Options) (*ipam.Networkview, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*ipam.Networkview, *http.Response, string, error)
}

type networkviewService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
	uddiClient *uddiclient.APIClient
}

func NewNetworkviewService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) NetworkviewService {
	return &networkviewService{
		backend:    backend,
		niosClient: nios,
		uddiClient: uddi,
	}
}

// Create creates a new Networkview and returns the created object
func (s *networkviewService) Create(ctx context.Context, obj *ipam.Networkview, opts *core.Options) (*ipam.Networkview, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *networkviewService) createNIOS(ctx context.Context, obj *ipam.Networkview, opts *core.Options) (*ipam.Networkview, *http.Response, error) {
	payload, err := common.MapTo[niosipam.Networkview](obj, mapper.NetworkviewNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.IPAMAPI.NetworkviewAPI.
		Create(ctx).
		Networkview(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateNetworkviewResponseAsObject.GetResult()

	return mapNIOSNetworkviewToResponse(&result), httpResp, nil
}

func (s *networkviewService) createUDDI(ctx context.Context, obj *ipam.Networkview, opts *core.Options) (*ipam.Networkview, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.IPSpace](obj, mapper.NetworkviewUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAddressManagementAPI.IpSpaceAPI.
		Create(ctx).
		Body(payload)

	if opts != nil && opts.Inherit != "" {
		req = req.Inherit(opts.Inherit)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDINetworkviewToResponse(&result), httpResp, nil
}

// Read retrieves a Networkview by ID
func (s *networkviewService) Read(ctx context.Context, id string, opts *core.Options) (*ipam.Networkview, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *networkviewService) readNIOS(ctx context.Context, id string, opts *core.Options) (*ipam.Networkview, *http.Response, error) {
	req := s.niosClient.IPAMAPI.NetworkviewAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetNetworkviewResponseObjectAsResult.GetResult()

	return mapNIOSNetworkviewToResponse(&result), httpResp, nil
}

func (s *networkviewService) readUDDI(ctx context.Context, id string, opts *core.Options) (*ipam.Networkview, *http.Response, error) {
	req := s.uddiClient.IPAddressManagementAPI.IpSpaceAPI.
		Read(ctx, id)

	if opts != nil && opts.Inherit != "" {
		req = req.Inherit(opts.Inherit)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDINetworkviewToResponse(&result), httpResp, nil
}

// Update modifies an existing Networkview and returns the updated object
func (s *networkviewService) Update(ctx context.Context, id string, obj *ipam.Networkview, opts *core.Options) (*ipam.Networkview, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *networkviewService) updateNIOS(ctx context.Context, id string, obj *ipam.Networkview, opts *core.Options) (*ipam.Networkview, *http.Response, error) {
	payload, err := common.MapTo[niosipam.Networkview](obj, mapper.NetworkviewNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.IPAMAPI.NetworkviewAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Networkview(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateNetworkviewResponseAsObject.GetResult()

	return mapNIOSNetworkviewToResponse(&result), httpResp, nil
}

func (s *networkviewService) updateUDDI(ctx context.Context, id string, obj *ipam.Networkview, opts *core.Options) (*ipam.Networkview, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.IPSpace](obj, mapper.NetworkviewUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAddressManagementAPI.IpSpaceAPI.
		Update(ctx, id).
		Body(payload)

	if opts != nil && opts.Inherit != "" {
		req = req.Inherit(opts.Inherit)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDINetworkviewToResponse(&result), httpResp, nil
}

// Delete removes a Networkview by ID
func (s *networkviewService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *networkviewService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.IPAMAPI.NetworkviewAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

func (s *networkviewService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.IPAddressManagementAPI.IpSpaceAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves Networkview objects based on filter options
func (s *networkviewService) List(ctx context.Context, opts *core.ListOptions) ([]*ipam.Networkview, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *networkviewService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*ipam.Networkview, *http.Response, string, error) {
	req := s.niosClient.IPAMAPI.NetworkviewAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.NetworkviewFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListNetworkviewResponseObject.GetResult()
	items := make([]*ipam.Networkview, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSNetworkviewToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListNetworkviewResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func (s *networkviewService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*ipam.Networkview, *http.Response, string, error) {
	req := s.uddiClient.IPAddressManagementAPI.IpSpaceAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.NetworkviewFilterFieldMap[core.BackendUDDI])
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
	items := make([]*ipam.Networkview, 0, len(results))
	for i := range results {
		items = append(items, mapUDDINetworkviewToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapNIOSNetworkviewToResponse(r *niosipam.Networkview) *ipam.Networkview {
	resp := &ipam.Networkview{
		Id: r.Ref,
	}
	resp.NIOS = &ipam.NIOSNetworkviewExt{
		CloudInfo:            r.CloudInfo,
		Comment:              r.Comment,
		DdnsDnsView:          r.DdnsDnsView,
		DdnsZonePrimaries:    r.DdnsZonePrimaries,
		FederatedRealms:      r.FederatedRealms,
		InternalForwardZones: r.InternalForwardZones,
		MgmPrivate:           r.MgmPrivate,
		Name:                 r.Name,
		RemoteForwardZones:   r.RemoteForwardZones,
		RemoteReverseZones:   r.RemoteReverseZones,
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

func mapUDDINetworkviewToResponse(r *uddiipam.IPSpace) *ipam.Networkview {
	resp := &ipam.Networkview{
		Id: r.Id,
	}
	resp.UDDI = &ipam.UDDINetworkviewExt{
		AsmConfig:                       r.AsmConfig,
		Comment:                         r.Comment,
		CompartmentId:                   r.CompartmentId,
		DdnsClientUpdate:                r.DdnsClientUpdate,
		DdnsConflictResolutionMode:      r.DdnsConflictResolutionMode,
		DdnsDomain:                      r.DdnsDomain,
		DdnsGenerateName:                r.DdnsGenerateName,
		DdnsGeneratedPrefix:             r.DdnsGeneratedPrefix,
		DdnsSendUpdates:                 r.DdnsSendUpdates,
		DdnsTtlPercent:                  r.DdnsTtlPercent,
		DdnsUpdateOnRenew:               r.DdnsUpdateOnRenew,
		DdnsUseConflictResolution:       r.DdnsUseConflictResolution,
		DefaultRealms:                   r.DefaultRealms,
		DhcpConfig:                      r.DhcpConfig,
		DhcpOptions:                     r.DhcpOptions,
		DhcpOptionsV6:                   r.DhcpOptionsV6,
		HeaderOptionFilename:            r.HeaderOptionFilename,
		HeaderOptionServerAddress:       r.HeaderOptionServerAddress,
		HeaderOptionServerName:          r.HeaderOptionServerName,
		HostnameRewriteChar:             r.HostnameRewriteChar,
		HostnameRewriteEnabled:          r.HostnameRewriteEnabled,
		HostnameRewriteRegex:            r.HostnameRewriteRegex,
		InheritanceSources:              r.InheritanceSources,
		Name:                            r.Name,
		VendorSpecificOptionOptionSpace: r.VendorSpecificOptionOptionSpace,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
