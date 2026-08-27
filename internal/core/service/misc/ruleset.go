package misc

import (
	"context"
	"fmt"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niosmisc "github.com/infobloxopen/infoblox-nios-go-client/misc"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/misc"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/misc"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
)

type RulesetService interface {
	Create(ctx context.Context, obj *misc.Ruleset, opts *core.Options) (*misc.Ruleset, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*misc.Ruleset, *http.Response, error)
	Update(ctx context.Context, id string, obj *misc.Ruleset, opts *core.Options) (*misc.Ruleset, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*misc.Ruleset, *http.Response, string, error)
}

type rulesetService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewRulesetService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) RulesetService {
	return &rulesetService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new Ruleset and returns the created object
func (s *rulesetService) Create(ctx context.Context, obj *misc.Ruleset, opts *core.Options) (*misc.Ruleset, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *rulesetService) createNIOS(ctx context.Context, obj *misc.Ruleset, opts *core.Options) (*misc.Ruleset, *http.Response, error) {
	payload, err := common.MapTo[niosmisc.Ruleset](obj, mapper.RulesetNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.MiscAPI.RulesetAPI.
		Create(ctx).
		Ruleset(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateRulesetResponseAsObject.GetResult()

	return mapNIOSRulesetToResponse(&result), httpResp, nil
}

// Read retrieves a Ruleset by ID
func (s *rulesetService) Read(ctx context.Context, id string, opts *core.Options) (*misc.Ruleset, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *rulesetService) readNIOS(ctx context.Context, id string, opts *core.Options) (*misc.Ruleset, *http.Response, error) {
	req := s.niosClient.MiscAPI.RulesetAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetRulesetResponseObjectAsResult.GetResult()

	return mapNIOSRulesetToResponse(&result), httpResp, nil
}

// Update modifies an existing Ruleset and returns the updated object
func (s *rulesetService) Update(ctx context.Context, id string, obj *misc.Ruleset, opts *core.Options) (*misc.Ruleset, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *rulesetService) updateNIOS(ctx context.Context, id string, obj *misc.Ruleset, opts *core.Options) (*misc.Ruleset, *http.Response, error) {
	payload, err := common.MapTo[niosmisc.Ruleset](obj, mapper.RulesetNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.MiscAPI.RulesetAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Ruleset(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateRulesetResponseAsObject.GetResult()

	return mapNIOSRulesetToResponse(&result), httpResp, nil
}

// Delete removes a Ruleset by ID
func (s *rulesetService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *rulesetService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.MiscAPI.RulesetAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves Ruleset objects based on filter options
func (s *rulesetService) List(ctx context.Context, opts *core.ListOptions) ([]*misc.Ruleset, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *rulesetService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*misc.Ruleset, *http.Response, string, error) {
	req := s.niosClient.MiscAPI.RulesetAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.RulesetFilterFieldMap[core.BackendNIOS])
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
		if opts.Paging == 1 {
			maxResults := opts.MaxResults
			if maxResults <= 0 {
				maxResults = core.DefaultListLimit
			}
			req = req.MaxResults(maxResults)
		} else if opts.MaxResults > 0 {
			req = req.MaxResults(opts.MaxResults)
		}
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, "", err
	}

	results := resp.ListRulesetResponseObject.GetResult()
	items := make([]*misc.Ruleset, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSRulesetToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListRulesetResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSRulesetToResponse(r *niosmisc.Ruleset) *misc.Ruleset {
	resp := &misc.Ruleset{
		Id: r.Ref,
	}
	resp.NIOS = &misc.NIOSRulesetExt{
		Comment:       r.Comment,
		Disabled:      r.Disabled,
		Name:          r.Name,
		NxdomainRules: r.NxdomainRules,
		Type:          r.Type,
	}
	return resp
}
