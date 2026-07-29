// List specific Auth Zones using filters
list "infoblox_zone_auth" "list_zones_using_filters" {
  provider = infoblox
  config {
    filters = {
      primary_type = "cloud"
    }
  }
  limit = 10
}

// List specific Auth Zones using Tags
list "infoblox_zone_auth" "list_zones_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Auth Zones with resource details included
list "infoblox_zone_auth" "list_zones_with_resource" {
  provider         = infoblox
  include_resource = true
}
