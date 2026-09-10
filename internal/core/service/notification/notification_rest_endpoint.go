package notification

import (
	"context"
	"fmt"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niosnotification "github.com/infobloxopen/infoblox-nios-go-client/notification"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/notification"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/notification"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
)

type NotificationRestEndpointService interface {
	Create(ctx context.Context, obj *notification.NotificationRestEndpoint, opts *core.Options) (*notification.NotificationRestEndpoint, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*notification.NotificationRestEndpoint, *http.Response, error)
	Update(ctx context.Context, id string, obj *notification.NotificationRestEndpoint, opts *core.Options) (*notification.NotificationRestEndpoint, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*notification.NotificationRestEndpoint, *http.Response, string, error)
}

type notificationRestEndpointService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewNotificationRestEndpointService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) NotificationRestEndpointService {
	return &notificationRestEndpointService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new NotificationRestEndpoint and returns the created object
func (s *notificationRestEndpointService) Create(ctx context.Context, obj *notification.NotificationRestEndpoint, opts *core.Options) (*notification.NotificationRestEndpoint, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *notificationRestEndpointService) createNIOS(ctx context.Context, obj *notification.NotificationRestEndpoint, opts *core.Options) (*notification.NotificationRestEndpoint, *http.Response, error) {
	payload, err := common.MapTo[niosnotification.NotificationRestEndpoint](obj, mapper.NotificationRestEndpointNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.NotificationAPI.NotificationRestEndpointAPI.
		Create(ctx).
		NotificationRestEndpoint(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateNotificationRestEndpointResponseAsObject.GetResult()

	return mapNIOSNotificationRestEndpointToResponse(&result), httpResp, nil
}

// Read retrieves a NotificationRestEndpoint by ID
func (s *notificationRestEndpointService) Read(ctx context.Context, id string, opts *core.Options) (*notification.NotificationRestEndpoint, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *notificationRestEndpointService) readNIOS(ctx context.Context, id string, opts *core.Options) (*notification.NotificationRestEndpoint, *http.Response, error) {
	req := s.niosClient.NotificationAPI.NotificationRestEndpointAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetNotificationRestEndpointResponseObjectAsResult.GetResult()

	return mapNIOSNotificationRestEndpointToResponse(&result), httpResp, nil
}

// Update modifies an existing NotificationRestEndpoint and returns the updated object
func (s *notificationRestEndpointService) Update(ctx context.Context, id string, obj *notification.NotificationRestEndpoint, opts *core.Options) (*notification.NotificationRestEndpoint, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *notificationRestEndpointService) updateNIOS(ctx context.Context, id string, obj *notification.NotificationRestEndpoint, opts *core.Options) (*notification.NotificationRestEndpoint, *http.Response, error) {
	payload, err := common.MapTo[niosnotification.NotificationRestEndpoint](obj, mapper.NotificationRestEndpointNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.NotificationAPI.NotificationRestEndpointAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		NotificationRestEndpoint(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateNotificationRestEndpointResponseAsObject.GetResult()

	return mapNIOSNotificationRestEndpointToResponse(&result), httpResp, nil
}

// Delete removes a NotificationRestEndpoint by ID
func (s *notificationRestEndpointService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *notificationRestEndpointService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.NotificationAPI.NotificationRestEndpointAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves NotificationRestEndpoint objects based on filter options
func (s *notificationRestEndpointService) List(ctx context.Context, opts *core.ListOptions) ([]*notification.NotificationRestEndpoint, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *notificationRestEndpointService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*notification.NotificationRestEndpoint, *http.Response, string, error) {
	req := s.niosClient.NotificationAPI.NotificationRestEndpointAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.NotificationRestEndpointFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListNotificationRestEndpointResponseObject.GetResult()
	items := make([]*notification.NotificationRestEndpoint, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSNotificationRestEndpointToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListNotificationRestEndpointResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSNotificationRestEndpointToResponse(r *niosnotification.NotificationRestEndpoint) *notification.NotificationRestEndpoint {
	resp := &notification.NotificationRestEndpoint{
		Id: r.Ref,
	}
	resp.NIOS = &notification.NIOSNotificationRestEndpointExt{
		ClientCertificateToken: r.ClientCertificateToken,
		Comment:                r.Comment,
		LogLevel:               r.LogLevel,
		Name:                   r.Name,
		OutboundMemberType:     r.OutboundMemberType,
		OutboundMembers:        r.OutboundMembers,
		Password:               r.Password,
		ServerCertValidation:   r.ServerCertValidation,
		SyncDisabled:           r.SyncDisabled,
		TemplateInstance:       r.TemplateInstance,
		Timeout:                r.Timeout,
		Uri:                    r.Uri,
		Username:               r.Username,
		VendorIdentifier:       r.VendorIdentifier,
		WapiUserName:           r.WapiUserName,
		WapiUserPassword:       r.WapiUserPassword,
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
