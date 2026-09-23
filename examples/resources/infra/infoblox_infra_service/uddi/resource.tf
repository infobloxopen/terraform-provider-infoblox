// Create a basic Infra Service
// Note: replace pool_id with a real InfraPool resource identifier (infra/pool/<id>)
resource "infoblox_infra_service" "example" {
  uddi = {
    name         = "example-infra-service"
    pool_id      = "infra/pool/<pool-id>"
    service_type = "dns"
  }
}

// Create an Infra Service with additional fields
resource "infoblox_infra_service" "example_with_options" {
  uddi = {
    name            = "example-dhcp-service"
    pool_id         = "infra/pool/<pool-id>"
    service_type    = "dhcp"
    description     = "DHCP service created by Terraform"
    desired_state   = "start"
    desired_version = "3.5.0"
    tags = {
      Site = "location-1"
    }
  }
}
