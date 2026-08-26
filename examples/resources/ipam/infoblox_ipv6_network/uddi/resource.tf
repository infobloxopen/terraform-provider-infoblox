// Create an IPAM IPv6 Network with Basic Fields
resource "infoblox_ipv6_network" "exampl_network_basic" {
  uddi = {
    address = "2001:db8:1ef8:e4ee::"
    cidr    = 64
    space   = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"

    // Other optional fields
    name    = "example_ipv6_network"
    comment = "Network for Site A"
    tags = {
      Site = "location-1"
    }
  }
}

resource "infoblox_ipv6_network" "example_network_additional" {
  uddi = {
    address = "2002:db8:1ef8:e4ee::"
    cidr    = 64
    space   = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"

    // Other optional fields
    name    = "example_ipv6_network_additional"
    comment = "Network for Site B"

    disable_dhcp = false
    renew_time   = 1800
    rebind_time  = 2700

    ddns_send_updates     = true
    ddns_domain           = "example.com"
    ddns_generate_name    = true
    ddns_generated_prefix = "myhost"

    tags = {
      Site = "location-2"
    }
  }
}
