data "infoblox_network_container" "example_by_attribute" {
  filters = {
    name = "example_network_container"
  }
}

data "infoblox_next_available_network_container" "example_next_available_nc" {
  id                  = data.infoblox_network_container.example_by_attribute.results.0.id
  address_block_count = 5
  cidr                = 27
}

data "infoblox_next_available_network_container" "example_next_available_nc_default_count" {
  id   = data.infoblox_network_container.example_by_attribute.results.0.id
  cidr = 24
}

data "infoblox_next_available_network_container" "example_next_available_nc_by_tag" {
  cidr                = 30
  address_block_count = 15
  tag_filters = {
    environment = "prd"
  }
}
