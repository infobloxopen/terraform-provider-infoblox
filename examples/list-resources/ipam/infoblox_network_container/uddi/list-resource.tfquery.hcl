// List specific Networkcontainers using filters
list "infoblox_network_container" "list_networkcontainer_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_address_block"
    }
  }
  limit = 10
}

// List specific Networkcontainers using Tags
list "infoblox_network_container" "list_networkcontainer_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Networkcontainers with resource details included
list "infoblox_network_container" "list_networkcontainer_with_resource" {
  provider         = infoblox
  include_resource = true
}
