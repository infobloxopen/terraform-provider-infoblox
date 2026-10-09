// Retrieve a specific VLAN by filters
data "infoblox_vlan" "get_ipam_vlan_using_filters" {
  filters = {
    name = "example_vlan"
  }
}

// Retrieve specific VLANs using Extensible Attributes
data "infoblox_vlan" "get_ipam_vlan_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all VLANs
data "infoblox_vlan" "get_all_ipam_vlans" {}
