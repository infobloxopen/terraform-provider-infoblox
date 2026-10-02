package discovery

// Infoblox DiscoveryCredentialgroup model
type DiscoveryCredentialgroup struct {
	Id   *string
	NIOS *NIOSDiscoveryCredentialgroupExt
}

// NIOSDiscoveryCredentialgroupExt - NIOS specific fields for DiscoveryCredentialgroup
type NIOSDiscoveryCredentialgroupExt struct {
	Name *string
}
