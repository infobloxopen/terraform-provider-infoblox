package dns

import (
	uddidnsconfig "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// Infoblox DnsServer model
type DnsServer struct {
	Id   *string
	UDDI *UDDIDnsServerExt
}

// UDDIDnsServerExt - UDDI specific fields for DnsServer
type UDDIDnsServerExt struct {
	AddEdnsOptionInOutgoingQuery                *bool
	AutoSortViews                               *bool
	Comment                                     *string
	CustomRootNs                                []uddidnsconfig.RootNS
	CustomRootNsEnabled                         *bool
	DnssecEnableValidation                      *bool
	DnssecEnabled                               *bool
	DnssecTrustAnchors                          []uddidnsconfig.TrustAnchor
	DnssecValidateExpiry                        *bool
	EcsEnabled                                  *bool
	EcsForwarding                               *bool
	EcsPrefixV4                                 *int64
	EcsPrefixV6                                 *int64
	EcsZones                                    []uddidnsconfig.ECSZone
	FilterAaaaAcl                               []uddidnsconfig.ACLItem
	FilterAaaaOnV4                              *string
	Forwarders                                  []uddidnsconfig.Forwarder
	ForwardersOnly                              *bool
	GssTsigEnabled                              *bool
	InheritanceSources                          *uddidnsconfig.ServerInheritance
	KerberosKeys                                []uddidnsconfig.KerberosKey
	LameTtl                                     *int64
	LogQueryResponse                            *bool
	MatchRecursiveOnly                          *bool
	MaxCacheTtl                                 *int64
	MaxNegativeTtl                              *int64
	MinimalResponses                            *bool
	Name                                        string
	Notify                                      *bool
	QueryAcl                                    []uddidnsconfig.ACLItem
	QueryPort                                   *int64
	RecursionAcl                                []uddidnsconfig.ACLItem
	RecursionEnabled                            *bool
	RecursiveClients                            *int64
	ResolverQueryTimeout                        *int64
	SecondaryAxfrQueryLimit                     *int64
	SecondarySoaQueryLimit                      *int64
	SortList                                    []uddidnsconfig.SortListItem
	SynthesizeAddressRecordsFromHttps           *bool
	Tags                                        map[string]any
	TransferAcl                                 []uddidnsconfig.ACLItem
	UpdateAcl                                   []uddidnsconfig.ACLItem
	UseForwardersForSubzones                    *bool
	UseRootForwardersForLocalResolutionWithB1td *bool
	Views                                       []uddidnsconfig.DisplayView
}
