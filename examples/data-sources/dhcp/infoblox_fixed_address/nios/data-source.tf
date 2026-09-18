// Retrieve a specific Fixed Address by filters
data "infoblox_fixed_address" "get_fixed_address_using_filters" {
  filters = {
    ipv4addr = "16.0.0.20"
  }
}

// Retrieve specific Fixed Addresses using Extensible Attributes
data "infoblox_fixed_address" "get_fixed_address_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Search for a Fixed Address by Microsoft Server
data "infoblox_fixed_address" "get_fixed_address_using_microsoft_server" {
  body = {
    ms_server = {
      struct   = "msdhcpserver"
      ipv4addr = "1.1.1.1" // Specify the IP address of the Microsoft DHCP server
    }
  }
}

// Retrieve all Fixed Addresses
data "infoblox_fixed_address" "get_all_fixed_address" {}
