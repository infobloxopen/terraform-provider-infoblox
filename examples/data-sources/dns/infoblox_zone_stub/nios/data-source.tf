// Retrieve a specific DNS Stub Zone by filters
data "infoblox_zone_stub" "get_zone_stub_using_filters" {
  filters = {
    fqdn = "example_stub_zone.example.com"
  }
}

// Retrieve specific DNS Stub Zones using Extensible Attributes
data "infoblox_zone_stub" "get_zone_stub_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all DNS Stub Zones
data "infoblox_zone_stub" "get_all_zone_stubs" {}
