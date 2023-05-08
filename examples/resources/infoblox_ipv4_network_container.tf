// statically allocated IPv4 network container, minimal set of parameters
resource "infoblox_ipv4_network_container" "v4net_c1" {
  cidr = "10.2.0.0/24"
}

// full set of parameters for statically allocated IPv4 network container
resource "infoblox_ipv4_network_container" "v4net_c2" {
  cidr = "10.2.0.0/24" // we may allocate the same IP address range but in another network view
  network_view = "nondefault_netview"
  comment = "one of our clients"
  ext_attrs = jsonencode({
    "Site" = "remote office"
    "Country" = "Australia"
  })
}

// full set of parameters for dynamic allocation of network containers
resource "infoblox_ipv4_network_container" "nc3" {
  parent_cidr = "25.0.0.0/24" // cidr must exists in the grid
  allocate_prefix_len = 26
  network_view = "nondefault_netview"
  comment = "one of our clients"
  ext_attrs = jsonencode({
    "Site" = "remote office"
    "Country" = "Australia"
  })
}