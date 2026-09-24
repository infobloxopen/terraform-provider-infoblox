// Create IPv6 Networks (Required as Parents)
resource "infoblox_ipv6_network" "parent_network1" {
  nios = {
    network      = "2001:db8:abcd:1::/64"
    network_view = "default"
    comment      = "Parent network for IPv6 shared network 1"
  }
}

resource "infoblox_ipv6_network" "parent_network2" {
  nios = {
    network      = "2001:db8:abcd:2::/64"
    network_view = "default"
    comment      = "Parent network for IPv6 shared network 1"
  }
}

resource "infoblox_ipv6_network" "parent_network3" {
  nios = {
    network      = "2001:db8:abcd:3::/64"
    network_view = "default"
    comment      = "Parent network for IPv6 shared network 2"
  }
}

resource "infoblox_ipv6_network" "parent_network4" {
  nios = {
    network      = "2001:db8:abcd:4::/64"
    network_view = "default"
    comment      = "Parent network for IPv6 shared network 2"
  }
}

// Create an IPv6 Shared Network with Basic Fields
resource "infoblox_ipv6_sharednetwork" "example_ipv6_sharednetwork_basic_fields" {
  nios = {
    name = "example_ipv6_shared_network1"
    networks = [
      infoblox_ipv6_network.parent_network1.id,
      infoblox_ipv6_network.parent_network2.id
    ]
    comment = "Created by Terraform"
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create an IPv6 Shared Network with Additional Fields
resource "infoblox_ipv6_sharednetwork" "example_ipv6_sharednetwork_additional_fields" {
  nios = {
    name = "example_ipv6_shared_network2"
    networks = [
      infoblox_ipv6_network.parent_network3.id,
      infoblox_ipv6_network.parent_network4.id
    ]
    comment = "IPv6 shared network with additional fields"

    enable_ddns                = true
    ddns_domainname            = "ddns.example.com"
    ddns_generate_hostname     = true
    ddns_server_always_updates = true
    ddns_use_option81          = true
    ddns_ttl                   = 3600

    preferred_lifetime          = 27000
    valid_lifetime              = 43200
    update_dns_on_lease_renewal = true

    domain_name_servers = ["2001:db8::53", "2001:db8::54"]
    options = [
      {
        name  = "domain-name"
        num   = 15
        value = "example.com"
      },
      {
        name         = "dhcp6.subscriber-id"
        value        = "subscriber-id"
        vendor_class = "DHCPv6"
      }
    ]

    disable = false
    ext_attrs = {
      Site = "location-1"
    }
  }
}
