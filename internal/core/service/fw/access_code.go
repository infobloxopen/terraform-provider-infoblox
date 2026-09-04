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

type AccessCodeService interface {
	Create(ctx context.Context, obj *fw.AccessCode, opts *core.Options) (*fw.AccessCode, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*fw.AccessCode, *http.Response, error)
	Update(ctx context.Context, id string, obj *fw.AccessCode, opts *core.Options) (*fw.AccessCode, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*fw.AccessCode, *http.Response, string, error)
}

type accessCodeService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewAccessCodeService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) AccessCodeService {
	return &accessCodeService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new AccessCode and returns the created object
func (s *accessCodeService) Create(ctx context.Context, obj *fw.AccessCode, opts *core.Options) (*fw.AccessCode, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *accessCodeService) createUDDI(ctx context.Context, obj *fw.AccessCode, opts *core.Options) (*fw.AccessCode, *http.Response, error) {
	payload, err := common.MapTo[uddifw.AccessCode](obj, mapper.AccessCodeUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.FWAPI.AccessCodesAPI.
		CreateAccessCode(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResults()

	return mapUDDIAccessCodeToResponse(&result), httpResp, nil
}

// Read retrieves a AccessCode by ID
func (s *accessCodeService) Read(ctx context.Context, id string, opts *core.Options) (*fw.AccessCode, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *accessCodeService) readUDDI(ctx context.Context, id string, opts *core.Options) (*fw.AccessCode, *http.Response, error) {
	req := s.uddiClient.FWAPI.AccessCodesAPI.
		ReadAccessCode(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResults()

	return mapUDDIAccessCodeToResponse(&result), httpResp, nil
}

// Update modifies an existing AccessCode and returns the updated object
func (s *accessCodeService) Update(ctx context.Context, id string, obj *fw.AccessCode, opts *core.Options) (*fw.AccessCode, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *accessCodeService) updateUDDI(ctx context.Context, id string, obj *fw.AccessCode, opts *core.Options) (*fw.AccessCode, *http.Response, error) {
	payload, err := common.MapTo[uddifw.AccessCode](obj, mapper.AccessCodeUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.FWAPI.AccessCodesAPI.
		UpdateAccessCode(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResults()

	return mapUDDIAccessCodeToResponse(&result), httpResp, nil
}

// Delete removes a AccessCode by ID
func (s *accessCodeService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *accessCodeService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.FWAPI.AccessCodesAPI.
		DeleteSingleAccessCodes(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves AccessCode objects based on filter options
func (s *accessCodeService) List(ctx context.Context, opts *core.ListOptions) ([]*fw.AccessCode, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *accessCodeService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*fw.AccessCode, *http.Response, string, error) {
	req := s.uddiClient.FWAPI.AccessCodesAPI.ListAccessCodes(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.AccessCodeFilterFieldMap[core.BackendUDDI])
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
	items := make([]*fw.AccessCode, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIAccessCodeToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDIAccessCodeToResponse(r *uddifw.AccessCode) *fw.AccessCode {
	resp := &fw.AccessCode{
		Id: r.AccessKey,
	}
	resp.UDDI = &fw.UDDIAccessCodeExt{
		AccessKey:   r.AccessKey,
		Activation:  r.Activation,
		CreatedTime: r.CreatedTime,
		Description: r.Description,
		Expiration:  r.Expiration,
		Name:        r.Name,
		PolicyIds:   r.PolicyIds,
		Rules:       r.Rules,
		UpdatedTime: r.UpdatedTime,
	}
	return resp
}
