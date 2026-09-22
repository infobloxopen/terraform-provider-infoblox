// List specific VLAN Views using filters
list "infoblox_vlan_view" "list_vlanviews_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_vlan_view"
    }
  }
  limit = 10
}

// List specific VLAN Views using Extensible Attributes
list "infoblox_vlan_view" "list_vlanviews_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List VLAN Views with resource details included
list "infoblox_vlan_view" "list_vlanviews_with_resource" {
  provider         = infoblox
  include_resource = true
}
