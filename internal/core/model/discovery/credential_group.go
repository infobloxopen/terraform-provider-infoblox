package discovery

// Infoblox CredentialGroup model
type CredentialGroup struct {
	Id   *string
	NIOS *NIOSCredentialGroupExt
}

// NIOSCredentialGroupExt - NIOS specific fields for CredentialGroup
type NIOSCredentialGroupExt struct {
	Name *string
}
