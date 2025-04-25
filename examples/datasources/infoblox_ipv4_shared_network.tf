resource "infoblox_ipv4_shared_network" "shared_network" {
  name = "shared-network1"
  networks = ["37.12.3.0/24"]
}

data "infoblox_ipv4_shared_network" "shared_network_read" {
  filters = {
    "network_view" = "default"
  }
  depends_on = [infoblox_ipv4_shared_network.shared_network]
}

resource "infoblox_ipv4_shared_network" "shared_network2" {
  name = "shared-network14"
  comment = "test ipv4 shared network record"
  networks = ["27.12.3.0/24","27.13.3.0/24"]
  network_view = "default"
  disable = false
  ext_attrs = jsonencode({
    "Site" = "Osaka"
  })
  use_options = true
  options {
    name = "domain-name-servers"
    num = 6
    value = "12.22.33.44"
    use_option = false
  }
}

data "infoblox_ipv4_shared_network" "shared_network_read2" {
  filters = {
    "*Site" = "Osaka"
  }
  depends_on = [infoblox_ipv4_shared_network.shared_network2]
}


