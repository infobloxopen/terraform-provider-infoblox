// Retrieve a specific network container using filters
data "infoblox_network_container" "example_by_attribute" {
  filters = {
    "name" = "example_address_block"
  }
}

// Retrieve a specific network container using tags
data "infoblox_network_container" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all network containers
data "infoblox_network_container" "example_all" {}
