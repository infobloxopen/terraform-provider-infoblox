// Retrieve a Fixed Address filtered by an attribute
data "infoblox_fixed_address" "example_by_attribute" {
  filters = {
    name = "example_fixed_address"
  }
}

// Retrieve Fixed Addresses by tag
data "infoblox_fixed_address" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all Fixed Addresses
data "infoblox_fixed_address" "example_all" {}
