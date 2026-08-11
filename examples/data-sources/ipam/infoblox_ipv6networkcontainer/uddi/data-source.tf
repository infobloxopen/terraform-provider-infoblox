// Get Address blocks filtered by an attribute
data "infoblox_ipv6networkcontainer" "example_by_attribute" {
  filters = {
    "name" = "example_subnet"
  }
}

// Get Address blocks filtered by tag
data "infoblox_ipv6networkcontainer" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Get all Address blocks
data "infoblox_ipv6networkcontainer" "example_all" {}
