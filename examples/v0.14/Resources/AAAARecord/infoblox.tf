terraform {
  # Required providers block for Terraform v0.14.7
  required_providers {
    infoblox = {
      source  = "terraform-providers/infoblox"
      version = ">= 1.0"
    }
    # Specify the required cloud providers. Here AWS is an example.
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

# Create record using an IPv6 Address
resource "infoblox_aaaa_record" "aaaa_record_1" {
  network_view = "default"
  dns_view="default"            

  ipv6_addr="2000:01"           # "2000::/64" network MUST exist at NIOS
  fqdn="aaaa_record.aws.com"    # "aws.com" zone MUST exist at NIOS
  ttl = 3600

  comment = "AAAA record created"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "CMP Type" = "Terraform"
    "Cloud API Owned" = "True"
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}

# Create record using an IPv6 CIDR
resource "infoblox_aaaa_record" "aaaa_record_2" {
  network_view = "non_default" # non_default network view MUST exist at NIOS
  dns_view="non_default"            # non_default network view MUST exist at NIOS
  cidr = "2000::/64"                # "2000::/64" network MUST exist at NIOS
  fqdn="aaaa_record.aws.com"        # "aws.com" zone MUST exist at NIOS
  ttl = 3600

  comment = "AAAA record created"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "CMP Type" = "Terraform"
    "Cloud API Owned" = "True"
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}

# Create record with default network view and dns view
resource "infoblox_aaaa_record" "aaaa_record_3" {
  ipv6_addr="2000:03"           # "2000::/64" network MUST exist at NIOS
  fqdn="aaaa_record.aws.com"    # "aws.com" zone MUST exist at NIOS
} 