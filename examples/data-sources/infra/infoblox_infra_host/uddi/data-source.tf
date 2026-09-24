// Retrieve Infra Hosts filtered by an attribute
data "infoblox_infra_host" "example_by_attribute" {
  filters = {
    "display_name" = "example_host"
  }
}

// Retrieve Infra Hosts filtered by tag
data "infoblox_infra_host" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all Infra Hosts
data "infoblox_infra_host" "example_all" {}
