package dns

import (
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	uddidnsconfig "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// Infoblox View model
type View struct {
	Id   *string
	NIOS *NIOSViewExt
	UDDI *UDDIViewExt
}

// NIOSViewExt - NIOS specific fields for View
type NIOSViewExt struct {
	BlacklistAction                     *string
	BlacklistLogQuery                   *bool
	BlacklistRedirectAddresses          []string
	BlacklistRedirectTtl                *int64
	BlacklistRulesets                   []string
	CloudInfo                           *niosdns.ViewCloudInfo
	Comment                             *string
	CustomRootNameServers               []niosdns.ViewCustomRootNameServers
	DdnsForceCreationTimestampUpdate    *bool
	DdnsPrincipalGroup                  *string
	DdnsPrincipalTracking               *bool
	DdnsRestrictPatterns                *bool
	DdnsRestrictPatternsList            []string
	DdnsRestrictProtected               *bool
	DdnsRestrictSecure                  *bool
	DdnsRestrictStatic                  *bool
	Disable                             *bool
	Dns64Enabled                        *bool
	Dns64Groups                         []string
	DnssecEnabled                       *bool
	DnssecExpiredSignaturesEnabled      *bool
	DnssecNegativeTrustAnchors          []string
	DnssecTrustedKeys                   []niosdns.ViewDnssecTrustedKeys
	DnssecValidationEnabled             *bool
	EdnsUdpSize                         *int64
	EnableBlacklist                     *bool
	EnableFixedRrsetOrderFqdns          *bool
	EnableMatchRecursiveOnly            *bool
	ExtAttrs                            map[string]any
	FilterAaaa                          *string
	FilterAaaaList                      []niosdns.ViewFilterAaaaList
	FixedRrsetOrderFqdns                []niosdns.ViewFixedRrsetOrderFqdns
	ForwardOnly                         *bool
	Forwarders                          []string
	LastQueriedAcl                      []niosdns.ViewLastQueriedAcl
	MatchClients                        []niosdns.ViewMatchClients
	MatchDestinations                   []niosdns.ViewMatchDestinations
	MaxCacheTtl                         *int64
	MaxNcacheTtl                        *int64
	MaxUdpSize                          *int64
	Name                                *string
	NetworkView                         *string
	NotifyDelay                         *int64
	NxdomainLogQuery                    *bool
	NxdomainRedirect                    *bool
	NxdomainRedirectAddresses           []string
	NxdomainRedirectAddressesV6         []string
	NxdomainRedirectTtl                 *int64
	NxdomainRulesets                    []string
	Recursion                           *bool
	ResponseRateLimiting                *niosdns.ViewResponseRateLimiting
	RootNameServerType                  *string
	RpzDropIpRuleEnabled                *bool
	RpzDropIpRuleMinPrefixLengthIpv4    *int64
	RpzDropIpRuleMinPrefixLengthIpv6    *int64
	RpzQnameWaitRecurse                 *bool
	ScavengingSettings                  *niosdns.ViewScavengingSettings
	Sortlist                            []niosdns.ViewSortlist
	UseBlacklist                        *bool
	UseDdnsForceCreationTimestampUpdate *bool
	UseDdnsPatternsRestriction          *bool
	UseDdnsPrincipalSecurity            *bool
	UseDdnsRestrictProtected            *bool
	UseDdnsRestrictStatic               *bool
	UseDns64                            *bool
	UseDnssec                           *bool
	UseEdnsUdpSize                      *bool
	UseFilterAaaa                       *bool
	UseFixedRrsetOrderFqdns             *bool
	UseForwarders                       *bool
	UseMaxCacheTtl                      *bool
	UseMaxNcacheTtl                     *bool
	UseMaxUdpSize                       *bool
	UseNxdomainRedirect                 *bool
	UseRecursion                        *bool
	UseResponseRateLimiting             *bool
	UseRootNameServer                   *bool
	UseRpzDropIpRule                    *bool
	UseRpzQnameWaitRecurse              *bool
	UseScavengingSettings               *bool
	UseSortlist                         *bool
}

// UDDIViewExt - UDDI specific fields for View
type UDDIViewExt struct {
	AddEdnsOptionInOutgoingQuery                *bool
	Comment                                     *string
	CompartmentId                               *string
	CustomRootNs                                []uddidnsconfig.RootNS
	CustomRootNsEnabled                         *bool
	Disabled                                    *bool
	DnssecEnableValidation                      *bool
	DnssecEnabled                               *bool
	DnssecTrustAnchors                          []uddidnsconfig.TrustAnchor
	DnssecValidateExpiry                        *bool
	DtcConfig                                   *uddidnsconfig.DTCConfig
	EcsEnabled                                  *bool
	EcsForwarding                               *bool
	EcsPrefixV4                                 *int64
	EcsPrefixV6                                 *int64
	EcsZones                                    []uddidnsconfig.ECSZone
	EdnsUdpSize                                 *int64
	FilterAaaaAcl                               []uddidnsconfig.ACLItem
	FilterAaaaOnV4                              *string
	Forwarders                                  []uddidnsconfig.Forwarder
	ForwardersOnly                              *bool
	GssTsigEnabled                              *bool
	InheritanceSources                          *uddidnsconfig.ViewInheritance
	IpSpaces                                    []string
	LameTtl                                     *int64
	MatchClientsAcl                             []uddidnsconfig.ACLItem
	MatchDestinationsAcl                        []uddidnsconfig.ACLItem
	MatchRecursiveOnly                          *bool
	MaxCacheTtl                                 *int64
	MaxNegativeTtl                              *int64
	MaxUdpSize                                  *int64
	MinimalResponses                            *bool
	Name                                        string
	Notify                                      *bool
	QueryAcl                                    []uddidnsconfig.ACLItem
	RecursionAcl                                []uddidnsconfig.ACLItem
	RecursionEnabled                            *bool
	SortList                                    []uddidnsconfig.SortListItem
	SynthesizeAddressRecordsFromHttps           *bool
	Tags                                        map[string]any
	TransferAcl                                 []uddidnsconfig.ACLItem
	UpdateAcl                                   []uddidnsconfig.ACLItem
	UseForwardersForSubzones                    *bool
	UseRootForwardersForLocalResolutionWithB1td *bool
	ZoneAuthority                               *uddidnsconfig.ZoneAuthority
}
