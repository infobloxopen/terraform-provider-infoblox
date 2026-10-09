// Create a basic DHCP Hardware Filter matching specific MAC addresses
resource "infoblox_hardware_filter" "example" {
  uddi = {
    name      = "example-hardware-filter"
    addresses = ["12:34:56:78:9a:bc"]
  }
}

// Create a DHCP Option Space (Required as Parent of the Option Definition)
resource "infoblox_dhcp_optionspace" "example" {
  uddi = {
    name = "example_option_space"
  }
}

// Create a DHCP Option Definition (Required for dhcp_options)
resource "infoblox_dhcp_optiondefinition" "example" {
  uddi = {
    code         = 150
    name         = "example_option_definition"
    type         = "address4"
    option_space = infoblox_dhcp_optionspace.example.id
  }
}

// Create a DHCP Hardware Filter with all optional fields
resource "infoblox_hardware_filter" "example_with_options" {
  uddi = {
    name    = "example-hardware-filter-full"
    comment = "Hardware filter for specific MAC addresses"
    addresses = [
      "12:34:56:78:9a:bc",
      "ab:cd:ef:12:34:56"
    ]

    # DHCP options handed out to clients matching this filter
    dhcp_options = [
      {
        type         = "option"
        option_code  = infoblox_dhcp_optiondefinition.example.id
        option_value = "192.168.10.1"
      }
    ]

    # PXE boot options
    header_option_filename       = "pxelinux.0"
    header_option_server_address = "192.168.1.10"
    header_option_server_name    = "tf-infoblox.example.com."

    lease_time = 3600
    role       = "values"

    tags = {
      Site = "location-1"
    }
  }
}
