// Retrieve a specific VLAN Range by filters
data "infoblox_vlan_range" "get_ipam_vlanrange_using_filters" {
  filters = {
    name = "example_vlan_range"
  }
}

// Retrieve specific VLAN Ranges using Extensible Attributes
data "infoblox_vlan_range" "get_ipam_vlanranges_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all VLAN Ranges
data "infoblox_vlan_range" "get_all_ipam_vlanranges" {}
