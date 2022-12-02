// statically allocated IPv6 network container, minimal set of parameters
resource "infoblox_ipv6_network_container" "v6net_c1" {
  cidr = "2002:1f93:0:1::/96"
}

// full set of parameters for statically allocated IPv6 network container
resource "infoblox_ipv6_network_container" "v6net_c2" {
  cidr = "2002:1f93:0:2::/96"
  network_view = "nondefault_netview"
  comment = "new generation network segment"
  ext_attrs = jsonencode({
    "Site" = "space station"
    "Country" = "Earth orbit"
  })
}

// so far, we do not support dynamic allocation of network containers
