package cloud

import (
	"context"
	"fmt"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	nioscloud "github.com/infobloxopen/infoblox-nios-go-client/cloud"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/cloud"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/cloud"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
)

type AwsuserService interface {
	Create(ctx context.Context, obj *cloud.Awsuser, opts *core.Options) (*cloud.Awsuser, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*cloud.Awsuser, *http.Response, error)
	Update(ctx context.Context, id string, obj *cloud.Awsuser, opts *core.Options) (*cloud.Awsuser, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*cloud.Awsuser, *http.Response, string, error)
}

type awsuserService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewAwsuserService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) AwsuserService {
	return &awsuserService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new Awsuser and returns the created object
func (s *awsuserService) Create(ctx context.Context, obj *cloud.Awsuser, opts *core.Options) (*cloud.Awsuser, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *awsuserService) createNIOS(ctx context.Context, obj *cloud.Awsuser, opts *core.Options) (*cloud.Awsuser, *http.Response, error) {
	payload, err := common.MapTo[nioscloud.Awsuser](obj, mapper.AwsuserNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.CloudAPI.AwsuserAPI.
		Create(ctx).
		Awsuser(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateAwsuserResponseAsObject.GetResult()

	return mapNIOSAwsuserToResponse(&result), httpResp, nil
}

// Read retrieves a Awsuser by ID
func (s *awsuserService) Read(ctx context.Context, id string, opts *core.Options) (*cloud.Awsuser, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *awsuserService) readNIOS(ctx context.Context, id string, opts *core.Options) (*cloud.Awsuser, *http.Response, error) {
	req := s.niosClient.CloudAPI.AwsuserAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetAwsuserResponseObjectAsResult.GetResult()

	return mapNIOSAwsuserToResponse(&result), httpResp, nil
}

// Update modifies an existing Awsuser and returns the updated object
func (s *awsuserService) Update(ctx context.Context, id string, obj *cloud.Awsuser, opts *core.Options) (*cloud.Awsuser, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *awsuserService) updateNIOS(ctx context.Context, id string, obj *cloud.Awsuser, opts *core.Options) (*cloud.Awsuser, *http.Response, error) {
	payload, err := common.MapTo[nioscloud.Awsuser](obj, mapper.AwsuserNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.CloudAPI.AwsuserAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Awsuser(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateAwsuserResponseAsObject.GetResult()

	return mapNIOSAwsuserToResponse(&result), httpResp, nil
}

// Delete removes a Awsuser by ID
func (s *awsuserService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *awsuserService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.CloudAPI.AwsuserAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves Awsuser objects based on filter options
func (s *awsuserService) List(ctx context.Context, opts *core.ListOptions) ([]*cloud.Awsuser, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *awsuserService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*cloud.Awsuser, *http.Response, string, error) {
	req := s.niosClient.CloudAPI.AwsuserAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.AwsuserFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListAwsuserResponseObject.GetResult()
	items := make([]*cloud.Awsuser, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSAwsuserToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListAwsuserResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSAwsuserToResponse(r *nioscloud.Awsuser) *cloud.Awsuser {
	resp := &cloud.Awsuser{
		Id: r.Ref,
	}
	resp.NIOS = &cloud.NIOSAwsuserExt{
		AccessKeyId:     r.AccessKeyId,
		AccountId:       r.AccountId,
		GovcloudEnabled: r.GovcloudEnabled,
		Name:            r.Name,
		NiosUserName:    r.NiosUserName,
		SecretAccessKey: r.SecretAccessKey,
	}
	return resp
}
