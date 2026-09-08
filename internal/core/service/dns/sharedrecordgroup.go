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

type SharedrecordgroupService interface {
	Create(ctx context.Context, obj *dns.Sharedrecordgroup, opts *core.Options) (*dns.Sharedrecordgroup, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dns.Sharedrecordgroup, *http.Response, error)
	Update(ctx context.Context, id string, obj *dns.Sharedrecordgroup, opts *core.Options) (*dns.Sharedrecordgroup, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dns.Sharedrecordgroup, *http.Response, string, error)
}

type sharedrecordgroupService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewSharedrecordgroupService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) SharedrecordgroupService {
	return &sharedrecordgroupService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new Sharedrecordgroup and returns the created object
func (s *sharedrecordgroupService) Create(ctx context.Context, obj *dns.Sharedrecordgroup, opts *core.Options) (*dns.Sharedrecordgroup, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordgroupService) createNIOS(ctx context.Context, obj *dns.Sharedrecordgroup, opts *core.Options) (*dns.Sharedrecordgroup, *http.Response, error) {
	payload, err := common.MapTo[niosdns.Sharedrecordgroup](obj, mapper.SharedrecordgroupNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.SharedrecordgroupAPI.
		Create(ctx).
		Sharedrecordgroup(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateSharedrecordgroupResponseAsObject.GetResult()

	return mapNIOSSharedrecordgroupToResponse(&result), httpResp, nil
}

// Read retrieves a Sharedrecordgroup by ID
func (s *sharedrecordgroupService) Read(ctx context.Context, id string, opts *core.Options) (*dns.Sharedrecordgroup, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordgroupService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dns.Sharedrecordgroup, *http.Response, error) {
	req := s.niosClient.DNSAPI.SharedrecordgroupAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetSharedrecordgroupResponseObjectAsResult.GetResult()

	return mapNIOSSharedrecordgroupToResponse(&result), httpResp, nil
}

// Update modifies an existing Sharedrecordgroup and returns the updated object
func (s *sharedrecordgroupService) Update(ctx context.Context, id string, obj *dns.Sharedrecordgroup, opts *core.Options) (*dns.Sharedrecordgroup, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordgroupService) updateNIOS(ctx context.Context, id string, obj *dns.Sharedrecordgroup, opts *core.Options) (*dns.Sharedrecordgroup, *http.Response, error) {
	payload, err := common.MapTo[niosdns.Sharedrecordgroup](obj, mapper.SharedrecordgroupNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.SharedrecordgroupAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Sharedrecordgroup(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateSharedrecordgroupResponseAsObject.GetResult()

	return mapNIOSSharedrecordgroupToResponse(&result), httpResp, nil
}

// Delete removes a Sharedrecordgroup by ID
func (s *sharedrecordgroupService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordgroupService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DNSAPI.SharedrecordgroupAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves Sharedrecordgroup objects based on filter options
func (s *sharedrecordgroupService) List(ctx context.Context, opts *core.ListOptions) ([]*dns.Sharedrecordgroup, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *sharedrecordgroupService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dns.Sharedrecordgroup, *http.Response, string, error) {
	req := s.niosClient.DNSAPI.SharedrecordgroupAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.SharedrecordgroupFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListSharedrecordgroupResponseObject.GetResult()
	items := make([]*dns.Sharedrecordgroup, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSSharedrecordgroupToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListSharedrecordgroupResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSSharedrecordgroupToResponse(r *niosdns.Sharedrecordgroup) *dns.Sharedrecordgroup {
	resp := &dns.Sharedrecordgroup{
		Id: r.Ref,
	}
	resp.NIOS = &dns.NIOSSharedrecordgroupExt{
		Comment:             r.Comment,
		Name:                r.Name,
		RecordNamePolicy:    r.RecordNamePolicy,
		UseRecordNamePolicy: r.UseRecordNamePolicy,
		ZoneAssociations:    r.ZoneAssociations,
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
