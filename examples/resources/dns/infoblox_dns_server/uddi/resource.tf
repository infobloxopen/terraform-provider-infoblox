resource "infoblox_dns_server" "example_server" {
  uddi = {
    name = "example_dns_server"

    // Other Optional fields
    comment = "An example server"
    tags = {
      Site = "location-1"
    }
  }
}
