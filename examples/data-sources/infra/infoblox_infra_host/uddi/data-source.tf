// Get hosts filtered by an attribute
data "infoblox_infra_host" "example_by_attribute" {
  filters = {
    "name" = "example_host"
  }
}

// Get hosts filtered by tag
data "infoblox_infra_host" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Get all hosts
data "infoblox_infra_host" "example_all" {}
