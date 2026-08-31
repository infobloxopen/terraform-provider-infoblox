// Get Auth NSGs filtered by an attribute
data "infoblox_auth_nsg" "example_by_attribute" {
  filters = {
    "name" = "example_auth_nsg"
  }
}

// Get Auth NSGs filtered by tag
data "infoblox_auth_nsg" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Get all Auth NSGs
data "infoblox_auth_nsg" "example_all" {}
