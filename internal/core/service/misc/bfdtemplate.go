package misc

import (
	"context"
	"fmt"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niosmisc "github.com/infobloxopen/infoblox-nios-go-client/misc"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/misc"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/misc"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
)

type BfdtemplateService interface {
	Create(ctx context.Context, obj *misc.Bfdtemplate, opts *core.Options) (*misc.Bfdtemplate, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*misc.Bfdtemplate, *http.Response, error)
	Update(ctx context.Context, id string, obj *misc.Bfdtemplate, opts *core.Options) (*misc.Bfdtemplate, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*misc.Bfdtemplate, *http.Response, string, error)
}

type bfdtemplateService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewBfdtemplateService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) BfdtemplateService {
	return &bfdtemplateService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new Bfdtemplate and returns the created object
func (s *bfdtemplateService) Create(ctx context.Context, obj *misc.Bfdtemplate, opts *core.Options) (*misc.Bfdtemplate, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *bfdtemplateService) createNIOS(ctx context.Context, obj *misc.Bfdtemplate, opts *core.Options) (*misc.Bfdtemplate, *http.Response, error) {
	payload, err := common.MapTo[niosmisc.Bfdtemplate](obj, mapper.BfdtemplateNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.MiscAPI.BfdtemplateAPI.
		Create(ctx).
		Bfdtemplate(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateBfdtemplateResponseAsObject.GetResult()

	return mapNIOSBfdtemplateToResponse(&result), httpResp, nil
}

// Read retrieves a Bfdtemplate by ID
func (s *bfdtemplateService) Read(ctx context.Context, id string, opts *core.Options) (*misc.Bfdtemplate, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *bfdtemplateService) readNIOS(ctx context.Context, id string, opts *core.Options) (*misc.Bfdtemplate, *http.Response, error) {
	req := s.niosClient.MiscAPI.BfdtemplateAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetBfdtemplateResponseObjectAsResult.GetResult()

	return mapNIOSBfdtemplateToResponse(&result), httpResp, nil
}

// Update modifies an existing Bfdtemplate and returns the updated object
func (s *bfdtemplateService) Update(ctx context.Context, id string, obj *misc.Bfdtemplate, opts *core.Options) (*misc.Bfdtemplate, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *bfdtemplateService) updateNIOS(ctx context.Context, id string, obj *misc.Bfdtemplate, opts *core.Options) (*misc.Bfdtemplate, *http.Response, error) {
	payload, err := common.MapTo[niosmisc.Bfdtemplate](obj, mapper.BfdtemplateNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.MiscAPI.BfdtemplateAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Bfdtemplate(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateBfdtemplateResponseAsObject.GetResult()

	return mapNIOSBfdtemplateToResponse(&result), httpResp, nil
}

// Delete removes a Bfdtemplate by ID
func (s *bfdtemplateService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *bfdtemplateService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.MiscAPI.BfdtemplateAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves Bfdtemplate objects based on filter options
func (s *bfdtemplateService) List(ctx context.Context, opts *core.ListOptions) ([]*misc.Bfdtemplate, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *bfdtemplateService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*misc.Bfdtemplate, *http.Response, string, error) {
	req := s.niosClient.MiscAPI.BfdtemplateAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.BfdtemplateFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListBfdtemplateResponseObject.GetResult()
	items := make([]*misc.Bfdtemplate, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSBfdtemplateToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListBfdtemplateResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSBfdtemplateToResponse(r *niosmisc.Bfdtemplate) *misc.Bfdtemplate {
	resp := &misc.Bfdtemplate{
		Id: r.Ref,
	}
	resp.NIOS = &misc.NIOSBfdtemplateExt{
		AuthenticationKey:   r.AuthenticationKey,
		AuthenticationKeyId: r.AuthenticationKeyId,
		AuthenticationType:  r.AuthenticationType,
		DetectionMultiplier: r.DetectionMultiplier,
		MinRxInterval:       r.MinRxInterval,
		MinTxInterval:       r.MinTxInterval,
		Name:                r.Name,
	}
	return resp
}
