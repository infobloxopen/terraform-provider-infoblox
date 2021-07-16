terraform {
  # Required providers block for Terraform v0.14.7
  required_providers {
    infoblox = {
      source  = "terraform-providers/infoblox"
      version = ">= 1.0"
    }
  }
}

# Create CNAME record for VM
resource "infoblox_cname_record" "ib_cname_record"{
  dns_view = "default"

  canonical = "CanonicalTestName.xyz.com"
  alias = "AliasTestName.xyz.com"

  ttl = 3600

  comment = "CNAME record created"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "CMP Type" = "Terraform"
    "Cloud API Owned" = "True"
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}
