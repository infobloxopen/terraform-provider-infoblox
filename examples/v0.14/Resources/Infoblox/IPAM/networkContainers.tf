terraform {
  # Required providers block for Terraform v0.14.7
  required_providers {
    infoblox = {
      source = "infobloxopen/infoblox"
      version = ">=2.0"
    }
  }
}

# Create a network container in Infoblox Grid
resource "infoblox_ipv4_network_container" "IPv4_nw_c" {
  network_view="default"
  cidr = "10.0.0.0/16"
  comment = "tf IPv4 network container"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}

resource "infoblox_ipv6_network_container" "IPv6_nw_c" {
  network_view="default"
  cidr = "2000::0/54"
  comment = "tf IPv6 network container"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}