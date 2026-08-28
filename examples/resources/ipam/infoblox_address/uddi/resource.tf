// Create a Network View (Required as Parent)
resource "infoblox_network_view" "example" {
  uddi = {
    name = "example_ip_space"
  }
}

resource "infoblox_network_container" "example" {
  uddi = {
    address = "10.0.0.0"
    cidr    = 8
    space   = infoblox_network_view.example.id
  }
}

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
    comment = "reservation for Site A"
    hwaddr  = "00:11:22:33:44:55"
    names = [{
      name = "bby-1"
      type = "user"
    }]
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
  depends_on = [infoblox_network.example]
}

// Next available address in a range
// TODO: drop this once infoblox_range is onboarded.
resource "infoblox_address" "example_na_range" {
  uddi = {
    space = "ipam/ip_space/00c2c546-6ad1-11f1-88c3-1e3cda826891"
    dynamic_allocation = {
      next_available_id = "ipam/range/8e6ec141-a2d7-11f1-829e-02fb57fee572"
    }
  }
}
