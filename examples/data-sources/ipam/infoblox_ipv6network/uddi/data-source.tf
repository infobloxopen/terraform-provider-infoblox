// Get subnets filtered by an attribute
data "infoblox_ipv6network" "example_by_attribute" {
  filters = {
    "address" = "2001:db8:1ef8:e4ee::"
    "cidr"    = "64"
  }
}

// Get subnets filtered by tag
data "infoblox_ipv6network" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Get all subnets
data "infoblox_ipv6network" "example_all" {}
