# Create an IPv4 network container in Infoblox Grid
resource "infoblox_ipv4_network_container" "IPv4_nw_c" {
  network_view="default"

  cidr = "10.0.0.0/16"
  comment = "tf IPv4 network container"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    Location = "Test loc."
    Site = "Test site"
  })
}

