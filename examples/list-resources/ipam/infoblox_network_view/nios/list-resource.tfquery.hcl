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

// List specific Network Views using Extensible Attributes
list "infoblox_network_view" "list_network_views_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List Network Views with resource details included
list "infoblox_network_view" "list_network_views_with_resource" {
  provider         = infoblox
  include_resource = true
}
