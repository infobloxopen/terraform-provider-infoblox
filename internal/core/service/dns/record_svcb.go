package dns

import (
	"context"
	"fmt"
	"maps"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	uddidnsdata "github.com/infobloxopen/universal-ddi-go-client/dnsdata"
)

type RecordSvcbService interface {
	Create(ctx context.Context, obj *dns.RecordSvcb, opts *core.Options) (*dns.RecordSvcb, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dns.RecordSvcb, *http.Response, error)
	Update(ctx context.Context, id string, obj *dns.RecordSvcb, opts *core.Options) (*dns.RecordSvcb, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dns.RecordSvcb, *http.Response, string, error)
}

type recordSvcbService struct {
	backend    core.BackendType
	uddiClient *uddiclient.APIClient
}

func NewRecordSvcbService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) RecordSvcbService {
	return &recordSvcbService{
		backend:    backend,
		uddiClient: uddi,
	}
}

// Create creates a new RecordSvcb and returns the created object
func (s *recordSvcbService) Create(ctx context.Context, obj *dns.RecordSvcb, opts *core.Options) (*dns.RecordSvcb, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordSvcbService) createUDDI(ctx context.Context, obj *dns.RecordSvcb, opts *core.Options) (*dns.RecordSvcb, *http.Response, error) {
	payload, err := common.MapTo[uddidnsdata.Record](obj, mapper.RecordSvcbUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSDataAPI.RecordAPI.
		Create(ctx).
		Body(payload)

	if opts != nil && opts.Inherit != "" {
		req = req.Inherit(opts.Inherit)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIRecordSvcbToResponse(&result), httpResp, nil
}

// Read retrieves a RecordSvcb by ID
func (s *recordSvcbService) Read(ctx context.Context, id string, opts *core.Options) (*dns.RecordSvcb, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordSvcbService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dns.RecordSvcb, *http.Response, error) {
	req := s.uddiClient.DNSDataAPI.RecordAPI.
		Read(ctx, id)

	if opts != nil && opts.Inherit != "" {
		req = req.Inherit(opts.Inherit)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIRecordSvcbToResponse(&result), httpResp, nil
}

// Update modifies an existing RecordSvcb and returns the updated object
func (s *recordSvcbService) Update(ctx context.Context, id string, obj *dns.RecordSvcb, opts *core.Options) (*dns.RecordSvcb, *http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordSvcbService) updateUDDI(ctx context.Context, id string, obj *dns.RecordSvcb, opts *core.Options) (*dns.RecordSvcb, *http.Response, error) {
	payload, err := common.MapTo[uddidnsdata.Record](obj, mapper.RecordSvcbUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSDataAPI.RecordAPI.
		Update(ctx, id).
		Body(payload)

	if opts != nil && opts.Inherit != "" {
		req = req.Inherit(opts.Inherit)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIRecordSvcbToResponse(&result), httpResp, nil
}

// Delete removes a RecordSvcb by ID
func (s *recordSvcbService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordSvcbService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.DNSDataAPI.RecordAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves RecordSvcb objects based on filter options
func (s *recordSvcbService) List(ctx context.Context, opts *core.ListOptions) ([]*dns.RecordSvcb, *http.Response, string, error) {
	switch s.backend {
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *recordSvcbService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dns.RecordSvcb, *http.Response, string, error) {
	req := s.uddiClient.DNSDataAPI.RecordAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.RecordSvcbFilterFieldMap[core.BackendUDDI])
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
	items := make([]*dns.RecordSvcb, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIRecordSvcbToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapUDDIRecordSvcbToResponse(r *uddidnsdata.Record) *dns.RecordSvcb {
	resp := &dns.RecordSvcb{
		Id: r.Id,
	}
	resp.UDDI = &dns.UDDIRecordSvcbExt{
		AbsoluteNameSpec:   r.AbsoluteNameSpec,
		Comment:            r.Comment,
		Disabled:           r.Disabled,
		InheritanceSources: r.InheritanceSources,
		NameInZone:         r.NameInZone,
		Options:            r.Options,
		Rdata:              r.Rdata,
		Ttl:                r.Ttl,
		Type:               r.Type,
		View:               r.View,
		Zone:               r.Zone,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
