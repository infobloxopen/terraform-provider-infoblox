package ipam

import (
	"context"
	"fmt"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/ipam"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
)

type Ipv6networkcontainerService interface {
	Create(ctx context.Context, obj *ipam.Ipv6networkcontainer, opts *core.Options) (*ipam.Ipv6networkcontainer, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*ipam.Ipv6networkcontainer, *http.Response, error)
	Update(ctx context.Context, id string, obj *ipam.Ipv6networkcontainer, opts *core.Options) (*ipam.Ipv6networkcontainer, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*ipam.Ipv6networkcontainer, *http.Response, string, error)
}

type ipv6networkcontainerService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewIpv6networkcontainerService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) Ipv6networkcontainerService {
	return &ipv6networkcontainerService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new Ipv6networkcontainer and returns the created object
func (s *ipv6networkcontainerService) Create(ctx context.Context, obj *ipam.Ipv6networkcontainer, opts *core.Options) (*ipam.Ipv6networkcontainer, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6networkcontainerService) createNIOS(ctx context.Context, obj *ipam.Ipv6networkcontainer, opts *core.Options) (*ipam.Ipv6networkcontainer, *http.Response, error) {
	payload, err := common.MapTo[niosipam.Ipv6networkcontainer](obj, mapper.Ipv6networkcontainerNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.IPAMAPI.Ipv6networkcontainerAPI.
		Create(ctx).
		Ipv6networkcontainer(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateIpv6networkcontainerResponseAsObject.GetResult()

	return mapNIOSIpv6networkcontainerToResponse(&result), httpResp, nil
}

// Read retrieves a Ipv6networkcontainer by ID
func (s *ipv6networkcontainerService) Read(ctx context.Context, id string, opts *core.Options) (*ipam.Ipv6networkcontainer, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6networkcontainerService) readNIOS(ctx context.Context, id string, opts *core.Options) (*ipam.Ipv6networkcontainer, *http.Response, error) {
	req := s.niosClient.IPAMAPI.Ipv6networkcontainerAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetIpv6networkcontainerResponseObjectAsResult.GetResult()

	return mapNIOSIpv6networkcontainerToResponse(&result), httpResp, nil
}

// Update modifies an existing Ipv6networkcontainer and returns the updated object
func (s *ipv6networkcontainerService) Update(ctx context.Context, id string, obj *ipam.Ipv6networkcontainer, opts *core.Options) (*ipam.Ipv6networkcontainer, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6networkcontainerService) updateNIOS(ctx context.Context, id string, obj *ipam.Ipv6networkcontainer, opts *core.Options) (*ipam.Ipv6networkcontainer, *http.Response, error) {
	payload, err := common.MapTo[niosipam.Ipv6networkcontainer](obj, mapper.Ipv6networkcontainerNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.IPAMAPI.Ipv6networkcontainerAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Ipv6networkcontainer(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateIpv6networkcontainerResponseAsObject.GetResult()

	return mapNIOSIpv6networkcontainerToResponse(&result), httpResp, nil
}

// Delete removes a Ipv6networkcontainer by ID
func (s *ipv6networkcontainerService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6networkcontainerService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.IPAMAPI.Ipv6networkcontainerAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves Ipv6networkcontainer objects based on filter options
func (s *ipv6networkcontainerService) List(ctx context.Context, opts *core.ListOptions) ([]*ipam.Ipv6networkcontainer, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *ipv6networkcontainerService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*ipam.Ipv6networkcontainer, *http.Response, string, error) {
	req := s.niosClient.IPAMAPI.Ipv6networkcontainerAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.Ipv6networkcontainerFilterFieldMap[core.BackendNIOS])
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
		if opts.Paging == 1 {
			maxResults := opts.MaxResults
			if maxResults <= 0 {
				maxResults = core.DefaultListLimit
			}
			req = req.MaxResults(maxResults)
		} else if opts.MaxResults > 0 {
			req = req.MaxResults(opts.MaxResults)
		}
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, "", err
	}

	results := resp.ListIpv6networkcontainerResponseObject.GetResult()
	items := make([]*ipam.Ipv6networkcontainer, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSIpv6networkcontainerToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListIpv6networkcontainerResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSIpv6networkcontainerToResponse(r *niosipam.Ipv6networkcontainer) *ipam.Ipv6networkcontainer {
	resp := &ipam.Ipv6networkcontainer{
		Id: r.Ref,
	}
	resp.NIOS = &ipam.NIOSIpv6networkcontainerExt{
		AutoCreateReversezone:            r.AutoCreateReversezone,
		CloudInfo:                        r.CloudInfo,
		Comment:                          r.Comment,
		DdnsDomainname:                   r.DdnsDomainname,
		DdnsEnableOptionFqdn:             r.DdnsEnableOptionFqdn,
		DdnsGenerateHostname:             r.DdnsGenerateHostname,
		DdnsServerAlwaysUpdates:          r.DdnsServerAlwaysUpdates,
		DdnsTtl:                          r.DdnsTtl,
		DeleteReason:                     r.DeleteReason,
		DiscoveryBasicPollSettings:       r.DiscoveryBasicPollSettings,
		DiscoveryBlackoutSetting:         r.DiscoveryBlackoutSetting,
		DiscoveryMember:                  r.DiscoveryMember,
		DomainNameServers:                r.DomainNameServers,
		EnableDdns:                       r.EnableDdns,
		EnableDiscovery:                  r.EnableDiscovery,
		EnableImmediateDiscovery:         r.EnableImmediateDiscovery,
		FederatedRealms:                  r.FederatedRealms,
		LogicFilterRules:                 r.LogicFilterRules,
		MgmPrivate:                       r.MgmPrivate,
		NetworkView:                      r.NetworkView,
		Options:                          r.Options,
		PortControlBlackoutSetting:       r.PortControlBlackoutSetting,
		PreferredLifetime:                r.PreferredLifetime,
		RestartIfNeeded:                  r.RestartIfNeeded,
		RirOrganization:                  r.RirOrganization,
		RirRegistrationAction:            r.RirRegistrationAction,
		RirRegistrationStatus:            r.RirRegistrationStatus,
		SamePortControlDiscoveryBlackout: r.SamePortControlDiscoveryBlackout,
		SendRirRequest:                   r.SendRirRequest,
		SubscribeSettings:                r.SubscribeSettings,
		Unmanaged:                        r.Unmanaged,
		UpdateDnsOnLeaseRenewal:          r.UpdateDnsOnLeaseRenewal,
		UseBlackoutSetting:               r.UseBlackoutSetting,
		UseDdnsDomainname:                r.UseDdnsDomainname,
		UseDdnsEnableOptionFqdn:          r.UseDdnsEnableOptionFqdn,
		UseDdnsGenerateHostname:          r.UseDdnsGenerateHostname,
		UseDdnsTtl:                       r.UseDdnsTtl,
		UseDiscoveryBasicPollingSettings: r.UseDiscoveryBasicPollingSettings,
		UseDomainNameServers:             r.UseDomainNameServers,
		UseEnableDdns:                    r.UseEnableDdns,
		UseEnableDiscovery:               r.UseEnableDiscovery,
		UseLogicFilterRules:              r.UseLogicFilterRules,
		UseMgmPrivate:                    r.UseMgmPrivate,
		UseOptions:                       r.UseOptions,
		UsePreferredLifetime:             r.UsePreferredLifetime,
		UseSubscribeSettings:             r.UseSubscribeSettings,
		UseUpdateDnsOnLeaseRenewal:       r.UseUpdateDnsOnLeaseRenewal,
		UseValidLifetime:                 r.UseValidLifetime,
		UseZoneAssociations:              r.UseZoneAssociations,
		ValidLifetime:                    r.ValidLifetime,
		ZoneAssociations:                 r.ZoneAssociations,
	}
	if r.Network != nil {
		resp.NIOS.Network = r.Network.String
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
