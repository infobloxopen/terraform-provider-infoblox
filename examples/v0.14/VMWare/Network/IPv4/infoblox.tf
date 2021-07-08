# Create a network in Infoblox Griby passing the CIDR
resource "infoblox_ipv4_network" "ipv4_network"{
  network_view = "default"

  parent_cidr = infoblox_ipv4_network_container.IPv4_nw_c.cidr
  allocate_prefix_len = 24
  reserve_ip = 2

  comment = "tf IPv4 network"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Network Name" = "ipv4-tf-network"
    Location = "Test loc."
    Site = "Test site"
  })
}

