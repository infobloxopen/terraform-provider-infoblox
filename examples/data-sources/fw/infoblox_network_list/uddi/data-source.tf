// Retrieve a specific Network List by name
data "infoblox_network_list" "get_network_list_using_filters" {
  filters = {
    name = "example_network_list"
  }
}

// Retrieve all Network Lists
data "infoblox_network_list" "get_all_network_lists" {}
