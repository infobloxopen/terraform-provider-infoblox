package dns

import (
	"context"
	"fmt"
	"maps"
	"net/http"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/common"
	mapper "github.com/infobloxopen/terraform-provider-infoblox/internal/core/mapper/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	uddidnsconfig "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

type ZoneAuthService interface {
	Create(ctx context.Context, obj *dns.ZoneAuth, opts *core.Options) (*dns.ZoneAuth, *http.Response, error)
	Read(ctx context.Context, id string, opts *core.Options) (*dns.ZoneAuth, *http.Response, error)
	Update(ctx context.Context, id string, obj *dns.ZoneAuth, opts *core.Options) (*dns.ZoneAuth, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	List(ctx context.Context, opts *core.ListOptions) ([]*dns.ZoneAuth, *http.Response, string, error)
}

type zoneAuthService struct {
	backend    core.BackendType
	niosClient *niosclient.APIClient
	uddiClient *uddiclient.APIClient
}

func NewZoneAuthService(backend core.BackendType, nios *niosclient.APIClient, uddi *uddiclient.APIClient) ZoneAuthService {
	return &zoneAuthService{
		backend:    backend,
		niosClient: nios,
		uddiClient: uddi,
	}
}

// Create creates a new ZoneAuth and returns the created object
func (s *zoneAuthService) Create(ctx context.Context, obj *dns.ZoneAuth, opts *core.Options) (*dns.ZoneAuth, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.createNIOS(ctx, obj, opts)
	case core.BackendUDDI:
		return s.createUDDI(ctx, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *zoneAuthService) createNIOS(ctx context.Context, obj *dns.ZoneAuth, opts *core.Options) (*dns.ZoneAuth, *http.Response, error) {
	payload, err := common.MapTo[niosdns.ZoneAuth](obj, mapper.ZoneAuthNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.ZoneAuthAPI.
		Create(ctx).
		ZoneAuth(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.CreateZoneAuthResponseAsObject.GetResult()

	return mapNIOSZoneAuthToResponse(&result), httpResp, nil
}

func (s *zoneAuthService) createUDDI(ctx context.Context, obj *dns.ZoneAuth, opts *core.Options) (*dns.ZoneAuth, *http.Response, error) {
	payload, err := common.MapTo[uddidnsconfig.AuthZone](obj, mapper.ZoneAuthUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSConfigurationAPI.AuthZoneAPI.
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

	return mapUDDIZoneAuthToResponse(&result), httpResp, nil
}

// Read retrieves a ZoneAuth by ID
func (s *zoneAuthService) Read(ctx context.Context, id string, opts *core.Options) (*dns.ZoneAuth, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.readNIOS(ctx, id, opts)
	case core.BackendUDDI:
		return s.readUDDI(ctx, id, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *zoneAuthService) readNIOS(ctx context.Context, id string, opts *core.Options) (*dns.ZoneAuth, *http.Response, error) {
	req := s.niosClient.DNSAPI.ZoneAuthAPI.
		Read(ctx, core.ExtractNIOSRef(id)).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetZoneAuthResponseObjectAsResult.GetResult()

	return mapNIOSZoneAuthToResponse(&result), httpResp, nil
}

func (s *zoneAuthService) readUDDI(ctx context.Context, id string, opts *core.Options) (*dns.ZoneAuth, *http.Response, error) {
	req := s.uddiClient.DNSConfigurationAPI.AuthZoneAPI.
		Read(ctx, id)

	if opts != nil && opts.Inherit != "" {
		req = req.Inherit(opts.Inherit)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.GetResult()

	return mapUDDIZoneAuthToResponse(&result), httpResp, nil
}

// Update modifies an existing ZoneAuth and returns the updated object
func (s *zoneAuthService) Update(ctx context.Context, id string, obj *dns.ZoneAuth, opts *core.Options) (*dns.ZoneAuth, *http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.updateNIOS(ctx, id, obj, opts)
	case core.BackendUDDI:
		return s.updateUDDI(ctx, id, obj, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *zoneAuthService) updateNIOS(ctx context.Context, id string, obj *dns.ZoneAuth, opts *core.Options) (*dns.ZoneAuth, *http.Response, error) {
	payload, err := common.MapTo[niosdns.ZoneAuth](obj, mapper.ZoneAuthNIOSFieldMap)
	if err != nil {
		return nil, nil, err
	}
	if obj.NIOS != nil && obj.NIOS.ExtAttrs != nil {
		if err := common.ProcessExtAttrs(obj.NIOS, &payload); err != nil {
			return nil, nil, err
		}
	}

	req := s.niosClient.DNSAPI.ZoneAuthAPI.
		Update(ctx, core.ExtractNIOSRef(id)).
		ZoneAuth(payload).
		ReturnAsObject(1)

	if opts != nil && opts.ReturnFields != "" {
		req = req.ReturnFieldsPlus(opts.ReturnFields)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	result := resp.UpdateZoneAuthResponseAsObject.GetResult()

	return mapNIOSZoneAuthToResponse(&result), httpResp, nil
}

func (s *zoneAuthService) updateUDDI(ctx context.Context, id string, obj *dns.ZoneAuth, opts *core.Options) (*dns.ZoneAuth, *http.Response, error) {
	payload, err := common.MapTo[uddidnsconfig.AuthZone](obj, mapper.ZoneAuthUDDIFieldMap)
	if err != nil {
		return nil, nil, err
	}

	req := s.uddiClient.DNSConfigurationAPI.AuthZoneAPI.
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

	return mapUDDIZoneAuthToResponse(&result), httpResp, nil
}

// Delete removes a ZoneAuth by ID
func (s *zoneAuthService) Delete(ctx context.Context, id string) (*http.Response, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.deleteNIOS(ctx, id)
	case core.BackendUDDI:
		return s.deleteUDDI(ctx, id)
	default:
		return nil, fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *zoneAuthService) deleteNIOS(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.niosClient.DNSAPI.ZoneAuthAPI.
		Delete(ctx, core.ExtractNIOSRef(id)).
		Execute()
	return httpResp, err
}

func (s *zoneAuthService) deleteUDDI(ctx context.Context, id string) (*http.Response, error) {
	httpResp, err := s.uddiClient.DNSConfigurationAPI.AuthZoneAPI.
		Delete(ctx, id).
		Execute()
	return httpResp, err
}

// List retrieves ZoneAuth objects based on filter options
func (s *zoneAuthService) List(ctx context.Context, opts *core.ListOptions) ([]*dns.ZoneAuth, *http.Response, string, error) {
	switch s.backend {
	case core.BackendNIOS:
		return s.listNIOS(ctx, opts)
	case core.BackendUDDI:
		return s.listUDDI(ctx, opts)
	default:
		return nil, nil, "", fmt.Errorf("unsupported backend: %s", s.backend)
	}
}

func (s *zoneAuthService) listNIOS(ctx context.Context, opts *core.ListOptions) ([]*dns.ZoneAuth, *http.Response, string, error) {
	req := s.niosClient.DNSAPI.ZoneAuthAPI.
		List(ctx).
		ReturnAsObject(1)

	if opts != nil {
		if opts.ReturnFields != "" {
			req = req.ReturnFieldsPlus(opts.ReturnFields)
		}
		if len(opts.Filters) > 0 {
			translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.ZoneAuthFilterFieldMap[core.BackendNIOS])
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

	results := resp.ListZoneAuthResponseObject.GetResult()
	items := make([]*dns.ZoneAuth, 0, len(results))
	for i := range results {
		items = append(items, mapNIOSZoneAuthToResponse(&results[i]))
	}

	var nextPageID string
	if ap := resp.ListZoneAuthResponseObject.AdditionalProperties; ap != nil {
		if npID, ok := ap["next_page_id"]; ok {
			if npIDStr, ok := npID.(string); ok {
				nextPageID = npIDStr
			}
		}
	}

	return items, httpResp, nextPageID, nil
}

func (s *zoneAuthService) listUDDI(ctx context.Context, opts *core.ListOptions) ([]*dns.ZoneAuth, *http.Response, string, error) {
	req := s.uddiClient.DNSConfigurationAPI.AuthZoneAPI.List(ctx)
	req = req.Limit(core.DefaultListLimit)

	if opts != nil {
		var filters []string
		for k, v := range opts.InternalFilters {
			filters = append(filters, core.FilterExpr(k, v))
		}
		translatedFilters := core.TranslateFilterKeys(opts.Filters, mapper.ZoneAuthFilterFieldMap[core.BackendUDDI])
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
	items := make([]*dns.ZoneAuth, 0, len(results))
	for i := range results {
		items = append(items, mapUDDIZoneAuthToResponse(&results[i]))
	}

	return items, httpResp, "", nil
}

func mapNIOSZoneAuthToResponse(r *niosdns.ZoneAuth) *dns.ZoneAuth {
	resp := &dns.ZoneAuth{
		Id: r.Ref,
	}
	resp.NIOS = &dns.NIOSZoneAuthExt{
		AllowActiveDir:                      r.AllowActiveDir,
		AllowFixedRrsetOrder:                r.AllowFixedRrsetOrder,
		AllowGssTsigForUnderscoreZone:       r.AllowGssTsigForUnderscoreZone,
		AllowGssTsigZoneUpdates:             r.AllowGssTsigZoneUpdates,
		AllowQuery:                          r.AllowQuery,
		AllowTransfer:                       r.AllowTransfer,
		AllowUpdate:                         r.AllowUpdate,
		AllowUpdateForwarding:               r.AllowUpdateForwarding,
		CloudInfo:                           r.CloudInfo,
		Comment:                             r.Comment,
		CopyXferToNotify:                    r.CopyXferToNotify,
		CreatePtrForBulkHosts:               r.CreatePtrForBulkHosts,
		CreatePtrForHosts:                   r.CreatePtrForHosts,
		CreateUnderscoreZones:               r.CreateUnderscoreZones,
		DdnsForceCreationTimestampUpdate:    r.DdnsForceCreationTimestampUpdate,
		DdnsPrincipalGroup:                  r.DdnsPrincipalGroup,
		DdnsPrincipalTracking:               r.DdnsPrincipalTracking,
		DdnsRestrictPatterns:                r.DdnsRestrictPatterns,
		DdnsRestrictPatternsList:            r.DdnsRestrictPatternsList,
		DdnsRestrictProtected:               r.DdnsRestrictProtected,
		DdnsRestrictSecure:                  r.DdnsRestrictSecure,
		DdnsRestrictStatic:                  r.DdnsRestrictStatic,
		Disable:                             r.Disable,
		DisableForwarding:                   r.DisableForwarding,
		DisplayDomain:                       r.DisplayDomain,
		DnsIntegrityEnable:                  r.DnsIntegrityEnable,
		DnsIntegrityFrequency:               r.DnsIntegrityFrequency,
		DnsIntegrityMember:                  r.DnsIntegrityMember,
		DnsIntegrityVerboseLogging:          r.DnsIntegrityVerboseLogging,
		DnssecKeyParams:                     r.DnssecKeyParams,
		DnssecKeys:                          r.DnssecKeys,
		DoHostAbstraction:                   r.DoHostAbstraction,
		EffectiveCheckNamesPolicy:           r.EffectiveCheckNamesPolicy,
		ExternalPrimaries:                   r.ExternalPrimaries,
		ExternalSecondaries:                 r.ExternalSecondaries,
		Fqdn:                                r.Fqdn,
		GridPrimary:                         r.GridPrimary,
		GridSecondaries:                     r.GridSecondaries,
		ImportFrom:                          r.ImportFrom,
		LastQueriedAcl:                      r.LastQueriedAcl,
		Locked:                              r.Locked,
		MemberSoaMnames:                     r.MemberSoaMnames,
		MsAdIntegrated:                      r.MsAdIntegrated,
		MsAllowTransfer:                     r.MsAllowTransfer,
		MsAllowTransferMode:                 r.MsAllowTransferMode,
		MsDcNsRecordCreation:                r.MsDcNsRecordCreation,
		MsDdnsMode:                          r.MsDdnsMode,
		MsPrimaries:                         r.MsPrimaries,
		MsSecondaries:                       r.MsSecondaries,
		MsSyncDisabled:                      r.MsSyncDisabled,
		NotifyDelay:                         r.NotifyDelay,
		NsGroup:                             r.NsGroup,
		Prefix:                              r.Prefix,
		RecordNamePolicy:                    r.RecordNamePolicy,
		RemoveSubzones:                      r.RemoveSubzones,
		RestartIfNeeded:                     r.RestartIfNeeded,
		ScavengingSettings:                  r.ScavengingSettings,
		SetSoaSerialNumber:                  r.SetSoaSerialNumber,
		SoaDefaultTtl:                       r.SoaDefaultTtl,
		SoaEmail:                            r.SoaEmail,
		SoaExpire:                           r.SoaExpire,
		SoaNegativeTtl:                      r.SoaNegativeTtl,
		SoaRefresh:                          r.SoaRefresh,
		SoaRetry:                            r.SoaRetry,
		SoaSerialNumber:                     r.SoaSerial,
		Srgs:                                r.Srgs,
		UpdateForwarding:                    r.UpdateForwarding,
		UseAllowActiveDir:                   r.UseAllowActiveDir,
		UseAllowQuery:                       r.UseAllowQuery,
		UseAllowTransfer:                    r.UseAllowTransfer,
		UseAllowUpdate:                      r.UseAllowUpdate,
		UseAllowUpdateForwarding:            r.UseAllowUpdateForwarding,
		UseCheckNamesPolicy:                 r.UseCheckNamesPolicy,
		UseCopyXferToNotify:                 r.UseCopyXferToNotify,
		UseDdnsForceCreationTimestampUpdate: r.UseDdnsForceCreationTimestampUpdate,
		UseDdnsPatternsRestriction:          r.UseDdnsPatternsRestriction,
		UseDdnsPrincipalSecurity:            r.UseDdnsPrincipalSecurity,
		UseDdnsRestrictProtected:            r.UseDdnsRestrictProtected,
		UseDdnsRestrictStatic:               r.UseDdnsRestrictStatic,
		UseDnssecKeyParams:                  r.UseDnssecKeyParams,
		UseExternalPrimary:                  r.UseExternalPrimary,
		UseGridZoneTimer:                    r.UseGridZoneTimer,
		UseImportFrom:                       r.UseImportFrom,
		UseNotifyDelay:                      r.UseNotifyDelay,
		UseRecordNamePolicy:                 r.UseRecordNamePolicy,
		UseScavengingSettings:               r.UseScavengingSettings,
		UseSoaEmail:                         r.UseSoaEmail,
		View:                                r.View,
		ZoneFormat:                          r.ZoneFormat,
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

func mapUDDIZoneAuthToResponse(r *uddidnsconfig.AuthZone) *dns.ZoneAuth {
	resp := &dns.ZoneAuth{
		Id: r.Id,
	}
	resp.UDDI = &dns.UDDIZoneAuthExt{
		Comment:                  r.Comment,
		CompartmentId:            r.CompartmentId,
		Disabled:                 r.Disabled,
		ExternalPrimaries:        r.ExternalPrimaries,
		ExternalSecondaries:      r.ExternalSecondaries,
		Fqdn:                     r.Fqdn,
		GssTsigEnabled:           r.GssTsigEnabled,
		InheritanceSources:       r.InheritanceSources,
		InitialSoaSerial:         r.InitialSoaSerial,
		InternalSecondaries:      r.InternalSecondaries,
		Notify:                   r.Notify,
		Nsgs:                     r.Nsgs,
		Parent:                   r.Parent,
		PrimaryType:              r.PrimaryType,
		QueryAcl:                 r.QueryAcl,
		TransferAcl:              r.TransferAcl,
		UpdateAcl:                r.UpdateAcl,
		UseForwardersForSubzones: r.UseForwardersForSubzones,
		View:                     r.View,
		ZoneAuthority:            r.ZoneAuthority,
	}
	if r.Tags != nil {
		tags := make(map[string]any, len(r.Tags))
		maps.Copy(tags, r.Tags)
		resp.UDDI.Tags = tags
	}
	return resp
}
