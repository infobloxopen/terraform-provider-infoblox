// List specific VLAN Ranges using filters
list "infoblox_vlanrange" "list_vlanranges_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_vlan_range"
    }
  }
  limit = 10
}

// List specific VLAN Ranges using Extensible Attributes
list "infoblox_vlanrange" "list_vlanranges_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List VLAN Ranges with resource details included
list "infoblox_vlanrange" "list_vlanranges_with_resource" {
  provider         = infoblox
  include_resource = true
}

