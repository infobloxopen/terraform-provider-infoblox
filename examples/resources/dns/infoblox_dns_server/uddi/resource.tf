# Manage a UDDI DNS Server
resource "infoblox_dns_server" "example" {
  uddi = {
    name = "example_server"

    // Other Optional fields
    comment = "An example server"
    tags = {
      Site = "location-1"
    }
    custom_root_ns = [{ address = "192.168.10.10", fqdn = "tf-example.com." }]
    ecs_enabled    = true
    ecs_zones = [
      {
        access = "allow"
        fqdn   = "tf-infoblox.com."
      }
    ]
  }
}
