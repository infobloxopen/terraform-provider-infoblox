package security

import (
	"context"
	"fmt"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niossecurity "github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/security"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/security"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
)

type AdminuserService interface {
	Create(ctx context.Context, obj *security.Adminuser, opts *core.Options) (*security.Adminuser, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*security.Adminuser, *http.Response, error)
	Update(ctx context.Context, id string, obj *security.Adminuser, opts *core.Options) (*security.Adminuser, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*security.Adminuser, *http.Response, string, error)
}

type adminuserService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewAdminuserService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) AdminuserService {
	return &adminuserService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new Adminuser and returns the created object
func (s *adminuserService) Create(ctx context.Context, obj *security.Adminuser, opts *core.Options) (*security.Adminuser, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *adminuserService) createNIOS(ctx context.Context, obj *security.Adminuser, opts *core.Options) (*security.Adminuser, *http.Response, error) {
	payload, err := common.MapTo[niossecurity.Adminuser](obj, mapper.AdminuserNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.SecurityAPI.AdminuserAPI.
		Create(ctx).
		Adminuser(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateAdminuserResponseAsObject.GetResult()

	return mapNIOSAdminuserToResponse(&result), httpResp, nil
}

// Read retrieves a Adminuser by ID
func (s *adminuserService) Read(ctx context.Context, id string, opts *core.Options) (*security.Adminuser, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *adminuserService) readNIOS(ctx context.Context, id string, opts *core.Options) (*security.Adminuser, *http.Response, error) {
	req := s.niosClient.SecurityAPI.AdminuserAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	req = req.ProxySearch(core.GetProxySearch())
	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetAdminuserResponseObjectAsResult.GetResult()

	return mapNIOSAdminuserToResponse(&result), httpResp, nil
}

// Update modifies an existing Adminuser and returns the updated object
func (s *adminuserService) Update(ctx context.Context, id string, obj *security.Adminuser, opts *core.Options) (*security.Adminuser, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *adminuserService) updateNIOS(ctx context.Context, id string, obj *security.Adminuser, opts *core.Options) (*security.Adminuser, *http.Response, error) {
	payload, err := common.MapTo[niossecurity.Adminuser](obj, mapper.AdminuserNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.SecurityAPI.AdminuserAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		Adminuser(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateAdminuserResponseAsObject.GetResult()

	return mapNIOSAdminuserToResponse(&result), httpResp, nil
}

// Delete removes a Adminuser by ID
func (s *adminuserService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *adminuserService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.SecurityAPI.AdminuserAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves Adminuser objects based on filter options
func (s *adminuserService) List(ctx context.Context, opts *core.ListOptions) ([]*security.Adminuser, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *adminuserService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*security.Adminuser, *http.Response, string, error) {
	req := s.niosClient.SecurityAPI.AdminuserAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.AdminuserFilterFieldMap[core.BackendNIOS])
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

	req = req.ProxySearch(core.GetProxySearch())
	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, "", err
	}

	results := resp.ListAdminuserResponseObject.GetResult()
	items := make([]*security.Adminuser, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSAdminuserToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListAdminuserResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSAdminuserToResponse(r *niossecurity.Adminuser) *security.Adminuser {
	resp := &security.Adminuser{
		Id: r.Ref,
	}
	resp.NIOS = &security.NIOSAdminuserExt{
		AdminGroups:                     r.AdminGroups,
		AuthMethod:                      r.AuthMethod,
		AuthType:                        r.AuthType,
		CaCertificateIssuer:             r.CaCertificateIssuer,
		ClientCertificateSerialNumber:   r.ClientCertificateSerialNumber,
		Comment:                         r.Comment,
		Disable:                         r.Disable,
		Email:                           r.Email,
		EnableCertificateAuthentication: r.EnableCertificateAuthentication,
		Name:                            r.Name,
		Password:                        r.Password,
		SshKeys:                         r.SshKeys,
		TimeZone:                        r.TimeZone,
		UseSshKeys:                      r.UseSshKeys,
		UseTimeZone:                     r.UseTimeZone,
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
