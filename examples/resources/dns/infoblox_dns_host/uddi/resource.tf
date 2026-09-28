resource "infoblox_dns_host" "example_dns_host" {
  uddi = {

    // Other Optional fields
    tags = {
      Site = "location-1"
    }
  }
}
