data "infoblox_ipv4_network" "net1" {
  filters = {
    network = "10.1.0.0/24"
    network_view = "nondefault_netview"
  }

  depends_on = [infoblox_ipv4_network.net1]
}


# Search by extensible is available, just specify the EA key starting with asterisk in filters
data "infoblox_ipv4_network" "net1" {
  filters = {
    "*Building" = "Cali"
  }
}
