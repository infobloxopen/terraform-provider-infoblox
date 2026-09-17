resource "infoblox_custom_redirect" "example" {
  uddi = {
    name = "example_custom_redirect"
    data = "156.2.3.10"
  }
}
