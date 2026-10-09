// List specific Auth Zones using filters
list "infoblox_zone_auth" "list_zones_using_filters" {
  provider = infoblox
  config {
    filters = {
      view = "default"
    }
  }
  limit = 10
}

// List specific Auth Zones using Extensible Attributes
list "infoblox_zone_auth" "list_zones_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List Auth Zones with resource details included
list "infoblox_zone_auth" "list_zones_with_resource" {
  provider         = infoblox
  include_resource = true
}
