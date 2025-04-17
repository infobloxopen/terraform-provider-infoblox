// shared network with minimum set of parameters
resource "infoblox_ipv4_shared_network" "shared_network_min_parameters" {
  name = "shared-network1"
  networks = ["37.12.3.0/24"]
}

// shared network with full set of parameters
resource "infoblox_ipv4_shared_network" "shared_network_full_parameters" {
  name = "shared-network2"
  comment = "test ipv4 shared network record"
  networks = ["31.12.3.0/24","31.13.3.0/24"]
  network_view = "view2"
  disable = false
  ext_attrs = jsonencode({
    "Site" = "Tokyo"
  })
  use_options = false
  options {
    name = "domain-name-servers"
    value = "11.22.33.44"
    vendor_class = "DHCP"
    num = 6
    use_option = true
  }
}
