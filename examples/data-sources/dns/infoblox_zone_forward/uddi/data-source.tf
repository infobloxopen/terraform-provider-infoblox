// Get forward zones filtered by an attribute
data "infoblox_zone_forward" "example_by_attribute" {
  filters = {
    fqdn = "domain.com."
  }
}

// Get forward zones filtered by tag
data "infoblox_zone_forward" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Get all forward zones
data "infoblox_zone_forward" "example_all" {}
