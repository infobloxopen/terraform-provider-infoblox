// Get DHCP fixed address filtered by an attribute
data "infoblox_ipv6fixedaddress" "example_by_attribute" {
  filters = {
    name = "example_fixed_address"
  }
}

// Get DHCP fixed address by tag
data "infoblox_ipv6fixedaddress" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Get all fixed address
data "infoblox_ipv6fixedaddress" "example_all" {}
