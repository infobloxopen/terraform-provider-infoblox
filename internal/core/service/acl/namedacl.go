package acl

import (
	"context"
	"fmt"
	"maps"
	"net/http"

	niosacl "github.com/infobloxopen/infoblox-nios-go-client/acl"
	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/acl"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/acl"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	uddidnsconfig "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

type NamedaclService interface {
	Create(ctx context.Context, obj *acl.Namedacl, opts *core.Options) (*acl.Namedacl, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*acl.Namedacl, *http.Response, error)
	Update(ctx context.Context, id string, obj *acl.Namedacl, opts *core.Options) (*acl.Namedacl, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*acl.Namedacl, *http.Response, string, error)
}

type namedaclService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
	uddiClient *uddiclient.APIClient
}

func NewNamedaclService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) NamedaclService {
	return &namedaclService{
		backend:    backend,
		niosClient: nios,
		uddiClient: uddi,
	}
}

// Create creates a new Namedacl and returns the created object
func (s *namedaclService) Create(ctx context.Context, obj *acl.Namedacl, opts *core.Options) (*acl.Namedacl, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *namedaclService) createNIOS(ctx context.Context, obj *acl.Namedacl, opts *core.Options) (*acl.Namedacl, *http.Response, error) {
	payload, err := common.MapTo[niosacl.Namedacl](obj, mapper.NamedaclNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.ACLAPI.NamedaclAPI.
		Create(ctx).
		Namedacl(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateNamedaclResponseAsObject.GetResult()

	return mapNIOSNamedaclToResponse(&result), httpResp, nil
}

func (s *namedaclService) createUDDI(ctx context.Context, obj *acl.Namedacl, opts *core.Options) (*acl.Namedacl, *http.Response, error) {
	payload, err := common.MapTo[uddidnsconfig.ACL](obj, mapper.NamedaclUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSConfigurationAPI.AclAPI.
		Create(ctx).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDINamedaclToResponse(&result), httpResp, nil
}

// Read retrieves a Namedacl by ID
func (s *namedaclService) Read(ctx context.Context, id string, opts *core.Options) (*acl.Namedacl, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *namedaclService) readNIOS(ctx context.Context, id string, opts *core.Options) (*acl.Namedacl, *http.Response, error) {
	req := s.niosClient.ACLAPI.NamedaclAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetNamedaclResponseObjectAsResult.GetResult()

	return mapNIOSNamedaclToResponse(&result), httpResp, nil
}

func (s *namedaclService) readUDDI(ctx context.Context, id string, opts *core.Options) (*acl.Namedacl, *http.Response, error) {
	req := s.uddiClient.DNSConfigurationAPI.AclAPI.
		Read(ctx, id)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDINamedaclToResponse(&result), httpResp, nil
}

// Update modifies an existing Namedacl and returns the updated object
func (s *namedaclService) Update(ctx context.Context, id string, obj *acl.Namedacl, opts *core.Options) (*acl.Namedacl, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *namedaclService) updateNIOS(ctx context.Context, id string, obj *acl.Namedacl, opts *core.Options) (*acl.Namedacl, *http.Response, error) {
	payload, err := common.MapTo[niosacl.Namedacl](obj, mapper.NamedaclNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.ACLAPI.NamedaclAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Namedacl(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateNamedaclResponseAsObject.GetResult()

	return mapNIOSNamedaclToResponse(&result), httpResp, nil
}

func (s *namedaclService) updateUDDI(ctx context.Context, id string, obj *acl.Namedacl, opts *core.Options) (*acl.Namedacl, *http.Response, error) {
	payload, err := common.MapTo[uddidnsconfig.ACL](obj, mapper.NamedaclUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSConfigurationAPI.AclAPI.
		Update(ctx, id).
		Body(payload)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDINamedaclToResponse(&result), httpResp, nil
}

// Delete removes a Namedacl by ID
func (s *namedaclService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *namedaclService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.ACLAPI.NamedaclAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

func (s *namedaclService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.DNSConfigurationAPI.AclAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves Namedacl objects based on filter options
func (s *namedaclService) List(ctx context.Context, opts *core.ListOptions) ([]*acl.Namedacl, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *namedaclService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*acl.Namedacl, *http.Response, string, error) {
	req := s.niosClient.ACLAPI.NamedaclAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.NamedaclFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListNamedaclResponseObject.GetResult()
	items := make([]*acl.Namedacl, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSNamedaclToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListNamedaclResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func (s *namedaclService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*acl.Namedacl, *http.Response, string, error) {
	req := s.uddiClient.DNSConfigurationAPI.AclAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.NamedaclFilterFieldMap[core.BackendUDDI])
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
	items := make([]*acl.Namedacl, 0, len(results))
	for i := range results {
		items = append(items, mapUDDINamedaclToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapNIOSNamedaclToResponse(r *niosacl.Namedacl) *acl.Namedacl {
	resp := &acl.Namedacl{
		Id: r.Ref,
	}
	resp.NIOS = &acl.NIOSNamedaclExt{
		AccessList: r.AccessList,
		Comment:    r.Comment,
		Name:       r.Name,
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

func mapUDDINamedaclToResponse(r *uddidnsconfig.ACL) *acl.Namedacl {
	resp := &acl.Namedacl{
		Id: r.Id,
	}
	resp.UDDI = &acl.UDDINamedaclExt{
		Comment:       r.Comment,
		CompartmentId: r.CompartmentId,
		List:          r.List,
		Name:          r.Name,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
