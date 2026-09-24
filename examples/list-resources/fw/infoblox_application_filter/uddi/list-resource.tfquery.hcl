// List specific Application Filters using filters
list "infoblox_application_filter" "list_application_filter_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example-app-filter-by-name"
    }
  }
  limit = 10
}

// List specific Application Filters using Tags
list "infoblox_application_filter" "list_application_filter_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Application Filters with resource details included
list "infoblox_application_filter" "list_application_filter_with_resource" {
  provider         = infoblox
  include_resource = true
}
