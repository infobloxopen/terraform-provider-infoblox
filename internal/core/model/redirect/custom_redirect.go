package redirect

// Infoblox CustomRedirect model
type CustomRedirect struct {
	Id   *int32
	UDDI *UDDICustomRedirectExt
}

// UDDICustomRedirectExt - UDDI specific fields for CustomRedirect
type UDDICustomRedirectExt struct {
	Data *string
	Name *string
}
