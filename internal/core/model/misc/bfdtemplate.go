package misc

// Infoblox Bfdtemplate model
type Bfdtemplate struct {
	Id   *string
	NIOS *NIOSBfdtemplateExt
}

// NIOSBfdtemplateExt - NIOS specific fields for Bfdtemplate
type NIOSBfdtemplateExt struct {
	AuthenticationKey   *string
	AuthenticationKeyId *int64
	AuthenticationType  *string
	DetectionMultiplier *int64
	MinRxInterval       *int64
	MinTxInterval       *int64
	Name                *string
}
