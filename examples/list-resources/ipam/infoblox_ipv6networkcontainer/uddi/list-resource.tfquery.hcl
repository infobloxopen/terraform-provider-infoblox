// List specific Ipv6 Network Containers using filters
list "infoblox_ipv6networkcontainer" "list_ipv6networkcontainer_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_ipv6_network_container"
    }
  }
  limit = 10
}

// List specific Ipv6 Network Containers using Tags
list "infoblox_ipv6networkcontainer" "list_ipv6networkcontainer_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Ipv6 Network Containers with resource details included
list "infoblox_ipv6networkcontainer" "list_ipv6networkcontainer_with_resource" {
  provider         = infoblox
  include_resource = true
}
