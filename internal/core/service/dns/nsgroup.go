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

type NsgroupService interface {
	Create(ctx context.Context, obj *dns.Nsgroup, opts *core.Options) (*dns.Nsgroup, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dns.Nsgroup, *http.Response, error)
	Update(ctx context.Context, id string, obj *dns.Nsgroup, opts *core.Options) (*dns.Nsgroup, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dns.Nsgroup, *http.Response, string, error)
}

type nsgroupService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewNsgroupService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) NsgroupService {
	return &nsgroupService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new Nsgroup and returns the created object
func (s *nsgroupService) Create(ctx context.Context, obj *dns.Nsgroup, opts *core.Options) (*dns.Nsgroup, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *nsgroupService) createNIOS(ctx context.Context, obj *dns.Nsgroup, opts *core.Options) (*dns.Nsgroup, *http.Response, error) {
	payload, err := common.MapTo[niosdns.Nsgroup](obj, mapper.NsgroupNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.NsgroupAPI.
		Create(ctx).
		Nsgroup(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateNsgroupResponseAsObject.GetResult()

	return mapNIOSNsgroupToResponse(&result), httpResp, nil
}

// Read retrieves a Nsgroup by ID
func (s *nsgroupService) Read(ctx context.Context, id string, opts *core.Options) (*dns.Nsgroup, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *nsgroupService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dns.Nsgroup, *http.Response, error) {
	req := s.niosClient.DNSAPI.NsgroupAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetNsgroupResponseObjectAsResult.GetResult()

	return mapNIOSNsgroupToResponse(&result), httpResp, nil
}

// Update modifies an existing Nsgroup and returns the updated object
func (s *nsgroupService) Update(ctx context.Context, id string, obj *dns.Nsgroup, opts *core.Options) (*dns.Nsgroup, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *nsgroupService) updateNIOS(ctx context.Context, id string, obj *dns.Nsgroup, opts *core.Options) (*dns.Nsgroup, *http.Response, error) {
	payload, err := common.MapTo[niosdns.Nsgroup](obj, mapper.NsgroupNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.NsgroupAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Nsgroup(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateNsgroupResponseAsObject.GetResult()

	return mapNIOSNsgroupToResponse(&result), httpResp, nil
}

// Delete removes a Nsgroup by ID
func (s *nsgroupService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *nsgroupService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DNSAPI.NsgroupAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves Nsgroup objects based on filter options
func (s *nsgroupService) List(ctx context.Context, opts *core.ListOptions) ([]*dns.Nsgroup, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *nsgroupService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dns.Nsgroup, *http.Response, string, error) {
	req := s.niosClient.DNSAPI.NsgroupAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.NsgroupFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListNsgroupResponseObject.GetResult()
	items := make([]*dns.Nsgroup, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSNsgroupToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListNsgroupResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSNsgroupToResponse(r *niosdns.Nsgroup) *dns.Nsgroup {
	resp := &dns.Nsgroup{
		Id: r.Ref,
	}
	resp.NIOS = &dns.NIOSNsgroupExt{
		Comment:             r.Comment,
		ExternalPrimaries:   r.ExternalPrimaries,
		ExternalSecondaries: r.ExternalSecondaries,
		GridPrimary:         r.GridPrimary,
		GridSecondaries:     r.GridSecondaries,
		IsGridDefault:       r.IsGridDefault,
		IsMultimaster:       r.IsMultimaster,
		Name:                r.Name,
		UseExternalPrimary:  r.UseExternalPrimary,
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
