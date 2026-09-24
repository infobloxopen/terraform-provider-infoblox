// Manage Network View ( Required as Parent )
resource "infoblox_network_view" "parent_view" {
  uddi = {
    name = "example_network_view"
  }
}

// Manage IPv4 Network ( Required as Parent )
resource "infoblox_network" "parent_network" {
  uddi = {
    address = "10.0.0.0"
    cidr    = 24
    space   = infoblox_network_view.parent_view.id
  }
}

// Manage IPAM Hosts with Basic Fields
resource "infoblox_ipam_host" "ipam_host" {
  uddi = {
    name = "example_ipam_host"
  }
}

// Manage IPAM Hosts with Additional Fields
resource "infoblox_ipam_host" "ipam_host_with_additional_fields" {
  uddi = {
    name      = "example_ipam_host_full"
    comment   = "IPAM Hosts example"
    addresses = [{ address = "10.0.0.1", space = infoblox_network_view.parent_view.id }]

    tags = {
      Site = "location-1"
    }
  }
}

// Manage IPAM Hosts with Next Available Address
resource "infoblox_ipam_host" "ipam_host_with_na_address" {
  uddi = {
    name      = "example_ipam_host_full"
    comment   = "IPAM Hosts Example with Next Available Address"
    addresses = [{ next_available_id = infoblox_network.parent_network.id }]

  }
}