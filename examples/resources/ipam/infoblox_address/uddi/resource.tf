// Create a Network View (Required as Parent)
resource "infoblox_network_view" "example" {
  uddi = {
    name = "example_nw_view"
  }
}

// Create a Network Container (Parent for next-available allocation)
resource "infoblox_network_container" "example" {
  uddi = {
    address = "10.0.0.0"
    cidr    = 8
    space   = infoblox_network_view.example.id
  }
}

// Create a Network (Parent for next-available allocation)
resource "infoblox_network" "example" {
  uddi = {
    address = "10.1.0.0"
    cidr    = 24
    space   = infoblox_network_view.example.id
  }
  depends_on = [infoblox_network_container.example]
}

// Static address
resource "infoblox_address" "example" {
  uddi = {
    address = "10.1.0.5"
    space   = infoblox_network_view.example.id

    // Other optional fields
    comment   = "Reservation for Site A"
    hwaddr    = "00:11:22:33:44:55"
    interface = "eth0"
    names = [{
      name = "bby-1"
      type = "user"
    }]
    external_keys = {
      key1 = "value1"
    }
    tags = {
      Site = "location-1"
    }
  }
  depends_on = [infoblox_network.example]
}

// Next available address in a network
resource "infoblox_address" "example_na_network" {
  uddi = {
    space = infoblox_network_view.example.id
    dynamic_allocation = {
      next_available_id = infoblox_network.example.id
    }

    // Other optional fields
    comment = "dynamically allocated from a network"
    tags = {
      Site = "location-2"
    }
  }
}

// Next available address in a network container
resource "infoblox_address" "example_na_network_container" {
  uddi = {
    space = infoblox_network_view.example.id
    dynamic_allocation = {
      next_available_id = infoblox_network_container.example.id
    }
  }
  depends_on = [infoblox_network.example, infoblox_address.example_na_network]
}

// Next available address in a range
// TODO: drop this once infoblox_range is onboarded.
resource "infoblox_address" "example_na_range" {
  uddi = {
    space = "ipam/ip_space/<>"
    dynamic_allocation = {
      next_available_id = "ipam/ip_space/<>"
    }
  }
}
