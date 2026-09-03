// Get delegation zone filtered by an attribute
data "infoblox_zone_delegated" "example_by_attribute" {
  filters = {
    fqdn = "del.domain.com."
  }
}

// Get delegation zones filtered by tag
data "infoblox_zone_delegated" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Get all delegation zones
data "infoblox_zone_delegated" "example_all" {}
