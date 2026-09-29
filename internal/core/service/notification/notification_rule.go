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

type NotificationRuleService interface {
	Create(ctx context.Context, obj *notification.NotificationRule, opts *core.Options) (*notification.NotificationRule, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*notification.NotificationRule, *http.Response, error)
	Update(ctx context.Context, id string, obj *notification.NotificationRule, opts *core.Options) (*notification.NotificationRule, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*notification.NotificationRule, *http.Response, string, error)
}

type notificationRuleService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewNotificationRuleService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) NotificationRuleService {
	return &notificationRuleService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new NotificationRule and returns the created object
func (s *notificationRuleService) Create(ctx context.Context, obj *notification.NotificationRule, opts *core.Options) (*notification.NotificationRule, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *notificationRuleService) createNIOS(ctx context.Context, obj *notification.NotificationRule, opts *core.Options) (*notification.NotificationRule, *http.Response, error) {
	payload, err := common.MapTo[niosnotification.NotificationRule](obj, mapper.NotificationRuleNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.NotificationAPI.NotificationRuleAPI.
		Create(ctx).
		NotificationRule(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateNotificationRuleResponseAsObject.GetResult()

	return mapNIOSNotificationRuleToResponse(&result), httpResp, nil
}

// Read retrieves a NotificationRule by ID
func (s *notificationRuleService) Read(ctx context.Context, id string, opts *core.Options) (*notification.NotificationRule, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *notificationRuleService) readNIOS(ctx context.Context, id string, opts *core.Options) (*notification.NotificationRule, *http.Response, error) {
	req := s.niosClient.NotificationAPI.NotificationRuleAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetNotificationRuleResponseObjectAsResult.GetResult()

	return mapNIOSNotificationRuleToResponse(&result), httpResp, nil
}

// Update modifies an existing NotificationRule and returns the updated object
func (s *notificationRuleService) Update(ctx context.Context, id string, obj *notification.NotificationRule, opts *core.Options) (*notification.NotificationRule, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *notificationRuleService) updateNIOS(ctx context.Context, id string, obj *notification.NotificationRule, opts *core.Options) (*notification.NotificationRule, *http.Response, error) {
	payload, err := common.MapTo[niosnotification.NotificationRule](obj, mapper.NotificationRuleNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.niosClient.NotificationAPI.NotificationRuleAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		NotificationRule(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateNotificationRuleResponseAsObject.GetResult()

	return mapNIOSNotificationRuleToResponse(&result), httpResp, nil
}

// Delete removes a NotificationRule by ID
func (s *notificationRuleService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *notificationRuleService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.NotificationAPI.NotificationRuleAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves NotificationRule objects based on filter options
func (s *notificationRuleService) List(ctx context.Context, opts *core.ListOptions) ([]*notification.NotificationRule, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *notificationRuleService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*notification.NotificationRule, *http.Response, string, error) {
	req := s.niosClient.NotificationAPI.NotificationRuleAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.NotificationRuleFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListNotificationRuleResponseObject.GetResult()
	items := make([]*notification.NotificationRule, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSNotificationRuleToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListNotificationRuleResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSNotificationRuleToResponse(r *niosnotification.NotificationRule) *notification.NotificationRule {
	resp := &notification.NotificationRule{
		Id: r.Ref,
	}
	resp.NIOS = &notification.NIOSNotificationRuleExt{
		AllMembers:                       r.AllMembers,
		Comment:                          r.Comment,
		Disable:                          r.Disable,
		EnableEventDeduplication:         r.EnableEventDeduplication,
		EnableEventDeduplicationLog:      r.EnableEventDeduplicationLog,
		EventDeduplicationFields:         r.EventDeduplicationFields,
		EventDeduplicationLookbackPeriod: r.EventDeduplicationLookbackPeriod,
		EventPriority:                    r.EventPriority,
		EventType:                        r.EventType,
		ExpressionList:                   r.ExpressionList,
		Name:                             r.Name,
		NotificationAction:               r.NotificationAction,
		NotificationTarget:               r.NotificationTarget,
		PublishSettings:                  r.PublishSettings,
		ScheduledEvent:                   r.ScheduledEvent,
		SelectedMembers:                  r.SelectedMembers,
		TemplateInstance:                 r.TemplateInstance,
		UsePublishSettings:               r.UsePublishSettings,
	}
	return resp
}
