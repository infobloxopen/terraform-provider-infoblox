data "infoblox_ipv4_network_container" "nc2" {
  cidr = "10.2.0.0/24"
  network_view = "nondefault_netview"

  depends_on = [infoblox_ipv4_network_container.nc2]
}
