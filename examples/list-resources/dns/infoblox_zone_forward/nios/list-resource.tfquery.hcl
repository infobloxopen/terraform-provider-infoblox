// List specific Forward Zones using filters
list "infoblox_zone_forward" "list_forward_zones_using_filters" {
  provider = infoblox
  config {
    filters = {
      fqdn = "example.com"
    }
  }
  limit = 10
}

// List specific Forward Zones using Extensible Attributes
list "infoblox_zone_forward" "list_forward_zones_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List Forward Zones with resource details included
list "infoblox_zone_forward" "list_forward_zones_with_resource" {
  provider         = infoblox
  include_resource = true
}
