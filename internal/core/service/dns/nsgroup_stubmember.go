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

type NsgroupStubmemberService interface {
	Create(ctx context.Context, obj *dns.NsgroupStubmember, opts *core.Options) (*dns.NsgroupStubmember, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dns.NsgroupStubmember, *http.Response, error)
	Update(ctx context.Context, id string, obj *dns.NsgroupStubmember, opts *core.Options) (*dns.NsgroupStubmember, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dns.NsgroupStubmember, *http.Response, string, error)
}

type nsgroupStubmemberService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewNsgroupStubmemberService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) NsgroupStubmemberService {
	return &nsgroupStubmemberService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new NsgroupStubmember and returns the created object
func (s *nsgroupStubmemberService) Create(ctx context.Context, obj *dns.NsgroupStubmember, opts *core.Options) (*dns.NsgroupStubmember, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *nsgroupStubmemberService) createNIOS(ctx context.Context, obj *dns.NsgroupStubmember, opts *core.Options) (*dns.NsgroupStubmember, *http.Response, error) {
	payload, err := common.MapTo[niosdns.NsgroupStubmember](obj, mapper.NsgroupStubmemberNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.NsgroupStubmemberAPI.
		Create(ctx).
		NsgroupStubmember(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateNsgroupStubmemberResponseAsObject.GetResult()

	return mapNIOSNsgroupStubmemberToResponse(&result), httpResp, nil
}

// Read retrieves a NsgroupStubmember by ID
func (s *nsgroupStubmemberService) Read(ctx context.Context, id string, opts *core.Options) (*dns.NsgroupStubmember, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *nsgroupStubmemberService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dns.NsgroupStubmember, *http.Response, error) {
	req := s.niosClient.DNSAPI.NsgroupStubmemberAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetNsgroupStubmemberResponseObjectAsResult.GetResult()

	return mapNIOSNsgroupStubmemberToResponse(&result), httpResp, nil
}

// Update modifies an existing NsgroupStubmember and returns the updated object
func (s *nsgroupStubmemberService) Update(ctx context.Context, id string, obj *dns.NsgroupStubmember, opts *core.Options) (*dns.NsgroupStubmember, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *nsgroupStubmemberService) updateNIOS(ctx context.Context, id string, obj *dns.NsgroupStubmember, opts *core.Options) (*dns.NsgroupStubmember, *http.Response, error) {
	payload, err := common.MapTo[niosdns.NsgroupStubmember](obj, mapper.NsgroupStubmemberNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.NsgroupStubmemberAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		NsgroupStubmember(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateNsgroupStubmemberResponseAsObject.GetResult()

	return mapNIOSNsgroupStubmemberToResponse(&result), httpResp, nil
}

// Delete removes a NsgroupStubmember by ID
func (s *nsgroupStubmemberService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *nsgroupStubmemberService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DNSAPI.NsgroupStubmemberAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves NsgroupStubmember objects based on filter options
func (s *nsgroupStubmemberService) List(ctx context.Context, opts *core.ListOptions) ([]*dns.NsgroupStubmember, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *nsgroupStubmemberService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dns.NsgroupStubmember, *http.Response, string, error) {
	req := s.niosClient.DNSAPI.NsgroupStubmemberAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.NsgroupStubmemberFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListNsgroupStubmemberResponseObject.GetResult()
	items := make([]*dns.NsgroupStubmember, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSNsgroupStubmemberToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListNsgroupStubmemberResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSNsgroupStubmemberToResponse(r *niosdns.NsgroupStubmember) *dns.NsgroupStubmember {
	resp := &dns.NsgroupStubmember{
		Id: r.Ref,
	}
	resp.NIOS = &dns.NIOSNsgroupStubmemberExt{
		Comment:     r.Comment,
		Name:        r.Name,
		StubMembers: r.StubMembers,
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
