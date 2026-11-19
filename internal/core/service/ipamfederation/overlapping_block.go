package ipamfederation

import (
	"context"
	"fmt"
	"maps"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/ipamfederation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/ipamfederation"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	uddiipamfederation "github.com/infobloxopen/universal-ddi-go-client/ipamfederation"
)

type OverlappingBlockService interface {
	Create(ctx context.Context, obj *ipamfederation.OverlappingBlock, opts *core.Options) (*ipamfederation.OverlappingBlock, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*ipamfederation.OverlappingBlock, *http.Response, error)
	Update(ctx context.Context, id string, obj *ipamfederation.OverlappingBlock, opts *core.Options) (*ipamfederation.OverlappingBlock, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*ipamfederation.OverlappingBlock, *http.Response, string, error)
}

type overlappingBlockService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewOverlappingBlockService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) OverlappingBlockService {
	return &overlappingBlockService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new OverlappingBlock and returns the created object
func (s *overlappingBlockService) Create(ctx context.Context, obj *ipamfederation.OverlappingBlock, opts *core.Options) (*ipamfederation.OverlappingBlock, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *overlappingBlockService) createUDDI(ctx context.Context, obj *ipamfederation.OverlappingBlock, opts *core.Options) (*ipamfederation.OverlappingBlock, *http.Response, error) {
	payload, err := common.MapTo[uddiipamfederation.OverlappingBlock](obj, mapper.OverlappingBlockUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAMFederationAPI.OverlappingBlockAPI.
		Create(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIOverlappingBlockToResponse(&result), httpResp, nil
}

// Read retrieves a OverlappingBlock by ID
func (s *overlappingBlockService) Read(ctx context.Context, id string, opts *core.Options) (*ipamfederation.OverlappingBlock, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *overlappingBlockService) readUDDI(ctx context.Context, id string, opts *core.Options) (*ipamfederation.OverlappingBlock, *http.Response, error) {
	req := s.uddiClient.IPAMFederationAPI.OverlappingBlockAPI.
		Read(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIOverlappingBlockToResponse(&result), httpResp, nil
}

// Update modifies an existing OverlappingBlock and returns the updated object
func (s *overlappingBlockService) Update(ctx context.Context, id string, obj *ipamfederation.OverlappingBlock, opts *core.Options) (*ipamfederation.OverlappingBlock, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *overlappingBlockService) updateUDDI(ctx context.Context, id string, obj *ipamfederation.OverlappingBlock, opts *core.Options) (*ipamfederation.OverlappingBlock, *http.Response, error) {
	payload, err := common.MapTo[uddiipamfederation.OverlappingBlock](obj, mapper.OverlappingBlockUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAMFederationAPI.OverlappingBlockAPI.
		Update(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIOverlappingBlockToResponse(&result), httpResp, nil
}

// Delete removes a OverlappingBlock by ID
func (s *overlappingBlockService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *overlappingBlockService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.IPAMFederationAPI.OverlappingBlockAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves OverlappingBlock objects based on filter options
func (s *overlappingBlockService) List(ctx context.Context, opts *core.ListOptions) ([]*ipamfederation.OverlappingBlock, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *overlappingBlockService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*ipamfederation.OverlappingBlock, *http.Response, string, error) {
	req := s.uddiClient.IPAMFederationAPI.OverlappingBlockAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.OverlappingBlockFilterFieldMap[core.BackendUDDI])
		for k, v := range translatedFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		if len(filters) > 0 {
			req = req.Filter(core.JoinFilters(filters))
		}

		if len(opts.TagFilter) > 0 {
			var tfilters []string
			for k, v := range opts.TagFilter {
				tfilters = append(tfilters, "'"+k+"'=='"+v+"'")
			}
			req = req.Tfilter(core.JoinFilters(tfilters))
		}

		if opts.Offset > 0 {
			req = req.Offset(opts.Offset)
		}

		if opts.Limit > 0 {
			req = req.Limit(opts.Limit)
		}
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, "", err
	}

	results := resp.GetResults()
	items := make([]*ipamfederation.OverlappingBlock, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIOverlappingBlockToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDIOverlappingBlockToResponse(r *uddiipamfederation.OverlappingBlock) *ipamfederation.OverlappingBlock {
	resp := &ipamfederation.OverlappingBlock{
		Id: r.Id,
	}
	resp.UDDI = &ipamfederation.UDDIOverlappingBlockExt{
		Address:          r.Address,
		Cidr:             r.Cidr,
		Comment:          r.Comment,
		FederatedPoolId:  r.FederatedPoolId,
		FederatedRealm:   r.FederatedRealm,
		Name:             r.Name,
		NetworkCompliant: r.NetworkCompliant,
		Parent:           r.Parent,
		Protocol:         r.Protocol,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
