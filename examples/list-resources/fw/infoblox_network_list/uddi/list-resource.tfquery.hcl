// List specific Network Lists using filters
list "infoblox_network_list" "list_network_list_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_network_list"
    }
  }
  limit = 10
}

// List Network Lists with resource details included
list "infoblox_network_list" "list_network_list_with_resource" {
  provider         = infoblox
  include_resource = true
}
