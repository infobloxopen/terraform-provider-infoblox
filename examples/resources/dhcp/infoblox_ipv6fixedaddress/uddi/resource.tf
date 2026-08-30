resource "infoblox_ipv6fixedaddress" "example_fixed_address" {
  uddi = {
    name        = "example_fixed_address"
    address     = "192.168.1.1"
    match_type  = "mac"
    match_value = "00:00:00:00:00:00"
    comment     = "Example Fixed Address created by the terraform provider"
    tags = {
      Site = "location-1"
    }
  }
}

// Address using Next available IP
resource "infoblox_ipv6fixedaddress" "example_fixed_address_na" {
  uddi = {
    name        = "example_fixed_address"
    match_type  = "mac"
    match_value = "00:00:00:00:00:01"
    comment     = "Example Fixed Address created by the terraform provider"
    tags = {
      Site = "location-1"
    }
  }
}
