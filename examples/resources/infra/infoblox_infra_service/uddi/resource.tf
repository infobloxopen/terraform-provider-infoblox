// Manage an Infra Host (Required as Parent)
resource "infoblox_infra_host" "example" {
  uddi = {
    display_name  = "example-host"
    serial_number = "1234"
    tags = {
      "host/serial_number" = "1234"
    }
  }
}

// Create an Infra Service with basic fields
resource "infoblox_infra_service" "example" {
  uddi = {
    name         = "example-infra-service"
    pool_id      = infoblox_infra_host.example.uddi.pool_id
    service_type = "dns"
  }
}

// Create an Infra Service with additional fields
resource "infoblox_infra_service" "example_with_options" {
  uddi = {
    name            = "example-dhcp-service"
    pool_id         = infoblox_infra_host.example.uddi.pool_id
    service_type    = "dhcp"
    description     = "DHCP service created by Terraform"
    desired_state   = "start"
    desired_version = "3.5.0"
    tags = {
      Site = "location-1"
    }
  }
}
