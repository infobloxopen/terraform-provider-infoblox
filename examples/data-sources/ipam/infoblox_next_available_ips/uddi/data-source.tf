data "infoblox_network_container" "example_by_attribute" {
  filters = {
    name = "example_network_container"
  }
}

data "infoblox_network" "example_by_attribute" {
  filters = {
    name = "example_network"
  }
}

data "infoblox_range" "example_by_attribute" {
  filters = {
    name = "example_range"
  }
}

data "infoblox_next_available_ips" "example_next_ip_nc" {
  id       = data.infoblox_network_container.example_by_attribute.results.0.id
  ip_count = 5
}

data "infoblox_next_available_ips" "example_next_ip_network" {
  id       = data.infoblox_network.example_by_attribute.results.0.id
  ip_count = 5
}

data "infoblox_next_available_ips" "example_next_ip_range" {
  id       = data.infoblox_range.example_by_attribute.results.0.id
  ip_count = 5
}

data "infoblox_next_available_ips" "example_next_ip_nc_by_tag" {
  tag_filters = {
    name = "example_network_container"
  }
  ip_count      = 5
  resource_type = "address_block"
}

data "infoblox_next_available_ips" "example_next_ip_network_by_tag" {
  tag_filters = {
    name = "example_network"
  }
  ip_count      = 5
  resource_type = "subnet"
}

data "infoblox_next_available_ips" "example_next_ip_range_by_tag" {
  tag_filters = {
    name = "example_range"
  }
  ip_count      = 5
  resource_type = "range"
}
