// Get subnets filtered by an attribute
data "infoblox_network" "example_by_attribute" {
  filters = {
    "address" = "10.0.0.0"
    "cidr"    = "24"
  }
}

// Get subnets filtered by tag
data "infoblox_network" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Get all subnets
data "infoblox_network" "example_all" {}
