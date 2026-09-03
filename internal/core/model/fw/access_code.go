package fw

import (
	"time"

	uddifw "github.com/infobloxopen/universal-ddi-go-client/fw"
)

// Infoblox AccessCode model
type AccessCode struct {
	Id   *string
	UDDI *UDDIAccessCodeExt
}

// UDDIAccessCodeExt - UDDI specific fields for AccessCode
type UDDIAccessCodeExt struct {
	AccessKey   *string
	Activation  *time.Time
	Description *string
	Expiration  *time.Time
	Name        *string
	PolicyIds   []int32
	Rules       []uddifw.AccessCodeRule
}
