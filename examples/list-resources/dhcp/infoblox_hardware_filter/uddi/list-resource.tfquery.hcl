// List specific Hardware Filters using filters
list "infoblox_hardware_filter" "list_hardware_filter_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example-hardware-filter"
    }
  }
  limit = 10
}

// List specific Hardware Filters using tags
list "infoblox_hardware_filter" "list_hardware_filter_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      environment = "production"
    }
  }
}

// List Hardware Filters with resource details included
list "infoblox_hardware_filter" "list_hardware_filter_with_resource" {
  provider         = infoblox
  include_resource = true
}
