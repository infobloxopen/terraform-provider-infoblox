// Retrieve a specific IPAM Vlan View by filters
data "infoblox_vlanview" "get_ipam_vlanview_using_filters" {
  filters = {
    name = "example_vlan_view"
  }
}
// Retrieve specific IPAM Vlan Views using Extensible Attributes
data "infoblox_vlanview" "get_ipam_vlanview_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all IPAM Vlan Views
data "infoblox_vlanview" "get_all_ipam_vlanview" {}
