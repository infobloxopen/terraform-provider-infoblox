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

type VlanService interface {
	Create(ctx context.Context, obj *ipam.Vlan, opts *core.Options) (*ipam.Vlan, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*ipam.Vlan, *http.Response, error)
	Update(ctx context.Context, id string, obj *ipam.Vlan, opts *core.Options) (*ipam.Vlan, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*ipam.Vlan, *http.Response, string, error)
}

type vlanService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewVlanService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) VlanService {
	return &vlanService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new Vlan and returns the created object
func (s *vlanService) Create(ctx context.Context, obj *ipam.Vlan, opts *core.Options) (*ipam.Vlan, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *vlanService) createNIOS(ctx context.Context, obj *ipam.Vlan, opts *core.Options) (*ipam.Vlan, *http.Response, error) {
	payload, err := common.MapTo[niosipam.Vlan](obj, mapper.VlanNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.IPAMAPI.VlanAPI.
		Create(ctx).
		Vlan(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateVlanResponseAsObject.GetResult()

	return mapNIOSVlanToResponse(&result), httpResp, nil
}

// Read retrieves a Vlan by ID
func (s *vlanService) Read(ctx context.Context, id string, opts *core.Options) (*ipam.Vlan, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *vlanService) readNIOS(ctx context.Context, id string, opts *core.Options) (*ipam.Vlan, *http.Response, error) {
	req := s.niosClient.IPAMAPI.VlanAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetVlanResponseObjectAsResult.GetResult()

	return mapNIOSVlanToResponse(&result), httpResp, nil
}

// Update modifies an existing Vlan and returns the updated object
func (s *vlanService) Update(ctx context.Context, id string, obj *ipam.Vlan, opts *core.Options) (*ipam.Vlan, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *vlanService) updateNIOS(ctx context.Context, id string, obj *ipam.Vlan, opts *core.Options) (*ipam.Vlan, *http.Response, error) {
	payload, err := common.MapTo[niosipam.Vlan](obj, mapper.VlanNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.IPAMAPI.VlanAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Vlan(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateVlanResponseAsObject.GetResult()

	return mapNIOSVlanToResponse(&result), httpResp, nil
}

// Delete removes a Vlan by ID
func (s *vlanService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *vlanService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.IPAMAPI.VlanAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves Vlan objects based on filter options
func (s *vlanService) List(ctx context.Context, opts *core.ListOptions) ([]*ipam.Vlan, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *vlanService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*ipam.Vlan, *http.Response, string, error) {
	req := s.niosClient.IPAMAPI.VlanAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.VlanFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListVlanResponseObject.GetResult()
	items := make([]*ipam.Vlan, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSVlanToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListVlanResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSVlanToResponse(r *niosipam.Vlan) *ipam.Vlan {
	resp := &ipam.Vlan{
		Id: r.Ref,
	}
	resp.NIOS = &ipam.NIOSVlanExt{
		Comment:     r.Comment,
		Contact:     r.Contact,
		Department:  r.Department,
		Description: r.Description,
		Id:          r.Id,
		Name:        r.Name,
		Parent:      r.Parent,
		Reserved:    r.Reserved,
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
