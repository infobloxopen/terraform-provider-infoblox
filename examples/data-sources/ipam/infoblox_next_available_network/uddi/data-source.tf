data "infoblox_network_container" "example_by_attribute" {
  filters = {
    name = "example_network_container"
  }
}

data "infoblox_next_available_network" "example_next_available_network" {
  id           = data.infoblox_network_container.example_by_attribute.results.0.id
  cidr         = 29
  subnet_count = 5
}

data "infoblox_next_available_network" "example_next_available_network_by_tag" {
  cidr         = 30
  subnet_count = 15
  tag_filters = {
    environment = "prd"
  }
}
