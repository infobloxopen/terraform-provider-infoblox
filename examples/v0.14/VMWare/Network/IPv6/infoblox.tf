# Create an IPv6 Network in Infoblox Grid when CIDR is passed
resource "infoblox_ipv6_network" "ipv6_network"{
  network_view = "default"

  parent_cidr = infoblox_ipv6_network_container.IPv6_nw_c.cidr
  allocate_prefix_len = 64
  reserve_ipv6 = 3

  comment = "tf IPv6 network"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Network Name" = "ipv6-tf-network"
    Location = "Test loc."
    Site = "Test site"
  })
}
