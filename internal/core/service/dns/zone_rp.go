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

type ZoneRpService interface {
	Create(ctx context.Context, obj *dns.ZoneRp, opts *core.Options) (*dns.ZoneRp, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dns.ZoneRp, *http.Response, error)
	Update(ctx context.Context, id string, obj *dns.ZoneRp, opts *core.Options) (*dns.ZoneRp, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dns.ZoneRp, *http.Response, string, error)
}

type zoneRpService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
}

func NewZoneRpService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) ZoneRpService {
	return &zoneRpService{
		backend:    backend,
		niosClient: nios,
	}
}

// Create creates a new ZoneRp and returns the created object
func (s *zoneRpService) Create(ctx context.Context, obj *dns.ZoneRp, opts *core.Options) (*dns.ZoneRp, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *zoneRpService) createNIOS(ctx context.Context, obj *dns.ZoneRp, opts *core.Options) (*dns.ZoneRp, *http.Response, error) {
	payload, err := common.MapTo[niosdns.ZoneRp](obj, mapper.ZoneRpNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.ZoneRpAPI.
		Create(ctx).
		ZoneRp(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateZoneRpResponseAsObject.GetResult()

	return mapNIOSZoneRpToResponse(&result), httpResp, nil
}

// Read retrieves a ZoneRp by ID
func (s *zoneRpService) Read(ctx context.Context, id string, opts *core.Options) (*dns.ZoneRp, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *zoneRpService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dns.ZoneRp, *http.Response, error) {
	req := s.niosClient.DNSAPI.ZoneRpAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetZoneRpResponseObjectAsResult.GetResult()

	return mapNIOSZoneRpToResponse(&result), httpResp, nil
}

// Update modifies an existing ZoneRp and returns the updated object
func (s *zoneRpService) Update(ctx context.Context, id string, obj *dns.ZoneRp, opts *core.Options) (*dns.ZoneRp, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *zoneRpService) updateNIOS(ctx context.Context, id string, obj *dns.ZoneRp, opts *core.Options) (*dns.ZoneRp, *http.Response, error) {
	payload, err := common.MapTo[niosdns.ZoneRp](obj, mapper.ZoneRpNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.ZoneRpAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		ZoneRp(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateZoneRpResponseAsObject.GetResult()

	return mapNIOSZoneRpToResponse(&result), httpResp, nil
}

// Delete removes a ZoneRp by ID
func (s *zoneRpService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *zoneRpService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DNSAPI.ZoneRpAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

// List retrieves ZoneRp objects based on filter options
func (s *zoneRpService) List(ctx context.Context, opts *core.ListOptions) ([]*dns.ZoneRp, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *zoneRpService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dns.ZoneRp, *http.Response, string, error) {
	req := s.niosClient.DNSAPI.ZoneRpAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.ZoneRpFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListZoneRpResponseObject.GetResult()
	items := make([]*dns.ZoneRp, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSZoneRpToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListZoneRpResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func mapNIOSZoneRpToResponse(r *niosdns.ZoneRp) *dns.ZoneRp {
	resp := &dns.ZoneRp{
		Id: r.Ref,
	}
	resp.NIOS = &dns.NIOSZoneRpExt{
		Comment:                          r.Comment,
		Disable:                          r.Disable,
		ExternalPrimaries:                r.ExternalPrimaries,
		ExternalSecondaries:              r.ExternalSecondaries,
		FireeyeRuleMapping:               r.FireeyeRuleMapping,
		Fqdn:                             r.Fqdn,
		GridPrimary:                      r.GridPrimary,
		GridSecondaries:                  r.GridSecondaries,
		Locked:                           r.Locked,
		LogRpz:                           r.LogRpz,
		MemberSoaMnames:                  r.MemberSoaMnames,
		NsGroup:                          r.NsGroup,
		Prefix:                           r.Prefix,
		RecordNamePolicy:                 r.RecordNamePolicy,
		RpzDropIpRuleEnabled:             r.RpzDropIpRuleEnabled,
		RpzDropIpRuleMinPrefixLengthIpv4: r.RpzDropIpRuleMinPrefixLengthIpv4,
		RpzDropIpRuleMinPrefixLengthIpv6: r.RpzDropIpRuleMinPrefixLengthIpv6,
		RpzPolicy:                        r.RpzPolicy,
		RpzSeverity:                      r.RpzSeverity,
		RpzType:                          r.RpzType,
		SetSoaSerialNumber:               r.SetSoaSerialNumber,
		SoaDefaultTtl:                    r.SoaDefaultTtl,
		SoaEmail:                         r.SoaEmail,
		SoaExpire:                        r.SoaExpire,
		SoaNegativeTtl:                   r.SoaNegativeTtl,
		SoaRefresh:                       r.SoaRefresh,
		SoaRetry:                         r.SoaRetry,
		SoaSerialNumber:                  r.SoaSerial,
		SubstituteName:                   r.SubstituteName,
		UseExternalPrimary:               r.UseExternalPrimary,
		UseGridZoneTimer:                 r.UseGridZoneTimer,
		UseLogRpz:                        r.UseLogRpz,
		UseRecordNamePolicy:              r.UseRecordNamePolicy,
		UseRpzDropIpRule:                 r.UseRpzDropIpRule,
		UseSoaEmail:                      r.UseSoaEmail,
		View:                             r.View,
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
