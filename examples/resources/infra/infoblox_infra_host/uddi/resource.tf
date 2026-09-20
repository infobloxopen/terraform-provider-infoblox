// Create a Network View ( Required as Parent )
resource "infoblox_network_view" "parent_space" {
  uddi = {
    name = "example-space"
  }
}

resource "infoblox_infra_host" "example" {
  uddi = {
    display_name = "example_host"

    // Other Optional fields
    description   = "An example host"
    serial_number = "1234"
    ip_space      = infoblox_network_view.parent_space.id
    tags = {
      Site                 = "location-1"
      "host/serial_number" = "1234"
    }
  }
}
