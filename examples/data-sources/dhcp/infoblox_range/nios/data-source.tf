// Create a Network View (Required as Parent)
resource "infoblox_network_view" "example_network_view" {
  nios = {
    name = "example-range-view"
  }
}

// Create an IPv4 Network inside the Network View (Required as Parent of the Range)
resource "infoblox_network" "example_network" {
  nios = {
    network      = "10.0.0.0/24"
    network_view = infoblox_network_view.example_network_view.nios.name
  }
}

// Create a DHCP Range inside the Network
resource "infoblox_range" "example_range" {
  nios = {
    start_addr   = "10.0.0.170"
    end_addr     = "10.0.0.180"
    network      = infoblox_network.example_network.nios.network
    network_view = infoblox_network.example_network.nios.network_view
    comment      = "Example Range created by the terraform provider"
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Retrieve a specific DHCP range by filters
data "infoblox_range" "get_range_with_filter" {
  filters = {
    start_addr   = infoblox_range.example_range.nios.start_addr
    network_view = infoblox_range.example_range.nios.network_view
  }
}

// Retrieve specific DHCP ranges using Extensible Attributes
data "infoblox_range" "get_range_with_extattr_filter" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all DHCP ranges
data "infoblox_range" "get_all_ranges" {}
