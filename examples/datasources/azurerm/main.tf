provider "infoblox" {
}

resource "infoblox_network" "test" {
  network_name = "test"
  cidr         = "10.0.23.0/24"
  reserve_ip   = 2
  tenant_id    = "default"
}

data "infoblox_network" "test" {
  cidr      = infoblox_network.test.cidr
  tenant_id = "default"
}

output "id" {
  value = data.infoblox_network.test.id
}

output "network_name" {
  value = data.infoblox_network.test.network_name
}
