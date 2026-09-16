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

type NsgroupForwardingmemberService interface {
	Create(ctx context.Context, obj *dns.NsgroupForwardingmember, opts *core.Options) (*dns.NsgroupForwardingmember, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dns.NsgroupForwardingmember, *http.Response, error)
	Update(ctx context.Context, id string, obj *dns.NsgroupForwardingmember, opts *core.Options) (*dns.NsgroupForwardingmember, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dns.NsgroupForwardingmember, *http.Response, string, error)
}

type nsgroupForwardingmemberService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewNsgroupForwardingmemberService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) NsgroupForwardingmemberService {
	return &nsgroupForwardingmemberService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new NsgroupForwardingmember and returns the created object
func (s *nsgroupForwardingmemberService) Create(ctx context.Context, obj *dns.NsgroupForwardingmember, opts *core.Options) (*dns.NsgroupForwardingmember, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *nsgroupForwardingmemberService) createNIOS(ctx context.Context, obj *dns.NsgroupForwardingmember, opts *core.Options) (*dns.NsgroupForwardingmember, *http.Response, error) {
	payload, err := common.MapTo[niosdns.NsgroupForwardingmember](obj, mapper.NsgroupForwardingmemberNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.NsgroupForwardingmemberAPI.
		Create(ctx).
		NsgroupForwardingmember(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateNsgroupForwardingmemberResponseAsObject.GetResult()

	return mapNIOSNsgroupForwardingmemberToResponse(&result), httpResp, nil
}

// Read retrieves a NsgroupForwardingmember by ID
func (s *nsgroupForwardingmemberService) Read(ctx context.Context, id string, opts *core.Options) (*dns.NsgroupForwardingmember, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *nsgroupForwardingmemberService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dns.NsgroupForwardingmember, *http.Response, error) {
	req := s.niosClient.DNSAPI.NsgroupForwardingmemberAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetNsgroupForwardingmemberResponseObjectAsResult.GetResult()

	return mapNIOSNsgroupForwardingmemberToResponse(&result), httpResp, nil
}

// Update modifies an existing NsgroupForwardingmember and returns the updated object
func (s *nsgroupForwardingmemberService) Update(ctx context.Context, id string, obj *dns.NsgroupForwardingmember, opts *core.Options) (*dns.NsgroupForwardingmember, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *nsgroupForwardingmemberService) updateNIOS(ctx context.Context, id string, obj *dns.NsgroupForwardingmember, opts *core.Options) (*dns.NsgroupForwardingmember, *http.Response, error) {
	payload, err := common.MapTo[niosdns.NsgroupForwardingmember](obj, mapper.NsgroupForwardingmemberNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.NsgroupForwardingmemberAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		NsgroupForwardingmember(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateNsgroupForwardingmemberResponseAsObject.GetResult()

	return mapNIOSNsgroupForwardingmemberToResponse(&result), httpResp, nil
}

// Delete removes a NsgroupForwardingmember by ID
func (s *nsgroupForwardingmemberService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *nsgroupForwardingmemberService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DNSAPI.NsgroupForwardingmemberAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves NsgroupForwardingmember objects based on filter options
func (s *nsgroupForwardingmemberService) List(ctx context.Context, opts *core.ListOptions) ([]*dns.NsgroupForwardingmember, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *nsgroupForwardingmemberService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dns.NsgroupForwardingmember, *http.Response, string, error) {
	req := s.niosClient.DNSAPI.NsgroupForwardingmemberAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.NsgroupForwardingmemberFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListNsgroupForwardingmemberResponseObject.GetResult()
	items := make([]*dns.NsgroupForwardingmember, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSNsgroupForwardingmemberToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListNsgroupForwardingmemberResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSNsgroupForwardingmemberToResponse(r *niosdns.NsgroupForwardingmember) *dns.NsgroupForwardingmember {
	resp := &dns.NsgroupForwardingmember{
		Id: r.Ref,
	}
	resp.NIOS = &dns.NIOSNsgroupForwardingmemberExt{
		Comment:           r.Comment,
		ForwardingServers: r.ForwardingServers,
		Name:              r.Name,
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
