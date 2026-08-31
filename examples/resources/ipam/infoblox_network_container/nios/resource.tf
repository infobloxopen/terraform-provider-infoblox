// Create an IPV4 Network Container with Basic Fields
resource "infoblox_network_container" "networkcontainer_with_basic_fields" {
  nios = {
    network      = "10.0.0.0/24"
    network_view = "default"
    comment      = "Created by Terraform"
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create an IPV4 Network Container with Additional Fields
resource "infoblox_network_container" "networkcontainer_with_additional_fields" {
  nios = {
    network = "11.0.0.0/24"
    comment = "Complete network container example with all possible writable attributes"

    // BOOTP/PXE settings 
    bootfile   = "pxelinux.0"
    bootserver = "192.168.1.10"

    // DDNS settings
    enable_ddns                 = true
    ddns_domainname             = "example.com"
    ddns_generate_hostname      = true
    ddns_ttl                    = 3600
    ddns_update_fixed_addresses = true
    ddns_use_option81           = true

    // Email and notification settings
    email_list = ["admin@example.com", "network@example.com"]

    // Water mark settings
    high_water_mark       = 95
    high_water_mark_reset = 85
    low_water_mark        = 10
    low_water_mark_reset  = 20

    // Extensible attributes
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create an IPV4 Network Container with Dynamic Allocation
resource "infoblox_network_container" "example_func_call" {
  nios = {
    dynamic_allocation = {
      network = "88.175.0.0/21"
      cidr    = 24
    }
  }
}
