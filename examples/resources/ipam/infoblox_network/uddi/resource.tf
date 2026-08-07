// Static address
resource "infoblox_network" "example" {
  uddi = {
    address = "10.0.0.0"
    cidr    = 24

    // Other optional fields
    name    = "example_subnet"
    comment = "Subnet for Site A"
    tags = {
      Site = "location-1"
    }
  }
}

// Next available subnet
resource "infoblox_network" "example_na_s" {
  uddi = {
    cidr = 24

    // Other optional fields
    name    = "example_subnet"
    comment = "Subnet for Site A"
    tags = {
      Site = "location-1"
    }
  }
}
