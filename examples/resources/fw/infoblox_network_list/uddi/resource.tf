// Create a Network List with Basic Fields
resource "infoblox_network_list" "create_network_list" {
  uddi = {
    name = "example_network_list"
    addr_block = [
      {
        address = "192.0.2.0/24"
      },
    ]
  }
}

// Create a Network List with Additional Fields
resource "infoblox_network_list" "create_network_list_with_additional_fields" {
  uddi = {
    name        = "example_network_list_advanced"
    description = "Network list for protecting external networks"

    addr_block = [
      {
        address     = "198.52.100.0/24"
        description = "Primary external network block"
      },
      {
        address     = "203.0.114.0/25"
        description = "Secondary external network block"
      },
    ]
  }
}

// Import an existing Network List by its ID
import {
  to = infoblox_network_list.create_network_list_with_additional_fields
  id = "772542"
}
