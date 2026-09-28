// Create a Network View ( Required as Parent )
resource "infoblox_network_view" "example" {
  uddi = {
    name    = "example"
    comment = "Example IP space created by the terraform provider"
    tags = {
      Site = "location-1"
    }
  }
}

// Create a Network ( Required as Parent )
resource "infoblox_network" "test" {
  uddi = {
    address = "10.0.0.0"
    cidr    = 24
    space   = infoblox_network_view.test.id
  }
}

// Create Fixed Address with Basic Fields
resource "infoblox_fixed_address" "example_fixed_address" {
  uddi = {
    name        = "example_fixed_address"
    address     = "10.0.0.222"
    ip_space    = infoblox_network_view.example.id
    match_type  = "mac"
    match_value = "aa:bb:cc:dd:ee:ff"
    comment     = "Example Fixed Address created by the terraform provider"
    tags = {
      Site = "location-1"
    }
  }
  depends_on = [infoblox_network.test.id]
}

// Create Fixed Address using Next available IP
resource "infoblox_fixed_address" "example_fixed_address_na" {
  uddi = {
    name               = "example_fixed_address2"
    ip_space           = infoblox_network_view.example.id
    dynamic_allocation = { next_available_id = infoblox_network.test.id }
    match_type         = "mac"
    match_value        = "00:00:00:00:00:01"
    comment            = "Example Fixed Address created by the terraform provider"
    tags = {
      Site = "location-1"
    }
  }
}
