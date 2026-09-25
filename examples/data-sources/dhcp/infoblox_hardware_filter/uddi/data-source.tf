// Retrieve specific Hardware Filters using filters
data "infoblox_hardware_filter" "get_by_name" {
  filters = {
    name = "example-hardware-filter"
  }
}

// Retrieve specific Hardware Filters using tag filters
data "infoblox_hardware_filter" "get_by_tag" {
  tag_filters = {
    environment = "production"
  }
}

// Retrieve all Hardware Filters
data "infoblox_hardware_filter" "get_all" {}
