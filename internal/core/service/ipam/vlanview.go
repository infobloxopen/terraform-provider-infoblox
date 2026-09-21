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

type VlanviewService interface {
	Create(ctx context.Context, obj *ipam.Vlanview, opts *core.Options) (*ipam.Vlanview, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*ipam.Vlanview, *http.Response, error)
	Update(ctx context.Context, id string, obj *ipam.Vlanview, opts *core.Options) (*ipam.Vlanview, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*ipam.Vlanview, *http.Response, string, error)
}

type vlanviewService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewVlanviewService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) VlanviewService {
	return &vlanviewService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new Vlanview and returns the created object
func (s *vlanviewService) Create(ctx context.Context, obj *ipam.Vlanview, opts *core.Options) (*ipam.Vlanview, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *vlanviewService) createNIOS(ctx context.Context, obj *ipam.Vlanview, opts *core.Options) (*ipam.Vlanview, *http.Response, error) {
	payload, err := common.MapTo[niosipam.Vlanview](obj, mapper.VlanviewNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.IPAMAPI.VlanviewAPI.
		Create(ctx).
		Vlanview(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateVlanviewResponseAsObject.GetResult()

	return mapNIOSVlanviewToResponse(&result), httpResp, nil
}

// Read retrieves a Vlanview by ID
func (s *vlanviewService) Read(ctx context.Context, id string, opts *core.Options) (*ipam.Vlanview, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *vlanviewService) readNIOS(ctx context.Context, id string, opts *core.Options) (*ipam.Vlanview, *http.Response, error) {
	req := s.niosClient.IPAMAPI.VlanviewAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetVlanviewResponseObjectAsResult.GetResult()

	return mapNIOSVlanviewToResponse(&result), httpResp, nil
}

// Update modifies an existing Vlanview and returns the updated object
func (s *vlanviewService) Update(ctx context.Context, id string, obj *ipam.Vlanview, opts *core.Options) (*ipam.Vlanview, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *vlanviewService) updateNIOS(ctx context.Context, id string, obj *ipam.Vlanview, opts *core.Options) (*ipam.Vlanview, *http.Response, error) {
	payload, err := common.MapTo[niosipam.Vlanview](obj, mapper.VlanviewNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.IPAMAPI.VlanviewAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Vlanview(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateVlanviewResponseAsObject.GetResult()

	return mapNIOSVlanviewToResponse(&result), httpResp, nil
}

// Delete removes a Vlanview by ID
func (s *vlanviewService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *vlanviewService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.IPAMAPI.VlanviewAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves Vlanview objects based on filter options
func (s *vlanviewService) List(ctx context.Context, opts *core.ListOptions) ([]*ipam.Vlanview, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *vlanviewService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*ipam.Vlanview, *http.Response, string, error) {
	req := s.niosClient.IPAMAPI.VlanviewAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.VlanviewFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListVlanviewResponseObject.GetResult()
	items := make([]*ipam.Vlanview, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSVlanviewToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListVlanviewResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSVlanviewToResponse(r *niosipam.Vlanview) *ipam.Vlanview {
	resp := &ipam.Vlanview{
		Id: r.Ref,
	}
	resp.NIOS = &ipam.NIOSVlanviewExt{
		AllowRangeOverlapping: r.AllowRangeOverlapping,
		Comment:               r.Comment,
		EndVlanId:             r.EndVlanId,
		Name:                  r.Name,
		PreCreateVlan:         r.PreCreateVlan,
		StartVlanId:           r.StartVlanId,
		VlanNamePrefix:        r.VlanNamePrefix,
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
