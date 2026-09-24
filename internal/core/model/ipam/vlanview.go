package ipam

// Infoblox Vlanview model
type Vlanview struct {
	Id   *string
	NIOS *NIOSVlanviewExt
}

// NIOSVlanviewExt - NIOS specific fields for Vlanview
type NIOSVlanviewExt struct {
	AllowRangeOverlapping *bool
	Comment               *string
	EndVlanId             *int64
	ExtAttrs              map[string]any
	Name                  *string
	PreCreateVlan         *bool
	StartVlanId           *int64
	VlanNamePrefix        *string
}
