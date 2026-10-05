// List specific Category Filters using filters
list "infoblox_category_filter" "list_category_filter_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_category_filter"
    }
  }
  limit = 10
}

// List specific Category Filters using Tags
list "infoblox_category_filter" "list_category_filter_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Category Filters with resource details included
list "infoblox_category_filter" "list_category_filter_with_resource" {
  provider         = infoblox
  include_resource = true
}
