// Retrieve specific Category Filters by filters
data "infoblox_category_filter" "get_category_filter_using_filters" {
  filters = {
    name = "example_category_filter"
  }
}

// Retrieve specific Category Filters using Tags
data "infoblox_category_filter" "get_category_filter_using_tag_filters" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all Category Filters
data "infoblox_category_filter" "get_all_category_filters" {}
