terraform {
  # Required providers block for Terraform v0.14.7
  required_providers {
    infoblox = {
      source = "infobloxopen/infoblox"
      version = ">=2.0"
    }
  }
}

# Allocate a network in Infoblox Grid under provided parent CIDR
resource "infoblox_ipv4_network" "ipv4_network"{
  network_view = "default"

  cidr = "10.0.0.0/24"
  reserve_ip = 2

  comment = "tf IPv4 network"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Network Name" = "ipv4-tf-network"
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}

resource "infoblox_ipv6_network" "ipv6_network"{
  network_view = "default"

  cidr = "2000::0/64"
  reserve_ipv6 = 3

  comment = "tf IPv6 network"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Network Name" = "ipv6-tf-network"
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}