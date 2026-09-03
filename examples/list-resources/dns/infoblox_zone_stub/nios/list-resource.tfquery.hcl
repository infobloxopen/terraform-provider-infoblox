// List specific Stub Zones using filters
list "infoblox_zone_stub" "list_stub_zones_using_filters" {
  provider = infoblox
  config {
    filters = {
      fqdn = "example_stub_zone.example.com"
    }
  }
  limit = 10
}

// List specific Stub Zones using Extensible Attributes
list "infoblox_zone_stub" "list_stub_zones_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List Stub Zones with resource details included
list "infoblox_zone_stub" "list_stub_zones_with_resource" {
  provider         = infoblox
  include_resource = true
}
