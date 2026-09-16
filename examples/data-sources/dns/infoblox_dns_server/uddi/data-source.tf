// Get DNS Servers filtered by an attribute
data "infoblox_dns_server" "example_by_attribute" {
  filters = {
    "name" = "example_server"
  }
}

// Get DNS Servers filtered by tag
data "infoblox_dns_server" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Get all DNS Servers
data "infoblox_dns_server" "example_all" {
}
