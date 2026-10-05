// Create a Custom Redirect with basic Fields
resource "infoblox_custom_redirect" "basic_custom_redirect" {
  uddi = {
    name = "example_custom_redirect"
    data = "156.2.3.10"
  }
}
