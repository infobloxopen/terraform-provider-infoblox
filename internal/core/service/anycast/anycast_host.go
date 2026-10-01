package anycast

import (
	"context"
	"fmt"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/anycast"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/anycast"
	uddianycast "github.com/infobloxopen/universal-ddi-go-client/anycast"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
)

type AnycastHostService interface {
	Create(ctx context.Context, obj *anycast.AnycastHost, opts *core.Options) (*anycast.AnycastHost, *http.Response, error)
	Read(ctx context.Context, id int64, opts *core.Options) (*anycast.AnycastHost, *http.Response, error)
	Update(ctx context.Context, id int64, obj *anycast.AnycastHost, opts *core.Options) (*anycast.AnycastHost, *http.Response, error)
	Delete(ctx context.Context, id int64) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*anycast.AnycastHost, *http.Response, string, error)
}

type anycastHostService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewAnycastHostService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) AnycastHostService {
	return &anycastHostService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new AnycastHost and returns the created object
func (s *anycastHostService) Create(ctx context.Context, obj *anycast.AnycastHost, opts *core.Options) (*anycast.AnycastHost, *http.Response, error) {
	switch s.backend {
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

// Read retrieves a AnycastHost by ID
func (s *anycastHostService) Read(ctx context.Context, id int64, opts *core.Options) (*anycast.AnycastHost, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *anycastHostService) readUDDI(ctx context.Context, id int64, opts *core.Options) (*anycast.AnycastHost, *http.Response, error) {
	req := s.uddiClient.AnycastAPI.OnPremAnycastManagerAPI.
		GetOnpremHost(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResults()

	return mapUDDIAnycastHostToResponse(&result), httpResp, nil
}

// Update modifies an existing AnycastHost and returns the updated object
func (s *anycastHostService) Update(ctx context.Context, id int64, obj *anycast.AnycastHost, opts *core.Options) (*anycast.AnycastHost, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *anycastHostService) updateUDDI(ctx context.Context, id int64, obj *anycast.AnycastHost, opts *core.Options) (*anycast.AnycastHost, *http.Response, error) {
	payload, err := common.MapTo[uddianycast.OnpremHost](obj, mapper.AnycastHostUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.AnycastAPI.OnPremAnycastManagerAPI.
		UpdateOnpremHost(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResults()

	return mapUDDIAnycastHostToResponse(&result), httpResp, nil
}

// Delete removes a AnycastHost by ID
func (s *anycastHostService) Delete(ctx context.Context, id int64) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *anycastHostService) deleteUDDI(ctx context.Context, id int64) (*http.Response, error) {
	// This endpoint declares a delete response body, so Execute returns it too.
	_, httpResp, err := s.uddiClient.AnycastAPI.OnPremAnycastManagerAPI.
		DeleteOnpremHost(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves AnycastHost objects based on filter options
func (s *anycastHostService) List(ctx context.Context, opts *core.ListOptions) ([]*anycast.AnycastHost, *http.Response, string, error) {
	switch s.backend {
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func mapUDDIAnycastHostToResponse(r *uddianycast.OnpremHost) *anycast.AnycastHost {
	resp := &anycast.AnycastHost{
		Id: r.Id,
	}
	resp.UDDI = &anycast.UDDIAnycastHostExt{
		AnycastConfigRefs: r.AnycastConfigRefs,
		ConfigBgp:         r.ConfigBgp,
		ConfigOspf:        r.ConfigOspf,
		ConfigOspfv3:      r.ConfigOspfv3,
		CreatedAt:         r.CreatedAt,
		IpAddress:         r.IpAddress,
		Ipv6Address:       r.Ipv6Address,
		Name:              r.Name,
		UpdatedAt:         r.UpdatedAt,
	}
	return resp
}
