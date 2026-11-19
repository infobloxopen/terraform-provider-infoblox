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

type FederatedBlockService interface {
	Create(ctx context.Context, obj *ipamfederation.FederatedBlock, opts *core.Options) (*ipamfederation.FederatedBlock, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*ipamfederation.FederatedBlock, *http.Response, error)
	Update(ctx context.Context, id string, obj *ipamfederation.FederatedBlock, opts *core.Options) (*ipamfederation.FederatedBlock, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*ipamfederation.FederatedBlock, *http.Response, string, error)
}

type federatedBlockService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewFederatedBlockService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) FederatedBlockService {
	return &federatedBlockService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new FederatedBlock and returns the created object
func (s *federatedBlockService) Create(ctx context.Context, obj *ipamfederation.FederatedBlock, opts *core.Options) (*ipamfederation.FederatedBlock, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *federatedBlockService) createUDDI(ctx context.Context, obj *ipamfederation.FederatedBlock, opts *core.Options) (*ipamfederation.FederatedBlock, *http.Response, error) {
	payload, err := common.MapTo[uddiipamfederation.FederatedBlock](obj, mapper.FederatedBlockUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAMFederationAPI.FederatedBlockAPI.
		Create(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIFederatedBlockToResponse(&result), httpResp, nil
}

// Read retrieves a FederatedBlock by ID
func (s *federatedBlockService) Read(ctx context.Context, id string, opts *core.Options) (*ipamfederation.FederatedBlock, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *federatedBlockService) readUDDI(ctx context.Context, id string, opts *core.Options) (*ipamfederation.FederatedBlock, *http.Response, error) {
	req := s.uddiClient.IPAMFederationAPI.FederatedBlockAPI.
		Read(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIFederatedBlockToResponse(&result), httpResp, nil
}

// Update modifies an existing FederatedBlock and returns the updated object
func (s *federatedBlockService) Update(ctx context.Context, id string, obj *ipamfederation.FederatedBlock, opts *core.Options) (*ipamfederation.FederatedBlock, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *federatedBlockService) updateUDDI(ctx context.Context, id string, obj *ipamfederation.FederatedBlock, opts *core.Options) (*ipamfederation.FederatedBlock, *http.Response, error) {
	payload, err := common.MapTo[uddiipamfederation.FederatedBlock](obj, mapper.FederatedBlockUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAMFederationAPI.FederatedBlockAPI.
		Update(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIFederatedBlockToResponse(&result), httpResp, nil
}

// Delete removes a FederatedBlock by ID
func (s *federatedBlockService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *federatedBlockService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.IPAMFederationAPI.FederatedBlockAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves FederatedBlock objects based on filter options
func (s *federatedBlockService) List(ctx context.Context, opts *core.ListOptions) ([]*ipamfederation.FederatedBlock, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *federatedBlockService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*ipamfederation.FederatedBlock, *http.Response, string, error) {
	req := s.uddiClient.IPAMFederationAPI.FederatedBlockAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.FederatedBlockFilterFieldMap[core.BackendUDDI])
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
	items := make([]*ipamfederation.FederatedBlock, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIFederatedBlockToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDIFederatedBlockToResponse(r *uddiipamfederation.FederatedBlock) *ipamfederation.FederatedBlock {
	resp := &ipamfederation.FederatedBlock{
		Id: r.Id,
	}
	resp.UDDI = &ipamfederation.UDDIFederatedBlockExt{
		Address:           r.Address,
		AllocationV4:      r.AllocationV4,
		Cidr:              r.Cidr,
		Comment:           r.Comment,
		CreatedAt:         r.CreatedAt,
		FederatedPoolId:   r.FederatedPoolId,
		FederatedRealm:    r.FederatedRealm,
		Metadata:          r.Metadata,
		Name:              r.Name,
		NetworkCompliance: r.NetworkCompliance,
		NetworkCompliant:  r.NetworkCompliant,
		Parent:            r.Parent,
		Protocol:          r.Protocol,
		Region:            r.Region,
		State:             r.State,
		UpdatedAt:         r.UpdatedAt,
		Utilization:       r.Utilization,
		UtilizationV6:     r.UtilizationV6,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
