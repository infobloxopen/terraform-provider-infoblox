// Get IP Spaces filtered by an attribute
data "infoblox_networkview" "example_by_attribute" {
  filters = {
    "name" = "example_ip_space"
  }
}

// Get IP Spaces filtered by tag
data "infoblox_networkview" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Get all IP Spaces
data "infoblox_networkview" "example_all" {}
