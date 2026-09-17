data "infoblox_application_filter" "get_by_name" {
  filters = { name = "example-app-filter-by-name" }
}

data "infoblox_application_filter" "get_by_tag" {
  tag_filters = { Site = "location-1" }
}

data "infoblox_application_filter" "get_all" {}
