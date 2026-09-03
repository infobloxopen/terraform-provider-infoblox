// Retrieve a specific DNS Zone Stub by filters
data "infoblox_zone_stub" "get_zone_stub_using_filters" {
  filters = {
    fqdn = "example_stub_zone.example.com"
  }
}

// Retrieve specific DNS Zone Stub using Extensible Attributes
data "infoblox_zone_stub" "get_zone_stub_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all DNS Zone Stub
data "infoblox_zone_stub" "get_all_zone_stub" {}
