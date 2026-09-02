// Retrieve a specific IPv6 Fixed Address by filters
data "infoblox_ipv6fixedaddress" "example_by_attribute" {
  filters = {
    name = "example_fixed_address"
  }
}

// Retrieve specific IPv6 Fixed Address using Tags
data "infoblox_ipv6fixedaddress" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all IPv6 Fixed Addresses
data "infoblox_ipv6fixedaddress" "example_all" {}
