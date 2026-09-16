// Retrieve a specific Network View by filters
data "infoblox_network_view" "get_network_views_using_filters" {
  filters = {
    "name" = "example_network_view"
  }
}

// Retrieve specific Network Views using Tags
data "infoblox_network_view" "get_network_views_using_tags" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all Network Views
data "infoblox_network_view" "get_all_network_views" {}
