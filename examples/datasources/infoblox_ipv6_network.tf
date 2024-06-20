resource "infoblox_ipv6_network" "ipv6net1" {
  cidr = "2002:1f93:0:4::/96"
  reserve_ipv6 = 10
  gateway = "2002:1f93:0:4::1"
  comment = "let's try IPv6"
  ext_attrs = jsonencode({
    "Site" = "Antarctica"
  })
}

data "infoblox_ipv6_network" "readNet1" {
  filters = {
    network = "2002:1f93:0:4::/96"
  }
  depends_on = [infoblox_ipv6_network.ipv6net1]
}

data "infoblox_ipv6_network" "readnet2" {
  filters = {
    "*Site" = "Antarctica"
  }
  depends_on = [infoblox_ipv6_network.ipv6net1]
}

output "ipv6net_res" {
  value = data.infoblox_ipv6_network.readNet1
}

output "ipv6net_res1" {
  value = data.infoblox_ipv6_network.readnet2
}