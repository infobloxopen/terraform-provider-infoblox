terraform {
  # Required providers block for Terraform v0.14.7
  required_providers {
    infoblox = {
      source  = "terraform-providers/infoblox"
      version = ">= 1.0"
    }
  }
}

resource "infoblox_ipv4_network" "ipv4_network"{
  network_view = "default"
  cidr = "10.0.0.0/24"

  comment = "tf IPv4 network updated"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Network Name" = "TestDataSource"
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}

data "infoblox_ipv4_network" "test" {
  network_view = "default"
  cidr = infoblox_ipv4_network.ipv4_network.cidr
}

output "id" {
  value = data.infoblox_ipv4_network.test
}

output "comment" {
  value = data.infoblox_ipv4_network.test.comment
}

output "ext_attrs" {
  value = data.infoblox_ipv4_network.test.ext_attrs
}
