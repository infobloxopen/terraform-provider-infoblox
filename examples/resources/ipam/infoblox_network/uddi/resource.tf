// Create a Network View (Required as Parent)
resource "infoblox_network_view" "example" {
  uddi = {
    name = "example_nw_view"
  }
}

// Create a Network Container (Parent for next-available allocation)
resource "infoblox_network_container" "example" {
  uddi = {
    address = "10.10.0.0"
    cidr    = 16
    space   = infoblox_network_view.example.id
  }
}

// Static address
resource "infoblox_network" "example" {
  uddi = {
    address = "10.0.0.0"
    cidr    = 24
    space   = infoblox_network_view.example.id

    // Other optional fields
    name    = "example_subnet"
    comment = "Subnet for Site A"
    tags = {
      Site = "location-1"
    }
  }
}

// Next available subnet — the address is allocated from the parent address block
resource "infoblox_network" "example_na_s" {
  uddi = {
    cidr  = 24
    space = infoblox_network_view.example.id
    dynamic_allocation = {
      next_available_id = infoblox_network_container.example.id
    }

    // Other optional fields
    name    = "example_subnet"
    comment = "Subnet for Site A"
    tags = {
      Site = "location-1"
    }
  }
}
