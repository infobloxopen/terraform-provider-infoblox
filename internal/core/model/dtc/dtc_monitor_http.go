package dtc

import (
	uddidtc "github.com/infobloxopen/universal-ddi-go-client/dtc"
)

// Infoblox DtcMonitorHttp model
type DtcMonitorHttp struct {
	Id   *string
	NIOS *NIOSDtcMonitorHttpExt
	UDDI *UDDIDtcMonitorHttpExt
}

// NIOSDtcMonitorHttpExt - NIOS specific fields for DtcMonitorHttp
type NIOSDtcMonitorHttpExt struct {
	Ciphers             *string
	ClientCert          *string
	Comment             *string
	ContentCheck        *string
	ContentCheckInput   *string
	ContentCheckOp      *string
	ContentCheckRegex   *string
	ContentExtractGroup *int64
	ContentExtractType  *string
	ContentExtractValue *string
	EnableSni           *bool
	ExtAttrs            map[string]any
	Interval            *int64
	Name                *string
	Port                *int64
	Request             *string
	Result              *string
	ResultCode          *int64
	RetryDown           *int64
	RetryUp             *int64
	Secure              *bool
	Timeout             *int64
	ValidateCert        *bool
}

// UDDIDtcMonitorHttpExt - UDDI specific fields for DtcMonitorHttp
type UDDIDtcMonitorHttpExt struct {
	CheckResponseBody           *bool
	CheckResponseBodyNegative   *bool
	CheckResponseBodyRegex      *string
	CheckResponseHeader         *bool
	CheckResponseHeaderNegative *bool
	CheckResponseHeaderRegexes  []uddidtc.HeaderRegex
	Codes                       *string
	Comment                     *string
	Disabled                    *bool
	Https                       *bool
	Interval                    *int64
	Metadata                    *uddidtc.Metadata
	Name                        string
	Port                        int64
	Request                     *string
	RetryDown                   *int64
	RetryUp                     *int64
	Tags                        map[string]any
	Timeout                     *int64
}
