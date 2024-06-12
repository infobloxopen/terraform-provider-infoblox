resource "infoblox_ipv6_network_container" "nc1" {
  cidr = "2002:1f93:0:2::/96"
  comment = "new generation network segment"
  ext_attrs = jsonencode({
    "Site" = "space station"
  })
}

data "infoblox_ipv6_network_container" "nc2" {
  filters = {
    network = "2002:1f93:0:2::/96"
  }

  depends_on = [infoblox_ipv6_network_container.nc1]
}

data "infoblox_ipv6_network_container" "nc_ea_search" {
  filters = {
    "*Site" = "space station"
  }
  depends_on = [infoblox_ipv6_network_container.nc1]
}

output "nc1_output" {
  value = data.infoblox_ipv6_network_container.nc2
}

output "nc1_comment" {
  value = data.infoblox_ipv6_network_container.nc_ea_search
}