resource "infoblox_view" "example" {
  uddi = {
    name    = "example"
    comment = "Example IP space created by the terraform provider"
    tags = {
      Site = "location-1"
    }
  }
}

resource "infoblox_range" "example" {
  uddi = {
    start = "192.168.1.15"
    end   = "192.168.1.30"
    space = infoblox_view.example.id

    // Other optional fields
    name    = "example"
    comment = "Example Range created by the terraform provider"
    tags = {
      Site = "location-1"
    }
    exclusion_ranges = [
      {
        start = "192.168.1.17"
        end   = "192.168.1.20"
      }
    ]
  }
}
