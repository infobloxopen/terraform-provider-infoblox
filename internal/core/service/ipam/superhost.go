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

type SuperhostService interface {
	Create(ctx context.Context, obj *ipam.Superhost, opts *core.Options) (*ipam.Superhost, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*ipam.Superhost, *http.Response, error)
	Update(ctx context.Context, id string, obj *ipam.Superhost, opts *core.Options) (*ipam.Superhost, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*ipam.Superhost, *http.Response, string, error)
}

type superhostService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewSuperhostService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) SuperhostService {
	return &superhostService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new Superhost and returns the created object
func (s *superhostService) Create(ctx context.Context, obj *ipam.Superhost, opts *core.Options) (*ipam.Superhost, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *superhostService) createNIOS(ctx context.Context, obj *ipam.Superhost, opts *core.Options) (*ipam.Superhost, *http.Response, error) {
	payload, err := common.MapTo[niosipam.Superhost](obj, mapper.SuperhostNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.IPAMAPI.SuperhostAPI.
		Create(ctx).
		Superhost(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateSuperhostResponseAsObject.GetResult()

	return mapNIOSSuperhostToResponse(&result), httpResp, nil
}

// Read retrieves a Superhost by ID
func (s *superhostService) Read(ctx context.Context, id string, opts *core.Options) (*ipam.Superhost, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *superhostService) readNIOS(ctx context.Context, id string, opts *core.Options) (*ipam.Superhost, *http.Response, error) {
	req := s.niosClient.IPAMAPI.SuperhostAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetSuperhostResponseObjectAsResult.GetResult()

	return mapNIOSSuperhostToResponse(&result), httpResp, nil
}

// Update modifies an existing Superhost and returns the updated object
func (s *superhostService) Update(ctx context.Context, id string, obj *ipam.Superhost, opts *core.Options) (*ipam.Superhost, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *superhostService) updateNIOS(ctx context.Context, id string, obj *ipam.Superhost, opts *core.Options) (*ipam.Superhost, *http.Response, error) {
	payload, err := common.MapTo[niosipam.Superhost](obj, mapper.SuperhostNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.IPAMAPI.SuperhostAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Superhost(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateSuperhostResponseAsObject.GetResult()

	return mapNIOSSuperhostToResponse(&result), httpResp, nil
}

// Delete removes a Superhost by ID
func (s *superhostService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *superhostService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.IPAMAPI.SuperhostAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves Superhost objects based on filter options
func (s *superhostService) List(ctx context.Context, opts *core.ListOptions) ([]*ipam.Superhost, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *superhostService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*ipam.Superhost, *http.Response, string, error) {
	req := s.niosClient.IPAMAPI.SuperhostAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.SuperhostFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListSuperhostResponseObject.GetResult()
	items := make([]*ipam.Superhost, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSSuperhostToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListSuperhostResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSSuperhostToResponse(r *niosipam.Superhost) *ipam.Superhost {
	resp := &ipam.Superhost{
		Id: r.Ref,
	}
	resp.NIOS = &ipam.NIOSSuperhostExt{
		Comment:                 r.Comment,
		DeleteAssociatedObjects: r.DeleteAssociatedObjects,
		DhcpAssociatedObjects:   r.DhcpAssociatedObjects,
		Disabled:                r.Disabled,
		DnsAssociatedObjects:    r.DnsAssociatedObjects,
		Name:                    r.Name,
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
