package dns

import (
	"context"
	"fmt"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
)

type RecordHostService interface {
	Create(ctx context.Context, obj *dns.RecordHost, opts *core.Options) (*dns.RecordHost, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dns.RecordHost, *http.Response, error)
	Update(ctx context.Context, id string, obj *dns.RecordHost, opts *core.Options) (*dns.RecordHost, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dns.RecordHost, *http.Response, string, error)
}

type recordHostService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewRecordHostService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) RecordHostService {
	return &recordHostService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new RecordHost and returns the created object
func (s *recordHostService) Create(ctx context.Context, obj *dns.RecordHost, opts *core.Options) (*dns.RecordHost, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordHostService) createNIOS(ctx context.Context, obj *dns.RecordHost, opts *core.Options) (*dns.RecordHost, *http.Response, error) {
	payload, err := common.MapTo[niosdns.RecordHost](obj, mapper.RecordHostNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.RecordHostAPI.
		Create(ctx).
		RecordHost(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateRecordHostResponseAsObject.GetResult()

	return mapNIOSRecordHostToResponse(&result), httpResp, nil
}

// Read retrieves a RecordHost by ID
func (s *recordHostService) Read(ctx context.Context, id string, opts *core.Options) (*dns.RecordHost, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordHostService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dns.RecordHost, *http.Response, error) {
	req := s.niosClient.DNSAPI.RecordHostAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetRecordHostResponseObjectAsResult.GetResult()

	return mapNIOSRecordHostToResponse(&result), httpResp, nil
}

// Update modifies an existing RecordHost and returns the updated object
func (s *recordHostService) Update(ctx context.Context, id string, obj *dns.RecordHost, opts *core.Options) (*dns.RecordHost, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordHostService) updateNIOS(ctx context.Context, id string, obj *dns.RecordHost, opts *core.Options) (*dns.RecordHost, *http.Response, error) {
	payload, err := common.MapTo[niosdns.RecordHost](obj, mapper.RecordHostNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.RecordHostAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		RecordHost(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateRecordHostResponseAsObject.GetResult()

	return mapNIOSRecordHostToResponse(&result), httpResp, nil
}

// Delete removes a RecordHost by ID
func (s *recordHostService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordHostService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DNSAPI.RecordHostAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves RecordHost objects based on filter options
func (s *recordHostService) List(ctx context.Context, opts *core.ListOptions) ([]*dns.RecordHost, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordHostService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dns.RecordHost, *http.Response, string, error) {
	req := s.niosClient.DNSAPI.RecordHostAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.RecordHostFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListRecordHostResponseObject.GetResult()
	items := make([]*dns.RecordHost, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSRecordHostToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListRecordHostResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSRecordHostToResponse(r *niosdns.RecordHost) *dns.RecordHost {
	resp := &dns.RecordHost{
		Id: r.Ref,
	}
	resp.NIOS = &dns.NIOSRecordHostExt{
		Aliases:                  r.Aliases,
		AllowTelnet:              r.AllowTelnet,
		CliCredentials:           r.CliCredentials,
		CloudInfo:                r.CloudInfo,
		Comment:                  r.Comment,
		ConfigureForDns:          r.ConfigureForDns,
		DdnsProtected:            r.DdnsProtected,
		DeviceDescription:        r.DeviceDescription,
		DeviceLocation:           r.DeviceLocation,
		DeviceType:               r.DeviceType,
		DeviceVendor:             r.DeviceVendor,
		Disable:                  r.Disable,
		DisableDiscovery:         r.DisableDiscovery,
		DnsAliases:               r.DnsAliases,
		EnableImmediateDiscovery: r.EnableImmediateDiscovery,
		Ipv4addrs:                r.Ipv4addrs,
		Ipv6addrs:                r.Ipv6addrs,
		Name:                     r.Name,
		NetworkView:              r.NetworkView,
		RestartIfNeeded:          r.RestartIfNeeded,
		RrsetOrder:               r.RrsetOrder,
		Snmp3Credential:          r.Snmp3Credential,
		SnmpCredential:           r.SnmpCredential,
		Ttl:                      r.Ttl,
		UseCliCredentials:        r.UseCliCredentials,
		UseDnsEaInheritance:      r.UseDnsEaInheritance,
		UseSnmp3Credential:       r.UseSnmp3Credential,
		UseSnmpCredential:        r.UseSnmpCredential,
		UseTtl:                   r.UseTtl,
		View:                     r.View,
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
