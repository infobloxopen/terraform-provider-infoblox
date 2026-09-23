// List specific VLANs using filters
list "infoblox_vlan" "list_vlans_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_vlan"
    }
  }
  limit = 10
}

// List specific VLANs using Extensible Attributes
list "infoblox_vlan" "list_vlans_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List VLANs with resource details included
list "infoblox_vlan" "list_vlans_with_resource" {
  provider         = infoblox
  include_resource = true
}
