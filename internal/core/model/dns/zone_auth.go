package dns

import (
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	uddidnsconfig "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// Infoblox ZoneAuth model
type ZoneAuth struct {
	Id   *string
	NIOS *NIOSZoneAuthExt
	UDDI *UDDIZoneAuthExt
}

// NIOSZoneAuthExt - NIOS specific fields for ZoneAuth
type NIOSZoneAuthExt struct {
	AllowActiveDir                      []niosdns.ZoneAuthAllowActiveDir
	AllowFixedRrsetOrder                *bool
	AllowGssTsigForUnderscoreZone       *bool
	AllowGssTsigZoneUpdates             *bool
	AllowQuery                          []niosdns.ZoneAuthAllowQuery
	AllowTransfer                       []niosdns.ZoneAuthAllowTransfer
	AllowUpdate                         []niosdns.ZoneAuthAllowUpdate
	AllowUpdateForwarding               *bool
	CloudInfo                           *niosdns.ZoneAuthCloudInfo
	Comment                             *string
	CopyXferToNotify                    *bool
	CreatePtrForBulkHosts               *bool
	CreatePtrForHosts                   *bool
	CreateUnderscoreZones               *bool
	DdnsForceCreationTimestampUpdate    *bool
	DdnsPrincipalGroup                  *string
	DdnsPrincipalTracking               *bool
	DdnsRestrictPatterns                *bool
	DdnsRestrictPatternsList            []string
	DdnsRestrictProtected               *bool
	DdnsRestrictSecure                  *bool
	DdnsRestrictStatic                  *bool
	Disable                             *bool
	DisableForwarding                   *bool
	DnsIntegrityEnable                  *bool
	DnsIntegrityFrequency               *int64
	DnsIntegrityMember                  *string
	DnsIntegrityVerboseLogging          *bool
	DnssecKeyParams                     *niosdns.ZoneAuthDnssecKeyParams
	DnssecKeys                          []niosdns.ZoneAuthDnssecKeys
	DoHostAbstraction                   *bool
	EffectiveCheckNamesPolicy           *string
	ExtAttrs                            map[string]any
	ExternalPrimaries                   []niosdns.ZoneAuthExternalPrimaries
	ExternalSecondaries                 []niosdns.ZoneAuthExternalSecondaries
	Fqdn                                *string
	GridPrimary                         []niosdns.ZoneAuthGridPrimary
	GridSecondaries                     []niosdns.ZoneAuthGridSecondaries
	ImportFrom                          *string
	LastQueriedAcl                      []niosdns.ZoneAuthLastQueriedAcl
	Locked                              *bool
	MemberSoaMnames                     []niosdns.ZoneAuthMemberSoaMnames
	MsAdIntegrated                      *bool
	MsAllowTransfer                     []niosdns.ZoneAuthMsAllowTransfer
	MsAllowTransferMode                 *string
	MsDcNsRecordCreation                []niosdns.ZoneAuthMsDcNsRecordCreation
	MsDdnsMode                          *string
	MsPrimaries                         []niosdns.ZoneAuthMsPrimaries
	MsSecondaries                       []niosdns.ZoneAuthMsSecondaries
	MsSyncDisabled                      *bool
	NotifyDelay                         *int64
	NsGroup                             *string
	Prefix                              *string
	RecordNamePolicy                    *string
	RemoveSubzones                      *bool
	RestartIfNeeded                     *bool
	ScavengingSettings                  *niosdns.ZoneAuthScavengingSettings
	SetSoaSerialNumber                  *bool
	SoaDefaultTtl                       *int64
	SoaEmail                            *string
	SoaExpire                           *int64
	SoaNegativeTtl                      *int64
	SoaRefresh                          *int64
	SoaRetry                            *int64
	SoaSerialNumber                     *int64
	Srgs                                []string
	UpdateForwarding                    []niosdns.ZoneAuthUpdateForwarding
	UseAllowActiveDir                   *bool
	UseAllowQuery                       *bool
	UseAllowTransfer                    *bool
	UseAllowUpdate                      *bool
	UseAllowUpdateForwarding            *bool
	UseCheckNamesPolicy                 *bool
	UseCopyXferToNotify                 *bool
	UseDdnsForceCreationTimestampUpdate *bool
	UseDdnsPatternsRestriction          *bool
	UseDdnsPrincipalSecurity            *bool
	UseDdnsRestrictProtected            *bool
	UseDdnsRestrictStatic               *bool
	UseDnssecKeyParams                  *bool
	UseExternalPrimary                  *bool
	UseGridZoneTimer                    *bool
	UseImportFrom                       *bool
	UseNotifyDelay                      *bool
	UseRecordNamePolicy                 *bool
	UseScavengingSettings               *bool
	UseSoaEmail                         *bool
	View                                *string
	ZoneFormat                          *string
}

// UDDIZoneAuthExt - UDDI specific fields for ZoneAuth
type UDDIZoneAuthExt struct {
	Comment                  *string
	CompartmentId            *string
	Disabled                 *bool
	ExternalPrimaries        []uddidnsconfig.ExternalPrimary
	ExternalSecondaries      []uddidnsconfig.ExternalSecondary
	Fqdn                     *string
	GssTsigEnabled           *bool
	InheritanceSources       *uddidnsconfig.AuthZoneInheritance
	InitialSoaSerial         *int64
	InternalSecondaries      []uddidnsconfig.InternalSecondary
	Notify                   *bool
	Nsgs                     []string
	Parent                   *string
	PrimaryType              *string
	QueryAcl                 []uddidnsconfig.ACLItem
	Tags                     map[string]any
	TransferAcl              []uddidnsconfig.ACLItem
	UpdateAcl                []uddidnsconfig.ACLItem
	UseForwardersForSubzones *bool
	View                     *string
	ZoneAuthority            *uddidnsconfig.ZoneAuthority
}
