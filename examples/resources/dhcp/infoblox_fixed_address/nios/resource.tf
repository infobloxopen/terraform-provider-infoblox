// Create an IPV4 Network (Required as Parent)
resource "infoblox_network" "parent_network" {
  nios = {
    network      = "16.0.0.0/24"
    network_view = "default"
    comment      = "Parent network for DHCP fixed addresses"
  }
}

// Create Fixed Address with Basic Fields
resource "infoblox_fixed_address" "create_fixed_address_basic" {
  nios = {
    ipv4addr     = "16.0.0.10"
    match_client = "MAC_ADDRESS"
    mac          = "00:1a:2b:3c:4d:5e"

    // Extensible Attributes
    ext_attrs = {
      Site = "location-1"
    }
  }
  depends_on = [infoblox_network.parent_network]
}

// Create Fixed Address with Additional Fields
resource "infoblox_fixed_address" "create_fixed_address_additional" {
  nios = {
    // Basic Fields
    ipv4addr     = "16.0.0.20"
    match_client = "MAC_ADDRESS"
    mac          = "00:6a:7b:8c:9d:5e"

    // Additional Fields
    comment = "Fixed Address created with additional fields"

    bootfile = "pxelinux.0"

    enable_ddns = true

    pxe_lease_time = 3600

    device_location = "APJ"
    device_type     = "Server"

    options = [
      {
        name  = "time-offset"
        num   = 2
        value = "50"
      },
      {
        name  = "dhcp-lease-time"
        num   = 51
        value = "7200"
      },
      {
        name  = "domain-name-servers"
        num   = 6
        value = "8.8.8.8,8.8.4.4"
      }
    ]
    // Extensible Attributes
    ext_attrs = {
      Site = "location-1"
    }
  }
  depends_on = [infoblox_network.parent_network]
}

// Create Fixed Address using dynamic allocation to retrieve ipv4addr
resource "infoblox_fixed_address" "create_fixed_address_with_func_call" {
  nios = {
    match_client     = "CIRCUIT_ID"
    agent_circuit_id = 250
    dynamic_allocation = {
      network      = "16.0.0.0/24"
      network_view = "default"
    }
    comment = "Fixed Address created with ipv4addr retrieved via function call"
  }
  depends_on = [infoblox_network.parent_network]
}
