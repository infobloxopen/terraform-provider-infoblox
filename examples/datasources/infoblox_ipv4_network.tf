data "infoblox_ipv4_network" "net1" {
  cidr = "10.1.0.0/24"
  network_view = "nondefault_netview" // required, even if it is the same as the default value
}
