// Associate a pre-existing DHCP on-prem host with a DHCP Config Profile
resource "infoblox_dhcp_host" "basic" {
  id = "dhcp/host/465008"
  uddi = {
    server = "dhcp/server/090985ed-7e7d-11f1-85df-a6a5c5dbbe56"
  }
}

// Associate a second DHCP on-prem host and link it to a specific DHCP Config Profile
resource "infoblox_dhcp_host" "with_server" {
  id = "dhcp/host/471711"
  uddi = {
    server = "dhcp/server/090985ed-7e7d-11f1-85df-a6a5c5dbbe56"
  }
}
