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

type Ipv6fixedaddressService interface {
	Create(ctx context.Context, obj *dhcp.Ipv6fixedaddress, opts *core.Options) (*dhcp.Ipv6fixedaddress, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dhcp.Ipv6fixedaddress, *http.Response, error)
	Update(ctx context.Context, id string, obj *dhcp.Ipv6fixedaddress, opts *core.Options) (*dhcp.Ipv6fixedaddress, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Ipv6fixedaddress, *http.Response, string, error)
}

type ipv6fixedaddressService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
	uddiClient *uddiclient.APIClient
}

func NewIpv6fixedaddressService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) Ipv6fixedaddressService {
	return &ipv6fixedaddressService{
		backend:    backend,
		niosClient: nios,
		uddiClient: uddi,
	}
}

// Create creates a new Ipv6fixedaddress and returns the created object
func (s *ipv6fixedaddressService) Create(ctx context.Context, obj *dhcp.Ipv6fixedaddress, opts *core.Options) (*dhcp.Ipv6fixedaddress, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6fixedaddressService) createNIOS(ctx context.Context, obj *dhcp.Ipv6fixedaddress, opts *core.Options) (*dhcp.Ipv6fixedaddress, *http.Response, error) {
	payload, err := common.MapTo[niosdhcp.Ipv6fixedaddress](obj, mapper.Ipv6fixedaddressNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if payload.FuncCall != nil && payload.Ipv6addr == nil {
		payload.Ipv6addr = &niosdhcp.Ipv6fixedaddressIpv6addr{}
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DHCPAPI.Ipv6fixedaddressAPI.
		Create(ctx).
		Ipv6fixedaddress(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateIpv6fixedaddressResponseAsObject.GetResult()

	return mapNIOSIpv6fixedaddressToResponse(&result), httpResp, nil
}

func (s *ipv6fixedaddressService) createUDDI(ctx context.Context, obj *dhcp.Ipv6fixedaddress, opts *core.Options) (*dhcp.Ipv6fixedaddress, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.FixedAddress](obj, mapper.Ipv6fixedaddressUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAddressManagementAPI.FixedAddressAPI.
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

	return mapUDDIIpv6fixedaddressToResponse(&result), httpResp, nil
}

// Read retrieves a Ipv6fixedaddress by ID
func (s *ipv6fixedaddressService) Read(ctx context.Context, id string, opts *core.Options) (*dhcp.Ipv6fixedaddress, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6fixedaddressService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dhcp.Ipv6fixedaddress, *http.Response, error) {
	req := s.niosClient.DHCPAPI.Ipv6fixedaddressAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetIpv6fixedaddressResponseObjectAsResult.GetResult()

	return mapNIOSIpv6fixedaddressToResponse(&result), httpResp, nil
}

func (s *ipv6fixedaddressService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dhcp.Ipv6fixedaddress, *http.Response, error) {
	req := s.uddiClient.IPAddressManagementAPI.FixedAddressAPI.
		Read(ctx, id)

	if opts != nil && opts.Inherit != "" {
		req = req.Inherit(opts.Inherit)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIIpv6fixedaddressToResponse(&result), httpResp, nil
}

// Update modifies an existing Ipv6fixedaddress and returns the updated object
func (s *ipv6fixedaddressService) Update(ctx context.Context, id string, obj *dhcp.Ipv6fixedaddress, opts *core.Options) (*dhcp.Ipv6fixedaddress, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6fixedaddressService) updateNIOS(ctx context.Context, id string, obj *dhcp.Ipv6fixedaddress, opts *core.Options) (*dhcp.Ipv6fixedaddress, *http.Response, error) {
	payload, err := common.MapTo[niosdhcp.Ipv6fixedaddress](obj, mapper.Ipv6fixedaddressNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DHCPAPI.Ipv6fixedaddressAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Ipv6fixedaddress(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateIpv6fixedaddressResponseAsObject.GetResult()

	return mapNIOSIpv6fixedaddressToResponse(&result), httpResp, nil
}

func (s *ipv6fixedaddressService) updateUDDI(ctx context.Context, id string, obj *dhcp.Ipv6fixedaddress, opts *core.Options) (*dhcp.Ipv6fixedaddress, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.FixedAddress](obj, mapper.Ipv6fixedaddressUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAddressManagementAPI.FixedAddressAPI.
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

	return mapUDDIIpv6fixedaddressToResponse(&result), httpResp, nil
}

// Delete removes a Ipv6fixedaddress by ID
func (s *ipv6fixedaddressService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6fixedaddressService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DHCPAPI.Ipv6fixedaddressAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

func (s *ipv6fixedaddressService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.IPAddressManagementAPI.FixedAddressAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves Ipv6fixedaddress objects based on filter options
func (s *ipv6fixedaddressService) List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Ipv6fixedaddress, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6fixedaddressService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Ipv6fixedaddress, *http.Response, string, error) {
	req := s.niosClient.DHCPAPI.Ipv6fixedaddressAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.Ipv6fixedaddressFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListIpv6fixedaddressResponseObject.GetResult()
	items := make([]*dhcp.Ipv6fixedaddress, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSIpv6fixedaddressToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListIpv6fixedaddressResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func (s *ipv6fixedaddressService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Ipv6fixedaddress, *http.Response, string, error) {
	req := s.uddiClient.IPAddressManagementAPI.FixedAddressAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.Ipv6fixedaddressFilterFieldMap[core.BackendUDDI])
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
	items := make([]*dhcp.Ipv6fixedaddress, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIIpv6fixedaddressToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapNIOSIpv6fixedaddressToResponse(r *niosdhcp.Ipv6fixedaddress) *dhcp.Ipv6fixedaddress {
	resp := &dhcp.Ipv6fixedaddress{
		Id: r.Ref,
	}
	resp.NIOS = &dhcp.NIOSIpv6fixedaddressExt{
		AddressType:              r.AddressType,
		AllowTelnet:              r.AllowTelnet,
		CliCredentials:           r.CliCredentials,
		CloudInfo:                r.CloudInfo,
		Comment:                  r.Comment,
		DeviceDescription:        r.DeviceDescription,
		DeviceLocation:           r.DeviceLocation,
		DeviceType:               r.DeviceType,
		DeviceVendor:             r.DeviceVendor,
		Disable:                  r.Disable,
		DisableDiscovery:         r.DisableDiscovery,
		DomainName:               r.DomainName,
		DomainNameServers:        r.DomainNameServers,
		Duid:                     r.Duid,
		EnableImmediateDiscovery: r.EnableImmediateDiscovery,
		Ipv6prefix:               r.Ipv6prefix,
		Ipv6prefixBits:           r.Ipv6prefixBits,
		LogicFilterRules:         r.LogicFilterRules,
		MacAddress:               r.MacAddress,
		MatchClient:              r.MatchClient,
		Name:                     r.Name,
		Network:                  r.Network,
		NetworkView:              r.NetworkView,
		Options:                  r.Options,
		PreferredLifetime:        r.PreferredLifetime,
		ReservedInterface:        r.ReservedInterface,
		RestartIfNeeded:          r.RestartIfNeeded,
		Snmp3Credential:          r.Snmp3Credential,
		SnmpCredential:           r.SnmpCredential,
		Template:                 r.Template,
		UseCliCredentials:        r.UseCliCredentials,
		UseDomainName:            r.UseDomainName,
		UseDomainNameServers:     r.UseDomainNameServers,
		UseLogicFilterRules:      r.UseLogicFilterRules,
		UseOptions:               r.UseOptions,
		UsePreferredLifetime:     r.UsePreferredLifetime,
		UseSnmp3Credential:       r.UseSnmp3Credential,
		UseSnmpCredential:        r.UseSnmpCredential,
		UseValidLifetime:         r.UseValidLifetime,
		ValidLifetime:            r.ValidLifetime,
	}
	if r.Ipv6addr != nil {
		resp.NIOS.Ipv6addr = r.Ipv6addr.String
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

func mapUDDIIpv6fixedaddressToResponse(r *uddiipam.FixedAddress) *dhcp.Ipv6fixedaddress {
	resp := &dhcp.Ipv6fixedaddress{
		Id: r.Id,
	}
	resp.UDDI = &dhcp.UDDIIpv6fixedaddressExt{
		Address:                   r.Address,
		Comment:                   r.Comment,
		DhcpOptions:               r.DhcpOptions,
		DisableDhcp:               r.DisableDhcp,
		HeaderOptionFilename:      r.HeaderOptionFilename,
		HeaderOptionServerAddress: r.HeaderOptionServerAddress,
		HeaderOptionServerName:    r.HeaderOptionServerName,
		Hostname:                  r.Hostname,
		InheritanceParent:         r.InheritanceParent,
		InheritanceSources:        r.InheritanceSources,
		IpSpace:                   r.IpSpace,
		MatchType:                 r.MatchType,
		MatchValue:                r.MatchValue,
		Name:                      r.Name,
		Parent:                    r.Parent,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
