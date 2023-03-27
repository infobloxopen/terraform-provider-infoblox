// statically allocated IPv4 network container, minimal set of parameters
resource "infoblox_ipv4_network_container" "nc1" {
  cidr = "10.2.0.0/24"
}

// full set of parameters for statically allocated IPv4 network container
resource "infoblox_ipv4_network_container" "nc2" {
  cidr = "10.2.0.0/24" // we may allocate the same IP address range but in another network view
  network_view = "nondefault_netview"
  comment = "one of our clients"
  ext_attrs = jsonencode({
    "Site" = "remote office"
    "Country" = "Australia"
  })
}

// so far, we do not support dynamic allocation of network containers
