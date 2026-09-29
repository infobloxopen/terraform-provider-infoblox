package infra

import (
	"time"

	uddiinfraprovision "github.com/infobloxopen/universal-ddi-go-client/infraprovision"
)

// Infoblox JoinToken model
type JoinToken struct {
	Id   *string
	UDDI *UDDIJoinTokenExt
}

// UDDIJoinTokenExt - UDDI specific fields for JoinToken
type UDDIJoinTokenExt struct {
	DeletedAt   *time.Time
	Description *string
	ExpiresAt   *time.Time
	JoinToken   *string
	LastUsedAt  *time.Time
	Name        *string
	Status      *uddiinfraprovision.JoinTokenJoinTokenStatus
	Tags        map[string]any
	TokenId     *string
	UseCounter  *int64
}
