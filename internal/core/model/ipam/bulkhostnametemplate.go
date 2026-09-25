package ipam

// Infoblox Bulkhostnametemplate model
type Bulkhostnametemplate struct {
	Id   *string
	NIOS *NIOSBulkhostnametemplateExt
}

// NIOSBulkhostnametemplateExt - NIOS specific fields for Bulkhostnametemplate
type NIOSBulkhostnametemplateExt struct {
	TemplateFormat *string
	TemplateName   *string
}
