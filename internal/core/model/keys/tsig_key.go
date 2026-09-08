package keys

// Infoblox TsigKey model
type TsigKey struct {
	Id   *string
	UDDI *UDDITsigKeyExt
}

// UDDITsigKeyExt - UDDI specific fields for TsigKey
type UDDITsigKeyExt struct {
	Algorithm *string
	Comment   *string
	Name      string
	Secret    string
	Tags      map[string]any
}
