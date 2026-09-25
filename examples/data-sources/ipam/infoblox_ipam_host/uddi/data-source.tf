// Retrieve IPAM Hosts filtered by an attribute
data "infoblox_ipam_host" "example_by_attribute" {
  filters = {
    address = "10.1.0.5"
  }
}

// Retrieve IPAM Hosts filtered by tag
data "infoblox_ipam_host" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all IPAM Hosts
data "infoblox_ipam_host" "example_all" {}
