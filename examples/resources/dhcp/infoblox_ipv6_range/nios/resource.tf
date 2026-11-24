// Create an IPv6 Network that contains the ranges (Required as Parent)
resource "infoblox_ipv6_network" "example" {
  nios = {
    network = "15::/64"
  }
}

// Create DHCP IPv6 Range with Address Type set to "ADDRESS" (default)
resource "infoblox_ipv6_range" "create_ipv6range_with_default_address_type" {
  nios = {
    network      = infoblox_ipv6_network.example.nios.network
    start_addr   = "15::10"
    end_addr     = "15::20"
    address_type = "ADDRESS"
  }
}

// Create DHCP IPv6 Range with Address Type set to "PREFIX"
resource "infoblox_ipv6_range" "create_ipv6range_with_address_type_prefix" {
  nios = {
    network           = infoblox_ipv6_network.example.nios.network
    ipv6_start_prefix = "15:0:0:1000::"
    ipv6_end_prefix   = "15:0:0:1fff::"
    ipv6_prefix_bits  = 80
    address_type      = "PREFIX"
  }
}

// Create DHCP IPv6 Range with Address Type set to "BOTH"
resource "infoblox_ipv6_range" "create_ipv6range_with_address_type_both" {
  nios = {
    network           = infoblox_ipv6_network.example.nios.network
    start_addr        = "15::25"
    end_addr          = "15::35"
    ipv6_start_prefix = "15:0:0:2000::"
    ipv6_end_prefix   = "15:0:0:2fff::"
    ipv6_prefix_bits  = 80
    address_type      = "BOTH"
  }
}

// Create DHCP IPv6 Range with Additional Fields
resource "infoblox_ipv6_range" "create_ipv6range_with_additional_fields" {
  nios = {
    network    = infoblox_ipv6_network.example.nios.network
    start_addr = "15::40"
    end_addr   = "15::50"

    // Additional Fields
    comment = "DHCP IPv6 Range created by Terraform"
    ext_attrs = {
      Site = "location-1"
    }
    exclude = [{
      start_address = "15::40",
      end_address   = "15::45",
      comment       = "Exclude range 15::40 - 15::45",
    }]
  }
}
