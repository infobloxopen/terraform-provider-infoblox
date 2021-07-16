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