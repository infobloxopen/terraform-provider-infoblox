package grid

import (
	"context"
	"fmt"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niosgrid "github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/grid"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/grid"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
)

type ExtensibleattributedefService interface {
	Create(ctx context.Context, obj *grid.Extensibleattributedef, opts *core.Options) (*grid.Extensibleattributedef, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*grid.Extensibleattributedef, *http.Response, error)
	Update(ctx context.Context, id string, obj *grid.Extensibleattributedef, opts *core.Options) (*grid.Extensibleattributedef, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*grid.Extensibleattributedef, *http.Response, string, error)
}

type extensibleattributedefService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewExtensibleattributedefService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) ExtensibleattributedefService {
	return &extensibleattributedefService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new Extensibleattributedef and returns the created object
func (s *extensibleattributedefService) Create(ctx context.Context, obj *grid.Extensibleattributedef, opts *core.Options) (*grid.Extensibleattributedef, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *extensibleattributedefService) createNIOS(ctx context.Context, obj *grid.Extensibleattributedef, opts *core.Options) (*grid.Extensibleattributedef, *http.Response, error) {
	payload, err := common.MapTo[niosgrid.Extensibleattributedef](obj, mapper.ExtensibleattributedefNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.GridAPI.ExtensibleattributedefAPI.
		Create(ctx).
		Extensibleattributedef(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateExtensibleattributedefResponseAsObject.GetResult()

	return mapNIOSExtensibleattributedefToResponse(&result), httpResp, nil
}

// Read retrieves a Extensibleattributedef by ID
func (s *extensibleattributedefService) Read(ctx context.Context, id string, opts *core.Options) (*grid.Extensibleattributedef, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *extensibleattributedefService) readNIOS(ctx context.Context, id string, opts *core.Options) (*grid.Extensibleattributedef, *http.Response, error) {
	req := s.niosClient.GridAPI.ExtensibleattributedefAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetExtensibleattributedefResponseObjectAsResult.GetResult()

	return mapNIOSExtensibleattributedefToResponse(&result), httpResp, nil
}

// Update modifies an existing Extensibleattributedef and returns the updated object
func (s *extensibleattributedefService) Update(ctx context.Context, id string, obj *grid.Extensibleattributedef, opts *core.Options) (*grid.Extensibleattributedef, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *extensibleattributedefService) updateNIOS(ctx context.Context, id string, obj *grid.Extensibleattributedef, opts *core.Options) (*grid.Extensibleattributedef, *http.Response, error) {
	payload, err := common.MapTo[niosgrid.Extensibleattributedef](obj, mapper.ExtensibleattributedefNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.GridAPI.ExtensibleattributedefAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Extensibleattributedef(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateExtensibleattributedefResponseAsObject.GetResult()

	return mapNIOSExtensibleattributedefToResponse(&result), httpResp, nil
}

// Delete removes a Extensibleattributedef by ID
func (s *extensibleattributedefService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *extensibleattributedefService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.GridAPI.ExtensibleattributedefAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves Extensibleattributedef objects based on filter options
func (s *extensibleattributedefService) List(ctx context.Context, opts *core.ListOptions) ([]*grid.Extensibleattributedef, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *extensibleattributedefService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*grid.Extensibleattributedef, *http.Response, string, error) {
	req := s.niosClient.GridAPI.ExtensibleattributedefAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.ExtensibleattributedefFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListExtensibleattributedefResponseObject.GetResult()
	items := make([]*grid.Extensibleattributedef, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSExtensibleattributedefToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListExtensibleattributedefResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSExtensibleattributedefToResponse(r *niosgrid.Extensibleattributedef) *grid.Extensibleattributedef {
	resp := &grid.Extensibleattributedef{
		Id: r.Ref,
	}
	resp.NIOS = &grid.NIOSExtensibleattributedefExt{
		AllowedObjectTypes: r.AllowedObjectTypes,
		Comment:            r.Comment,
		DefaultValue:       r.DefaultValue,
		DescendantsAction:  r.DescendantsAction,
		Flags:              r.Flags,
		ListValues:         r.ListValues,
		Max:                r.Max,
		Min:                r.Min,
		Name:               r.Name,
		Type:               r.Type,
	}
	return resp
}
