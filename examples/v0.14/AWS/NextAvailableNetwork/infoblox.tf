terraform {
  # Required providers block for Terraform v0.14.7
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
    infoblox = {
      source  = "terraform-providers/infoblox"
      version = ">= 1.0"
    }
  }
}

# Create a network container in Infoblox Grid
resource "infoblox_ipv4_network_container" "IPv4_nw_c" {
  network_view="default"

  cidr = aws_vpc.vpc.cidr_block
  comment = "tf IPv4 network container"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}

resource "infoblox_ipv6_network_container" "IPv6_nw_c" {
  network_view="default"

  cidr = aws_vpc.vpc.ipv6_cidr_block
  comment = "tf IPv6 network container"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}


# Allocate a network in Infoblox Grid under provided parent CIDR
resource "infoblox_ipv4_network" "ipv4_network"{
  network_view = "default"

  parent_cidr = infoblox_ipv4_network_container.IPv4_nw_c.cidr
  allocate_prefix_len = 24
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

  parent_cidr = infoblox_ipv6_network_container.IPv6_nw_c.cidr
  allocate_prefix_len = 64
  reserve_ipv6 = 3

  comment = "tf IPv6 network"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Network Name" = "ipv6-tf-network"
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}