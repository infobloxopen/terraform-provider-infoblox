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

type VlanrangeService interface {
	Create(ctx context.Context, obj *ipam.Vlanrange, opts *core.Options) (*ipam.Vlanrange, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*ipam.Vlanrange, *http.Response, error)
	Update(ctx context.Context, id string, obj *ipam.Vlanrange, opts *core.Options) (*ipam.Vlanrange, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*ipam.Vlanrange, *http.Response, string, error)
}

type vlanrangeService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewVlanrangeService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) VlanrangeService {
	return &vlanrangeService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new Vlanrange and returns the created object
func (s *vlanrangeService) Create(ctx context.Context, obj *ipam.Vlanrange, opts *core.Options) (*ipam.Vlanrange, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *vlanrangeService) createNIOS(ctx context.Context, obj *ipam.Vlanrange, opts *core.Options) (*ipam.Vlanrange, *http.Response, error) {
	payload, err := common.MapTo[niosipam.Vlanrange](obj, mapper.VlanrangeNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.IPAMAPI.VlanrangeAPI.
		Create(ctx).
		Vlanrange(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateVlanrangeResponseAsObject.GetResult()

	return mapNIOSVlanrangeToResponse(&result), httpResp, nil
}

// Read retrieves a Vlanrange by ID
func (s *vlanrangeService) Read(ctx context.Context, id string, opts *core.Options) (*ipam.Vlanrange, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *vlanrangeService) readNIOS(ctx context.Context, id string, opts *core.Options) (*ipam.Vlanrange, *http.Response, error) {
	req := s.niosClient.IPAMAPI.VlanrangeAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetVlanrangeResponseObjectAsResult.GetResult()

	return mapNIOSVlanrangeToResponse(&result), httpResp, nil
}

// Update modifies an existing Vlanrange and returns the updated object
func (s *vlanrangeService) Update(ctx context.Context, id string, obj *ipam.Vlanrange, opts *core.Options) (*ipam.Vlanrange, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *vlanrangeService) updateNIOS(ctx context.Context, id string, obj *ipam.Vlanrange, opts *core.Options) (*ipam.Vlanrange, *http.Response, error) {
	payload, err := common.MapTo[niosipam.Vlanrange](obj, mapper.VlanrangeNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.IPAMAPI.VlanrangeAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Vlanrange(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateVlanrangeResponseAsObject.GetResult()

	return mapNIOSVlanrangeToResponse(&result), httpResp, nil
}

// Delete removes a Vlanrange by ID
func (s *vlanrangeService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *vlanrangeService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.IPAMAPI.VlanrangeAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves Vlanrange objects based on filter options
func (s *vlanrangeService) List(ctx context.Context, opts *core.ListOptions) ([]*ipam.Vlanrange, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *vlanrangeService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*ipam.Vlanrange, *http.Response, string, error) {
	req := s.niosClient.IPAMAPI.VlanrangeAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.VlanrangeFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListVlanrangeResponseObject.GetResult()
	items := make([]*ipam.Vlanrange, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSVlanrangeToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListVlanrangeResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSVlanrangeToResponse(r *niosipam.Vlanrange) *ipam.Vlanrange {
	resp := &ipam.Vlanrange{
		Id: r.Ref,
	}
	resp.NIOS = &ipam.NIOSVlanrangeExt{
		Comment:        r.Comment,
		DeleteVlans:    r.DeleteVlans,
		EndVlanId:      r.EndVlanId,
		Name:           r.Name,
		PreCreateVlan:  r.PreCreateVlan,
		StartVlanId:    r.StartVlanId,
		VlanNamePrefix: r.VlanNamePrefix,
		VlanView:       r.VlanView,
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
