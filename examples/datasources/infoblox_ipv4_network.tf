data "infoblox_ipv4_network" "net1" {
  cidr = "10.1.0.0/24"
  network_view = "nondefault_netview"

  depends_on = [infoblox_ipv4_network.net1]
}
