// Retrieve a specific DNS view by filters
data "infoblox_view" "get_view_using_filters" {
  filters = {
    name = "example_custom_view"
  }
}

// Retrieve specific DNS view using Extensible Attributes
data "infoblox_view" "get_view_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all DNS views
data "infoblox_view" "get_all_views" {}
