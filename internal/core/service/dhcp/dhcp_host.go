package dhcp

import (
	"context"
	"fmt"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dhcp"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

type DhcpHostService interface {
	Create(ctx context.Context, obj *dhcp.DhcpHost, opts *core.Options) (*dhcp.DhcpHost, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dhcp.DhcpHost, *http.Response, error)
	Update(ctx context.Context, id string, obj *dhcp.DhcpHost, opts *core.Options) (*dhcp.DhcpHost, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.DhcpHost, *http.Response, string, error)
}

type dhcpHostService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewDhcpHostService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) DhcpHostService {
	return &dhcpHostService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new DhcpHost and returns the created object
func (s *dhcpHostService) Create(ctx context.Context, obj *dhcp.DhcpHost, opts *core.Options) (*dhcp.DhcpHost, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dhcpHostService) createUDDI(ctx context.Context, obj *dhcp.DhcpHost, opts *core.Options) (*dhcp.DhcpHost, *http.Response, error) {
	// DhcpHost is a system-managed object that pre-exists; Create applies config via Update.
	if obj.Id == nil {
		return nil, nil, fmt.Errorf("id is required: DHCP hosts are pre-existing system objects")
	}
	return s.updateUDDI(ctx, *obj.Id, obj, opts)
}

// Read retrieves a DhcpHost by ID
func (s *dhcpHostService) Read(ctx context.Context, id string, opts *core.Options) (*dhcp.DhcpHost, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dhcpHostService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dhcp.DhcpHost, *http.Response, error) {
	req := s.uddiClient.IPAddressManagementAPI.DhcpHostAPI.
		Read(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIDhcpHostToResponse(&result), httpResp, nil
}

// Update modifies an existing DhcpHost and returns the updated object
func (s *dhcpHostService) Update(ctx context.Context, id string, obj *dhcp.DhcpHost, opts *core.Options) (*dhcp.DhcpHost, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dhcpHostService) updateUDDI(ctx context.Context, id string, obj *dhcp.DhcpHost, opts *core.Options) (*dhcp.DhcpHost, *http.Response, error) {
	// Id is the URL path parameter; strip it from the body to avoid API rejection ("Id is read-only").
	stripped := *obj
	stripped.Id = nil
	payload, err := common.MapTo[uddiipam.Host](&stripped, mapper.DhcpHostUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.IPAddressManagementAPI.DhcpHostAPI.
		Update(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIDhcpHostToResponse(&result), httpResp, nil
}

// Delete removes a DhcpHost by ID
func (s *dhcpHostService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dhcpHostService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	// DhcpHost has no DELETE endpoint. Destroy disassociates by clearing server.
	payload := uddiipam.Host{Server: nil}
	_, httpResp, err := s.uddiClient.IPAddressManagementAPI.DhcpHostAPI.
		Update(ctx, id).
		Body(payload).
		Execute()
	return httpResp, err
}

// List retrieves DhcpHost objects based on filter options
func (s *dhcpHostService) List(ctx context.Context, opts *core.ListOptions) ([]*dhcp.DhcpHost, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *dhcpHostService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dhcp.DhcpHost, *http.Response, string, error) {
	req := s.uddiClient.IPAddressManagementAPI.DhcpHostAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, k+"=='"+v+"'")
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.DhcpHostFilterFieldMap[core.BackendUDDI])
		for k, v := range translatedFilters {
			filters = append(filters, k+"=='"+v+"'")
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
	items := make([]*dhcp.DhcpHost, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIDhcpHostToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDIDhcpHostToResponse(r *uddiipam.Host) *dhcp.DhcpHost {
	resp := &dhcp.DhcpHost{Id: r.Id}
	resp.UDDI = &dhcp.UDDIDhcpHostExt{
		Address:          r.Address,
		AnycastAddresses: r.AnycastAddresses,
		Comment:          r.Comment,
		CurrentVersion:   r.CurrentVersion,
		IpSpace:          r.IpSpace,
		Name:             r.Name,
		Ophid:            r.Ophid,
		ProviderId:       r.ProviderId,
		Server:           r.Server,
		Tags:             r.Tags,
		Type:             r.Type,
	}
	return resp
}
