// Create a Network View (Required as Parent)
resource "infoblox_network_view" "parent_network_view" {
  uddi = {
    name = "example_network_view"
  }
}

// Create an IPv6 Network (Required as Parent)
resource "infoblox_ipv6_network" "parent_network" {
  uddi = {
    address = "2001:db8:abcd:1231::"
    cidr    = 64
    space   = infoblox_network_view.parent_network_view.id
    comment = "Parent Network for the fixed addresses below"
  }
}

// Create an IPv6 Fixed Address with Basic Fields
resource "infoblox_ipv6_fixed_address" "example_fixed_address" {
  uddi = {
    name        = "example_fixed_address"
    ip_space    = infoblox_network_view.parent_network_view.id
    address     = "2001:db8:abcd:1231::10"
    match_type  = "mac"
    match_value = "00:00:00:00:00:00"
  }
  depends_on = [infoblox_ipv6_network.parent_network]
}

// Create an IPv6 Fixed Address with Additional Fields
resource "infoblox_ipv6_fixed_address" "example_fixed_address_additional" {
  uddi = {
    // Basic Fields
    name        = "example_fixed_address_additional"
    ip_space    = infoblox_network_view.parent_network_view.id
    address     = "2001:db8:abcd:1231::11"
    match_type  = "duid"
    match_value = "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"

    // Additional Fields
    comment      = "Example Fixed Address created by the terraform provider"
    hostname     = "example-host"
    disable_dhcp = false

    // Tags
    tags = {
      Site = "location-1"
    }
  }
  depends_on = [infoblox_ipv6_network.parent_network]
}
