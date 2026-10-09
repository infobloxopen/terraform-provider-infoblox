// Create an Auth Zone (Required as Parent)
resource "infoblox_zone_auth" "example" {
  nios = {
    fqdn = "example.com"
  }
}

// Create a Network (Required for Dynamic Allocation Examples)
resource "infoblox_network" "example_network" {
  nios = {
    network = "13.0.0.0/24"
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create a Host Record with a static IPv4 address
resource "infoblox_record_host" "example_static" {
  nios = {
    name    = "host-1.${infoblox_zone_auth.example.nios.fqdn}"
    view    = "default"
    comment = "This is a test host record"
    ipv4addrs = [
      {
        ipv4addr = "10.0.0.18"
      }
    ]
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create a dual-stack Host Record
resource "infoblox_record_host" "example_dual_stack" {
  nios = {
    name              = "host-2.${infoblox_zone_auth.example.nios.fqdn}"
    view              = "default"
    configure_for_dns = true
    ipv4addrs = [
      {
        ipv4addr = "10.0.0.19"
      }
    ]
    ipv6addrs = [
      {
        ipv6addr = "2002:1f93::12:2"
      }
    ]
  }
}

// Create a Host Record using dynamic allocation
resource "infoblox_record_host" "example_dynamic" {
  nios = {
    name = "host-3.${infoblox_zone_auth.example.nios.fqdn}"
    view = "default"
    ipv4addrs = [
      {
        dynamic_allocation = {
          network      = infoblox_network.example_network.nios.network
          network_view = "default"
          exclude      = ["13.0.0.1", "13.0.0.2"]
        }
      }
    ]
  }
}
