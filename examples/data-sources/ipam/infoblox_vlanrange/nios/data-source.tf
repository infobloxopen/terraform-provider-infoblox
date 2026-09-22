// Retrieve a specific IPAM Vlan Range by filters
data "infoblox_vlanrange" "get_ipam_vlanrange_using_filters" {
  filters = {
    name = "example_vlan_range"
  }
}

// Retrieve specific IPAM Vlan Ranges using Extensible Attributes
data "infoblox_vlanrange" "get_ipam_vlanranges_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all IPAM Vlan Ranges
data "infoblox_vlanrange" "get_all_ipam_vlanranges" {}
