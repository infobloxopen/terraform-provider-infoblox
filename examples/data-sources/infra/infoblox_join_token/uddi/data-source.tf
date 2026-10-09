// Retrieve Join Tokens filtered by an attribute
data "infoblox_join_token" "example_by_attribute" {
  filters = {
    "name" = "example_join_token"
  }
}

// Retrieve Join Tokens filtered by tag
data "infoblox_join_token" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all Join Tokens
data "infoblox_join_token" "example_all" {}
