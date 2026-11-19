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

type ReservedBlockService interface {
	Create(ctx context.Context, obj *ipamfederation.ReservedBlock, opts *core.Options) (*ipamfederation.ReservedBlock, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*ipamfederation.ReservedBlock, *http.Response, error)
	Update(ctx context.Context, id string, obj *ipamfederation.ReservedBlock, opts *core.Options) (*ipamfederation.ReservedBlock, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*ipamfederation.ReservedBlock, *http.Response, string, error)
}

type reservedBlockService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewReservedBlockService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) ReservedBlockService {
	return &reservedBlockService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new ReservedBlock and returns the created object
func (s *reservedBlockService) Create(ctx context.Context, obj *ipamfederation.ReservedBlock, opts *core.Options) (*ipamfederation.ReservedBlock, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *reservedBlockService) createUDDI(ctx context.Context, obj *ipamfederation.ReservedBlock, opts *core.Options) (*ipamfederation.ReservedBlock, *http.Response, error) {
	payload, err := common.MapTo[uddiipamfederation.ReservedBlock](obj, mapper.ReservedBlockUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAMFederationAPI.ReservedBlockAPI.
		Create(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIReservedBlockToResponse(&result), httpResp, nil
}

// Read retrieves a ReservedBlock by ID
func (s *reservedBlockService) Read(ctx context.Context, id string, opts *core.Options) (*ipamfederation.ReservedBlock, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *reservedBlockService) readUDDI(ctx context.Context, id string, opts *core.Options) (*ipamfederation.ReservedBlock, *http.Response, error) {
	req := s.uddiClient.IPAMFederationAPI.ReservedBlockAPI.
		Read(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIReservedBlockToResponse(&result), httpResp, nil
}

// Update modifies an existing ReservedBlock and returns the updated object
func (s *reservedBlockService) Update(ctx context.Context, id string, obj *ipamfederation.ReservedBlock, opts *core.Options) (*ipamfederation.ReservedBlock, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *reservedBlockService) updateUDDI(ctx context.Context, id string, obj *ipamfederation.ReservedBlock, opts *core.Options) (*ipamfederation.ReservedBlock, *http.Response, error) {
	payload, err := common.MapTo[uddiipamfederation.ReservedBlock](obj, mapper.ReservedBlockUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAMFederationAPI.ReservedBlockAPI.
		Update(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIReservedBlockToResponse(&result), httpResp, nil
}

// Delete removes a ReservedBlock by ID
func (s *reservedBlockService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *reservedBlockService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.IPAMFederationAPI.ReservedBlockAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves ReservedBlock objects based on filter options
func (s *reservedBlockService) List(ctx context.Context, opts *core.ListOptions) ([]*ipamfederation.ReservedBlock, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *reservedBlockService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*ipamfederation.ReservedBlock, *http.Response, string, error) {
	req := s.uddiClient.IPAMFederationAPI.ReservedBlockAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.ReservedBlockFilterFieldMap[core.BackendUDDI])
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
	items := make([]*ipamfederation.ReservedBlock, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIReservedBlockToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDIReservedBlockToResponse(r *uddiipamfederation.ReservedBlock) *ipamfederation.ReservedBlock {
	resp := &ipamfederation.ReservedBlock{
		Id: r.Id,
	}
	resp.UDDI = &ipamfederation.UDDIReservedBlockExt{
		Address:          r.Address,
		Cidr:             r.Cidr,
		Comment:          r.Comment,
		CreatedAt:        r.CreatedAt,
		FederatedPoolId:  r.FederatedPoolId,
		FederatedRealm:   r.FederatedRealm,
		Metadata:         r.Metadata,
		Name:             r.Name,
		NetworkCompliant: r.NetworkCompliant,
		Parent:           r.Parent,
		Protocol:         r.Protocol,
		Region:           r.Region,
		UpdatedAt:        r.UpdatedAt,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
