// Get all IPv6 Ranges
data "infoblox_ipv6_range" "example_all_ranges" {}

// Get a specific IPv6 Range by start and end values
data "infoblox_ipv6_range" "example_range_by_start_end" {
  filters = {
    start = "2001:db8:1::15",
    end   = "2001:db8:1::30"
  }
}

// Get a specific IPv6 Range by name
data "infoblox_ipv6_range" "example_range_by_name" {
  filters = {
    name = "example"
  }
}

// Get IPv6 Ranges by tag
data "infoblox_ipv6_range" "example_range_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}
