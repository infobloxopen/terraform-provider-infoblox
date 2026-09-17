// Retrieve a specific IPAM Vlan by filters
data "infoblox_vlan" "get_ipam_vlan_using_filters" {
  filters = {
    name = "example_vlan"
  }
}
// Retrieve specific IPAM Vlan using Extensible Attributes
data "infoblox_vlan" "get_ipam_vlan_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all IPAM Vlans
data "infoblox_vlan" "get_all_ipam_vlans" {}
