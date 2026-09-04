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

type NsgroupForwardstubserverService interface {
	Create(ctx context.Context, obj *dns.NsgroupForwardstubserver, opts *core.Options) (*dns.NsgroupForwardstubserver, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dns.NsgroupForwardstubserver, *http.Response, error)
	Update(ctx context.Context, id string, obj *dns.NsgroupForwardstubserver, opts *core.Options) (*dns.NsgroupForwardstubserver, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dns.NsgroupForwardstubserver, *http.Response, string, error)
}

type nsgroupForwardstubserverService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewNsgroupForwardstubserverService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) NsgroupForwardstubserverService {
	return &nsgroupForwardstubserverService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new NsgroupForwardstubserver and returns the created object
func (s *nsgroupForwardstubserverService) Create(ctx context.Context, obj *dns.NsgroupForwardstubserver, opts *core.Options) (*dns.NsgroupForwardstubserver, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *nsgroupForwardstubserverService) createNIOS(ctx context.Context, obj *dns.NsgroupForwardstubserver, opts *core.Options) (*dns.NsgroupForwardstubserver, *http.Response, error) {
	payload, err := common.MapTo[niosdns.NsgroupForwardstubserver](obj, mapper.NsgroupForwardstubserverNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.NsgroupForwardstubserverAPI.
		Create(ctx).
		NsgroupForwardstubserver(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateNsgroupForwardstubserverResponseAsObject.GetResult()

	return mapNIOSNsgroupForwardstubserverToResponse(&result), httpResp, nil
}

// Read retrieves a NsgroupForwardstubserver by ID
func (s *nsgroupForwardstubserverService) Read(ctx context.Context, id string, opts *core.Options) (*dns.NsgroupForwardstubserver, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *nsgroupForwardstubserverService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dns.NsgroupForwardstubserver, *http.Response, error) {
	req := s.niosClient.DNSAPI.NsgroupForwardstubserverAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetNsgroupForwardstubserverResponseObjectAsResult.GetResult()

	return mapNIOSNsgroupForwardstubserverToResponse(&result), httpResp, nil
}

// Update modifies an existing NsgroupForwardstubserver and returns the updated object
func (s *nsgroupForwardstubserverService) Update(ctx context.Context, id string, obj *dns.NsgroupForwardstubserver, opts *core.Options) (*dns.NsgroupForwardstubserver, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *nsgroupForwardstubserverService) updateNIOS(ctx context.Context, id string, obj *dns.NsgroupForwardstubserver, opts *core.Options) (*dns.NsgroupForwardstubserver, *http.Response, error) {
	payload, err := common.MapTo[niosdns.NsgroupForwardstubserver](obj, mapper.NsgroupForwardstubserverNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.NsgroupForwardstubserverAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		NsgroupForwardstubserver(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateNsgroupForwardstubserverResponseAsObject.GetResult()

	return mapNIOSNsgroupForwardstubserverToResponse(&result), httpResp, nil
}

// Delete removes a NsgroupForwardstubserver by ID
func (s *nsgroupForwardstubserverService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *nsgroupForwardstubserverService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DNSAPI.NsgroupForwardstubserverAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves NsgroupForwardstubserver objects based on filter options
func (s *nsgroupForwardstubserverService) List(ctx context.Context, opts *core.ListOptions) ([]*dns.NsgroupForwardstubserver, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *nsgroupForwardstubserverService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dns.NsgroupForwardstubserver, *http.Response, string, error) {
	req := s.niosClient.DNSAPI.NsgroupForwardstubserverAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.NsgroupForwardstubserverFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListNsgroupForwardstubserverResponseObject.GetResult()
	items := make([]*dns.NsgroupForwardstubserver, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSNsgroupForwardstubserverToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListNsgroupForwardstubserverResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSNsgroupForwardstubserverToResponse(r *niosdns.NsgroupForwardstubserver) *dns.NsgroupForwardstubserver {
	resp := &dns.NsgroupForwardstubserver{
		Id: r.Ref,
	}
	resp.NIOS = &dns.NIOSNsgroupForwardstubserverExt{
		Comment:         r.Comment,
		ExternalServers: r.ExternalServers,
		Name:            r.Name,
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
