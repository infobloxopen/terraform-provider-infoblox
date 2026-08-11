// List specific Ipv6Networkcontainers using filters
list "infoblox_ipv6networkcontainer" "list_ipv6networkcontainer_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_subnet"
    }
  }
  limit = 10
}

// List specific Ipv6Networkcontainers using Tags
list "infoblox_ipv6networkcontainer" "list_ipv6networkcontainer_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Ipv6Networkcontainers with resource details included
list "infoblox_ipv6networkcontainer" "list_ipv6networkcontainer_with_resource" {
  provider         = infoblox
  include_resource = true
}
