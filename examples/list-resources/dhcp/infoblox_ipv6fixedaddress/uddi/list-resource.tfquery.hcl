// List specific Ipv6Fixedaddresses using filters
list "infoblox_ipv6fixedaddress" "list_ipv6fixedaddress_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_fixed_address"
    }
  }
  limit = 10
}

// List specific Ipv6Fixedaddresses using Tags
list "infoblox_ipv6fixedaddress" "list_ipv6fixedaddress_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Ipv6Fixedaddresses with resource details included
list "infoblox_ipv6fixedaddress" "list_ipv6fixedaddress_with_resource" {
  provider         = infoblox
  include_resource = true
}
