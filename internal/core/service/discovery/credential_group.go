package discovery

import (
	"context"
	"fmt"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niosdiscovery "github.com/infobloxopen/infoblox-nios-go-client/discovery"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/discovery"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/discovery"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
)

type CredentialGroupService interface {
	Create(ctx context.Context, obj *discovery.CredentialGroup, opts *core.Options) (*discovery.CredentialGroup, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*discovery.CredentialGroup, *http.Response, error)
	Update(ctx context.Context, id string, obj *discovery.CredentialGroup, opts *core.Options) (*discovery.CredentialGroup, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*discovery.CredentialGroup, *http.Response, string, error)
}

type credentialGroupService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewCredentialGroupService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) CredentialGroupService {
	return &credentialGroupService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new CredentialGroup and returns the created object
func (s *credentialGroupService) Create(ctx context.Context, obj *discovery.CredentialGroup, opts *core.Options) (*discovery.CredentialGroup, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *credentialGroupService) createNIOS(ctx context.Context, obj *discovery.CredentialGroup, opts *core.Options) (*discovery.CredentialGroup, *http.Response, error) {
	payload, err := common.MapTo[niosdiscovery.DiscoveryCredentialgroup](obj, mapper.CredentialGroupNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.DiscoveryAPI.DiscoveryCredentialgroupAPI.
		Create(ctx).
		DiscoveryCredentialgroup(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateDiscoveryCredentialgroupResponseAsObject.GetResult()

	return mapNIOSCredentialGroupToResponse(&result), httpResp, nil
}

// Read retrieves a CredentialGroup by ID
func (s *credentialGroupService) Read(ctx context.Context, id string, opts *core.Options) (*discovery.CredentialGroup, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *credentialGroupService) readNIOS(ctx context.Context, id string, opts *core.Options) (*discovery.CredentialGroup, *http.Response, error) {
	req := s.niosClient.DiscoveryAPI.DiscoveryCredentialgroupAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetDiscoveryCredentialgroupResponseObjectAsResult.GetResult()

	return mapNIOSCredentialGroupToResponse(&result), httpResp, nil
}

// Update modifies an existing CredentialGroup and returns the updated object
func (s *credentialGroupService) Update(ctx context.Context, id string, obj *discovery.CredentialGroup, opts *core.Options) (*discovery.CredentialGroup, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *credentialGroupService) updateNIOS(ctx context.Context, id string, obj *discovery.CredentialGroup, opts *core.Options) (*discovery.CredentialGroup, *http.Response, error) {
	payload, err := common.MapTo[niosdiscovery.DiscoveryCredentialgroup](obj, mapper.CredentialGroupNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.DiscoveryAPI.DiscoveryCredentialgroupAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		DiscoveryCredentialgroup(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateDiscoveryCredentialgroupResponseAsObject.GetResult()

	return mapNIOSCredentialGroupToResponse(&result), httpResp, nil
}

// Delete removes a CredentialGroup by ID
func (s *credentialGroupService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *credentialGroupService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DiscoveryAPI.DiscoveryCredentialgroupAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves CredentialGroup objects based on filter options
func (s *credentialGroupService) List(ctx context.Context, opts *core.ListOptions) ([]*discovery.CredentialGroup, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *credentialGroupService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*discovery.CredentialGroup, *http.Response, string, error) {
	req := s.niosClient.DiscoveryAPI.DiscoveryCredentialgroupAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.CredentialGroupFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListDiscoveryCredentialgroupResponseObject.GetResult()
	items := make([]*discovery.CredentialGroup, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSCredentialGroupToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListDiscoveryCredentialgroupResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSCredentialGroupToResponse(r *niosdiscovery.DiscoveryCredentialgroup) *discovery.CredentialGroup {
	resp := &discovery.CredentialGroup{
		Id: r.Ref,
	}
	resp.NIOS = &discovery.NIOSCredentialGroupExt{
		Name: r.Name,
	}
	return resp
}
