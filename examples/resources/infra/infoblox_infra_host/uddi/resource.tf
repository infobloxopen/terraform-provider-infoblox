resource "infoblox_view" "current" {
  uddi = {
    name = "example-space"
  }
}

resource "infoblox_infra_host" "example" {
  uddi = {
    display_name = "example_host"

    // Other Optional fields
    description = "An example host"
    ip_space    = infoblox_view.current.id
    tags = {
      Site = "location-1"
    }
  }
}
