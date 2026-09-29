// Retrieve a specific Application Filter by name
data "infoblox_application_filter" "get_by_name" {
  filters = { name = "example-app-filter-by-name" }
}

// Retrieve Application Filters matching a tag value
data "infoblox_application_filter" "get_by_tag" {
  tag_filters = { Site = "location-1" }
}

// Retrieve all Application Filters
data "infoblox_application_filter" "get_all" {}
