// Retrieve a specific VLAN View by filters
data "infoblox_vlan_view" "get_ipam_vlanview_using_filters" {
  filters = {
    name = "example_vlan_view"
  }
}
// Retrieve specific VLAN Views using Extensible Attributes
data "infoblox_vlan_view" "get_ipam_vlanview_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all VLAN Views
data "infoblox_vlan_view" "get_all_ipam_vlanviews" {}
