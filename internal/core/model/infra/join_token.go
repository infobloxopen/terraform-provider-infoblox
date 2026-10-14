package infra

import (
	"time"
)

// Infoblox JoinToken model
type JoinToken struct {
	Id   *string
	UDDI *UDDIJoinTokenExt
}

// UDDIJoinTokenExt - UDDI specific fields for JoinToken
type UDDIJoinTokenExt struct {
	Description *string
	ExpiresAt   *time.Time
	JoinToken   *string
	Name        *string
	Tags        map[string]any
	TokenId     *string
}
