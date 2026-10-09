// Retrieve specific Auth Zones using filters
data "infoblox_zone_auth" "get_zone_using_filters" {
  filters = {
    fqdn = "example.com."
  }
}

// Retrieve specific Auth Zones using Tags
data "infoblox_zone_auth" "get_zone_using_tag_filters" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all Auth Zones
data "infoblox_zone_auth" "get_all_zones" {}
