data "infoblox_ipv4_network_container" "nc2" {
  filters = {
    network      = "10.2.0.0/24"
    network_view = "nondefault_netview"
  }

  depends_on = [infoblox_ipv4_network_container.nc2]
}


// Search by extensible is available, just specify the EA key starting with asterisk in filters
data "infoblox_ipv4_network_container" "nc_ea_search" {
  filters = {
    "*Building" = "Cali"
  }
}
