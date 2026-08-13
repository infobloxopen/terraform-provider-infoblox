package dns

import (
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
)

// Infoblox RecordAlias model
type RecordAlias struct {
	Id   *string
	NIOS *NIOSRecordAliasExt
}

// NIOSRecordAliasExt - NIOS specific fields for RecordAlias
type NIOSRecordAliasExt struct {
	CloudInfo  *niosdns.RecordAliasCloudInfo
	Comment    *string
	Creator    *string
	Disable    *bool
	ExtAttrs   map[string]any
	Name       *string
	TargetName *string
	TargetType *string
	Ttl        *int64
	UseTtl     *bool
	View       *string
}
