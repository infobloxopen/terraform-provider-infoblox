// statically allocated IPv6 network, minimal set of parameters
resource "infoblox_ipv6_network" "net1" {
  cidr = "2002:1f93:0:3::/96"
}

// full set of parameters for statically allocated IPv6 network
resource "infoblox_ipv6_network" "net2" {
  cidr = "2002:1f93:0:4::/96"
  network_view = "nondefault_netview"
  reserve_ipv6 = 10
  gateway = "2002:1f93:0:4::1"
  comment = "let's try IPv6"
  ext_attrs = jsonencode({
    "Site" = "somewhere in Antarctica"
  })
}

// full set of parameters for dynamically allocated IPv6 network
resource "infoblox_ipv6_network" "net3" {
  parent_cidr = infoblox_ipv6_network_container.nc1.cidr // reference to the resource from another example
  allocate_prefix_len = 100 // 96 (existing network container) + 4 (new network), prefix
  network_view = "default" // we may omit this but it is not a mistake to specify explicitly
  reserve_ipv6 = 20
  gateway = "none" // no gateway defined for this network
  comment = "the network for the Test Lab"
  ext_attrs = jsonencode({
    "Site" = "small inner cluster"
  })
}
