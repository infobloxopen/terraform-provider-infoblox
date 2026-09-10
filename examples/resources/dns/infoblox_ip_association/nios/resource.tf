// Create an Auth Zone (Required as Parent)
resource "infoblox_zone_auth" "example" {
  nios = {
    fqdn = "example.com"
  }
}

// Create a Network (Required for Dynamic Allocation Example)
resource "infoblox_network" "example_network" {
  nios = {
    network = "13.0.0.0/24"
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Static IPv4 host record with a MAC address, DHCP left off
resource "infoblox_record_host" "static" {
  nios = {
    name              = "host1.${infoblox_zone_auth.example.nios.fqdn}"
    view              = "default"
    configure_for_dns = true
    ipv4addrs = [
      {
        ipv4addr = "10.101.1.110"
      }
    ]
    ext_attrs = {
      Site = "location-1"
    }
  }
}

resource "infoblox_ip_association" "static" {
  nios = {
    record_host_id     = infoblox_record_host.static.id
    mac                = "aa:bb:cc:11:22:33"
    configure_for_dhcp = false
  }
}

// Dual-stack host record with DHCP enabled, leasing the IPv6 address on its DUID
resource "infoblox_record_host" "dual_stack" {
  nios = {
    name              = "host2.${infoblox_zone_auth.example.nios.fqdn}"
    view              = "default"
    configure_for_dns = true
    ipv4addrs = [
      {
        ipv4addr = "10.101.1.112"
      }
    ]
    ipv6addrs = [
      {
        ipv6addr = "2002:1f93::12:2"
      }
    ]
  }
}

resource "infoblox_ip_association" "dual_stack" {
  nios = {
    record_host_id     = infoblox_record_host.dual_stack.id
    mac                = "aa:bb:cc:11:22:44"
    duid               = "00:03:00:01:aa:bb:cc:11:22:44"
    match_client       = "DUID"
    configure_for_dhcp = true
  }
}

// Host record whose address is allocated from a network, then associated
resource "infoblox_record_host" "dynamic" {
  nios = {
    name              = "host3.${infoblox_zone_auth.example.nios.fqdn}"
    view              = "default"
    configure_for_dns = true
    ipv4addrs = [
      {
        dynamic_allocation = {
          network      = infoblox_network.example_network.nios.network
          network_view = "default"
          exclude      = ["13.10.0.1", "13.10.0.2"]
        }
      }
    ]
  }
}

resource "infoblox_ip_association" "dynamic" {
  nios = {
    record_host_id     = infoblox_record_host.dynamic.id
    mac                = "aa:bb:cc:11:22:55"
    configure_for_dhcp = true
  }
}
