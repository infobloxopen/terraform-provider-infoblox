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

type BulkhostnametemplateService interface {
	Create(ctx context.Context, obj *ipam.Bulkhostnametemplate, opts *core.Options) (*ipam.Bulkhostnametemplate, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*ipam.Bulkhostnametemplate, *http.Response, error)
	Update(ctx context.Context, id string, obj *ipam.Bulkhostnametemplate, opts *core.Options) (*ipam.Bulkhostnametemplate, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*ipam.Bulkhostnametemplate, *http.Response, string, error)
}

type bulkhostnametemplateService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewBulkhostnametemplateService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) BulkhostnametemplateService {
	return &bulkhostnametemplateService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new Bulkhostnametemplate and returns the created object
func (s *bulkhostnametemplateService) Create(ctx context.Context, obj *ipam.Bulkhostnametemplate, opts *core.Options) (*ipam.Bulkhostnametemplate, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *bulkhostnametemplateService) createNIOS(ctx context.Context, obj *ipam.Bulkhostnametemplate, opts *core.Options) (*ipam.Bulkhostnametemplate, *http.Response, error) {
	payload, err := common.MapTo[niosipam.Bulkhostnametemplate](obj, mapper.BulkhostnametemplateNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.IPAMAPI.BulkhostnametemplateAPI.
		Create(ctx).
		Bulkhostnametemplate(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateBulkhostnametemplateResponseAsObject.GetResult()

	return mapNIOSBulkhostnametemplateToResponse(&result), httpResp, nil
}

// Read retrieves a Bulkhostnametemplate by ID
func (s *bulkhostnametemplateService) Read(ctx context.Context, id string, opts *core.Options) (*ipam.Bulkhostnametemplate, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *bulkhostnametemplateService) readNIOS(ctx context.Context, id string, opts *core.Options) (*ipam.Bulkhostnametemplate, *http.Response, error) {
	req := s.niosClient.IPAMAPI.BulkhostnametemplateAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetBulkhostnametemplateResponseObjectAsResult.GetResult()

	return mapNIOSBulkhostnametemplateToResponse(&result), httpResp, nil
}

// Update modifies an existing Bulkhostnametemplate and returns the updated object
func (s *bulkhostnametemplateService) Update(ctx context.Context, id string, obj *ipam.Bulkhostnametemplate, opts *core.Options) (*ipam.Bulkhostnametemplate, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *bulkhostnametemplateService) updateNIOS(ctx context.Context, id string, obj *ipam.Bulkhostnametemplate, opts *core.Options) (*ipam.Bulkhostnametemplate, *http.Response, error) {
	payload, err := common.MapTo[niosipam.Bulkhostnametemplate](obj, mapper.BulkhostnametemplateNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.IPAMAPI.BulkhostnametemplateAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Bulkhostnametemplate(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateBulkhostnametemplateResponseAsObject.GetResult()

	return mapNIOSBulkhostnametemplateToResponse(&result), httpResp, nil
}

// Delete removes a Bulkhostnametemplate by ID
func (s *bulkhostnametemplateService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *bulkhostnametemplateService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.IPAMAPI.BulkhostnametemplateAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves Bulkhostnametemplate objects based on filter options
func (s *bulkhostnametemplateService) List(ctx context.Context, opts *core.ListOptions) ([]*ipam.Bulkhostnametemplate, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *bulkhostnametemplateService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*ipam.Bulkhostnametemplate, *http.Response, string, error) {
	req := s.niosClient.IPAMAPI.BulkhostnametemplateAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.BulkhostnametemplateFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListBulkhostnametemplateResponseObject.GetResult()
	items := make([]*ipam.Bulkhostnametemplate, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSBulkhostnametemplateToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListBulkhostnametemplateResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSBulkhostnametemplateToResponse(r *niosipam.Bulkhostnametemplate) *ipam.Bulkhostnametemplate {
	resp := &ipam.Bulkhostnametemplate{
		Id: r.Ref,
	}
	resp.NIOS = &ipam.NIOSBulkhostnametemplateExt{
		TemplateFormat: r.TemplateFormat,
		TemplateName:   r.TemplateName,
	}
	return resp
}
