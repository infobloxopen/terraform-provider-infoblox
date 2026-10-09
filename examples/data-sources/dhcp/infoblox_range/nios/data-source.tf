// Retrieve a specific DHCP Range by filters
data "infoblox_range" "get_range_with_filter" {
  filters = {
    start_addr   = "10.0.0.10"
    network_view = "example-range-view"
  }
}

// Retrieve specific DHCP Ranges using Extensible Attributes
data "infoblox_range" "get_range_with_extattr_filter" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all DHCP Ranges
data "infoblox_range" "get_all_ranges" {}
