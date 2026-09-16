// List specific Network Views using filters
list "infoblox_network_view" "list_network_views_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_network_view"
    }
  }
  limit = 10
}

// List specific Network Views using Tags
list "infoblox_network_view" "list_network_views_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Network Views with resource details included
list "infoblox_network_view" "list_network_views_with_resource" {
  provider         = infoblox
  include_resource = true
}
