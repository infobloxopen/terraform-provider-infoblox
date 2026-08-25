// Retrieve a specific Network View by filters
data "infoblox_networkview" "get_network_views_using_filters" {
  filters = {
    name = "example_network_view"
  }
}

// Retrieve specific Network Views using Extensible Attributes
data "infoblox_networkview" "get_network_views_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all Network Views
data "infoblox_networkview" "get_all_network_views" {}
