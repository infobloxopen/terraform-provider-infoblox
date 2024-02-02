// statically allocated IPv6 network container, minimal set of parameters
resource "infoblox_ipv6_network_container" "nc1" {
  cidr = "2002:1f93:0:1::/96"
}

// full set of parameters for statically allocated IPv6 network container
resource "infoblox_ipv6_network_container" "nc2" {
  cidr = "2002:1f93:0:2::/96"
  network_view = "nondefault_netview"
  comment = "new generation network segment"
  ext_attrs = jsonencode({
    "Site" = "space station"
    "Country" = "Earth orbit"
  })
}

// full set of parameters for dynamic allocation of network containers
resource "infoblox_ipv6_network_container" "ncv6" {
  parent_cidr = infoblox_ipv6_network_container.nc2.cidr
  allocate_prefix_len = 97
  network_view = "default"
  comment = "dynamic allocation of network container"
  ext_attrs = jsonencode({
    "Tenant ID" = "terraform_test_tenant"
    Site = "Test site"
  })
}
