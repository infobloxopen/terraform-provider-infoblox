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

type FixedaddressService interface {
	Create(ctx context.Context, obj *dhcp.Fixedaddress, opts *core.Options) (*dhcp.Fixedaddress, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dhcp.Fixedaddress, *http.Response, error)
	Update(ctx context.Context, id string, obj *dhcp.Fixedaddress, opts *core.Options) (*dhcp.Fixedaddress, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Fixedaddress, *http.Response, string, error)
}

type fixedaddressService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
	uddiClient *uddiclient.APIClient
}

func NewFixedaddressService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) FixedaddressService {
	return &fixedaddressService{
		backend:    backend,
		niosClient: nios,
		uddiClient: uddi,
	}
}

// Create creates a new Fixedaddress and returns the created object
func (s *fixedaddressService) Create(ctx context.Context, obj *dhcp.Fixedaddress, opts *core.Options) (*dhcp.Fixedaddress, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *fixedaddressService) createNIOS(ctx context.Context, obj *dhcp.Fixedaddress, opts *core.Options) (*dhcp.Fixedaddress, *http.Response, error) {
	payload, err := common.MapTo[niosdhcp.Fixedaddress](obj, mapper.FixedaddressNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if payload.FuncCall != nil && payload.Ipv4addr == nil {
		payload.Ipv4addr = &niosdhcp.FixedaddressIpv4addr{}
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DHCPAPI.FixedaddressAPI.
		Create(ctx).
		Fixedaddress(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateFixedaddressResponseAsObject.GetResult()

	return mapNIOSFixedaddressToResponse(&result), httpResp, nil
}

func (s *fixedaddressService) createUDDI(ctx context.Context, obj *dhcp.Fixedaddress, opts *core.Options) (*dhcp.Fixedaddress, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.FixedAddress](obj, mapper.FixedaddressUDDIFieldMap)
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

	return mapUDDIFixedaddressToResponse(&result), httpResp, nil
}

// Read retrieves a Fixedaddress by ID
func (s *fixedaddressService) Read(ctx context.Context, id string, opts *core.Options) (*dhcp.Fixedaddress, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *fixedaddressService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dhcp.Fixedaddress, *http.Response, error) {
	req := s.niosClient.DHCPAPI.FixedaddressAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetFixedaddressResponseObjectAsResult.GetResult()

	return mapNIOSFixedaddressToResponse(&result), httpResp, nil
}

func (s *fixedaddressService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dhcp.Fixedaddress, *http.Response, error) {
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

	return mapUDDIFixedaddressToResponse(&result), httpResp, nil
}

// Update modifies an existing Fixedaddress and returns the updated object
func (s *fixedaddressService) Update(ctx context.Context, id string, obj *dhcp.Fixedaddress, opts *core.Options) (*dhcp.Fixedaddress, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *fixedaddressService) updateNIOS(ctx context.Context, id string, obj *dhcp.Fixedaddress, opts *core.Options) (*dhcp.Fixedaddress, *http.Response, error) {
	payload, err := common.MapTo[niosdhcp.Fixedaddress](obj, mapper.FixedaddressNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DHCPAPI.FixedaddressAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Fixedaddress(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateFixedaddressResponseAsObject.GetResult()

	return mapNIOSFixedaddressToResponse(&result), httpResp, nil
}

func (s *fixedaddressService) updateUDDI(ctx context.Context, id string, obj *dhcp.Fixedaddress, opts *core.Options) (*dhcp.Fixedaddress, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.FixedAddress](obj, mapper.FixedaddressUDDIFieldMap)
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

	return mapUDDIFixedaddressToResponse(&result), httpResp, nil
}

// Delete removes a Fixedaddress by ID
func (s *fixedaddressService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *fixedaddressService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DHCPAPI.FixedaddressAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

func (s *fixedaddressService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.IPAddressManagementAPI.FixedAddressAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves Fixedaddress objects based on filter options
func (s *fixedaddressService) List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Fixedaddress, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *fixedaddressService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Fixedaddress, *http.Response, string, error) {
	req := s.niosClient.DHCPAPI.FixedaddressAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.FixedaddressFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListFixedaddressResponseObject.GetResult()
	items := make([]*dhcp.Fixedaddress, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSFixedaddressToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListFixedaddressResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func (s *fixedaddressService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Fixedaddress, *http.Response, string, error) {
	req := s.uddiClient.IPAddressManagementAPI.FixedAddressAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.FixedaddressFilterFieldMap[core.BackendUDDI])
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
	items := make([]*dhcp.Fixedaddress, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIFixedaddressToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapNIOSFixedaddressToResponse(r *niosdhcp.Fixedaddress) *dhcp.Fixedaddress {
	resp := &dhcp.Fixedaddress{
		Id: r.Ref,
	}
	resp.NIOS = &dhcp.NIOSFixedaddressExt{
		AgentCircuitId:                 r.AgentCircuitId,
		AgentRemoteId:                  r.AgentRemoteId,
		AllowTelnet:                    r.AllowTelnet,
		AlwaysUpdateDns:                r.AlwaysUpdateDns,
		Bootfile:                       r.Bootfile,
		Bootserver:                     r.Bootserver,
		CliCredentials:                 r.CliCredentials,
		ClientIdentifierPrependZero:    r.ClientIdentifierPrependZero,
		CloudInfo:                      r.CloudInfo,
		Comment:                        r.Comment,
		DdnsDomainname:                 r.DdnsDomainname,
		DdnsHostname:                   r.DdnsHostname,
		DenyBootp:                      r.DenyBootp,
		DeviceDescription:              r.DeviceDescription,
		DeviceLocation:                 r.DeviceLocation,
		DeviceType:                     r.DeviceType,
		DeviceVendor:                   r.DeviceVendor,
		DhcpClientIdentifier:           r.DhcpClientIdentifier,
		Disable:                        r.Disable,
		DisableDiscovery:               r.DisableDiscovery,
		EnableDdns:                     r.EnableDdns,
		EnableImmediateDiscovery:       r.EnableImmediateDiscovery,
		EnablePxeLeaseTime:             r.EnablePxeLeaseTime,
		IgnoreDhcpOptionListRequest:    r.IgnoreDhcpOptionListRequest,
		LogicFilterRules:               r.LogicFilterRules,
		Mac:                            r.Mac,
		MatchClient:                    r.MatchClient,
		MsOptions:                      r.MsOptions,
		MsServer:                       r.MsServer,
		Name:                           r.Name,
		Network:                        r.Network,
		NetworkView:                    r.NetworkView,
		Nextserver:                     r.Nextserver,
		Options:                        r.Options,
		PxeLeaseTime:                   r.PxeLeaseTime,
		ReservedInterface:              r.ReservedInterface,
		RestartIfNeeded:                r.RestartIfNeeded,
		Snmp3Credential:                r.Snmp3Credential,
		SnmpCredential:                 r.SnmpCredential,
		Template:                       r.Template,
		UseBootfile:                    r.UseBootfile,
		UseBootserver:                  r.UseBootserver,
		UseCliCredentials:              r.UseCliCredentials,
		UseDdnsDomainname:              r.UseDdnsDomainname,
		UseDenyBootp:                   r.UseDenyBootp,
		UseEnableDdns:                  r.UseEnableDdns,
		UseIgnoreDhcpOptionListRequest: r.UseIgnoreDhcpOptionListRequest,
		UseLogicFilterRules:            r.UseLogicFilterRules,
		UseMsOptions:                   r.UseMsOptions,
		UseNextserver:                  r.UseNextserver,
		UseOptions:                     r.UseOptions,
		UsePxeLeaseTime:                r.UsePxeLeaseTime,
		UseSnmp3Credential:             r.UseSnmp3Credential,
		UseSnmpCredential:              r.UseSnmpCredential,
	}
	if r.Ipv4addr != nil {
		resp.NIOS.Ipv4addr = r.Ipv4addr.String
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

func mapUDDIFixedaddressToResponse(r *uddiipam.FixedAddress) *dhcp.Fixedaddress {
	resp := &dhcp.Fixedaddress{
		Id: r.Id,
	}
	resp.UDDI = &dhcp.UDDIFixedaddressExt{
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
