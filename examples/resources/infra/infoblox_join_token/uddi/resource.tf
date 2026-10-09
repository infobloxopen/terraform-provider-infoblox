resource "infoblox_join_token" "example" {
  uddi = {
    name = "example_join_token"

    // Other optional fields
    description = "Join token for Site A"
    tags = {
      Site = "location-1"
    }
  }
}
