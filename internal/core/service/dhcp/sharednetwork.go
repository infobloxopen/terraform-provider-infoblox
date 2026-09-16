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
)

type SharednetworkService interface {
	Create(ctx context.Context, obj *dhcp.Sharednetwork, opts *core.Options) (*dhcp.Sharednetwork, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dhcp.Sharednetwork, *http.Response, error)
	Update(ctx context.Context, id string, obj *dhcp.Sharednetwork, opts *core.Options) (*dhcp.Sharednetwork, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Sharednetwork, *http.Response, string, error)
}

type sharednetworkService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewSharednetworkService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) SharednetworkService {
	return &sharednetworkService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new Sharednetwork and returns the created object
func (s *sharednetworkService) Create(ctx context.Context, obj *dhcp.Sharednetwork, opts *core.Options) (*dhcp.Sharednetwork, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharednetworkService) createNIOS(ctx context.Context, obj *dhcp.Sharednetwork, opts *core.Options) (*dhcp.Sharednetwork, *http.Response, error) {
	payload, err := common.MapTo[niosdhcp.Sharednetwork](obj, mapper.SharednetworkNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DHCPAPI.SharednetworkAPI.
		Create(ctx).
		Sharednetwork(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateSharednetworkResponseAsObject.GetResult()

	return mapNIOSSharednetworkToResponse(&result), httpResp, nil
}

// Read retrieves a Sharednetwork by ID
func (s *sharednetworkService) Read(ctx context.Context, id string, opts *core.Options) (*dhcp.Sharednetwork, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharednetworkService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dhcp.Sharednetwork, *http.Response, error) {
	req := s.niosClient.DHCPAPI.SharednetworkAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetSharednetworkResponseObjectAsResult.GetResult()

	return mapNIOSSharednetworkToResponse(&result), httpResp, nil
}

// Update modifies an existing Sharednetwork and returns the updated object
func (s *sharednetworkService) Update(ctx context.Context, id string, obj *dhcp.Sharednetwork, opts *core.Options) (*dhcp.Sharednetwork, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharednetworkService) updateNIOS(ctx context.Context, id string, obj *dhcp.Sharednetwork, opts *core.Options) (*dhcp.Sharednetwork, *http.Response, error) {
	payload, err := common.MapTo[niosdhcp.Sharednetwork](obj, mapper.SharednetworkNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DHCPAPI.SharednetworkAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Sharednetwork(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateSharednetworkResponseAsObject.GetResult()

	return mapNIOSSharednetworkToResponse(&result), httpResp, nil
}

// Delete removes a Sharednetwork by ID
func (s *sharednetworkService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharednetworkService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DHCPAPI.SharednetworkAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves Sharednetwork objects based on filter options
func (s *sharednetworkService) List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Sharednetwork, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharednetworkService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Sharednetwork, *http.Response, string, error) {
	req := s.niosClient.DHCPAPI.SharednetworkAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.SharednetworkFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListSharednetworkResponseObject.GetResult()
	items := make([]*dhcp.Sharednetwork, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSSharednetworkToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListSharednetworkResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSSharednetworkToResponse(r *niosdhcp.Sharednetwork) *dhcp.Sharednetwork {
	resp := &dhcp.Sharednetwork{
		Id: r.Ref,
	}
	resp.NIOS = &dhcp.NIOSSharednetworkExt{
		Authority:                      r.Authority,
		Bootfile:                       r.Bootfile,
		Bootserver:                     r.Bootserver,
		Comment:                        r.Comment,
		DdnsGenerateHostname:           r.DdnsGenerateHostname,
		DdnsServerAlwaysUpdates:        r.DdnsServerAlwaysUpdates,
		DdnsTtl:                        r.DdnsTtl,
		DdnsUpdateFixedAddresses:       r.DdnsUpdateFixedAddresses,
		DdnsUseOption81:                r.DdnsUseOption81,
		DenyBootp:                      r.DenyBootp,
		Disable:                        r.Disable,
		EnableDdns:                     r.EnableDdns,
		EnablePxeLeaseTime:             r.EnablePxeLeaseTime,
		IgnoreClientIdentifier:         r.IgnoreClientIdentifier,
		IgnoreDhcpOptionListRequest:    r.IgnoreDhcpOptionListRequest,
		IgnoreId:                       r.IgnoreId,
		IgnoreMacAddresses:             r.IgnoreMacAddresses,
		LeaseScavengeTime:              r.LeaseScavengeTime,
		LogicFilterRules:               r.LogicFilterRules,
		Name:                           r.Name,
		NetworkView:                    r.NetworkView,
		Networks:                       r.Networks,
		Nextserver:                     r.Nextserver,
		Options:                        r.Options,
		PxeLeaseTime:                   r.PxeLeaseTime,
		UpdateDnsOnLeaseRenewal:        r.UpdateDnsOnLeaseRenewal,
		UseAuthority:                   r.UseAuthority,
		UseBootfile:                    r.UseBootfile,
		UseBootserver:                  r.UseBootserver,
		UseDdnsGenerateHostname:        r.UseDdnsGenerateHostname,
		UseDdnsTtl:                     r.UseDdnsTtl,
		UseDdnsUpdateFixedAddresses:    r.UseDdnsUpdateFixedAddresses,
		UseDdnsUseOption81:             r.UseDdnsUseOption81,
		UseDenyBootp:                   r.UseDenyBootp,
		UseEnableDdns:                  r.UseEnableDdns,
		UseIgnoreClientIdentifier:      r.UseIgnoreClientIdentifier,
		UseIgnoreDhcpOptionListRequest: r.UseIgnoreDhcpOptionListRequest,
		UseIgnoreId:                    r.UseIgnoreId,
		UseLeaseScavengeTime:           r.UseLeaseScavengeTime,
		UseLogicFilterRules:            r.UseLogicFilterRules,
		UseNextserver:                  r.UseNextserver,
		UseOptions:                     r.UseOptions,
		UsePxeLeaseTime:                r.UsePxeLeaseTime,
		UseUpdateDnsOnLeaseRenewal:     r.UseUpdateDnsOnLeaseRenewal,
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
