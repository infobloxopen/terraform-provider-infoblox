// Get join tokens filtered by an attribute
data "infoblox_join_token" "example_by_attribute" {
  filters = {
    "name" = "example_join_token"
  }
}

// Get join tokens filtered by tag
data "infoblox_join_token" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Get all join tokens
data "infoblox_join_token" "example_all" {}
