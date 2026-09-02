package grid

// Infoblox Natgroup model
type Natgroup struct {
	Id   *string
	NIOS *NIOSNatgroupExt
}

// NIOSNatgroupExt - NIOS specific fields for Natgroup
type NIOSNatgroupExt struct {
	Comment *string
	Name    *string
}
