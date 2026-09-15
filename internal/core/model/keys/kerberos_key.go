package keys

// Infoblox KerberosKey model
type KerberosKey struct {
	Id   *string
	UDDI *UDDIKerberosKeyExt
}

// UDDIKerberosKeyExt - UDDI specific fields for KerberosKey
type UDDIKerberosKeyExt struct {
	Algorithm  *string
	Comment    *string
	Domain     *string
	Principal  *string
	Tags       map[string]any
	UploadedAt *string
	Version    *int64
}
