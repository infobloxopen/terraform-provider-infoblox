data "infoblox_ipv4_network_container" "net_container1" {
  cidr = "10.2.0.0/24"
  network_view = "nondefault_netview" // required, even if it is the same as the default value
}
