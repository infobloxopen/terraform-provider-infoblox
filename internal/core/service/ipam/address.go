package ipam

import (
	"context"
	"fmt"
	"maps"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/ipam"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

type AddressService interface {
	Create(ctx context.Context, obj *ipam.Address, opts *core.Options) (*ipam.Address, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*ipam.Address, *http.Response, error)
	Update(ctx context.Context, id string, obj *ipam.Address, opts *core.Options) (*ipam.Address, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*ipam.Address, *http.Response, string, error)
}

type addressService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewAddressService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) AddressService {
	return &addressService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new Address and returns the created object
func (s *addressService) Create(ctx context.Context, obj *ipam.Address, opts *core.Options) (*ipam.Address, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *addressService) createUDDI(ctx context.Context, obj *ipam.Address, opts *core.Options) (*ipam.Address, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.Address](obj, mapper.AddressUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAddressManagementAPI.AddressAPI.
		Create(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIAddressToResponse(&result), httpResp, nil
}

// Read retrieves a Address by ID
func (s *addressService) Read(ctx context.Context, id string, opts *core.Options) (*ipam.Address, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *addressService) readUDDI(ctx context.Context, id string, opts *core.Options) (*ipam.Address, *http.Response, error) {
	req := s.uddiClient.IPAddressManagementAPI.AddressAPI.
		Read(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIAddressToResponse(&result), httpResp, nil
}

// Update modifies an existing Address and returns the updated object
func (s *addressService) Update(ctx context.Context, id string, obj *ipam.Address, opts *core.Options) (*ipam.Address, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *addressService) updateUDDI(ctx context.Context, id string, obj *ipam.Address, opts *core.Options) (*ipam.Address, *http.Response, error) {
	payload, err := common.MapTo[uddiipam.Address](obj, mapper.AddressUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAddressManagementAPI.AddressAPI.
		Update(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIAddressToResponse(&result), httpResp, nil
}

// Delete removes a Address by ID
func (s *addressService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *addressService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.IPAddressManagementAPI.AddressAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves Address objects based on filter options
func (s *addressService) List(ctx context.Context, opts *core.ListOptions) ([]*ipam.Address, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *addressService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*ipam.Address, *http.Response, string, error) {
	req := s.uddiClient.IPAddressManagementAPI.AddressAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.AddressFilterFieldMap[core.BackendUDDI])
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
	items := make([]*ipam.Address, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIAddressToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDIAddressToResponse(r *uddiipam.Address) *ipam.Address {
	resp := &ipam.Address{
		Id: r.Id,
	}
	resp.UDDI = &ipam.UDDIAddressExt{
		Address:      r.Address,
		Comment:      r.Comment,
		ExternalKeys: r.ExternalKeys,
		Host:         r.Host,
		Hwaddr:       r.Hwaddr,
		Interface:    r.Interface,
		Names:        r.Names,
		Parent:       r.Parent,
		Range:        r.Range,
		Space:        r.Space,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
