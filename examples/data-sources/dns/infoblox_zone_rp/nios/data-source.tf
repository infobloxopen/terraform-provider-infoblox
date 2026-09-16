// Retrieve a specific Response Policy Zone using filters
data "infoblox_zone_rp" "get_zone_rp_using_filters" {
  filters = {
    view = "default"
    fqdn = "example1.com"
  }
}

// Retrieve Response Policy Zones using Extensible Attributes
data "infoblox_zone_rp" "get_zones_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all Response Policy Zones
data "infoblox_zone_rp" "get_all_zone_rp" {}
