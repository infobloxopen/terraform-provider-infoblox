// Retrieve specific Auth Zones using filters
data "infoblox_zone_auth" "get_zone_using_filters" {
  filters = {
    fqdn = "example.com"
    view = "default"
  }
}

// Retrieve specific Auth Zones using Extensible Attributes
data "infoblox_zone_auth" "get_zone_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all Auth Zones
data "infoblox_zone_auth" "get_all_zones" {}
