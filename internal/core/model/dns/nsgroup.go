package dns

import (
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
)

// Infoblox Nsgroup model
type Nsgroup struct {
	Id   *string
	NIOS *NIOSNsgroupExt
}

// NIOSNsgroupExt - NIOS specific fields for Nsgroup
type NIOSNsgroupExt struct {
	Comment             *string
	ExtAttrs            map[string]any
	ExternalPrimaries   []niosdns.NsgroupExternalPrimaries
	ExternalSecondaries []niosdns.NsgroupExternalSecondaries
	GridPrimary         []niosdns.NsgroupGridPrimary
	GridSecondaries     []niosdns.NsgroupGridSecondaries
	IsGridDefault       *bool
	IsMultimaster       *bool
	Name                *string
	UseExternalPrimary  *bool
}
