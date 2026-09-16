// Get addresses filtered by an attribute
data "infoblox_address" "example_by_attribute" {
  filters = {
    address = "10.1.0.5"
  }
}

// Get addresses filtered by tag
data "infoblox_address" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Get all addresses
data "infoblox_address" "example_all" {}
