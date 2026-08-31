package dns

import (
	"context"
	"fmt"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
)

type RecordAliasService interface {
	Create(ctx context.Context, obj *dns.RecordAlias, opts *core.Options) (*dns.RecordAlias, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dns.RecordAlias, *http.Response, error)
	Update(ctx context.Context, id string, obj *dns.RecordAlias, opts *core.Options) (*dns.RecordAlias, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dns.RecordAlias, *http.Response, string, error)
}

type recordAliasService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewRecordAliasService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) RecordAliasService {
	return &recordAliasService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new RecordAlias and returns the created object
func (s *recordAliasService) Create(ctx context.Context, obj *dns.RecordAlias, opts *core.Options) (*dns.RecordAlias, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordAliasService) createNIOS(ctx context.Context, obj *dns.RecordAlias, opts *core.Options) (*dns.RecordAlias, *http.Response, error) {
	payload, err := common.MapTo[niosdns.RecordAlias](obj, mapper.RecordAliasNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.RecordAliasAPI.
		Create(ctx).
		RecordAlias(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateRecordAliasResponseAsObject.GetResult()

	return mapNIOSRecordAliasToResponse(&result), httpResp, nil
}

// Read retrieves a RecordAlias by ID
func (s *recordAliasService) Read(ctx context.Context, id string, opts *core.Options) (*dns.RecordAlias, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordAliasService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dns.RecordAlias, *http.Response, error) {
	req := s.niosClient.DNSAPI.RecordAliasAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetRecordAliasResponseObjectAsResult.GetResult()

	return mapNIOSRecordAliasToResponse(&result), httpResp, nil
}

// Update modifies an existing RecordAlias and returns the updated object
func (s *recordAliasService) Update(ctx context.Context, id string, obj *dns.RecordAlias, opts *core.Options) (*dns.RecordAlias, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordAliasService) updateNIOS(ctx context.Context, id string, obj *dns.RecordAlias, opts *core.Options) (*dns.RecordAlias, *http.Response, error) {
	payload, err := common.MapTo[niosdns.RecordAlias](obj, mapper.RecordAliasNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.RecordAliasAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		RecordAlias(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateRecordAliasResponseAsObject.GetResult()

	return mapNIOSRecordAliasToResponse(&result), httpResp, nil
}

// Delete removes a RecordAlias by ID
func (s *recordAliasService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordAliasService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DNSAPI.RecordAliasAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves RecordAlias objects based on filter options
func (s *recordAliasService) List(ctx context.Context, opts *core.ListOptions) ([]*dns.RecordAlias, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordAliasService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dns.RecordAlias, *http.Response, string, error) {
	req := s.niosClient.DNSAPI.RecordAliasAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.RecordAliasFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListRecordAliasResponseObject.GetResult()
	items := make([]*dns.RecordAlias, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSRecordAliasToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListRecordAliasResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSRecordAliasToResponse(r *niosdns.RecordAlias) *dns.RecordAlias {
	resp := &dns.RecordAlias{
		Id: r.Ref,
	}
	resp.NIOS = &dns.NIOSRecordAliasExt{
		CloudInfo:  r.CloudInfo,
		Comment:    r.Comment,
		Creator:    r.Creator,
		Disable:    r.Disable,
		Name:       r.Name,
		TargetName: r.TargetName,
		TargetType: r.TargetType,
		Ttl:        r.Ttl,
		UseTtl:     r.UseTtl,
		View:       r.View,
	}
	if r.ExtAttrs != nil {
		attrs := make(map[string]any, len(*r.ExtAttrs))
		for k, v := range *r.ExtAttrs {
			attrs[k] = core.StringifyEAValue(v.Value)
		}
		resp.NIOS.ExtAttrs = attrs
	}
	return resp
}
