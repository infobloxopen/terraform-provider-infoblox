package fw

import (
	"context"
	"fmt"
	"maps"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/fw"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/fw"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	uddifw "github.com/infobloxopen/universal-ddi-go-client/fw"
)

type SecurityPolicyService interface {
	Create(ctx context.Context, obj *fw.SecurityPolicy, opts *core.Options) (*fw.SecurityPolicy, *http.Response, error)
	Read(ctx context.Context, id int32, opts *core.Options) (*fw.SecurityPolicy, *http.Response, error)
	Update(ctx context.Context, id int32, obj *fw.SecurityPolicy, opts *core.Options) (*fw.SecurityPolicy, *http.Response, error)
	Delete(ctx context.Context, id int32) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*fw.SecurityPolicy, *http.Response, string, error)
}

type securityPolicyService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewSecurityPolicyService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) SecurityPolicyService {
	return &securityPolicyService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new SecurityPolicy and returns the created object
func (s *securityPolicyService) Create(ctx context.Context, obj *fw.SecurityPolicy, opts *core.Options) (*fw.SecurityPolicy, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *securityPolicyService) createUDDI(ctx context.Context, obj *fw.SecurityPolicy, opts *core.Options) (*fw.SecurityPolicy, *http.Response, error) {
	payload, err := common.MapTo[uddifw.SecurityPolicy](obj, mapper.SecurityPolicyUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.FWAPI.SecurityPoliciesAPI.
		CreateSecurityPolicy(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResults()

	return mapUDDISecurityPolicyToResponse(&result), httpResp, nil
}

// Read retrieves a SecurityPolicy by ID
func (s *securityPolicyService) Read(ctx context.Context, id int32, opts *core.Options) (*fw.SecurityPolicy, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *securityPolicyService) readUDDI(ctx context.Context, id int32, opts *core.Options) (*fw.SecurityPolicy, *http.Response, error) {
	req := s.uddiClient.FWAPI.SecurityPoliciesAPI.
		ReadSecurityPolicy(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResults()

	return mapUDDISecurityPolicyToResponse(&result), httpResp, nil
}

// Update modifies an existing SecurityPolicy and returns the updated object
func (s *securityPolicyService) Update(ctx context.Context, id int32, obj *fw.SecurityPolicy, opts *core.Options) (*fw.SecurityPolicy, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *securityPolicyService) updateUDDI(ctx context.Context, id int32, obj *fw.SecurityPolicy, opts *core.Options) (*fw.SecurityPolicy, *http.Response, error) {
	payload, err := common.MapTo[uddifw.SecurityPolicy](obj, mapper.SecurityPolicyUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.FWAPI.SecurityPoliciesAPI.
		UpdateSecurityPolicy(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResults()

	return mapUDDISecurityPolicyToResponse(&result), httpResp, nil
}

// Delete removes a SecurityPolicy by ID
func (s *securityPolicyService) Delete(ctx context.Context, id int32) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *securityPolicyService) deleteUDDI(ctx context.Context, id int32) (*http.Response, error) {
	httpResp, err := s.uddiClient.FWAPI.SecurityPoliciesAPI.
		DeleteSingleSecurityPolicy(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves SecurityPolicy objects based on filter options
func (s *securityPolicyService) List(ctx context.Context, opts *core.ListOptions) ([]*fw.SecurityPolicy, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *securityPolicyService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*fw.SecurityPolicy, *http.Response, string, error) {
	req := s.uddiClient.FWAPI.SecurityPoliciesAPI.ListSecurityPolicies(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.SecurityPolicyFilterFieldMap[core.BackendUDDI])
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
	items := make([]*fw.SecurityPolicy, 0, len(results))
	for i := range results {
		items = append(items, mapUDDISecurityPolicyToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDISecurityPolicyToResponse(r *uddifw.SecurityPolicy) *fw.SecurityPolicy {
	resp := &fw.SecurityPolicy{
		Id: r.Id,
	}
	resp.UDDI = &fw.UDDISecurityPolicyExt{
		AccessCodes:         r.AccessCodes,
		DefaultAction:       r.DefaultAction,
		DefaultRedirectName: r.DefaultRedirectName,
		Description:         r.Description,
		DfpServices:         r.DfpServices,
		Dfps:                r.Dfps,
		Ecs:                 r.Ecs,
		Name:                r.Name,
		NetAddressDfps:      r.NetAddressDfps,
		NetworkLists:        r.NetworkLists,
		OnpremResolve:       r.OnpremResolve,
		Precedence:          r.Precedence,
		RoamingDeviceGroups: r.RoamingDeviceGroups,
		Rules:               r.Rules,
		SafeSearch:          r.SafeSearch,
		UserGroups:          r.UserGroups,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
