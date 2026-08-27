// Associate a pre-existing DHCP on-prem host with a DHCP Config Profile
resource "infoblox_dhcp_host" "basic" {
  id = "dhcp/host/1415016"
  uddi = {
    server = "dhcp/server/4489b6fb-9b56-11ef-9ea8-b29a8ad0d125"
  }
}

// Associate a DHCP on-prem host and link it to a specific DHCP Config Profile
resource "infoblox_dhcp_host" "with_server" {
  id = "dhcp/host/1390921"
  uddi = {
    server = "dhcp/server/4489b6fb-9b56-11ef-9ea8-b29a8ad0d125"
  }
}
