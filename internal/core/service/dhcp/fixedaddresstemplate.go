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

type FixedaddresstemplateService interface {
	Create(ctx context.Context, obj *dhcp.Fixedaddresstemplate, opts *core.Options) (*dhcp.Fixedaddresstemplate, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dhcp.Fixedaddresstemplate, *http.Response, error)
	Update(ctx context.Context, id string, obj *dhcp.Fixedaddresstemplate, opts *core.Options) (*dhcp.Fixedaddresstemplate, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Fixedaddresstemplate, *http.Response, string, error)
}

type fixedaddresstemplateService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewFixedaddresstemplateService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) FixedaddresstemplateService {
	return &fixedaddresstemplateService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new Fixedaddresstemplate and returns the created object
func (s *fixedaddresstemplateService) Create(ctx context.Context, obj *dhcp.Fixedaddresstemplate, opts *core.Options) (*dhcp.Fixedaddresstemplate, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *fixedaddresstemplateService) createNIOS(ctx context.Context, obj *dhcp.Fixedaddresstemplate, opts *core.Options) (*dhcp.Fixedaddresstemplate, *http.Response, error) {
	payload, err := common.MapTo[niosdhcp.Fixedaddresstemplate](obj, mapper.FixedaddresstemplateNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DHCPAPI.FixedaddresstemplateAPI.
		Create(ctx).
		Fixedaddresstemplate(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateFixedaddresstemplateResponseAsObject.GetResult()

	return mapNIOSFixedaddresstemplateToResponse(&result), httpResp, nil
}

// Read retrieves a Fixedaddresstemplate by ID
func (s *fixedaddresstemplateService) Read(ctx context.Context, id string, opts *core.Options) (*dhcp.Fixedaddresstemplate, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *fixedaddresstemplateService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dhcp.Fixedaddresstemplate, *http.Response, error) {
	req := s.niosClient.DHCPAPI.FixedaddresstemplateAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	req = req.ProxySearch(core.GetProxySearch())
	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetFixedaddresstemplateResponseObjectAsResult.GetResult()

	return mapNIOSFixedaddresstemplateToResponse(&result), httpResp, nil
}

// Update modifies an existing Fixedaddresstemplate and returns the updated object
func (s *fixedaddresstemplateService) Update(ctx context.Context, id string, obj *dhcp.Fixedaddresstemplate, opts *core.Options) (*dhcp.Fixedaddresstemplate, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *fixedaddresstemplateService) updateNIOS(ctx context.Context, id string, obj *dhcp.Fixedaddresstemplate, opts *core.Options) (*dhcp.Fixedaddresstemplate, *http.Response, error) {
	payload, err := common.MapTo[niosdhcp.Fixedaddresstemplate](obj, mapper.FixedaddresstemplateNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DHCPAPI.FixedaddresstemplateAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Fixedaddresstemplate(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateFixedaddresstemplateResponseAsObject.GetResult()

	return mapNIOSFixedaddresstemplateToResponse(&result), httpResp, nil
}

// Delete removes a Fixedaddresstemplate by ID
func (s *fixedaddresstemplateService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *fixedaddresstemplateService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DHCPAPI.FixedaddresstemplateAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves Fixedaddresstemplate objects based on filter options
func (s *fixedaddresstemplateService) List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Fixedaddresstemplate, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *fixedaddresstemplateService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dhcp.Fixedaddresstemplate, *http.Response, string, error) {
	req := s.niosClient.DHCPAPI.FixedaddresstemplateAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.FixedaddresstemplateFilterFieldMap[core.BackendNIOS])
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

	req = req.ProxySearch(core.GetProxySearch())
	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, "", err
	}

	results := resp.ListFixedaddresstemplateResponseObject.GetResult()
	items := make([]*dhcp.Fixedaddresstemplate, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSFixedaddresstemplateToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListFixedaddresstemplateResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSFixedaddresstemplateToResponse(r *niosdhcp.Fixedaddresstemplate) *dhcp.Fixedaddresstemplate {
	resp := &dhcp.Fixedaddresstemplate{
		Id: r.Ref,
	}
	resp.NIOS = &dhcp.NIOSFixedaddresstemplateExt{
		Bootfile:                       r.Bootfile,
		Bootserver:                     r.Bootserver,
		Comment:                        r.Comment,
		DdnsDomainname:                 r.DdnsDomainname,
		DdnsHostname:                   r.DdnsHostname,
		DenyBootp:                      r.DenyBootp,
		EnableDdns:                     r.EnableDdns,
		EnablePxeLeaseTime:             r.EnablePxeLeaseTime,
		IgnoreDhcpOptionListRequest:    r.IgnoreDhcpOptionListRequest,
		LogicFilterRules:               r.LogicFilterRules,
		Name:                           r.Name,
		Nextserver:                     r.Nextserver,
		NumberOfAddresses:              r.NumberOfAddresses,
		Offset:                         r.Offset,
		Options:                        r.Options,
		PxeLeaseTime:                   r.PxeLeaseTime,
		UseBootfile:                    r.UseBootfile,
		UseBootserver:                  r.UseBootserver,
		UseDdnsDomainname:              r.UseDdnsDomainname,
		UseDenyBootp:                   r.UseDenyBootp,
		UseEnableDdns:                  r.UseEnableDdns,
		UseIgnoreDhcpOptionListRequest: r.UseIgnoreDhcpOptionListRequest,
		UseLogicFilterRules:            r.UseLogicFilterRules,
		UseNextserver:                  r.UseNextserver,
		UseOptions:                     r.UseOptions,
		UsePxeLeaseTime:                r.UsePxeLeaseTime,
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
