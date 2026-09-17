package fw

import (
	"context"
	"fmt"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/fw"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/fw"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	uddifw "github.com/infobloxopen/universal-ddi-go-client/fw"
)

type NetworkListService interface {
	Create(ctx context.Context, obj *fw.NetworkList, opts *core.Options) (*fw.NetworkList, *http.Response, error)
	Read(ctx context.Context, id int32, opts *core.Options) (*fw.NetworkList, *http.Response, error)
	Update(ctx context.Context, id int32, obj *fw.NetworkList, opts *core.Options) (*fw.NetworkList, *http.Response, error)
	Delete(ctx context.Context, id int32) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*fw.NetworkList, *http.Response, string, error)
}

type networkListService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewNetworkListService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) NetworkListService {
	return &networkListService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new NetworkList and returns the created object
func (s *networkListService) Create(ctx context.Context, obj *fw.NetworkList, opts *core.Options) (*fw.NetworkList, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *networkListService) createUDDI(ctx context.Context, obj *fw.NetworkList, opts *core.Options) (*fw.NetworkList, *http.Response, error) {
	payload, err := common.MapTo[uddifw.NetworkList](obj, mapper.NetworkListUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.FWAPI.NetworkListsAPI.
		CreateNetworkList(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResults()

	return mapUDDINetworkListToResponse(&result), httpResp, nil
}

// Read retrieves a NetworkList by ID
func (s *networkListService) Read(ctx context.Context, id int32, opts *core.Options) (*fw.NetworkList, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *networkListService) readUDDI(ctx context.Context, id int32, opts *core.Options) (*fw.NetworkList, *http.Response, error) {
	req := s.uddiClient.FWAPI.NetworkListsAPI.
		ReadNetworkList(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResults()

	return mapUDDINetworkListToResponse(&result), httpResp, nil
}

// Update modifies an existing NetworkList and returns the updated object
func (s *networkListService) Update(ctx context.Context, id int32, obj *fw.NetworkList, opts *core.Options) (*fw.NetworkList, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *networkListService) updateUDDI(ctx context.Context, id int32, obj *fw.NetworkList, opts *core.Options) (*fw.NetworkList, *http.Response, error) {
	payload, err := common.MapTo[uddifw.NetworkList](obj, mapper.NetworkListUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.FWAPI.NetworkListsAPI.
		UpdateNetworkList(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResults()

	return mapUDDINetworkListToResponse(&result), httpResp, nil
}

// Delete removes a NetworkList by ID
func (s *networkListService) Delete(ctx context.Context, id int32) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *networkListService) deleteUDDI(ctx context.Context, id int32) (*http.Response, error) {
	httpResp, err := s.uddiClient.FWAPI.NetworkListsAPI.
		DeleteSingleNetworkLists(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves NetworkList objects based on filter options
func (s *networkListService) List(ctx context.Context, opts *core.ListOptions) ([]*fw.NetworkList, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *networkListService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*fw.NetworkList, *http.Response, string, error) {
	req := s.uddiClient.FWAPI.NetworkListsAPI.ListNetworkLists(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.NetworkListFilterFieldMap[core.BackendUDDI])
		for k, v := range translatedFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		if len(filters) > 0 {
			req = req.Filter(core.JoinFilters(filters))
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
	items := make([]*fw.NetworkList, 0, len(results))
	for i := range results {
		items = append(items, mapUDDINetworkListToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDINetworkListToResponse(r *uddifw.NetworkList) *fw.NetworkList {
	resp := &fw.NetworkList{
		Id: r.Id,
	}
	resp.UDDI = &fw.UDDINetworkListExt{
		AddrBlock:   r.AddrBlock,
		Description: r.Description,
		Items:       r.Items,
		Name:        r.Name,
	}
	return resp
}
