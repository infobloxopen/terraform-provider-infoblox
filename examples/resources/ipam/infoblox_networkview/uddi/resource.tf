resource "infoblox_networkview" "example" {
  uddi = {
    name = "example_ip_space"
  }
}

resource "infoblox_networkview" "example_tags" {
  uddi = {
    name    = "example_ip_space_tags"
    comment = "Example IP space with tags created by the terraform provider"
    tags = {
      Site = "location-1"
    }
  }
}
