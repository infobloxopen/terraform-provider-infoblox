package dns

import (
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
)

// Infoblox ZoneRp model
type ZoneRp struct {
	Id   *string
	NIOS *NIOSZoneRpExt
}

// NIOSZoneRpExt - NIOS specific fields for ZoneRp
type NIOSZoneRpExt struct {
	Comment                          *string
	Disable                          *bool
	ExtAttrs                         map[string]any
	ExternalPrimaries                []niosdns.ZoneRpExternalPrimaries
	ExternalSecondaries              []niosdns.ZoneRpExternalSecondaries
	FireeyeRuleMapping               *niosdns.ZoneRpFireeyeRuleMapping
	Fqdn                             *string
	GridPrimary                      []niosdns.ZoneRpGridPrimary
	GridSecondaries                  []niosdns.ZoneRpGridSecondaries
	Locked                           *bool
	LogRpz                           *bool
	MemberSoaMnames                  []niosdns.ZoneRpMemberSoaMnames
	NsGroup                          *string
	Prefix                           *string
	RecordNamePolicy                 *string
	RpzDropIpRuleEnabled             *bool
	RpzDropIpRuleMinPrefixLengthIpv4 *int64
	RpzDropIpRuleMinPrefixLengthIpv6 *int64
	RpzPolicy                        *string
	RpzSeverity                      *string
	RpzType                          *string
	SetSoaSerialNumber               *bool
	SoaDefaultTtl                    *int64
	SoaEmail                         *string
	SoaExpire                        *int64
	SoaNegativeTtl                   *int64
	SoaRefresh                       *int64
	SoaRetry                         *int64
	SoaSerialNumber                  *int64
	SubstituteName                   *string
	UseExternalPrimary               *bool
	UseGridZoneTimer                 *bool
	UseLogRpz                        *bool
	UseRecordNamePolicy              *bool
	UseRpzDropIpRule                 *bool
	UseSoaEmail                      *bool
	View                             *string
}
