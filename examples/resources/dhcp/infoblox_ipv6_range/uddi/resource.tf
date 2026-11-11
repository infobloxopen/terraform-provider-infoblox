// Create an IP Space (Required as Parent)
resource "infoblox_network_view" "example" {
  uddi = {
    name    = "example"
    comment = "Example IP space created by the terraform provider"
    tags = {
      Site = "location-1"
    }
  }
}

// Create an IPv6 Network that contains the range (Required as Parent)
resource "infoblox_ipv6_network" "example" {
  uddi = {
    address = "2001:db8:1::"
    cidr    = 64
    space   = infoblox_network_view.example.id
  }
}

// Create an IPv6 Range
resource "infoblox_ipv6_range" "example" {
  uddi = {
    start = "2001:db8:1::15"
    end   = "2001:db8:1::30"
    space = infoblox_network_view.example.id

    // Other optional fields
    name    = "example"
    comment = "Example IPv6 Range created by the terraform provider"
    tags = {
      Site = "location-1"
    }
    exclusion_ranges = [
      {
        start = "2001:db8:1::17"
        end   = "2001:db8:1::20"
      }
    ]
  }
  depends_on = [infoblox_ipv6_network.example]
}
