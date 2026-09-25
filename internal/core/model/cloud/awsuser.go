package cloud

// Infoblox Awsuser model
type Awsuser struct {
	Id   *string
	NIOS *NIOSAwsuserExt
}

// NIOSAwsuserExt - NIOS specific fields for Awsuser
type NIOSAwsuserExt struct {
	AccessKeyId     *string
	AccountId       *string
	GovcloudEnabled *bool
	Name            *string
	NiosUserName    *string
	SecretAccessKey *string
}
